package commands

import "github.com/spf13/cobra"

// Documentation for init command
// This command will be used to initialize a new minimal_api project
// Here is the basic structure of a project
// Recommended folder structure for project:
/*
project/
|--main.go
|--routes/
|  |--routes.go
|--handlers/
|  |--handlers.go
|--models/
|  |--models.go
|--config/
|  |--config.go
// Other things that dev wants to add
*/

// initCmd is a command to initialize a new minimal_api project
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init is a command to initialize a new minimal_api project",
	Long:  `init is a command to initialize a new minimal_api project, this will create a new minimal_api project with default configuration`,
}

// Create project skeleton
// This func will create a boilerplate for minimal_api project
func createProjectSkeleton() {
	// Dirs
}
