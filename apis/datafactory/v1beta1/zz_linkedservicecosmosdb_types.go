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

type LinkedServiceCosmosDBInitParameters struct {

	// The endpoint of the Azure CosmosDB account. Required if connection_string is unspecified.
	AccountEndpoint *string `json:"accountEndpoint,omitempty" tf:"account_endpoint,omitempty"`

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The name of the database. Required if connection_string is unspecified.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceCosmosDBObservation struct {

	// The endpoint of the Azure CosmosDB account. Required if connection_string is unspecified.
	AccountEndpoint *string `json:"accountEndpoint,omitempty" tf:"account_endpoint,omitempty"`

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The name of the database. Required if connection_string is unspecified.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceCosmosDBParameters struct {

	// The endpoint of the Azure CosmosDB account. Required if connection_string is unspecified.
	// +kubebuilder:validation:Optional
	AccountEndpoint *string `json:"accountEndpoint,omitempty" tf:"account_endpoint,omitempty"`

	// The account key of the Azure Cosmos DB account. Required if connection_string is unspecified.
	// +kubebuilder:validation:Optional
	AccountKeySecretRef *v1.SecretKeySelector `json:"accountKeySecretRef,omitempty" tf:"-"`

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string. Required if account_endpoint, account_key, and database are unspecified.
	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef *v1.SecretKeySelector `json:"connectionStringSecretRef,omitempty" tf:"-"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The name of the database. Required if connection_string is unspecified.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// LinkedServiceCosmosDBSpec defines the desired state of LinkedServiceCosmosDB
type LinkedServiceCosmosDBSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceCosmosDBParameters `json:"forProvider"`
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
	InitProvider LinkedServiceCosmosDBInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceCosmosDBStatus defines the observed state of LinkedServiceCosmosDB.
type LinkedServiceCosmosDBStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceCosmosDBObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LinkedServiceCosmosDB is the Schema for the LinkedServiceCosmosDBs API. Manages a Linked Service (connection) between a CosmosDB and Azure Data Factory using SQL API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceCosmosDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceCosmosDBSpec   `json:"spec"`
	Status            LinkedServiceCosmosDBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceCosmosDBList contains a list of LinkedServiceCosmosDBs
type LinkedServiceCosmosDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceCosmosDB `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceCosmosDB_Kind             = "LinkedServiceCosmosDB"
	LinkedServiceCosmosDB_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceCosmosDB_Kind}.String()
	LinkedServiceCosmosDB_KindAPIVersion   = LinkedServiceCosmosDB_Kind + "." + CRDGroupVersion.String()
	LinkedServiceCosmosDB_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceCosmosDB_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceCosmosDB{}, &LinkedServiceCosmosDBList{})
}
