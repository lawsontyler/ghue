package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/lawsontyler/ghue/cli/config"
	"github.com/lawsontyler/ghue/cli/groups"
	"github.com/lawsontyler/ghue/cli/info"
	"github.com/lawsontyler/ghue/cli/internal"
	"github.com/lawsontyler/ghue/cli/lights"
	"github.com/lawsontyler/ghue/cli/rules"
	"github.com/lawsontyler/ghue/cli/scenes"
	"github.com/lawsontyler/ghue/cli/schedules"
	"github.com/lawsontyler/ghue/cli/sensors"
	"github.com/lawsontyler/ghue/cli/update"
	"github.com/lawsontyler/ghue/cli/version"
)

var rootCmd = &cobra.Command{
	Use:   "ghue",
	Short: "Hue Cli",
	Long:  `Golang Hue Cli`,
}

func main() {
	addCommands()
	rootCmd.PersistentFlags().BoolVarP(&internal.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&internal.Format, "format", "f", "pretty", "choose format output. One of 'json', 'yaml' and 'pretty'")
	rootCmd.PersistentFlags().StringVarP(&config.ConfigFile, "configFile", "c", internal.Home+"/.ghue/config.json", "configuration file, default is "+internal.Home+"/.ghue/config.json")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//AddCommands adds child commands to the root command rootCmd.
func addCommands() {
	rootCmd.AddCommand(config.Cmd)
	rootCmd.AddCommand(lights.Cmd)
	rootCmd.AddCommand(groups.Cmd)
	rootCmd.AddCommand(schedules.Cmd)
	rootCmd.AddCommand(sensors.Cmd)
	rootCmd.AddCommand(scenes.Cmd)
	rootCmd.AddCommand(info.Cmd)
	rootCmd.AddCommand(rules.Cmd)
	rootCmd.AddCommand(update.Cmd)
	rootCmd.AddCommand(version.Cmd)
}
