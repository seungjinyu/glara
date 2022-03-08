/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"log"

	"github.com/seungjinyu/glara/glarautils"
	"github.com/seungjinyu/glara/settings"
	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart an kubernetes resource",
	Long: `Restart Method
	Required arguments
	- env 
	  whether the command is executed inside or outside the cluster
	  inside = IN , outside = OUT 
	- namespace
	  The namespace where the pod is in
	- pod name 
	  The name of the pod you would like to inspect, not meaning the full name but the basic name 
	- pod owner reference
	  The type of the pod owner reference

	[FORMAT]

	  glara inspect [ENV] [NAMESPACE] [PODNAME] 

	[USAGE EXAMPLE]

	  glara inspect OUT mongodb mongo "SSL peer certificate validation failed"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 4 {
			return errors.New("requried more arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		KUBE_ENV := args[0]
		namespace := args[1]
		resourceName := args[2]
		resourceType := args[3]
		var kubecli settings.ClientSetInstance
		// var result []string
		settings.ClientSetting(&kubecli, KUBE_ENV)
		var err error
		switch resourceType {
		case "rs":
			err = glarautils.RestartReplicaSet(namespace, resourceName, kubecli.Clientset)
		case "stf":
			err = glarautils.RestartStatefulSet(namespace, resourceName, kubecli.Clientset)
		}

		if err != nil {
			log.Println(err)
		} else {
			log.Println("[Pods has been deleted]")
		}
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
