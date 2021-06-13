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
	"log"

	"github.com/spf13/cobra"
)

// Append adds a member to a collection for a given key.
func AppendMember(k, v string) {
	conn := GetCacheConn()

	_, err := conn.Do("APPEND", k, v + ",")
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	defer conn.Close()
	fmt.Println("Added " + k + " to " + v)
}

// ADDCmd represents the ADD command
var ADDCmd = &cobra.Command{
	Use:   "ADD",
	Short: "Add a key and value",
	Long:  "Add a member to a collection for a given key.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("please supply a key and value to add")
			return 
		}
		AppendMember(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(ADDCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ADDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ADDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
