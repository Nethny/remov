FROM ubuntu
FROM golang:latest
EXPOSE 8001:8001
WORKDIR /works
RUN git clone https://github.com/rmrk-team/rmrk2-examples
COPY Receiver.go /works/rmrk2-examples/projects/scripts/Receiver.go
COPY Settings.json /works/rmrk2-examples/projects/scripts/Settings.json
COPY ABI /works/rmrk2-examples/projects/scripts/ABI
COPY hub.go /works/rmrk2-examples/projects/scripts/hub.go
COPY account.env /works/rmrk2-examples/projects/scripts/.env
COPY constants.ts /works/rmrk2-examples/projects/scripts/constants.ts
COPY simple-send.ts /works/rmrk2-examples/projects/scripts/simple-send.ts
COPY run-simple-script.ts /works/rmrk2-examples/projects/scripts/run-simple-script.js
WORKDIR /works/rmrk2-examples/projects/scripts
RUN apt update
RUN apt install -y nodejs
RUN apt install -y npm
RUN npm install typescript
RUN go get github.com/ethereum/go-ethereum
RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest
RUN abigen --pkg main --abi ABI --type Hub --out hub.go
RUN go get -u github.com/ethereum/go-ethereum/accounts/abi/bind
RUN go get -u github.com/ethereum/go-ethereum/common
RUN go get -u github.com/ethereum/go-ethereum/crypto
RUN go get -u github.com/ethereum/go-ethereum/ethclient
RUN go mod init main
RUN go get github.com/ethereum/go-ethereum/accounts/abi/bind

RUN npm install polkadot
CMD go run main
