// Copyright 2018 Google Inc. All Rights Reserved.
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

// This code was autogenerated. Do not edit directly.

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	stable_v1alpha1 "agones.dev/agones/pkg/apis/stable/v1alpha1"
	versioned "agones.dev/agones/pkg/client/clientset/versioned"
	internalinterfaces "agones.dev/agones/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "agones.dev/agones/pkg/client/listers/stable/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GameServerInformer provides access to a shared informer and lister for
// GameServers.
type GameServerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GameServerLister
}

type gameServerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGameServerInformer constructs a new informer for GameServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGameServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGameServerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGameServerInformer constructs a new informer for GameServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGameServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StableV1alpha1().GameServers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StableV1alpha1().GameServers(namespace).Watch(options)
			},
		},
		&stable_v1alpha1.GameServer{},
		resyncPeriod,
		indexers,
	)
}

func (f *gameServerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGameServerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gameServerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&stable_v1alpha1.GameServer{}, f.defaultInformer)
}

func (f *gameServerInformer) Lister() v1alpha1.GameServerLister {
	return v1alpha1.NewGameServerLister(f.Informer().GetIndexer())
}
