package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtSecret = "latex"

var TokenExpireTime = 3600 * 7

var TitleLatex = "\\documentclass{article} % article 文档\n\\usepackage[UTF8]{ctex}  % 使用宏包(为了能够显示汉字)\n% 设置页面的环境,a4纸张大小，左右上下边距信息\n\\usepackage[a4paper,left=10mm,right=10mm,top=15mm,bottom=15mm]{geometry}\n\n\\title{南京邮电大学数据结构实验报告}   % 文章标题    \n\n\\makeatletter\n\\newcommand\\dlmu[2][4cm]{\\hskip1pt\\underline{\\hb@xt@ #1{\\hss#2\\hss}}\\hskip3pt}\n\\makeatother\n\\begin{document}\n\\maketitle          % 添加这一句才能够显示标题等信息\n\t\\begin{center}\n\t\t\\zihao{3}\n\t\t\\begin{tabular}{rl}\n\t\t\t\t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{题目}    \\hspace{0.2cm}\t\\dlmu[5.5cm]{"

var ClassLatex = "} \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{学院}    \\hspace{0.2cm}\t\\dlmu[5.5cm]{软件学院} \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{专业班级}\t\\hspace{0.2cm}\t\t\\dlmu[5.5cm]{"

var NumNameLatex = "}      \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{届次}\t\\hspace{0.2cm}\t\\dlmu[5.5cm]{2021届}   \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{学号姓名}\t\\hspace{0.2cm}\t\\dlmu[5.5cm]{"

var TeacherLatex = "}   \\\\\n\t\t\t   \t\t\t\t\t\t\t\t\t &\\makebox[4em][s]{指导教师}\t\\hspace{0.2cm}\t\\dlmu[5.5cm]{"

var ContentLatex = "}   \\\\\n\t\t\\end{tabular}\n\t\\end{center}\n\n\n\n\n\n\\section{实验内容}"

var EndLatex = "\\end{document}"
