FROM golang:1.20.5

ENV VOICERSS_API_KEY=key
ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app

EXPOSE 8080

CMD ["app"]
