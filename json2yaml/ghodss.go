package main

import (
	"fmt"
	"github.com/ghodss/yaml"
)

//https://github.com/ghodss/yaml

const s3 = `Services:
-   Orders:
    -   ID: $save ID1
        SupplierOrderCode: $SupplierOrderCode
    -   ID: $save ID2
        SupplierOrderCode: 111111
line: 2
`
const s4 = ``

func main() {
	//j := []byte(`{"name": "John", "age": 30}`)
	j := []byte(s4)
	y, err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
	/* Output:
	name: John
	age: 30
	*/
	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j2))
	/* Output:
	{"age":30,"name":"John"}
	*/
}
