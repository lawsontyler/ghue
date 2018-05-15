package groups

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/sdk/groups"
	"strconv"
	"github.com/lawsontyler/ghue/sdk/common"
)

var (
	name      string
	lights    []int
	groupType string
)

func init() {
	cmdGroupsCreate.Flags().IntSliceVarP(&lights, "light","l", []int{}, "The Light IDs to add to the group")
	cmdGroupsCreate.Flags().StringVar(&name, "name", "", "Name of the Group.")
	cmdGroupsCreate.Flags().StringVar(&groupType, "group-type", "", "Group Type: lightgroup")

	cmdGroupsCreate.MarkFlagRequired("name")
	cmdGroupsCreate.MarkFlagRequired("light")
}

var cmdGroupsCreate = &cobra.Command{
	Use:   "create",
	Short: "Create group: ghue groups create ...",
	Long:  `Create group: ghue groups create ...`,

	Run: func(cmd *cobra.Command, args []string) {
		createGroupCmd(config.ReadConfig())
	},
}

func createGroupCmd(connection *common.Connection) {
	// I'm intentionally getting this as ints off the command line
	// I figure, why not let Cobra take care of the validation?  Converting it to strings is easy.

	var lightStrings []string

	for _, element := range lights {
		lightStrings = append(lightStrings, strconv.Itoa(element))
	}

	group := groups.Create{
		Name: name,
		Lights: lightStrings,
	}

	// I _could_ restrict this to the currently known types...but seeing as the Hue API is going to change, I'll just
	// let the hub return the error relative to its own API version.
	if groupType != "" {
		group.Type = groupType
	}

	result, errHUE, err := groups.CreateAPI(connection, &group)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}

