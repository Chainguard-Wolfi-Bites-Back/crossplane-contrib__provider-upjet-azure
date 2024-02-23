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

type BlobInventoryPolicyInitParameters struct {

	// One or more rules blocks as defined below.
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// The ID of the storage account to apply this Blob Inventory Policy to. Changing this forces a new Storage Blob Inventory Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

type BlobInventoryPolicyObservation struct {

	// The ID of the Storage Blob Inventory Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more rules blocks as defined below.
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// The ID of the storage account to apply this Blob Inventory Policy to. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type BlobInventoryPolicyParameters struct {

	// One or more rules blocks as defined below.
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// The ID of the storage account to apply this Blob Inventory Policy to. Changing this forces a new Storage Blob Inventory Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

type FilterInitParameters struct {

	// A set of blob types. Possible values are blockBlob, appendBlob, and pageBlob. The storage account with is_hns_enabled is true doesn't support pageBlob.
	// +listType=set
	BlobTypes []*string `json:"blobTypes,omitempty" tf:"blob_types,omitempty"`

	// A set of strings for blob prefixes to be excluded. Maximum of 10 blob prefixes.
	// +listType=set
	ExcludePrefixes []*string `json:"excludePrefixes,omitempty" tf:"exclude_prefixes,omitempty"`

	// Includes blob versions in blob inventory or not? Defaults to false.
	IncludeBlobVersions *bool `json:"includeBlobVersions,omitempty" tf:"include_blob_versions,omitempty"`

	// Includes deleted blobs in blob inventory or not? Defaults to false.
	IncludeDeleted *bool `json:"includeDeleted,omitempty" tf:"include_deleted,omitempty"`

	// Includes blob snapshots in blob inventory or not? Defaults to false.
	IncludeSnapshots *bool `json:"includeSnapshots,omitempty" tf:"include_snapshots,omitempty"`

	// A set of strings for blob prefixes to be matched. Maximum of 10 blob prefixes.
	// +listType=set
	PrefixMatch []*string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`
}

type FilterObservation struct {

	// A set of blob types. Possible values are blockBlob, appendBlob, and pageBlob. The storage account with is_hns_enabled is true doesn't support pageBlob.
	// +listType=set
	BlobTypes []*string `json:"blobTypes,omitempty" tf:"blob_types,omitempty"`

	// A set of strings for blob prefixes to be excluded. Maximum of 10 blob prefixes.
	// +listType=set
	ExcludePrefixes []*string `json:"excludePrefixes,omitempty" tf:"exclude_prefixes,omitempty"`

	// Includes blob versions in blob inventory or not? Defaults to false.
	IncludeBlobVersions *bool `json:"includeBlobVersions,omitempty" tf:"include_blob_versions,omitempty"`

	// Includes deleted blobs in blob inventory or not? Defaults to false.
	IncludeDeleted *bool `json:"includeDeleted,omitempty" tf:"include_deleted,omitempty"`

	// Includes blob snapshots in blob inventory or not? Defaults to false.
	IncludeSnapshots *bool `json:"includeSnapshots,omitempty" tf:"include_snapshots,omitempty"`

	// A set of strings for blob prefixes to be matched. Maximum of 10 blob prefixes.
	// +listType=set
	PrefixMatch []*string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`
}

type FilterParameters struct {

	// A set of blob types. Possible values are blockBlob, appendBlob, and pageBlob. The storage account with is_hns_enabled is true doesn't support pageBlob.
	// +kubebuilder:validation:Optional
	// +listType=set
	BlobTypes []*string `json:"blobTypes" tf:"blob_types,omitempty"`

	// A set of strings for blob prefixes to be excluded. Maximum of 10 blob prefixes.
	// +kubebuilder:validation:Optional
	// +listType=set
	ExcludePrefixes []*string `json:"excludePrefixes,omitempty" tf:"exclude_prefixes,omitempty"`

	// Includes blob versions in blob inventory or not? Defaults to false.
	// +kubebuilder:validation:Optional
	IncludeBlobVersions *bool `json:"includeBlobVersions,omitempty" tf:"include_blob_versions,omitempty"`

	// Includes deleted blobs in blob inventory or not? Defaults to false.
	// +kubebuilder:validation:Optional
	IncludeDeleted *bool `json:"includeDeleted,omitempty" tf:"include_deleted,omitempty"`

	// Includes blob snapshots in blob inventory or not? Defaults to false.
	// +kubebuilder:validation:Optional
	IncludeSnapshots *bool `json:"includeSnapshots,omitempty" tf:"include_snapshots,omitempty"`

	// A set of strings for blob prefixes to be matched. Maximum of 10 blob prefixes.
	// +kubebuilder:validation:Optional
	// +listType=set
	PrefixMatch []*string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`
}

type RulesInitParameters struct {

	// A filter block as defined above. Can only be set when the scope is Blob.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// The format of the inventory files. Possible values are Csv and Parquet.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The name which should be used for this Blob Inventory Policy Rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The inventory schedule applied by this rule. Possible values are Daily and Weekly.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// A list of fields to be included in the inventory. See the Azure API reference for all the supported fields.
	SchemaFields []*string `json:"schemaFields,omitempty" tf:"schema_fields,omitempty"`

	// The scope of the inventory for this rule. Possible values are Blob and Container.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The storage container name to store the blob inventory files for this rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// Reference to a Container in storage to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameRef *v1.Reference `json:"storageContainerNameRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameSelector *v1.Selector `json:"storageContainerNameSelector,omitempty" tf:"-"`
}

type RulesObservation struct {

	// A filter block as defined above. Can only be set when the scope is Blob.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// The format of the inventory files. Possible values are Csv and Parquet.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The name which should be used for this Blob Inventory Policy Rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The inventory schedule applied by this rule. Possible values are Daily and Weekly.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// A list of fields to be included in the inventory. See the Azure API reference for all the supported fields.
	SchemaFields []*string `json:"schemaFields,omitempty" tf:"schema_fields,omitempty"`

	// The scope of the inventory for this rule. Possible values are Blob and Container.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The storage container name to store the blob inventory files for this rule.
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`
}

type RulesParameters struct {

	// A filter block as defined above. Can only be set when the scope is Blob.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// The format of the inventory files. Possible values are Csv and Parquet.
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`

	// The name which should be used for this Blob Inventory Policy Rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The inventory schedule applied by this rule. Possible values are Daily and Weekly.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`

	// A list of fields to be included in the inventory. See the Azure API reference for all the supported fields.
	// +kubebuilder:validation:Optional
	SchemaFields []*string `json:"schemaFields" tf:"schema_fields,omitempty"`

	// The scope of the inventory for this rule. Possible values are Blob and Container.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// The storage container name to store the blob inventory files for this rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// Reference to a Container in storage to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameRef *v1.Reference `json:"storageContainerNameRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameSelector *v1.Selector `json:"storageContainerNameSelector,omitempty" tf:"-"`
}

// BlobInventoryPolicySpec defines the desired state of BlobInventoryPolicy
type BlobInventoryPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BlobInventoryPolicyParameters `json:"forProvider"`
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
	InitProvider BlobInventoryPolicyInitParameters `json:"initProvider,omitempty"`
}

// BlobInventoryPolicyStatus defines the observed state of BlobInventoryPolicy.
type BlobInventoryPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BlobInventoryPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BlobInventoryPolicy is the Schema for the BlobInventoryPolicys API. Manages a Storage Blob Inventory Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BlobInventoryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rules) || (has(self.initProvider) && has(self.initProvider.rules))",message="spec.forProvider.rules is a required parameter"
	Spec   BlobInventoryPolicySpec   `json:"spec"`
	Status BlobInventoryPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlobInventoryPolicyList contains a list of BlobInventoryPolicys
type BlobInventoryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlobInventoryPolicy `json:"items"`
}

// Repository type metadata.
var (
	BlobInventoryPolicy_Kind             = "BlobInventoryPolicy"
	BlobInventoryPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BlobInventoryPolicy_Kind}.String()
	BlobInventoryPolicy_KindAPIVersion   = BlobInventoryPolicy_Kind + "." + CRDGroupVersion.String()
	BlobInventoryPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BlobInventoryPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BlobInventoryPolicy{}, &BlobInventoryPolicyList{})
}
