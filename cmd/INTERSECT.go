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

// intersect finds duplicate members for 2 given keys.
func intersect(key1, key2 string) {
	conn := GetCacheConn()
	
	//get first key
	mems1 := GetMembers(conn, key1)
	mems2 := GetMembers(conn, key2)

	// create map to improve search efficiency
	memsMap := make(map[string]string)
	for _, mem1 := range mems1 {
		memsMap[mem1] = mem1
	}
	
	// compare vals
	for _, mem2 := range mems2 {
		if memsMap[mem2] == mem2 {
			fmt.Printf("found equal values: %s\n", mem2)
		}
	}
}

// INTERSECTCmd
var INTERSECTCmd = &cobra.Command{
	Use:   "INTERSECT",
	Short: "INTERSECT",
	Long:  `INTERSECT`,
	Run: func(cmd *cobra.Command, args []string) {
		intersect(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(INTERSECTCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// KEYEXISTSCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// KEYEXISTSCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
