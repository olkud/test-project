FROM golang:1.20

WORKDIR /app

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

COPY go.sum ./

COPY go.mod ./

RUN go mod download

COPY . ./

COPY internal ./

RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -build="go build -o /test-project" -command="/test-project"

#RUN CGO_ENABLED=0 GOOS=linux go build -o /test-project

EXPOSE 8080

#CMD ["/test-project"]


