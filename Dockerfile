FROM golang:1.20-alpine3.18
LABEL name="groupie-tracker"
LABEL description=""
LABEL authors="moussadieng"
RUN mkdir /app
RUN apk update && apk add bash && apk add tree 
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]