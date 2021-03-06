// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

const (
	// APIVersion is the version of this API
	APIVersion = "vlabs"
)

// the orchestrators supported by vlabs
const (
	// DCOS is the string constant for DCOS orchestrator type and defaults to DCOS188
	DCOS string = "DCOS"
	// Swarm is the string constant for the Swarm orchestrator type
	Swarm string = "Swarm"
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes string = "Kubernetes"
	// SwarmMode is the string constant for the Swarm Mode orchestrator type
	SwarmMode string = "SwarmMode"
)

// the OSTypes supported by vlabs
const (
	Windows OSType = "Windows"
	Linux   OSType = "Linux"
)

// the LinuxDistros supported by vlabs
const (
	Ubuntu          Distro = "ubuntu"
	Ubuntu1804      Distro = "ubuntu-18.04"
	RHEL            Distro = "rhel"
	CoreOS          Distro = "coreos"
	AKS             Distro = "aks"
	AKSDockerEngine Distro = "aks-docker-engine"
	AKS1804         Distro = "aks-18.04"
)

// validation values
const (
	// MinAgentCount are the minimum number of agents per agent pool
	MinAgentCount = 1
	// MaxAgentCount are the maximum number of agents per agent pool
	MaxAgentCount = 100
	// MinPort specifies the minimum tcp port to open
	MinPort = 1
	// MaxPort specifies the maximum tcp port to open
	MaxPort = 65535
	// MaxDisks specifies the maximum attached disks to add to the cluster
	MaxDisks = 4
	// MinDiskSizeGB specifies the minimum attached disk size
	MinDiskSizeGB = 1
	// MaxDiskSizeGB specifies the maximum attached disk size
	MaxDiskSizeGB = 1023
	// MinIPAddressCount specifies the minimum number of IP addresses per network interface
	MinIPAddressCount = 1
	// MaxIPAddressCount specifies the maximum number of IP addresses per network interface
	MaxIPAddressCount = 256
)

// Availability profiles
const (
	// AvailabilitySet means that the vms are in an availability set
	AvailabilitySet = "AvailabilitySet"
	// VirtualMachineScaleSets means that the vms are in a virtual machine scaleset
	VirtualMachineScaleSets = "VirtualMachineScaleSets"
)

// storage profiles
const (
	// StorageAccount means that the nodes use raw storage accounts for their os and attached volumes
	StorageAccount = "StorageAccount"
	// ManagedDisks means that the nodes use managed disks for their os and attached volumes
	ManagedDisks = "ManagedDisks"
)

// Supported container runtimes
const (
	Docker          = "docker"
	ClearContainers = "clear-containers"
	KataContainers  = "kata-containers"
	Containerd      = "containerd"
)

var (
	// NetworkPluginValues holds the valid values for network plugin implementation
	NetworkPluginValues = [...]string{"", "kubenet", "azure", NetworkPluginCilium, "flannel"}

	// NetworkPolicyValues holds the valid values for a network policy
	// "azure" and "none" are there for backwards-compatibility
	NetworkPolicyValues = [...]string{"", "calico", NetworkPolicyCilium, "azure", "none"}

	// ContainerRuntimeValues holds the valid values for container runtimes
	ContainerRuntimeValues = [...]string{"", Docker, ClearContainers, KataContainers, Containerd}
)

// Kubernetes configuration
const (
	// KubernetesMinMaxPods is the minimum valid value for MaxPods, necessary for running kube-system pods
	KubernetesMinMaxPods = 5
)

// vlabs default configuration
const (
	// DefaultNetworkPlugin defines the network plugin to use by default
	DefaultNetworkPlugin = "azure"
	// DefaultNetworkPluginWindows defines the network plugin to use by default for clusters with Windows agent pools
	DefaultNetworkPluginWindows = "azure"
	// DefaultNetworkPolicy defines the network policy to use by default
	DefaultNetworkPolicy = ""
	// NetworkPolicyCilium is the string expression for cilium network policy config option
	NetworkPolicyCilium = "cilium"
	// NetworkPluginCilium is the string expression for cilium network policy config option
	NetworkPluginCilium = NetworkPolicyCilium
)

const (
	// AgentPoolProfileRoleEmpty is the empty role
	AgentPoolProfileRoleEmpty AgentPoolProfileRole = ""
	// AgentPoolProfileRoleInfra is the infra role
	AgentPoolProfileRoleInfra AgentPoolProfileRole = "infra"
)

const (
	// AzureStackCloud is a const string reference identifier for Azure Stack cloud
	AzureStackCloud = "AzureStackCloud"
)

const (
	//AzureADIdentitySystem Identity System
	AzureADIdentitySystem = "azure_ad"
	//ADFS Identity System
	ADFSIdentitySystem = "adfs"
)

const (
	//ClientSecretAuthMethod Authentication method
	ClientSecretAuthMethod = "client_secret"
	//ClientCertificateAuthMethod Authentication method
	ClientCertificateAuthMethod = "client_certificate"
)
