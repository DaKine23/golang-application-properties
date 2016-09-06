package golangapplicationproperties

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ApplicationProperties is the Singleton Value for props to init them only once systemwide
var ApplicationProperties Properties

// Properties is the type of ApplicationProperties the map contains all the data from the file
type Properties struct {
	FilePath    string
	PropertyMap map[string]string
	InitTime    time.Time
	Debug       bool
}

func (ap *Properties) check(e error) {
	if e != nil && ap.Debug {
		fmt.Print("GAP debug :")
		fmt.Println(e)
	}
}

var propertyMap = make(map[string]string)

// GetProps returns the content of the property file that was read at InitTime
func (ap *Properties) GetProps() map[string]string {

	return ap.PropertyMap
}

func (ap *Properties) setFilePath(newFilePath string) {
	ap.FilePath = newFilePath
}

// Init initialize the properties usually only needs to be done once the property file changes or application starts
func (ap *Properties) Init(newFilePath string, debug bool) Properties {
	ApplicationProperties.Debug = debug
	ApplicationProperties.setFilePath(newFilePath)
	ApplicationProperties.readProps()
	ApplicationProperties.InitTime = time.Now().Local().UTC()
	return ApplicationProperties
}

//GetInitTime gets the time when last init has taken place
func (ap *Properties) GetInitTime() time.Time {
	return ap.InitTime
}

// Current gets the ApplicationProperties Singleton Object
func (ap *Properties) Current() Properties {
	return ApplicationProperties
}

func (ap *Properties) readProps() error {

	ap.PropertyMap = make(map[string]string)

	file, err := os.Open(ap.FilePath)
	ap.check(err)

	reader := bufio.NewReader(file)

	for {

		lineAsBytes, _, err := reader.ReadLine()

		if err != nil {
			if err.Error() != "EOF" {
				ap.check(err)
			}
			break
		}
		lineAsString := strings.Split(string(lineAsBytes), "=")
		ap.check(err)
		if len(strings.TrimSpace(lineAsString[0])) > 0 {
			ap.PropertyMap[strings.TrimSpace(lineAsString[0])] = strings.TrimSpace(lineAsString[len(lineAsString)-1])
		}

	}

	file.Close()
	return nil

}

//IsInitialized tells you if the property file was already loaded
func (ap *Properties) IsInitialized() bool {
	return (len(ap.PropertyMap) > 0)
}
