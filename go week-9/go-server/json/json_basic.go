package fakejson

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
	Address
}

type Address struct {
	HouseName string `json:"house_name"`
	Pincode   int    `json:"pincode"`
	Area      string `json:"area"`
}

type Employee struct {
	FullName   string `json:"full_name"`
	EmployeeId int    `json:"emp_id"`
	Age        int    `json:"age"`
	Post       string `json:"post"`
	IsActive   bool   `json:"isactive"`
}

func jsonpro() {
	p := Person{Name: "Tharun", Age: 20, Address: Address{
		HouseName: "Mra 39",
		Pincode:   683110,
		Area:      "Aluva",
	}}
	fmt.Println(p)

	jsondata, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsondata))

	// mn := Employee{
	// 	FullName:   "Tharun",
	// 	EmployeeId: 2299,
	// 	Age:        20,
	// 	Post:       "Manager",
	// 	IsActive:   true,
	// }

	// js, e := json.Marshal(mn)
	// if e != nil {

	// }
	// fmt.Println(string(js))

	jsondata1 := `{"full_name":"Tharun V S","emp_id":2399,"age":20,"post":"CEO","isactive":true}`

	var mainemp Employee

	er := json.Unmarshal([]byte(jsondata1), &mainemp)
	if er != nil {
		fmt.Println("error :", er)
	}
	fmt.Printf("%+v", mainemp)
	fmt.Println()
}
