package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator", // 短注释
	Long: `A Fast and Flexible Static Site Generator built with	
					love by spf13 and friends in Go.
					Complete documentation is available at http://hugo.spf13.com`, // 详细注释
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPreRun")
	},

	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PreRun")
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("PreRunE")
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run")
	},

	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PostRun")
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPostRun")
	},
}

func init() {
	rootCmd.Flags().StringP("baseURL", "b", "", "hostname (and path) to the root, e.g. https://spf13.com/")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
