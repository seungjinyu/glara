/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		KUBE_ENV := args[0]
		namespace := args[1]
		statefulsetName := args[1]
		var kubecli settings.ClientSetInstance
		settings.ClientSetting(&kubecli, KUBE_ENV)

		result := glarautils.RestartStatefulSet(namespace, statefulsetName, kubecli.Clientset)
		log.Println("result is ", result)
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
