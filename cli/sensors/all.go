package sensors

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/sensors"
	"github.com/lawsontyler/ghue/sdk/common"
)

var cmdSensorsAll = &cobra.Command{
	Use:   "all",
	Short: "Get All sensors: ghue sensors all",
	Long:  `Get all sensors: ghue sensors all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd(config.ReadConfig())
	},
}

func allCmd(connection *common.Connection) {
	result, errHUE, err := sensors.GetAllSensors(connection)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
