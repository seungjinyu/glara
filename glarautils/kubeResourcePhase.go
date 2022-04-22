package glarautils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/seungjinyu/glara/errorHandler"
	"github.com/seungjinyu/glara/models"
	"github.com/seungjinyu/glara/settings"
)

// InspectPodLogPhase
func InspectPodLogPhase(namespace, pod, rStr string, kubecli settings.ClientSetInstance) error {

	for {
		resultStack := models.NewGlaraPodInfoStack()

		fmt.Println("Inspect called namespace:", namespace, " pod: ", pod, " rStr: ", rStr)
		totalPodStack, err := GetGlaraPodListInfo(
			kubecli.Clientset,
			namespace,
		)

		errorHandler.PrintError(err)

		if totalPodStack != nil {
			resultStack = checkLogStack(totalPodStack, resultStack, pod, rStr)
			log.Println("The stack is checked")
			if resultStack != nil {
				log.Println("Checking")
				// TOTALPODSTOCHECK := len(*resultStack)
				deletePodProcess(resultStack, namespace, pod, rStr, kubecli)
				// if TOTALPODSTOCHECK > 0 {
				// 	log.Println("Sending message to slack")
				// 	SendmsgToSlack(TOTALPODSTOCHECK, namespace)
				// }

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

// DeleteTaskCompletedPhase
func DeleteTaskCompletedPhase(namespace, pod, taskstatus string, kubecli settings.ClientSetInstance) error {
	for {
		resultStack := models.NewGlaraPodInfoStack()

		fmt.Println("Inspect called namespace: ", namespace, " task status ", taskstatus)
		// resultStack = checkTaskStack()
		if resultStack != nil {

		}
	}
	// return nil

}
