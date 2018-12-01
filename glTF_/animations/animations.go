package animations

import (
	"./channels"
	"./samplers"
)

type T struct {
	Channels []channels.T
	Name     string
	Samplers []samplers.T
}

type T_channels = channels.T
type T_samplers = samplers.T
