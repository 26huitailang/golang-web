/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/app/service"
	"github.com/26huitailang/golang_web/library/mycrypto"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
	nickname string
)

// userCmd represents the server command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "user cli",
	Long: `user cli:

create -u USERNAME -p PASSWORD -nick NICKNAME
`,
}

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create user",
	Long:  `-u USERNAME -p PASSWORD -nick NICKNAME`,
	Args: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("args: %v", args)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		encPwd := mycrypto.Password(password).Encrypt(nil)
		user := &model.User{Username: username, Password: encPwd, Nickname: nickname}
		user, err := service.UserService.CreateUser(user)
		if err != nil {
			fmt.Printf("create failed: %v", err.Error())
			return
		}
		fmt.Printf("created: %v", user)
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	userCreateCmd.Flags().StringVarP(&username, "username", "u", "admin", "username")
	userCreateCmd.Flags().StringVarP(&password, "password", "p", "123123", "password")
	userCreateCmd.Flags().StringVarP(&nickname, "nickname", "n", "", "nickname, default ''")
}
