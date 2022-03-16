/*
Copyright 2021 TriggerMesh Inc.

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
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/typed/flow/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeFlowV1alpha1 struct {
	*testing.Fake
}

func (c *FakeFlowV1alpha1) JQTransformations(namespace string) v1alpha1.JQTransformationInterface {
	return &FakeJQTransformations{c, namespace}
}

func (c *FakeFlowV1alpha1) Synchronizers(namespace string) v1alpha1.SynchronizerInterface {
	return &FakeSynchronizers{c, namespace}
}

func (c *FakeFlowV1alpha1) Transformations(namespace string) v1alpha1.TransformationInterface {
	return &FakeTransformations{c, namespace}
}

func (c *FakeFlowV1alpha1) XMLToJSONTransformations(namespace string) v1alpha1.XMLToJSONTransformationInterface {
	return &FakeXMLToJSONTransformations{c, namespace}
}

func (c *FakeFlowV1alpha1) XSLTTransformations(namespace string) v1alpha1.XSLTTransformationInterface {
	return &FakeXSLTTransformations{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFlowV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
