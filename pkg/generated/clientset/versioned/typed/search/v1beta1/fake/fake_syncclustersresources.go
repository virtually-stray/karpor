/*
Copyright The Karbour Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/KusionStack/karbour/pkg/apis/search/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSyncClustersResourceses implements SyncClustersResourcesInterface
type FakeSyncClustersResourceses struct {
	Fake *FakeSearchV1beta1
}

var syncclustersresourcesesResource = schema.GroupVersionResource{Group: "search.karbour.com", Version: "v1beta1", Resource: "syncclustersresourceses"}

var syncclustersresourcesesKind = schema.GroupVersionKind{Group: "search.karbour.com", Version: "v1beta1", Kind: "SyncClustersResources"}

// Get takes name of the syncClustersResources, and returns the corresponding syncClustersResources object, and an error if there is any.
func (c *FakeSyncClustersResourceses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SyncClustersResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(syncclustersresourcesesResource, name), &v1beta1.SyncClustersResources{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SyncClustersResources), err
}

// List takes label and field selectors, and returns the list of SyncClustersResourceses that match those selectors.
func (c *FakeSyncClustersResourceses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SyncClustersResourcesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(syncclustersresourcesesResource, syncclustersresourcesesKind, opts), &v1beta1.SyncClustersResourcesList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.SyncClustersResourcesList{ListMeta: obj.(*v1beta1.SyncClustersResourcesList).ListMeta}
	for _, item := range obj.(*v1beta1.SyncClustersResourcesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested syncClustersResourceses.
func (c *FakeSyncClustersResourceses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(syncclustersresourcesesResource, opts))
}

// Create takes the representation of a syncClustersResources and creates it.  Returns the server's representation of the syncClustersResources, and an error, if there is any.
func (c *FakeSyncClustersResourceses) Create(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.CreateOptions) (result *v1beta1.SyncClustersResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(syncclustersresourcesesResource, syncClustersResources), &v1beta1.SyncClustersResources{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SyncClustersResources), err
}

// Update takes the representation of a syncClustersResources and updates it. Returns the server's representation of the syncClustersResources, and an error, if there is any.
func (c *FakeSyncClustersResourceses) Update(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.UpdateOptions) (result *v1beta1.SyncClustersResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(syncclustersresourcesesResource, syncClustersResources), &v1beta1.SyncClustersResources{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SyncClustersResources), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSyncClustersResourceses) UpdateStatus(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.UpdateOptions) (*v1beta1.SyncClustersResources, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(syncclustersresourcesesResource, "status", syncClustersResources), &v1beta1.SyncClustersResources{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SyncClustersResources), err
}

// Delete takes name of the syncClustersResources and deletes it. Returns an error if one occurs.
func (c *FakeSyncClustersResourceses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(syncclustersresourcesesResource, name, opts), &v1beta1.SyncClustersResources{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSyncClustersResourceses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(syncclustersresourcesesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.SyncClustersResourcesList{})
	return err
}

// Patch applies the patch and returns the patched syncClustersResources.
func (c *FakeSyncClustersResourceses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SyncClustersResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(syncclustersresourcesesResource, name, pt, data, subresources...), &v1beta1.SyncClustersResources{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SyncClustersResources), err
}