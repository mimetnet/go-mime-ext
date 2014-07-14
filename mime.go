package mimeext

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var extLock sync.RWMutex

var extensions = map[string]string{
	"application/octet-stream": "",
	"application/pdf":          ".pdf",
	"application/x-gif":        ".gif",
	"application/x-javascript": ".js",
	"application/x-png":        ".png",
	"application/xml":          ".xml",
	"image/gif":                ".gif",
	"image/jpeg":               ".jpg",
	"image/png":                ".png",
	"image/tiff":               ".tif",
	"image/tiff-fx":            ".tif",
	"text/css":                 ".css",
	"text/html":                ".htm",
	"text/plain":               ".txt",
	"text/xml":                 ".xml",
}

func strip(str string) string {
	return strings.Trim(strings.Split(str, ";")[0], " ")
}

func ExtensionByType(ctype string) string {
	ctype = strip(ctype)

	extLock.RLock()
	ext := extensions[ctype]
	extLock.RUnlock()
	return ext
}

func AddTypeExtension(contentType string, ext string) error {
	if contentType == "" {
		return errors.New("contentType cannot be ''")
	} else if ext == "" || ext[0] != '.' {
		return fmt.Errorf(`ext "%s" missing dot`, ext)
	} else if 1 == len(ext) {
		return fmt.Errorf(`ext "%s" must start with a dot and be followed by one or more chars`, ext)
	}

	contentType = strip(contentType)

	extLock.Lock()
	extensions[contentType] = ext
	extLock.Unlock()

	return nil
}
