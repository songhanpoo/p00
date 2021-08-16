package help

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

var DATETIME = time.Now()
var formatDate = DATETIME.Format("02-01-2006 15:04:05 Monday")

/*
Validate regex domain ( support only a sub )
Example :
	example.com 		-> true
	sub.example.com -> true
	moresub.sub.example.com -> false
*/
func ValidateDomainName(domain string) bool {
	RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
]{2,3})$`)

	return RegExp.MatchString(domain)
}

type HttpVar struct {
	Url        string
	Method     string
	AttrHeader map[string][]string
	Wrapper    interface{}
}

/*
Constructor HttpVar
	default url
	default GET
	default [string][]string
	default interface{}
*/
func NewHttpVar(url string) *HttpVar {
	return &HttpVar{
		Url:        url,
		Method:     "GET",
		AttrHeader: make(map[string][]string),
		Wrapper:    nil,
	}
}

/*
Request function
@flow1 simple request -> return []byte ( without wrapper )
@flow2 advanced request -> return wrapper input JSON
*/
func Req(h *HttpVar) []byte {
	req, err := http.NewRequest(
		h.Method, //method
		h.Url,    //url
		nil,      //body
	)

	if err != nil {
		log.Printf("Could not request a url. %v", err)
	}

	req.Header = h.AttrHeader
	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}
	out, err := ioutil.ReadAll(res.Body)

	//flow1 -> return byte
	if h.Wrapper == nil {
		return out
	}
	//flow2 -> return null bytes with value into wrapper
	if err := json.Unmarshal([]byte(out), &h.Wrapper); err != nil {
		log.Fatal(err)
	}
	return []byte{}
}
