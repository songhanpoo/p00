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
	"strings"
	"github.com/songhanpoo/p00/common"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		domain, _ := cmd.Flags().GetString("domain")
		fHostRecord, _ := cmd.Flags().GetBool("find-host-record")
		extractLinks, _ := cmd.Flags().GetBool("extract-links")
		if (domain != "") && (help.ValidateDomainName(domain)) {
			if fHostRecord {
				fHostRecords(domain)
			} else if extractLinks {
				pageLinks(domain)
			} else {
				dnsLookup(domain)
				cloudflareHelper.GetZone("techtank9.com")
			}
		} else {
			fmt.Printf("Domain Name %s is invalid\n", domain)
		}
	},
}

func init() {
	rootCmd.AddCommand(dnsCmd)
	//
	dnsCmd.PersistentFlags().String("domain", "", "Use this DNS lookup tool to easily view the standard DNS records for a domain.")

	dnsCmd.Flags().BoolP("find-host-record", "f", false, "Find all Forward DNS (A) records for a domain.\nEnter a domain name and search for all subdomains associated with that domain.\nA handy reconnaissance tool when assessing an organisations security.")

	dnsCmd.Flags().BoolP("ip-lookup", "i", false, "Find the location of an IP address with this GeoIP lookup tool.")

	dnsCmd.Flags().BoolP("extract-links", "e", false, "This tool will parse the html of a website and extract links from the page. The hrefs or \"page links\" are displayed in plain text for easy copying or review.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func dnsLookup(domain string) {
	url := fmt.Sprintf("https://api.hackertarget.com/dnslookup/?q=%s", domain)
	initReq := help.NewHttpVar(url)
	resp := help.Req(initReq)
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Records"})
	d := strings.Split(string(resp), "\n")

	for _, v := range d {
		row := []interface{}{}
		row = append(row, strings.Trim(v, " "))
		tw.AppendRow(row)
	}
	tw.SetCaption("Result from " + help.DATETIME.Format("02-01-2006 15:04:05 Monday") + " / DD-MM-YYYY")

	fmt.Println(tw.Render())
}

func fHostRecords(domain string) {
	url := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", domain)
	initReq := help.NewHttpVar(url)
	resp := help.Req(initReq)
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Domain", "Ip Address"})
	d := strings.Split(string(resp), "\n")
	for _, v := range d {
		tmp := strings.Split(string(v), ",")
		row := []interface{}{}
		for _, indicate := range tmp {
			row = append(row, indicate)
		}
		tw.AppendRow(row)
	}
	tw.SetCaption("Result from " + help.DATETIME.Format("02-01-2006 15:04:05 Monday") + " / DD-MM-YYYY")
	fmt.Println(tw.Render())
}

func pageLinks(domain string) {
	url := fmt.Sprintf("https://api.hackertarget.com/pagelinks/?q=%s", domain)
	initReq := help.NewHttpVar(url)
	resp := help.Req(initReq)
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Links"})
	d := strings.Split(string(resp), "\n")
	for _, v := range d {
		row := []interface{}{}
		row = append(row, v)
		tw.AppendRow(row)
	}
	tw.SetCaption("Result from " + help.DATETIME.Format("02-01-2006 15:04:05 Monday") + " / DD-MM-YYYY")
	fmt.Println(tw.Render())
}
