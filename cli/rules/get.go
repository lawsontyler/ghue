package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/rules"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/factory"
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
			getCmd(config.ReadConfig(), args[0])
		}
	},
}

func getCmd(connection *common.Connection, id string) {
	client := factory.GetSdkClient(connection)
	result, errHUE, err := rules.GetRule(client, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
