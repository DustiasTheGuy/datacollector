package main

import (
	"fmt"

	"github.com/DustiasTheGuy/datacollector/datacollector"
)

func main() {
	dc := datacollector.New()

	fmt.Println(dc)
}
