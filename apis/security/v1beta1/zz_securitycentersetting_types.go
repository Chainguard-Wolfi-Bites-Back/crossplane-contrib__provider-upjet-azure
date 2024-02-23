// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type SecurityCenterSettingInitParameters struct {

	// Boolean flag to enable/disable data access.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The setting to manage. Possible values are MCAS , WDATP and SENTINEL. Changing this forces a new resource to be created.
	SettingName *string `json:"settingName,omitempty" tf:"setting_name,omitempty"`
}

type SecurityCenterSettingObservation struct {

	// Boolean flag to enable/disable data access.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The subscription security center setting id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The setting to manage. Possible values are MCAS , WDATP and SENTINEL. Changing this forces a new resource to be created.
	SettingName *string `json:"settingName,omitempty" tf:"setting_name,omitempty"`
}

type SecurityCenterSettingParameters struct {

	// Boolean flag to enable/disable data access.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The setting to manage. Possible values are MCAS , WDATP and SENTINEL. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SettingName *string `json:"settingName,omitempty" tf:"setting_name,omitempty"`
}

// SecurityCenterSettingSpec defines the desired state of SecurityCenterSetting
type SecurityCenterSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityCenterSettingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SecurityCenterSettingInitParameters `json:"initProvider,omitempty"`
}

// SecurityCenterSettingStatus defines the observed state of SecurityCenterSetting.
type SecurityCenterSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityCenterSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityCenterSetting is the Schema for the SecurityCenterSettings API. Manages the Data Access Settings for Azure Security Center.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.settingName) || (has(self.initProvider) && has(self.initProvider.settingName))",message="spec.forProvider.settingName is a required parameter"
	Spec   SecurityCenterSettingSpec   `json:"spec"`
	Status SecurityCenterSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterSettingList contains a list of SecurityCenterSettings
type SecurityCenterSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterSetting `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterSetting_Kind             = "SecurityCenterSetting"
	SecurityCenterSetting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityCenterSetting_Kind}.String()
	SecurityCenterSetting_KindAPIVersion   = SecurityCenterSetting_Kind + "." + CRDGroupVersion.String()
	SecurityCenterSetting_GroupVersionKind = CRDGroupVersion.WithKind(SecurityCenterSetting_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterSetting{}, &SecurityCenterSettingList{})
}
