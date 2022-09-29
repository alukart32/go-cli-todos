/*
Copyright © 2022 alukart32 alukart32work@gmail.com
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		res := make(chan error)
		go manager.List(res)
		defer close(res)

		if err := <-res; err != nil {
			log.Printf("\t%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
