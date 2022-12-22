// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"github.com/spf13/cobra"

	"github.com/ava-labs/xsvm/cmd/chain/create"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:   "chain",
		Short: "Manages XS chains",
	}
	c.AddCommand(
		create.Command(),
	)
	return c
}