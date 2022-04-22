/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"html/template"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var fields string
var title string
var swalId string

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Generate Angular Sweetalert html form",
	Long: `Generate Angular Sweetalert html form. 
	Syntax:
	ng-sweetalert-gen html -t "title" -i formId -f "field1Name:field1HtmlType,field2Name:field2HtmlType,...,fieldNName:fieldNHtmlType"
	
	Usage:
	ng-sweetalert-gen html -t "Create student" -i student -f "name:text,age:number,message:textarea"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		split := strings.Split(fields, ",")
		fmt.Println("html split : ", split)
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
		fmt.Println("html elements : ", elements)
		tmplFile, _ := swal_form.ReadFile("templates/swal_form.html.gotmpl")
		tmpl, err := template.New("swal_form").Funcs(template.FuncMap{
			"toLower": strings.ToLower,
			"toTitle": func(i interface{}) string {
				return strings.Title(fmt.Sprintf("%v", i))
			},
			"isTextArea": func(e FormElementType) bool {
				return e == Textarea
			},
			"isSelect": func(e FormElementType) bool {
				return e == Select
			},
		}).Parse(string(tmplFile))

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if err = tmpl.Execute(os.Stdout, map[string]interface{}{"title": title, "swalId": swalId, "elements": elements}); err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(htmlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// htmlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// htmlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	htmlCmd.Flags().StringVarP(&title, "title", "t", "", "Sweetalert form title")
	htmlCmd.MarkFlagRequired("title")
	htmlCmd.Flags().StringVarP(&swalId, "id", "i", "", "Sweetalert form id")
	htmlCmd.MarkFlagRequired("swalId")
	htmlCmd.Flags().StringVarP(&fields, "fields", "f", "", "Sweetalert form fields")
	htmlCmd.MarkFlagRequired("fields")
}
