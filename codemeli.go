package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["codemeli"]

	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'codemeli' is missing")
		return
	}
	if codeMeliValidator(keys[0]) {
		fmt.Fprintf(w, "Url Param 'codemeli' is Valid: "+string(keys[0]))
	} else {
		fmt.Fprintf(w, "Url Param 'codemeli' is: "+string(keys[0])+" and it is Invalid")
	}
}

func codeMeliValidator(key string) bool {
	sumValid := 0
	keyR := reverse(key)
	var controlNumber int64
	if len(keyR) != 9 {
		splitedText := strings.SplitAfter(keyR, "")
		for k, v := range splitedText {
			if k != 0 {
				v, _ := strconv.ParseInt(v, 10, 64)
				sumValid = sumValid + (int(v) * (k + 1))
			} else {
				controlNumber, _ = strconv.ParseInt(v, 10, 64)
			}
		}
		remaining := sumValid % 11
		if remaining < 2 && int(controlNumber) == remaining {
			return true
		} else if int(controlNumber) == (11 - remaining) {
			return true
		}
	}
	return false
}
