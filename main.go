package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	// Decode
	sd, _ := base64.StdEncoding.DecodeString(secret)

	//Convert to string
	sd_convert_to_str := string(sd)

	// Reverse string
	sd_reverse_str := reverseString(sd_convert_to_str)

	// Replace
	result := strings.ReplaceAll(sd_reverse_str, ":", " ")
	fmt.Printf("Result: %s\n", result)

}
