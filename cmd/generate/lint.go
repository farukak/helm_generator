/*
Copyright © 2022 FARUK AK <KAKURAF@GMAIL.COM>

*/
package generate

import (
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	GenerateCmd.AddCommand(lintCmd)
}
