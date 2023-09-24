FROM golang:1.13.8-alpine3.11
LABEL name="groupie-tracker"
LABEL description=""
LABEL authors="moussadieng"
RUN mkdir /app
RUN apk update && apk add bash && apk add tree 
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]