FROM golang:1.17 as builder

COPY . /dice
WORKDIR /dice/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzCore.a && \
    clang -fsanitize=fuzzer fuzzCore.a -o fuzzCore.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /dice/mayhem/fuzzCore.libfuzzer /