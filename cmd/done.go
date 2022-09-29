/*
Copyright © 2022 alukart32 alukart32work@gmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"alukart32.com/todos/pkg/validator"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Remove a todo from list by id",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}

		for _, a := range args {
			if !validator.ValidateArgByRegex(a, validator.NUMBER_REGX) {
				return fmt.Errorf("invalid todo task specified: %s", a)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		res := make(chan error)
		defer close(res)

		go manager.Done(res, args[0])

		if err := <-res; err != nil {
			log.Printf("\t%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
