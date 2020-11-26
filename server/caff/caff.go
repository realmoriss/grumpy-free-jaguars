package caff

import (
	"context"
	"errors"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type CaffParser struct {
	fullPath, dirname string
}

var (
	ErrParserPath = errors.New("fullPath must be an absolute path to parser binary")
)

func NewParser(fullPath string) (parser *CaffParser, err error) {
	parser = &CaffParser{}

	parser.fullPath = strings.TrimSpace(fullPath)
	parser.dirname = strings.TrimSpace(path.Dir(fullPath))

	if len(parser.dirname) < 1 || !path.IsAbs(fullPath) {
		err = ErrParserPath
	}

	return
}

var (
	ErrParserStart       = errors.New("could not launch parser")
	ErrParserCrash       = errors.New("parser crashed")
	ErrMissingField      = errors.New("missing field from output")
	ErrTruncatedPixels   = errors.New("read fewer pixel bytes than expected")
	ErrDimensionTooLarge = errors.New("dimension out of allowed bounds")
)

func (parser CaffParser) ParseCAFF(ctx context.Context, input io.Reader) (*image.RGBA, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, parser.fullPath)
	cmd.Dir = parser.dirname

	// TODO: optionally log stderr somewhere in case of non-zero status
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	cmd.Stdin = input

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	defer cmd.Process.Kill()

	img, errParse := ParseOutput(ctx, stdout)

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	if errParse != nil {
		return nil, err
	}

	return img, nil
}

var (
	ErrZeroDimension = errors.New("one of the image dimensions is zero")
	ErrTooManyBytes  = errors.New("encountered more decoded bytes than expected")
)

func ParseOutput(ctx context.Context, caffOutput io.Reader) (*image.RGBA, error) {
	width, err := readInt(caffOutput)
	if err != nil {
		return nil, err
	}

	height, err := readInt(caffOutput)
	if err != nil {
		return nil, err
	}

	if height < 1 || width < 1 {
		return nil, ErrZeroDimension
	}

	pixBytes, err := ioutil.ReadAll(caffOutput)
	if err != nil {
		return nil, err
	}

	nPixels := width * height
	expectedBytes := 3 * nPixels // RGB

	if len(pixBytes) < expectedBytes {
		return nil, ErrTruncatedPixels
	}
	if len(pixBytes) > expectedBytes {
		return nil, ErrTooManyBytes
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	srcBuff := []uint8(pixBytes)
	for i := 0; i < nPixels; i++ {
		pixel := srcBuff[i*3 : i*3+3]
		x := i % width
		y := i / width
		pixelColor := color.RGBA{pixel[0], pixel[1], pixel[2], 0xFF}
		img.SetRGBA(x, y, pixelColor)
	}

	return img, nil
}

func readInt(from io.Reader) (int, error) {
	acc := []rune{}
readDigits:
	for {
		var buff [1]byte
		n, err := from.Read(buff[:])

		if n != len(buff) || err != nil {
			return 0, ErrMissingField
		}

		asRune := rune(buff[0])
		switch {
		case asRune >= rune('0') && asRune <= rune('9'):
			acc = append(acc, asRune)
		case unicode.IsSpace(asRune):
			break readDigits
		default:
			return 0, ErrMissingField
		}
	}

	asStr := string(acc)
	asInteger, err := strconv.ParseInt(asStr /*base*/, 10, 32 /*bits*/)
	if err != nil {
		return 0, err
	}

	return int(asInteger), nil
}
