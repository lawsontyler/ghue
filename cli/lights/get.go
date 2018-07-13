package lights

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/lights"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

var cmdLightsGet = &cobra.Command{
	Use:     "get",
	Short:   "Get light attributes and state: ghue light <id>",
	Long:    `Get light attributes and state: ghue light <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue lights state --help")
		} else {
			getCmd(sdk_client.GetSdkClient(config.ReadConfig()), args[0])
		}
	},
}

func getCmd(client *sdk_client.SdkClient, id string) {
	result, errHUE, err := lights.GetLight(client, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
