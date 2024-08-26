package cmd

import (
	"github.com/spf13/cobra"

	"github.com/gkwa/fogslayer/core"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running hello command")
		core.Hello(logger)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
