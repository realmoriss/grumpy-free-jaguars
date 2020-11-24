package caff

import (
    "context"
    "image"
    "image/png"
    "os"
    "path"
    "testing"
)

const fixturesBasePath = "../fixtures"
const parserPath = "/home/sarkadicsa/Documents/BME/MSc/Courses/Computer Security/2020/hw/grumpy-free-jaguars/libcaff/build/caff"

// for debugging
func dumpPNG(img image.Image, path string) {
    outf, _ := os.Create(path)
    defer outf.Close()

    png.Encode(outf, img)
}

func TestParsesProto(t *testing.T) {
    type TestCase struct{
        fixture string
        expectedWith, expectedHeight int
    }

    cases := []TestCase{
        {"1.caff-out", 1000, 667},
    }

    for _, testcase := range cases {
        t.Run(testcase.fixture, func(t *testing.T) {
            fixtureFile := openFixture(testcase.fixture)
            defer fixtureFile.Close()

            img, err := ParseOutput(context.Background(), fixtureFile)
            if err != nil {
                t.Error(err)
                return
            }

            bounds := img.Bounds()
            switch {
            case bounds.Max.X != testcase.expectedWith: fallthrough
            case bounds.Max.Y != testcase.expectedHeight:
                t.Errorf("unexpected bounds")
            }
        })
    }
}

func openFixture(name string) (*os.File) {
    filePath := path.Join(fixturesBasePath, name)
    f, err := os.Open(filePath)
    if err != nil { panic(err) }

    return f
}
