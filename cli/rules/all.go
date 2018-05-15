package rules

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/rules"
	"github.com/lawsontyler/ghue/sdk/factory"
)

var cmdRulesAll = &cobra.Command{
	Use:   "all",
	Short: "Get All rules: ghue rules all",
	Long:  `Get all rules: ghue rules all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd(factory.GetSdkClient(config.ReadConfig()))
	},
}

func allCmd(client *factory.SdkClient) {
	result, errHUE, err := rules.GetAllRules(client)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
