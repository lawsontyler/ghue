package info

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/info"
)

var cmdInfoTimezones = &cobra.Command{
	Use:   "timezones",
	Short: "Get All Timezones: ghue info timezones",
	Long:  `Get All Timezones: ghue info timezones`,
	Run: func(cmd *cobra.Command, args []string) {
		allTimezonesCmd()
	},
}

func allTimezonesCmd() {
	connection := config.ReadConfig()
	result, errHUE, err := info.GetAllTimezones(connection)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
