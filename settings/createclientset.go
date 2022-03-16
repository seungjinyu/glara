package settings

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/seungjinyu/glara/errorHandler"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	certutil "k8s.io/client-go/util/cert"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
)

func ClientSetting(cli *ClientSetInstance, kubeEnv string) {
	KUBE_ENV := kubeEnv
	fmt.Println("KUBE ENV is : [", KUBE_ENV, "]")
	if KUBE_ENV == "OUT" {

		err := cli.CreateOutClientSet()
		errorHandler.ConfigError(err)

	} else {

		err := cli.CreateInClientSet()
		errorHandler.ConfigError(err)
	}
}

type ClientSetInstance struct {
	Clientset *kubernetes.Clientset
}

func (c *ClientSetInstance) CreateInClientSet() error {
	// config, err := rest.InClusterConfig()

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	const (
		tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
		rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	)

	token := os.Getenv("serviceaccount-token")
	host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")

	tlsClientConfig := rest.TLSClientConfig{}

	if _, err := certutil.NewPool(rootCAFile); err != nil {
		klog.Errorf("Expected to load root CA config from %s, but got err: %v", rootCAFile, err)
	} else {
		tlsClientConfig.CAFile = rootCAFile
	}

	config := &rest.Config{
		Host:            "https://" + net.JoinHostPort(host, port),
		TLSClientConfig: tlsClientConfig,
		BearerToken:     string(token),
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err.Error())
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
		log.Println(err.Error())
	}

	// fmt.Println(config)
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Println(err.Error())
	}
	c.Clientset = clientset
	return err
}
