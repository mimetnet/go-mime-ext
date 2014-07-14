
# go-mime-ext

Add functions that golang’s mime should have to compliment TypeByExtension, and AddExtensionType.

 - AddTypeExtension
 - ExtensionByType

## Installation

```
$ go get github.com/mimetnet/go-mime-ext
```

## Example

```go
import (
    "fmt"
    "net/http"
    "os"

    . "github.com/mimetnet/go-mime-ext"
)

head := make([]byte, 512)
read, _ := os.Open(filename)

defer read.Close()

read.Read(head)
read.Seek(0, 0)

ctype := http.DetectContentType(head)

// if ctype = application/pdf
// fext = .pdf
fext := ExtensionByType(ctype)
```

# License

 MIT
