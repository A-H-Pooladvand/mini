package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func ToJson(v any) {
	result, err := json.Marshal(v)
	result, err = json.MarshalIndent(v, "", "  ")

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("---------------------------------------------------------------------")
	fmt.Println()
	fmt.Printf(string(result))
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------")
}
