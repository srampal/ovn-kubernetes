// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/clusternetworkpolicy/v1/apis/clientset/versioned/typed/clusternetworkpolicy/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeK8sV1 struct {
	*testing.Fake
}

func (c *FakeK8sV1) ClusterNetworkPolicies(namespace string) v1.ClusterNetworkPolicyInterface {
	return &FakeClusterNetworkPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeK8sV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
