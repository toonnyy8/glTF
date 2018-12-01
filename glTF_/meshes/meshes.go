package meshes

import "./primitives"

type T struct {
	Name       string
	Primitives []primitives.T
}

func (this *T) Init() {
}
