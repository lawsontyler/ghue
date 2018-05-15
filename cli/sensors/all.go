package sensors

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/sensors"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/factory"
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
	client := factory.GetSdkClient(connection)
	result, errHUE, err := sensors.GetAllSensors(client)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
