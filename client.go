package main

import (
	"fmt"
	"lab3/hash"
	"lab3/structure"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	firstServer  = "3000"
	secondServer = "3001"
	thirdServer  = "3002"
)

// The input whichServer is an Integer, regarded as 3000, 3001 or 3002
func PutKey(pair structure.KeyValuePair, whichServer string) bool {
	urlPath := "http://localhost:"
	urlPath += whichServer
	urlPath += "/keys/" + convertInt(pair.Key) + "/" + pair.Value
	fmt.Println(urlPath)
	client := &http.Client{}
	request, err := http.NewRequest("PUT", urlPath, nil)
	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Put key operation wrong ", err)
		panic(err)
	}
	if response != nil {
		return true
	} else {
		return false
	}
}

func convertInt(key int) string {
	return strconv.Itoa(key)
}

func main() {

	// Initialize hash ring
	cHashRing := Hash.NewConsistent()
	//NewNode(id, ip, weight)
	cHashRing.Add(Hash.NewNode(0, "http://localhost:3000", 1))
	cHashRing.Add(Hash.NewNode(1, "http://localhost:3001", 1))
	cHashRing.Add(Hash.NewNode(2, "http://localhost:3002", 1))

	countMap := make(map[string]int, 0)
	// which string maps to which server
	belongMap := make(map[structure.KeyValuePair]string)

	var array = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var allPairs []structure.KeyValuePair
	// Make pairs
	for i := 1; i < 11; i++ {
		newPair := structure.KeyValuePair{
			i,
			array[i-1],
		}
		allPairs = append(allPairs, newPair)
	}
	// Assign pairs to different server
	for i := 0; i < 10; i++ {
		k := cHashRing.Get(allPairs[i].Value)
		//fmt.Println(array[i] + " ==> " + k.Ip)
		belongMap[allPairs[i]] = k.Ip
		if _, ok := countMap[k.Ip]; ok {
			countMap[k.Ip] += 1
		} else {
			countMap[k.Ip] = 1
		}
	}

	for k, v := range countMap {
		fmt.Println("Node IP:", k, " count:", v)
	}

	for k, v := range belongMap {
		if strings.Contains(v, "3000") {
			PutKey(k, firstServer)
		} else if strings.Contains(v, "3001") {
			PutKey(k, secondServer)
		} else if strings.Contains(v, "3002") {
			PutKey(k, thirdServer)
		}
	}

}
