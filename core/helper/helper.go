package helper

import (
	"bufio"
	"core/define"
	"core/internal/types"
	"core/models"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//var wg sync.WaitGroup

func GenerateToken(id int, identity string, name string, second int) (string, error) {
	user := &define.UserClaim{
		Id:             id,
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)

	tokenString, err := token.SignedString([]byte(define.JwtSecret))
	if err != nil {
		log.Printf("token.SignedString err:%v", err)
		return "", err
	}
	return tokenString, nil
}

func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtSecret), nil
	})
	if err != nil {
		log.Printf("jwt.ParseWithClaims err:%v", err)
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}

	return uc, nil
}

func UUid() string {
	return uuid.NewV4().String()
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func PdfGenerate(pdf *models.Pdf, filename string, filepath string) error {

	semester := "二"
	if time.Now().Month() >= 9 && time.Now().Month() <= 12 {
		semester = "一"
	}
	header := define.T1 + string(time.Now().Year()) +
		define.T2 + semester +
		define.T3 + pdf.Title +
		define.T4 + pdf.Subject +
		define.T5 + pdf.GroupLeaderName +
		define.T6 + pdf.GroupMemberName +
		define.T7 + pdf.Teacher +
		define.T8 + pdf.Company +
		define.T9 + string(time.Now().Year()) + string(time.Now().Month()) + string(time.Now().Day())

	LeaderWorkDivideModel := ""
	for _, strings := range pdf.LeaderWorkDivide {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			LeaderWorkDivideModel += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			LeaderWorkDivideModel += tabletmp
		}
	}

	MemberWorkDivide := ""
	for _, strings := range pdf.MemberWorkDivide {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			MemberWorkDivide += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			MemberWorkDivide += tabletmp
		}
	}

	Requirement := ""
	for _, strings := range pdf.Requirement {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			Requirement += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			Requirement += tabletmp
		}
	}

	DemandAnalysis := ""
	for _, strings := range pdf.DemandAnalysis {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			DemandAnalysis += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			DemandAnalysis += tabletmp
		}
	}

	OutlineDesign := ""
	for _, strings := range pdf.OutlineDesign {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			OutlineDesign += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			OutlineDesign += tabletmp
		}
	}

	SourceCode := ""
	for _, strings := range pdf.SourceCode {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			SourceCode += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			SourceCode += tabletmp
		}
	}

	TestAndResult := ""
	for _, strings := range pdf.TestAndResult {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			TestAndResult += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			TestAndResult += tabletmp
		}
	}

	Question := ""
	for _, strings := range pdf.Question {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			Question += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			Question += tabletmp
		}
	}

	Summary := ""
	for _, strings := range pdf.Summary {
		if strings.Type == "paragraph" {
			tmp := define.Suojin + strings.TextContent[0] + define.Hanghuang
			Summary += tmp
		}
		if strings.Type == "table" {
			//row, _ := strconv.Atoi(strings.TextContent[0])
			column, _ := strconv.Atoi(strings.TextContent[1])

			tabletmp := define.Tbg
			for i := 0; i < column; i++ {
				tmp := define.Tcow
				tabletmp += tmp
			}

			tabletmp += define.Tdp

			for i := 2; i < len(strings.TextContent); i++ {
				tmp := ""
				if (i-1)%column == 0 {
					tmp = strings.TextContent[i] + "\\" + "\n\t\\hline "
				} else {
					tmp = strings.TextContent[i] + "&"
				}
				tabletmp += tmp
			}

			tabletmp += define.Ted

			Summary += tabletmp
		}
	}

	modelArticle := header + LeaderWorkDivideModel + MemberWorkDivide + Requirement + DemandAnalysis + OutlineDesign + SourceCode + TestAndResult + Question + Summary

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(modelArticle)
	if err != nil {
		log.Fatal(err)
		return err
	}
	writer.Flush()
	//go 开启协程
	//wg.Add(1)
	go osPdf(filename)

	//pdfPath := "/pdf/" + filename

	//wg.Wait()

	return nil
}

func osPdf(filename string) {
	getwd, _ := os.Getwd()
	f := getwd + "/file/njupt" + filename
	c2 := exec.Command("xelatex", "-output-directory=pdf", f)
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c2.Wait()

	//wg.Done()
}

func Upload(c *gin.Context, path string) {
	c.Header("Content-Type", "application/pdf") // 这里是压缩文件类型 .zip
	c.File(path)
}

func Hash(filepath string) string {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		fmt.Errorf("读取文件失败！")
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x\n", sum)
}

//func Download(w http.ResponseWriter, r *http.Request, filepath string) {
//	filename := r.FormValue("filename")
//	f, err := ioutil.ReadFile(filepath + ".pdf")
//	if err != nil {
//		fmt.Fprintln(w, "文件下载失败", err)
//		return
//	}
//	h := w.Header()
//	h.Set("Content-type", "application/pdf")
//	h.Set("Content-Disposition", "attachment;filename="+filename)
//	w.Write(f)
//}

func Download(w http.ResponseWriter, r *http.Request) {
	filepath := r.FormValue("filepath")
	fmt.Println(filepath)
	resp := new(types.PdfReply)
	f, err := ioutil.ReadFile("./pdf/" + filepath + ".pdf")
	if err != nil {
		fmt.Errorf("文件下载失败", err)
		resp.ResponseData.Code = 501
		resp.ResponseData.Msg = "文件下载失败"

		return
	}
	h := w.Header()
	h.Set("Content-type", "application/pdf")
	h.Set("Content-Disposition", "attachment;filename="+filepath+".pdf")
	w.Write(f)
}
