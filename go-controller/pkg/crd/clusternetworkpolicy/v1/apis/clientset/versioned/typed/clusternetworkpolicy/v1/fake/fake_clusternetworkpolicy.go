// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	clusternetworkpolicyv1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/clusternetworkpolicy/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterNetworkPolicies implements ClusterNetworkPolicyInterface
type FakeClusterNetworkPolicies struct {
	Fake *FakeK8sV1
	ns   string
}

var clusternetworkpoliciesResource = schema.GroupVersionResource{Group: "k8s.ovn.org", Version: "v1", Resource: "clusternetworkpolicies"}

var clusternetworkpoliciesKind = schema.GroupVersionKind{Group: "k8s.ovn.org", Version: "v1", Kind: "ClusterNetworkPolicy"}

// Get takes name of the clusterNetworkPolicy, and returns the corresponding clusterNetworkPolicy object, and an error if there is any.
func (c *FakeClusterNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *clusternetworkpolicyv1.ClusterNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusternetworkpoliciesResource, c.ns, name), &clusternetworkpolicyv1.ClusterNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusternetworkpolicyv1.ClusterNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of ClusterNetworkPolicies that match those selectors.
func (c *FakeClusterNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *clusternetworkpolicyv1.ClusterNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusternetworkpoliciesResource, clusternetworkpoliciesKind, c.ns, opts), &clusternetworkpolicyv1.ClusterNetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clusternetworkpolicyv1.ClusterNetworkPolicyList{ListMeta: obj.(*clusternetworkpolicyv1.ClusterNetworkPolicyList).ListMeta}
	for _, item := range obj.(*clusternetworkpolicyv1.ClusterNetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterNetworkPolicies.
func (c *FakeClusterNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusternetworkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a clusterNetworkPolicy and creates it.  Returns the server's representation of the clusterNetworkPolicy, and an error, if there is any.
func (c *FakeClusterNetworkPolicies) Create(ctx context.Context, clusterNetworkPolicy *clusternetworkpolicyv1.ClusterNetworkPolicy, opts v1.CreateOptions) (result *clusternetworkpolicyv1.ClusterNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusternetworkpoliciesResource, c.ns, clusterNetworkPolicy), &clusternetworkpolicyv1.ClusterNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusternetworkpolicyv1.ClusterNetworkPolicy), err
}

// Update takes the representation of a clusterNetworkPolicy and updates it. Returns the server's representation of the clusterNetworkPolicy, and an error, if there is any.
func (c *FakeClusterNetworkPolicies) Update(ctx context.Context, clusterNetworkPolicy *clusternetworkpolicyv1.ClusterNetworkPolicy, opts v1.UpdateOptions) (result *clusternetworkpolicyv1.ClusterNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusternetworkpoliciesResource, c.ns, clusterNetworkPolicy), &clusternetworkpolicyv1.ClusterNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusternetworkpolicyv1.ClusterNetworkPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterNetworkPolicies) UpdateStatus(ctx context.Context, clusterNetworkPolicy *clusternetworkpolicyv1.ClusterNetworkPolicy, opts v1.UpdateOptions) (*clusternetworkpolicyv1.ClusterNetworkPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusternetworkpoliciesResource, "status", c.ns, clusterNetworkPolicy), &clusternetworkpolicyv1.ClusterNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusternetworkpolicyv1.ClusterNetworkPolicy), err
}

// Delete takes name of the clusterNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeClusterNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusternetworkpoliciesResource, c.ns, name), &clusternetworkpolicyv1.ClusterNetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusternetworkpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &clusternetworkpolicyv1.ClusterNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched clusterNetworkPolicy.
func (c *FakeClusterNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clusternetworkpolicyv1.ClusterNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusternetworkpoliciesResource, c.ns, name, pt, data, subresources...), &clusternetworkpolicyv1.ClusterNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusternetworkpolicyv1.ClusterNetworkPolicy), err
}
