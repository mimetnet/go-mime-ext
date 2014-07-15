
# go-mime-ext

Add functions that compliment mime’s TypeByExtension, and AddExtensionType by
providing the other side of the coin. This allows us to convert Content-Type’s
into file extensions.

 - AddTypeExtension
 - ExtensionByType

## Installation

```
$ go get github.com/mimetnet/go-mime-ext
```

[![build status][1]][2]

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

// ctype = application/pdf
ctype := http.DetectContentType(head)

// fext = .pdf
fext := ExtensionByType(ctype)
```

# License

 MIT


  [1]: https://api.travis-ci.org/mimetnet/go-mime-ext.png
  [2]: https://travis-ci.org/mimetnet/go-mime-ext
