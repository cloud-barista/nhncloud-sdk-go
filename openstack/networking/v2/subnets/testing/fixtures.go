package testing

const SubnetListResult = `
{
    "subnets": [
        {
            "name": "private-subnet",
            "enable_dhcp": true,
            "network_id": "db193ab3-96e3-4cb3-8fc5-05f4296d0324",
            "tenant_id": "26a7980765d0414dbc1fc1f88cdb7e6e",
            "dns_nameservers": [],
            "allocation_pools": [
                {
                    "start": "10.0.0.2",
                    "end": "10.0.0.254"
                }
            ],
            "host_routes": [],
            "ip_version": 4,
            "gateway_ip": "10.0.0.1",
            "cidr": "10.0.0.0/24",
            "id": "08eae331-0402-425a-923c-34f7cfe39c1b"
        },
        {
            "name": "my_subnet",
            "enable_dhcp": true,
            "network_id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
            "tenant_id": "4fd44f30292945e481c7b8a0c8908869",
            "dns_nameservers": [],
            "allocation_pools": [
                {
                    "start": "192.0.0.2",
                    "end": "192.255.255.254"
                }
            ],
            "host_routes": [],
            "ip_version": 4,
            "gateway_ip": "192.0.0.1",
            "cidr": "192.0.0.0/8",
            "id": "54d6f61d-db07-451c-9ab3-b9609b6b6f0b"
        },
        {
            "name": "my_gatewayless_subnet",
            "enable_dhcp": true,
            "network_id": "d32019d3-bc6e-4319-9c1d-6722fc136a23",
            "tenant_id": "4fd44f30292945e481c7b8a0c8908869",
            "dns_nameservers": [],
            "allocation_pools": [
                {
                    "start": "192.168.1.2",
                    "end": "192.168.1.254"
                }
            ],
            "host_routes": [],
            "ip_version": 4,
            "gateway_ip": null,
            "cidr": "192.168.1.0/24",
            "id": "54d6f61d-db07-451c-9ab3-b9609b6b6f0c"
        },
        {
            "name": "my_subnet_with_subnetpool",
            "enable_dhcp": false,
            "network_id": "d32019d3-bc6e-4319-9c1d-6722fc136a23",
            "tenant_id": "4fd44f30292945e481c7b8a0c8908869",
            "dns_nameservers": [],
            "allocation_pools": [
                {
                    "start": "10.11.12.2",
                    "end": "10.11.12.254"
                }
            ],
            "host_routes": [],
            "ip_version": 4,
            "gateway_ip": null,
            "cidr": "10.11.12.0/24",
            "id": "38186a51-f373-4bbc-838b-6eaa1aa13eac",
            "subnetpool_id": "b80340c7-9960-4f67-a99c-02501656284b"
        }
    ]
}
`
