# Grumpy Free Jaguars

This is a CAFF (CrySys Animated File Format) webstore, made as a homework.

## Overview
The CAFF webstore is a multi-component web service which is deployed in docker containers which helps minimizing the attack surface. The example deployment is described in the `docker-compose.yml` file. The `gocaff` service uses a custom built Alpine Linux based image. The image contains the web application (`server`) which is implemented in Go language. The `libcaff` parser and the `caff` utility is also included in this image. The `front` service is a customized Nginx container which is used to limit the size of the uploaded files and handle connection timeouts. Currently TLS is not enabled because the service is not deployed on a public site, but configuring it in Nginx is a well-documented and fairly easy task.

## Deployment
The service uses Docker compose for the example deployment.

The images can be automatically built with the Docker engine by issuing `docker-compose build`
To deploy the services issue the `docker-compose up` command. The service is listening on the HTTP port `8080` by default.

The code has a native component for parsing CAFF files. To run checks (AFL fuzzer and valgrind) for memory management errors on it, run `docker build -f Dockerfile.testtools .`. Note that AFL can run for a really long time.

## Code security
The CAFF parser has been inspected for memory-management errors, using valgrind and AFL. To repeat these checks, please follow the steps above.

The service had been scanned using Zed Attack Proxy and SQLmap. No true positive failures have been discovered.

The development process followed the [OWASP's Go web application secure coding practices](https://raw.githubusercontent.com/OWASP/Go-SCP/master/dist/go-webapp-scp.pdf) to maintain the security standard. The Go compiler and linter was used as a static analysis tool.

## CAFF parser
The CAFF parser library (`libcaff.so`) is implemented in C++. The library uses [Kaitai Struct](https://kaitai.io) for parsing the animation files. A sample application (`caff`) is provided to demonstrate the usage of the library.

### Build requirements
To build the `libcaff` library and the example application, make sure the following software are installed:
- GNU Make, GCC
- pkg-config
- [Googletest](https://github.com/google/googletest) framework (Debian/Ubuntu package: `libgtest-dev`, Arch package: `gtest`)
    - Note: Google Test seems to be broken for Debian-based distros at the time of writing. It should be fine otherwise.
- Docker with BuildKit support (`18.09` or higher)

To build the library simply issue the `make` command in the `libcaff` directory. The outputed files can be found under the `libcaff/build` directory.

### CAFF application usage
The `caff` application reads a CAFF file from the standard input and outputs the preview on the standard output.
The first line of the output is the width and height of the image, separated with a space character. The image data is outputed in raw format after the newline character.

Example for creating a CAFF preview with ImageMagick:
```sh
$ ./caff < 1.caff > /tmp/1.preview
$ head -n 1 /tmp/1.preview | read WIDTH HEIGHT && tail -n +1 /tmp/1.preview | convert -size ${WIDTH}x${HEIGHT} -depth 8 - 1.jpeg
```

## Work shares
- Sarkadi Csaba: 33,333%
- Schulcz Ferenc: 33,333%
- Telek István: 33,333%