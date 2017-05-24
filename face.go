package cognitive

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type FaceDetectResult struct {
	FaceId         string           `json:"faceId"`
	FaceRectangle  Rectangle        `json:"faceRectangle"`
	FaceLandmarks  map[string]Point `json:"faceLandmarks"`
	FaceAttributes struct {
		Age        float64            `json:"age"`
		Gender     string             `json:"gender"`
		Smile      float64            `json:"smile"`
		FacialHair map[string]float64 `json:"facialHair"`
		Glasses    string             `json:"glasses"`
		HeadPose   map[string]float64 `json:"headPose"`
		Emotion    map[string]float64 `json:"emotion"`
	} `json:"faceAttributes"`
}

type FaceFindSimilarRequest1 struct {
	FaceId                    string `json:"faceId"`
	FaceListId                string `json:"faceListId"`
	MaxNumOfCandidateReturned string `json:"maxNumOfCandidatesReturned"`
	Mode                      string `json:"mode"`
}

type FaceFindSimilarRequest2 struct {
	FaceId                    string   `json:"faceId"`
	FaceIds                   []string `json:"faceIds"`
	MaxNumOfCandidateReturned string   `json:"maxNumOfCandidatesReturned"`
	Mode                      string   `json:"mode"`
}

type FaceFindSimilarResult struct {
	PersistedFaceId string  `json:"persistedFaceId"`
	FaceId          string  `json:"faceId"`
	Confidence      float64 `json:"confidence"`
}

type FaceGroupRequest struct {
	FaceIds []string `json:"faceIds"`
}

type FaceGroupResult struct {
	Groups     [][]string `json:"groups"`
	MessyGroup []string   `json:"messyGroup"`
}

type FaceIdentifyRequest1 struct {
	FaceIds                   []string `json:"faceIds"`
	PersonGroupId             string   `json:"personGroupId"`
	MaxNumOfCandidateReturned int      `json:"maxNumOfCandidatesReturned"`
	ConfidenceThreshold       float64  `json:"confidenceThreshold"`
}

type FaceIdentifyRequest2 struct {
	FaceIds                   []string `json:"faceIds"`
	PersonGroupId             string   `json:"personGroupId"`
	MaxNumOfCandidateReturned int      `json:"maxNumOfCandidatesReturned"`
}

type FaceIdentifyResult struct {
	FaceId     string `json:"faceId"`
	Candidates []struct {
		PersonId   string  `json:"personId"`
		Confidence float64 `json:"confidence"`
	} `json:"candidates"`
}

type FaceVerifyRequest1 struct {
	FaceId1 string `json:"faceId1"`
	FaceId2 string `json:"faceId2"`
}

type FaceVerifyRequest2 struct {
	FaceId        string `json:"faceId"`
	PersonGroupId string `json:"personGroupId"`
	PersonId      string `json:"personId"`
}

type FaceVerifyResult struct {
	IsIdentifical bool    `json:"isIdentical"`
	Confidence    float64 `json:"confidence"`
}

type FaceAddToListResult struct {
	PersistedFaceId string `json:"persistedFaceId"`
}

type FaceCreateFaceListRequest struct {
	Name     string `json:"name"`
	UserData string `json:"userData"`
}

type FaceFacesResult struct {
	FaceListId     string `json:"faceListId"`
	Name           string `json:"name"`
	UserData       string `json:"userData"`
	PersistedFaces []struct {
		PersistedFaceId string `json:"persistedFaceId"`
		UserData        string `json:"userData"`
	} `json:"persistedFaces"`
}

type FaceListResult struct {
	FaceListId string `json:"faceListId"`
	Name       string `json:"name"`
	UserData   string `json:"userData"`
}

type FaceUpdateFaceListRequest struct {
	Name     string `json:"name"`
	UserData string `json:"userData"`
}

type FaceAddPersonFaceResult struct {
	PersistedFaceId string `json:"persistedFaceId"`
}

type FaceCreatePersonRequest struct {
	Name     string `json:"name"`
	UserData string `json:"userData"`
}

type FaceCreatePersonResult struct {
	PersonId string `json:"personId"`
}

type FaceGetPersonResult struct {
	PersonId         string   `json:"personId"`
	PersistedFaceIds []string `json:"persistedFaceIds"`
	Name             string   `json:"name"`
	UserData         string   `json:"userData"`
}

type FaceGetPersonFaceResult struct {
	PersistedFaceId string `json:"persistedFaceId"`
	UserData        string `json:"userData"`
}

