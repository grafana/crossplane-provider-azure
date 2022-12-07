/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	resourcegroup "github.com/upbound/provider-azure/internal/controller/azure/resourcegroup"
	configuration "github.com/upbound/provider-azure/internal/controller/dbformysql/configuration"
	firewallrule "github.com/upbound/provider-azure/internal/controller/dbformysql/firewallrule"
	flexibledatabase "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibledatabase"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserverfirewallrule"
	server "github.com/upbound/provider-azure/internal/controller/dbformysql/server"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/dbformysql/virtualnetworkrule"
	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/configuration"
	database "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/database"
	firewallruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/firewallrule"
	flexibleserverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserver"
	flexibleserverconfigurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverfirewallrule"
	serverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/server"
	serverkey "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/serverkey"
	virtualnetworkruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/virtualnetworkrule"
	key "github.com/upbound/provider-azure/internal/controller/keyvault/key"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/keyvault/managedstorageaccount"
	vault "github.com/upbound/provider-azure/internal/controller/keyvault/vault"
	privatednszone "github.com/upbound/provider-azure/internal/controller/network/privatednszone"
	subnet "github.com/upbound/provider-azure/internal/controller/network/subnet"
	virtualnetwork "github.com/upbound/provider-azure/internal/controller/network/virtualnetwork"
	providerconfig "github.com/upbound/provider-azure/internal/controller/providerconfig"
	account "github.com/upbound/provider-azure/internal/controller/storage/account"
	blob "github.com/upbound/provider-azure/internal/controller/storage/blob"
	container "github.com/upbound/provider-azure/internal/controller/storage/container"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcegroup.Setup,
		configuration.Setup,
		firewallrule.Setup,
		flexibledatabase.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverfirewallrule.Setup,
		server.Setup,
		virtualnetworkrule.Setup,
		activedirectoryadministrator.Setup,
		configurationdbforpostgresql.Setup,
		database.Setup,
		firewallruledbforpostgresql.Setup,
		flexibleserverdbforpostgresql.Setup,
		flexibleserverconfigurationdbforpostgresql.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallruledbforpostgresql.Setup,
		serverdbforpostgresql.Setup,
		serverkey.Setup,
		virtualnetworkruledbforpostgresql.Setup,
		key.Setup,
		managedstorageaccount.Setup,
		vault.Setup,
		privatednszone.Setup,
		subnet.Setup,
		virtualnetwork.Setup,
		providerconfig.Setup,
		account.Setup,
		blob.Setup,
		container.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
