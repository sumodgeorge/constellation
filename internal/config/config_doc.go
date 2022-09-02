// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by hack/docgen tool. DO NOT EDIT.

package config

import (
	"github.com/talos-systems/talos/pkg/machinery/config/encoder"
)

var (
	ConfigDoc         encoder.Doc
	UpgradeConfigDoc  encoder.Doc
	UserKeyDoc        encoder.Doc
	FirewallRuleDoc   encoder.Doc
	ProviderConfigDoc encoder.Doc
	AzureConfigDoc    encoder.Doc
	GCPConfigDoc      encoder.Doc
	QEMUConfigDoc     encoder.Doc
)

func init() {
	ConfigDoc.Type = "Config"
	ConfigDoc.Comments[encoder.LineComment] = "Config defines configuration used by CLI."
	ConfigDoc.Description = "Config defines configuration used by CLI."
	ConfigDoc.Fields = make([]encoder.Doc, 10)
	ConfigDoc.Fields[0].Name = "version"
	ConfigDoc.Fields[0].Type = "string"
	ConfigDoc.Fields[0].Note = ""
	ConfigDoc.Fields[0].Description = "Schema version of this configuration file."
	ConfigDoc.Fields[0].Comments[encoder.LineComment] = "Schema version of this configuration file."
	ConfigDoc.Fields[1].Name = "autoscalingNodeGroupMin"
	ConfigDoc.Fields[1].Type = "int"
	ConfigDoc.Fields[1].Note = ""
	ConfigDoc.Fields[1].Description = "Minimum number of worker nodes in autoscaling group."
	ConfigDoc.Fields[1].Comments[encoder.LineComment] = "Minimum number of worker nodes in autoscaling group."
	ConfigDoc.Fields[2].Name = "autoscalingNodeGroupMax"
	ConfigDoc.Fields[2].Type = "int"
	ConfigDoc.Fields[2].Note = ""
	ConfigDoc.Fields[2].Description = "Maximum number of worker nodes in autoscaling group."
	ConfigDoc.Fields[2].Comments[encoder.LineComment] = "Maximum number of worker nodes in autoscaling group."
	ConfigDoc.Fields[3].Name = "stateDiskSizeGB"
	ConfigDoc.Fields[3].Type = "int"
	ConfigDoc.Fields[3].Note = ""
	ConfigDoc.Fields[3].Description = "Size (in GB) of a node's disk to store the non-volatile state."
	ConfigDoc.Fields[3].Comments[encoder.LineComment] = "Size (in GB) of a node's disk to store the non-volatile state."
	ConfigDoc.Fields[4].Name = "ingressFirewall"
	ConfigDoc.Fields[4].Type = "Firewall"
	ConfigDoc.Fields[4].Note = ""
	ConfigDoc.Fields[4].Description = "Ingress firewall rules for node network."
	ConfigDoc.Fields[4].Comments[encoder.LineComment] = "Ingress firewall rules for node network."
	ConfigDoc.Fields[5].Name = "egressFirewall"
	ConfigDoc.Fields[5].Type = "Firewall"
	ConfigDoc.Fields[5].Note = ""
	ConfigDoc.Fields[5].Description = "Egress firewall rules for node network."
	ConfigDoc.Fields[5].Comments[encoder.LineComment] = "Egress firewall rules for node network."

	ConfigDoc.Fields[5].AddExample("", Firewall{{Name: "rule#1", Description: "the first rule", Protocol: "tcp", IPRange: "0.0.0.0/0", FromPort: 443, ToPort: 443}})
	ConfigDoc.Fields[6].Name = "provider"
	ConfigDoc.Fields[6].Type = "ProviderConfig"
	ConfigDoc.Fields[6].Note = ""
	ConfigDoc.Fields[6].Description = "Supported cloud providers and their specific configurations."
	ConfigDoc.Fields[6].Comments[encoder.LineComment] = "Supported cloud providers and their specific configurations."
	ConfigDoc.Fields[7].Name = "sshUsers"
	ConfigDoc.Fields[7].Type = "[]UserKey"
	ConfigDoc.Fields[7].Note = ""
	ConfigDoc.Fields[7].Description = "Create SSH users on Constellation nodes."
	ConfigDoc.Fields[7].Comments[encoder.LineComment] = "Create SSH users on Constellation nodes."

	ConfigDoc.Fields[7].AddExample("", []UserKey{{Username: "Alice", PublicKey: "ssh-rsa AAAAB3NzaC...5QXHKW1rufgtJeSeJ8= alice@domain.com"}})
	ConfigDoc.Fields[8].Name = "kubernetesVersion"
	ConfigDoc.Fields[8].Type = "string"
	ConfigDoc.Fields[8].Note = ""
	ConfigDoc.Fields[8].Description = "Kubernetes version installed in the cluster."
	ConfigDoc.Fields[8].Comments[encoder.LineComment] = "Kubernetes version installed in the cluster."
	ConfigDoc.Fields[9].Name = "upgrade"
	ConfigDoc.Fields[9].Type = "UpgradeConfig"
	ConfigDoc.Fields[9].Note = ""
	ConfigDoc.Fields[9].Description = "Configuration to apply during constellation upgrade."
	ConfigDoc.Fields[9].Comments[encoder.LineComment] = "Configuration to apply during constellation upgrade."

	ConfigDoc.Fields[9].AddExample("", UpgradeConfig{Image: "", Measurements: Measurements{}})

	UpgradeConfigDoc.Type = "UpgradeConfig"
	UpgradeConfigDoc.Comments[encoder.LineComment] = "UpgradeConfig defines configuration used during constellation upgrade."
	UpgradeConfigDoc.Description = "UpgradeConfig defines configuration used during constellation upgrade."

	UpgradeConfigDoc.AddExample("", UpgradeConfig{Image: "", Measurements: Measurements{}})
	UpgradeConfigDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Config",
			FieldName: "upgrade",
		},
	}
	UpgradeConfigDoc.Fields = make([]encoder.Doc, 2)
	UpgradeConfigDoc.Fields[0].Name = "image"
	UpgradeConfigDoc.Fields[0].Type = "string"
	UpgradeConfigDoc.Fields[0].Note = ""
	UpgradeConfigDoc.Fields[0].Description = "Updated machine image to install on all nodes."
	UpgradeConfigDoc.Fields[0].Comments[encoder.LineComment] = "Updated machine image to install on all nodes."
	UpgradeConfigDoc.Fields[1].Name = "measurements"
	UpgradeConfigDoc.Fields[1].Type = "Measurements"
	UpgradeConfigDoc.Fields[1].Note = ""
	UpgradeConfigDoc.Fields[1].Description = "Measurements of the updated image."
	UpgradeConfigDoc.Fields[1].Comments[encoder.LineComment] = "Measurements of the updated image."

	UserKeyDoc.Type = "UserKey"
	UserKeyDoc.Comments[encoder.LineComment] = "UserKey describes a user that should be created with corresponding public SSH key."
	UserKeyDoc.Description = "UserKey describes a user that should be created with corresponding public SSH key."

	UserKeyDoc.AddExample("", []UserKey{{Username: "Alice", PublicKey: "ssh-rsa AAAAB3NzaC...5QXHKW1rufgtJeSeJ8= alice@domain.com"}})
	UserKeyDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Config",
			FieldName: "sshUsers",
		},
	}
	UserKeyDoc.Fields = make([]encoder.Doc, 2)
	UserKeyDoc.Fields[0].Name = "username"
	UserKeyDoc.Fields[0].Type = "string"
	UserKeyDoc.Fields[0].Note = ""
	UserKeyDoc.Fields[0].Description = "Username of new SSH user."
	UserKeyDoc.Fields[0].Comments[encoder.LineComment] = "Username of new SSH user."
	UserKeyDoc.Fields[1].Name = "publicKey"
	UserKeyDoc.Fields[1].Type = "string"
	UserKeyDoc.Fields[1].Note = ""
	UserKeyDoc.Fields[1].Description = "Public key of new SSH user."
	UserKeyDoc.Fields[1].Comments[encoder.LineComment] = "Public key of new SSH user."

	FirewallRuleDoc.Type = "FirewallRule"
	FirewallRuleDoc.Comments[encoder.LineComment] = ""
	FirewallRuleDoc.Description = ""
	FirewallRuleDoc.Fields = make([]encoder.Doc, 6)
	FirewallRuleDoc.Fields[0].Name = "name"
	FirewallRuleDoc.Fields[0].Type = "string"
	FirewallRuleDoc.Fields[0].Note = ""
	FirewallRuleDoc.Fields[0].Description = "Name of rule."
	FirewallRuleDoc.Fields[0].Comments[encoder.LineComment] = "Name of rule."
	FirewallRuleDoc.Fields[1].Name = "description"
	FirewallRuleDoc.Fields[1].Type = "string"
	FirewallRuleDoc.Fields[1].Note = ""
	FirewallRuleDoc.Fields[1].Description = "Description for rule."
	FirewallRuleDoc.Fields[1].Comments[encoder.LineComment] = "Description for rule."
	FirewallRuleDoc.Fields[2].Name = "protocol"
	FirewallRuleDoc.Fields[2].Type = "string"
	FirewallRuleDoc.Fields[2].Note = ""
	FirewallRuleDoc.Fields[2].Description = "Protocol, such as 'udp' or 'tcp'."
	FirewallRuleDoc.Fields[2].Comments[encoder.LineComment] = "Protocol, such as 'udp' or 'tcp'."
	FirewallRuleDoc.Fields[3].Name = "iprange"
	FirewallRuleDoc.Fields[3].Type = "string"
	FirewallRuleDoc.Fields[3].Note = ""
	FirewallRuleDoc.Fields[3].Description = "CIDR range for which this rule is applied."
	FirewallRuleDoc.Fields[3].Comments[encoder.LineComment] = "CIDR range for which this rule is applied."
	FirewallRuleDoc.Fields[4].Name = "fromport"
	FirewallRuleDoc.Fields[4].Type = "int"
	FirewallRuleDoc.Fields[4].Note = ""
	FirewallRuleDoc.Fields[4].Description = "Start port of a range."
	FirewallRuleDoc.Fields[4].Comments[encoder.LineComment] = "Start port of a range."
	FirewallRuleDoc.Fields[5].Name = "toport"
	FirewallRuleDoc.Fields[5].Type = "int"
	FirewallRuleDoc.Fields[5].Note = ""
	FirewallRuleDoc.Fields[5].Description = "End port of a range, or 0 if a single port is given by fromport."
	FirewallRuleDoc.Fields[5].Comments[encoder.LineComment] = "End port of a range, or 0 if a single port is given by fromport."

	ProviderConfigDoc.Type = "ProviderConfig"
	ProviderConfigDoc.Comments[encoder.LineComment] = "ProviderConfig are cloud-provider specific configuration values used by the CLI."
	ProviderConfigDoc.Description = "ProviderConfig are cloud-provider specific configuration values used by the CLI.\nFields should remain pointer-types so custom specific configs can nil them\nif not required.\n"
	ProviderConfigDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Config",
			FieldName: "provider",
		},
	}
	ProviderConfigDoc.Fields = make([]encoder.Doc, 3)
	ProviderConfigDoc.Fields[0].Name = "azure"
	ProviderConfigDoc.Fields[0].Type = "AzureConfig"
	ProviderConfigDoc.Fields[0].Note = ""
	ProviderConfigDoc.Fields[0].Description = "Configuration for Azure as provider."
	ProviderConfigDoc.Fields[0].Comments[encoder.LineComment] = "Configuration for Azure as provider."
	ProviderConfigDoc.Fields[1].Name = "gcp"
	ProviderConfigDoc.Fields[1].Type = "GCPConfig"
	ProviderConfigDoc.Fields[1].Note = ""
	ProviderConfigDoc.Fields[1].Description = "Configuration for Google Cloud as provider."
	ProviderConfigDoc.Fields[1].Comments[encoder.LineComment] = "Configuration for Google Cloud as provider."
	ProviderConfigDoc.Fields[2].Name = "qemu"
	ProviderConfigDoc.Fields[2].Type = "QEMUConfig"
	ProviderConfigDoc.Fields[2].Note = ""
	ProviderConfigDoc.Fields[2].Description = "Configuration for QEMU as provider."
	ProviderConfigDoc.Fields[2].Comments[encoder.LineComment] = "Configuration for QEMU as provider."

	AzureConfigDoc.Type = "AzureConfig"
	AzureConfigDoc.Comments[encoder.LineComment] = "AzureConfig are Azure specific configuration values used by the CLI."
	AzureConfigDoc.Description = "AzureConfig are Azure specific configuration values used by the CLI."
	AzureConfigDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "ProviderConfig",
			FieldName: "azure",
		},
	}
	AzureConfigDoc.Fields = make([]encoder.Doc, 15)
	AzureConfigDoc.Fields[0].Name = "subscription"
	AzureConfigDoc.Fields[0].Type = "string"
	AzureConfigDoc.Fields[0].Note = ""
	AzureConfigDoc.Fields[0].Description = "Subscription ID of the used Azure account. See: https://docs.microsoft.com/en-us/azure/azure-portal/get-subscription-tenant-id#find-your-azure-subscription"
	AzureConfigDoc.Fields[0].Comments[encoder.LineComment] = "Subscription ID of the used Azure account. See: https://docs.microsoft.com/en-us/azure/azure-portal/get-subscription-tenant-id#find-your-azure-subscription"
	AzureConfigDoc.Fields[1].Name = "tenant"
	AzureConfigDoc.Fields[1].Type = "string"
	AzureConfigDoc.Fields[1].Note = ""
	AzureConfigDoc.Fields[1].Description = "Tenant ID of the used Azure account. See: https://docs.microsoft.com/en-us/azure/azure-portal/get-subscription-tenant-id#find-your-azure-ad-tenant"
	AzureConfigDoc.Fields[1].Comments[encoder.LineComment] = "Tenant ID of the used Azure account. See: https://docs.microsoft.com/en-us/azure/azure-portal/get-subscription-tenant-id#find-your-azure-ad-tenant"
	AzureConfigDoc.Fields[2].Name = "location"
	AzureConfigDoc.Fields[2].Type = "string"
	AzureConfigDoc.Fields[2].Note = ""
	AzureConfigDoc.Fields[2].Description = "Azure datacenter region to be used. See: https://docs.microsoft.com/en-us/azure/availability-zones/az-overview#azure-regions-with-availability-zones"
	AzureConfigDoc.Fields[2].Comments[encoder.LineComment] = "Azure datacenter region to be used. See: https://docs.microsoft.com/en-us/azure/availability-zones/az-overview#azure-regions-with-availability-zones"
	AzureConfigDoc.Fields[3].Name = "image"
	AzureConfigDoc.Fields[3].Type = "string"
	AzureConfigDoc.Fields[3].Note = ""
	AzureConfigDoc.Fields[3].Description = "Machine image used to create Constellation nodes."
	AzureConfigDoc.Fields[3].Comments[encoder.LineComment] = "Machine image used to create Constellation nodes."
	AzureConfigDoc.Fields[4].Name = "instanceType"
	AzureConfigDoc.Fields[4].Type = "string"
	AzureConfigDoc.Fields[4].Note = ""
	AzureConfigDoc.Fields[4].Description = "Virtual machine instance type to use for Constellation nodes."
	AzureConfigDoc.Fields[4].Comments[encoder.LineComment] = "Virtual machine instance type to use for Constellation nodes."
	AzureConfigDoc.Fields[5].Name = "stateDiskType"
	AzureConfigDoc.Fields[5].Type = "string"
	AzureConfigDoc.Fields[5].Note = ""
	AzureConfigDoc.Fields[5].Description = "Type of a node's state disk. The type influences boot time and I/O performance. See: https://docs.microsoft.com/en-us/azure/virtual-machines/disks-types#disk-type-comparison"
	AzureConfigDoc.Fields[5].Comments[encoder.LineComment] = "Type of a node's state disk. The type influences boot time and I/O performance. See: https://docs.microsoft.com/en-us/azure/virtual-machines/disks-types#disk-type-comparison"
	AzureConfigDoc.Fields[6].Name = "resourceGroup"
	AzureConfigDoc.Fields[6].Type = "string"
	AzureConfigDoc.Fields[6].Note = ""
	AzureConfigDoc.Fields[6].Description = "Resource group to use."
	AzureConfigDoc.Fields[6].Comments[encoder.LineComment] = "Resource group to use."
	AzureConfigDoc.Fields[7].Name = "userAssignedIdentity"
	AzureConfigDoc.Fields[7].Type = "string"
	AzureConfigDoc.Fields[7].Note = ""
	AzureConfigDoc.Fields[7].Description = "Authorize spawned VMs to access Azure API."
	AzureConfigDoc.Fields[7].Comments[encoder.LineComment] = "Authorize spawned VMs to access Azure API."
	AzureConfigDoc.Fields[8].Name = "appClientID"
	AzureConfigDoc.Fields[8].Type = "string"
	AzureConfigDoc.Fields[8].Note = ""
	AzureConfigDoc.Fields[8].Description = "Application client ID of the Active Directory app registration."
	AzureConfigDoc.Fields[8].Comments[encoder.LineComment] = "Application client ID of the Active Directory app registration."
	AzureConfigDoc.Fields[9].Name = "clientSecretValue"
	AzureConfigDoc.Fields[9].Type = "string"
	AzureConfigDoc.Fields[9].Note = ""
	AzureConfigDoc.Fields[9].Description = "Client secret value of the Active Directory app registration credentials."
	AzureConfigDoc.Fields[9].Comments[encoder.LineComment] = "Client secret value of the Active Directory app registration credentials."
	AzureConfigDoc.Fields[10].Name = "measurements"
	AzureConfigDoc.Fields[10].Type = "Measurements"
	AzureConfigDoc.Fields[10].Note = ""
	AzureConfigDoc.Fields[10].Description = "Expected confidential VM measurements."
	AzureConfigDoc.Fields[10].Comments[encoder.LineComment] = "Expected confidential VM measurements."
	AzureConfigDoc.Fields[11].Name = "enforcedMeasurements"
	AzureConfigDoc.Fields[11].Type = "[]uint32"
	AzureConfigDoc.Fields[11].Note = ""
	AzureConfigDoc.Fields[11].Description = "List of values that should be enforced to be equal to the ones from the measurement list. Any non-equal values not in this list will only result in a warning."
	AzureConfigDoc.Fields[11].Comments[encoder.LineComment] = "List of values that should be enforced to be equal to the ones from the measurement list. Any non-equal values not in this list will only result in a warning."
	AzureConfigDoc.Fields[12].Name = "idKeyDigest"
	AzureConfigDoc.Fields[12].Type = "string"
	AzureConfigDoc.Fields[12].Note = ""
	AzureConfigDoc.Fields[12].Description = "Expected value for the field 'idkeydigest' in the AMD SEV-SNP attestation report. Only usable with ConfidentialVMs. See 4.6 and 7.3 in: https://www.amd.com/system/files/TechDocs/56860.pdf"
	AzureConfigDoc.Fields[12].Comments[encoder.LineComment] = "Expected value for the field 'idkeydigest' in the AMD SEV-SNP attestation report. Only usable with ConfidentialVMs. See 4.6 and 7.3 in: https://www.amd.com/system/files/TechDocs/56860.pdf"
	AzureConfigDoc.Fields[13].Name = "enforceIdKeyDigest"
	AzureConfigDoc.Fields[13].Type = "bool"
	AzureConfigDoc.Fields[13].Note = ""
	AzureConfigDoc.Fields[13].Description = "Enforce the specified idKeyDigest value during remote attestation."
	AzureConfigDoc.Fields[13].Comments[encoder.LineComment] = "Enforce the specified idKeyDigest value during remote attestation."
	AzureConfigDoc.Fields[14].Name = "confidentialVM"
	AzureConfigDoc.Fields[14].Type = "bool"
	AzureConfigDoc.Fields[14].Note = ""
	AzureConfigDoc.Fields[14].Description = "Use VMs with security type Confidential VM. If set to false, Trusted Launch VMs will be used instead. See: https://docs.microsoft.com/en-us/azure/confidential-computing/confidential-vm-overview"
	AzureConfigDoc.Fields[14].Comments[encoder.LineComment] = "Use VMs with security type Confidential VM. If set to false, Trusted Launch VMs will be used instead. See: https://docs.microsoft.com/en-us/azure/confidential-computing/confidential-vm-overview"

	GCPConfigDoc.Type = "GCPConfig"
	GCPConfigDoc.Comments[encoder.LineComment] = "GCPConfig are GCP specific configuration values used by the CLI."
	GCPConfigDoc.Description = "GCPConfig are GCP specific configuration values used by the CLI."
	GCPConfigDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "ProviderConfig",
			FieldName: "gcp",
		},
	}
	GCPConfigDoc.Fields = make([]encoder.Doc, 9)
	GCPConfigDoc.Fields[0].Name = "project"
	GCPConfigDoc.Fields[0].Type = "string"
	GCPConfigDoc.Fields[0].Note = ""
	GCPConfigDoc.Fields[0].Description = "GCP project. See: https://support.google.com/googleapi/answer/7014113?hl=en"
	GCPConfigDoc.Fields[0].Comments[encoder.LineComment] = "GCP project. See: https://support.google.com/googleapi/answer/7014113?hl=en"
	GCPConfigDoc.Fields[1].Name = "region"
	GCPConfigDoc.Fields[1].Type = "string"
	GCPConfigDoc.Fields[1].Note = ""
	GCPConfigDoc.Fields[1].Description = "GCP datacenter region. See: https://cloud.google.com/compute/docs/regions-zones#available"
	GCPConfigDoc.Fields[1].Comments[encoder.LineComment] = "GCP datacenter region. See: https://cloud.google.com/compute/docs/regions-zones#available"
	GCPConfigDoc.Fields[2].Name = "zone"
	GCPConfigDoc.Fields[2].Type = "string"
	GCPConfigDoc.Fields[2].Note = ""
	GCPConfigDoc.Fields[2].Description = "GCP datacenter zone. See: https://cloud.google.com/compute/docs/regions-zones#available"
	GCPConfigDoc.Fields[2].Comments[encoder.LineComment] = "GCP datacenter zone. See: https://cloud.google.com/compute/docs/regions-zones#available"
	GCPConfigDoc.Fields[3].Name = "image"
	GCPConfigDoc.Fields[3].Type = "string"
	GCPConfigDoc.Fields[3].Note = ""
	GCPConfigDoc.Fields[3].Description = "Machine image used to create Constellation nodes."
	GCPConfigDoc.Fields[3].Comments[encoder.LineComment] = "Machine image used to create Constellation nodes."
	GCPConfigDoc.Fields[4].Name = "instanceType"
	GCPConfigDoc.Fields[4].Type = "string"
	GCPConfigDoc.Fields[4].Note = ""
	GCPConfigDoc.Fields[4].Description = "Virtual machine instance type to use for Constellation nodes."
	GCPConfigDoc.Fields[4].Comments[encoder.LineComment] = "Virtual machine instance type to use for Constellation nodes."
	GCPConfigDoc.Fields[5].Name = "stateDiskType"
	GCPConfigDoc.Fields[5].Type = "string"
	GCPConfigDoc.Fields[5].Note = ""
	GCPConfigDoc.Fields[5].Description = "Type of a node's state disk. The type influences boot time and I/O performance. See: https://cloud.google.com/compute/docs/disks#disk-types"
	GCPConfigDoc.Fields[5].Comments[encoder.LineComment] = "Type of a node's state disk. The type influences boot time and I/O performance. See: https://cloud.google.com/compute/docs/disks#disk-types"
	GCPConfigDoc.Fields[6].Name = "serviceAccountKeyPath"
	GCPConfigDoc.Fields[6].Type = "string"
	GCPConfigDoc.Fields[6].Note = ""
	GCPConfigDoc.Fields[6].Description = "Path of service account key file. For needed service account roles, see https://constellation-docs.edgeless.systems/constellation/getting-started/install#authorization"
	GCPConfigDoc.Fields[6].Comments[encoder.LineComment] = "Path of service account key file. For needed service account roles, see https://constellation-docs.edgeless.systems/constellation/getting-started/install#authorization"
	GCPConfigDoc.Fields[7].Name = "measurements"
	GCPConfigDoc.Fields[7].Type = "Measurements"
	GCPConfigDoc.Fields[7].Note = ""
	GCPConfigDoc.Fields[7].Description = "Expected confidential VM measurements."
	GCPConfigDoc.Fields[7].Comments[encoder.LineComment] = "Expected confidential VM measurements."
	GCPConfigDoc.Fields[8].Name = "enforcedMeasurements"
	GCPConfigDoc.Fields[8].Type = "[]uint32"
	GCPConfigDoc.Fields[8].Note = ""
	GCPConfigDoc.Fields[8].Description = "List of values that should be enforced to be equal to the ones from the measurement list. Any non-equal values not in this list will only result in a warning."
	GCPConfigDoc.Fields[8].Comments[encoder.LineComment] = "List of values that should be enforced to be equal to the ones from the measurement list. Any non-equal values not in this list will only result in a warning."

	QEMUConfigDoc.Type = "QEMUConfig"
	QEMUConfigDoc.Comments[encoder.LineComment] = ""
	QEMUConfigDoc.Description = ""
	QEMUConfigDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "ProviderConfig",
			FieldName: "qemu",
		},
	}
	QEMUConfigDoc.Fields = make([]encoder.Doc, 2)
	QEMUConfigDoc.Fields[0].Name = "measurements"
	QEMUConfigDoc.Fields[0].Type = "Measurements"
	QEMUConfigDoc.Fields[0].Note = ""
	QEMUConfigDoc.Fields[0].Description = "Measurement used to enable measured boot."
	QEMUConfigDoc.Fields[0].Comments[encoder.LineComment] = "Measurement used to enable measured boot."
	QEMUConfigDoc.Fields[1].Name = "enforcedMeasurements"
	QEMUConfigDoc.Fields[1].Type = "[]uint32"
	QEMUConfigDoc.Fields[1].Note = ""
	QEMUConfigDoc.Fields[1].Description = "List of values that should be enforced to be equal to the ones from the measurement list. Any non-equal values not in this list will only result in a warning."
	QEMUConfigDoc.Fields[1].Comments[encoder.LineComment] = "List of values that should be enforced to be equal to the ones from the measurement list. Any non-equal values not in this list will only result in a warning."
}

