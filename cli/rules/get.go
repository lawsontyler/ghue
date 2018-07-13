package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/rules"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

var cmdRulesGet = &cobra.Command{
	Use:     "get",
	Short:   "Get rule attributes and state: ghue rule <id>",
	Long:    `Get rule attributes and state: ghue rule <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue rules state --help")
		} else {
			getCmd(sdk_client.GetSdkClient(config.ReadConfig()), args[0])
		}
	},
}

func getCmd(client *sdk_client.SdkClient, id string) {
	result, errHUE, err := rules.GetRule(client, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
