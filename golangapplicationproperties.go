package golangapplicationproperties // import "dakine/golangapplicationproperties"

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ApplicationProperties Properties

type Properties struct {
	FilePath    string
	PropertyMap map[string]string
}

func check(e error) {
	if e != nil {
		fmt.Print("check :")
		fmt.Println(e)
	}
}

var propertyMap = make(map[string]string)

func (ap *Properties) GetProps() map[string]string {

	return ap.PropertyMap
}

func (ap *Properties) setFilePath(newFilePath string) {
	ap.FilePath = newFilePath
}

func (ap *Properties) Init(newFilePath string) Properties {
	ApplicationProperties.setFilePath(newFilePath)
	ApplicationProperties.readProps()
	return ApplicationProperties
}

func (ap *Properties) Current() Properties {
	return ApplicationProperties
}

func (ap *Properties) readProps() error {

	ap.PropertyMap = make(map[string]string)

	file, err := os.Open(ap.FilePath)
	check(err)

	reader := bufio.NewReader(file)

	for {

		lineAsBytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineAsString := strings.Split(string(lineAsBytes), "=")
		check(err)
		if len(strings.TrimSpace(lineAsString[0])) > 0 {
			ap.PropertyMap[strings.TrimSpace(lineAsString[0])] = strings.TrimSpace(lineAsString[len(lineAsString)-1])
		}

	}

	file.Close()
	return nil

}

func (ap *Properties) IsInitialized() bool {
	return (len(ap.PropertyMap) > 0)
}

//func main() {
//	a, b := IsInitialized()
//	println(a, b)
//	Init("./application.properties")
//	a, b = IsInitialized()
//	println(a, b)
//	m := GetProps()

//	println(m["zalos.sproc.datasource.maxConnectionAgeInSeconds"])

//}
