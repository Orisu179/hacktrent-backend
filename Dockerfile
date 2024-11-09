FROM golang:1.22-alpine
LABEL authors="Tyler Li"
RUN apk --update add git
RUN mkdir /app
WORKDIR /app
COPY . /app
CMD go mod download
CMD go run ./cmd/web -addr=":4000"
EXPOSE 4000