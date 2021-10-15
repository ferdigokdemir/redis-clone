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
 * @apiName GetItemHandler
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
 *     HTTP/1.1
 *     {
 *       "error": "KeyNotFound"
 *     }
 */
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
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

/**
 * @api {post} /setItem Save Item
 * @apiName SetItemHandler
 * @apiGroup Item
 *
 * @apiBody {String} key Key of the item to set
 * @apiBody {String} value Value of the item to set
 *
 * @apiSuccess {String} success Status of the operation
 * @apiSuccess {String} data Value of the item
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "success": true,
 *       "data": null
 *     }
 *
 * @apiError SetItemError Item could not be set
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1
 *     {
 *	 	 "success": false,
 *       "error": "SetItemError"
 *     }
 */
func SetItemHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	LogRequest(r)
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	data[item.Key] = item.Value
	SuccessResponse(w, nil)
}

/**
 * @api {post} /flushItems Flush all items
 * @apiName FlushItemsHandler
 * @apiGroup Item
 *
 * @apiSuccess {String} success Status of the operation
 * @apiSuccess {String} data Value of the operation
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "success": true,
 *       "data": null
 *     }
 *
 * @apiError FlushItemsError Items could not be flushed
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1
 *     {
 *	 	 "success": false,
 *       "error": "FlushItemsError"
 *     }
 */
func FlushItemsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
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
	http.HandleFunc("/api/v1/getItem", GetItemHandler)
	http.HandleFunc("/api/v1/setItem", SetItemHandler)
	http.HandleFunc("/api/v1/flushItems", FlushItemsHandler)
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

func EnableCors(w *http.ResponseWriter) {
	cors := os.Getenv("ENABLE_CORS")
	if cors == "true" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-Methods", "POST")
		(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}
