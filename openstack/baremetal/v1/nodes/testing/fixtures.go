package testing

import "github.com/abramad-labs/gophercloud-utils-x/openstack/baremetal/v1/nodes"

const IgnitionConfig = `
{
    "ignition": {
        "version": "2.2.0"
    },
    "systemd": {
        "units": [
            {
                "enabled": true,
                "name": "example.service"
            }
        ]
    }
}
`

const OpenstackMetaDataJSON = `
{
    "availability_zone": "nova",
    "hostname": "test.novalocal",
    "public_keys": {
        "mykey": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDBqUfVvCSez0/Wfpd8dLLgZXV9GtXQ7hnMN+Z0OWQUyebVEHey1CXuin0uY1cAJMhUq8j98SiW+cU0sU4J3x5l2+xi1bodDm1BtFWVeLIOQINpfV1n8fKjHB+ynPpe1F6tMDvrFGUlJs44t30BrujMXBe8Rq44cCk6wqyjATA3rQ== Generated by Nova\n"
    }
}
`

const NetworkDataJSON = `
"services": [
    {
        "type": "dns",
        "address": "8.8.8.8"
    },
    {
        "type": "dns",
        "address": "8.8.4.4"
    }
]
`

const CloudInitString = `
#cloud-init

groups:
  - ubuntu: [root,sys]
  - cloud-users
`

var (
	IgnitionUserData = nodes.UserDataMap{
		"ignition": map[string]string{
			"version": "2.2.0",
		},
		"systemd": map[string]interface{}{
			"units": []map[string]interface{}{{
				"name":    "example.service",
				"enabled": true,
			},
			},
		},
	}

	OpenStackMetaData = map[string]interface{}{
		"availability_zone": "nova",
		"hostname":          "test.novalocal",
		"public_keys": map[string]string{
			"mykey": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDBqUfVvCSez0/Wfpd8dLLgZXV9GtXQ7hnMN+Z0OWQUyebVEHey1CXuin0uY1cAJMhUq8j98SiW+cU0sU4J3x5l2+xi1bodDm1BtFWVeLIOQINpfV1n8fKjHB+ynPpe1F6tMDvrFGUlJs44t30BrujMXBe8Rq44cCk6wqyjATA3rQ== Generated by Nova\n",
		},
	}

	NetworkData = map[string]interface{}{
		"services": []map[string]string{
			{
				"type":    "dns",
				"address": "8.8.8.8",
			},
			{
				"type":    "dns",
				"address": "8.8.4.4",
			},
		},
	}

	CloudInitUserData = nodes.UserDataString(CloudInitString)

	ConfigDrive = nodes.ConfigDrive{
		UserData:    IgnitionUserData,
		MetaData:    OpenStackMetaData,
		NetworkData: NetworkData,
	}

	ConfigDriveVersioned = nodes.ConfigDrive{
		UserData:    IgnitionUserData,
		MetaData:    OpenStackMetaData,
		NetworkData: NetworkData,
		Version:     "2018-10-10",
	}
)
