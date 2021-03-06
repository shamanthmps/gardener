// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	garden_v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	versioned "github.com/gardener/gardener/pkg/client/garden/clientset/versioned"
	internalinterfaces "github.com/gardener/gardener/pkg/client/garden/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/gardener/gardener/pkg/client/garden/listers/garden/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BackupInfrastructureInformer provides access to a shared informer and lister for
// BackupInfrastructures.
type BackupInfrastructureInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.BackupInfrastructureLister
}

type backupInfrastructureInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBackupInfrastructureInformer constructs a new informer for BackupInfrastructure type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBackupInfrastructureInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBackupInfrastructureInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBackupInfrastructureInformer constructs a new informer for BackupInfrastructure type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBackupInfrastructureInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GardenV1beta1().BackupInfrastructures(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GardenV1beta1().BackupInfrastructures(namespace).Watch(options)
			},
		},
		&garden_v1beta1.BackupInfrastructure{},
		resyncPeriod,
		indexers,
	)
}

func (f *backupInfrastructureInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBackupInfrastructureInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *backupInfrastructureInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&garden_v1beta1.BackupInfrastructure{}, f.defaultInformer)
}

func (f *backupInfrastructureInformer) Lister() v1beta1.BackupInfrastructureLister {
	return v1beta1.NewBackupInfrastructureLister(f.Informer().GetIndexer())
}
