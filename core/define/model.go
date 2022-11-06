package define

var T1 = "\\documentclass[12pt, a4paper, oneside]{ctexart}\n\n\\usepackage{amsmath, amsthm, amssymb, appendix, bm, graphicx, hyperref, mathrsfs, multirow}\n\\usepackage{fancyhdr}\n\\usepackage{zhnumber} % 汉字序号\n\\renewcommand{\\thesection}{\\zhnum{section}、}\n\\renewcommand{\\thesubsection}{\\arabic{section}.\\arabic{subsection}}\n\n\\linespread{1.5}\n\\newtheorem{theorem}{定理}[section]\n\\newtheorem{definition}[theorem]{定义}\n\\newtheorem{lemma}[theorem]{引理}\n\\newtheorem{corollary}[theorem]{推论}\n\\newtheorem{example}[theorem]{例}\n\\newtheorem{proposition}[theorem]{命题}\n\\renewcommand{\\abstractname}{\\Large\\textbf{摘要}}\n\n% 取消居中\n\\CTEXsetup[format={\\Large\\bfseries}]{section}\n\\pagestyle{fancy}\n\\fancyhead{}\n\\cfoot{\\thepage}\n\\renewcommand{\\headrulewidth}{0pt}\n\n\\begin{document}\n\n\\thispagestyle{empty}\n\n\\begin{figure}[t]\n    \\centering\n    \n\\end{figure}\n\n\\begin{center}\n    {\\heiti \\zihao{-0} \\textbf{程序设计报告} }\n\\end{center}\n\\begin{center}\n    \\normalsize \\zihao{-3} \\textbf{（"

var T2 = "学年/\\;第"

var T3 = "学期）}\n\\end{center}\n\n\\vspace{4em}\n\n\\begin{center}\n    {\\heiti \\zihao{2} \\textbf{题 \\quad 目：} \\zihao{3} \\underline{\\textbf{"

var T4 = "}}}\n\\end{center}\n\n\\newlength{\\changdu}\n\\settowidth{\\changdu}{\\large{组员 \\quad 学号姓名}}\n\n\\begin{table}[b]\n    \\centering\n    \\large\n    \\begin{tabular}{ll}\n        {\\heiti \\textbf{专\\hspace{\\fill}业}} &\\underline{"

var T5 = "} \\\\\n        {\\heiti \\textbf{组长\\hspace{\\fill}学号姓名}}&\\underline{"

var T6 = "} \\\\\n        {\\heiti \\textbf{组员\\hspace{\\fill}学号姓名}}&\\underline{"

var T7 = "} \\\\\n        {\\heiti \\textbf{指\\hspace{\\fill}导\\hspace{\\fill}教\\hspace{\\fill}师}}&\\underline{"

var T8 = "} \\\\\n        {\\heiti \\textbf{指\\hspace{\\fill}导\\hspace{\\fill}单\\hspace{\\fill}位}}&\\underline{"

var T9 = "} \\\\\n        {\\heiti \\textbf{日\\hspace{\\fill}期}} &\\underline{"

var T10 = "} \n    \\end{tabular}\n\\end{table}\n\n\\newpage\n\n% 表格部分\n\\begin{tabular}{|c|r|r|r|r|r|r|}\n\t\\hline\n\t\\multirow{3}{*}{成员分工} & 组长 & \\multicolumn{5}{|c|}{"

var T11 = "}{\\; \\hspace{20em} \\; } \\\\ \\cline{2-7} \n        &  组员 & \\multicolumn{5}{|c|}{"

var T12 = "}{\\; \\hspace{20em} \\; } \\\\ \\cline{2-7} \n        &  组员& \\multicolumn{5}{|c|}{"

var T13 = "}{\\; \\hspace{20em} \\; } \\\\ \\hline\n    \n    \\multirow{2}{*}{评分细则} & 组长 & \\multicolumn{5}{|c|}{xx} \\\\ \\cline{2-7}\n        &  组员 & \\multicolumn{5}{|c|}{\\; \\hspace{20em} \\; } \\\\ \\cline{2-7} \n        &  组员& \\multicolumn{5}{|c|}{\\; \\hspace{20em} \\; } \\\\ \\hline\n    \n    \\multirow{2}{*}{简短评语} & \\multicolumn{6}{|c|}{xx} \\\\ \n        &   \\multicolumn{6}{|c|}{\\; \\hspace{20em} \\; } \\\\\n        &   \\multicolumn{6}{|c|}{\\; \\hspace{20em} \\; } \\\\ \n        &   \\multicolumn{6}{|c|}{\\; \\hspace{20em} \\; } \\\\ \\hline"

var T135 = " \\multirow{2}{*}{评分等级} \n        & \\multicolumn{2}{|c|}{\\; \\hspace{5em} \\; } & \\multicolumn{2}{|c|}{\\; \\hspace{5em} \\; } & \\multicolumn{2}{|c|}{\\; \\hspace{5em} \\; }\n        \\\\ \\cline{2-6}\n        & \\multicolumn{2}{|c|}{xx} & \\multicolumn{2}{|c|}{xx} & \\multicolumn{2}{|c|}{xx} \n        \\\\ \\cline{2-7}\n\n        & \\multicolumn{2}{|c|}{\\; \\hspace{5em} \\; } & \\multicolumn{2}{|c|}{\\; \\hspace{5em} \\; } & \\multicolumn{2}{|c|}{\\; \\hspace{5em} \\; }\n        \\\\ \\cline{2-6}\n        & \\multicolumn{2}{|c|}{xx} & \\multicolumn{2}{|c|}{xx} & \\multicolumn{2}{|c|}{xx} \n        \\\\ \\hline"

var T14 = "{备注} & \\multicolumn{6}{|c|}{评分等级共五种：优秀、良好、中等、及格、不及格} \\\\ \\hline \n\\end{tabular}\n\n\\newpage\n\\begin{center}\n    {\\songti \\zihao{0} \\textbf{{"

var T15 = "程序设计报告} }}\n\\end{center}\n\n\\section{课题内容和要求}\n{"

var T16 = "}\n\n\\section{概要设计}\n{"

var T17 = "}\n\n\n\\section{源程序代码}\n{"

var T18 = "}\n\n\n\\section{测试数据及其结果分析}\n{"

var T19 = "}\n\n\\section{调试过程中的问题}\n{"

var T20 = "}\n\n\\section{课程设计总结}\n{"

var T21 = "}\n\n\\newpage\n\n\\begin{thebibliography}{99}\n    \\bibitem{a}作者. \\emph{文献}[M]. 地点:出版社,年份.\n    \\bibitem{b}作者. \\emph{文献}[M]. 地点:出版社,年份.\n\\end{thebibliography}\n\n\\newpage\n\n\\begin{appendices}\n    \\renewcommand{\\thesection}{\\Alph{section}}\n    \\section{附录标题}\n        这里是附录. \n\\end{appendices}\n\n\\end{document}"

var Suojin = "\\indent\\setlength{\\parindent}{2em}"

var Hanghuang = "\\"

var Tbg = "\\begin{tabular}{|"

var Tcow = "c|"

var Tdp = "}\n\t\\hline "

var Ted = "\\hline\n\\end{tabular}\\"
