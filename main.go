package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	for index := 0; index < 100; index++ {

		fmt.Fprintf(w, "Hello World! %s", time.Now())
	}
}

func main() {
	postMe("http://40.113.149.11/v1/order/")
}

func postMe(url string) {

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

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}
