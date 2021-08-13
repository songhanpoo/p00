package common

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"log"
	"encoding/json"
)

func reqGET(baseAPI string) []byte {
	req, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
			log.Printf("Could not request a baseAPI. %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read res body. %v", err)
	}

	return raw
}

func fetchJSON(url string, wrapper interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	return dec.Decode(wrapper)
}

func validateDomainName(domain string) bool {
	RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
]{2,3})$`)

	return RegExp.MatchString(domain)
}
