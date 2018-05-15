package scenes

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/scenes"
	"github.com/lawsontyler/ghue/sdk/factory"
)

var cmdScenesAll = &cobra.Command{
	Use:   "all",
	Short: "Get All scenes: ghue scenes all",
	Long:  `Get all scenes: ghue scenes all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd(factory.GetSdkClient(config.ReadConfig()))
	},
}

func allCmd(client *factory.SdkClient) {
	result, errHUE, err := scenes.GetAllScenes(client)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
