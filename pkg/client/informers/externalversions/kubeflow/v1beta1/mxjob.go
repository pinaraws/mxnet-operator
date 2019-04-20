// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

// This file was automatically generated by informer-gen

package v1beta1

import (
	time "time"

	mxnetv1beta1 "github.com/pinaraws/mxnet-operator/pkg/apis/mxnet/v1beta1"
	versioned "github.com/pinaraws/mxnet-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/pinaraws/mxnet-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/pinaraws/mxnet-operator/pkg/client/listers/kubeflow/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MXJobInformer provides access to a shared informer and lister for
// MXJobs.
type MXJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.MXJobLister
}

type mXJobInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewMXJobInformer constructs a new informer for MXJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMXJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.KubeflowV1beta1().MXJobs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.KubeflowV1beta1().MXJobs(namespace).Watch(options)
			},
		},
		&mxnetv1beta1.MXJob{},
		resyncPeriod,
		indexers,
	)
}

func defaultMXJobInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewMXJobInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *mXJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mxnetv1beta1.MXJob{}, defaultMXJobInformer)
}

func (f *mXJobInformer) Lister() v1beta1.MXJobLister {
	return v1beta1.NewMXJobLister(f.Informer().GetIndexer())
}
