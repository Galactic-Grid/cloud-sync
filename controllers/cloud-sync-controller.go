package controller

import (
	"context"
	"time"

	"github.com/Galactic-Grid/cloud-sync/api/v1alpha1"
	"github.com/go-logr/logr"
	"golang.org/x/time/rate"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Application Reconciler reconciles a Application object
type ApplicationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

// Reconcile implements reconcile.Reconciler.
func (r *ApplicationReconciler) Reconcile(context.Context, reconcile.Request) (reconcile.Result, error) {
	panic("unimplemented")
}

// SetupWithManager sets up the controller with the Manager.
func (r *ApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// Create a new rate limiter that combines exponential backoff with a maximum rate
	rateLimiter := workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(5*time.Millisecond, 30*time.Second),
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)

	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(controller.Options{
			RateLimiter:             rateLimiter,
			MaxConcurrentReconciles: 5,
		}).
		For(&v1alpha1.Application{}).
		Complete(r)
}
