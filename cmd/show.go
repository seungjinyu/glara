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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show command shows whether the namespace has an pod",
	Long: `show command shows whether the namespace has an pod required arguments
	Required arguments
	  namespace
	  - The namespace where the pod is included
	
	[FORMAT]

	  glara inspect [ENV][NAMESPACE]

	[USAGE EXAMPLE]

	  glara inspect OUT mongodb

	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("there are few arguments, please try to enter in the required arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		KUBE_ENV := args[0]
		namespace := args[1]
		var kubecli settings.ClientSetInstance
		settings.ClientSetting(&kubecli, KUBE_ENV)

		result := glarautils.GetglaraPodListInfo(
			kubecli.Clientset,
			namespace,
		)
		for _, v := range result.InfoList {
			log.Printf("|%7s|%50s|\n",
				"PODNAME",
				v.PodName,
			)
			// log.Println("[PODLOG]: ", v.PodLog)
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// showCmd.PersistentFlags().String("ENV", "IN", "KUBE ENV string")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
