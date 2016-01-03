package translation
import (
	"io/ioutil"
	"strings"
	"path/filepath"
	"os"
	"fmt"
	"bufio"
)

type Translator interface {
	InputExtension() string
	OutputExtension() string
	fmt.Stringer
	Translate(input map[string]string) (output map[string]string)
}


func RunTranslator(translator Translator) {
	if len(os.Args) < 2 {
		translateStdin(translator)
	} else {
		translateFiles(translator, os.Args[1:])
	}
}

func translateStdin(translator Translator) {
	input, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {panic(err)}
	inputString := strings.Replace(string(input), "\r\n", "\n", -1)
	output := translator.Translate(map[string]string{"Stdin":inputString})
	for _, outputString := range output {
		os.Stdout.WriteString(outputString)
	}
}

func translateFiles(translator Translator, filenames []string) {
	input := readFiles(translator, filenames)
	output := translator.Translate(input)
	writeFiles(translator, output)
}

func readFiles(translator Translator, filenames []string) map[string]string {
	input := make(map[string]string, 0)
	for _, filename := range filenames {
		if !strings.HasSuffix(filename, "." + translator.InputExtension()) {
			panic("input file " + filename + " has an unexpected file extension")
		}
		fileContent, err := ioutil.ReadFile(filename)
		if err != nil {panic(err)}
		baseFilename := strings.TrimSuffix(filename, filepath.Ext(filename))
		input[baseFilename] = strings.Replace(string(fileContent), "\r\n", "\n", -1)
	}
	return input
}

func writeFiles(translator Translator, output map[string]string) {
	for outputBaseName, outputString := range output {
		outputFilename := outputBaseName + "." + translator.OutputExtension()
		ioutil.WriteFile(outputFilename, []byte(outputString), 0644)
	}
}