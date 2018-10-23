package HttpMethodModule

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../dbConnectModule"
	"../structModule"
	"../validationModule"
	"github.com/julienschmidt/httprouter"
)

func GetMyBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
