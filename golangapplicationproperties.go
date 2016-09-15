package golangapplicationproperties

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// ApplicationProperties is the Singleton Value for props to init them only once systemwide
var ApplicationProperties Properties

// Properties is the type of ApplicationProperties the map contains all the data from the file
type Properties struct {
	FilePath    string
	PropertyMap map[string]string
	InitTime    time.Time
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
func (ap *Properties) Init(newFilePath string) (Properties, error) {
	ApplicationProperties.setFilePath(newFilePath)
	err := ApplicationProperties.readProps()
	if err != nil {
		return ApplicationProperties, err
	}
	ApplicationProperties.InitTime = time.Now().Local().UTC()
	return ApplicationProperties, err
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
	if err != nil {
		return errors.Wrap(err, "GAP could not open the property file")
	}

	reader := bufio.NewReader(file)

	for {

		lineAsBytes, _, err := reader.ReadLine()

		if err != nil {
			if err != io.EOF {
				return errors.Wrap(err, "GAP could not read the line of a property file")
			}
			break
		}
		lineAsString := strings.Split(string(lineAsBytes), "=")

		if len(strings.TrimSpace(lineAsString[0])) > 0 && len(lineAsString) > 1 && !strings.HasPrefix(lineAsString[0], "#") {
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
