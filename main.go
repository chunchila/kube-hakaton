package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"text/template"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {

	var h = `
	<!DOCTYPE html>
<html lang="en">
<head>
  <title>Bootstrap 4 Website Example</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.0/umd/popper.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/js/bootstrap.min.js"></script>
  <style>
  .fakeimg {
      height: 200px;
      background: #aaa;
  }
  </style>
</head>
<body>

<table style="width:100%">
  <tr>
    <th>Firstname</th>
    <th>Lastname</th> 
    <th>Age</th>
  </tr>
  
  <tr>
  sd
  </tr>
  <tr>
  ddf
  </tr>
  <tr>
  dffd
  </tr>
  <tr>

  dfd
  </tr>
    <td>Jill</td>
	<td>
	<form action="http://40.113.149.11/v1/order/" method="POST">
    Email :<input type="text" name="EmailAddress" value="guyravid94@gmail.com"><br>
    ID :<input type="text" name="ID" value="1"><br>
    PreferredLanguage :<input type="text" name="PreferredLanguage" value="en"><br>
    Product :<input type="text" name="Product" value="Mickey"><br>
    Partition :<input type="text" name="Partition" value="Mickey"><br>
    Source :<input type="text" name="Source" value="Mickey"><br>
    Status :<input type="text" name="Status" value="Mickey"><br>
    Total :<input type="text" name="Total" value="2242"><br>
	Submit :<input type="submit" value="Submit" class="btn-success">
	
	</form>
	</td> 
    <td>50</td>
  </tr>
  <tr>
    <td>Eve</td>
    <td>Jackson</td> 
    <td>94</td>
  </tr>
</table>

	


</body>
</html>
`

	t, _ := template.New("foo").Parse(h)
	t.Execute(w, "sds")
}

func main() {

	log.Println("this is sparta")
	http.HandleFunc("/", greet)
	log.Fatal(http.ListenAndServe(":5000", nil))

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
