package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	jar, err := cookiejar.New(nil) // 쿠키를 저장할 cookiejar 인스턴스를 만든다
	if err != nil {
		panic(err)
	}

	client := http.Client{ // 쿠키를 저장할 수 있는 http.Client 인스턴스 생성
		Jar: jar,
	}

	for i := 0; i < 2; i++ { // 첫 번째 접근에서 쿠키를 받고, 두 번째 이후의 액세스에서 쿠키를 보내는 구조이므로 두 번 접근한다
		resp, err := client.Get("http://localhost:18888/cookie") // http.Get() 대신 작성한 클라이언트의 Get() 메서드로 액세스
		if err != nil {
			panic(err)
		}

		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}

		log.Println(string(dump))
	}
}
