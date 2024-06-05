package RPC

import (
	"encoding/json"
	"fmt"
	"log"
)

type TestAPI int

type TestObject struct {
	Title string
	Value string
}

func (a *TestAPI) Echo(value string, reply *string) error {
	*reply = fmt.Sprintf("%s : %s", value, value)
	return nil
}

func (a *TestAPI) GetObject(title string, reply *TestObject) error {
	obj := TestObject{
		Title: title,
		Value: "RAWRRR",
	}
	*reply = obj

	return nil
}

func (a *TestAPI) ReceiveObject(obj TestObject, _ *string) error {
	js, _ := json.Marshal(obj)
	log.Printf("Received object: %s\n", string(js))

	return nil
}
