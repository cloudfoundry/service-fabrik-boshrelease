package apiutils

import (
	"bytes"
	"encoding/json"
	"iaas-utils/ha/common/constants"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func InvokeRESTAPI(httpMethod string, apiURL string, reqHeader map[string]string, requestBody interface{}) (string, string, bool) {

	var httpClient *http.Client = &http.Client{}
	var request *http.Request
	var response *http.Response
	var responseStatus string
	var err error
	var responseStatusSuccessful bool

	log.Println("URL being hit is ", httpMethod, " ", apiURL)

	if requestBody == nil {
		request, err = http.NewRequest(httpMethod, apiURL, nil)
	} else {
		jsonValue, _ := json.Marshal(requestBody)
		request, err = http.NewRequest(httpMethod, apiURL, bytes.NewBuffer(jsonValue))
	}
	if err != nil {
		log.Println("Error occured while creating new request")
		return "Error occured while creating new HTTP Request", "", false
	}

	for key, value := range reqHeader {
		request.Header.Set(key, value)
	}

	response, err = httpClient.Do(request)
	if err != nil {
		log.Println("Error occured while caling the api: ", apiURL, "Error is ", err.Error())
		return err.Error(), "", false
	}

	responseStatus = strings.TrimSpace(strings.ToUpper(response.Status))
	log.Println("HTTP Response Status Code: ", responseStatus)
	// Check if the response is successful by mathcing for 2xx
	responseStatusSuccessful, _ = regexp.MatchString("^2[0-9][0-9] ", responseStatus)
	if !responseStatusSuccessful {
		log.Println("API call was not successful : ", apiURL)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("HTTP Status : ", responseStatus, "Error occurred while reading response body ", err.Error())
		return err.Error(), "", false
	}
	//	log.Println("response Body:", string(body))
	defer response.Body.Close()

	if responseStatusSuccessful {
		return string(body), responseStatus, true
	} else {
		log.Println("Response body:", string(body))
		return string(body), responseStatus, false
	}

}

func InvokeOAuthAPI(httpMethod string, apiURL string, requestBody map[string]string) (string, bool) {

	var data url.Values = url.Values{}
	var httpClient *http.Client = &http.Client{}
	var request *http.Request
	var response *http.Response
	var responseStatus string
	var err error

	for key, value := range requestBody {
		if key != "" {
			data.Add(key, value)
		}
	}

	log.Println("URL being hit is ", httpMethod, apiURL)
	request, err = http.NewRequest(httpMethod, apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Println("Error occured while creating new request")
		return "Error occured while creating new HTTP Request", false
	}
	// set appropriate content type
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// All required params have been set - call the api.
	response, err = httpClient.Do(request)
	if err != nil {
		log.Println("Error occured while caling the api: ", apiURL, "Error is ", err.Error())
		return err.Error(), false
	}

	responseStatus = response.Status
	log.Println("HTTP Response Status Code: ", responseStatus)
	if response.Status != constants.HTTP_STATUS_OK {
		log.Println("API call was not successful : ", apiURL)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("HTTP Status : ", responseStatus, "Error occurred while reading response body ", err.Error())
		return err.Error(), false
	}
	// log.Println("response Body:", string(body))
	defer response.Body.Close()

	if response.Status == constants.HTTP_STATUS_OK {
		return string(body), true
	} else {
		log.Println("Response body:", string(body))
		return string(body), false
	}

}
