package helper

import (
	"io"
	"log"
	"net/http"
	"os"
)

func HandlerUploadPic(r *http.Request) error {
	file, header, err := r.FormFile("File")
	if err != nil {
		log.Fatal("r.FormFile err", err)
		return err
	}
	log.Printf("selected file name is %s", header.Filename)
	// 将文件拷贝到指定路径下，或者其他文件操作
	filepath := "./pic/" + header.Filename
	dst, err := os.Create(filepath)
	if err != nil {
		log.Fatal("os.Create err", err)
		return err
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Fatal("io.Copy(dst, file)", err)
		return err
	}

	return nil
}
