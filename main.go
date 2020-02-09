package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", responseHandler)
	port := "8080"
	fmt.Println("listen to localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func responseHandler(w http.ResponseWriter, r *http.Request) {

	var (
		statusCode = 200
		err        error
	)

	values := r.URL.Query()

	msg := values.Get("msg")
	if msg == "" {
		msg = "nenhuma mensagem passada"
	}

	code := values.Get("status")
	if code != "" {
		statusCode, err = strconv.Atoi(code)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	headersSet := values.Get("header")
	if headersSet != "" {
		headers := strings.Split(headersSet, ",")
		for _, header := range headers {
			h := strings.Split(header, "-")
			if len(h)%2 == 0 {
				w.Header().Set(h[0], h[1])
			}
		}
	}

	jsonValue := values.Get("json")
	if jsonValue == "" {
		http.Error(w, msg, statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	w.Write([]byte(jsonValue))

	return
}
