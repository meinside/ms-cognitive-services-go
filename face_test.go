package cognitive

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func init() {
	//IsVerbose = true // XXX - if you wanna see verbose messages, uncomment it

	// read keys
	initTestKeys()
}

func TestFace(t *testing.T) {
	var faceId1, faceId2 string

	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["celebrity-face-image"]); err == nil {
		if result, err := FaceDetect(
			WestUS,
			testKeys["face-subscription-key"],
			imgBytes,
			true,
			true,
			[]string{"age", "gender", "headPose", "smile", "facialHair", "glasses", "emotion"},
		); err == nil {
			fmt.Printf("FaceDetect() => %+v\n", result)

			faceId1 = result[0].FaceId
			faceIds := []string{}
			for _, r := range result {
				faceIds = append(faceIds, r.FaceId)
			}

			// find similar
			if result, err := FaceFindSimilar(
				WestUS,
				testKeys["face-subscription-key"],
				faceId1,
				"",
				faceIds,
				1,
				"",
			); err == nil {
				fmt.Printf("FaceFindSimilar() => %+v\n", result)
			} else {
				t.Errorf("FaceFindSimilar() failed: %s\n", err)
			}
		} else {
			t.Errorf("FaceDetect() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}

	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["face-image1"]); err == nil {
		if result, err := FaceDetect(
			WestUS,
			testKeys["face-subscription-key"],
			imgBytes,
			true,
			true,
			[]string{"age", "gender", "headPose", "smile", "facialHair", "glasses", "emotion"},
		); err == nil {
			faceId2 = result[0].FaceId

			// group
			if result, err := FaceGroup(
				WestUS,
				testKeys["face-subscription-key"],
				[]string{faceId1, faceId2},
			); err == nil {
				fmt.Printf("FaceGroup() => %+v\n", result)
			} else {
				t.Errorf("FaceGroup() failed: %s\n", err)
			}
		} else {
			t.Errorf("FaceDetect() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}

	// verify
	if result, err := FaceVerify(
		WestUS,
		testKeys["face-subscription-key"],
		FaceVerifyRequest1{
			FaceId1: faceId1,
			FaceId2: faceId2,
		},
	); err == nil {
		fmt.Printf("FaceVerify() => %+v\n", result)
	} else {
		t.Errorf("FaceVerify() failed: %s\n", err)
	}
}

func TestFace_List(t *testing.T) {
	newFaceListId := "test-00001"
	newFaceListName := "test-list"

	// create a face list
	if err := FaceCreateFaceList(
		WestUS,
		testKeys["face-subscription-key"],
		newFaceListId,
		newFaceListName,
		"this face list is for test",
	); err == nil {
		fmt.Printf("FaceCreateFaceList() => success\n")

		// add a face to list
		if imgBytes, err := ioutil.ReadFile(testKeys["face-image1"]); err == nil {
			if result, err := FaceAddFaceToList(
				WestUS,
				testKeys["face-subscription-key"],
				imgBytes,
				newFaceListId,
				"this face is for test",
				Rectangle{},
			); err == nil {
				fmt.Printf("FaceAddFaceToList() => %+v\n", result)
			} else {
				t.Errorf("FaceAddFaceToList() failed: %s\n", err)
			}
		} else {
			fmt.Printf("File read error.\n")
		}

		// list faces
		if result, err := FaceGetFaces(
			WestUS,
			testKeys["face-subscription-key"],
			newFaceListId,
		); err == nil {
			fmt.Printf("FaceGetFaces() => %+v\n", result)

			for _, f := range result.PersistedFaces {
				// delete a face from a list
				if err := FaceDeleteFace(
					WestUS,
					testKeys["face-subscription-key"],
					newFaceListId,
					f.PersistedFaceId,
				); err == nil {
					fmt.Printf("FaceDeleteFace() => success\n")
				} else {
					t.Errorf("FaceDeleteFace() failed: %s\n", err)
				}
			}
		} else {
			t.Errorf("FaceGetFaces() failed: %s\n", err)
		}

		// get face lists
		if result, err := FaceGetLists(
			WestUS,
			testKeys["face-subscription-key"],
		); err == nil {
			fmt.Printf("FaceGetLists() => %+v\n", result)

			for _, l := range result {
				// update a face list
				if err := FaceUpdateFaceList(
					WestUS,
					testKeys["face-subscription-key"],
					l.FaceListId,
					"new-name",
					"new-data",
				); err == nil {
					fmt.Printf("FaceUpdateFaceList() => success\n")
				} else {
					t.Errorf("FaceUpdateFaceList() failed: %s\n", err)
				}

				// delete a face list
				if err := FaceDeleteFaceList(
					WestUS,
					testKeys["face-subscription-key"],
					l.FaceListId,
				); err == nil {
					fmt.Printf("FaceDeleteFaceList() => success\n")
				} else {
					t.Errorf("FaceDeleteFaceList() failed: %s\n", err)
				}
			}
		} else {
			t.Errorf("FaceGetLists() failed: %s\n", err)
		}
	} else {
		t.Errorf("FaceCreateFaceList() failed: %s\n", err)
	}
}

func TestFace_Person_PersonGroup(t *testing.T) {
	newPersonGroupId := "test-group-00001"
	newPersonGroupName := "test-group"

	// create a person group
	if err := FaceCreatePersonGroup(
		WestUS,
		testKeys["face-subscription-key"],
		newPersonGroupId,
		newPersonGroupName,
		"this person group is for test",
	); err == nil {
		fmt.Printf("FaceCreatePersonGroup() => success\n")

		// update a person group
		if err := FaceUpdatePersonGroup(
			WestUS,
			testKeys["face-subscription-key"],
			newPersonGroupId,
			"new-name",
			"new-data",
		); err == nil {
			fmt.Printf("FaceUpdatePersonGroup() => success\n")
		} else {
			t.Errorf("FaceUpdatePersonGroup() failed: %s\n", err)
		}

		// create a person
		if imgBytes, err := ioutil.ReadFile(testKeys["face-image1"]); err == nil {
			if result, err := FaceCreatePerson(
				WestUS,
				testKeys["face-subscription-key"],
				imgBytes,
				newPersonGroupId,
				"test-person",
				"this person is for test",
			); err == nil {
				fmt.Printf("FaceCreatePerson() => %+v\n", result)

				// update a person
				personId := result.PersonId
				if err := FaceUpdatePerson(
					WestUS,
					testKeys["face-subscription-key"],
					newPersonGroupId,
					personId,
					"test-person-updated",
					"this person is still for test",
				); err == nil {
					fmt.Printf("FaceUpdatePerson() => success\n")
				} else {
					t.Errorf("FaceUpdatePerson() failed: %s\n", err)
				}

				// add a person face
				if imgBytes, err := ioutil.ReadFile(testKeys["face-image2"]); err == nil {
					if result, err := FaceAddPersonFace(
						WestUS,
						testKeys["face-subscription-key"],
						imgBytes,
						newPersonGroupId,
						personId,
						"another face",
						Rectangle{},
					); err == nil {
						fmt.Printf("FaceAddPersonFace() => %+v\n", result)

						persistedFaceId := result.PersistedFaceId

						// update a person face
						if err := FaceUpdatePersonFace(
							WestUS,
							testKeys["face-subscription-key"],
							newPersonGroupId,
							personId,
							persistedFaceId,
							"updated face",
						); err == nil {
							fmt.Printf("FaceUpdatePersonFace() => success\n")
						} else {
							t.Errorf("FaceUpdatePersonFace() failed: %s\n", err)
						}
					} else {
						t.Errorf("FaceAddPersonFace() failed: %s\n", err)
					}
				} else {
					fmt.Printf("File read error.\n")
				}

			} else {
				t.Errorf("FaceCreatePerson() failed: %s\n", err)
			}
		} else {
			fmt.Printf("File read error.\n")
		}

		// train person group
		if err := FaceTrainPersonGroup(
			WestUS,
			testKeys["face-subscription-key"],
			newPersonGroupId,
		); err == nil {
			fmt.Printf("FaceTrainPersonGroup() => success\n")
		} else {
			t.Errorf("FaceTrainPersonGroup() failed: %s\n", err)
		}

		// get person group training status
		if result, err := FaceGetPersonGroupTrainingStatus(
			WestUS,
			testKeys["face-subscription-key"],
			newPersonGroupId,
		); err == nil {
			fmt.Printf("FaceGetPersonGroupTrainingStatus() => %+v\n", result)
		} else {
			t.Errorf("FaceGetPersonGroupTrainingStatus() failed: %s\n", err)
		}

		// identify
		if imgBytes, err := ioutil.ReadFile(testKeys["face-image1"]); err == nil {
			if result, err := FaceDetect(
				WestUS,
				testKeys["face-subscription-key"],
				imgBytes,
				true,
				true,
				[]string{"age", "gender", "headPose", "smile", "facialHair", "glasses", "emotion"},
			); err == nil {
				faceId := result[0].FaceId

				if result, err := FaceIdentify(
					WestUS,
					testKeys["face-subscription-key"],
					[]string{faceId},
					newPersonGroupId,
					5,
					-1,
				); err == nil {
					fmt.Printf("FaceIdentify() => %+v\n", result)
				} else {
					t.Errorf("FaceIdentify() failed: %s\n", err)
				}
			} else {
				t.Errorf("FaceDetect() failed: %s\n", err)
			}
		} else {
			fmt.Printf("File read error.\n")
		}

		// list persons in a person group
		if result, err := FaceGetPersons(
			WestUS,
			testKeys["face-subscription-key"],
			newPersonGroupId,
		); err == nil {
			fmt.Printf("FaceGetPersons() => %+v\n", result)

			person := result[0]

			// get a person
			if result, err := FaceGetPerson(
				WestUS,
				testKeys["face-subscription-key"],
				newPersonGroupId,
				person.PersonId,
			); err == nil {
				fmt.Printf("FaceGetPerson() => %+v\n", result)
			} else {
				t.Errorf("FaceGetPerson() failed: %s\n", err)
			}

			// get a person face
			if result, err := FaceGetPersonFace(
				WestUS,
				testKeys["face-subscription-key"],
				newPersonGroupId,
				person.PersonId,
				person.PersistedFaceIds[0],
			); err == nil {
				fmt.Printf("FaceGetPersonFace() => %+v\n", result)
			} else {
				t.Errorf("FaceGetPersonFace() failed: %s\n", err)
			}

			// delete a person face
			if err := FaceDeletePersonFace(
				WestUS,
				testKeys["face-subscription-key"],
				newPersonGroupId,
				person.PersonId,
				person.PersistedFaceIds[0],
			); err == nil {
				fmt.Printf("FaceDeletePersonFace() => success\n")
			} else {
				t.Errorf("FaceDeletePersonFace() failed: %s\n", err)
			}

			// delete a person
			if err := FaceDeletePerson(
				WestUS,
				testKeys["face-subscription-key"],
				newPersonGroupId,
				person.PersonId,
			); err == nil {
				fmt.Printf("FaceDeletePerson() => success\n")
			} else {
				t.Errorf("FaceDeletePerson() failed: %s\n", err)
			}
		} else {
			t.Errorf("FaceGetPersons() failed: %s\n", err)
		}

		// list person groups
		if result, err := FaceGetPersonGroups(
			WestUS,
			testKeys["face-subscription-key"],
			"",
			100,
		); err == nil {
			fmt.Printf("FaceGetPersonGroups() => %+v\n", result)
		} else {
			t.Errorf("FaceGetPersonGroups() failed: %s\n", err)
		}

		// get a person group
		if result, err := FaceGetPersonGroup(
			WestUS,
			testKeys["face-subscription-key"],
			newPersonGroupId,
		); err == nil {
			fmt.Printf("FaceGetPersonGroup() => %+v\n", result)
		} else {
			t.Errorf("FaceGetPersonGroup() failed: %s\n", err)
		}

		// delete a person group
		if err := FaceDeletePersonGroup(
			WestUS,
			testKeys["face-subscription-key"],
			newPersonGroupId,
		); err == nil {
			fmt.Printf("FaceDeletePersonGroup() => success\n")
		} else {
			t.Errorf("FaceDeletePersonGroup() failed: %s\n", err)
		}
	} else {
		t.Errorf("FaceCreatePersonGroup() failed: %s\n", err)
	}
}
