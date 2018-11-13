package main

import (
	_ "github.com/cosmos/cosmos-sdk/client/lcd/statik"
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	"github.com/spf13/cobra"
)


// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "decode",
		Short: "decode chain tx",
	}
)

// go run main.go parse --remote=http://13.57.187.196:26657 --height=31998
func main() {
	cobra.EnableCommandSorting = false
	cdc := app.MakeCodec()


	// TODO: setup keybase, viper object, etc. to be passed into
	// the below functions and eliminate global vars, like we do
	// with the cdc
	rootCmd.AddCommand(ParseTxCmd(cdc))


	err := rootCmd.Execute()

	if err != nil {
		// handle with #870
		panic(err)
	}
}