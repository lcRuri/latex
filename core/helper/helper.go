package helper

import (
	"bufio"
	"core/define"
	"core/internal/types"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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

func PdfGenerate(filepath, filename, content string) error {

	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal("os.OpenFile", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
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
	f := getwd + "/file/" + filename
	c2 := exec.Command("xelatex", "-output-directory=pdf", f)
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c2.Wait()

	//wg.Done()
}

func HandleUpload(w http.ResponseWriter, request *http.Request) (filepath, filename string) {
	//文件上传只允许POST方法
	if request.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}

	//从表单中读取文件
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		_, _ = io.WriteString(w, "Read file error")
		return
	}
	//defer 结束时关闭文件
	defer file.Close()
	filename = fileHeader.Filename
	log.Println("filename: " + fileHeader.Filename)

	//创建文件
	filepath = define.BaseUploadPath + "/" + fileHeader.Filename
	newFile, err := os.Create(filepath)
	if err != nil {
		_, _ = io.WriteString(w, "Create file error")
		return
	}
	//defer 结束时关闭文件
	defer newFile.Close()

	//将文件写到本地
	_, err = io.Copy(newFile, file)
	if err != nil {
		_, _ = io.WriteString(w, "Write file error")
		return
	}
	_, _ = io.WriteString(w, "Upload success")

	return filepath, filename
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
