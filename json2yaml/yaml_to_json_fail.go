package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ss = `Services:
-   Orders:
    -   ID: $save ID1
        SupplierOrderCode: $SupplierOrderCode
    -   ID: $save ID2
        SupplierOrderCode: 111111
`
const ss2 = `Services: hello
Services2: hello`

func main() {
	var body map[string]interface{}
	err := yaml.Unmarshal([]byte(ss), &body)
	fmt.Println(err)
	b, err2 := json.Marshal(body)
	fmt.Println(string(b))
	fmt.Println(err2)
}
