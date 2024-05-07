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
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *LogAnalyticsSolution) ResolveReferences( // ResolveReferences of this LogAnalyticsSolution.
	ctx context.Context, c client.Reader) error {
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
	{
		m, l, err = apisresolver.GetManagedResource("operationalinsights.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.WorkspaceNameRef,
			Selector:     mg.Spec.ForProvider.WorkspaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceName")
	}
	mg.Spec.ForProvider.WorkspaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("operationalinsights.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WorkspaceResourceIDRef,
			Selector:     mg.Spec.ForProvider.WorkspaceResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceResourceID")
	}
	mg.Spec.ForProvider.WorkspaceResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("operationalinsights.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WorkspaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.WorkspaceNameRef,
			Selector:     mg.Spec.InitProvider.WorkspaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WorkspaceName")
	}
	mg.Spec.InitProvider.WorkspaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WorkspaceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("operationalinsights.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WorkspaceResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.WorkspaceResourceIDRef,
			Selector:     mg.Spec.InitProvider.WorkspaceResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WorkspaceResourceID")
	}
	mg.Spec.InitProvider.WorkspaceResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WorkspaceResourceIDRef = rsp.ResolvedReference

	return nil
}
