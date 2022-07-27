// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	"time"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumIdentitiesGetter has a method to return a CiliumIdentityInterface.
// A group's client should implement this interface.
type CiliumIdentitiesGetter interface {
	CiliumIdentities() CiliumIdentityInterface
}

// CiliumIdentityInterface has methods to work with CiliumIdentity resources.
type CiliumIdentityInterface interface {
	Create(*v2.CiliumIdentity) (*v2.CiliumIdentity, error)
	Update(*v2.CiliumIdentity) (*v2.CiliumIdentity, error)
	UpdateStatus(*v2.CiliumIdentity) (*v2.CiliumIdentity, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.CiliumIdentity, error)
	List(opts v1.ListOptions) (*v2.CiliumIdentityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.CiliumIdentity, err error)
	CiliumIdentityExpansion
}

// ciliumIdentities implements CiliumIdentityInterface
type ciliumIdentities struct {
	client rest.Interface
}

// newCiliumIdentities returns a CiliumIdentities
func newCiliumIdentities(c *CiliumV2Client) *ciliumIdentities {
	return &ciliumIdentities{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumIdentity, and returns the corresponding ciliumIdentity object, and an error if there is any.
func (c *ciliumIdentities) Get(name string, options v1.GetOptions) (result *v2.CiliumIdentity, err error) {
	result = &v2.CiliumIdentity{}
	err = c.client.Get().
		Resource("ciliumidentities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumIdentities that match those selectors.
func (c *ciliumIdentities) List(opts v1.ListOptions) (result *v2.CiliumIdentityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.CiliumIdentityList{}
	err = c.client.Get().
		Resource("ciliumidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumIdentities.
func (c *ciliumIdentities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ciliumIdentity and creates it.  Returns the server's representation of the ciliumIdentity, and an error, if there is any.
func (c *ciliumIdentities) Create(ciliumIdentity *v2.CiliumIdentity) (result *v2.CiliumIdentity, err error) {
	result = &v2.CiliumIdentity{}
	err = c.client.Post().
		Resource("ciliumidentities").
		Body(ciliumIdentity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ciliumIdentity and updates it. Returns the server's representation of the ciliumIdentity, and an error, if there is any.
func (c *ciliumIdentities) Update(ciliumIdentity *v2.CiliumIdentity) (result *v2.CiliumIdentity, err error) {
	result = &v2.CiliumIdentity{}
	err = c.client.Put().
		Resource("ciliumidentities").
		Name(ciliumIdentity.Name).
		Body(ciliumIdentity).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ciliumIdentities) UpdateStatus(ciliumIdentity *v2.CiliumIdentity) (result *v2.CiliumIdentity, err error) {
	result = &v2.CiliumIdentity{}
	err = c.client.Put().
		Resource("ciliumidentities").
		Name(ciliumIdentity.Name).
		SubResource("status").
		Body(ciliumIdentity).
		Do().
		Into(result)
	return
}

// Delete takes name of the ciliumIdentity and deletes it. Returns an error if one occurs.
func (c *ciliumIdentities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumidentities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumIdentities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumidentities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ciliumIdentity.
func (c *ciliumIdentities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.CiliumIdentity, err error) {
	result = &v2.CiliumIdentity{}
	err = c.client.Patch(pt).
		Resource("ciliumidentities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
