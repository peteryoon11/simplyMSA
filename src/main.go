package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"../pkg/dbConnectModule"
	"../pkg/structModule"
	"../pkg/validationModule"
	"github.com/julienschmidt/httprouter"
)

func getMyBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//fmt.Println(ps.ByName("test"))
	//mem := Member{"Alex", 10, true}

	test, err := ioutil.ReadAll(r.Body)

	//	var testmem Member
	var testAuth structModule.ValAPIKey

	err = json.Unmarshal(test, &testAuth)

	var respondUser structModule.Response

	var tempBookArray []structModule.EBookInfo

	fmt.Println(testAuth)

	if validationModule.CheckAPIToken(testAuth) {
		tempBookArray, err = dbConnectModule.GetMyOwnBook(testAuth)
		respondUser = structModule.Response{200, "Respond Success", tempBookArray}
	} else {

		respondUser = structModule.Response{403, "Invalid Auth key or Expire key", tempBookArray}
	}

	temp, err := json.Marshal(respondUser)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println()
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(temp)

}
func main() {
	// have to include
	// 1. server 와 client 역할을 같이 해야 하나?
	// 2. 토큰 발행은 얼마 마다 해야 하나?
	fmt.Println(os.Args)
	router := httprouter.New()
	/* 	router.GET("/", Index)
	   	router.GET("/hello/:name", Hello) */
	//router.POST("/getAllBook", getAllBook)
	//router.POST(GetBookInfo.Uri(), getMyBook)
	//router.POST(GetBookInfo.URI(), getMyBook)

	/* Create = PUT with a new URI
	POST to a base URI returning a newly created URI
	Read   = GET
	Update = PUT with an existing URI
	Delete = DELETE
	*/
	router.GET("/getMyBook", getMyBook)
	router.POST("/getMyBook", getMyBook)
	router.DELETE("/getMyBook", getMyBook)

	//router.POST("/getUserInfo", getUser)
	//router.POST("/getUserInfo/:test", getUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
