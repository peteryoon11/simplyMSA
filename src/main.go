package main

import (
	"log"
	"net/http"

	"../pkg/HttpMethodModule"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// have to include
	// 1. server 와 client 역할을 같이 해야 하나?
	// 2. 토큰 발행은 얼마 마다 해야 하나?
	//fmt.Println(os.Args)
	router := httprouter.New()
	/* 	router.GET("/", Index)
	   	router.GET("/hello/:name", Hello) */
	//router.POST("/getAllBook", getAllBook)
	//router.POST(GetBookInfo.Uri(), getMyBook)
	//router.POST(GetBookInfo.URI(), getMyBook)

	/*
		Create = PUT with a new URI
		POST to a base URI returning a newly created URI
		Read   = GET
		Update = PUT with an existing URI
		Delete = DELETE
	*/

	router.PUT("/Book", HttpMethodModule.GetMyBook)
	// 최초에 등록? 할때

	router.GET("/Book", HttpMethodModule.GetMyBook)
	// 조회만

	router.PUT("/Book", HttpMethodModule.GetMyBook)
	// 책 추가

	router.POST("/Book", HttpMethodModule.GetMyBook)
	// 책 정보 변경 / update

	router.DELETE("/Book", HttpMethodModule.GetMyBook)

	// 책 삭제

	//router.POST("/getUserInfo", getUser)
	//router.POST("/getUserInfo/:test", getUser)

	log.Fatal(http.ListenAndServe(":8090", router))
}
