package model

import (
	"fmt"

	"github.com/zimmski/nethead"
)

type URL struct {
	ID  uint64
	URL string
}

var urlID uint64 = 1

var urls []*URL

func init() {
	Create(&URL{URL: "http://a"})
	Create(&URL{URL: "http://b"})
	Create(&URL{URL: "http://c"})
}

func validate(url *URL) []*nethead.Error {
	var errs []*nethead.Error

	if url.URL == "" {
		errs = append(errs, nethead.NewError(nethead.Required, fmt.Sprintf("%q is required", "URL")))
	}

	return errs
}

func Create(url *URL) (*URL, []*nethead.Error) {
	errs := validate(url)
	if len(errs) != 0 {
		return nil, errs
	}

	url.ID = urlID

	urlID++

	urls = append(urls, url)

	return url, nil
}

func CreateForm() map[string]interface{} {
	return map[string]interface{}{
		"URL": map[string]interface{}{
			"type":     "string",
			"required": true,
		},
	}
}

func All() []*URL {
	return urls
}

func Delete(id uint64) *nethead.Error {
	for i, url := range urls {
		if url.ID == id {
			urls = append(urls[:i], urls[i+1:]...)

			return nil
		}
	}

	return nethead.NewError(nethead.NotFound, fmt.Sprintf("URL %d not found", id))
}

func Edit(data *URL) (*URL, []*nethead.Error) {
	url, err := One(data.ID)
	if err != nil {
		return nil, []*nethead.Error{err}
	}

	errs := validate(data)
	if len(errs) != 0 {
		return nil, errs
	}

	url.URL = data.URL

	return url, nil
}

func EditForm() map[string]interface{} {
	return CreateForm()
}

func One(id uint64) (*URL, *nethead.Error) {
	for _, url := range urls {
		if url.ID == id {
			return url, nil
		}
	}

	return nil, nethead.NewError(nethead.NotFound, fmt.Sprintf("URL %d not found", id))
}
