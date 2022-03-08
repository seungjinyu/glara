package glarautils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/seungjinyu/glara/models"
	"github.com/seungjinyu/glara/settings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// RestareResources restarts the resoucres by there ownerreferences type
// func RestartResources(namespace, resourceName, resourceType string, clientset *kubernetes.Clientset) error {

// 	// if resourceType == "Statefulset"{

// 	// }

// 	result, _ := clientset.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})

// 	var rList []string

// 	for _, v := range result.Items {
// 		log.Println(v.GetName())
// 		if strings.Contains(v.GetName(), resourceName) {
// 			rList = append(rList, v.GetName())
// 		}
// 	}

// 	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// var podList []string

// 	for _, v := range pods.Items {
// 		for _, j := range rList {
// 			if strings.Contains(v.Name, j) {
// 				// podList = append(podList, v.Name)
// 				err = clientset.CoreV1().Pods(namespace).Delete(
// 					context.TODO(),
// 					j,
// 					metav1.DeleteOptions{},
// 				)
// 				return err
// 			}
// 		}
// 	}

// 	return nil

// }

// RestartStatefulSet restarts the stateful set
func RestartStatefulSet(namespace, statefulsetName string, clientset *kubernetes.Clientset) error {

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

	// var podList []string

	for _, v := range pods.Items {
		for _, j := range sfList {
			if strings.Contains(v.Name, j) {
				// podList = append(podList, v.Name)
				err = clientset.CoreV1().Pods(namespace).Delete(
					context.TODO(),
					j,
					metav1.DeleteOptions{},
				)
				return err
			}
		}
	}
	return nil
	// return podList
}

// RestartStatefulSet restarts the replicaset
func RestartReplicaSet(namespace, ReplicaSets string, clientset *kubernetes.Clientset) error {

	result, _ := clientset.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{})

	var rsList []string

	for _, v := range result.Items {
		if strings.Contains(v.GetName(), ReplicaSets) {
			rsList = append(rsList, v.GetName())
		}
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// var podList []string

	for _, v := range pods.Items {
		for _, j := range rsList {
			if strings.Contains(v.Name, j) {
				// podList = append(podList, v.Name)
				err = clientset.CoreV1().Pods(namespace).Delete(
					context.TODO(),
					v.Name,
					metav1.DeleteOptions{},
				)
				return err
			}
		}
	}

	return nil
	// return podList
}

func InspectPod(KUBE_ENV, namespace, pod, rStr string, kubecli settings.ClientSetInstance) error {

	for {
		var tmpPodList []models.GlaraPodInfo
		fmt.Println("Inspect called namespace:", namespace, " pod: ", pod)
		datas := GetglaraPodListInfo(
			kubecli.Clientset,
			namespace,
		)
		for _, v := range datas.InfoList {
			if strings.Contains(v.PodName, pod) && strings.Contains(v.PodLog, rStr) {
				tmpPodList = append(tmpPodList, v)
			}
		}

		if len(tmpPodList) != 0 {
			for _, v := range tmpPodList {
				inspectResult := strings.Contains(v.PodLog, rStr)
				log.Printf("|%7s|%50s|%10s|%5s|%4s|%12s|\n",
					"PODNAME", v.PodName,
					"LOG CONTAIN", strconv.FormatBool(inspectResult),
					"TYPE", v.OwnerReference)
				switch v.OwnerReference {
				case "StatefulSet":
					log.Println("StatefulSet")
					err := RestartStatefulSet(namespace, pod, kubecli.Clientset)
					if err != nil {
						log.Println(err)
					}

				case "ReplicaSet":
					log.Println("ReplicaSet")
					err := RestartReplicaSet(namespace, pod, kubecli.Clientset)
					if err != nil {
						log.Println(err)
					}
				}

			}
		} else {
			return errors.New("there is no element for the given condition please try to set up an another condition")
		}

		time.Sleep(time.Second * 5)
	}
}

// func InspectLog(PodLog, rStr string) bool {
// 	result := strings.Contains(PodLog, rStr)
// 	return result
// }
