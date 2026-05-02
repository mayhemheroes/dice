FROM golang:1.25 AS builder

COPY . /dice
WORKDIR /dice

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN GOFLAGS=-mod=mod go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN GOFLAGS=-mod=mod go get github.com/AdaLogics/go-fuzz-headers
RUN apt-get update && apt-get install -y clang

RUN cd mayhem && go-fuzz-build -libfuzzer -o fuzzCore.a github.com/dicedb/dice/mayhem && \
    clang -fsanitize=fuzzer fuzzCore.a -o fuzzCore.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /dice/mayhem/fuzzCore.libfuzzer /
