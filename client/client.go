package client

import (
	"src/event"
	"src/service"
)

type Client interface {
	service.Service
	event.EventTarget
}