type FaceGetPersonsResult struct {
	PersonId         string   `json:"personId"`
	Name             string   `json:"name"`
	UserData         string   `json:"userData"`
	PersistedFaceIds []string `json:"persistedFaceIds"`
}

type FaceUpdatePersonRequest struct {
	Name     string `json:"name"`
	UserData string `json:"userData"`
}

type FaceUpdatePersonFaceRequest struct {
	UserData string `json:"userData"`
}

type FaceCreatePersonGroupRequest struct {
	Name     string `json:"name"`
	UserData string `json:"userData"`
}

type FaceGetPersonGroupResult struct {
	PersonGroupId string `json:"personGroupId"`
	Name          string `json:"name"`
	UserData      string `json:"userData"`
}

type FaceGetPersonGroupTrainingStatusResult struct {
	Status string `json:"status"`
	/*
		// XXX - in format of: '1/3/2017 4:11:35 AM'
		CreatedDateTime    time.Time `json:"createdDateTime"`
		LastActionDateTime time.Time `json:"lastActionDateTime"`
	*/
	CreatedDateTime    string `json:"createdDateTime"`
	LastActionDateTime string `json:"lastActionDateTime"`
	Message            string `json:"message"`
}

type FaceGetPersonGroupsResult struct {
	PersonGroupId string `json:"personGroupId"`
	Name          string `json:"name"`
	UserData      string `json:"userData"`
}

type FaceUpdatePersonGroupRequest struct {
	Name     string `json:"name"`
	UserData string `json:"userData"`
}

// Face API: Detect
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395236
//
// location             : API location
// key                  : subscription key for this API
// image                : string(image url) or []byte(image bytes array)
// returnFaceId         : (default: true)
// returnFaceLandmarks  : (default: false)
// returnFaceAttributes : "age", "gender", "headPose", "smile", "facialHair", "glasses", or "emotion"
func FaceDetect(
	location ApiLocation,
	key string,
	image interface{},
	returnFaceId bool,
	returnFaceLandmarks bool,
	returnFaceAttributes []string,
) (processResult []FaceDetectResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/detect"

	// params
	params := map[string]string{}
	if !returnFaceId {
		params["returnFaceId"] = "false"
	}
	if returnFaceLandmarks {
		params["returnFaceLandmarks"] = "true"
	}
	if len(returnFaceAttributes) > 0 {
		params["returnFaceAttributes"] = strings.Join(returnFaceAttributes, ",")
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, image)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return []FaceDetectResult{}, err
}

// Face API: Find Similar
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395237
//
// location                   : API location
// key                        : subscription key for this API
// faceId                     : (can get from FaceDetect())
// faceListId                 : (can get from FaceListCreate())
// faceIds                    : (can get from FaceDetect())
// maxNumOfCandidatesReturned : 1 - 1000 (default: 20)
// mode                       : "matchPerson" or "matchFace" (default: "matchPerson")
func FaceFindSimilar(
	location ApiLocation,
	key string,
	faceId string,
	faceListId string,
	faceIds []string,
	maxNumOfCandidatesReturned int,
	mode string,
) (processResult []FaceFindSimilarResult, err error) {
	if faceListId != "" && len(faceIds) > 0 {
		return []FaceFindSimilarResult{}, fmt.Errorf("faceListId and faceIds cannot be provided at the same time")
	}

	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/findsimilars"

	// json object
	var obj interface{}
	if maxNumOfCandidatesReturned < 1 || maxNumOfCandidatesReturned > 1000 {
		maxNumOfCandidatesReturned = 20
	}
	if mode == "" {
		mode = "matchPerson"
	}
	if faceListId != "" {
		obj = FaceFindSimilarRequest1{
			FaceId:                    faceId,
			FaceListId:                faceListId,
			MaxNumOfCandidateReturned: strconv.Itoa(maxNumOfCandidatesReturned),
			Mode: mode,
		}
	} else if len(faceIds) > 0 {
		obj = FaceFindSimilarRequest2{
			FaceId:                    faceId,
			FaceIds:                   faceIds,
			MaxNumOfCandidateReturned: strconv.Itoa(maxNumOfCandidatesReturned),
			Mode: mode,
		}
	} else {
		return []FaceFindSimilarResult{}, fmt.Errorf("Both faceListId and faceIds are not provided")
	}

	var result []byte
	result, err = httpPost(apiUrl, key, nil, obj)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return []FaceFindSimilarResult{}, err
}

