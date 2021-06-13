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
	"errors"
	"fmt"
	"log"
	"github.com/spf13/cobra"
)

// RemoveMember removes all members for a key and removes the key from the dictionary. Returns an error if the key does not exist.
func RemoveAll(key string) error {
	conn := GetCacheConn()
	defer conn.Close()
	
	if KeyExists(conn, key) == false {
		return errors.New("ERROR, key does not exist")
	}

	_, err := conn.Do("DEL", key)
    return err
}

// REMOVEALLCmd represents the REMOVE command
var REMOVEALLCmd = &cobra.Command{
	Use:   "REMOVEALL",
	Short: "Removes a key and values",
	Long: `Removes a key and values`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("please supply a key to remove values for")
			return 
		}
		e := RemoveAll(args[0])
		if e != nil {
			fmt.Println(e)
		}
	},
}

func init() {
	rootCmd.AddCommand(REMOVEALLCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// REMOVECmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// REMOVECmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
