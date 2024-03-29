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

// https://www.tujigu.com/x/82/
// themeCmd represents the theme command
var themeCmd = &cobra.Command{
	Use:   "theme",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("theme called")
	},
}

// themeDownloadCmd represents the themeDownload command
var themeDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("themeDownload called")
		t := downloadsuite.NewTheme(firstPage, folderSave)
		t.DownloadOneTheme()
	},
}

func init() {
	rootCmd.AddCommand(themeCmd)
	themeCmd.AddCommand(themeDownloadCmd)

	dir, _ := os.Getwd()
	themeDownloadCmd.Flags().StringVarP(&folderSave, "folder", "f", dir, "folder to save contents")
	themeDownloadCmd.Flags().StringVarP(&firstPage, "url", "u", "", "suite url")
	if err := suiteDownloadCmd.MarkFlagRequired("url"); err != nil {
		fmt.Errorf("required url %s", err)
	}
}