// Face API: Group
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395238
//
// location : API location
// key      : subscription key for this API
// faceIds  : (can get from FaceDetect())
func FaceGroup(
	location ApiLocation,
	key string,
	faceIds []string,
) (processResult FaceGroupResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/group"

	// json object
	obj := FaceGroupRequest{
		FaceIds: faceIds,
	}

	var result []byte
	result, err = httpPost(apiUrl, key, nil, obj)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceGroupResult{}, err
}

// Face API: Identify
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395239
//
// location                   : API location
// key                        : subscription key for this API
// faceIds                    : (can get from FaceDetect())
// personGroupId              : (can get from FacePersonGroup())
// maxNumOfCandidatesReturned : 1 - 5 (default: 1)
// confidenceThreshold        : 0.0 - 1.0 (default: set automatically)
func FaceIdentify(
	location ApiLocation,
	key string,
	faceIds []string,
	personGroupId string,
	maxNumOfCandidatesReturned int,
	confidenceThreshold float64,
) (processResult []FaceIdentifyResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/identify"

	// json object
	var obj interface{}
	if maxNumOfCandidatesReturned < 1 || maxNumOfCandidatesReturned > 5 {
		maxNumOfCandidatesReturned = 1
	}
	if confidenceThreshold >= 0.0 && confidenceThreshold <= 1.0 {
		obj = FaceIdentifyRequest1{
			FaceIds:                   faceIds,
			PersonGroupId:             personGroupId,
			MaxNumOfCandidateReturned: maxNumOfCandidatesReturned,
			ConfidenceThreshold:       confidenceThreshold,
		}
	} else {
		obj = FaceIdentifyRequest2{
			FaceIds:                   faceIds,
			PersonGroupId:             personGroupId,
			MaxNumOfCandidateReturned: maxNumOfCandidatesReturned,
		}
	}

	var result []byte
	result, err = httpPost(apiUrl, key, nil, obj)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return []FaceIdentifyResult{}, err
}

// Face API: Verify
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039523a
//
// location : API location
// key      : subscription key for this API
// obj      : FaceVerifyRequest1(face-to-face) or FaceVerifyRequest2(face-to-person)
func FaceVerify(
	location ApiLocation,
	key string,
	obj interface{},
) (processResult FaceVerifyResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/verify"

	// json object
	switch obj.(type) {
	case FaceVerifyRequest1, FaceVerifyRequest2: // ok
	default: // not ok
		return FaceVerifyResult{}, fmt.Errorf("Provided object type is not supported: %T", obj)
	}

	var result []byte
	result, err = httpPost(apiUrl, key, nil, obj)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceVerifyResult{}, err
}

// Face API: Add a Face to a Face List
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395250
//
// location   : API location
// key        : subscription key for this API
// image      : string(image url) or []byte(image bytes array)
// faceListId : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
// userData   : max 1kb
// targetFace : if there're more than 1 faces, it should be passed
func FaceAddFaceToList(
	location ApiLocation,
	key string,
	image interface{},
	faceListId string,
	userData string,
	targetFace Rectangle,
) (processResult FaceAddToListResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/facelists/" + faceListId + "/persistedFaces"

	// params
	params := map[string]string{}
	if userData != "" {
		params["userData"] = userData
	}
	if targetFace.Width > 0 && targetFace.Height > 0 {
		params["targetFace"] = fmt.Sprintf("%d,%d,%d,%d", targetFace.Left, targetFace.Top, targetFace.Width, targetFace.Height)
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, image)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceAddToListResult{}, err
}

// Face API: Create a Face List
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039524b
//
// location   : API location
// key        : subscription key for this API
// faceListId : id of a new face list
// name       : name of a new face list
// userData   : max 16kb
func FaceCreateFaceList(
	location ApiLocation,
	key string,
	faceListId string,
	name string,
	userData string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/facelists/" + faceListId

	// json object
	obj := FaceCreateFaceListRequest{
		Name:     name,
		UserData: userData,
	}

	_, err = httpPut(apiUrl, key, nil, obj)

	return err
}

// Face API: Delete a Face from a Face List
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395251
//
// location        : API location
// key             : subscription key for this API
// faceListId      : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
// persistedFaceId : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
func FaceDeleteFace(
	location ApiLocation,
	key string,
	faceListId string,
	persistedFaceId string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/facelists/" + faceListId + "/persistedFaces/" + persistedFaceId

	// params
	params := map[string]string{
		"faceListId":      faceListId,
		"persistedFaceId": persistedFaceId,
	}

	_, err = httpDelete(apiUrl, key, params)

	return err
}

