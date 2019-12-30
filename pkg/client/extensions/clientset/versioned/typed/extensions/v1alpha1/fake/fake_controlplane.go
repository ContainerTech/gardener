/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeControlPlanes implements ControlPlaneInterface
type FakeControlPlanes struct {
	Fake *FakeExtensionsV1alpha1
	ns   string
}

var controlplanesResource = schema.GroupVersionResource{Group: "extensions.gardener.cloud", Version: "v1alpha1", Resource: "controlplanes"}

var controlplanesKind = schema.GroupVersionKind{Group: "extensions.gardener.cloud", Version: "v1alpha1", Kind: "ControlPlane"}

// Get takes name of the controlPlane, and returns the corresponding controlPlane object, and an error if there is any.
func (c *FakeControlPlanes) Get(name string, options v1.GetOptions) (result *v1alpha1.ControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(controlplanesResource, c.ns, name), &v1alpha1.ControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlane), err
}

// List takes label and field selectors, and returns the list of ControlPlanes that match those selectors.
func (c *FakeControlPlanes) List(opts v1.ListOptions) (result *v1alpha1.ControlPlaneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(controlplanesResource, controlplanesKind, c.ns, opts), &v1alpha1.ControlPlaneList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ControlPlaneList{ListMeta: obj.(*v1alpha1.ControlPlaneList).ListMeta}
	for _, item := range obj.(*v1alpha1.ControlPlaneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested controlPlanes.
func (c *FakeControlPlanes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(controlplanesResource, c.ns, opts))

}

// Create takes the representation of a controlPlane and creates it.  Returns the server's representation of the controlPlane, and an error, if there is any.
func (c *FakeControlPlanes) Create(controlPlane *v1alpha1.ControlPlane) (result *v1alpha1.ControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(controlplanesResource, c.ns, controlPlane), &v1alpha1.ControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlane), err
}

// Update takes the representation of a controlPlane and updates it. Returns the server's representation of the controlPlane, and an error, if there is any.
func (c *FakeControlPlanes) Update(controlPlane *v1alpha1.ControlPlane) (result *v1alpha1.ControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(controlplanesResource, c.ns, controlPlane), &v1alpha1.ControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlane), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeControlPlanes) UpdateStatus(controlPlane *v1alpha1.ControlPlane) (*v1alpha1.ControlPlane, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(controlplanesResource, "status", c.ns, controlPlane), &v1alpha1.ControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlane), err
}

// Delete takes name of the controlPlane and deletes it. Returns an error if one occurs.
func (c *FakeControlPlanes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(controlplanesResource, c.ns, name), &v1alpha1.ControlPlane{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeControlPlanes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(controlplanesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ControlPlaneList{})
	return err
}

// Patch applies the patch and returns the patched controlPlane.
func (c *FakeControlPlanes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(controlplanesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ControlPlane), err
}
