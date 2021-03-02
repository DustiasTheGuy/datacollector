package datacollector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type DataCollector struct {
	WriteToFile bool
}

func New() *DataCollector {

	dc := new(DataCollector)

	return dc
}

// Collect should collect all information being requested
func (d *DataCollector) Collect() error {
	if err := d.FindAllLinks("https://stackoverflow.com/questions/13573269/convert-string-to-byte"); err != nil {
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
	regex := regexp.MustCompile(`href=".*"`)
	//regex := regexp.MustCompile(`<a href=".*">`)
	strs := regex.FindAllString(b, -1)

	var out []byte

	for i := 0; i < len(strs); i++ {
		out = append(out, []byte(fmt.Sprintf("%s\n", strs[i]))...)
	}

	d.Output(out)
	return nil
}

func (d *DataCollector) Output(data []byte) {
	if d.WriteToFile {
		if err := ioutil.WriteFile("./output/output.html", data, 0644); err != nil {
			panic(err)
		}
	}
}
