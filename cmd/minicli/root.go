package minicli

import (
	"github.com/pol-cova/minimal_api/cmd/minicli/commands"
	"github.com/spf13/cobra"
)

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

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(commands.InitCmd)

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version is a command to get the version of mapi",
	Long:  `version is a command to get the version of mapi`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("mapi version 0.1.0\n")
	},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "info is a command to get the information about mapi",
	Long:  `info is a command to get the information about mapi`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("mapi is a cli tool for minimal_api")
		cmd.Println("This helps user to create, build and run the a simple minimal_api project")
		cmd.Println("For more information visit: https://github.com/pol-cova/minimal_api")
		cmd.Println("Created by: pol-cova")
	},
}
