/*
Copyright 2021 The Kruise Authors.

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
	context "context"

	appsv1alpha1 "github.com/openkruise/kruise/apis/apps/v1alpha1"
	scheme "github.com/openkruise/kruise/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SidecarSetsGetter has a method to return a SidecarSetInterface.
// A group's client should implement this interface.
type SidecarSetsGetter interface {
	SidecarSets() SidecarSetInterface
}

// SidecarSetInterface has methods to work with SidecarSet resources.
type SidecarSetInterface interface {
	Create(ctx context.Context, sidecarSet *appsv1alpha1.SidecarSet, opts v1.CreateOptions) (*appsv1alpha1.SidecarSet, error)
	Update(ctx context.Context, sidecarSet *appsv1alpha1.SidecarSet, opts v1.UpdateOptions) (*appsv1alpha1.SidecarSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, sidecarSet *appsv1alpha1.SidecarSet, opts v1.UpdateOptions) (*appsv1alpha1.SidecarSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*appsv1alpha1.SidecarSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*appsv1alpha1.SidecarSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appsv1alpha1.SidecarSet, err error)
	SidecarSetExpansion
}

// sidecarSets implements SidecarSetInterface
type sidecarSets struct {
	*gentype.ClientWithList[*appsv1alpha1.SidecarSet, *appsv1alpha1.SidecarSetList]
}

// newSidecarSets returns a SidecarSets
func newSidecarSets(c *AppsV1alpha1Client) *sidecarSets {
	return &sidecarSets{
		gentype.NewClientWithList[*appsv1alpha1.SidecarSet, *appsv1alpha1.SidecarSetList](
			"sidecarsets",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *appsv1alpha1.SidecarSet { return &appsv1alpha1.SidecarSet{} },
			func() *appsv1alpha1.SidecarSetList { return &appsv1alpha1.SidecarSetList{} },
		),
	}
}
