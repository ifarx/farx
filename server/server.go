package server

import (
	"src/event"
	"src/service"
)

type Server interface {
	service.Service
	event.EventTarget
}
