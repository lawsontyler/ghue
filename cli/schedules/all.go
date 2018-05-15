package schedules

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/schedules"
	"github.com/lawsontyler/ghue/sdk/factory"
)

var cmdSchedulesAll = &cobra.Command{
	Use:   "all",
	Short: "Get All schedules: ghue schedules all",
	Long:  `Get all schedules: ghue schedules all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd(factory.GetSdkClient(config.ReadConfig()))
	},
}

func allCmd(client *factory.SdkClient) {
	result, errHUE, err := schedules.GetAllSchedules(client)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
