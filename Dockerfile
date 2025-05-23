FROM golang:1.24.2-alpine3.21

RUN apk add --no-cache \
    bash \
    brotli \
    git

RUN wget https://github.com/WebAssembly/binaryen/releases/download/version_123/binaryen-version_123-x86_64-linux.tar.gz && \
    tar -xvzf binaryen-version_123-x86_64-linux.tar.gz && \
    cp binaryen-version_123/bin/wasm-opt /usr/local/bin && \
    rm -r binaryen-version_123 binaryen-version_123-x86_64-linux.tar.gz

RUN wget https://github.com/sass/dart-sass/releases/download/1.77.8/dart-sass-1.77.8-linux-x64-musl.tar.gz && \
    tar -xvzf dart-sass-1.77.8-linux-x64-musl.tar.gz && \
    mv dart-sass / && \
    rm dart-sass-1.77.8-linux-x64-musl.tar.gz

ENV PATH="${PATH}:/dart-sass"
