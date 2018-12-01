package scenes

type T struct {
	Name  string
	Nodes map[int]int
}

func (this *T) Init() {
	this.Nodes = make(map[int]int)
}
