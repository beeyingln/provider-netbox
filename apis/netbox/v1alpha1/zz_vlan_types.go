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

type VlanObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VlanParameters struct {

	// Defaults to `""`.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RoleID *float64 `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// +kubebuilder:validation:Optional
	SiteID *float64 `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// Defaults to `active`.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Required
	Tags []*string `json:"tags" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Required
	Vid *float64 `json:"vid" tf:"vid,omitempty"`
}

// VlanSpec defines the desired state of Vlan
type VlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VlanParameters `json:"forProvider"`
}

// VlanStatus defines the observed state of Vlan.
type VlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vlan is the Schema for the Vlans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netboxjet}
type Vlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VlanSpec   `json:"spec"`
	Status            VlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VlanList contains a list of Vlans
type VlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vlan `json:"items"`
}

// Repository type metadata.
var (
	Vlan_Kind             = "Vlan"
	Vlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vlan_Kind}.String()
	Vlan_KindAPIVersion   = Vlan_Kind + "." + CRDGroupVersion.String()
	Vlan_GroupVersionKind = CRDGroupVersion.WithKind(Vlan_Kind)
)

func init() {
	SchemeBuilder.Register(&Vlan{}, &VlanList{})
}
