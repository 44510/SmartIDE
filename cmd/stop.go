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
	"os"
	"os/exec"

	"github.com/leansoftX/smartide-cli/lib/common"
	"github.com/leansoftX/smartide-cli/lib/i18n"
	"github.com/spf13/cobra"
) // stopCmd represents the stop command

var instanceI18nStop = i18n.GetInstance().Stop

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: instanceI18nStop.Info.Help_short,
	Long:  instanceI18nStop.Info.Help_long,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(instanceI18nStop.Info.Info_start)

		//1. 获取docker compose的文件内容
		var yamlFileCongfig YamlFileConfig
		yamlFileCongfig.GetConfig()
		dockerCompose, _, _ := yamlFileCongfig.ConvertToDockerCompose()
		servicename := yamlFileCongfig.Workspace.DevContainer.ServiceName

		fmt.Printf("docker-compose servicename: %v \n", servicename)

		yamlFilePath, _ := dockerCompose.GetDockerComposeFile(servicename)

		pwd, _ := os.Getwd()
		composeCmd := exec.Command("docker-compose", "-f", yamlFilePath, "--project-directory", pwd, "stop")
		composeCmd.Stdout = os.Stdout
		composeCmd.Stderr = os.Stderr
		if composeCmdErr := composeCmd.Run(); composeCmdErr != nil {
			common.SmartIDELog.Fatal(composeCmdErr)
		}

		fmt.Println(instanceI18nStop.Info.Info_end)

	},
}

func init() {
	//rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
