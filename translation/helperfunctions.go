package translation
import (
"strings"
"os"
"io/ioutil"
)

func TrimComment(s string) string {
	commentIndex := strings.Index(s, "//")
	if commentIndex == -1 {
		return s
	} else {
		return s[:strings.Index(s, "//")]
	}
}

func TrimBlanks(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	return s
}

func LoadFileContent(path string) string {
	addFile := os.Getenv("GOPATH") + "/" + path
	content, _ := ioutil.ReadFile(addFile)
	return strings.Replace(string(content), "\r", "", -1)
}
