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

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

// GetMembers returns the collection of strings for the given key.
func GetKeys(r redis.Conn) []string {
	var conn redis.Conn
	if r == nil {
		conn = GetCacheConn()
		defer conn.Close()
	} else {
		conn = r
	}

	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		fmt.Printf("error getting KEYS: %s", err)
	}
	
	return keys
}

// KEYSCmd represents the KEYS command
var KEYSCmd = &cobra.Command{
	Use:   "KEYS",
	Short: "Get all keys",
	Long:  "Returns a collection of keys",
	Run: func(cmd *cobra.Command, args []string) {
		keys := GetKeys(nil)
		for index, key := range keys {
			fmt.Printf("%d) %s\n", index + 1, key)
		}
	},
}

func init() {
	rootCmd.AddCommand(KEYSCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MEMBEREXISTSCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MEMBEREXISTSCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
