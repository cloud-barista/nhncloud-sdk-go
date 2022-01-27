package resourceproviders

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

const (
	apiName = "resource_providers"
)

func resourceProvidersListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func getResourceProviderUsagesURL(client *gophercloud.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "usages")
}

func getResourceProviderInventoriesURL(client *gophercloud.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "inventories")
}

func getResourceProviderAllocationsURL(client *gophercloud.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "allocations")
}

func getResourceProviderTraitsURL(client *gophercloud.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "traits")
}
