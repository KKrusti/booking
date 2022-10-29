#FROM golang:1.19
#WORKDIR /app
#COPY go.mod go.sum ./
#RUN go mod download
#EXPOSE 3000
#CMD go run main.go
# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.19-alpine

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./
COPY . ./


# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY *.go ./

# compile application
RUN go build -o /booking

# tells Docker that the container listens on specified network ports at runtime
EXPOSE 3000

# command to be used to execute when the image is used to start a container
CMD [ "/booking" ]

