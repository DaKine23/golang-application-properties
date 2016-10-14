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

// New initialize the properties usually only needs to be done once the property file changes or application starts
func New(newFilePath string) (*Properties, error) {
	ApplicationProperties.FilePath = newFilePath
	err := ApplicationProperties.readProps()
	if err != nil {
		return &ApplicationProperties, err
	}
	ApplicationProperties.InitTime = time.Now().Local().UTC()
	return &ApplicationProperties, err
}

func (ap *Properties) readProps() error {

	ap.PropertyMap = make(map[string]string)

	file, err := os.Open(ap.FilePath)
	if err != nil {
		return errors.Wrap(err, "GAP could not open the property file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {

		lineAsBytes, _, err := reader.ReadLine()

		if err != nil {
			if err != io.EOF {
				return errors.Wrap(err, "GAP could not read the line of a property file")
			}
			break
		}
		// ignore comments at the end of the line
		lineAsStringWithComments := strings.SplitN(string(lineAsBytes), "#", 2)

		// ignore all "=" after the first
		lineAsString := strings.SplitN(lineAsStringWithComments[0], "=", 2)

		if len(strings.TrimSpace(lineAsString[0])) > 0 && len(lineAsString) > 1 && !strings.HasPrefix(lineAsString[0], "#") {
			ap.PropertyMap[strings.TrimSpace(lineAsString[0])] = strings.TrimSpace(lineAsString[len(lineAsString)-1])
		}

	}

	return nil

}

//IsInitialized tells you if the property file was already loaded
func IsInitialized() bool {
	return (len(ApplicationProperties.PropertyMap) > 0)
}
