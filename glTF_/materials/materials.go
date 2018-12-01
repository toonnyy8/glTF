package materials

type T struct {
	Name                 string
	PbrMetallicRoughness struct {
		BaseColorFactor [4]float64
		MetallicFactor  float64
		RoughnessFactor float64
	}
}
