package scenes

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/scenes"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

var cmdScenesGet = &cobra.Command{
	Use:     "get",
	Short:   "Get scene attributes and state: ghue scene <id>",
	Long:    `Get scene attributes and state: ghue scene <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue scenes state --help")
		} else {
			getCmd(sdk_client.GetSdkClient(config.ReadConfig()), args[0])
		}
	},
}

func getCmd(client *sdk_client.SdkClient, id string) {
	result, errHUE, err := scenes.GetScene(client, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
