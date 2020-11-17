package helper

import (
	"encoding/json"
	"log"
)

// M istilah
type M map[string]interface{}

//type paramBody, return value "objPar" is JSON format
type paramBody struct {
	UserID   string    `json:"userID"`
	Function string    `json:"function"`
	ObjPar   parObjPar `json:"objPar"`
}

type parObjPar struct {
	QueryString struct {
		Query string `json:"query"`
	} `json:"query_string"`
}

func mainJson2() {
	b, err := json.Marshal(M{"query": M{"query_string": M{"query": "query goes here"}}})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("   As Map:", string(b))

	var q paramBody
	q.ObjPar.QueryString.Query = "query in a struct"
	b, err = json.Marshal(q)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("As Struct:", string(b))
}
