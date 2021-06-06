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

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

// MEMBEREXISTSCmd represents the MEMBEREXISTS command
var MEMBERSCmd = &cobra.Command{
	Use:   "MEMBERS",
	Short: "Get values for key",
	Long:  `Get values for key`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MEMBERS called")
		conn, err := redis.Dial("tcp", ":6379")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		n, err := conn.Do("GET", args[0])
		if err != nil {
			fmt.Printf("error: %s", err)
		}
		fmt.Println(string(n.([]byte)))
	},
}

func init() {
	rootCmd.AddCommand(MEMBERSCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MEMBEREXISTSCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MEMBEREXISTSCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
