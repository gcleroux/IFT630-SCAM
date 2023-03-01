FROM golang:1.20-bullseye

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /scam

CMD [ "/scam" ]
