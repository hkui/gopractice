package list

import (
	"io/ioutil"
	"net/http"
	"os"
)

func Handle(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/static/"):]

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)

	return nil
}
