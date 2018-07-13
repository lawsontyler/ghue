package sensors

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/sensors"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

var cmdSensorsAll = &cobra.Command{
	Use:   "all",
	Short: "Get All sensors: ghue sensors all",
	Long:  `Get all sensors: ghue sensors all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd(sdk_client.GetSdkClient(config.ReadConfig()))
	},
}

func allCmd(client *sdk_client.SdkClient) {
	result, errHUE, err := sensors.GetAllSensors(client)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
