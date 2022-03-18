package glarautils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/seungjinyu/glara/errorHandler"
	"github.com/seungjinyu/glara/models"
	"github.com/seungjinyu/glara/settings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeletePodInKube(namespace, PodName string, clientset *kubernetes.Clientset) error {

	err := clientset.CoreV1().Pods(namespace).Delete(
		context.TODO(),
		PodName,
		metav1.DeleteOptions{},
	)

	return err
}

// checkStack is used to check the elements in the stack and push them to an anothor stack
func checkStack(gs *models.GlaraPodInfoStack, rs *models.GlaraPodInfoStack, pod, rStr string) *models.GlaraPodInfoStack {

	if !gs.IsEmpty() {
		tmp := gs.Pop()
		if strings.Contains(tmp.PodName, pod) {
			if strings.Contains(tmp.PodLog, rStr) {
				rs.Push(tmp)
			} else {
				log.Println(tmp.PodName, " does not contains that log")
			}
		} else {
			log.Println(tmp.PodName, "Pod name is not included")
		}
		checkStack(gs, rs, pod, rStr)
	}

	return rs

}

func deletePodProcess(rs *models.GlaraPodInfoStack, namespace, pod, rStr string, kubecli settings.ClientSetInstance) {
	if !rs.IsEmpty() {
		tmp := rs.Pop()

		log.Printf("|%7s|%50s|%4s|%12s|\n",
			"PODNAME", tmp.PodName,
			"TYPE", tmp.OwnerReference,
		)

		err := DeletePodInKube(namespace, tmp.PodName, kubecli.Clientset)
		errorHandler.PrintError(err)
		deletePodProcess(rs, namespace, pod, rStr, kubecli)

	} else {
		log.Println("The stack is inspected")
	}
}

// InspectPod inspects the pods and returns an error if there is no pod.
func InspectPod(namespace, pod, rStr string, kubecli settings.ClientSetInstance) error {

	for {
		resultStack := models.NewGlaraPodInfoStack()

		fmt.Println("Inspect called namespace:", namespace, " pod: ", pod, " rStr: ", rStr)
		totalPodStack, err := GetGlaraPodListInfo(
			kubecli.Clientset,
			namespace,
		)

		errorHandler.PrintError(err)

		if totalPodStack != nil {
			resultStack = checkStack(totalPodStack, resultStack, pod, rStr)
			log.Println("The stack is checked")
			if resultStack != nil {
				log.Println("Checking")
				TOTALPODSTOCHECK := strconv.Itoa(len(*resultStack))
				deletePodProcess(resultStack, namespace, pod, rStr, kubecli)
				if TOTALPODSTOCHECK != "0" {
					SendmsgToSlack(TOTALPODSTOCHECK, namespace)
				}

			} else {
				log.Println("Stack is empty")
			}

		}

		intervalTime, err := strconv.Atoi(fmt.Sprintf("%s", os.Getenv("INTERVAL_TIME")))
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * time.Duration(intervalTime))
	}
}

func SendmsgToSlack(TOTALPODSTOCHECK, namespace string) {
	payload := Payload{
		Parse:       "",
		Username:    "Glara-" + os.Getenv("CLUSTER_NAME"),
		IconUrl:     "",
		IconEmoji:   ":high_brightness:",
		Channel:     "",
		Text:        "Glara deleted " + TOTALPODSTOCHECK + " pods in " + namespace,
		LinkNames:   "",
		Attachments: []Attachment{},
		UnfurlLinks: false,
		UnfurlMedia: false,
		Markdown:    false,
	}
	url := os.Getenv("SLACK_URL")
	payload.SendSlack(url)
}

func DeleteCompletedTask(namespace, pod, rStr string, kubecli settings.ClientSetInstance) error {

	return nil

}
