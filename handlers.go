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

/**
 * @api {post} /getItem Request Item by Key
 * @apiName GetItem
 * @apiGroup Item
 *
 * @apiBody {String} key Key of the item to get
 *
 * @apiSuccess {String} key Key of the item
 * @apiSuccess {String} value Value of the item
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "key": "name",
 *       "value": "Ferdi Gökdemir"
 *     }
 *
 * @apiError KeyNotFound Key not found
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error": "KeyNotFound"
 *     }
 */

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	LogRequest(r)
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	if value, ok := data[item.Key]; ok {
		item.Value = value
		SuccessResponse(w, item)
	} else {
		ErrorResponse(w, "KeyNotFound")
	}
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	data[item.Key] = item.Value
	SuccessResponse(w, nil)
}

func FlushHandler(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	data = make(map[string]string)
	err := os.Remove("db.json")
	if err != nil {
		ErrorResponse(w, err.Error())
	} else {
		SuccessResponse(w, nil)
	}
}

func StartHttpRouter() {
	http.HandleFunc("/api/v1/getItem", GetHandler)
	http.HandleFunc("/api/v1/setItem", SetHandler)
	http.HandleFunc("/api/v1/flushDatabase", FlushHandler)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}

func LogRequest(r *http.Request) {
	requestID := fmt.Sprintf("%d", time.Now().UnixNano())
	logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
}

func ErrorResponse(w http.ResponseWriter, err string) {
	json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": err})
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": data})
}
