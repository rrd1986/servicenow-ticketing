package cmd

import (
	"github.com/rrd1986/servicenow-ticketing/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var InputFile string

// filesCmd represents the files command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create servicenow ticket.",
	Long:  `Create servicenow ticket.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Debug {
			utils.LogFlags()
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVarP(&InputFile, "inputfile", "i", "input.txt", "input file to be processed")
	viper.BindPFlag("inputfile", createCmd.PersistentFlags().Lookup("inputfile"))

}