// Face API: Delete a Face List
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039524f
//
// location        : API location
// key             : subscription key for this API
// faceListId      : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
func FaceDeleteFaceList(
	location ApiLocation,
	key string,
	faceListId string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/facelists/" + faceListId

	// params
	params := map[string]string{
		"faceListId": faceListId,
	}

	_, err = httpDelete(apiUrl, key, params)

	return err
}

// Face API: Get a Face List
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039524c
//
// location   : API location
// key        : subscription key for this API
// faceListId : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
func FaceGetFaces(
	location ApiLocation,
	key string,
	faceListId string,
) (processResult FaceFacesResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/facelists/" + faceListId

	var result []byte
	result, err = httpGet(apiUrl, key, nil)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceFacesResult{}, err
}

// Face API: List Face Lists
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039524d
//
// location   : API location
// key        : subscription key for this API
func FaceGetLists(
	location ApiLocation,
	key string,
) (processResult []FaceListResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/facelists"

	var result []byte
	result, err = httpGet(apiUrl, key, nil)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return []FaceListResult{}, err
}

// Face API: Update a Face List
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039524e
//
// location   : API location
// key        : subscription key for this API
// faceListId : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
// name       : max length = 128
// userData   : max 16kb
func FaceUpdateFaceList(
	location ApiLocation,
	key string,
	faceListId string,
	name string,
	userData string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/facelists/" + faceListId

	// object
	obj := FaceUpdateFaceListRequest{
		Name:     name,
		UserData: userData,
	}

	_, err = httpPatch(apiUrl, key, nil, obj)

	return err
}

// Face API: Add a Person Face
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039523b
//
// location      : API location
// key           : subscription key for this API
// image         : string(image url) or []byte(image bytes array)
// personGroupId :
// personId      :
// userData      : max 1kb
// targetFace    :
func FaceAddPersonFace(
	location ApiLocation,
	key string,
	image interface{},
	personGroupId string,
	personId string,
	userData string,
	targetFace Rectangle,
) (processResult FaceAddPersonFaceResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons/" + personId + "/persistedFaces"

	// params
	params := map[string]string{}
	if userData != "" {
		params["userData"] = userData
	}
	if targetFace.Width > 0 && targetFace.Height > 0 {
		params["targetFace"] = fmt.Sprintf("%d,%d,%d,%d", targetFace.Left, targetFace.Top, targetFace.Width, targetFace.Height)
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, image)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceAddPersonFaceResult{}, err
}

// Face API: Create a Person
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039523c
//
// location      : API location
// key           : subscription key for this API
// image         : string(image url) or []byte(image bytes array)
// personGroupId :
// name          : max length = 128
// userData      : max 16kb
func FaceCreatePerson(
	location ApiLocation,
	key string,
	image interface{},
	personGroupId string,
	name string,
	userData string,
) (processResult FaceCreatePersonResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons"

	// params
	obj := FaceCreatePersonRequest{
		Name:     name,
		UserData: userData,
	}

	var result []byte
	result, err = httpPost(apiUrl, key, nil, obj)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceCreatePersonResult{}, err
}

// Face API: Delete a Person
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039523d
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
// personId      :
func FaceDeletePerson(
	location ApiLocation,
	key string,
	personGroupId string,
	personId string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons/" + personId

	_, err = httpDelete(apiUrl, key, nil)

	return err
}

// Face API: Delete a Person Face
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039523e
//
// location        : API location
// key             : subscription key for this API
// personGroupId   :
// personId        :
// persistedFaceId :
func FaceDeletePersonFace(
	location ApiLocation,
	key string,
	personGroupId string,
	personId string,
	persistedFaceId string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons/" + personId + "/persistedFaces/" + persistedFaceId

	_, err = httpDelete(apiUrl, key, nil)

	return err
}

// Face API: Get a Person
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039523f
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
// personId      :
func FaceGetPerson(
	location ApiLocation,
	key string,
	personGroupId string,
	personId string,
) (processResult FaceGetPersonResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons/" + personId

	var result []byte
	result, err = httpGet(apiUrl, key, nil)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceGetPersonResult{}, err
}

