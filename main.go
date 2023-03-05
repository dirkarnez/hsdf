package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type HSDF struct {
	location string
}

func newHSDF(location string) *HSDF {

	err := os.Mkdir(location, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return &HSDF{location}
}

/*
{
     "firstName": "John",
     "lastName": "Smith",
     "sex": "male",
     "age": 25,
     "address":
     {
         "streetAddress": "21 2nd Street",
         "city": "New York",
         "state": "NY",
         "postalCode": "10021"
     },
     "phoneNumber":
     [
         {
           "type": "home",
           "number": "212 555-1234"
         },
         {
           "type": "fax",
           "number": "646 555-4567"
         }
     ]
 }
*/

type HSDFCurrentNode struct {
	currentKey string
}

func (hsdf HSDF) CreateArray(key string, arrayValue []string) error {
	err := os.MkdirAll(strings.Join([]string{hsdf.location, key}, "/"), 0755)
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range arrayValue {
		ioutil.WriteFile(strings.Join([]string{hsdf.location, key, fmt.Sprintf("%d.json", i)}, "/"), []byte(fmt.Sprintf("\"%s\"", s)), 0644)
	}
	return nil
}

func (hsdf HSDF) CreateObject(key string, objectCreated func(currentNode HSDFCurrentNode)) error {
	err := os.MkdirAll(strings.Join([]string{hsdf.location, key}, "/"), 0755)
	if err != nil {
		log.Fatal(err)
	}

	objectCreated(struct{ currentKey string }{
		currentKey: key,
	})

	return nil
}

func (hsdf HSDF) CreateTupleString(currentNode HSDFCurrentNode, key string, value string) error {
	return ioutil.WriteFile(strings.Join([]string{hsdf.location, currentNode.currentKey, fmt.Sprintf("%s.json", key)}, "/"), []byte(fmt.Sprintf("\"%s\"", value)), 0644)
}

func main() {
	hsdfObject := newHSDF("testobject")
	/*
		"phoneNumber": ["a", "b", "c"]
	*/
	hsdfObject.CreateArray("phoneNumber", []string{"a", "b", "c"})
	/*
	   "address": {
	       "streetAddress": "21 2nd Street",
	       "city": "New York",
	       "state": "NY",
	       "postalCode": "10021"
	   }
	*/
	hsdfObject.CreateObject("address", func(currentNode HSDFCurrentNode) {
		hsdfObject.CreateTupleString(currentNode, "streetAddress", "21 2nd Street")
		hsdfObject.CreateTupleString(currentNode, "city", "New York")
		hsdfObject.CreateTupleString(currentNode, "state", "NY")
		hsdfObject.CreateTupleString(currentNode, "postalCode", "10021")
	})
	fmt.Println("done")
}
