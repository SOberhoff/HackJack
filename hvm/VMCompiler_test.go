package hvm
import (
	"testing"
	. "gopkg.in/check.v1"
"translation"
	"fmt"
)

func Test(t *testing.T) {TestingT(t)}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestVM(t *C) {
	translation := translation.LoadFileContent("resources/hvm/StackArithmetic/SimpleAdd/SimpleAdd.vm")
	fmt.Print(new(VMCompiler).Translate(translation))
}