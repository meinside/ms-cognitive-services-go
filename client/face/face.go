package face

// Wrapper Client for Face functions

import (
	"github.com/meinside/ms-cognitive-services-go"
)

type Client struct {
	Location cognitive.ApiLocation
	ApiKey   string
}

func NewClient(apiKey string) *Client {
	return &Client{
		Location: cognitive.WestUS,
		ApiKey:   apiKey,
	}
}

func NewClientWithLocation(location cognitive.ApiLocation, apiKey string) *Client {
	return &Client{
		Location: location,
		ApiKey:   apiKey,
	}
}

// Detect
//
// image               : string(image url) or []byte(image bytes array)
// returnFaceId        : (default: true)
// returnFaceLandmarks : (default: false)
// faceAttributes      : "age", "gender", "headPose", "smile", "facialHair", "glasses", or "emotion"
func (c *Client) Detect(
	image interface{},
	returnFaceId bool,
	returnFaceLandmarks bool,
	returnFaceAttributes []string,
) (processResult []cognitive.FaceDetectResult, err error) {
	return cognitive.FaceDetect(
		c.Location,
		c.ApiKey,
		image,
		returnFaceId,
		returnFaceLandmarks,
		returnFaceAttributes,
	)
}

// Find Similar
//
// faceId                     : (can get from FaceDetect())
// faceListId                 : (can get from FaceListCreate())
// faceIds                    : (can get from FaceDetect())
// maxNumOfCandidatesReturned : 1 - 1000 (default: 20)
// mode                       : "matchPerson" or "matchFace" (default: "matchPerson")
func (c *Client) FindSimilar(
	faceId string,
	faceListId string,
	faceIds []string,
	maxNumOfCandidatesReturned int,
	mode string,
) (processResult []cognitive.FaceFindSimilarResult, err error) {
	return cognitive.FaceFindSimilar(
		c.Location,
		c.ApiKey,
		faceId,
		faceListId,
		faceIds,
		maxNumOfCandidatesReturned,
		mode,
	)
}

// Group
//
// faceIds : (can get from FaceDetect())
func (c *Client) Group(
	faceIds []string,
) (processResult cognitive.FaceGroupResult, err error) {
	return cognitive.FaceGroup(
		c.Location,
		c.ApiKey,
		faceIds,
	)
}

// Identify
//
// faceIds                    : (can get from FaceDetect())
// personGroupId              : (can get from FacePersonGroup())
// maxNumOfCandidatesReturned : 1 - 5 (default: 1)
// confidenceThreshold        : 0.0 - 1.0 (default: set automatically)
func (c *Client) Identify(
	faceIds []string,
	personGroupId string,
	maxNumOfCandidatesReturned int,
	confidenceThreshold float32,
) (processResult []cognitive.FaceIdentifyResult, err error) {
	return cognitive.FaceIdentify(
		c.Location,
		c.ApiKey,
		faceIds,
		personGroupId,
		maxNumOfCandidatesReturned,
		confidenceThreshold,
	)
}

// Verify
//
// obj : FaceVerifyRequest1(face-to-face) or FaceVerifyRequest2(face-to-person)
func (c *Client) Verify(
	obj interface{},
) (processResult cognitive.FaceVerifyResult, err error) {
	return cognitive.FaceVerify(
		c.Location,
		c.ApiKey,
		obj,
	)
}

// Add a Face to a Face List
//
// image      : string(image url) or []byte(image bytes array)
// faceListId : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
// userData   : max 1kb
// targetFace : if there're more than 1 faces, it should be passed
func (c *Client) AddFaceToList(
	image interface{},
	faceListId string,
	userData string,
	targetFace cognitive.Rectangle,
) (processResult cognitive.FaceAddToListResult, err error) {
	return cognitive.FaceAddFaceToList(
		c.Location,
		c.ApiKey,
		image,
		faceListId,
		userData,
		targetFace,
	)
}

