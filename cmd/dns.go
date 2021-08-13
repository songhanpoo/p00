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
	"regexp"
	"time"
	"strings"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joho/godotenv"
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"github.com/spf13/cobra"
)
// Global variable
var _urlBase    = "https://api.hackertarget.com"
var dt          = time.Now()
var formatDatet = dt.Format("02-01-2006 15:04:05 Monday")
// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS lookup, Reverse DNS",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		domain, _ := cmd.Flags().GetString("domain")
		ip, _ := cmd.Flags().GetString("ip")
		fHostRecord, _ := cmd.Flags().GetBool("find-host-record")
		ipLookup, _ := cmd.Flags().GetBool("ip-lookup")
		extractLinks,_ := cmd.Flags().GetBool("extract-links")

		if (domain != "") && (validateDomainName(domain)) {
			if fHostRecord {
				fHostRecords(domain)
			} else if extractLinks {
				pagelinks(domain)
			} else {
				dnsLookup(domain)
			}
		} else {
			fmt.Printf("Domain Name %s is invalid\n", domain)
		}

		if ip != "" {
			if ipLookup {
				geoIP(ip)
			} else {
				reverseDNS(ip)
			}
		}

	},
}


func init() {
	rootCmd.AddCommand(dnsCmd)

	dnsCmd.PersistentFlags().String("domain", "", "Use this DNS lookup tool to easily view the standard DNS records for a domain.")

	dnsCmd.Flags().BoolP("find-host-record", "f", false, "Find all Forward DNS (A) records for a domain.\nEnter a domain name and search for all subdomains associated with that domain.\nA handy reconnaissance tool when assessing an organisations security.")

	dnsCmd.PersistentFlags().String("ip", "", "Discover the reverse DNS entries for an IP address,a range of IP addresses or a domain name.\nIP based reverse DNS lookups will resolve the IP addresses in real time,\nwhile the domain name or hostname search uses a cached database.")

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
	url := _urlBase + "/dnslookup/?q=" + domain
	fmt.Println(url)
	responseBytes := reqGET(url)

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Records"})
	d := strings.Split(string(responseBytes), "\n")
	
	for _, v := range d {
		row := []interface{}{}
		row = append(row,strings.Trim(v," "))
		tw.AppendRow(row)
	}
	tw.SetCaption("Result from " + dt.Format("02-01-2006 15:04:05 Monday") + " / DD-MM-YYYY")

	fmt.Println(tw.Render())

}

func reverseDNS(ip string) {
	url := _urlBase + "/reversedns/?q=" + ip
	fmt.Println(url)
	responseBytes := reqGET(url)

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Ip Address","Domain"})
	d := strings.Split(string(responseBytes), "\n")
	
	for _, v := range d {
		tmp := strings.Split(string(v), " ")
		row := []interface{}{}
		for _, indicate := range tmp {
			row = append(row, indicate)
		}
		tw.AppendRow(row)
	}
	tw.SetCaption("Result from " + dt.Format("02-01-2006 15:04:05 Monday") + " / DD-MM-YYYY")
	fmt.Println(tw.Render())
}

func fHostRecords(domain string) {
	url := _urlBase + "/hostsearch/?q=" + domain
	fmt.Println(url)
	responseBytes := reqGET(url)
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Domain","Ip Address"})
	d := strings.Split(string(responseBytes), "\n")
	for _, v := range d {
		tmp := strings.Split(string(v), ",")
		row := []interface{}{}
		for _, indicate := range tmp {
			row = append(row, indicate)
		}
		tw.AppendRow(row)
	}
	tw.SetCaption("Result from " + dt.Format("02-01-2006 15:04:05 Monday") + " / DD-MM-YYYY")
	fmt.Println(tw.Render())
}

func geoIP(ip string) {
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

func pagelinks(domain string) {
	url := _urlBase + "/pagelinks/?q=" + domain
	fmt.Println(url)
	responseBytes := reqGET(url)
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Links"})
	d := strings.Split(string(responseBytes), "\n")
	for _, v := range d {
		row := []interface{}{}
		row = append(row, v)
		tw.AppendRow(row)
	}
	tw.SetCaption("Result from " + dt.Format("02-01-2006 15:04:05 Monday") + " / DD-MM-YYYY")
	fmt.Println(tw.Render())
}

func reqGET(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
			log.Printf("Could not request a baseAPI. %v", err)
	}

	// request.Header.Add("Accept", "application/json")
	// request.Header.Add("User-Agent", "P00 CLI (https://github.com/songhanpoo/p00)")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

func validateDomainName(domain string) bool {

	RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
]{2,3})$`)

	return RegExp.MatchString(domain)
}