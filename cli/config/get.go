package config

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/config"
	"github.com/lawsontyler/ghue/sdk/factory"
)

var cmdConfigGet = &cobra.Command{
	Use:     "get",
	Short:   "Get configuration: ghue configuration get",
	Long:    `Get configuration: ghue configuration get`,
	Aliases: []string{"show"},
	Run: func(cmd *cobra.Command, args []string) {
		getCmd(factory.GetSdkClient(ReadConfig()))
	},
}

func getCmd(client *factory.SdkClient) {
	result, errHUE, err := config.Get(client)

	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
