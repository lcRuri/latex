package test

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"testing"
)

func Upload(c *gin.Context, path string) {
	c.Header("Content-Type", "application/pdf") // 这里是压缩文件类型 .zip
	c.File(path)
}

func TestFile(t *testing.T) {

	file, err := os.Open("./file/njupt.tex")
	defer file.Close()
	if err != nil {
		fmt.Errorf("读取文件失败！")
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	sum := hash.Sum(nil)

	fmt.Printf("%x\n", sum)
}
