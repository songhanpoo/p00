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
	// "regexp"
	// "time"
	"strings"
	"github.com/jedib0t/go-pretty/v6/table"
	// "net/http"
	// "io/ioutil"
	"fmt"
	// "log"
	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ip called")
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)

	ipCmd.PersistentFlags().String("ip", "", "Discover the reverse DNS entries for an IP address,a range of IP addresses or a domain name.\nIP based reverse DNS lookups will resolve the IP addresses in real time,\nwhile the domain name or hostname search uses a cached database.")

	ipCmd.Flags().BoolP("ip-lookup", "i", false, "Find the location of an IP address with this GeoIP lookup tool.")

	ipCmd.Flags().BoolP("extract-links", "e", false, "This tool will parse the html of a website and extract links from the page. The hrefs or \"page links\" are displayed in plain text for easy copying or review.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func geoIP(ip string) {
	fetchURL := fmt.Sprintf("https://dns.bufferover.run/dns?q=.%s", domain)
	url := _urlBase + "/geoip/?q=" + ip
	fmt.Println(url)
	responseBytes := reqGET(url)
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Domain","Ip Address"})
	d := strings.Split(string(responseBytes), "\n")
	for _, v := range d {
		tmp := strings.Split(string(v), ":")
		row := []interface{}{}
		for _, indicate := range tmp {
			row = append(row, indicate)
		}
		tw.AppendRow(row)
	}
	tw.SetCaption("Result from " + dt.Format("02-01-2006 15:04:05 Monday") + " / DD-MM-YYYY")
	fmt.Println(tw.Render())
}