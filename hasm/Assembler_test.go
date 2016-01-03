package hasm
import (
	. "gopkg.in/check.v1"
	"testing"
	"strings"
	"translation"
)

func Test(t *testing.T) {TestingT(t)}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestAdd(t *C) {
	locs := ConvertToLoCs(translation.LoadFileContent("resources/hasm/add/Add.asm"))
	expectedLines := strings.Split(translation.LoadFileContent("resources/hasm/add/Add.hack"), "\n")
	for _, loc := range locs {
		if loc.targetIndex >= 0 {
			sourcePrefix := loc.String() + " -> "
			t.Assert(sourcePrefix + loc.translation, Equals, sourcePrefix + expectedLines[loc.targetIndex])
		}
	}
}

func (s *MySuite) TestMax(t *C) {
	locs := ConvertToLoCs(translation.LoadFileContent("resources/hasm/max/Max.asm"))
	expectedLines := strings.Split(translation.LoadFileContent("resources/hasm/max/Max.hack"), "\n")
	for _, loc := range locs {
		if loc.targetIndex >= 0 {
			sourcePrefix := loc.String() + " -> "
			t.Assert(sourcePrefix + loc.translation, Equals, sourcePrefix + expectedLines[loc.targetIndex])
		}
	}
}

func (s *MySuite) TestMaxL(t *C) {
	locs := ConvertToLoCs(translation.LoadFileContent("resources/hasm/max/MaxL.asm"))
	expectedLines := strings.Split(translation.LoadFileContent("resources/hasm/max/MaxL.hack"), "\n")
	for _, loc := range locs {
		if loc.targetIndex >= 0 {
			sourcePrefix := loc.String() + " -> "
			t.Assert(sourcePrefix + loc.translation, Equals, sourcePrefix + expectedLines[loc.targetIndex])
		}
	}
}

func (s *MySuite) TestPong(t *C) {
	locs := ConvertToLoCs(translation.LoadFileContent("resources/hasm/pong/Pong.asm"))
	expectedLines := strings.Split(translation.LoadFileContent("resources/hasm/pong/Pong.hack"), "\n")
	for _, loc := range locs {
		if loc.targetIndex >= 0 {
			sourcePrefix := loc.String() + " -> "
			t.Assert(sourcePrefix + loc.translation, Equals, sourcePrefix + expectedLines[loc.targetIndex])
		}
	}
}

func (s *MySuite) TestPongL(t *C) {
	locs := ConvertToLoCs(translation.LoadFileContent("resources/hasm/pong/PongL.asm"))
	expectedLines := strings.Split(translation.LoadFileContent("resources/hasm/pong/PongL.hack"), "\n")
	for _, loc := range locs {
		if loc.targetIndex >= 0 {
			sourcePrefix := loc.String() + " -> "
			t.Assert(sourcePrefix + loc.translation, Equals, sourcePrefix + expectedLines[loc.targetIndex])
		}
	}
}

func (s *MySuite) TestRect(t *C) {
	locs := ConvertToLoCs(translation.LoadFileContent("resources/hasm/rect/Rect.asm"))
	expectedLines := strings.Split(translation.LoadFileContent("resources/hasm/rect/Rect.hack"), "\n")
	for _, loc := range locs {
		if loc.targetIndex >= 0 {
			sourcePrefix := loc.String() + " -> "
			t.Assert(sourcePrefix + loc.translation, Equals, sourcePrefix + expectedLines[loc.targetIndex])
		}
	}
}

func (s *MySuite) TestRectL(t *C) {
	locs := ConvertToLoCs(translation.LoadFileContent("resources/hasm/rect/RectL.asm"))
	expectedLines := strings.Split(translation.LoadFileContent("resources/hasm/rect/RectL.hack"), "\n")
	for _, loc := range locs {
		if loc.targetIndex >= 0 {
			sourcePrefix := loc.String() + " -> "
			t.Assert(sourcePrefix + loc.translation, Equals, sourcePrefix + expectedLines[loc.targetIndex])
		}
	}
}

func (s *MySuite) TestNegativeConstant(t *C) {
	testString := "@-1"
	translation := new(Assembler).Translate(testString)
	t.Assert(translation, Equals, "0111111111111111\n")
}
