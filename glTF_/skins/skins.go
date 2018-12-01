package skins

type T struct {
	InverseBindMatrices int
	Joints              map[int]int
	Skeleton            int
}

func (this *T) Init() {
	this.Joints = make(map[int]int)
}
