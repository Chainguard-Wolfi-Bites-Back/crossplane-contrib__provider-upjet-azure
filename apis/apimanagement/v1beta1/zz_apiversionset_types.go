/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type APIVersionSetObservation struct {

	// The ID of the API Version Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIVersionSetParameters struct {

	// The name of the API Management Service in which the API Version Set should exist. May only contain alphanumeric characters and dashes up to 50 characters in length. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The description of API Version Set.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of this API Version Set.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	// +kubebuilder:validation:Optional
	VersionHeaderName *string `json:"versionHeaderName,omitempty" tf:"version_header_name,omitempty"`

	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	// +kubebuilder:validation:Optional
	VersionQueryName *string `json:"versionQueryName,omitempty" tf:"version_query_name,omitempty"`

	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are Header, Query and Segment.
	// +kubebuilder:validation:Required
	VersioningScheme *string `json:"versioningScheme" tf:"versioning_scheme,omitempty"`
}

// APIVersionSetSpec defines the desired state of APIVersionSet
type APIVersionSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIVersionSetParameters `json:"forProvider"`
}

// APIVersionSetStatus defines the observed state of APIVersionSet.
type APIVersionSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIVersionSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIVersionSet is the Schema for the APIVersionSets API. Manages an API Version Set within an API Management Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type APIVersionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIVersionSetSpec   `json:"spec"`
	Status            APIVersionSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIVersionSetList contains a list of APIVersionSets
type APIVersionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIVersionSet `json:"items"`
}

// Repository type metadata.
var (
	APIVersionSet_Kind             = "APIVersionSet"
	APIVersionSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIVersionSet_Kind}.String()
	APIVersionSet_KindAPIVersion   = APIVersionSet_Kind + "." + CRDGroupVersion.String()
	APIVersionSet_GroupVersionKind = CRDGroupVersion.WithKind(APIVersionSet_Kind)
)

func init() {
	SchemeBuilder.Register(&APIVersionSet{}, &APIVersionSetList{})
}