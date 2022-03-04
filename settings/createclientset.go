package settings

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/seungjinyu/glara/errorHandler"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func ClientSetting(cli *ClientSetInstance, kubeEnv string) {
	KUBE_ENV := kubeEnv
	if KUBE_ENV == "OUT" {
		fmt.Println(KUBE_ENV)
		err := cli.CreateOutClientSet()
		errorHandler.ConfigError(err)

	} else {
		fmt.Println(KUBE_ENV)
		err := cli.CreateInClientSet()
		errorHandler.ConfigError(err)
	}
}

type ClientSetInstance struct {
	Clientset *kubernetes.Clientset
	Appenv    string
}

func (c *ClientSetInstance) CreateInClientSet() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	c.Clientset = clientset

	return err
}

func (c *ClientSetInstance) CreateOutClientSet() error {

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional)  absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// fmt.Println("kubeconfig path:", *kubeconfig)

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(config)
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}
	c.Clientset = clientset
	return err
}
