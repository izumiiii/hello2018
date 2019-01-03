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

package versioned

import (
	samplecontrollerv1alpha "github.com/takaishi/hello2018/hello-custom-resource/my-sample-controller/pkg/client/clientset/versioned/typed/foo/v1alpha"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	SamplecontrollerV1alpha() samplecontrollerv1alpha.SamplecontrollerV1alphaInterface
	// Deprecated: please explicitly pick a version if possible.
	Samplecontroller() samplecontrollerv1alpha.SamplecontrollerV1alphaInterface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	samplecontrollerV1alpha *samplecontrollerv1alpha.SamplecontrollerV1alphaClient
}

// SamplecontrollerV1alpha retrieves the SamplecontrollerV1alphaClient
func (c *Clientset) SamplecontrollerV1alpha() samplecontrollerv1alpha.SamplecontrollerV1alphaInterface {
	return c.samplecontrollerV1alpha
}

// Deprecated: Samplecontroller retrieves the default version of SamplecontrollerClient.
// Please explicitly pick a version.
func (c *Clientset) Samplecontroller() samplecontrollerv1alpha.SamplecontrollerV1alphaInterface {
	return c.samplecontrollerV1alpha
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.samplecontrollerV1alpha, err = samplecontrollerv1alpha.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.samplecontrollerV1alpha = samplecontrollerv1alpha.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.samplecontrollerV1alpha = samplecontrollerv1alpha.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}