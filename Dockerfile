FROM frolvlad/alpine-gxx AS caff-builder

RUN apk --update add \
        git \
        make \
        cmake \
        pkgconf \
        gtest-dev

# kaitai lib build

WORKDIR /kaitai-cpp-src

RUN git clone https://github.com/kaitai-io/kaitai_struct_cpp_stl_runtime . && \
    git checkout 0.9

RUN cd build && \
    cmake .. && \
    make

# CAFF parser build

COPY ./libcaff /src

COPY ./fixtures /fixtures

RUN mkdir -p /src/build

RUN cp /kaitai-cpp-src/build/libkaitai_struct_cpp_stl_runtime.so /src/build/

WORKDIR /src

RUN make

FROM golang:alpine AS server-builder

RUN apk --update add \
    gcc \
    musl-dev

COPY ./server /src

WORKDIR /src

RUN go build

FROM alpine:latest

RUN apk add --update libstdc++ libgcc

COPY --from=caff-builder /src/build/caff /usr/bin/caff
COPY --from=caff-builder /src/build/libcaff.so /usr/lib/libcaff.so
COPY --from=caff-builder /src/build/libkaitai_struct_cpp_stl_runtime.so /usr/lib/libkaitai_struct_cpp_stl_runtime.so
COPY --from=server-builder /src/views /views
COPY --from=server-builder /src/server /usr/bin/server

RUN mkdir /data

WORKDIR /

ENTRYPOINT ["/usr/bin/server"]