// Create a Face List
//
// faceListId : id of a new face list
// name       : name of a new face list
// userData   : max 16kb
func (c *Client) CreateFaceList(
	faceListId string,
	name string,
	userData string,
) (err error) {
	return cognitive.FaceCreateFaceList(
		c.Location,
		c.ApiKey,
		faceListId,
		name,
		userData,
	)
}

// Delete a Face from a Face List
//
// faceListId      : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
// persistedFaceId : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
func (c *Client) DeleteFace(
	faceListId string,
	persistedFaceId string,
) (err error) {
	return cognitive.FaceDeleteFace(
		c.Location,
		c.ApiKey,
		faceListId,
		persistedFaceId,
	)
}

// Delete a Face List
//
// faceListId      : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
func (c *Client) DeleteFaceList(
	faceListId string,
) (err error) {
	return cognitive.FaceDeleteFaceList(
		c.Location,
		c.ApiKey,
		faceListId,
	)
}

// Get a Face List
//
// faceListId : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
func (c *Client) GetFaces(
	faceListId string,
) (processResult cognitive.FaceFacesResult, err error) {
	return cognitive.FaceGetFaces(
		c.Location,
		c.ApiKey,
		faceListId,
	)
}

// List Face Lists
func (c *Client) GetLists() (processResult []cognitive.FaceListResult, err error) {
	return cognitive.FaceGetLists(
		c.Location,
		c.ApiKey,
	)
}

// Update a Face List
//
// faceListId : (valid chars: letter in lower case or digit or '-' or '_', maximum length is 64)
// name       : max length = 128
// userData   : max 16kb
func (c *Client) UpdateFaceList(
	faceListId string,
	name string,
	userData string,
) (err error) {
	return cognitive.FaceUpdateFaceList(
		c.Location,
		c.ApiKey,
		faceListId,
		name,
		userData,
	)
}

// Add a Person Face
//
// image         : string(image url) or []byte(image bytes array)
// personGroupId :
// personId      :
// userData      : max 1kb
// targetFace    :
func (c *Client) AddPersonFace(
	image interface{},
	personGroupId string,
	personId string,
	userData string,
	targetFace cognitive.Rectangle,
) (processResult cognitive.FaceAddPersonFaceResult, err error) {
	return cognitive.FaceAddPersonFace(
		c.Location,
		c.ApiKey,
		image,
		personGroupId,
		personId,
		userData,
		targetFace,
	)
}

// Create a Person
//
// image         : string(image url) or []byte(image bytes array)
// personGroupId :
// name          : max length = 128
// userData      : max 16kb
func (c *Client) CreatePerson(
	image interface{},
	personGroupId string,
	name string,
	userData string,
) (processResult cognitive.FaceCreatePersonResult, err error) {
	return cognitive.FaceCreatePerson(
		c.Location,
		c.ApiKey,
		image,
		personGroupId,
		name,
		userData,
	)
}

// Delete a Person
//
// personGroupId :
// personId      :
func (c *Client) DeletePerson(
	personGroupId string,
	personId string,
) (err error) {
	return cognitive.FaceDeletePerson(
		c.Location,
		c.ApiKey,
		personGroupId,
		personId,
	)
}

// Delete a Person Face
//
// personGroupId   :
// personId        :
// persistedFaceId :
func (c *Client) DeletePersonFace(
	personGroupId string,
	personId string,
	persistedFaceId string,
) (err error) {
	return cognitive.FaceDeletePersonFace(
		c.Location,
		c.ApiKey,
		personGroupId,
		personId,
		persistedFaceId,
	)
}

// Get a Person
//
// personGroupId :
// personId      :
func (c *Client) GetPerson(
	personGroupId string,
	personId string,
) (processResult cognitive.FaceGetPersonResult, err error) {
	return cognitive.FaceGetPerson(
		c.Location,
		c.ApiKey,
		personGroupId,
		personId,
	)
}

