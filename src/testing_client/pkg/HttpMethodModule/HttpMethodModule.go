package HttpMethodModule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	//"../dbConnectModule"
	//"../structModule"
	//"../validationModule"
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

func RequestViaHttp(reqType string, raw_url string, headers map[string]string, data []byte, timeoutSeconds int, download_path string, logger *get_module_logger.Logger) (int, []byte, map[string][]string, error) {
	var reader io.Reader = nil

	timeout := time.Duration(time.Duration(timeoutSeconds) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	if data != nil && len(data) > 0 {
		reader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(reqType, raw_url, reader)

	if err != nil {
		logger.Println(err)
		return 0, nil, nil, err
	}

	if nil != reader {
		req.Header.Set("Content-Length", fmt.Sprintf("%d", len(data)))
	}

	req.Close = true
	// I strongly advise setting user agent as some servers ignore request without it
	req.Header.Set("User-Agent", "MSA")
	if headers != nil {
		logger.Println("Request headers : ", headers)
		for k, v := range headers {
			if strings.EqualFold("Host", k) {
				req.Host = v
				continue
			}
			req.Header.Set(k, v)
		}
	}

	var (
		statusCode int
		body       []byte
		header     http.Header //map[string][]string
	)

	if bJsonReq, err := json.Marshal(req); nil != err {
		//logger.Println(err)
		logger.Println("Do request ", req)
	} else {
		logger.Println("Do request( Json ) ", string(bJsonReq))
	}

	resp, err := client.Do(req)
	if nil != err {
		logger.Println(err)
		return 0, nil, nil, err
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode
	header = resp.Header

	if 0 < len(download_path) {
		out, err := os.Create(download_path)
		if nil != err {
			logger.Println(err)
			return 0, nil, nil, err
		}
		defer out.Close()
		n, err := io.Copy(out, resp.Body)
		if nil != err {
			logger.Println(err)
			return 0, nil, nil, err
		}
		logger.Println(n, " bytes downloaded to ", download_path)
		body = nil
	} else {
		body, _ = ioutil.ReadAll(resp.Body)

		//wait and see until the test

		//		if strings.Contains(header.Get("Content-Type"), "utf-8") {
		//			r, _ := utf8.DecodeRune(body)
		//			body =
		//		}

		logger.Println("Status : ", statusCode, "\nHeaders : ", header)

	}
	return statusCode, body, header, err

}
