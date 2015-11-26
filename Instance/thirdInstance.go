package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"lab3/structure"
	"net/http"
	"strconv"
)

var pairMap map[int]structure.KeyValuePair

func PutKeyPair(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var newPair structure.KeyValuePair
	key := convert(p.ByName("key"))
	value := p.ByName("value")

	newPair.Key = key
	newPair.Value = value
	pairMap[key] = newPair
	result, _ := json.Marshal(newPair)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", result)
}

func GetKeyPair(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	inputKey := convert(p.ByName("key"))
	//fmt.Println(inputKey)
	var findPair structure.KeyValuePair
	for key, value := range pairMap {
		if key == inputKey {
			findPair.Key = inputKey
			findPair.Value = value.Value
		}
	}
	result, _ := json.Marshal(findPair)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", result)
}

func GetAllPairs(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var pairs []structure.KeyValuePair
	for key, value := range pairMap {
		tempPair := structure.KeyValuePair{
			key,
			value.Value,
		}
		pairs = append(pairs, tempPair)
	}
	var allPair structure.AllPair
	allPair.Pairs = pairs

	result, _ := json.Marshal(allPair)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", result)
}

func convert(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	pairMap = make(map[int]structure.KeyValuePair)
	firstRouter := httprouter.New()
	firstRouter.PUT("/keys/:key/:value", PutKeyPair)
	firstRouter.GET("/keys/:key", GetKeyPair)
	firstRouter.GET("/keys", GetAllPairs)
	firstServer := http.Server{
		Addr:    "0.0.0.0:3002",
		Handler: firstRouter,
	}
	firstServer.ListenAndServe()
}
