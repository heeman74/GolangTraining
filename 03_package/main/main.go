package main

import (
	"fmt"

	"github.com/heeman74/GolangTraining/03_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG, olleH")) // function starts with uppercase letter will be "public"
	// and starts with lowercase letter will be "private"
	fmt.Println(stringutil.Myname)
}
