package url

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zimmski/nethead/response"

	"github.com/zimmski/urlshot/model"
)

type URL struct {
}

func New() *URL {
	return &URL{}
}

func (c *URL) UID() string {
	return "url"
}

func (c *URL) All(req *http.Request) response.Responder {
	data := model.All()

	return response.NewJSON(data)
}

func (c *URL) Create(req *http.Request) response.Responder {
	decoder := json.NewDecoder(req.Body)

	var data model.URL

	err := decoder.Decode(&data)
	if err != nil {
		return response.NewError(err)
	}

	url, errs := model.Create(&data)
	if len(errs) != 0 {
		return response.NewErrors(errs)
	}

	return response.NewJSON(url)
}

func (c *URL) CreateForm(req *http.Request) response.Responder {
	data := model.CreateForm()

	return response.NewJSON(data)
}

func (c *URL) Delete(req *http.Request) response.Responder {
	urlID, _ := strconv.ParseUint(mux.Vars(req)["urlID"], 10, 64)

	err := model.Delete(urlID)
	if err != nil {
		return response.NewError(err)
	}

	return response.NewJSON(nil)
}

func (c *URL) Edit(req *http.Request) response.Responder {
	urlID, _ := strconv.ParseUint(mux.Vars(req)["urlID"], 10, 64)

	decoder := json.NewDecoder(req.Body)

	var data model.URL

	err := decoder.Decode(&data)
	if err != nil {
		return response.NewError(err)
	}

	data.ID = urlID

	url, errs := model.Edit(&data)
	if len(errs) != 0 {
		return response.NewErrors(errs)
	}

	return response.NewJSON(url)
}

func (c *URL) EditForm(req *http.Request) response.Responder {
	data := model.EditForm()

	return response.NewJSON(data)
}

func (c *URL) One(req *http.Request) response.Responder {
	urlID, _ := strconv.ParseUint(mux.Vars(req)["urlID"], 10, 64)

	data, err := model.One(urlID)
	if err != nil {
		return response.NewError(err)
	}

	return response.NewJSON(data)
}
