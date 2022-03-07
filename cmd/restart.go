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
	Long:  `Testing for restart method`,
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
