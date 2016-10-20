# golangapplicationproperties
Provides easy access to Java style property files

GET

```bash
go get "github.com/DaKine23/golangapplicationproperties"
```
IMPORT

```golang
import "github.com/DaKine23/golangapplicationproperties"
```



USAGE

- place your file wherever you like most likely next to your go program
- e.g. **application.properties**
- add some values like **my.value = 1234** to the files

```golang
	if !golangapplicationproperties.IsInitialized() {
		golangapplicationproperties.NewProperties("application.properties")
	}

	_, err1 := fmt.Println(golangapplicationproperties.ApplicationProperties.PropertyMap["my.value"])

	if err1 != nil {
		panic(err1) //handle error
	}
    
    // or use it like that


	p, err2 := golangapplicationproperties.NewProperties("application.properties")

	if err2 != nil {
		panic(err2) //handle error
	}

	fmt.Println(p.PropertyMap["my.value"])


```
