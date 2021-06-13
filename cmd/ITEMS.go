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

// getItems returns all keys in the dictionary and all of their members.  Returns nothing if there are none.  Order is not guaranteed.
func getItems() string {
	conn := GetCacheConn()
	defer conn.Close()

	keys := GetKeys(conn)
	var items string
	for index, key := range keys {
		m := GetMembers(conn, key)
		items += fmt.Sprintf("%d) %s: %s\n", index + 1, key, m)
	}
	return items
}

// ITEMSCmd represents the ITEMS command
var ITEMSCmd = &cobra.Command{
	Use:   "ITEMS",
	Short: "Get everything",
	Long: `Get everything including keys and members`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getItems())
	},
}

func init() {
	rootCmd.AddCommand(ITEMSCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ITEMSCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ITEMSCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
