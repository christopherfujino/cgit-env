package service

import ()

type T interface {
	Register()
	Describe() Description
}

type Description struct {
	Endpoint string
	Text string
}
