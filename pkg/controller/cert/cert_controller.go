package cert

import (
	"context"
	"fmt"
	"time"

	certoperatorv1beta1 "github.com/fanfengqiang/cert-operator/pkg/apis/certoperator/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var logger = logf.Log.WithName("controller_cert")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Cert Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileCert{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("cert-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Cert
	err = c.Watch(&source.Kind{Type: &certoperatorv1beta1.Cert{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Secrets and requeue the owner Cert
	err = c.Watch(&source.Kind{Type: &corev1.Secret{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &certoperatorv1beta1.Cert{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileCert implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileCert{}

// ReconcileCert reconciles a Cert object
type ReconcileCert struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Cert object and makes changes based on the state read
// and what is in the Cert.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Secret as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileCert) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := logger.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Cert")

	// Fetch the Cert instance
	instance := &certoperatorv1beta1.Cert{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("Cert resource not found. Ignoring since object must be deleted")
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		reqLogger.Error(err, "Failed to get Cert")
		return reconcile.Result{}, err
	}

	// Check if this Secret already exists
	found := &corev1.Secret{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		dep, err := r.sercretForCert(instance)
		if err != nil {
			reqLogger.Error(err, "file to create a new cert!")
			return reconcile.Result{}, err
		}
		reqLogger.Info("Creating a new Secret", "Secret.Namespace", instance.Namespace, "Secret.Name", instance.Name)
		err = r.client.Create(context.TODO(), dep)
		if err != nil {
			reqLogger.Error(err, "Failed to create new Secret", "Secret.Namespace", dep.Namespace, "Secret.Name", dep.Name)
			return reconcile.Result{}, err
		}

		// Secret created successfully - don't requeue
		return reconcile.Result{RequeueAfter: time.Second * 10}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Secret")
		return reconcile.Result{}, err
	}

	// Ensure the secret is in validityPeriod

	loc, _ := time.LoadLocation("Local")
	formatTime, _ := time.ParseInLocation("2006-01-02-15-04-05", found.Annotations["updateTime"], loc)

	days := instance.Spec.ValidityPeriod - int(time.Now().Sub(formatTime).Minutes())
	if days < 0 {
		fmt.Println("renew status ValidityPeriod")
		dep, err := r.sercretForCert(instance)
		if err != nil {
			reqLogger.Error(err, "file to create a new cert!")
			return reconcile.Result{}, err
		}
		err = r.client.Update(context.TODO(), dep)
		if err != nil {
			reqLogger.Error(err, "Failed to update Secret", "Secret.Namespace", found.Namespace, "Deployment.Name", found.Name)
			return reconcile.Result{}, err
		}
		// Spec updated - return and requeue
		return reconcile.Result{Requeue: true}, nil
	}

	// Update RemainingValidDays if needed
	if days != instance.Status.RemainingValidDays {
		fmt.Println("renew status RemainingValidDays")
		instance.Status.RemainingValidDays = days
		err := r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update Cert status RemainingValidDays")
			return reconcile.Result{}, err
		}
	}

	// Update the Cret status with the secret createtime
	ctime := found.Annotations["updateTime"]
	// Update status if needed
	if ctime != instance.Status.SecretUpdateTime {
		fmt.Println("renew status SecretUpdateTime")
		instance.Status.SecretUpdateTime = ctime
		err := r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update Cert status SecretUpdateTime")
			return reconcile.Result{}, err
		}
	}

	// Secret already exists - don't requeue
	reqLogger.Info("Skip reconcile: Secret already exists", "secret.Namespace", found.Namespace, "secret.Name", found.Name)
	return reconcile.Result{RequeueAfter: time.Second * 10}, nil
}

func (r *ReconcileCert) sercretForCert(c *certoperatorv1beta1.Cert) (*corev1.Secret, error) {

	var aCert, err = CreateCert(c.Spec.Email, c.Spec.Domain, c.Spec.Provider, c.Spec.Envs)
	fmt.Println(c.Spec.Email)
	if err != nil {
		fmt.Println(err, "Failed to create a new Cert")
		return &corev1.Secret{}, err
	}

	labels := map[string]string{
		"certificateAuthority": "letsencrypt",
		"controller":           c.Name,
		"updateTime":           time.Unix(time.Now().Unix(), 0).Format("2006-01-02-15-04-05"),
	}
	fmt.Println("create a new secret")
	dep := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        c.Name,
			Namespace:   c.Namespace,
			Annotations: labels,
		},
		Data: map[string][]byte{
			"tls.crt": aCert["cert"],
			"tls.key": aCert["key"],
		},
	}
	// Set Cret instance as the owner and controller
	controllerutil.SetControllerReference(c, dep, r.scheme)
	return dep, nil
}
