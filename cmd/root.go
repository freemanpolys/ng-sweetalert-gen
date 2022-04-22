/*
Copyright © 2022 James Kokou GAGLO <freemanpolys@gmail.com>

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
	"strings"

	"github.com/metal3d/go-slugify"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var fields string
var title string
var swalId string
var toFile bool = true

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ng-sweetalert-gen",
	Short: "Generate Angular Sweetalert html and component form",
	Long: `Generate Angular Sweetalert html component form. 
	Syntax:
	ng-sweetalert-gen  -t "title" -i formId -f "field1Name:field1HtmlType,field2Name:field2HtmlType,...,fieldNName:fieldNHtmlType"
	
	Usage:
	ng-sweetalert-gen  -t "Create student" -i student -f "name:text,age:number,message:textarea"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		split := strings.Split(fields, ",")
		elements := make([]FormElement, 0)
		for _, v := range split {
			if strings.Contains(v, ":") {
				field := strings.Split(v, ":")
				element, err := GetHtmlFormInputType(field[1])
				if err != nil {
					fmt.Println("Field type '", field[1], "' not supported")
					os.Exit(0)
				}
				element.Value = field[0]
				elements = append(elements, element)
			} else {
				fmt.Println("Syntax error with field '", v, "'. Run 'ng-sweetalert-gen html -h' for help.")
				os.Exit(0)
			}
		}
		swalId = slugify.Marshal(swalId)
		data := map[string]interface{}{"title": title, "swalId": swalId, "elements": elements}

		htmlFile, err := swalForm.ReadFile("templates/swal_form.html.gotmpl")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		ProcessTmplFiles(".", swalId+"-swal-form.html", htmlFile, data, !toFile)

		tsFile, err := formGroup.ReadFile("templates/form_group.ts.gotmpl")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		ProcessTmplFiles(".", swalId+"-componnent.ts", tsFile, data, !toFile)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&title, "title", "t", "", "Sweetalert form title")
	rootCmd.MarkFlagRequired("title")
	rootCmd.Flags().StringVarP(&swalId, "id", "i", "", "Sweetalert form id")
	rootCmd.MarkFlagRequired("swalId")
	rootCmd.Flags().StringVarP(&fields, "fields", "f", "", "Sweetalert form fields")
	rootCmd.MarkFlagRequired("fields")
	//rootCmd.Flags().BoolVarP(&toFile, "write-to-file", "w", true, "Write generated output to file")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ng-sweetalert-gen" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ng-sweetalert-gen")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
