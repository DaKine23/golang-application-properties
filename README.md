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

USAGE (simple)

- place your file wherever you like most likely next to your go program
- e.g. **application.properties**
- add some values like **my.value = 1234** to the files

```golang
var p *golangapplicationproperties.Properties
// Init with filename and debugging false
ap = ap.Init("./application.properties", false)
//get properties as a map m [string] sting
m := ap.GetProps()
// should print "1234"
println(m["my.value"])
```

USAGE

```golang
var p *golangapplicationproperties.Properties
// get currently loaded properties (may somewhere else in the code)
ap := p.Current()
// check for "IsInitialized"
if !ap.IsInitialized() {
    //only init if not already done
    ap = ap.Init("./application.properties", false)
}
//get properties as a map m [string] sting
m := ap.GetProps()
// should print "1234"
println(m["my.value"])
```
