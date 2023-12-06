package day05

type seed struct {
	seed int
	data []int
}

func newSeed(id int) *seed {
	return &seed{
		seed: id,
		data: make([]int, 7),
	}
}
