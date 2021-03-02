package datacollector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type DataCollector struct {
}

func New() *DataCollector {

	dc := new(DataCollector)

	return dc
}

// Collect should collect all information being requested
func (d *DataCollector) Collect() error {
	if err := d.FindAllLinks("https://en.wikipedia.org/wiki/Nilgai"); err != nil {
		return err
	}

	return nil
}

// FindAllLinks grabs all links on a page and stores them in a slice
func (d *DataCollector) FindAllLinks(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	b := string(body)
	regex := regexp.MustCompile(`<a\s+(?:[^>]*?\s+)?href="([^"]*)">`)
	strs := regex.FindAllString(b, -1)

	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
	return nil
}
