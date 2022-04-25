package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Before marshalling :",t.String())

	mdata, err := asn1.Marshal(t)
	checkError(err)
	fmt.Println("Marshalled ok")

	var newtime = new(time.Time)
	_, err1 := asn1.Unmarshal(mdata, newtime)
	checkError(err1)

	fmt.Println("After marshal/unmarshal :",newtime.String())
	s := "hello \u00bc"

	fmt.Println("Before marshalling : ",s)
}
