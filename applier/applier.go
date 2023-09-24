package applier

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/dynamic"
)

func CreateOrUpdate(ctx context.Context, object *unstructured.Unstructured) error {
	dynClient, err := getClient()
	if err != nil {
		return err
	}
	_, err = dynClient.Resource(getGVR(object)).Create(ctx, object, metav1.CreateOptions{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func getGVR(object *unstructured.Unstructured) schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    object.GetObjectKind().GroupVersionKind().Group,
		Version:  object.GetAPIVersion(),
		Resource: object.GetKind(),
	}
}

func getClient() (*dynamic.DynamicClient, error) {
	config, err := getConfig()
	if err != nil {
		return nil, err
	}
	builtConfig, configErr := clientcmd.BuildConfigFromFlags("", config)
	if configErr != nil {
		return nil, err
	}
	dynClient := dynamic.NewForConfigOrDie(builtConfig)
	return dynClient, nil
}

func getConfig() (string, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("discover-k8s: error retrieving home directory: %s", err)
	}
	return filepath.Join(dir, ".kube", "config"), nil
}
