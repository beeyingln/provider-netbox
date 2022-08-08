/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProviderObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProviderParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`
}

// ProviderSpec defines the desired state of Provider
type ProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProviderParameters `json:"forProvider"`
}

// ProviderStatus defines the observed state of Provider.
type ProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Provider is the Schema for the Providers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netboxjet}
type Provider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProviderSpec   `json:"spec"`
	Status            ProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderList contains a list of Providers
type ProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Provider `json:"items"`
}

// Repository type metadata.
var (
	Provider_Kind             = "Provider"
	Provider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Provider_Kind}.String()
	Provider_KindAPIVersion   = Provider_Kind + "." + CRDGroupVersion.String()
	Provider_GroupVersionKind = CRDGroupVersion.WithKind(Provider_Kind)
)

func init() {
	SchemeBuilder.Register(&Provider{}, &ProviderList{})
}
