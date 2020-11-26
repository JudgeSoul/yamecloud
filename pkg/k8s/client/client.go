package client

import (
	"github.com/yametech/yamecloud/pkg/k8s/types"
	"k8s.io/client-go/dynamic"
	dynamicInformer "k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

const (
	// High enough QPS to fit all expected use cases.
	qps = 1e6
	// High enough Burst to fit all expected use cases.
	burst = 1e6
	// full sync cache resource time
	period = 30 * time.Second
)

var SharedCacheInformerFactory *CacheInformerFactory

func BuildClientSet(path string) (client.Interface, *rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, nil, err
	}
	cli, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	return cli, config, nil
}

func buildDynamicClientFromRest(clientCfg *rest.Config) (client.Interface, error) {
	clientCfg.QPS = qps
	clientCfg.Burst = burst
	dynClient, err := client.NewForConfig(clientCfg)
	if err != nil {
		return nil, err
	}
	return dynClient, nil
}

type CacheInformerFactory struct {
	Interface dynamic.Interface
	Informer  dynamicInformer.DynamicSharedInformerFactory
	stopChan  chan struct{}
}

func NewCacheInformerFactory(resLister types.ResourceLister, restConf *rest.Config) (*CacheInformerFactory, error) {

	if SharedCacheInformerFactory != nil {
		return SharedCacheInformerFactory, nil
	}

	client, err := buildDynamicClientFromRest(restConf)
	if err != nil {
		return nil, err
	}

	stop := make(chan struct{})
	sharedInformerFactory := dynamicInformer.NewDynamicSharedInformerFactory(client, period)

	resLister.Ranges(sharedInformerFactory, stop)

	sharedInformerFactory.Start(stop)

	SharedCacheInformerFactory =
		&CacheInformerFactory{
			client,
			sharedInformerFactory,
			stop,
		}

	return SharedCacheInformerFactory, nil
}

func CreateInClusterConfig() (*kubernetes.Clientset, *rest.Config, error) {
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, nil, err
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, nil, err
	}
	return clientSet, restConfig, nil
}
