package main_test

import (
	"fmt"
	"github.com/onsi/gomega/gexec"
	"os/exec"
	"sigs.k8s.io/kind/pkg/apis/config/defaults"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cmd"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var smbCsiDriverPath string
var session *gexec.Session
var smbBrokerCompiledPath string
var kubeConfigPath string
var nodeName string

func TestSmbCsiDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}

var _ = BeforeSuite(func() {
	SetDefaultEventuallyTimeout(10 * time.Minute)
	nodeName = "default-smb-csi-driver-test-node"
	kubeConfigPath = "/tmp/csi-kubeconfig"

	createK8sCluster()
})

var _ = AfterSuite(func() {
	provider := cluster.NewProvider(
		cluster.ProviderWithLogger(cmd.NewLogger()),
	)

	_ = provider.Delete(nodeName, kubeConfigPath)
})

func createK8sCluster() {
	provider := cluster.NewProvider(
		cluster.ProviderWithLogger(cmd.NewLogger()),
	)

	// Check if the cluster name already exists
	n, err := provider.ListNodes(nodeName)
	Expect(err).NotTo(HaveOccurred())
	Expect(n).To(HaveLen(0), "node(s) already exist for a cluster with the name "+nodeName)

	// create the cluster
	err = provider.Create(
		nodeName,
		cluster.CreateWithRawConfig([]byte(`kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches: 
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."registry:5000"]
    endpoint = ["http://registry:5000"]
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    apiVersion: kubeadm.k8s.io/v1beta2
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
        authorization-mode: "AlwaysAllow"
  extraMounts:                                                
    - hostPath: /tmp/                                         
      containerPath: /var/lib/kubelet/plugins/csi-smbplugin   
    - hostPath: /tmp/                                         
      containerPath: /var/lib/kubelet/plugins_registry   
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
  - containerPort: 443
    hostPort: 443`)),
		cluster.CreateWithNodeImage(defaults.Image), // There's a v1.13 image = kindest/node:v1.13.12
		cluster.CreateWithRetain(true),
		cluster.CreateWithWaitForReady(10*time.Minute),
		cluster.CreateWithKubeconfigPath(kubeConfigPath),
		cluster.CreateWithDisplayUsage(true),
		cluster.CreateWithDisplaySalutation(true),
	)
	Expect(err).NotTo(HaveOccurred())

	kubeContext := "kind-" + nodeName
	kubectl("cluster-info", "--context", kubeContext, "--kubeconfig", kubeConfigPath)
	kubectl("apply", "-f", "./assets/ingress-nginx")
	kubectl("patch", "deployments", "-n", "ingress-nginx", "nginx-ingress-controller", "-p", `{"spec":{"template":{"spec":{"containers":[{"name":"nginx-ingress-controller","ports":[{"containerPort":80,"hostPort":80},{"containerPort":443,"hostPort":443}]}],"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}`)
	runTestCommand("bash", "-c", "./assets/setup-local-registry.sh "+nodeName+"-control-plane")

	kubectl("apply", "--kustomize", "./base")
}

func kubectl(cmd ...string) string {
	return runTestCommand("kubectl", cmd...)
}

func runTestCommand(name string, cmds ...string) string {
	command := exec.Command(name, cmds...)
	command.Env = append(command.Env, "KUBECONFIG="+kubeConfigPath)
	fmt.Println(fmt.Sprintf("Running %v", command.Args))

	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0), string(session.Out.Contents()))
	return string(session.Out.Contents()) + string(session.Err.Contents())
}