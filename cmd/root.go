/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package cmd

import (
	"os"

	"github.com/edoardottt/depsdev/pkg/depsdev"
	"github.com/edoardottt/depsdev/pkg/output"
	"github.com/spf13/cobra"
)

var api *depsdev.API

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "depsdev",
	Short: output.ShortDescription,
	Long:  output.Banner,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		api = depsdev.NewAPI()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(depsCmd)
	rootCmd.AddCommand(advisoryCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(graphCmd)
	rootCmd.AddCommand(reqsCmd)
}
