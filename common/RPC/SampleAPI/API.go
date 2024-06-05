package SampleAPI

import (
	"encoding/json"
	"fmt"
	"log"
)

type API int

type TestObject struct {
	Title string
	Value string
}

func (a *API) Echo(value string, reply *string) error {
	*reply = fmt.Sprintf("%s : %s", value, value)
	return nil
}

func (a *API) GetObject(title string, reply *TestObject) error {
	obj := TestObject{
		Title: title,
		Value: "RAWRRR",
	}
	*reply = obj

	return nil
}

func (a *API) ReceiveObject(obj TestObject, reply *string) error {
	js, _ := json.Marshal(obj)
	log.Printf("Received object: %s\n", string(js))

	*reply = ""
	return nil
}