// Get a Person Face
//
// personGroupId   :
// personId        :
// persistedFaceId :
func (c *Client) GetPersonFace(
	personGroupId string,
	personId string,
	persistedFaceId string,
) (processResult cognitive.FaceGetPersonFaceResult, err error) {
	return cognitive.FaceGetPersonFace(
		c.Location,
		c.ApiKey,
		personGroupId,
		personId,
		persistedFaceId,
	)
}

// List Persons in a Person Group
//
// personGroupId :
func (c *Client) GetPersons(
	personGroupId string,
) (processResult []cognitive.FaceGetPersonsResult, err error) {
	return cognitive.FaceGetPersons(
		c.Location,
		c.ApiKey,
		personGroupId,
	)
}

// Update a Person
//
// personGroupId :
// personId      :
// name          : max length = 128
// userData      : max 16kb
func (c *Client) UpdatePerson(
	personGroupId string,
	personId string,
	name string,
	userData string,
) (err error) {
	return cognitive.FaceUpdatePerson(
		c.Location,
		c.ApiKey,
		personGroupId,
		personId,
		name,
		userData,
	)
}

// Update a Person Face
//
// personGroupId   :
// personId        :
// persistedFaceId :
// userData        : max 1kb
func (c *Client) UpdatePersonFace(
	personGroupId string,
	personId string,
	persistedFaceId string,
	userData string,
) (err error) {
	return cognitive.FaceUpdatePersonFace(
		c.Location,
		c.ApiKey,
		personGroupId,
		personId,
		persistedFaceId,
		userData,
	)
}

// Create a Person Group
//
// personGroupId :
// name          : max length = 128
// userData      : max 16kb
func (c *Client) CreatePersonGroup(
	personGroupId string,
	name string,
	userData string,
) (err error) {
	return cognitive.FaceCreatePersonGroup(
		c.Location,
		c.ApiKey,
		personGroupId,
		name,
		userData,
	)
}

// Delete a Person Group
//
// personGroupId   :
func (c *Client) DeletePersonGroup(
	personGroupId string,
) (err error) {
	return cognitive.FaceDeletePersonGroup(
		c.Location,
		c.ApiKey,
		personGroupId,
	)
}

// Get a Person Group
//
// personGroupId :
func (c *Client) GetPersonGroup(
	personGroupId string,
) (processResult cognitive.FaceGetPersonGroupResult, err error) {
	return cognitive.FaceGetPersonGroup(
		c.Location,
		c.ApiKey,
		personGroupId,
	)
}

// Get Person Group Training Status
//
// personGroupId :
func (c *Client) GetPersonGroupTrainingStatus(
	personGroupId string,
) (processResult cognitive.FaceGetPersonGroupTrainingStatusResult, err error) {
	return cognitive.FaceGetPersonGroupTrainingStatus(
		c.Location,
		c.ApiKey,
		personGroupId,
	)
}

// List Person Groups
//
// start    : 0 - 64 characters
// top      : 1 - 1000 (default 1000)
func (c *Client) GetPersonGroups(
	start string,
	top int,
) (processResult []cognitive.FaceGetPersonGroupsResult, err error) {
	return cognitive.FaceGetPersonGroups(
		c.Location,
		c.ApiKey,
		start,
		top,
	)
}

// Train Person Group
//
// personGroupId :
func (c *Client) TrainPersonGroup(
	personGroupId string,
) (err error) {
	return cognitive.FaceTrainPersonGroup(
		c.Location,
		c.ApiKey,
		personGroupId,
	)
}

// Update a Person Group
//
// personGroupId :
func (c *Client) UpdatePersonGroup(
	personGroupId string,
	name string,
	userData string,
) (err error) {
	return cognitive.FaceUpdatePersonGroup(
		c.Location,
		c.ApiKey,
		personGroupId,
		name,
		userData,
	)
}
