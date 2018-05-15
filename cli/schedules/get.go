package schedules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/schedules"
	"github.com/lawsontyler/ghue/sdk/factory"
)

var cmdSchedulesGet = &cobra.Command{
	Use:     "get",
	Short:   "Get schedule attributes and state: ghue schedule <id>",
	Long:    `Get schedule attributes and state: ghue schedule <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue schedules state --help")
		} else {
			getCmd(factory.GetSdkClient(config.ReadConfig()), args[0])
		}
	},
}

func getCmd(client *factory.SdkClient, id string) {
	result, errHUE, err := schedules.GetSchedule(client, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
