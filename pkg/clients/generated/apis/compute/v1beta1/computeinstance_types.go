// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstanceAccessConfig struct {
	/*  */
	// +optional
	NatIpRef *v1alpha1.ResourceRef `json:"natIpRef,omitempty"`

	/* The networking tier used for configuring this instance. One of PREMIUM or STANDARD. */
	// +optional
	NetworkTier *string `json:"networkTier,omitempty"`

	/* The DNS domain name for the public PTR record. */
	// +optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`
}

type InstanceAdvancedMachineFeatures struct {
	/* Whether to enable nested virtualization or not. */
	// +optional
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`

	/* The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed. */
	// +optional
	ThreadsPerCore *int `json:"threadsPerCore,omitempty"`

	/* The number of physical cores to expose to an instance. Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance\'s nominal CPU count and the underlying platform\'s SMT width. */
	// +optional
	VisibleCoreCount *int `json:"visibleCoreCount,omitempty"`
}

type InstanceAliasIpRange struct {
	/* The IP CIDR range represented by this alias IP range. */
	IpCidrRange string `json:"ipCidrRange"`

	/* The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. */
	// +optional
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty"`
}

type InstanceAttachedDisk struct {
	/* Name with which the attached disk is accessible under /dev/disk/by-id/. */
	// +optional
	DeviceName *string `json:"deviceName,omitempty"`

	/* A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set. */
	// +optional
	DiskEncryptionKeyRaw *InstanceDiskEncryptionKeyRaw `json:"diskEncryptionKeyRaw,omitempty"`

	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource. */
	// +optional
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty"`

	/*  */
	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE". */
	// +optional
	Mode *string `json:"mode,omitempty"`

	/*  */
	SourceDiskRef v1alpha1.ResourceRef `json:"sourceDiskRef"`
}

type InstanceBootDisk struct {
	/* Immutable. Whether the disk will be auto-deleted when the instance is deleted. */
	// +optional
	AutoDelete *bool `json:"autoDelete,omitempty"`

	/* Immutable. Name with which attached disk will be accessible under /dev/disk/by-id/. */
	// +optional
	DeviceName *string `json:"deviceName,omitempty"`

	/* Immutable. A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set. */
	// +optional
	DiskEncryptionKeyRaw *InstanceDiskEncryptionKeyRaw `json:"diskEncryptionKeyRaw,omitempty"`

	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource. */
	// +optional
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty"`

	/* Immutable. Parameters with which a disk was created alongside the instance. */
	// +optional
	InitializeParams *InstanceInitializeParams `json:"initializeParams,omitempty"`

	/*  */
	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* Immutable. Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE". */
	// +optional
	Mode *string `json:"mode,omitempty"`

	/*  */
	// +optional
	SourceDiskRef *v1alpha1.ResourceRef `json:"sourceDiskRef,omitempty"`
}

type InstanceConfidentialInstanceConfig struct {
	/* Defines whether the instance should have confidential compute enabled. */
	EnableConfidentialCompute bool `json:"enableConfidentialCompute"`
}

type InstanceDiskEncryptionKeyRaw struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *InstanceValueFrom `json:"valueFrom,omitempty"`
}

type InstanceGuestAccelerator struct {
	/* Immutable. The number of the guest accelerator cards exposed to this instance. */
	Count int `json:"count"`

	/* Immutable. The accelerator type resource exposed to this instance. E.g. nvidia-tesla-k80. */
	Type string `json:"type"`
}

type InstanceInitializeParams struct {
	/* Immutable. A set of key/value label pairs assigned to the disk. */
	// +optional
	Labels *InstanceLabels `json:"labels,omitempty"`

	/* Immutable. The size of the image in gigabytes. */
	// +optional
	Size *int `json:"size,omitempty"`

	/*  */
	// +optional
	SourceImageRef *v1alpha1.ResourceRef `json:"sourceImageRef,omitempty"`

	/* Immutable. The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type InstanceIpv6AccessConfig struct {
	/* The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. The field is output only, an IPv6 address from a subnetwork associated with the instance will be allocated dynamically. */
	// +optional
	ExternalIpv6 *string `json:"externalIpv6,omitempty"`

	/* The prefix length of the external IPv6 range. */
	// +optional
	ExternalIpv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty"`

	/* The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6. */
	NetworkTier string `json:"networkTier"`

	/* The domain name to be used when creating DNSv6 records for the external IPv6 ranges. */
	// +optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty"`
}

