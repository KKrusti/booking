# booking kata

Booking is a Golang kata. The required tasks can be found at instructions.md file in this same source level

## Installation

To build the application run the following command from the root folder of the project

```bash
go build .
```

If it's desired to execute the tests in the project and check the code coverage run the following command
```bash
go test ./... --cover
```

## Running the project locally
To run the project locally type the following command after building it
```
go run main.go
```
The application will start and will be served on port 3000

If it's desired to test the available endpoints it can be done via Curl or any other tool like Postman
```
curl -i -X POST -H "Content-Type: application/json"
   -d "[
    {
        \"request_id\": \"bookata_XY123\",
        \"check_in\": \"2020-01-01\",
        \"nights\": 5,
        \"selling_rate\": 200,
        \"margin\": 20
    },
    {
        \"request_id\": \"kayete_PP234\",
        \"check_in\": \"2020-01-04\",
        \"nights\": 4,
        \"selling_rate\": 156,
        \"margin\": 22
    }
]"
   http://localhost:3000/v1/stats  
``` 
this call should return 
```
HTTP/1.1 200 OK
Date: Sat, 29 Oct 2022 00:48:31 GMT
Content-Type: application/json
Content-Length: 49

{"avg_night":8.29,"min_night":8,"max_night":8.58}
``` 

## Running project in Docker
To run the project in a docker container there are two possibilities
### Docker
To build the docker image execute the following command
```
docker build -t <image-name> .
```
Once the docker image is build it can be executed using the command
```
docker run -p 3000:3000 -it <image-name>
```
Once the docker container is running the same curl command can be used against the API exposed on port 3000

### Docker-compose
To make it simpler the docker container can be run using docker compose at the project root folder
```
docker-compose up
```
To shut down the docker container use the following command
```
docker-compose down
```
In case the image is needed to rebuild use the commands
```
docker-compose build
docker-compose up
```