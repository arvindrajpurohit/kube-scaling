package kube

import (
	appsv1e "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func scalereplicaset(name string, namespace string, replicas int32) *appsv1e.Scale {
	c := appsv1e.Scale{TypeMeta: metav1.TypeMeta{
		Kind:       "Replicaset",
		APIVersion: "v2alpha1",
	},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1e.ScaleSpec{
			Replicas: replicas,
		},
		Status: appsv1e.ScaleStatus{},
	}
	return &c

}
