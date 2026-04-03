package urltool

import (
	"net/url"
	"path"
)

func GetBasePath(targetUrl string) (string, error) {
	myUrl, error := url.Parse(targetUrl)
	if error != nil {
		return "", error
	}
	basePath := path.Base(myUrl.Path)
	return basePath, nil
}
