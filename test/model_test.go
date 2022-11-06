package test

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var titleLatex = "\\documentclass{article} % article 文档\n\\usepackage[UTF8]{ctex}  % 使用宏包(为了能够显示汉字)\n% 设置页面的环境,a4纸张大小，左右上下边距信息\n\\usepackage[a4paper,left=10mm,right=10mm,top=15mm,bottom=15mm]{geometry}\n\n\\title{南京邮电大学数据结构实验报告}   % 文章标题    \n\n\\makeatletter\n\\newcommand\\dlmu[2][4cm]{\\hskip1pt\\underline{\\hb@xt@ #1{\\hss#2\\hss}}\\hskip3pt}\n\\makeatother\n\\begin{document}\n\\maketitle          % 添加这一句才能够显示标题等信息\n\t\\begin{center}\n\t\t\\zihao{3}\n\t\t\\begin{tabular}{rl}\n\t\t\t\t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{题目}    \\hspace{0.2cm}\t\\dlmu[5.5cm]{"

var classLatex = "} \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{学院}    \\hspace{0.2cm}\t\\dlmu[5.5cm]{软件学院} \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{专业班级}\t\\hspace{0.2cm}\t\t\\dlmu[5.5cm]{"

var numNameLatex = "}      \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{届次}\t\\hspace{0.2cm}\t\\dlmu[5.5cm]{2021届}   \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{学号姓名}\t\\hspace{0.2cm}\t\\dlmu[5.5cm]{"

var teacherLatex = "}   \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{指导教师}\t\\hspace{0.2cm}\t\\dlmu[5.5cm]{"

var contentLatex = "}   \\\\\n\t\t\\end{tabular}\n\t\\end{center}\n\n\n\n\n\n\\section{实验内容}"

var endLatex = "\\end{document}"

func TestConnect(t *testing.T) {

	title := "xxxxx"
	class := "xxxxx"
	numName := "xxxxxx"
	teacher := "xxxxxx"
	content := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	modelArticle := titleLatex + title + classLatex + class + numNameLatex + numName + teacherLatex + teacher + contentLatex + content + endLatex

	filepath := "./file/njupt.tex"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(modelArticle)
	if err != nil {
		t.Fatal(err)
	}
	writer.Flush()
	fmt.Println("success write")
}

func TestCom(t *testing.T) {
	filename := "njupt.tex"
	cmdLatex(filename)

}
func cmdLatex(filename string) {
	getwd, _ := os.Getwd()
	f := getwd + "/file/" + filename
	c2 := exec.Command("xelatex", "-output-directory=file", f)
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c2.Wait()

}
