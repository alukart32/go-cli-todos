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

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `The add command inserts new tasks into the list.
Each task is a set of words containing letters or numbers`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		for _, a := range args {
			if !validator.ValidateArgByRegex(a, validator.WORD_REGX) {
				return fmt.Errorf("invalid todo specified: %s", a)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		res := make(chan error)
		defer close(res)

		go manager.Add(res, args)
		if err := <-res; err != nil {
			log.Fatalf("\t%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
