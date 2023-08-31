/*
Copyright 2023 The KServe Authors.

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

package v1alpha1

import (
	"context"
	time "time"

	servingv1alpha1 "github.com/kserve/kserve/pkg/apis/serving/v1alpha1"
	versioned "github.com/kserve/kserve/pkg/clientv1alpha1/clientset/versioned"
	internalinterfaces "github.com/kserve/kserve/pkg/clientv1alpha1/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kserve/kserve/pkg/clientv1alpha1/listers/serving/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterServingRuntimeInformer provides access to a shared informer and lister for
// ClusterServingRuntimes.
type ClusterServingRuntimeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterServingRuntimeLister
}

type clusterServingRuntimeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterServingRuntimeInformer constructs a new informer for ClusterServingRuntime type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterServingRuntimeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterServingRuntimeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterServingRuntimeInformer constructs a new informer for ClusterServingRuntime type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterServingRuntimeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha1().ClusterServingRuntimes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha1().ClusterServingRuntimes(namespace).Watch(context.TODO(), options)
			},
		},
		&servingv1alpha1.ClusterServingRuntime{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterServingRuntimeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterServingRuntimeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterServingRuntimeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servingv1alpha1.ClusterServingRuntime{}, f.defaultInformer)
}

func (f *clusterServingRuntimeInformer) Lister() v1alpha1.ClusterServingRuntimeLister {
	return v1alpha1.NewClusterServingRuntimeLister(f.Informer().GetIndexer())
}
