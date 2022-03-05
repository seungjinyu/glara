package glarautils

import (
	"context"
	"log"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// RestartStatefulSet restarts the stateful set
func RestartStatefulSet(namespace, statefulsetName string, clientset *kubernetes.Clientset) []string {

	result, _ := clientset.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})

	var sfList []string

	for _, v := range result.Items {
		log.Println(v.GetName())
		if strings.Contains(v.GetName(), statefulsetName) {
			sfList = append(sfList, v.GetName())
		}
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var podList []string

	for _, v := range pods.Items {
		for _, j := range sfList {
			if strings.Contains(v.Name, j) {
				podList = append(podList, v.Name)
			}
		}
	}
	return podList
}

// RestartStatefulSet restarts the replicaset
func RestartReplicaSet(namespace, ReplicaSets string, clientset *kubernetes.Clientset) []string {

	result, _ := clientset.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{})

	var sfList []string

	for _, v := range result.Items {
		log.Println(v.GetName())
		if strings.Contains(v.GetName(), ReplicaSets) {
			sfList = append(sfList, v.GetName())
		}

	}
	return sfList
}
