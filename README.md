# Grumpy Free Jaguars

## CAFF parser
The CAFF parser library (`libcaff.so`) is implemented in C++. The library uses [Kaitai Struct](https://kaitai.io) for parsing the animation files. A sample application (`caff`) is provided to demonstrate the usage of the library.

### Build requirements
To build the library and the example application make sure the following software are installed:
- GNU Make, GCC
- [Googletest](https://github.com/google/googletest) framework (Debian/Ubuntu package: `googletest`, Arch package: `gtest`)
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