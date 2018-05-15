package sensors

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/sensors"
	"github.com/lawsontyler/ghue/sdk/common"
)

var cmdSensorsGet = &cobra.Command{
	Use:     "get",
	Short:   "Get sensor attributes and state: ghue sensor <id>",
	Long:    `Get sensor attributes and state: ghue sensor <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue sensors state --help")
		} else {
			getCmd(config.ReadConfig(), args[0])
		}
	},
}

func getCmd(connection *common.Connection, id string) {
	result, errHUE, err := sensors.GetSensor(connection, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
