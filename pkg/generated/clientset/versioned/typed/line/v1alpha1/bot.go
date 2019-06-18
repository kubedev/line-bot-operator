/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/kairen/line-bot-operator/pkg/apis/line/v1alpha1"
	scheme "github.com/kairen/line-bot-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BotsGetter has a method to return a BotInterface.
// A group's client should implement this interface.
type BotsGetter interface {
	Bots(namespace string) BotInterface
}

// BotInterface has methods to work with Bot resources.
type BotInterface interface {
	Create(*v1alpha1.Bot) (*v1alpha1.Bot, error)
	Update(*v1alpha1.Bot) (*v1alpha1.Bot, error)
	UpdateStatus(*v1alpha1.Bot) (*v1alpha1.Bot, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Bot, error)
	List(opts v1.ListOptions) (*v1alpha1.BotList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Bot, err error)
	BotExpansion
}

// bots implements BotInterface
type bots struct {
	client rest.Interface
	ns     string
}

// newBots returns a Bots
func newBots(c *LineV1alpha1Client, namespace string) *bots {
	return &bots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bot, and returns the corresponding bot object, and an error if there is any.
func (c *bots) Get(name string, options v1.GetOptions) (result *v1alpha1.Bot, err error) {
	result = &v1alpha1.Bot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Bots that match those selectors.
func (c *bots) List(opts v1.ListOptions) (result *v1alpha1.BotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bots.
func (c *bots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a bot and creates it.  Returns the server's representation of the bot, and an error, if there is any.
func (c *bots) Create(bot *v1alpha1.Bot) (result *v1alpha1.Bot, err error) {
	result = &v1alpha1.Bot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bots").
		Body(bot).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bot and updates it. Returns the server's representation of the bot, and an error, if there is any.
func (c *bots) Update(bot *v1alpha1.Bot) (result *v1alpha1.Bot, err error) {
	result = &v1alpha1.Bot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bots").
		Name(bot.Name).
		Body(bot).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bots) UpdateStatus(bot *v1alpha1.Bot) (result *v1alpha1.Bot, err error) {
	result = &v1alpha1.Bot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bots").
		Name(bot.Name).
		SubResource("status").
		Body(bot).
		Do().
		Into(result)
	return
}

// Delete takes name of the bot and deletes it. Returns an error if one occurs.
func (c *bots) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bots").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bots").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bot.
func (c *bots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Bot, err error) {
	result = &v1alpha1.Bot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bots").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
