#Go Notes

##Overview

Go is designed for public repo integration (ex Github).  Point $GOPATH to your projects

###Commands

For a project in /src/&lt;repo&gt;/&lt;user&gt;/my_project

* Compile/install: `$ go install`
* run with `$ $GOPATH/bin/my_project`
* Build libraries `$ go build`

Import library with:

```
import (
	…
	"<repo>/<user>/my_library"
)
```

##Other Notes & Questions

* `[]byte` - byte slice?