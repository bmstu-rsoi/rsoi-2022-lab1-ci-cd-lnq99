FROM golang:alpine as build-stage

WORKDIR /app

COPY server/go.mod ./

RUN go mod download

COPY server .

RUN go mod tidy

RUN go build -o api ./cmd/main.go


FROM alpine as production-stage

WORKDIR /app

COPY --from=build-stage /app/api .

EXPOSE 8080

CMD ./api
