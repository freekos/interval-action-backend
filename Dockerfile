FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV PORT=8080

RUN go build -o build ./cmd

RUN rm -rf go.* && \
  rm -rf vendor && \
  go clean -cache && \
  apt-get clean

EXPOSE 8080

CMD ["/app/build"]
