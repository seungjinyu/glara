package glarautils

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/seungjinyu/glara/errorHandler"
	"github.com/seungjinyu/glara/models"
	"github.com/seungjinyu/glara/settings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// func UpdateDeploymentsInKube(namespace string, deployName *v1.Deployment, clientset *kubernetes.Clientset) (*v1.Deployment, error) {
// 	deploy, err := clientset.AppsV1().Deployments(namespace).Update(
// 		context.TODO(),
// 		deployName,
// 		metav1.UpdateOptions{},
// 	)
// 	return deploy, err
// }

// DeletePodInKube
func DeletePodInKube(namespace, PodName string, clientset *kubernetes.Clientset) error {

	err := clientset.CoreV1().Pods(namespace).Delete(
		context.TODO(),
		PodName,
		metav1.DeleteOptions{},
	)

	return err
}

// deletePodInProcess
func deletePodProcess(rs *models.GlaraPodInfoStack, namespace, pod, rStr string, kubecli settings.ClientSetInstance) {
	if !rs.IsEmpty() {
		tmp := rs.Pop()

		log.Printf("|%7s|%50s|%4s|%12s|\n",
			"PODNAME", tmp.PodName,
			"TYPE", tmp.OwnerReference,
		)

		err := DeletePodInKube(namespace, tmp.PodName, kubecli.Clientset)
		errorHandler.PrintError(err)
		SendmsgToSlack(tmp.PodName, namespace)
		deletePodProcess(rs, namespace, pod, rStr, kubecli)

	} else {
		log.Println("The stack is inspected")
	}
}

// checkLogStack is used to check the elements in the stack and push them to an anothor stack
func checkLogStack(gs *models.GlaraPodInfoStack, rs *models.GlaraPodInfoStack, pod, rStr string) *models.GlaraPodInfoStack {

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
		checkLogStack(gs, rs, pod, rStr)
	}

	return rs

}

func SendmsgToSlack(PodName string, namespace string) {
	payload := Payload{
		Parse:       "",
		Username:    "Glara-" + os.Getenv("CLUSTER_NAME"),
		IconUrl:     "",
		IconEmoji:   ":high_brightness:",
		Channel:     "",
		Text:        "Glara deleted " + PodName + " pod in " + namespace,
		LinkNames:   "",
		Attachments: []Attachment{},
		UnfurlLinks: false,
		UnfurlMedia: false,
		Markdown:    false,
	}

	url := os.Getenv("SLACK_URL")

	if url != "" {
		payload.SendSlack(url)
	} else {
		log.Println("Error with SLACK URL")
	}

}
