FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/claisne/snippetdb

ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET uwlUl3JoVtz4fpDM

# Replace this with actual PostgreSQL DSN.
ENV DATABASE_DRIVER postgres
ENV DSN postgres://postgres:postgres@localhost:5432/snippetdb?sslmode=disable

WORKDIR /go/src/github.com/claisne/snippetdb

# RUN godep go build
RUN go get
RUN go build

EXPOSE 8888
