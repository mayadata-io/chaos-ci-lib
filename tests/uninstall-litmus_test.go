package tests

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/mayadata-io/chaos-ci-lib/pkg"
	chaosTypes "github.com/mayadata-io/chaos-ci-lib/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes"
	scheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog"

	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	chaosClient "github.com/litmuschaos/chaos-operator/pkg/client/clientset/versioned/typed/litmuschaos/v1alpha1"
)

func TestUninstallLitmus(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "BDD test")
}

var _ = BeforeSuite(func() {
	var err error

	chaosTypes.Config, err = chaosTypes.GetKubeConfig()
	if err != nil {
		Expect(err).To(BeNil(), "Failed to get kubeconfig client")
	}
	chaosTypes.Client, err = kubernetes.NewForConfig(chaosTypes.Config)
	if err != nil {
		Expect(err).To(BeNil(), "failed to get client")
	}
	chaosTypes.ClientSet, err = chaosClient.NewForConfig(chaosTypes.Config)
	if err != nil {
		Expect(err).To(BeNil(), "failed to get clientSet")
	}
	err = v1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		fmt.Println(err)
	}

})

//BDD Tests to delete litmus
var _ = Describe("BDD of Litmus cleanup", func() {

	// BDD TEST CASE 1
	Context("Check for the Litmus components", func() {

		It("Should check for deletion of Litmus", func() {

			var err error
			//Deleting all chaosengines
			By("Deleting all chaosengine")
			err = exec.Command("kubectl", "delete", "chaosengine", "-n", pkg.GetEnv("APP_NS", "default"), "--all").Run()
			Expect(err).To(BeNil(), "Failed to delete chaosengine")
			klog.Info("All chaosengine deleted successfully")

			//Deleting all chaosexperiment
			By("Deleting all chaosexperiment")
			err = exec.Command("kubectl", "delete", "chaosexperiment", "-n", pkg.GetEnv("APP_NS", "default"), "--all").Run()
			Expect(err).To(BeNil(), "Failed to delete chaosexperiment")
			klog.Info("All chaosexperiment deleted successfully")

			//Deleting all chaosresults
			By("Deleting all chaosresults")
			err = exec.Command("kubectl", "delete", "chaosresult", "-n", pkg.GetEnv("APP_NS", "default"), "--all").Run()
			Expect(err).To(BeNil(), "Failed to delete chaosresult")
			klog.Info("All chaosresult deleted successfully")

			//Deleting crds
			By("Delete chaosengine crd")
			err = exec.Command("kubectl", "delete", "-f", chaosTypes.LitmusCrd).Run()
			Expect(err).To(BeNil(), "Failed to delete crds")
			klog.Info("Litmus crds deleted successfully")

			//Deleting litmus service account
			By("Delete Litmus service account")
			err = exec.Command("kubectl", "delete", "sa", "litmus", "-n", "litmus").Run()
			Expect(err).To(BeNil(), "Failed to delete litmus service account")
			klog.Info("Litmus service account deleted sucessfully")

			//Deleting litmus role
			By("Delete Litmus role")
			err = exec.Command("kubectl", "delete", "clusterrole", "litmus").Run()
			Expect(err).To(BeNil(), "Failed to delete litmus clusterrole")
			klog.Info("Litmus clusterrole deleted sucessfully")

			//Deleting litmus operator
			By("Delete Litmus operator")
			err = exec.Command("kubectl", "delete", "deploy", "chaos-operator-ce", "-n", "litmus").Run()
			Expect(err).To(BeNil(), "Failed to delete chaos operator")
			klog.Info("Litmus chaos operator deleted sucessfully")

			//Deleting litmus namespace
			By("Delete Litmus namespace")
			err = exec.Command("kubectl", "delete", "ns", "litmus").Run()
			Expect(err).To(BeNil(), "Failed to delete litmus namespace")
			klog.Info("Litmus namespace deleted sucessfully")

		})
	})
})
