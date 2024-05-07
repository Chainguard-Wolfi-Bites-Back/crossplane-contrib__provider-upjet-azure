// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Account.
	apisresolver "github.com/upbound/provider-azure/internal/apis"
)

func (mg *Account) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SnapshotPolicy.
func (mg *SnapshotPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AccountNameRef,
			Selector:     mg.Spec.ForProvider.AccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountName")
	}
	mg.Spec.ForProvider.AccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Volume.
func (mg *Volume) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AccountNameRef,
			Selector:     mg.Spec.ForProvider.AccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountName")
	}
	mg.Spec.ForProvider.AccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "Snapshot", "SnapshotList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CreateFromSnapshotResourceID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.CreateFromSnapshotResourceIDRef,
			Selector:     mg.Spec.ForProvider.CreateFromSnapshotResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CreateFromSnapshotResourceID")
	}
	mg.Spec.ForProvider.CreateFromSnapshotResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CreateFromSnapshotResourceIDRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.DataProtectionReplication != nil {
		{
			m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "Volume", "VolumeList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataProtectionReplication.RemoteVolumeResourceID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.DataProtectionReplication.RemoteVolumeResourceIDRef,
				Selector:     mg.Spec.ForProvider.DataProtectionReplication.RemoteVolumeResourceIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataProtectionReplication.RemoteVolumeResourceID")
		}
		mg.Spec.ForProvider.DataProtectionReplication.RemoteVolumeResourceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataProtectionReplication.RemoteVolumeResourceIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.DataProtectionSnapshotPolicy != nil {
		{
			m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "SnapshotPolicy", "SnapshotPolicyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataProtectionSnapshotPolicy.SnapshotPolicyID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.DataProtectionSnapshotPolicy.SnapshotPolicyIDRef,
				Selector:     mg.Spec.ForProvider.DataProtectionSnapshotPolicy.SnapshotPolicyIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataProtectionSnapshotPolicy.SnapshotPolicyID")
		}
		mg.Spec.ForProvider.DataProtectionSnapshotPolicy.SnapshotPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataProtectionSnapshotPolicy.SnapshotPolicyIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "Pool", "PoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PoolName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PoolNameRef,
			Selector:     mg.Spec.ForProvider.PoolNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PoolName")
	}
	mg.Spec.ForProvider.PoolName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PoolNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubnetIDRef,
			Selector:     mg.Spec.ForProvider.SubnetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "Snapshot", "SnapshotList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CreateFromSnapshotResourceID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.CreateFromSnapshotResourceIDRef,
			Selector:     mg.Spec.InitProvider.CreateFromSnapshotResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CreateFromSnapshotResourceID")
	}
	mg.Spec.InitProvider.CreateFromSnapshotResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CreateFromSnapshotResourceIDRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.DataProtectionReplication != nil {
		{
			m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "Volume", "VolumeList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataProtectionReplication.RemoteVolumeResourceID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.DataProtectionReplication.RemoteVolumeResourceIDRef,
				Selector:     mg.Spec.InitProvider.DataProtectionReplication.RemoteVolumeResourceIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DataProtectionReplication.RemoteVolumeResourceID")
		}
		mg.Spec.InitProvider.DataProtectionReplication.RemoteVolumeResourceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DataProtectionReplication.RemoteVolumeResourceIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.DataProtectionSnapshotPolicy != nil {
		{
			m, l, err = apisresolver.GetManagedResource("netapp.azure.upbound.io", "v1beta2", "SnapshotPolicy", "SnapshotPolicyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataProtectionSnapshotPolicy.SnapshotPolicyID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.DataProtectionSnapshotPolicy.SnapshotPolicyIDRef,
				Selector:     mg.Spec.InitProvider.DataProtectionSnapshotPolicy.SnapshotPolicyIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DataProtectionSnapshotPolicy.SnapshotPolicyID")
		}
		mg.Spec.InitProvider.DataProtectionSnapshotPolicy.SnapshotPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DataProtectionSnapshotPolicy.SnapshotPolicyIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubnetIDRef,
			Selector:     mg.Spec.InitProvider.SubnetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}
