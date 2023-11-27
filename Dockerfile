FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod verify && go mod download

COPY . .

RUN go build -v -o /server ./cmd/app/main.go

FROM scratch

COPY --from=build /server .

ENV API_PORT=3000
ENV DEBUG_PORT=3001
ENV SHUTDOWN_TIMEOUT_SEC=20
ENV ENVIRONMENT=development

CMD ["/server"]
