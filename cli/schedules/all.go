package schedules

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/schedules"
)

var cmdSchedulesAll = &cobra.Command{
	Use:   "all",
	Short: "Get All schedules: ghue schedules all",
	Long:  `Get all schedules: ghue schedules all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd()
	},
}

func allCmd() {
	connection := config.ReadConfig()
	result, errHUE, err := schedules.GetAllSchedules(connection)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
