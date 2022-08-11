/*
Copyright © 2022 FARUK AK <KAKURAF@GMAIL.COM>

*/
package generate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	image string
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Helm Chart",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	GenerateCmd.Flags().StringVarP(&image, "image", "i", "", "The image of Helm Chart")

	if err := GenerateCmd.MarkFlagRequired("image"); err != nil {
		fmt.Println(err)
	}
}
