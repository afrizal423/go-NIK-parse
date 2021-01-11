package main

import (
	"encoding/json"
	"fmt"

	"github.com/afrizal423/go-NIK-parse/ParseNIK"
)

func main() {
	tesnik := "3307114404920004"
	coba := ParseNIK.GetdataNIK(tesnik)
	data, _ := json.Marshal(coba)
	fmt.Println(string(data))
}
