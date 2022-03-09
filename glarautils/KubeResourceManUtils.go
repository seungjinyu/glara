package glarautils

import (
	"context"
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

// RestartStatefulSet restarts the stateful set
func RestartStatefulSet(namespace, StatefulSetPodName string, clientset *kubernetes.Clientset) error {
	err := clientset.CoreV1().Pods(namespace).Delete(
		context.TODO(),
		StatefulSetPodName,
		metav1.DeleteOptions{},
	)

	return err
}

func RestartDaemonSet(namespace, DaemonSetPodName string, clientset *kubernetes.Clientset) error {

	err := clientset.CoreV1().Pods(namespace).Delete(
		context.TODO(),
		DaemonSetPodName,
		metav1.DeleteOptions{},
	)
	return err

}

// RestartStatefulSet restarts the replicaset
func RestartReplicaSet(namespace, ReplicaSetPodName string, clientset *kubernetes.Clientset) error {

	err := clientset.CoreV1().Pods(namespace).Delete(
		context.TODO(),
		ReplicaSetPodName,
		metav1.DeleteOptions{},
	)

	return err
}

// InspectPod inspects the pods and returns an error if there is no pod.
func InspectPod(namespace, pod, rStr string, kubecli settings.ClientSetInstance) error {

	for {
		var tmpPodList []models.GlaraPodInfo
		fmt.Println("Inspect called namespace:", namespace, " pod: ", pod, " rStr: ", rStr)
		datas := GetglaraPodListInfo(
			kubecli.Clientset,
			namespace,
		)

		for _, v := range datas.InfoList {
			if strings.Contains(v.PodName, pod) {
				if strings.Contains(v.PodLog, rStr) {
					tmpPodList = append(tmpPodList, v)
				}
			}
		}

		for _, v := range tmpPodList {
			fmt.Println(v.PodName)
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
					err := RestartStatefulSet(namespace, v.PodName, kubecli.Clientset)
					if err != nil {
						log.Println(err)
					}
				case "ReplicaSet":
					err := RestartReplicaSet(namespace, v.PodName, kubecli.Clientset)
					if err != nil {
						log.Println(err)
					}
				case "DaemonSet":
					err := RestartDaemonSet(namespace, v.PodName, kubecli.Clientset)
					if err != nil {
						log.Println(err)
					}
				}

			}
		} else {
			log.Println("There is no pod that matches the condition")
		}

		time.Sleep(time.Second * 5)
	}
}

// func InspectLog(PodLog, rStr string) bool {
// 	result := strings.Contains(PodLog, rStr)
// 	return result
// }
