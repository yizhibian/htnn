/*
Copyright The HTNN Authors.

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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1 "mosn.io/htnn/types/apis/v1"
)

// FakeDynamicConfigs implements DynamicConfigInterface
type FakeDynamicConfigs struct {
	Fake *FakeApisV1
	ns   string
}

var dynamicconfigsResource = v1.SchemeGroupVersion.WithResource("dynamicconfigs")

var dynamicconfigsKind = v1.SchemeGroupVersion.WithKind("DynamicConfig")

// Get takes name of the dynamicConfig, and returns the corresponding dynamicConfig object, and an error if there is any.
func (c *FakeDynamicConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DynamicConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dynamicconfigsResource, c.ns, name), &v1.DynamicConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DynamicConfig), err
}

// List takes label and field selectors, and returns the list of DynamicConfigs that match those selectors.
func (c *FakeDynamicConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DynamicConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dynamicconfigsResource, dynamicconfigsKind, c.ns, opts), &v1.DynamicConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DynamicConfigList{ListMeta: obj.(*v1.DynamicConfigList).ListMeta}
	for _, item := range obj.(*v1.DynamicConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dynamicConfigs.
func (c *FakeDynamicConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dynamicconfigsResource, c.ns, opts))

}

// Create takes the representation of a dynamicConfig and creates it.  Returns the server's representation of the dynamicConfig, and an error, if there is any.
func (c *FakeDynamicConfigs) Create(ctx context.Context, dynamicConfig *v1.DynamicConfig, opts metav1.CreateOptions) (result *v1.DynamicConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dynamicconfigsResource, c.ns, dynamicConfig), &v1.DynamicConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DynamicConfig), err
}

// Update takes the representation of a dynamicConfig and updates it. Returns the server's representation of the dynamicConfig, and an error, if there is any.
func (c *FakeDynamicConfigs) Update(ctx context.Context, dynamicConfig *v1.DynamicConfig, opts metav1.UpdateOptions) (result *v1.DynamicConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dynamicconfigsResource, c.ns, dynamicConfig), &v1.DynamicConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DynamicConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDynamicConfigs) UpdateStatus(ctx context.Context, dynamicConfig *v1.DynamicConfig, opts metav1.UpdateOptions) (*v1.DynamicConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dynamicconfigsResource, "status", c.ns, dynamicConfig), &v1.DynamicConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DynamicConfig), err
}

// Delete takes name of the dynamicConfig and deletes it. Returns an error if one occurs.
func (c *FakeDynamicConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dynamicconfigsResource, c.ns, name, opts), &v1.DynamicConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDynamicConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dynamicconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DynamicConfigList{})
	return err
}

// Patch applies the patch and returns the patched dynamicConfig.
func (c *FakeDynamicConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DynamicConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dynamicconfigsResource, c.ns, name, pt, data, subresources...), &v1.DynamicConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DynamicConfig), err
}
