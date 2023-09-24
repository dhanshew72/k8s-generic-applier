package main

import (
	"context"
	"fmt"

	"github.com/dhanshew72/k8s-generic-applier/applier"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"k8s.io/apimachinery/pkg/runtime"
)

func main() {
	// cm := corev1.ConfigMap{
	// 	TypeMeta: metav1.TypeMeta{
	// 		Kind:       "ConfigMap",
	// 		APIVersion: "v1",
	// 	},
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:      "my-config-map",
	// 		Namespace: "default",
	// 	},

	// 	Data: map[string]string{"foo": "my-data"},
	// }
	cm := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-new-namespace",
		},
	}
	unstruct, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&cm)
	if err != nil {
		panic(err)
	}
	err = applier.CreateOrUpdate(context.Background(), &unstructured.Unstructured{Object: unstruct})
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello, world.")
}
