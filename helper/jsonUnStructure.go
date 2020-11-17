package helper

import (
	"encoding/json"
	"fmt"
)

// Bird
type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func mainJDON() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{})

	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
	//pigeon likes to perch on rocks
	//eagle bird of prey
}

func maptojson() {
	// The keys need to be strings, the values can be
	// any serializable value
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	// JSON encoding is done the same way as before
	data, _ := json.Marshal(birdData)
	fmt.Println(string(data))
}
