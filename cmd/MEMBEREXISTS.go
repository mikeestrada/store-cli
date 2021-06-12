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

// MemberExists returns whether a key exists or not.
func MemberExists(k, v string) bool {
	conn := GetCacheConn()
	defer conn.Close()

	members := GetMembers(k)
	for i := 0; i < len(members); i++ {
		if members[i] == v {
			return true
		}
	}
	return false

	// existsBytes, err := conn.Do("EXISTS", k)
	// if err != nil {
	// 	fmt.Printf("error checking if member %s exists: %s", k, err)
	// 	return false
	// }
	// exists := fmt.Sprint(existsBytes.(int64))
	// if exists == "1" {
	// 	return true
	// }
	// return false
}

// MEMBEREXISTSCmd represents the MEMBEREXISTS command
var MEMBEREXISTSCmd = &cobra.Command{
	Use:   "MEMBEREXISTS",
	Short: "Is the value here?",
	Long: `Is the value here?`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(MemberExists(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(MEMBEREXISTSCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MEMBEREXISTSCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MEMBEREXISTSCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
