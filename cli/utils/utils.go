package utils

import "github.com/lawsontyler/ghue/cli/internal"

func GetVerbose() bool {
	return internal.Verbose
}

func SetVerbose(v bool) {
	internal.Verbose = v
}

func GetOutputFormat() string {
	return internal.Format
}

func SetOutputFormat(f string) {
	internal.Format = f
}

func GetHomeDir() string {
	return internal.Home
}