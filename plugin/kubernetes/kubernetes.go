package kubernetes

import (
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Kubernetes struct {
	configPath string
	kubeConfig *rest.Config
	ClientSet  *kubernetes.Clientset
}

func NewKubernetes() *Kubernetes {
	return &Kubernetes{}
}

func (k *Kubernetes) GetConfigPath() string {
	return k.configPath
}

func (k *Kubernetes) GenerateConfigPath() {

	// generating config file path
	homeDirectory := homedir.HomeDir()

	configPath := filepath.Join(homeDirectory, ".kube", "config")
	log.Printf("Using kubeconfig: %s\n", configPath)

	k.configPath = configPath
}

// generating kubeconfig from current context
func (k *Kubernetes) ConfigureKubeConfig() error {
	var err error

	//building the config
	k.kubeConfig, err = clientcmd.BuildConfigFromFlags("", k.configPath)
	if err != nil {
		return err
	}
	return nil
}

// generating kubeconfig from given context
func (k *Kubernetes) ConfigWithKubeContext(kubeContext string) error {
	var err error
	k.kubeConfig, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: k.configPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: kubeContext,
		}).ClientConfig()
	if err != nil {
		return err
	}
	return nil
}

func (k *Kubernetes) ConfigureClientSet() error {

	// initializing client set
	clientSet, err := kubernetes.NewForConfig(k.kubeConfig)
	if err != nil {
		return err
	}
	k.ClientSet = clientSet
	return nil
}
