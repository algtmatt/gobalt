package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobalt/src/config"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Transport struct {
	opts config.Opts
	S    *Session
}

func New(c config.Opts, s *Session) *Transport {
	t := &Transport{
		opts: c,
		S:    s,
	}
	return t
}

//func (t *Transport) Fetch(service string, data ...map[string]string) (body string, err error) {
func (t *Transport) Fetch(service string, data ...map[string]string) (body []byte, err error) {

	var url = strings.Trim(t.opts.Url, "/") + ":" + strconv.Itoa(t.opts.Port) + service
	req := &http.Request{}
	var reader io.Reader
	// Build POST data
	if len(data) > 0 {
		postdata, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(postdata))
	} else {
		req, err = http.NewRequest("GET", url, reader)
	}
	fmt.Printf("During fetch found auth set to: %s\n", t.S.Token)
	if t.S.Token != "badauth" {
		fmt.Printf("Found token. Re-using token: %s\n", t.S.Token)
		req.Header.Set("X-Auth-Token", t.S.Token)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	d, err := ioutil.ReadAll(response.Body)
	fmt.Printf("Fetched %s and received %s\n", url, string(d))

	//return string(d), nil
	return d, nil

}
