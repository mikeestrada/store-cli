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

	"github.com/spf13/cobra"
)

// RemoveMember removes a member from a key.
func RemoveMember(key, val string) error {
	
	ke := KeyExists(key)
	me := MemberExists(key, val)
	
	if me == false {
		return errors.New("ERROR, member does not exist")
	}
	if ke && me {
		// get vals for key
		members := GetMembers(key)
		
		var newMembers []string
		for i := 0; i < len(members); i++ {
			// remove match by creating new slice
			if members[i] != val {
				newMembers = append(newMembers, members[i])
			}
		}
		// update vals for key
		InsertAll(key, newMembers)
		fmt.Println("REMOVED")
		
	}
	return nil
}

// REMOVECmd represents the REMOVE command
var REMOVECmd = &cobra.Command{
	Use:   "REMOVE",
	Short: "Remove a key and value",
	Long: `Remove a key and value`,
	Run: func(cmd *cobra.Command, args []string) {
		e := RemoveMember(args[0], args[1])
		if e != nil {
			fmt.Println(e)
		}
	},
}

func init() {
	rootCmd.AddCommand(REMOVECmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// REMOVECmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// REMOVECmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
