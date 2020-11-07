FROM openjdk:8-buster AS compile-stage

RUN apt-key adv --keyserver hkp://pool.sks-keyservers.net --recv 379CE192D401AB61

RUN apt-get update && \
    apt-get -y install apt-transport-https ca-certificates

RUN echo "deb https://dl.bintray.com/kaitai-io/debian jessie main" | tee /etc/apt/sources.list.d/kaitai.list

RUN apt-get update && \
    apt-get -y install kaitai-struct-compiler

COPY "." "/src"

WORKDIR "/src"

RUN kaitai-struct-compiler --target cpp_stl --outdir out caff.ksy

FROM scratch AS export-stage

COPY --from=compile-stage "/src/out" "/"