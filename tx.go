// Copyright 2018 The QOS Authors

package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
)

var (
	height  int64
	txindex string
	remote  string
)

func ParseTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse",
		Short: "parse tx",
		RunE: func(cmd *cobra.Command, args []string) error {

			tmc := client.NewHTTP(remote, "/websocket")
			b, err := tmc.Block(&height)
			if err != nil {
				return err
			}

			if b.Block.NumTxs == 0 {
				fmt.Printf("没有交易\n")
				return nil
			}

			txDecoder := auth.DefaultTxDecoder(cdc)

			for k, v := range b.Block.Data.Txs {
				tx, err := txDecoder(v)
				if err != nil {
					return err
				}

				fmt.Printf("tx%d, %+v\n", k, tx)
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&txindex, "txindex", "", "块中交易序号，逗号分割 1,2")
	cmd.PersistentFlags().StringVar(&remote, "remote", "127.0.0.1:26657", "node地址")
	cmd.PersistentFlags().Int64Var(&height, "height", 1, "the name of node")
	return cmd
}
