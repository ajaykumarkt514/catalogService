package utils

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func ParseBody(r *http.Request, x interface{})  {
	if body,err := ioutil.ReadAll(r.Body); err == nil{
		err := json.Unmarshal([]byte(body),x)
		if err != nil{
			log.Fatal("Unable to Unmarshal check json data")
			panic(err)
		}
	}
}
