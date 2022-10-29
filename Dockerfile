FROM golang:1.19-alpine

WORKDIR /app

COPY . ./

RUN go mod download

COPY *.go ./

RUN go build -o /booking

EXPOSE 3000

CMD [ "/booking" ]

