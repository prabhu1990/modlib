To use the modlib library, add an import in Go code and call it:

```
import "github.com/prabhu1990/modlib"

// ... later
s := modlib.Config()
```

To use the `modcmd` binary, do:

```
$ go get github.com/prabhu1990/modlib/cmd/modlib-client
$ modlib-client

$ go get github.com/prabhu1990/modlib/cmd/modlib-server
$ modlib-server
```
