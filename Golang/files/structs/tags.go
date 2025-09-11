package structs

import (
	"encoding/json"
	"fmt"
)
/*
`json:"page"`	--->  this is field name in json
*/
func Tags() {
	type response struct {
		Page  	 int 		  `json:"page"`
		Words []string	 `json:"words,omitempty"`
	}

	r := &response{Page: 1, Words: []string{"neel", "kumar"}}
	// r is pointer of response struct
	j, _ := json.Marshal(r)		// convert r ot json format

	fmt.Println(string(j))		// json to string

	// converting json back to struct
	var r2 response
	_ = json.Unmarshal(j, &r2)
	fmt.Println(r2)
}
