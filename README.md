# Grumpy Free Jaguars

This is a CAFF (CrySys Animated File Format) webstore, made as a homework.

## Build and use
Docker-compose is required for the project.

Building and running all takes place in docker-compose, so just issue the following command in the project root: `docker-compose up --build && yes | docker-compose rm`.

The code has a native component for parsing CAFF files. To run checks (AFL fuzzer and valgrind) for memory management errors on it, run `docker-compose --file docker-compose.testtools.yml up --build && yes | docker-compose --file docker-compose.testtools.yml rm`. Note that AFL can run for a really long time.

## Security tests
The CAFF parser has been inspected for memory-management errors, using valgrind and AFL. To repeat these checks, see above.

The service has been scanned using Zed Attack Proxy and SQLmap. No true positive problems have been discovered.

## CAFF parser
The CAFF parser library (`libcaff.so`) is implemented in C++. The library uses [Kaitai Struct](https://kaitai.io) for parsing the animation files. A sample application (`caff`) is provided to demonstrate the usage of the library.

### Build requirements
The parser is built along with the other components in docker-compose.

To only build the library and the example application make sure the following software are installed:
- GNU Make, GCC
- [Googletest](https://github.com/google/googletest) framework (Debian/Ubuntu package: `libgtest-dev`, Arch package: `gtest`)
    - Note: Google Test seems to be broken for Debian-based distros at the time of writing. It should be fine otherwise.
- Docker with BuildKit support (`18.09` or higher)

To build the library simply issue the `make` command in the `libcaff` directory. The outputed files can be found under the `libcaff/build` directory.

### Application usage
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