package algo



//QuickUnion ..
type QuickUnion interface {
	QuickFind
	Root(val int) int
}

type qu struct {
	Mapper []int
}

//NewQU ..
func NewQU(size int) QuickUnion {
	var qu qu
	qu.Mapper = make([]int, size, size)

	for i := range qu.Mapper {
		qu.Mapper[i] = i
	}

	return &qu
}

func (qu *qu) Union(lhs, rhs int) {
	qu.Mapper[qu.Root(lhs)] = qu.Root(rhs)
}

func (qu *qu) Connected(lhs, rhs int) bool {
	return qu.Root(lhs) == qu.Root(rhs)
}

//No error handling, should probably do something with out of range calls
func (qu *qu) Root(index int) int {
	result := index
	for i := 0; i < len(qu.Mapper); i++ {
		if i == qu.Mapper[i] {
			result = i
			break
		}
	}
	return result
}
