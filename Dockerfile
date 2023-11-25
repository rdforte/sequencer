FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod verify

COPY . .

RUN go build -v -o /server ./cmd/app/main.go

FROM scratch

COPY --from=build /server .

CMD ["/server"]



