/*
Package networks contains functionality for working with Neutron network
resources. A network is an isolated virtual layer-2 broadcast domain that is
typically reserved for the tenant who created it (unless you configure the
network to be shared). Tenants can create multiple networks until the
thresholds per-tenant quota is reached.

In the v2.0 Networking API, the network is the main entity. Ports and subnets
are always associated with a network.

Example to List Networks

	listOpts := networks.ListOpts{
		TenantID: "a99e9b4e620e4db09a2dfb6e42a01e66",
	}

	allPages, err := networks.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		panic(err)
	}

	for _, network := range allNetworks {
		fmt.Printf("%+v", network)
	}
*/
package networks
