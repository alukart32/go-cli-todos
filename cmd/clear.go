/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear todos list",
	Run: func(cmd *cobra.Command, args []string) {
		res := make(chan error)
		defer close(res)

		go manager.Clear(res)

		if err := <-res; err != nil {
			log.Printf("\t%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
