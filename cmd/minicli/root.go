package minicli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "mapi",
	Short: "mapi is a is a cli tool for minimal_api",
	Long:  `mapi is a is a cli tool for minimal_api, this helps user to create, build and run the a simple minimal_api project`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
