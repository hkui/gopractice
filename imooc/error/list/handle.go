package list

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const Prefix  = "/static/"
type AppHandle func( http.ResponseWriter,  *http.Request)error

type userError string

func (e userError) Error()string {
	return  e.Message()
}
func (e userError)Message() string  {
	return  string(e)
}

func Handle(writer http.ResponseWriter, request *http.Request) error {
	allPath:=request.URL.Path
	if strings.Index(allPath,Prefix)!=0{
		//return errors.New(fmt.Sprintf("must start with :%s\n",Prefix))
		return userError(fmt.Sprintf("must start with :%s\n",Prefix))
	}


	path := request.URL.Path[len(Prefix):]


	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	_, _ = writer.Write(all)

	return nil
}
