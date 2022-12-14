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

type ManufacturerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManufacturerParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`
}

// ManufacturerSpec defines the desired state of Manufacturer
type ManufacturerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManufacturerParameters `json:"forProvider"`
}

// ManufacturerStatus defines the observed state of Manufacturer.
type ManufacturerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManufacturerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Manufacturer is the Schema for the Manufacturers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netboxjet}
type Manufacturer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManufacturerSpec   `json:"spec"`
	Status            ManufacturerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManufacturerList contains a list of Manufacturers
type ManufacturerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Manufacturer `json:"items"`
}

// Repository type metadata.
var (
	Manufacturer_Kind             = "Manufacturer"
	Manufacturer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Manufacturer_Kind}.String()
	Manufacturer_KindAPIVersion   = Manufacturer_Kind + "." + CRDGroupVersion.String()
	Manufacturer_GroupVersionKind = CRDGroupVersion.WithKind(Manufacturer_Kind)
)

func init() {
	SchemeBuilder.Register(&Manufacturer{}, &ManufacturerList{})
}
