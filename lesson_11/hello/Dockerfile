FROM golang

ADD . /go/src/github.com/bukhavtsov/go-training-spring-2021/lesson_11/hello

RUN go install github.com/bukhavtsov/go-training-spring-2021/lesson_11/hello@latest

ENTRYPOINT /go/bin/hello

EXPOSE 8080

## How to run this application locally:
# docker build -t lesson_11 .
# docker run --publish 6060:8080 --name lesson_11_app --rm lesson_11
