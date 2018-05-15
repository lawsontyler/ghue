package info

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/info"
	"github.com/lawsontyler/ghue/sdk/factory"
)

var cmdInfoTimezones = &cobra.Command{
	Use:   "timezones",
	Short: "Get All Timezones: ghue info timezones",
	Long:  `Get All Timezones: ghue info timezones`,
	Run: func(cmd *cobra.Command, args []string) {
		allTimezonesCmd(factory.GetSdkClient(config.ReadConfig()))
	},
}

func allTimezonesCmd(client *factory.SdkClient) {
	result, errHUE, err := info.GetAllTimezones(client)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
