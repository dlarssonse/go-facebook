package facebook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

const (
	URL = "https://graph.facebook.com/"
)

// makeAPIRequest ...
func makeAPIRequest(method string, accessToken string, position int, path string, input interface{}, output interface{}) error {

	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	arguments := []string{}
	if accessToken != "" {
		arguments = append(arguments, "access_token="+accessToken)
	}

	if method == "GET" {
		if input != nil {
			typeOf := reflect.TypeOf(input)
			valueOf := reflect.ValueOf(input)
			if typeOf.Kind() == reflect.Ptr {
				typeOf = typeOf.Elem()
				valueOf = valueOf.Elem()
			}
			for i := 0; i < typeOf.NumField(); i++ {
				if fmt.Sprintf("%v", valueOf.Field(i)) != "<nil>" {
					jsonName := typeOf.Field(i).Tag.Get("json")
					jsonName = strings.Replace(jsonName, ",omitempty", "", -1)
					arguments = append(arguments, fmt.Sprintf("%v=%v", jsonName, valueOf.Field(i)))
				}
			}
		}
	}

	var req *http.Request
	if method == "GET" || method == "DELETE" {
		req, _ = http.NewRequest(method, URL+path+"?"+strings.Join(arguments, "&"), nil)
	} else {
		req, _ = http.NewRequest(method, URL+path, bytes.NewBuffer(data))
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 {
		if resp.StatusCode == 302 {
			return fmt.Errorf("status code is 302 (FOUND), probably a redirect request")
		} else if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 {
			body, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("status code is %d, body: %s", resp.StatusCode, body)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("status code is %d, body: %s", resp.StatusCode, body)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	if len(body) > 0 {
		if err := json.Unmarshal(body, output); err != nil {
			return err
		}
	}
	return nil
}
