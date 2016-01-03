package hasm
import (
	"strings"
	"strconv"
	"translation"
)

var instructionCount int

type LoC struct {
	input          string
	isLabel        bool
	isAInstruction bool
	isCInstruction bool
	isRedundant    bool
	sourceIndex    int
	targetIndex    int
	translation    string
}

func newLoC(input string, sourceIndex int) *LoC {
	input = translation.TrimBlanks(translation.TrimComment(input))
	loc := new(LoC)
	loc.input = input
	loc.isLabel = strings.HasPrefix(input, "(") && strings.HasSuffix(input, ")")
	loc.isAInstruction = strings.HasPrefix(input, "@")
	loc.isCInstruction = !loc.isLabel && !loc.isAInstruction && input != ""
	loc.isRedundant = !loc.isLabel && !loc.isAInstruction && !loc.isCInstruction
	loc.sourceIndex = sourceIndex
	if loc.isAInstruction || loc.isCInstruction {
		loc.targetIndex = instructionCount
		instructionCount++
	} else {
		loc.targetIndex = -1
	}
	return loc
}

func (loc *LoC) String() string {
	return strconv.FormatInt(int64(loc.sourceIndex), 10) + ": " + loc.input
}

func splitIntoLoCs(input string) []LoC {
	instructionCount = 0
	locs := make([]LoC, 0)
	for i, line := range strings.Split(input, "\n") {
		locs = append(locs, *newLoC(line, i))
	}
	return locs
}

func (loc *LoC) getLabel() string {
	if !loc.isLabel {panic(loc.input + " isn't a label")}
	return loc.input[1:len(loc.input) - 1]
}

func (loc *LoC) getImmediate() (immediate string, isVariable bool) {
	if !loc.isAInstruction {panic(loc.input + " isn't an A-instruction")}
	_, err := strconv.ParseInt(loc.input[1:], 10, 64)
	return loc.input[1:], err != nil
}

func (loc *LoC) findLabelTarget(locs []LoC) int {
	if !loc.isLabel {panic(loc.input + " isn't a label")}
	for _, followingLoc := range locs[loc.sourceIndex + 1 :] {
		if followingLoc.isAInstruction || followingLoc.isCInstruction {
			return followingLoc.targetIndex
		}
	}
	panic("no instruction following label " + loc.input + " found")
}

func (loc *LoC) splitCInstruction() (dest string, comp string, jmp string) {
	unprocessed := loc.input
	equalSignIndex := strings.Index(unprocessed, "=")
	semicolonIndex := strings.Index(unprocessed, ";")
	if equalSignIndex >= 0 {
		dest = unprocessed[:equalSignIndex]
		unprocessed = unprocessed[equalSignIndex + 1:]
	}
	if semicolonIndex >= 0 {
		jmp = unprocessed[semicolonIndex + 1:]
		unprocessed = unprocessed[:semicolonIndex]
	}
	comp = unprocessed
	return
}