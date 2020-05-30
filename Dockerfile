FROM golang:latest as builder

LABEL maintainer="Roger Coll <roger.coll.aumatell@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/ipGeoMongo/main.go

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"] 