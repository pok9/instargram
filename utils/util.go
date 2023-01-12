package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(val interface{}) {
	str, _ := json.MarshalIndent(val, "", "\t")
	fmt.Println(string(str))
}
