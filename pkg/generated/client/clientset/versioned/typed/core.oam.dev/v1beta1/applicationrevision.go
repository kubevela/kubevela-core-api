/*
Copyright 2021 The KubeVela Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	scheme "github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationRevisionsGetter has a method to return a ApplicationRevisionInterface.
// A group's client should implement this interface.
type ApplicationRevisionsGetter interface {
	ApplicationRevisions(namespace string) ApplicationRevisionInterface
}

// ApplicationRevisionInterface has methods to work with ApplicationRevision resources.
type ApplicationRevisionInterface interface {
	Create(ctx context.Context, applicationRevision *v1beta1.ApplicationRevision, opts v1.CreateOptions) (*v1beta1.ApplicationRevision, error)
	Update(ctx context.Context, applicationRevision *v1beta1.ApplicationRevision, opts v1.UpdateOptions) (*v1beta1.ApplicationRevision, error)
	UpdateStatus(ctx context.Context, applicationRevision *v1beta1.ApplicationRevision, opts v1.UpdateOptions) (*v1beta1.ApplicationRevision, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ApplicationRevision, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ApplicationRevisionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ApplicationRevision, err error)
	ApplicationRevisionExpansion
}

// applicationRevisions implements ApplicationRevisionInterface
type applicationRevisions struct {
	client rest.Interface
	ns     string
}

// newApplicationRevisions returns a ApplicationRevisions
func newApplicationRevisions(c *CoreV1beta1Client, namespace string) *applicationRevisions {
	return &applicationRevisions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationRevision, and returns the corresponding applicationRevision object, and an error if there is any.
func (c *applicationRevisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ApplicationRevision, err error) {
	result = &v1beta1.ApplicationRevision{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationrevisions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationRevisions that match those selectors.
func (c *applicationRevisions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ApplicationRevisionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ApplicationRevisionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationRevisions.
func (c *applicationRevisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a applicationRevision and creates it.  Returns the server's representation of the applicationRevision, and an error, if there is any.
func (c *applicationRevisions) Create(ctx context.Context, applicationRevision *v1beta1.ApplicationRevision, opts v1.CreateOptions) (result *v1beta1.ApplicationRevision, err error) {
	result = &v1beta1.ApplicationRevision{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationRevision).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a applicationRevision and updates it. Returns the server's representation of the applicationRevision, and an error, if there is any.
func (c *applicationRevisions) Update(ctx context.Context, applicationRevision *v1beta1.ApplicationRevision, opts v1.UpdateOptions) (result *v1beta1.ApplicationRevision, err error) {
	result = &v1beta1.ApplicationRevision{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationrevisions").
		Name(applicationRevision.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationRevision).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *applicationRevisions) UpdateStatus(ctx context.Context, applicationRevision *v1beta1.ApplicationRevision, opts v1.UpdateOptions) (result *v1beta1.ApplicationRevision, err error) {
	result = &v1beta1.ApplicationRevision{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationrevisions").
		Name(applicationRevision.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationRevision).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the applicationRevision and deletes it. Returns an error if one occurs.
func (c *applicationRevisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationrevisions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationRevisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationrevisions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched applicationRevision.
func (c *applicationRevisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ApplicationRevision, err error) {
	result = &v1beta1.ApplicationRevision{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationrevisions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}