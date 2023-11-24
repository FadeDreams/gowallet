FROM golang:1.21.3-alpine3.18
###############################
# DOCKER START STAGE
###############################
WORKDIR /usr/src/goapp/
USER ${USER}
ADD ./go.mod /usr/src/goapp/
ADD . /usr/src/goapp/

###############################
# DOCKER ENVIRONMENT STAGE
###############################
ENV GO111MODULE="on" \
  CGO_ENABLED="0" \
  GO_GC="off"

# DOCKER INSTALL & BUILD STAGE
###############################
RUN go mod download \
  && go mod tidy \
  && go mod verify \
  && go build -o main .

###############################
# DOCKER FINAL STAGE
###############################
EXPOSE 8000
CMD ["./main"]
