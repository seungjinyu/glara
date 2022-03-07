/*
Copyright © 2022 NAME HERE seungjinyu93@gmail.com

*/
package cmd

import (
	"errors"
	"log"

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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("arguments are missing check the cmd options\nglara inspect -h")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		KUBE_ENV := args[0]
		namespace := args[1]
		// namespace := args[0]
		pod := args[2]
		rStr := args[3]
		// fmt.Println(rStr)
		var kubecli settings.ClientSetInstance

		settings.ClientSetting(&kubecli, KUBE_ENV)

		err := glarautils.InspectPod(KUBE_ENV, namespace, pod, rStr, kubecli)

		if err != nil {
			log.Println(err)
		}

	},
}

func init() {

	rootCmd.AddCommand(inspectCmd)

}
