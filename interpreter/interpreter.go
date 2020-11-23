package interpreter

import (
	"fmt"
	"os"

	"github.com/hculpan/gopascal/parser"
)

func Interpreter(filename string) {
	fmt.Println("Reading", filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("ERROR: file '%s' does not exist\n", filename)
		return
	} else if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	parser := parser.NewParser(file)
	_, errors := parser.Parse()
	hasErrors := len(errors) > 0
	for _, v := range errors {
		fmt.Println(v)
	}

	if hasErrors {
		return
	}

	fmt.Println("Ready to interpret")
}
