/*
Copyright © 2022 NAME HERE seungjinyu93@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/seungjinyu/glara/glarautils"
	"github.com/seungjinyu/glara/settings"
	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect the logs of the pods and if there is a problem with the agent restarts the deployment",
	Long: `Inspect the logs of the pods and if there is a problem with the agent restarts the deployment
	Required arguments
	- env 
	  whether the command is executed inside or outside the cluster
	  inside = IN , outside = OUT 
	- namespace
	  The namespace where the pod is in
	- pod
	  The name of the pod you would like to inspect, not meaning the full name but the basic name 
	- logstring
	  The log which will occur the deployments to be restarted manually

	[FORMAT]

	  glara inspect [ENV] [NAMESPACE] [PODNAME] [logstring]

	[USAGE EXAMPLE]

	  glara inspect OUT mongodb mongo "SSL peer certificate validation failed"

	`,
	Run: func(cmd *cobra.Command, args []string) {

		KUBE_ENV := os.Getenv("KUBE_ENV")
		namespace := os.Getenv("NAMESPACE")
		// namespace := args[0]
		pod := os.Getenv("POD")
		rStr := os.Getenv("RESTARTLOG")

		var kubecli settings.ClientSetInstance
		fmt.Println("Client setting is now initialized")
		settings.ClientSetting(&kubecli, KUBE_ENV)
		fmt.Println("Client setting completed")
		glarautils.InspectPodLogPhase(namespace, pod, rStr, kubecli)

	},
}

func init() {

	rootCmd.AddCommand(inspectCmd)

}
