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
	modelArticle := define.T1 + pdf.T1 +
		define.T2 + pdf.T2 +
		define.T3 + pdf.T3 +
		define.T4 + pdf.T4 +
		define.T5 + pdf.T5 +
		define.T6 + pdf.T6 +
		define.T7 + pdf.T7 +
		define.T8 + pdf.T8 +
		define.T9 + pdf.T9 +
		define.T10 + pdf.T10 +
		define.T11 + pdf.T11 +
		define.T12 + pdf.T12 +
		define.T13 + pdf.T13 +
		define.T135 +
		define.T14 + pdf.T14 +
		define.T15 + pdf.T15 +
		define.T16 + pdf.T16 +
		define.T17 + pdf.T17 +
		define.T18 + pdf.T18 +
		define.T19 + pdf.T19 +
		define.T20 + pdf.T20 +
		define.T21

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
