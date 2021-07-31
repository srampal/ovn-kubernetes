package clusternetworkpolicy 

import (
	"time"
	"fmt"
	"k8s.io/klog/v2"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/client-go/tools/record"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	k8sclientset "k8s.io/client-go/kubernetes"
        okube "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/kube"
        cnpInformers "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/clusternetworkpolicy/v1/apis/informers/externalversions"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"

/***
        cnpclientset "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/clusternetworkpolicy/v1/apis/clientset/versioned"


	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/config"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/metrics"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/ovn/loadbalancer"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/util"
	"github.com/pkg/errors"

	discovery "k8s.io/api/discovery/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"

	coreinformers "k8s.io/client-go/informers/core/v1"
	discoveryinformers "k8s.io/client-go/informers/discovery/v1beta1"
	corelisters "k8s.io/client-go/listers/core/v1"
	discoverylisters "k8s.io/client-go/listers/discovery/v1beta1"


	utilnet "k8s.io/utils/net"
***/
)

const (
	// maxRetries is the number of times a object will be retried before it is dropped out of the queue.
	// With the current rate-limiter in use (5ms*2^(maxRetries-1)) the following numbers represent the
	// sequence of delays between successive queuings of an object.
	//
	// 5ms, 10ms, 20ms, 40ms, 80ms, 160ms, 320ms, 640ms, 1.3s, 2.6s, 5.1s, 10.2s, 20.4s, 41s, 82s
	maxRetries = 15

	controllerName = "ovn-cnp-controller"
)

// Main Cnp Controller struct  
type Controller struct {
/**
        // clientset for k8s core api groups/ resources
	k8sClient           clientset.Interface

        // clientset for cnp resources
        cnpClient           cnpclientset.Interface
**/

	// cnpSynced returns true if the shared informer has been synced at least once.
	cnpSynced cache.InformerSynced

	// Work queue of CNP updates that need to be processed. A channel is inappropriate here
	queue workqueue.RateLimitingInterface

	// workerLoopPeriod is the time between worker runs. The workers process the queue of changes.
	workerLoopPeriod time.Duration

        recorder record.EventRecorder
}

// NewController returns a new *Controller.
func NewController(ockube okube.Interface, 
                   k8sclient k8sclientset.Interface,
                   cnpInformerFactory cnpInformers.SharedInformerFactory,
                  ) *Controller {

	klog.V(4).Info("Creating new Cnp controller")

        if ockube == nil {
                fmt.Errorf("Error input starting new cnp controller ... ")
                return nil
        }

	broadcaster := record.NewBroadcaster()
	broadcaster.StartStructuredLogging(0)
	broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: k8sclient.CoreV1().Events("")})
	recorder := broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: controllerName})

	c := &Controller{
		queue:            workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerName),
		workerLoopPeriod: time.Second,
                recorder:         recorder,
	}

        // Handlers for CNP CRUD events

        klog.Info("Setting up event handlers for CNPs")

        cnpInformerFactory.K8s().V1().ClusterNetworkPolicies().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
                AddFunc:    c.onCnpAdd,
                UpdateFunc: nil,
                DeleteFunc: nil,
        })

	return c

} // NewController()


// handlers

// onCnpAdd queues the Cnp creation for processing.
func (c *Controller) onCnpAdd(obj interface{}) {
	klog.Infof("QQQQQ handling CNP Addition")
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		fmt.Errorf("couldn't get key for object %+v: %v", obj, err)
		utilruntime.HandleError(fmt.Errorf("couldn't get key for object %+v: %v", obj, err))
		return
	}
	klog.Infof("Queueing up Cluster Network Policy  %s", key)
	c.queue.Add(key)
}


// Run will not return until stopCh is closed. workers determines how many
// endpoints will be handled in parallel.
func (c *Controller) Run(workers int, stopCh <-chan struct{}, runRepair bool) error {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	klog.Infof("Starting Run of controller %s", controllerName)
	defer klog.Infof("Shutting down controller %s", controllerName)

/**
	// Wait for the caches to be synced
	klog.Info("Waiting for informer caches to sync")
	if !cache.WaitForNamedCacheSync(controllerName, stopCh, c.servicesSynced, c.endpointSlicesSynced) {
		return fmt.Errorf("error syncing cache")
	}

	if runRepair {
		// Run the repair controller only once
		// it keeps in sync Kubernetes and OVN
		// and handles removal of stale data on upgrades
		klog.Info("Remove stale OVN services")
		if err := c.repair.runOnce(); err != nil {
			klog.Errorf("Error repairing services: %v", err)
		}
	}
**/
	// Start the workers after the repair loop to avoid races
	klog.Info("Starting workers")
	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, c.workerLoopPeriod, stopCh)
	}

	<-stopCh
	return nil
}

// worker runs a worker thread that just dequeues items, processes them, and
// marks them done. You may run as many of these in parallel as you wish; the
// workqueue guarantees that they will not end up processing the same service
// at the same time.
func (c *Controller) worker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	klog.Infof(" HHHH Processing next CNP work item %s", controllerName)

	eKey, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(eKey)

	_ = c.syncCnps(eKey.(string))
/**
	c.handleErr(err, eKey)
**/

	return true
}


func (c *Controller) syncCnps(key string) error {

	klog.Infof("HHHH entered syncCnps with key %s ", key)

	startTime := time.Now()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
	        klog.Infof("HHHH error in SplitMetaNamespacw %s ", err.Error())
		return err
	}
	klog.Infof("HHHH Processing sync for name %s on namespace %s ", name, namespace)

	defer func() {
		klog.V(4).Infof("Finished syncing CNP %s on namespace %s : %v", name, namespace, time.Since(startTime))
	}()

        return nil
}
