FROM golang:1.16-alpine AS development

WORKDIR /api

RUN go get -u github.com/cosmtrek/air
RUN export PATH=$PATH:$(go env GOPATH)/bin

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 5000

CMD [ "air" ]
