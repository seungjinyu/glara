package glarautils

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"

	"github.com/seungjinyu/glara/models"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetglaraPodListInfo gets the information from the podlist and extract the datas and returns GlaraPodInfoList
func GetglaraPodListInfo(clientset *kubernetes.Clientset, namespace string) *models.GlaraPodInfoList {

	podlist, err := K8sPodList(clientset, namespace)

	if err != nil {
		log.Println(err)
	}

	result := extractDataFromPodList(podlist, clientset)

	return result
}

// ExtracDataFromPodList extracts data from the pod list
func extractDataFromPodList(pl *[]v1.Pod, clientset *kubernetes.Clientset) *models.GlaraPodInfoList {

	var tmp models.GlaraPodInfoList
	repl := pl
	tmp.InfoList = make([]models.GlaraPodInfo, len(*repl))

	// fmt.Println(len(repl))
	for i, value := range *repl {

		tmp.InfoList[i] = *extractDataFromPod(&value, clientset)
	}
	return &tmp
}

// ExtracDataFromPod extracts data from the pod
func extractDataFromPod(pd *v1.Pod, clientset *kubernetes.Clientset) *models.GlaraPodInfo {
	// fmt.Println(pd)
	podLog, err := getPodLogs(pd, clientset)
	if err != nil {
		log.Println(err)
	}
	tmp := &models.GlaraPodInfo{
		PodName: pd.GetName(),
		// PodLogs: pd.GetPod(),
		PodLog: podLog,
		// OwnerReference: pd.ObjectMeta.GetOwnerReferences()[0].Kind,
		OwnerReference: pd.OwnerReferences[0].Kind,
	}

	// fmt.Println(tmp.OwnerReference[0])
	// fmt.Println(tmp.PodLog, tmp.PodName)
	return tmp
}

// K8sPodList K8ss backs the pod instance List of the cluster by the kubernetes config file
func K8sPodList(clientset *kubernetes.Clientset, namespace string) (*[]v1.Pod, error) {

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		// log.Println(err.Error())
		return nil, err
	}

	if len(pods.Items) == 0 {
		return nil, errors.New("there is no pod for the requested option")
	}
	// fmt.Printf("There are %d pods \n", len(pods.Items))
	items := &pods.Items
	// result := model.GlaraPodInfo{}
	return items, nil
}

// getPodLogs Here is what we came up with,, eventually using client-go library:
func getPodLogs(pod *v1.Pod, clientset *kubernetes.Clientset) (string, error) {

	podLogOpts := v1.PodLogOptions{}
	nsPodsData := clientset.CoreV1().Pods(pod.Namespace)

	// fmt.Println(nsPodsData)
	req := nsPodsData.GetLogs(pod.Name, &podLogOpts)

	podLogs, err := req.Stream(context.TODO())

	if err != nil {
		return "error in opening stream", err
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)

	_, err = io.Copy(buf, podLogs)

	if err != nil {
		return "error in copy information from podLogs to buf", err
	}

	return buf.String(), nil
}

// // SaveGlaraPodInfoList saves the logs into a multiple *.log file
// func SaveGlaraPodInfoList(pil models.GlaraPodInfoList) {

// 	for _, v := range pil.InfoList {

// 		fileName := "./logs/" + v.PodName + ".log"
// 		tmp, err := os.Create(fileName)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		contents := v.PodName + "\n" + v.PodLog
// 		tmp.WriteString(contents)

// 		defer tmp.Close()
// 	}

// }

// // SaveGlaraPodInfo saves the logs into a *.log file
// func SaveGlaraPodInfo(pi models.GlaraPodInfo) {

// 	fileName := "./logs/" + pi.PodName + "log"
// 	tmp, err := os.Create(fileName)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	contents := pi.PodName + "\n" + pi.PodLog
// 	tmp.WriteString(contents)

// 	defer tmp.Close()

// }

// // GetGlaraPodInfo gets the information from the pod and extract the datas and returns GlaraPodInfo
// func GetGlaraPodInfo(clientset *kubernetes.Clientset, namespace, requestPodName string) models.GlaraPodInfo {

// 	GlaraPodInfo, err := K8sPod(clientset, namespace, requestPodName)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	result := extractDataFromPod(GlaraPodInfo, clientset)

// 	return result
// }

// K8sPod K8s the pod instance of the cluster by the kubernetes config file
// func K8sPod(clientset *kubernetes.Clientset, namespace, requestPodName string) (*v1.Pod, error) {

// 	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		// log.Println(err)
// 		return nil, err
// 	}

// 	for _, v := range pods.Items {
// 		if strings.Contains(v.Name, requestPodName) {
// 			return &v, nil
// 		}
// 	}

// 	err = errors.New("there is no pod for the requested pod")
// 	return nil, err
// }
