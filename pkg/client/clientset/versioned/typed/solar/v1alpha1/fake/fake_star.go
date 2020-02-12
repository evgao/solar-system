/*
 Copyright 2020 The Knative Authors

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "my.dev/solar-system/pkg/apis/solar/v1alpha1"
)

// FakeStars implements StarInterface
type FakeStars struct {
	Fake *FakeExampleV1alpha1
	ns   string
}

var starsResource = schema.GroupVersionResource{Group: "example.crd.com", Version: "v1alpha1", Resource: "stars"}

var starsKind = schema.GroupVersionKind{Group: "example.crd.com", Version: "v1alpha1", Kind: "Star"}

// Get takes name of the star, and returns the corresponding star object, and an error if there is any.
func (c *FakeStars) Get(name string, options v1.GetOptions) (result *v1alpha1.Star, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(starsResource, c.ns, name), &v1alpha1.Star{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Star), err
}

// List takes label and field selectors, and returns the list of Stars that match those selectors.
func (c *FakeStars) List(opts v1.ListOptions) (result *v1alpha1.StarList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(starsResource, starsKind, c.ns, opts), &v1alpha1.StarList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StarList{ListMeta: obj.(*v1alpha1.StarList).ListMeta}
	for _, item := range obj.(*v1alpha1.StarList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stars.
func (c *FakeStars) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(starsResource, c.ns, opts))

}

// Create takes the representation of a star and creates it.  Returns the server's representation of the star, and an error, if there is any.
func (c *FakeStars) Create(star *v1alpha1.Star) (result *v1alpha1.Star, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(starsResource, c.ns, star), &v1alpha1.Star{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Star), err
}

// Update takes the representation of a star and updates it. Returns the server's representation of the star, and an error, if there is any.
func (c *FakeStars) Update(star *v1alpha1.Star) (result *v1alpha1.Star, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(starsResource, c.ns, star), &v1alpha1.Star{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Star), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStars) UpdateStatus(star *v1alpha1.Star) (*v1alpha1.Star, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(starsResource, "status", c.ns, star), &v1alpha1.Star{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Star), err
}

// Delete takes name of the star and deletes it. Returns an error if one occurs.
func (c *FakeStars) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(starsResource, c.ns, name), &v1alpha1.Star{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStars) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(starsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StarList{})
	return err
}

// Patch applies the patch and returns the patched star.
func (c *FakeStars) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Star, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(starsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Star{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Star), err
}