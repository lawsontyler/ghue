package groups

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/groups"
	"github.com/lawsontyler/ghue/sdk/factory"
)

var cmdGroupsGet = &cobra.Command{
	Use:     "get",
	Short:   "Get group attributes and state: ghue group <id>",
	Long:    `Get group attributes and state: ghue group <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue groups state --help")
		} else {
			getCmd(factory.GetSdkClient(config.ReadConfig()), args[0])
		}
	},
}

func getCmd(client *factory.SdkClient, id string) {
	result, errHUE, err := groups.GetGroup(client, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
