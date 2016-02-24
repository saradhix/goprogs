package main

import "fmt"
import "encoding/json"

func main() {
	fmt.Printf("Hello, World\n")
	config := make(map[string]string)
	config["username"] = "tesla"
	config["password"] = ""
	config["hostname"] = "localhost"
	config["port"] = "5432"
	jsonstr, err := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(jsonstr))
	if err != nil {
		panic(err)
	}
	jsonstr = []byte(`{ "hostname": "localhost", "password": "", "port":"5432", "username": "tesla" }`)
	var f interface{}
	err = json.Unmarshal(jsonstr, &f)
	m := f.(map[string]interface{})
	fmt.Println(f)
	fmt.Println(add(10, 20))
	/*	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle", v)
		}
	}*/
	fmt.Println(m["username"])
	fmt.Println(m["port"])
	x := m["port"]
	fmt.Printf("x=%d type=%T\n", x, x)
	m["name"] = "Vijay"
	fmt.Println(m)
	js2, err := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(js2))
}

func add(x int, y int) int {
	return x + y
}
