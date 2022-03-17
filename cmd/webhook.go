/*
Copyright © 2022 NAME HERE seungjinyu93@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/seungjinyu/glara/glarautils"
	"github.com/spf13/cobra"
)

// webhookCmd represents the webhook command
var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webhook called")
		payload := glarautils.Payload{
			Text:      "Webhook Check",
			Username:  "Glara",
			IconEmoji: ":high_brightness:",
		}
		url := os.Getenv("SLACK_URL")
		payload.SendSlack(url)
	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)
}
