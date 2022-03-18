/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/seungjinyu/glara/glarautils"
	"github.com/seungjinyu/glara/settings"
	"github.com/spf13/cobra"
)

// deletecomtaskCmd represents the deletecomtask command
var deletecomtaskCmd = &cobra.Command{
	Use:   "deletecomtask",
	Short: "Delete the completed tasks",
	Long: `deletecomtask the logs of the pods and if there is a problem with the agent restarts the deployment
	Required arguments
	- env 
	  whether the command is executed inside or outside the cluster
	  inside = IN , outside = OUT 
	- namespace
	  The namespace where the pod is in
	- pod
	  The name of the pod you would like to deletecomtask, not meaning the full name but the basic name 
	- DELETESTATUS
	  The log which will occur the deployments to be delete manually

	[FORMAT]

	  glara deletecomtask [ENV] [NAMESPACE] [PODNAME] [DELETESTATUS]

	[USAGE EXAMPLE]

	  glara deletecomtask OUT default agent Successed
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deletecomtask called")
		KUBE_ENV := os.Getenv("KUBE_ENV")
		namespace := os.Getenv("NAMESPACE")
		// namespace := args[0]
		pod := os.Getenv("POD")
		rStr := os.Getenv("DELETESTATUS")

		var kubecli settings.ClientSetInstance
		fmt.Println("Client setting is now initialized")
		settings.ClientSetting(&kubecli, KUBE_ENV)
		fmt.Println("Client setting completed")
		glarautils.DeleteCompletedTask(namespace, pod, rStr, kubecli)

	},
}

func init() {
	rootCmd.AddCommand(deletecomtaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deletecomtaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deletecomtaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
