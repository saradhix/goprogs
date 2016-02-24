package main

import "fmt"
import "stringutil"
import "square"
import "encoding/json"

//import "encoding/json"
//import "gabs"

func main() {
	fmt.Printf("Hello, World\n")
	fmt.Printf(stringutil.Reverse("\nXXThis is baad"))
	fmt.Printf("%d\n", square.Sqr(15))
	//config := make(map[string]string)
	var config = make(map[string]string)
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
	for k, v := range m {
		fmt.Printf("Key=%s value=%s\n", k, v)
	}
	x, y := 10, 20
	fmt.Printf("%d %d ", x, y)
}
