package pods

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPods(t *testing.T) {

	testpod := "test-pod"

	clientSet := fake.NewSimpleClientset(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testpod,
			Namespace: "default",
		},
	})

	pods, err := GetPods(context.Background(), clientSet, "default")
	if err != nil {
		t.Errorf("[TestGetPods]: %s", err.Error())
	}

	if len(pods.Items) != 1 {
		t.Error()
	}
	if pods.Items[0].Name != testpod {
		t.Error()
	}

}
