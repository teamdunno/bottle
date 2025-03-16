FROM golang:1.23.5 AS builder
WORKDIR /src
COPY . .
RUN ["go", "install", "."]
CMD ["bottle"]