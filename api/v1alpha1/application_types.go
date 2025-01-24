package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// Source defines the source of the application
	Source ApplicationSource `json:"source"`

	// Destination defines the target cluster and namespace
	Destination ApplicationDestination `json:"destination"`

	// Project is the project name this application belongs to
	// +optional
	Project string `json:"project,omitempty"`

	// SyncPolicy controls when and how a sync will be performed
	// +optional
	SyncPolicy *SyncPolicy `json:"syncPolicy,omitempty"`
}

// ApplicationSource contains information about the application's source
type ApplicationSource struct {
	// RepoURL is the URL to the repository (Git or Helm) that contains the application manifests
	RepoURL string `json:"repoURL"`

	// Path is a directory path within the Git repository
	Path string `json:"path"`

	// TargetRevision defines the revision of the source to sync to
	// +optional
	TargetRevision string `json:"targetRevision,omitempty"`
}

// ApplicationDestination contains deployment destination information
type ApplicationDestination struct {
	// Server specifies the URL of the target cluster's Kubernetes API server
	Server string `json:"server"`

	// Namespace specifies the target namespace for the application's resources
	Namespace string `json:"namespace"`
}

// SyncPolicy controls when and how a sync will be performed
type SyncPolicy struct {
	// Automated defines if an application should be synced automatically
	// +optional
	Automated *SyncPolicyAutomated `json:"automated,omitempty"`
}

// SyncPolicyAutomated defines the policy for automated syncs
type SyncPolicyAutomated struct {
	// Prune specifies whether to delete resources from the cluster that are not found in Git
	// +optional
	Prune bool `json:"prune,omitempty"`

	// SelfHeal specifies whether to revert resources back to their desired state upon modification
	// +optional
	SelfHeal bool `json:"selfHeal,omitempty"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// Conditions is a list of currently observed application conditions
	Conditions []ApplicationCondition `json:"conditions,omitempty"`

	// Health contains information about the application's current health status
	Health HealthStatus `json:"health,omitempty"`

	// Sync contains information about the application's current sync status
	Sync SyncStatus `json:"sync,omitempty"`
}

// ApplicationCondition contains details about current application condition
type ApplicationCondition struct {
	// Type is an application condition type
	Type string `json:"type"`

	// Message contains human-readable message indicating details about condition
	Message string `json:"message"`

	// LastTransitionTime is the time the condition was last observed
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
}

// HealthStatus contains information about the currently observed health state
type HealthStatus struct {
	// Status holds the status code of the application or resource
	Status string `json:"status"`

	// Message is a human-readable informational message describing the health status
	Message string `json:"message,omitempty"`
}

// SyncStatus contains information about the currently observed sync state
type SyncStatus struct {
	// Status is the sync state of the application
	Status string `json:"status"`

	// Revision contains information about the revision that was last synced
	Revision string `json:"revision,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Sync Status",type="string",JSONPath=".status.sync.status"
//+kubebuilder:printcolumn:name="Health",type="string",JSONPath=".status.health.status"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
