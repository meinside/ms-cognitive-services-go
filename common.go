package cognitive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Rectangle struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ApiResponseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type ApiResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type OperationStatus struct {
	Status             string    `json:"status"`
	Progress           float32   `json:"progress"`
	CreatedDateTime    time.Time `json:"createdDateTime"`
	LastActionDateTime time.Time `json:"lastActionDateTime"`

	ProcessingResultJson string `json:"processingResult"`
	ResourceLocation     string `json:"resourceLocation"`

	Message string `json:"message"`
}

const (
	NumTries    = 100
	WaitSeconds = 10
)

type ApiLocation string

const (
	WestUS        ApiLocation = "westus"
	EastUS2       ApiLocation = "eastus2"
	WestCentralUS ApiLocation = "westcentralus"
	WestEurope    ApiLocation = "westeurope"
	SoutheastAsia ApiLocation = "southeastasia"
)

// for showing verbose messages
var IsVerbose bool = false

// for testing only
var testKeys map[string]string = nil

const (
	confFilenameForTest = "files/testconf.json"
)

func initTestKeys() {
	if testKeys == nil {
		if _, err := os.Stat(confFilenameForTest); !os.IsNotExist(err) {
			if data, err := ioutil.ReadFile(confFilenameForTest); err == nil {
				testKeys = map[string]string{}

				if err := json.Unmarshal(data, &testKeys); err != nil {
					log.Printf("Failed to parse config file: %s", err)

					testKeys = nil
				}
			}
		}
	}

	if testKeys == nil {
		panic("Failed to read config file for testing.")
	}
}

// http request with bytes
func httpRequest(method, url, key string, headers, params map[string]string, data []byte, contentType string) (response *http.Response, err error) {
	var req *http.Request
	if req, err = http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(data)); err == nil {
		// http headers
		if contentType == "" && len(data) > 0 {
			contentType = http.DetectContentType(data)
		}
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}
		req.Header.Set("Ocp-Apim-Subscription-Key", key)
		for k, v := range headers { // additional http headers
			req.Header.Set(k, v)
		}

		// get params
		query := req.URL.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()

		client := &http.Client{}
		return client.Do(req)
	}
	return nil, err
}

// http request with a json object
func requestJson(method, url, key string, headers, params map[string]string, object interface{}) (response *http.Response, err error) {
	var data []byte
	if data, err = json.Marshal(object); err == nil {
		return httpRequest(method, url, key, headers, params, data, "application/json")
	}
	return nil, err
}

// http get
func httpGet(url, key string, params map[string]string) (result []byte, err error) {
	var resp *http.Response
	if resp, err = httpRequest("get", url, key, nil, params, nil, ""); err == nil {
		defer resp.Body.Close()

		if result, err = ioutil.ReadAll(resp.Body); err == nil {
			if IsVerbose {
				log.Printf("<< request: %s, %+v", url, params)
				log.Printf(">> response: %s", string(result))
			}

			if resp.StatusCode == 200 {
				return result, nil
			} else {
				var errResp ApiResponseError
				if err = json.Unmarshal(result, &errResp); err == nil {
					if errResp.StatusCode > 0 {
						err = fmt.Errorf("HTTP %d; %s", errResp.StatusCode, errResp.Message)
					} else {
						var resp ApiResponse
						if err = json.Unmarshal(result, &resp); err == nil {
							err = fmt.Errorf("%s; %s", resp.Error.Code, resp.Error.Message)
						}
					}
				}
			}
		}
	}
	return []byte{}, err
}

// http post with an object
func httpPost(url, key string, params map[string]string, object interface{}) (result []byte, err error) {
	var resp *http.Response
	if resp, err = requestJson("post", url, key, nil, params, object); err == nil {
		defer resp.Body.Close()

		if result, err = ioutil.ReadAll(resp.Body); err == nil {
			if IsVerbose {
				log.Printf("<< request: %s, %+v, %+v", url, params, object)
				log.Printf(">> response: %s", string(result))
			}

			if resp.StatusCode == 200 {
				return result, nil
			} else if resp.StatusCode == 202 {
				location := resp.Header.Get("Operation-Location")

				if IsVerbose {
					log.Printf(">> operation location: %s", location)
				}

				return []byte(location), nil
			} else {
				var errResp ApiResponseError
				if err = json.Unmarshal(result, &errResp); err == nil {
					if errResp.StatusCode > 0 {
						err = fmt.Errorf("HTTP %d; %s", errResp.StatusCode, errResp.Message)
					} else {
						var resp ApiResponse
						if err = json.Unmarshal(result, &resp); err == nil {
							err = fmt.Errorf("%s; %s", resp.Error.Code, resp.Error.Message)
						}
					}
				}
			}
		}
	}
	return []byte{}, err
}