func (_ Config) Doc() *encoder.Doc {
	return &ConfigDoc
}

func (_ UpgradeConfig) Doc() *encoder.Doc {
	return &UpgradeConfigDoc
}

func (_ UserKey) Doc() *encoder.Doc {
	return &UserKeyDoc
}

func (_ FirewallRule) Doc() *encoder.Doc {
	return &FirewallRuleDoc
}

func (_ ProviderConfig) Doc() *encoder.Doc {
	return &ProviderConfigDoc
}

func (_ AzureConfig) Doc() *encoder.Doc {
	return &AzureConfigDoc
}

func (_ GCPConfig) Doc() *encoder.Doc {
	return &GCPConfigDoc
}

func (_ QEMUConfig) Doc() *encoder.Doc {
	return &QEMUConfigDoc
}

// GetConfigurationDoc returns documentation for the file ./config_doc.go.
func GetConfigurationDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "Configuration",
		Description: "This binary can be build from siderolabs/talos projects. Located at:\nhttps://github.com/siderolabs/talos/tree/master/hack/docgen\n",
		Structs: []*encoder.Doc{
			&ConfigDoc,
			&UpgradeConfigDoc,
			&UserKeyDoc,
			&FirewallRuleDoc,
			&ProviderConfigDoc,
			&AzureConfigDoc,
			&GCPConfigDoc,
			&QEMUConfigDoc,
		},
	}
}
