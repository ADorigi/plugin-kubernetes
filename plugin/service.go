package plugin

import (
	"context"

	"github.com/kaytu-io/kaytu/pkg/plugin/proto/src/golang"
	"github.com/kaytu-io/kaytu/pkg/plugin/sdk"
	"github.com/kaytu-io/plugin-kubernetes/plugin/kubernetes"
	"github.com/kaytu-io/plugin-kubernetes/plugin/kubernetes/pods"
	"github.com/kaytu-io/plugin-kubernetes/plugin/preferences"
	"github.com/kaytu-io/plugin-kubernetes/plugin/version"
)

type KubernetesPlugin struct {
	stream golang.Plugin_RegisterClient
	//processor processor2.Processor
}

func NewPlugin() *KubernetesPlugin {
	return &KubernetesPlugin{}
}

func (p *KubernetesPlugin) GetConfig() golang.RegisterConfig {
	return golang.RegisterConfig{
		Name:     "kaytu-io/plugin-kubernetes",
		Version:  version.VERSION,
		Provider: "aws",
		Commands: []*golang.Command{
			{
				Name:        "pods",
				Description: "", // needs to change
				Flags: []*golang.Flag{
					{
						Name:        "namespace",
						Default:     "default",
						Description: "Kubernetes namespace",
						Required:    false,
					},
				},
				DefaultPreferences: preferences.DefaultKubernetesPreferences,
			},
		},
	}
}

func (p *KubernetesPlugin) SetStream(stream golang.Plugin_RegisterClient) {
	p.stream = stream
}

func (p *KubernetesPlugin) StartProcess(command string, flags map[string]string, kaytuAccessToken string, jobQueue *sdk.JobQueue) error {

	// creating kubernetes object
	kube := kubernetes.NewKubernetes()

	// generating kubeconfig path
	kube.GenerateConfigPath()

	// configuring kubeconfig
	err := kube.ConfigureKubeConfig()
	if err != nil {
		return err
	}

	// configure clientset from kubeconfig
	kube.ConfigureClientSet()

	// get pods using CoreV1 client from clientset
	_, err = pods.GetPods(context.Background(), kube.ClientSet, flags["namespace"])
	if err != nil {
		return err
	}

	// to be extended

	return nil
}

func (p *KubernetesPlugin) ReEvaluate(evaluate *golang.ReEvaluate) {
	//p.processor.ReEvaluate(evaluate.Id, evaluate.Preferences)
}
