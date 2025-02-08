package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GitConfigSpec defines the desired state of GitConfig
type GitConfigSpec struct {
	// Repository URL of the Git repository
	RepoURL string `json:"repoURL"`

	// +optional
	// Secret containing the credentials to access the Git repository
	GitAuthSecret string `json:"gitAuthSecret"`
}

// GitConfigStatus defines the observed state of GitConfig

type GitConfigStatus struct {

	// Health indicates the health of the GitConfig
	Health string `json:"health,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitConfig is the Schema for the gitconfigs API
type GitConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitConfigSpec   `json:"spec,omitempty"`
	Status GitConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitConfigList contains a list of GitConfig
type GitConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitConfig `json:"items"`
}
