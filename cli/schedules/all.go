package schedules

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/schedules"
	"github.com/lawsontyler/ghue/sdk/common"
)

var cmdSchedulesAll = &cobra.Command{
	Use:   "all",
	Short: "Get All schedules: ghue schedules all",
	Long:  `Get all schedules: ghue schedules all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd(config.ReadConfig())
	},
}

func allCmd(connection *common.Connection) {
	result, errHUE, err := schedules.GetAllSchedules(connection)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
