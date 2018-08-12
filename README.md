# SUPERMARKET API CHALLENGE

#### About

	Restful go application desgined to manage a supermarket

   # Commands
##### This is restful api, you can send the following http requests to it
	-GET 	/
		gets a string for the name of the project and the running version of the api
	-GET 	/fetch
    	gets the current list of the produce held within the database
    -GET    /get/{code}
        gets the code being requested
    -POST 	/add
    	atempts to add a produce item from the database
        - example request method
        {
			"producecode": "A12T-4GH7-QPL9-3N4N",
			"name": "romaine lettuce",
			"unitprice": "4.00"
		}

    -DELETE	/remove/{code}
    	- Example url
    		/remove/A12T-4GH7-QPL9-3N4M


#### Preqisites

	This project is based on the go programing language
    follow instalation instructions from https://golang.org/dl/ for your platform

    If you wish to deploy the application you will also need docker

# Running the application locally

    You can run the application locally directly via go or via docker
    Application Runs on port 8081

    Running Localy:

    1) Clone this repo
    2) Set the GOPATH to the root of this dir
    3) cd into this dir such that main is in scope
    4) Get dependencies for this application (go get ./...)
    5) go run main.go

    Running Via Docker:

    1) run: docker pull docker pull raghusauce011/supermarketchallange
    2) run:	docker run -d -p 8081:8081 raghusauce011/supermarketchallange

 # Testing
  #### Running Unit Tests
  	To run unit tests, run "go test ./.." inside the project directory
  #### Running Integration Tests
    To run integration tests an instance of the application must be running

    run: go test ./supermarket_service -integration
# Deployment
#### Jenkins
	A Jenkinsfile for a pipeline exists within the root of the project.
    It can build and deploy the application to the kuberneties cluster.
    You will need to be on a linux env for this pipeline to work.
#### Manual Deployment
	Build
    	run: go build
        run: SupermarketAPI.exe/Supermarket.sh (depending on env)
    Dockerize
    	run: docker build -t "yourUser/yourRepo:yourTag" .
    Deploy
    	run: kubectl set image deployment/"your deployment"="yourUser/yourRepo:yourTag"
