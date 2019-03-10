package main

import (
	"fmt"
	"github.com/ixuzhi/testmode"
	testmode_v2 "github.com/ixuzhi/testmode/v2"
)

func main() {
	fmt.Println(testmode.Hi("user testmode case!"))
	fmt.Println(testmode_v2.Hi("user testmode case!","ch"))
}
