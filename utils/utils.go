package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}