package nodes

type T struct {
	Mesh        int
	Children    map[int]int
	Name        string
	Rotation    [3]float64
	Scale       [3]float64
	Skin        int
	Translation [3]float64
}

func (this *T) Init() {
	this.Children = make(map[int]int)
}
