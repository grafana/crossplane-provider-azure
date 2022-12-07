/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/storage/account"
	blob "github.com/upbound/provider-azure/internal/controller/storage/blob"
	container "github.com/upbound/provider-azure/internal/controller/storage/container"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/storage/managementpolicy"
)

// Setup_storage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		blob.Setup,
		container.Setup,
		managementpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