// Face API: Get a Person Face
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395240
//
// location        : API location
// key             : subscription key for this API
// personGroupId   :
// personId        :
// persistedFaceId :
func FaceGetPersonFace(
	location ApiLocation,
	key string,
	personGroupId string,
	personId string,
	persistedFaceId string,
) (processResult FaceGetPersonFaceResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons/" + personId + "/persistedFaces/" + persistedFaceId

	var result []byte
	result, err = httpGet(apiUrl, key, nil)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceGetPersonFaceResult{}, err
}

// Face API: List Persons in a Person Group
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395241
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
func FaceGetPersons(
	location ApiLocation,
	key string,
	personGroupId string,
) (processResult []FaceGetPersonsResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons"

	var result []byte
	result, err = httpGet(apiUrl, key, nil)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return []FaceGetPersonsResult{}, err
}

// Face API: Update a Person
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395242
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
// personId      :
// name          : max length = 128
// userData      : max 16kb
func FaceUpdatePerson(
	location ApiLocation,
	key string,
	personGroupId string,
	personId string,
	name string,
	userData string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons/" + personId

	// object
	obj := FaceUpdatePersonRequest{
		Name:     name,
		UserData: userData,
	}

	_, err = httpPatch(apiUrl, key, nil, obj)

	return err
}

// Face API: Update a Person Face
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395243
//
// location        : API location
// key             : subscription key for this API
// personGroupId   :
// personId        :
// persistedFaceId :
// userData        : max 1kb
func FaceUpdatePersonFace(
	location ApiLocation,
	key string,
	personGroupId string,
	personId string,
	persistedFaceId string,
	userData string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/persons/" + personId + "/persistedFaces/" + persistedFaceId

	// object
	obj := FaceUpdatePersonFaceRequest{
		UserData: userData,
	}

	_, err = httpPatch(apiUrl, key, nil, obj)

	return err
}

// Face API: Create a Person Group
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395244
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
// name          : max length = 128
// userData      : max 16kb
func FaceCreatePersonGroup(
	location ApiLocation,
	key string,
	personGroupId string,
	name string,
	userData string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId

	// json object
	obj := FaceCreatePersonGroupRequest{
		Name:     name,
		UserData: userData,
	}

	_, err = httpPut(apiUrl, key, nil, obj)

	return err
}

// Face API: Delete a Person Group
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395245
//
// location        : API location
// key             : subscription key for this API
// personGroupId   :
func FaceDeletePersonGroup(
	location ApiLocation,
	key string,
	personGroupId string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId

	_, err = httpDelete(apiUrl, key, nil)

	return err
}

// Face API: Get a Person Group
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395246
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
func FaceGetPersonGroup(
	location ApiLocation,
	key string,
	personGroupId string,
) (processResult FaceGetPersonGroupResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId

	var result []byte
	result, err = httpGet(apiUrl, key, nil)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceGetPersonGroupResult{}, err
}

// Face API: Get Person Group Training Status
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395247
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
func FaceGetPersonGroupTrainingStatus(
	location ApiLocation,
	key string,
	personGroupId string,
) (processResult FaceGetPersonGroupTrainingStatusResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/training"

	var result []byte
	result, err = httpGet(apiUrl, key, nil)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return FaceGetPersonGroupTrainingStatusResult{}, err
}

// Face API: List Person Groups
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395248
//
// location : API location
// key      : subscription key for this API
// start    : 0 - 64 characters
// top      : 1 - 1000 (default 1000)
func FaceGetPersonGroups(
	location ApiLocation,
	key string,
	start string,
	top int,
) (processResult []FaceGetPersonGroupsResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups"

	params := map[string]string{}
	if start != "" {
		params["start"] = start
	}
	if top < 1 || top > 1000 {
		top = 1000
	}
	params["top"] = strconv.Itoa(top)

	var result []byte
	result, err = httpGet(apiUrl, key, params)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return []FaceGetPersonGroupsResult{}, err
}

// Face API: Train Person Group
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395249
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
func FaceTrainPersonGroup(
	location ApiLocation,
	key string,
	personGroupId string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId + "/train"

	_, err = httpPost(apiUrl, key, nil, nil)

	return err
}

// Face API: Update a Person Group
//
// https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039524a
//
// location      : API location
// key           : subscription key for this API
// personGroupId :
func FaceUpdatePersonGroup(
	location ApiLocation,
	key string,
	personGroupId string,
	name string,
	userData string,
) (err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/face/v1.0/persongroups/" + personGroupId

	// object
	obj := FaceUpdatePersonGroupRequest{
		Name:     name,
		UserData: userData,
	}

	_, err = httpPatch(apiUrl, key, nil, obj)

	return err
}
