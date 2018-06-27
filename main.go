package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	for index := 0; index < 100; index++ {

		fmt.Fprintf(w, "Hello World! %s", time.Now())
	}
}

func main() {

	var wg sync.WaitGroup
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go postMe("http://40.113.149.11/v1/order/", &wg, index)
	}

	wg.Wait()
}

func postMe(url string, wg *sync.WaitGroup, index int) {

	defer wg.Done()
	var jsonStr = []byte(`"{"EmailAddress": "roman",  "ID": "string",  "PreferredLanguage": "string",  "Product": "string",  "Partition": "string",  "Source": "string",  "Status": "strinfgvfdvdfsvfg",  "Total": 1}"`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status, index)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}
