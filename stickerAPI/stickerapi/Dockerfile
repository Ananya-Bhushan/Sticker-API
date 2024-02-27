 FROM golang:1.18.1-alpine AS build
 RUN mkdir /app
 ADD . /app
 WORKDIR /app
#  COPY . .
 RUN go build -o main .
 FROM alpine:latest
 WORKDIR /app
 EXPOSE 9090
 COPY --from=build /app/main .
 CMD ["/app/main"]
 


