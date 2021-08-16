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
	"github.com/26huitailang/golang_web/app/service/downloadsuite"
	"os"

	"github.com/spf13/cobra"
)

// https://www.meituri.com/a/26718/
// https://www.tujigu.com/a/26718/
// suiteCmd represents the suite command
var suiteCmd = &cobra.Command{
	Use:   "suite",
	Short: "suite commands",
	Long: `suite commands to handle a suite's operation:

gws suite download --url xxxxx`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("suite called", firstPage)
	//},
}

var suiteDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download one suite",
	Long: `For example:

gws suite download --url HOST [--folder .]`,
	Args: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("args: %v", args)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
		operator := downloadsuite.NewMeituriSuite(firstPage, folderSave, &downloadsuite.MeituriParser{})
		s := downloadsuite.NewSuite(operator)
		s.Download()
	},
}

func init() {
	suiteCmd.AddCommand(suiteDownloadCmd)
	rootCmd.AddCommand(suiteCmd)

	dir, _ := os.Getwd()
	suiteDownloadCmd.Flags().StringVarP(&folderSave, "folder", "f", dir, "folder to save contents")
	suiteDownloadCmd.Flags().StringVarP(&firstPage, "url", "u", "", "suite url")
	if err := suiteDownloadCmd.MarkFlagRequired("url"); err != nil {
		fmt.Errorf("required url %s", err)
	}
}

func isValidURL(url string) bool {
	return true
}
