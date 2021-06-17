// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/clusternetworkpolicy/v1/apis/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterNetworkPolicies returns a ClusterNetworkPolicyInformer.
	ClusterNetworkPolicies() ClusterNetworkPolicyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterNetworkPolicies returns a ClusterNetworkPolicyInformer.
func (v *version) ClusterNetworkPolicies() ClusterNetworkPolicyInformer {
	return &clusterNetworkPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
