package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/go-logr/zapr"
	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
	"gopkg.in/alecthomas/kingpin.v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

func main() {
	var (
		app              = kingpin.New("Cloud-Sync", "Cloud-Sync for Deployments.")
		debug            = app.Flag("debug", "Enable debug mode").Default("false").Bool()
		maxReconcileRate = app.Flag("max-reconcile-rate", "The maximum number of concurrent reconciliation operations.").Default("10").Int()
		namespace        = app.Flag("namespace", "Namespace used to set as default scope in default secret store config.").Default("cloudsync").OverrideDefaultFromEnvar("POD_NAMESPACE").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	var zl *zap.Logger
	var err error
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = iso8601TimeEncoder

	if *debug {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		zl, err = config.Build()
	} else {
		zl, err = config.Build()
	}
	if err != nil {
		panic(err)
	}

	log := zapr.NewLogger(zl.Named("cloud-sync"))
	// SetLogger is required starting in controller-runtime 0.15.0.
	// https://github.com/kubernetes-sigs/controller-runtime/pull/2317
	ctrl.SetLogger(log)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(ratelimiter.LimitRESTConfig(cfg, *maxReconcileRate), ctrl.Options{

		// controller-runtime uses both ConfigMaps and Leases for leader
		// election by default. Leases expire after 15 seconds, with a
		// 10 second renewal deadline. We've observed leader loss due to
		// renewal deadlines being exceeded when under high load - i.e.
		// hundreds of reconciles per second and ~200rps to the API
		// server. Switching to Leases only and longer leases appears to
		// alleviate this.
		LeaderElection: false,
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")

	kingpin.FatalIfError(apis.AddToScheme(mgr.GetScheme()), "Cannot add terraform APIs to scheme")

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()

	metrics.Registry.MustRegister(metricRecorder)
	metrics.Registry.MustRegister(stateMetrics)

	o := controller.Options{
		Logger:                  log,
		MaxConcurrentReconciles: *maxReconcileRate,
		PollInterval:            *pollInterval,
		GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
		Features:                &feature.Flags{},
		MetricOptions: &controller.MetricOptions{
			PollStateMetricInterval: *pollStateMetricInterval,
			MRMetrics:               metricRecorder,
			MRStateMetrics:          stateMetrics,
		},
	}

	if *enableManagementPolicies {
		o.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	if *enableExternalSecretStores {
		o.Features.Enable(features.EnableAlphaExternalSecretStores)
		log.Info("Alpha feature enabled", "flag", features.EnableAlphaExternalSecretStores)
		o.ESSOptions = &controller.ESSOptions{}
		if *essTLSCertsPath != "" {
			log.Info("ESS TLS certificates path is set. Loading mTLS configuration.")
			tCfg, err := certificates.LoadMTLSConfig(filepath.Join(*essTLSCertsPath, "ca.crt"), filepath.Join(*essTLSCertsPath, "tls.crt"), filepath.Join(*essTLSCertsPath, "tls.key"), false)
			kingpin.FatalIfError(err, "Cannot load ESS TLS config.")

			o.ESSOptions.TLSConfig = tCfg
		}

		// Ensure default store config exists.
		kingpin.FatalIfError(resource.Ignore(kerrors.IsAlreadyExists, mgr.GetClient().Create(context.Background(), &v1beta1.StoreConfig{
			ObjectMeta: metav1.ObjectMeta{
				Name: "default",
			},
			Spec: v1beta1.StoreConfigSpec{
				// NOTE(turkenh): We only set required spec and expect optional
				// ones to properly be initialized with CRD level default values.
				SecretStoreConfig: xpv1.SecretStoreConfig{
					DefaultScope: *namespace,
				},
			},
		})), "cannot create default store config")
	}

	kingpin.FatalIfError(workspace.Setup(mgr, o, *timeout, *pollJitter), "Cannot setup application controllers")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

func iso8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(time.RFC3339)) // or use a custom ISO 8601 format
}
