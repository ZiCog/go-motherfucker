package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

//	"time"
)

// Message type to marshal/unmarshal
type MessageTx struct {
	Name string
	Body string
	Time int64
}

type MessageRx struct {
	Name string
	Body string
}

// An instance of message type
var outMessage = MessageTx{"Alice", "Hello", 1294706395881547000}


// An example of a json message
var jasonBytes = []byte(`{
    "firstName": "John",
    "lastName": "Smith",
    "age": 25,
    "address": {
        "streetAddress": "21 2nd Street",
        "city": "New York",
        "state": "NY",
        "postalCode": 10021,
	"subObject": {"start": "something"}
    },
    "phoneNumber": [
        {
            "type": "home",
            "number": "212 555-1234",
     	    "subObject": ["start", "something"]
        },
        {
            "type": "fax",
            "number": "646 555-4567"
        }
    ]
}`)

func main() {
	fmt.Println("JSON test")

	// Marshal the struct data into a JSON message as bytes
	msgBytes, err := json.Marshal(outMessage)
	fmt.Printf("JSON = %s\nError = %v\n", msgBytes, err)

	var inMessage MessageRx

	// Unmarshal the JSON bytes into our message structure
	err = json.Unmarshal(msgBytes, &inMessage)
	fmt.Printf("Name  = %v\n", inMessage.Name)
	fmt.Printf("Body  = %v\n", inMessage.Body)
	//fmt.Printf("Time  = %v\n", inMessage.Time)
	fmt.Printf("Error = %v\n", err)

	// The magical "emty interface" can hold any type !
	var i interface{}
	i = "a string"
	fmt.Println(i)
	i = 2011
	fmt.Println(i)
	i = 2.777
	fmt.Println(i)

	// A "type assertion" will panic if the wrong type and we don't check ok
	r, ok := i.(float64)
	if ok {
		fmt.Println("The circle's area", math.Pi*r*r)
	} else {
		fmt.Println("r is not float64")
	}

	// We can switch on the type of a variable
	switch i.(type) {
	case int:
		fmt.Println("An interger")
	case float64:
		fmt.Println("A 64 bit float")
	case string:
		fmt.Println("A string")
	default:
		fmt.Println("Don't know this type")
	}

	fmt.Println("--------------------------------------------\n")
	
	// So..We can parse any JSON string into an empty interface 
	//	b := []byte(`{"Name": "Julie",  "dada": null, "Live": true, "Age":6,  "Parents":["Gomez", "Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(jasonBytes, &f)
	fmt.Printf("f is:\n%v\n", f)

	m, ok := f.(map[string]interface{})
	fmt.Printf("m is:\n%v\n", m)
	if !ok {
		fmt.Printf("Oops, our json did not produce a map\n")
		os.Exit(1)
	}

	fmt.Println("--------------------------------------------")
	for key, value := range m {
		switch valueType := value.(type) {
		case nil:
			fmt.Println(key, "(nil) =", valueType)
		case bool:
			fmt.Println(key, "(bool) =", valueType)
		case string:
			fmt.Println(key, "(string) =", valueType)
		case float64:
			fmt.Println(key, "(float64) =", valueType)
		case []interface{}:
			fmt.Println(key, "(array) =")
			for key, value := range valueType {
				fmt.Println("    ", key, value)
			}
		case map[string]interface{}:
			fmt.Println(key, "(object) =")
			for key, value := range valueType {
				fmt.Println("    ", key, value)
			}
		default:
			fmt.Println(key, "(unknown) = ")
		}
	}
	fmt.Println("--------------------------------------------")
}
