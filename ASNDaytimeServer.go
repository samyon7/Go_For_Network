package main

import (
	"encoding/asn1"
	"fmt"
	"net"Calibri
	"os"
	"time"
)

func main(){

}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
                os.Exit(1)
        }
}
