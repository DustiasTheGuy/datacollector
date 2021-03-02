package main

import (
	"github.com/DustiasTheGuy/datacollector/datacollector"
)

func main() {
	dc := datacollector.New()
	dc.Collect()
}