type InstanceLabels struct {
}

type InstanceMetadata struct {
	/*  */
	Key string `json:"key"`

	/*  */
	Value string `json:"value"`
}

type InstanceNetworkInterface struct {
	/* Access configurations, i.e. IPs via which this instance can be accessed via the Internet. */
	// +optional
	AccessConfig []InstanceAccessConfig `json:"accessConfig,omitempty"`

	/* An array of alias IP ranges for this network interface. */
	// +optional
	AliasIpRange []InstanceAliasIpRange `json:"aliasIpRange,omitempty"`

	/* An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access. */
	// +optional
	Ipv6AccessConfig []InstanceIpv6AccessConfig `json:"ipv6AccessConfig,omitempty"`

	/* One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork. */
	// +optional
	Ipv6AccessType *string `json:"ipv6AccessType,omitempty"`

	/* The name of the interface. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* DEPRECATED. Although this field is still available, there is limited support. We recommend that you use `spec.networkInterface.networkIpRef` instead. */
	// +optional
	NetworkIp *string `json:"networkIp,omitempty"`

	/*  */
	// +optional
	NetworkIpRef *v1alpha1.ResourceRef `json:"networkIpRef,omitempty"`

	/*  */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET. */
	// +optional
	NicType *string `json:"nicType,omitempty"`

	/* Immutable. The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified. */
	// +optional
	QueueCount *int `json:"queueCount,omitempty"`

	/* The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used. */
	// +optional
	StackType *string `json:"stackType,omitempty"`

	/* The project in which the subnetwork belongs. */
	// +optional
	SubnetworkProject *string `json:"subnetworkProject,omitempty"`

	/*  */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
}

type InstanceNetworkPerformanceConfig struct {
	/* Immutable. The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT. */
	TotalEgressBandwidthTier string `json:"totalEgressBandwidthTier"`
}

type InstanceNodeAffinities struct {
	/*  */
	// +optional
	Value *InstanceValue `json:"value,omitempty"`
}

type InstanceReservationAffinity struct {
	/* Immutable. Specifies the label selector for the reservation to use. */
	// +optional
	SpecificReservation *InstanceSpecificReservation `json:"specificReservation,omitempty"`

	/* Immutable. The type of reservation from which this instance can consume resources. */
	Type string `json:"type"`
}

type InstanceScheduling struct {
	/* Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user). */
	// +optional
	AutomaticRestart *bool `json:"automaticRestart,omitempty"`

	/* Specifies the action GCE should take when SPOT VM is preempted. */
	// +optional
	InstanceTerminationAction *string `json:"instanceTerminationAction,omitempty"`

	/*  */
	// +optional
	MinNodeCpus *int `json:"minNodeCpus,omitempty"`

	/*  */
	// +optional
	NodeAffinities []InstanceNodeAffinities `json:"nodeAffinities,omitempty"`

	/* Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,. */
	// +optional
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty"`

	/* Immutable. Whether the instance is preemptible. */
	// +optional
	Preemptible *bool `json:"preemptible,omitempty"`

	/* Immutable. Whether the instance is spot. If this is set as SPOT. */
	// +optional
	ProvisioningModel *string `json:"provisioningModel,omitempty"`
}

type InstanceScratchDisk struct {
	/* The disk interface used for attaching this disk. One of SCSI or NVME. */
	Interface string `json:"interface"`
}

type InstanceServiceAccount struct {
	/* A list of service scopes. */
	Scopes []string `json:"scopes"`

	/*  */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
}

type InstanceShieldedInstanceConfig struct {
	/* Whether integrity monitoring is enabled for the instance. */
	// +optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	/* Whether secure boot is enabled for the instance. */
	// +optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	/* Whether the instance uses vTPM. */
	// +optional
	EnableVtpm *bool `json:"enableVtpm,omitempty"`
}

type InstanceSpecificReservation struct {
	/* Immutable. Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value. */
	Key string `json:"key"`

	/* Immutable. Corresponds to the label values of a reservation resource. */
	Values []string `json:"values"`
}

type InstanceValue struct {
}

type InstanceValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type ComputeInstanceSpec struct {
	/* Controls for advanced machine-related behavior features. */
	// +optional
	AdvancedMachineFeatures *InstanceAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	/* List of disks attached to the instance. */
	// +optional
	AttachedDisk []InstanceAttachedDisk `json:"attachedDisk,omitempty"`

	/* Immutable. The boot disk for the instance. */
	// +optional
	BootDisk *InstanceBootDisk `json:"bootDisk,omitempty"`

	/* Whether sending and receiving of packets with non-matching source or destination IPs is allowed. */
	// +optional
	CanIpForward *bool `json:"canIpForward,omitempty"`

	/* Immutable. The Confidential VM config being used by the instance.  on_host_maintenance has to be set to TERMINATE or this will fail to create. */
	// +optional
	ConfidentialInstanceConfig *InstanceConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`

	/* Whether deletion protection is enabled on this instance. */
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty"`

	/* Immutable. A brief description of the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Desired status of the instance. Either "RUNNING" or "TERMINATED". */
	// +optional
	DesiredStatus *string `json:"desiredStatus,omitempty"`

	/* Whether the instance has virtual displays enabled. */
	// +optional
	EnableDisplay *bool `json:"enableDisplay,omitempty"`

	/* Immutable. List of the type and count of accelerator cards attached to the instance. */
	// +optional
	GuestAccelerator []InstanceGuestAccelerator `json:"guestAccelerator,omitempty"`

	/* Immutable. A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created. */
	// +optional
	Hostname *string `json:"hostname,omitempty"`

	/*  */
	// +optional
	InstanceTemplateRef *v1alpha1.ResourceRef `json:"instanceTemplateRef,omitempty"`

	/* The machine type to create. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/*  */
	// +optional
	Metadata []InstanceMetadata `json:"metadata,omitempty"`

	/* Immutable. Metadata startup scripts made available within the instance. */
	// +optional
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty"`

	/* The minimum CPU platform specified for the VM instance. */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. The networks attached to the instance. */
	// +optional
	NetworkInterface []InstanceNetworkInterface `json:"networkInterface,omitempty"`

	/* Immutable. Configures network performance settings for the instance. If not specified, the instance will be created with its default network performance configuration. */
	// +optional
	NetworkPerformanceConfig *InstanceNetworkPerformanceConfig `json:"networkPerformanceConfig,omitempty"`

	/* Immutable. Specifies the reservations that this instance can consume from. */
	// +optional
	ReservationAffinity *InstanceReservationAffinity `json:"reservationAffinity,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/*  */
	// +optional
	ResourcePolicies []v1alpha1.ResourceRef `json:"resourcePolicies,omitempty"`

	/* The scheduling strategy being used by the instance. */
	// +optional
	Scheduling *InstanceScheduling `json:"scheduling,omitempty"`

	/* Immutable. The scratch disks attached to the instance. */
	// +optional
	ScratchDisk []InstanceScratchDisk `json:"scratchDisk,omitempty"`

	/* The service account to attach to the instance. */
	// +optional
	ServiceAccount *InstanceServiceAccount `json:"serviceAccount,omitempty"`

	/* The shielded vm config being used by the instance. */
	// +optional
	ShieldedInstanceConfig *InstanceShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	/* The list of tags attached to the instance. */
	// +optional
	Tags []string `json:"tags,omitempty"`

	/* Immutable. The zone of the instance. If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used. */
	// +optional
	Zone *string `json:"zone,omitempty"`
}

type ComputeInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The CPU platform used by this instance. */
	CpuPlatform string `json:"cpuPlatform,omitempty"`
	/* Current status of the instance. */
	CurrentStatus string `json:"currentStatus,omitempty"`
	/* The server-assigned unique identifier of this instance. */
	InstanceId string `json:"instanceId,omitempty"`
	/* The unique fingerprint of the labels. */
	LabelFingerprint string `json:"labelFingerprint,omitempty"`
	/* The unique fingerprint of the metadata. */
	MetadataFingerprint string `json:"metadataFingerprint,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* The URI of the created resource. */
	SelfLink string `json:"selfLink,omitempty"`
	/* The unique fingerprint of the tags. */
	TagsFingerprint string `json:"tagsFingerprint,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInstance is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceSpec   `json:"spec,omitempty"`
	Status ComputeInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInstanceList contains a list of ComputeInstance
type ComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstance{}, &ComputeInstanceList{})
}
