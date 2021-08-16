package cloudflareHelper

import (
	"log"
	"github.com/songhanpoo/p00/common"
)

func GetZone(url string) string {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones?name=%s&match=all",url)

	wrapper := struct {
		Success string `json:"success"`
		Result 	[]struct {
			ID                  string      `json:"id"`
			Name                string      `json:"name"`
			Status              string      `json:"status"`
			Paused              bool        `json:"paused"`
			Type                string      `json:"type"`
			DevelopmentMode     int         `json:"development_mode"`
			NameServers         []string    `json:"name_servers"`
			OriginalNameServers []string    `json:"original_name_servers"`
			OriginalRegistrar   string      `json:"original_registrar"`
			OriginalDnshost     interface{} `json:"original_dnshost"`
		}
	}

	tmpHeader :=  map[string][]string{
		"Content-Type": []string{"application/json"},
		"x-auth-email": []string{"dasdasdae@e2123ds.com"},
    "x-auth-key": []string{"50fb5b628a143da7bb7a43a346aef5624b0f3"},
	}
	

	tmpHttpVar := &HttpVar{
		Url: url,
		Method: "GET", //default GET
		AttrHeader: tmpHeader,
		Wrapper: &wrapper,
	}
	_ = help.Req(tmpHttpVar)

	fmt.Printf("%v",tmpHttpVar.Wrapper)
}