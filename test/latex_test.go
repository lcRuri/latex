package test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var latexHead = "\\documentclass{article}\n\\usepackage{xeCJK}\n\\usepackage{hyperref}\n"

var latexStart = "\\begin{document}\n"

var latexEnd = "\\end{document}\n"

var latexUrl = "\\url{"

func TestLatex(t *testing.T) {
	msg := "hello latex\n"
	url := "www.baidu.com"
	content := latexHead + latexStart + msg + latexUrl + url + "}" + "\n" + latexEnd

	filepath := "./file/la01.tex"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	writer.Flush()
	fmt.Println("success write")

}


