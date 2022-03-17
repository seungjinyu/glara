package glarautils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/seungjinyu/glara/models"
	"github.com/seungjinyu/glara/settings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeleteStatefulSetPod restarts the stateful set
func DeleteStatefulSetPod(namespace, StatefulSetPodName string, clientset *kubernetes.Clientset) error {
	err := clientset.CoreV1().Pods(namespace).Delete(
		context.TODO(),
		StatefulSetPodName,
		metav1.DeleteOptions{},
	)

	return err
}

func DeleteDaemonSetPod(namespace, DaemonSetPodName string, clientset *kubernetes.Clientset) error {

	err := clientset.CoreV1().Pods(namespace).Delete(
		context.TODO(),
		DaemonSetPodName,
		metav1.DeleteOptions{},
	)
	return err

}

// RestartStatefulSet restarts the replicaset
func DeleteReplicaSetPod(namespace, ReplicaSetPodName string, clientset *kubernetes.Clientset) error {

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
					err := DeleteStatefulSetPod(namespace, v.PodName, kubecli.Clientset)
					if err != nil {
						log.Println(err)
					}
				case "ReplicaSet":
					err := DeleteReplicaSetPod(namespace, v.PodName, kubecli.Clientset)
					if err != nil {
						log.Println(err)
					}
				case "DaemonSet":
					err := DeleteDaemonSetPod(namespace, v.PodName, kubecli.Clientset)
					if err != nil {
						log.Println(err)
					}
				}

			}
			payload := Payload{
				Text:      "Glara deleted " + strconv.Itoa(len(tmpPodList)) + " pods in " + namespace,
				Username:  "Glara-" + os.Getenv("CLUSTER_NAME"),
				IconEmoji: ":high_brightness:",
			}
			url := os.Getenv("SLACK_URL")
			payload.SendSlack(url)
		} else {
			log.Println("There is no pod that matches the condition")
		}

		intervalTime, err := strconv.Atoi(fmt.Sprintf("%s", os.Getenv("INTERVAL_TIME")))
		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second * time.Duration(intervalTime))
	}
}

// func InspectLog(PodLog, rStr string) bool {
// 	result := strings.Contains(PodLog, rStr)
// 	return result
// }
