package lights

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/lights"
)

var cmdLightsAll = &cobra.Command{
	Use:   "all",
	Short: "Get All lights: ghue lights all",
	Long:  `Get all lights: ghue lights all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd()
	},
}

func allCmd() {
	connection := config.ReadConfig()
	result, errHUE, err := lights.GetAllLights(connection)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
