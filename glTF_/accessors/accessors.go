package accessors

type T struct {
	BufferView    int
	ByteOffset    int
	ComponentType int
	Count         int
	Max           []float64
	Min           []float64
	Type          string
	Sparse        struct {
		Count   int
		Indices struct {
			BufferView    int
			ByteOffset    int
			ComponentType int
		}
		Values struct {
			BufferView int
			ByteOffset int
		}
	}
}
