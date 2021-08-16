// package cloudflareHelper

// import (
// 	"log"
// 	"github.com/songhanpoo/p00/common"
// )

// type Zone struct {
// 	Id string
// 	Path string `default:"match=all"`
// }



// func getZone(zone Zone) string {
// 	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records?%s",zone.Id, zone.Path)

// 	wrapper := struct {
// 		success string `json:"success"`
// 		result []string `json:"result"`
// 	}

// 	err := help.FetchJSON(url,&wrapper)

// 	fmt.Printf("%v",wrapper)

// }