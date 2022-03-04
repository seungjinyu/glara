package glarautils

import (
	"bytes"
	"context"
	"io"
	"log"
	"strings"

	"github.com/seungjinyu/glara/models"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// var defaultNamespace = "default"
// var defaultPodname = "coredns"

// ExtracDataFromPodList extracts data from the pod list
func extractDataFromPodList(pl []v1.Pod, clientset *kubernetes.Clientset) models.GlaraPodInfoList {

	var tmp models.GlaraPodInfoList
	repl := pl
	tmp.InfoList = make([]models.GlaraPodInfo, len(repl))

	// fmt.Println(len(repl))
	for i, value := range repl {

		tmp.InfoList[i] = extractDataFromPod(value, clientset)
	}
	return tmp
}

// ExtracDataFromPod extracts data from the pod
func extractDataFromPod(pi v1.Pod, clientset *kubernetes.Clientset) models.GlaraPodInfo {

	tmp := models.GlaraPodInfo{
		PodName:        pi.GetName(),
		PodLog:         getPodLogs(pi, clientset),
		OwnerReference: pi.ObjectMeta.GetOwnerReferences()[0].Kind,
	}
	// fmt.Println(tmp.PodLog, tmp.PodName)
	return tmp
}

// GetglaraPodListInfo gets the information from the podlist and extract the datas and returns GlaraPodInfoList
func GetglaraPodListInfo(clientset *kubernetes.Clientset, namespace string) models.GlaraPodInfoList {

	podlist := K8sPodList(clientset, namespace)
	if podlist == nil {
		log.Panic("The Request returned ZERO pod")
	}
	result := extractDataFromPodList(podlist, clientset)

	return result
}

// GetGlaraPodInfo gets the information from the pod and extract the datas and returns GlaraPodInfo
func GetGlaraPodInfo(clientset *kubernetes.Clientset, namespace, requestPodName string) models.GlaraPodInfo {

	GlaraPodInfo, err := K8sPod(clientset, namespace, requestPodName)
	if err != nil {
		panic(err)
	}
	result := extractDataFromPod(GlaraPodInfo, clientset)

	return result
}

// K8sPodList K8ss backs the pod instance List of the cluster by the kubernetes config file
func K8sPodList(clientset *kubernetes.Clientset, namespace string) []v1.Pod {

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	if len(pods.Items) == 0 {
		log.Printf("There are no pod from the request please try an another option and run the agent again")
		return nil
	}
	// fmt.Printf("There are %d pods \n", len(pods.Items))
	items := pods.Items
	// result := model.GlaraPodInfo{}
	return items

}

// K8sPod K8ss backs the pod instance of the cluster by the kubernetes config file
func K8sPod(clientset *kubernetes.Clientset, namespace, requestPodName string) (v1.Pod, error) {

	// if namespace == "" {
	// 	namespace = defaultNamespace
	// }
	// if requestPodName == "" {
	// 	requestPodName = defaultPodname
	// }

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, v := range pods.Items {
		if strings.Contains(v.Name, requestPodName) {
			return v, nil
		}
	}

	return v1.Pod{}, err
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

// getPodLogs Here is what we came up with,, eventually using client-go library:
func getPodLogs(pod v1.Pod, clientset *kubernetes.Clientset) string {

	podLogOpts := v1.PodLogOptions{}
	nsPodsData := clientset.CoreV1().Pods(pod.Namespace)

	// fmt.Println(nsPodsData)
	req := nsPodsData.GetLogs(pod.Name, &podLogOpts)

	podLogs, err := req.Stream(context.TODO())

	if err != nil {
		return "error in opening stream"
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)

	_, err = io.Copy(buf, podLogs)

	if err != nil {
		return "error in copy information from podLogs to buf"
	}

	return buf.String()
}