// http post with a bytes array
func httpPostBytes(url, key string, params map[string]string, bts []byte) (result []byte, err error) {
	var resp *http.Response
	if resp, err = httpRequest("post", url, key, nil, params, bts, "application/octet-stream"); err == nil {
		defer resp.Body.Close()

		if result, err = ioutil.ReadAll(resp.Body); err == nil {
			if IsVerbose {
				log.Printf("<< request: %s, %+v, %d bytes", url, params, len(bts))
				log.Printf(">> response: %s", string(result))
			}

			if resp.StatusCode == 200 {
				return result, nil
			} else if resp.StatusCode == 202 {
				location := resp.Header.Get("Operation-Location")

				if IsVerbose {
					log.Printf(">> operation location: %s", location)
				}

				return []byte(location), nil
			} else {
				var errResp ApiResponseError
				if err = json.Unmarshal(result, &errResp); err == nil {
					if errResp.StatusCode > 0 {
						err = fmt.Errorf("HTTP %d; %s", errResp.StatusCode, errResp.Message)
					} else {
						var resp ApiResponse
						if err = json.Unmarshal(result, &resp); err == nil {
							err = fmt.Errorf("%s; %s", resp.Error.Code, resp.Error.Message)
						}
					}
				}
			}
		}
	}
	return []byte{}, err
}

func postArg(url, key string, params map[string]string, arg interface{}) (result []byte, err error) {
	switch arg.(type) {
	case string: // => url
		if a, ok := arg.(string); ok {
			result, err = httpPost(
				url,
				key,
				params,
				struct {
					Url string `json:"url"`
				}{Url: a},
			)
		} else {
			err = fmt.Errorf("Could not convert given parameter to string")
		}
	case []byte: // => bytes array
		if a, ok := arg.([]byte); ok {
			result, err = httpPostBytes(
				url,
				key,
				params,
				a,
			)
		} else {
			err = fmt.Errorf("Could not convert given parameter to []byte")
		}
	default:
		err = fmt.Errorf("Given parameter type (%T) is not supported", arg)
	}
	return
}

// http put with an object
func httpPut(url, key string, params map[string]string, object interface{}) (result []byte, err error) {
	return httpMethod("put", url, key, params, object)
}

// http delete
func httpDelete(url, key string, params map[string]string) (result []byte, err error) {
	return httpMethod("delete", url, key, params, nil)
}

// http patch
func httpPatch(url, key string, params map[string]string, object interface{}) (result []byte, err error) {
	return httpMethod("patch", url, key, params, object)
}

func httpMethod(method, url, key string, params map[string]string, object interface{}) (result []byte, err error) {
	var resp *http.Response

	if IsVerbose {
		switch object.(type) {
		case []byte:
			if b, ok := object.([]byte); ok {
				log.Printf("<< request (HTTP %s): %s, %+v, %d bytes", method, url, params, len(b))
			}
		default:
			log.Printf("<< request (HTTP %s): %s, %+v, %+v", method, url, params, object)
		}
	}

	if resp, err = requestJson(method, url, key, nil, params, object); err == nil {
		defer resp.Body.Close()

		if result, err = ioutil.ReadAll(resp.Body); err == nil {
			if IsVerbose {
				log.Printf(">> response (HTTP %d): %s", resp.StatusCode, string(result))
			}

			if resp.StatusCode == 200 {
				return result, nil
			} else if resp.StatusCode == 202 {
				location := resp.Header.Get("Operation-Location")

				if IsVerbose {
					log.Printf(">> operation location: %s", location)
				}

				return []byte(location), nil
			} else {
				var errResp ApiResponseError
				if err = json.Unmarshal(result, &errResp); err == nil {
					if errResp.StatusCode > 0 {
						err = fmt.Errorf("HTTP %d; %s", errResp.StatusCode, errResp.Message)
					} else {
						var resp ApiResponse
						if err = json.Unmarshal(result, &resp); err == nil {
							err = fmt.Errorf("%s; %s", resp.Error.Code, resp.Error.Message)
						}
					}
				}
			}
		}
	}
	return []byte{}, err
}

// download a file from given url
func Download(url, key, localPath string) (err error) {
	var out *os.File
	if out, err = os.Create(localPath); err == nil {
		defer out.Close()

		var resp *http.Response
		if resp, err = httpRequest("get", url, key, nil, nil, nil, ""); err == nil {
			defer resp.Body.Close()

			_, err = io.Copy(out, resp.Body)
		}
	}
	return err
}
