# Build Stage
FROM golang:1.21 as build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production Stage
FROM scratch

COPY --from=build /app/main .

EXPOSE 3000

CMD [ "./main" ]