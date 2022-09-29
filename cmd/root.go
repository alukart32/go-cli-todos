/*
Copyright © 2022 alukart32

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"alukart32.com/todos/internal/usecase"
	"alukart32.com/todos/internal/usecase/repo"
	"alukart32.com/todos/pkg/boltdbx"
	"github.com/spf13/cobra"
)

var (
	manager usecase.ToDoListManager
	rootCmd = &cobra.Command{
		Use:   "todos",
		Short: "Tool that can be used to manage your TODOs",
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			boltdbx.Close()
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	repo := repo.NewToDoListRepo(boltdbx.New(), "todos")
	manager = usecase.NewManager(repo)
}
