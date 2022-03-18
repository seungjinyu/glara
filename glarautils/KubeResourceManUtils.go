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
func checkStack(gs *models.GlaraPodInfoStack, rs *models.GlaraPodInfoStack, pod, rStr string) *models.GlaraPodInfoStack {

	if !gs.IsEmpty() {
		tmp := gs.Pop()
		log.Println("trying to push ", tmp.PodName, " ", tmp.OwnerReference)
		if strings.Contains(tmp.PodName, pod) {
			if strings.Contains(tmp.PodLog, rStr) {
				fmt.Println("Pushing a new element")
				rs.Push(tmp)
			} else {
				log.Println(tmp.PodName, " does not contains that log")
			}
		} else {
			log.Println("Pod name is not included")
		}
		checkStack(gs, rs, pod, rStr)
	}

	return rs

}

func inspectOwnerReferenceStack(rs *models.GlaraPodInfoStack, namespace, pod, rStr string, kubecli settings.ClientSetInstance) {
	if !rs.IsEmpty() {
		tmp := rs.Pop()

		inspectResult := strings.Contains(tmp.PodLog, rStr)
		log.Printf("|%7s|%50s|%10s|%5s|%4s|%12s|\n",
			"PODNAME", tmp.PodName,
			"LOG CONTAIN", strconv.FormatBool(inspectResult),
			"TYPE", tmp.OwnerReference)
		switch tmp.OwnerReference {

		case "StatefulSet":
			err := DeleteStatefulSetPod(namespace, tmp.PodName, kubecli.Clientset)
			if err != nil {
				log.Println(err)
			}

		case "ReplicaSet":
			err := DeleteReplicaSetPod(namespace, tmp.PodName, kubecli.Clientset)
			if err != nil {
				log.Println(err)
			}

		case "DaemonSet":
			err := DeleteDaemonSetPod(namespace, tmp.PodName, kubecli.Clientset)
			if err != nil {
				log.Println(err)
			}

		}
		inspectOwnerReferenceStack(rs, namespace, pod, rStr, kubecli)

	} else {
		log.Println("The stack is inspected")
	}
}

// InspectPod inspects the pods and returns an error if there is no pod.
func InspectPod(namespace, pod, rStr string, kubecli settings.ClientSetInstance) error {

	for {
		resultStack := models.NewGlaraPodInfoStack()

		fmt.Println("Inspect called namespace:", namespace, " pod: ", pod, " rStr: ", rStr)
		totalPodStack := GetglaraPodListInfo(
			kubecli.Clientset,
			namespace,
		)

		if totalPodStack != nil {

			resultStack = checkStack(totalPodStack, resultStack, pod, rStr)
			log.Println("The stack is checked")
			if resultStack != nil {
				fmt.Println("Checking")
				TOTALPODSTOCHECK := strconv.Itoa(len(*resultStack))
				inspectOwnerReferenceStack(resultStack, namespace, pod, rStr, kubecli)
				payload := Payload{
					Text:      "Glara deleted " + TOTALPODSTOCHECK + " pods in " + namespace,
					Username:  "Glara-" + os.Getenv("CLUSTER_NAME"),
					IconEmoji: ":high_brightness:",
				}
				url := os.Getenv("SLACK_URL")
				payload.SendSlack(url)
			} else {
				fmt.Println("Stack is empty")
			}

		}

		intervalTime, err := strconv.Atoi(fmt.Sprintf("%s", os.Getenv("INTERVAL_TIME")))
		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second * time.Duration(intervalTime))
	}
}
