/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
)

func getAllMembers() []string {
	keys := GetKeys()
	var allMembers []string
	for _, key := range keys {
		mems := GetMembers(key)
		allMembers = append(allMembers, mems...)
	}
	return allMembers
}

// ALLMEMBERSCmd represents the ALLMEMBERS command
var ALLMEMBERSCmd = &cobra.Command{
	Use:   "ALLMEMBERS",
	Short: "Get all members (keys)",
	Long: `Get all members (keys)`,
	Run: func(cmd *cobra.Command, args []string) {
		members := getAllMembers()
		for i := 0; i < len(members); i++ {
			fmt.Printf("%d) %s \n", i + 1, members[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(ALLMEMBERSCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ALLMEMBERSCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ALLMEMBERSCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
