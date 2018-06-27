package main

import (
	"bytes"
	"flag"
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

    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">
  <!-- Bootstrap core CSS -->
    <link href="./vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom fonts for this template -->
    <link href="vendor/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">
    <link href="https://fonts.googleapis.com/css?family=Montserrat:400,700" rel="stylesheet" type="text/css">
    <link href="https://fonts.googleapis.com/css?family=Lato:400,700,400italic,700italic" rel="stylesheet" type="text/css">

    <!-- Plugin CSS -->
    <link href="./vendor/magnific-popup/magnific-popup.css" rel="stylesheet" type="text/css">

    <!-- Custom styles for this template -->
    <link href="./css/freelancer.min.css" rel="stylesheet">
    

</head>

<body id="page-top">

<!-- Navigation -->
<nav class="navbar navbar-expand-lg bg-secondary fixed-top text-uppercase" id="mainNav">
    <div class="container">
        <a class="navbar-brand js-scroll-trigger" href="#page-top">Place an Order</a>
        <button class="navbar-toggler navbar-toggler-right text-uppercase bg-primary text-white rounded" type="button"
                data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive"
                aria-expanded="false" aria-label="Toggle navigation">
            Menu
            <i class="fa fa-bars"></i>
        </button>
        <div class="collapse navbar-collapse" id="navbarResponsive">
            <ul class="navbar-nav ml-auto">
                <li class="nav-item mx-0 mx-lg-1">
                    <a class="nav-link py-3 px-0 px-lg-3 rounded js-scroll-trigger" href="#todaysList">Todays List</a>
                </li>
                <li class="nav-item mx-0 mx-lg-1">
                    <a class="nav-link py-3 px-0 px-lg-3 rounded js-scroll-trigger" href="#ordernow">Order Now</a>
                </li>
            </ul>
        </div>
    </div>
</nav>

<!-- Header -->
<header class="masthead bg-primary text-white text-center">
    <div class="container">
        <!--img class="img-fluid mb-5 d-block mx-auto" src="static/img/profile.png" alt=""-->
        <h1 class="text-uppercase mb-0">Order Your Book</h1>
        <hr class="star-light">
        <h2 class="font-weight-light mb-0">........</h2>
    </div>
</header>


<section id="todaysList">
    <div class="container">
        <h2 class="text-center text-uppercase text-secondary mb-0">Current Orfer</h2>
        <hr class="star-dark mb-5">
        <div class="row">
        </div>

        <form name="orderfrm ge" id="todaysListFrm" name="todaysListFrm" action="http://40.113.149.11/v1/order/"
        method="post">
        
                  <div class="control-group">
                      <div class="form-group floating-label-form-group controls mb-0 pb-2">
                        <label>Email</label>
                        Email
                        <input class="form-control" value="guyravid94@gmail.com" id="EmailAddress" name="EmailAddress" type="email" placeholder="EmailAddress" required="required" data-validation-required-message="Please enter your phone number.">
                        <p class="help-block text-danger"></p>
                      </div>
                      <div class="control-group">
                          <div class="form-group floating-label-form-group controls mb-0 pb-2">
                            <label>ID</label>
                            ID
                            <input class="form-control" value="1" id="ID" name="ID" type="text" placeholder="Id" required="required" data-validation-required-message="Please enter your Id number.">
                            <p class="help-block text-danger"></p>
                          </div>
                      </div>
                      <div class="control-group">
                          <div class="form-group floating-label-form-group controls mb-0 pb-2">
                            <label>PreferredLanguage</label>
                            PreferredLanguage
                            <input class="form-control" value="PreferredLanguage" id="PreferredLanguage" name="PreferredLanguage" type="text" placeholder="PreferredLanguage" required="required" data-validation-required-message="Please enter your Id number.">
                            <p class="help-block text-danger"></p>
                          </div>
                      </div>
                      <div class="control-group">
                          <div class="form-group floating-label-form-group controls mb-0 pb-2">
                            <label>Product</label>
                            Product
                            <input class="form-control" value="Product" id="Product" name="Product" type="text" placeholder="Product" required="required" data-validation-required-message="Please enter your Id number.">
                            <p class="help-block text-danger"></p>
                          </div>
                      </div>
                      <div class="control-group">
                          <div class="form-group floating-label-form-group controls mb-0 pb-2">
                            <label>Partition</label>
                            Partition
                            <input class="form-control" value="Partition" id="Partition" name="Partition" type="text" placeholder="Partition" required="required" data-validation-required-message="Please enter your Id number.">
                            <p class="help-block text-danger"></p>
                          </div>
                      </div>
                      <div class="control-group">
                          <div class="form-group floating-label-form-group controls mb-0 pb-2">
                            <label>Source</label>
                            Source
                            <input class="form-control" value="Source" id="Source" name="Source" type="text" placeholder="Source" required="required" data-validation-required-message="Please enter your Id number.">
                            <p class="help-block text-danger"></p>
                          </div>
                      </div>
                      <div class="control-group">
                          <div class="form-group floating-label-form-group controls mb-0 pb-2">
                            <label>Status</label>
                            Status
                            <input class="form-control" value="Status" id="Status" name="Status" type="text" placeholder="Status" required="required" data-validation-required-message="Please enter your Id number.">
                            <p class="help-block text-danger"></p>
                          </div>
                      </div>

                      <div class="control-group">
                          <div class="form-group floating-label-form-group controls mb-0 pb-2">
                            <label>Total</label>
                            Total
                            <input class="form-control" value="Total" id="Total" name="Total" type="text" placeholder="Total" required="required" data-validation-required-message="Please enter your Id number.">
                            <p class="help-block text-danger"></p>
                          </div>
                      </div>

                      <div class="form-group">
                          <button type="submit" class="btn btn-primary btn-xl" id="sendMessageButton">Send</button>
                        </div>
                  </div>

</section>

        
    </div>
</section>


<!-- Order Section -->



<!--
 Footer
<footer class="footer text-center">
    <div class="container">
        <div class="row">
            <div class="col-md-4 mb-5 mb-lg-0">
                <h4 class="text-uppercase mb-4">Location</h4>
                <p class="lead mb-0">2215 John Daniel Drive
                    <br>Clark, MO 65243</p>
            </div>
            <div class="col-md-4 mb-5 mb-lg-0">
                <h4 class="text-uppercase mb-4">Around the Web</h4>
                <ul class="list-inline mb-0">
                    <li class="list-inline-item">
                        <a class="btn btn-outline-light btn-social text-center rounded-circle" href="#">
                            <i class="fa fa-fw fa-facebook"></i>
                        </a>
                    </li>
                    <li class="list-inline-item">
                        <a class="btn btn-outline-light btn-social text-center rounded-circle" href="#">
                            <i class="fa fa-fw fa-google-plus"></i>
                        </a>
                    </li>
                    <li class="list-inline-item">
                        <a class="btn btn-outline-light btn-social text-center rounded-circle" href="#">
                            <i class="fa fa-fw fa-twitter"></i>
                        </a>
                    </li>
                    <li class="list-inline-item">
                        <a class="btn btn-outline-light btn-social text-center rounded-circle" href="#">
                            <i class="fa fa-fw fa-linkedin"></i>
                        </a>
                    </li>
                    <li class="list-inline-item">
                        <a class="btn btn-outline-light btn-social text-center rounded-circle" href="#">
                            <i class="fa fa-fw fa-dribbble"></i>
                        </a>
                    </li>
                </ul>
            </div>
            <div class="col-md-4">
                <h4 class="text-uppercase mb-4">About Freelancer</h4>
                <p class="lead mb-0">Freelance is a free to use, open source Bootstrap theme created by
                    Order Your Food.</p>
            </div>
        </div>
    </div>
</footer>-->

<div class="copyright py-4 text-center text-white">
    <div class="container">
        <small>Copyright &copy; Your MyWebsite 2018</small>
    </div>
</div>

<!-- Scroll to Top Button (Only visible on small and extra-small screen sizes) -->
<div class="scroll-to-top d-lg-none position-fixed ">
    <a class="js-scroll-trigger d-block text-center text-white rounded" href="#page-top">
        <i class="fa fa-chevron-up"></i>
    </a>
</div>


<!-- Portfolio Modal 1 -->
<div class="portfolio-modal mfp-hide" id="portfolio-modal-1">
    <div class="portfolio-modal-dialog bg-white">
        <a class="close-button d-none d-md-block portfolio-modal-dismiss" href="#">
            <i class="fa fa-3x fa-times"></i>
        </a>
        <div class="container text-center">
            <div class="row">
                <div class="col-lg-8 mx-auto">
                    <h2 class="text-secondary text-uppercase mb-0">Project Name</h2>
                    <hr class="star-dark mb-5">
                    <img class="img-fluid mb-5" src="static/img/portfolio/cabin.png" alt="">
                    <p class="mb-5">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia neque assumenda
                        ipsam nihil, molestias magnam, recusandae quos quis inventore quisquam velit asperiores, vitae?
                        Reprehenderit soluta, eos quod consequuntur itaque. Nam.</p>
                    <a class="btn btn-primary btn-lg rounded-pill portfolio-modal-dismiss" href="#">
                        <i class="fa fa-close"></i>
                        Close Project</a>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Portfolio Modal 2 -->
<div class="portfolio-modal mfp-hide" id="portfolio-modal-2">
    <div class="portfolio-modal-dialog bg-white">
        <a class="close-button d-none d-md-block portfolio-modal-dismiss" href="#">
            <i class="fa fa-3x fa-times"></i>
        </a>
        <div class="container text-center">
            <div class="row">
                <div class="col-lg-8 mx-auto">
                    <h2 class="text-secondary text-uppercase mb-0">Project Name</h2>
                    <hr class="star-dark mb-5">
                    <img class="img-fluid mb-5" src="static/img/portfolio/cake.png" alt="">
                    <p class="mb-5">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia neque assumenda
                        ipsam nihil, molestias magnam, recusandae quos quis inventore quisquam velit asperiores, vitae?
                        Reprehenderit soluta, eos quod consequuntur itaque. Nam.</p>
                    <a class="btn btn-primary btn-lg rounded-pill portfolio-modal-dismiss" href="#">
                        <i class="fa fa-close"></i>
                        Close Project</a>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Portfolio Modal 3 -->
<div class="portfolio-modal mfp-hide" id="portfolio-modal-3">
    <div class="portfolio-modal-dialog bg-white">
        <a class="close-button d-none d-md-block portfolio-modal-dismiss" href="#">
            <i class="fa fa-3x fa-times"></i>
        </a>
        <div class="container text-center">
            <div class="row">
                <div class="col-lg-8 mx-auto">
                    <h2 class="text-secondary text-uppercase mb-0">Project Name</h2>
                    <hr class="star-dark mb-5">
                    <img class="img-fluid mb-5" src="static/img/portfolio/circus.png" alt="">
                    <p class="mb-5">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia neque assumenda
                        ipsam nihil, molestias magnam, recusandae quos quis inventore quisquam velit asperiores, vitae?
                        Reprehenderit soluta, eos quod consequuntur itaque. Nam.</p>
                    <a class="btn btn-primary btn-lg rounded-pill portfolio-modal-dismiss" href="#">
                        <i class="fa fa-close"></i>
                        Close Project</a>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Portfolio Modal 4 -->
<div class="portfolio-modal mfp-hide" id="portfolio-modal-4">
    <div class="portfolio-modal-dialog bg-white">
        <a class="close-button d-none d-md-block portfolio-modal-dismiss" href="#">
            <i class="fa fa-3x fa-times"></i>
        </a>
        <div class="container text-center">
            <div class="row">
                <div class="col-lg-8 mx-auto">
                    <h2 class="text-secondary text-uppercase mb-0">Project Name</h2>
                    <hr class="star-dark mb-5">
                    <img class="img-fluid mb-5" src="static/img/portfolio/game.png" alt="">
                    <p class="mb-5">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia neque assumenda
                        ipsam nihil, molestias magnam, recusandae quos quis inventore quisquam velit asperiores, vitae?
                        Reprehenderit soluta, eos quod consequuntur itaque. Nam.</p>
                    <a class="btn btn-primary btn-lg rounded-pill portfolio-modal-dismiss" href="#">
                        <i class="fa fa-close"></i>
                        Close Project</a>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Portfolio Modal 5 -->
<div class="portfolio-modal mfp-hide" id="portfolio-modal-5">
    <div class="portfolio-modal-dialog bg-white">
        <a class="close-button d-none d-md-block portfolio-modal-dismiss" href="#">
            <i class="fa fa-3x fa-times"></i>
        </a>
        <div class="container text-center">
            <div class="row">
                <div class="col-lg-8 mx-auto">
                    <h2 class="text-secondary text-uppercase mb-0">Project Name</h2>
                    <hr class="star-dark mb-5">
                    <img class="img-fluid mb-5" src="static/img/portfolio/safe.png" alt="">
                    <p class="mb-5">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia neque assumenda
                        ipsam nihil, molestias magnam, recusandae quos quis inventore quisquam velit asperiores, vitae?
                        Reprehenderit soluta, eos quod consequuntur itaque. Nam.</p>
                    <a class="btn btn-primary btn-lg rounded-pill portfolio-modal-dismiss" href="#">
                        <i class="fa fa-close"></i>
                        Close Project</a>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Portfolio Modal 6 -->
<div class="portfolio-modal mfp-hide" id="portfolio-modal-6">
    <div class="portfolio-modal-dialog bg-white">
        <a class="close-button d-none d-md-block portfolio-modal-dismiss" href="#">
            <i class="fa fa-3x fa-times"></i>
        </a>
        <div class="container text-center">
            <div class="row">
                <div class="col-lg-8 mx-auto">
                    <h2 class="text-secondary text-uppercase mb-0">Project Name</h2>
                    <hr class="star-dark mb-5">
                    <img class="img-fluid mb-5" src="static/img/portfolio/submarine.png" alt="">
                    <p class="mb-5">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Mollitia neque assumenda
                        ipsam nihil, molestias magnam, recusandae quos quis inventore quisquam velit asperiores, vitae?
                        Reprehenderit soluta, eos quod consequuntur itaque. Nam.</p>
                    <a class="btn btn-primary btn-lg rounded-pill portfolio-modal-dismiss" href="#">
                        <i class="fa fa-close"></i>
                        Close Project</a>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Bootstrap core JavaScript -->
<script src="static/vendor/jquery/jquery.min.js"></script>
<script src="static/vendor/bootstrap/js/bootstrap.bundle.min.js"></script>

<!-- Plugin JavaScript -->
<script src="static/vendor/jquery-easing/jquery.easing.min.js"></script>
<script src="static/vendor/magnific-popup/jquery.magnific-popup.min.js"></script>

<!-- Contact Form JavaScript -->
<script src="static/js/jqBootstrapValidation.js"></script>
<script src="static/js/contact_me.js"></script>

<!-- Custom scripts for this template -->
<script src="static/js/freelancer.min.js"></script>

</body>

</html>
`

	t, _ := template.New("foo").Parse(h)
	t.Execute(w, "sds")
}

func main() {
	directory := flag.String("d", ".", "the directory of static file to host")
	http.Handle("/", http.FileServer(http.Dir(*directory)))
	// http.HandleFunc("/", greet)
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
