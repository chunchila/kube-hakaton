package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"text/template"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {

	var h = `<form action="http://40.113.149.11/v1/order/" method="POST">
    <input type="text" name="EmailAddress" value="guyravid94@gmail.com"><br>
    <input type="text" name="ID" value="1"><br>
    <input type="text" name="PreferredLanguage" value="en"><br>
    <input type="text" name="Product" value="Mickey"><br>
    <input type="text" name="Partition" value="Mickey"><br>
    <input type="text" name="Source" value="Mickey"><br>
    <input type="text" name="Status" value="Mickey"><br>
    <input type="text" name="Total" value="2242"><br>
    <input type="submit" value="Submit">
  </form>
`

	t, _ := template.New("foo").Parse(h)
	t.Execute(w, "sds")
}

func main() {

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)

}

func runAsyncPost() {

	var wg sync.WaitGroup
	for index := 0; index < 300; index++ {
		wg.Add(1)

		rand.Seed(time.Now().Unix())
		r := rand.Intn(4000)
		time.Sleep(time.Millisecond * 50)

		fmt.Println("This is the time selected for ", index, r)
		go postMe("http://40.113.149.11/v1/order/", &wg, index, r)
	}

	wg.Wait()
}

func postMe(url string, wg *sync.WaitGroup, index int, random int) {

	defer wg.Done()
	var jsonStr = []byte(`"{"EmailAddress": "guyravid94@gmail.com",  "ID": "string",  "PreferredLanguage": "string",  "Product": "string",  "Partition": "string",  "Source": "string",  "Status": "status",  "Total": 1}"`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	time.Sleep(time.Millisecond * time.Duration(random))
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
