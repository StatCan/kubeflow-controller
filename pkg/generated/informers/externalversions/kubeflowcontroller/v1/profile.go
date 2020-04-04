/*
Copyright 2020 Statistics Canada

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	kubeflowcontrollerv1 "k8s.io/kubeflow-controller/pkg/apis/kubeflowcontroller/v1"
	versioned "k8s.io/kubeflow-controller/pkg/generated/clientset/versioned"
	internalinterfaces "k8s.io/kubeflow-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "k8s.io/kubeflow-controller/pkg/generated/listers/kubeflowcontroller/v1"
)

// ProfileInformer provides access to a shared informer and lister for
// Profiles.
type ProfileInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProfileLister
}

type profileInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewProfileInformer constructs a new informer for Profile type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProfileInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProfileInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredProfileInformer constructs a new informer for Profile type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProfileInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().Profiles().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().Profiles().Watch(context.TODO(), options)
			},
		},
		&kubeflowcontrollerv1.Profile{},
		resyncPeriod,
		indexers,
	)
}

func (f *profileInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProfileInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *profileInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeflowcontrollerv1.Profile{}, f.defaultInformer)
}

func (f *profileInformer) Lister() v1.ProfileLister {
	return v1.NewProfileLister(f.Informer().GetIndexer())
}
