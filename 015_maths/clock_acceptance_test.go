package clockFace_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	clockFace "github.com/soudengwu/GoPractice/015_maths"
)

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	b := bytes.Buffer{}
	clockFace.SVGWriter(&b, tm)

	svg := clockFace.SVG{}
	xml.Unmarshal(b.Bytes(), &svg)


	want := clockFace.Line{150, 150, 150, 60}

	for _, line := range svg.Line {
		if line == want {
			return
		}
	}

	t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}
