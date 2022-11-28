package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buf bytes.Buffer                // 멀티파트부를 조립한 뒤 바이트열을 저장할 버퍼를 선언
	writer := multipart.NewWriter(&buf) // 멀티파트를 조합할 writer를 만든다
	writer.WriteField("name", "John")   // 파일 이외의 필드는 WriteField() 메서드로 등록한다

	// 파일을 읽는 조작.
	fileWriter, err := writer.CreateFormFile("file", "photo.jpg") // 개별 파일을 써넣을 io.Writer를 만든다.
	if err != nil {
		panic(err)
	}
	file, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(fileWriter, file) // 파일의 모든 콘텐츠를 io.Writer에 복사
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buf)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
