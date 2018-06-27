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

	var h = `
	<!DOCTYPE html>
	<html>
	<head>
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB" crossorigin="anonymous">
	<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js" integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49" crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js" integrity="sha384-smHYKdLADwkXOn1EmN1qk/HfnUcbVRZyYmZ4qpPea6sjB/pTJ0euyQp0Mk8ck+5T" crossorigin="anonymous"></script>
</head>
	<body>
	
	<form>
  <div class="form-group row">
    <label for="staticEmail" class="col-sm-2 col-form-label">Email</label>
    <div class="col-sm-10">
      <input type="text" readonly class="form-control-plaintext" id="staticEmail" value="email@example.com">
    </div>
  </div>
  <div class="form-group row">
    <label for="inputPassword" class="col-sm-2 col-form-label">Password</label>
    <div class="col-sm-10">
      <input type="password" class="form-control" id="inputPassword" placeholder="Password">
    </div>
  </div>
</form>

	<form action="http://40.113.149.11/v1/order/" method="POST">
    Email :<input type="text" name="EmailAddress" value="guyravid94@gmail.com"><br>
    ID :<input type="text" name="ID" value="1"><br>
    PreferredLanguage :<input type="text" name="PreferredLanguage" value="en"><br>
    Product :<input type="text" name="Product" value="Mickey"><br>
    Partition :<input type="text" name="Partition" value="Mickey"><br>
    Source :<input type="text" name="Source" value="Mickey"><br>
    Status :<input type="text" name="Status" value="Mickey"><br>
    Total :<input type="text" name="Total" value="2242"><br>
    Submit :<input type="submit" value="Submit">
  </form>
  </body>
  </html>
`

	t, _ := template.New("foo").Parse(h)
	t.Execute(w, "sds")
}

func main() {

	http.HandleFunc("/", greet)
	http.ListenAndServe(":5000", nil)

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
