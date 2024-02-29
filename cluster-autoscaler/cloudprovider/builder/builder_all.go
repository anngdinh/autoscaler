//go:build !gce && !aws && !azure && !magnum && !digitalocean && !clusterapi && !linode && !equinixmetal && !oci && !vultr && !tencentcloud && !scaleway && !externalgrpc && !rancher && !volcengine && !cloudstack
// +build !gce,!aws,!azure,!magnum,!digitalocean,!clusterapi,!linode,!equinixmetal,!oci,!vultr,!tencentcloud,!scaleway,!externalgrpc,!rancher,!volcengine,!cloudstack

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package builder

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/azure"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/cloudstack"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/clusterapi"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/digitalocean"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/equinixmetal"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/externalgrpc"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/gce"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/linode"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/magnum"
	oci "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/instancepools"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/rancher"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/scaleway"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/tencentcloud"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/vultr"
	"k8s.io/autoscaler/cluster-autoscaler/config"
	"k8s.io/client-go/informers"
)

// AvailableCloudProviders supported by the cloud provider builder.
var AvailableCloudProviders = []string{
	cloudprovider.AwsProviderName,
	cloudprovider.AzureProviderName,
	cloudprovider.GceProviderName,
	cloudprovider.CloudStackProviderName,
	cloudprovider.MagnumProviderName,
	cloudprovider.DigitalOceanProviderName,
	cloudprovider.ExternalGrpcProviderName,
	cloudprovider.OracleCloudProviderName,
	cloudprovider.ClusterAPIProviderName,
	cloudprovider.LinodeProviderName,
	cloudprovider.EquinixMetalProviderName,
	cloudprovider.VultrProviderName,
	cloudprovider.TencentcloudProviderName,
	cloudprovider.ScalewayProviderName,
	cloudprovider.RancherProviderName,
	cloudprovider.VolcengineProviderName,
}

// DefaultCloudProvider is GCE.
const DefaultCloudProvider = cloudprovider.GceProviderName

func buildCloudProvider(opts config.AutoscalingOptions,
	do cloudprovider.NodeGroupDiscoveryOptions,
	rl *cloudprovider.ResourceLimiter,
	informerFactory informers.SharedInformerFactory) cloudprovider.CloudProvider {
	switch opts.CloudProviderName {
	case cloudprovider.GceProviderName:
		return gce.BuildGCE(opts, do, rl)
	case cloudprovider.AwsProviderName:
		return aws.BuildAWS(opts, do, rl)
	case cloudprovider.AzureProviderName:
		return azure.BuildAzure(opts, do, rl)
	case cloudprovider.CloudStackProviderName:
		return cloudstack.BuildCloudStack(opts, do, rl)
	case cloudprovider.DigitalOceanProviderName:
		return digitalocean.BuildDigitalOcean(opts, do, rl)
	case cloudprovider.ExternalGrpcProviderName:
		return externalgrpc.BuildExternalGrpc(opts, do, rl)
	case cloudprovider.MagnumProviderName:
		return magnum.BuildMagnum(opts, do, rl)
	case cloudprovider.PacketProviderName, cloudprovider.EquinixMetalProviderName:
		return equinixmetal.BuildCloudProvider(opts, do, rl)
	case cloudprovider.ClusterAPIProviderName:
		return clusterapi.BuildClusterAPI(opts, do, rl)
	case cloudprovider.LinodeProviderName:
		return linode.BuildLinode(opts, do, rl)
	case cloudprovider.OracleCloudProviderName:
		return oci.BuildOCI(opts, do, rl)
	case cloudprovider.VultrProviderName:
		return vultr.BuildVultr(opts, do, rl)
	case cloudprovider.TencentcloudProviderName:
		return tencentcloud.BuildTencentcloud(opts, do, rl)
	case cloudprovider.ScalewayProviderName:
		return scaleway.BuildScaleway(opts, do, rl)
	case cloudprovider.RancherProviderName:
		return rancher.BuildRancher(opts, do, rl)
	case cloudprovider.VolcengineProviderName:
		return volcengine.BuildVolcengine(opts, do, rl)
	}
	return nil
}
