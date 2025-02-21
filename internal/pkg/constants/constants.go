/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package constants

import (
	"time"

	"github.com/containerd/containerd/defaults"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
)

const (
	// KernelParamUserData is the kernel parameter name for specifying the URL
	// to the user data.
	KernelParamUserData = "talos.userdata"

	// KernelParamPlatform is the kernel parameter name for specifying the
	// platform.
	KernelParamPlatform = "talos.platform"

	// KernelParamHostname is the kernel parameter name for specifying the
	// hostname.
	KernelParamHostname = "talos.hostname"

	// KernelParamDefaultInterface is the kernel parameter for specifying the
	// initial interface used to bootstrap the node
	KernelParamDefaultInterface = "talos.interface"

	// KernelCurrentRoot is the kernel parameter name for specifying the
	// current root partition.
	KernelCurrentRoot = "talos.root"

	// NewRoot is the path where the switchroot target is mounted.
	NewRoot = "/root"

	// BootPartitionLabel is the label of the partition to use for mounting at
	// the boot path.
	BootPartitionLabel = "ESP"

	// BootMountPoint is the label of the partition to use for mounting at
	// the boot path.
	BootMountPoint = "/boot"

	// DataPartitionLabel is the label of the partition to use for mounting at
	// the data path.
	DataPartitionLabel = "DATA"

	// DataMountPoint is the label of the partition to use for mounting at
	// the data path.
	DataMountPoint = "/var"

	// RootMountPoint is the label of the partition to use for mounting at
	// the root path.
	RootMountPoint = "/"

	// ISOFilesystemLabel is the label of the ISO file system for the Talos
	// installer.
	ISOFilesystemLabel = "TALOS"

	// PATH defines all locations where executables are stored.
	PATH = "/sbin:/bin:/usr/sbin:/usr/bin:/usr/local/sbin:/usr/local/bin:/opt/cni/bin"

	// CNICalico is used to specify Calico CNI.
	CNICalico = "calico"

	// CNIFlannel is used to specify Flannel CNI.
	CNIFlannel = "flannel"

	// KubeadmConfig is the path to the kubeadm manifest file.
	KubeadmConfig = "/etc/kubernetes/kubeadm-config.yaml"

	// KubeadmCACert is the path to the root CA certificate.
	KubeadmCACert = v1beta1.DefaultCertificatesDir + "/" + constants.CACertName

	// KubeadmCAKey is the path to the root CA private key.
	KubeadmCAKey = v1beta1.DefaultCertificatesDir + "/" + constants.CAKeyName

	// KubeadmSACert is the path to the SA certificate.
	KubeadmSACert = v1beta1.DefaultCertificatesDir + "/" + constants.ServiceAccountPublicKeyName

	// KubeadmSAKey is the path to the SA private key.
	KubeadmSAKey = v1beta1.DefaultCertificatesDir + "/" + constants.ServiceAccountPrivateKeyName

	// KubeadmFrontProxyCACert is the path to the front proxy CA certificate.
	KubeadmFrontProxyCACert = v1beta1.DefaultCertificatesDir + "/" + constants.FrontProxyCACertName

	// KubeadmFrontProxyCAKey is the path to the front proxy CA private key.
	KubeadmFrontProxyCAKey = v1beta1.DefaultCertificatesDir + "/" + constants.FrontProxyCAKeyName

	// KubeadmEtcdCACert is the path to the etcd CA certificate.
	KubeadmEtcdCACert = v1beta1.DefaultCertificatesDir + "/" + constants.EtcdCACertName

	// KubeadmEtcdCAKey is the path to the etcd CA private key.
	KubeadmEtcdCAKey = v1beta1.DefaultCertificatesDir + "/" + constants.EtcdCAKeyName

	// KubeadmEtcdPeerCert is the path to the etcd CA certificate.
	KubeadmEtcdPeerCert = v1beta1.DefaultCertificatesDir + "/" + constants.EtcdPeerCertName

	// KubeadmEtcdPeerKey is the path to the etcd CA private key.
	KubeadmEtcdPeerKey = v1beta1.DefaultCertificatesDir + "/" + constants.EtcdPeerKeyName

	// KubernetesVersion is the enforced target version of the control plane.
	KubernetesVersion = "v1.15.0"

	// KubernetesImage is the enforced hyperkube image to use for the control plane.
	KubernetesImage = "k8s.gcr.io/hyperkube:" + KubernetesVersion

	// UserDataPath is the path to the downloaded user data.
	UserDataPath = "/var/userdata.yaml"

	// UserDataCIData is the volume label for NoCloud cloud-init.
	// See https://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html#datasource-nocloud.
	UserDataCIData = "cidata"

	// UserDataGuestInfo is the name of the VMware guestinfo user data strategy.
	UserDataGuestInfo = "guestinfo"

	// VMwareGuestInfoUserDataKey is the guestinfo key used to provide a user data file.
	VMwareGuestInfoUserDataKey = "talos.userdata"

	// AuditPolicyPathInitramfs is the path to the audit-policy.yaml relative to initramfs.
	AuditPolicyPathInitramfs = "/etc/kubernetes/audit-policy.yaml"

	// EncryptionConfigInitramfsPath is the path to the EncryptionConfig relative to initramfs.
	EncryptionConfigInitramfsPath = "/etc/kubernetes/encryptionconfig.yaml"

	// EncryptionConfigRootfsPath is the path to the EncryptionConfig relative to rootfs.
	EncryptionConfigRootfsPath = "/etc/kubernetes/encryptionconfig.yaml"

	// OsdPort is the port for the osd service.
	OsdPort = 50000

	// TrustdPort is the port for the trustd service.
	TrustdPort = 50001

	// SystemContainerdNamespace is the Containerd namespace for Talos services.
	SystemContainerdNamespace = "system"

	// TalosConfigEnvVar is the environment variable for setting the Talos configuration file path.
	TalosConfigEnvVar = "TALOSCONFIG"

	// InitSocketPath is the path to file socket of init API
	InitSocketPath = "/run/system/init/init.sock"

	// KernelAsset defines a well known name for our kernel filename
	KernelAsset = "vmlinuz"

	// InitramfsAsset defines a well known name for our initramfs filename
	InitramfsAsset = "initramfs.xz"

	// RootfsAsset defines a well known name for our rootfs filename
	RootfsAsset = "rootfs.sqsh"

	// NodeCertFile is the filename where the current Talos Node Certificate may be found
	NodeCertFile = "/var/talos-node.crt"

	// NodeCertRenewalInterval is the default interval at which Talos Node Certifications should be renewed
	NodeCertRenewalInterval = 24 * time.Hour
)

// See https://linux.die.net/man/3/klogctl
const (
	// SYSLOG_ACTION_SIZE_BUFFER is a named type argument to klogctl.
	// nolint: golint
	SYSLOG_ACTION_SIZE_BUFFER = 10

	// SYSLOG_ACTION_READ_ALL is a named type argument to klogctl.
	// nolint: golint
	SYSLOG_ACTION_READ_ALL = 3
)

// Containerd
const (
	ContainerdAddress = defaults.DefaultAddress
)
