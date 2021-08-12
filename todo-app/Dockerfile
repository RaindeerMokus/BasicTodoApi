##
##  Build
##

FROM golang:1.16 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /todo-app

RUN apt update && apt install -y dos2unix
RUN find /todo-app -type f -print0 | xargs -0 dos2unix --


##
##  Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /todo-app /todo-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/todo-app" ]