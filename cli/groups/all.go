package groups

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/groups"
	"github.com/lawsontyler/ghue/sdk/factory"
)

var cmdGroupsAll = &cobra.Command{
	Use:   "all",
	Short: "Get All groups: ghue groups all",
	Long:  `Get all groups: ghue groups all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd(factory.GetSdkClient(config.ReadConfig()))
	},
}

func allCmd(client *factory.SdkClient) {
	result, errHUE, err := groups.GetAllGroups(client)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
