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

type DataSetKustoClusterInitParameters struct {

	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	KustoClusterID *string `json:"kustoClusterId,omitempty" tf:"kusto_cluster_id,omitempty"`

	// Reference to a Cluster in kusto to populate kustoClusterId.
	// +kubebuilder:validation:Optional
	KustoClusterIDRef *v1.Reference `json:"kustoClusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate kustoClusterId.
	// +kubebuilder:validation:Optional
	KustoClusterIDSelector *v1.Selector `json:"kustoClusterIdSelector,omitempty" tf:"-"`
}

type DataSetKustoClusterObservation struct {

	// The name of the Data Share Dataset.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The resource ID of the Data Share Kusto Cluster Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	KustoClusterID *string `json:"kustoClusterId,omitempty" tf:"kusto_cluster_id,omitempty"`

	// The location of the Kusto Cluster.
	KustoClusterLocation *string `json:"kustoClusterLocation,omitempty" tf:"kusto_cluster_location,omitempty"`

	// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`
}

type DataSetKustoClusterParameters struct {

	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KustoClusterID *string `json:"kustoClusterId,omitempty" tf:"kusto_cluster_id,omitempty"`

	// Reference to a Cluster in kusto to populate kustoClusterId.
	// +kubebuilder:validation:Optional
	KustoClusterIDRef *v1.Reference `json:"kustoClusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate kustoClusterId.
	// +kubebuilder:validation:Optional
	KustoClusterIDSelector *v1.Selector `json:"kustoClusterIdSelector,omitempty" tf:"-"`

	// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	// +crossplane:generate:reference:type=DataShare
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`

	// Reference to a DataShare to populate shareId.
	// +kubebuilder:validation:Optional
	ShareIDRef *v1.Reference `json:"shareIdRef,omitempty" tf:"-"`

	// Selector for a DataShare to populate shareId.
	// +kubebuilder:validation:Optional
	ShareIDSelector *v1.Selector `json:"shareIdSelector,omitempty" tf:"-"`
}

// DataSetKustoClusterSpec defines the desired state of DataSetKustoCluster
type DataSetKustoClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetKustoClusterParameters `json:"forProvider"`
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
	InitProvider DataSetKustoClusterInitParameters `json:"initProvider,omitempty"`
}

// DataSetKustoClusterStatus defines the observed state of DataSetKustoCluster.
type DataSetKustoClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetKustoClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DataSetKustoCluster is the Schema for the DataSetKustoClusters API. Manages a Data Share Kusto Cluster Dataset.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetKustoCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetKustoClusterSpec   `json:"spec"`
	Status            DataSetKustoClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetKustoClusterList contains a list of DataSetKustoClusters
type DataSetKustoClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetKustoCluster `json:"items"`
}

// Repository type metadata.
var (
	DataSetKustoCluster_Kind             = "DataSetKustoCluster"
	DataSetKustoCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetKustoCluster_Kind}.String()
	DataSetKustoCluster_KindAPIVersion   = DataSetKustoCluster_Kind + "." + CRDGroupVersion.String()
	DataSetKustoCluster_GroupVersionKind = CRDGroupVersion.WithKind(DataSetKustoCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetKustoCluster{}, &DataSetKustoClusterList{})
}
