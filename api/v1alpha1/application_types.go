package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {

	// Name of the application
	Name string `json:"name"`

	GitConfigRef GitConfigRef `json:"gitConfigRef"`

	// +optional
	ClusterConfigRef string `json:"clusterConfigRef,omitempty"`
}

type GitConfigRef struct {

	// Name of the GitConfig
	Name string `json:"name"`

	// Revision of the Git repository
	Revision string `json:"revision"`
}

type ClusterConfigRef struct {

	// Name of the ClusterConfig
	Name string `json:"name"`
}

// +k8s:deepcopy-gen=true

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// Phase represents the current phase of the application
	// +optional
	Phase string `json:"phase,omitempty"`

	// Conditions represent the latest available observations of the application's state
	// +optional
	Conditions []ApplicationCondition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen=true

// ApplicationCondition describes the state of an application at a certain point
type ApplicationCondition struct {
	// Type of application condition
	Type string `json:"type"`

	// Status of the condition, one of True, False, Unknown
	Status string `json:"status"`

	// Last time the condition transit from one status to another
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// Reason for the condition's last transition
	Reason string `json:"reason,omitempty"`

	// Message associated with the condition
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}
