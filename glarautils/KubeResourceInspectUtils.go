package glarautils

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/seungjinyu/glara/models"
	"github.com/seungjinyu/glara/settings"
)

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
				inspectResult := InspectLog(v.PodLog, rStr)
				log.Printf("|%7s|%50s|%10s|%5s|%4s|%12s|\n",
					"PODNAME", v.PodName,
					"LOG CONTAIN", strconv.FormatBool(inspectResult),
					"TYPE", v.OwnerReference)
				switch v.OwnerReference {
				case "StatefulSet":
					log.Println("StatefulSet")
					// case "replicaset":
					// 	log.Println("Replicaset ")
				}

			}
		} else {
			return errors.New("there is no element for the given condition please try to set up an another condition")
		}

		// time.Sleep(time.Second * 5)
	}
}

func InspectLog(PodLog, rStr string) bool {
	result := strings.Contains(PodLog, rStr)
	return result
}
