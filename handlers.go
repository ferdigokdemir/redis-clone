package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	w.Write([]byte(data[item.Key]))
	LogRequest(r)
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	data[item.Key] = item.Value
	LogRequest(r)
}

func FlushHandler(w http.ResponseWriter, r *http.Request) {
	data = make(map[string]string)
	os.Remove("db.json")
	LogRequest(r)
}

func StartHttpRouter() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/set", SetHandler)
	http.HandleFunc("/flush", FlushHandler)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}

func LogRequest(r *http.Request) {
	requestID := fmt.Sprintf("%d", time.Now().UnixNano())
	logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
}
