FROM ubuntu
FROM golang:latest
EXPOSE 8000:8000
COPY Watcher.go /works/Watcher.go
COPY Settings.json /works/Settings.json
COPY ABI /works/ABI
COPY hub.go /works/hub.go
WORKDIR /works
RUN go get github.com/ethereum/go-ethereum
RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest
RUN abigen --pkg main --abi ABI --type Hub --out hub.go
RUN go get -u github.com/ethereum/go-ethereum/accounts/abi/bind
RUN go get -u github.com/ethereum/go-ethereum/common
RUN go get -u github.com/ethereum/go-ethereum/crypto
RUN go get -u github.com/ethereum/go-ethereum/ethclient
RUN go mod init main
RUN go get github.com/ethereum/go-ethereum/accounts/abi/bind
CMD go run main
