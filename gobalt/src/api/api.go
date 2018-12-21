package api

import (
	"gobalt/src/config"
	"gobalt/src/transport"
)

type API struct {
	Transport transport.Transport
	Session   transport.Session
	Opts      config.Opts
}

func New(c config.Opts) *API {
	api := &API{}
	api.Transport = *transport.New(c, &api.Session)
	api.Opts = c
	api.Session.Token = "badauth"
	return api
}
