package cmd

import (
	"github.com/mobingi/gosdk/pkg/util/cmdline"
	"github.com/spf13/cobra"
)

func HelpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "help about any command",
		Long: `Help provides help for any command in the application.
Simply type '` + cmdline.Args0() + ` help [path to command]' for full details.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}
