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

func (mg *IOTHubDeviceUpdateAccount) ResolveReferences( // ResolveReferences of this IOTHubDeviceUpdateAccount.
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

	return nil
}

// ResolveReferences of this IOTHubDeviceUpdateInstance.
func (mg *IOTHubDeviceUpdateInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("deviceupdate.azure.upbound.io", "v1beta2", "IOTHubDeviceUpdateAccount", "IOTHubDeviceUpdateAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeviceUpdateAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DeviceUpdateAccountIDRef,
			Selector:     mg.Spec.ForProvider.DeviceUpdateAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeviceUpdateAccountID")
	}
	mg.Spec.ForProvider.DeviceUpdateAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeviceUpdateAccountIDRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.DiagnosticStorageAccount != nil {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DiagnosticStorageAccount.ID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.DiagnosticStorageAccount.IDRef,
				Selector:     mg.Spec.ForProvider.DiagnosticStorageAccount.IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DiagnosticStorageAccount.ID")
		}
		mg.Spec.ForProvider.DiagnosticStorageAccount.ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DiagnosticStorageAccount.IDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta2", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IOTHubID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.IOTHubIDRef,
			Selector:     mg.Spec.ForProvider.IOTHubIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IOTHubID")
	}
	mg.Spec.ForProvider.IOTHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IOTHubIDRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.DiagnosticStorageAccount != nil {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DiagnosticStorageAccount.ID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.DiagnosticStorageAccount.IDRef,
				Selector:     mg.Spec.InitProvider.DiagnosticStorageAccount.IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DiagnosticStorageAccount.ID")
		}
		mg.Spec.InitProvider.DiagnosticStorageAccount.ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DiagnosticStorageAccount.IDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta2", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IOTHubID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.IOTHubIDRef,
			Selector:     mg.Spec.InitProvider.IOTHubIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IOTHubID")
	}
	mg.Spec.InitProvider.IOTHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IOTHubIDRef = rsp.ResolvedReference

	return nil
}
