package algo

//QuickFind ..
type QuickFind interface {
	Union(lhs, rhs int)
	Connected(lhs, rhs int) bool
}

type qf struct {
	Mapper []int
}

//New returns a newly instantiated QuickFind instance
func New(size int) QuickFind {
	var qf qf
	qf.Mapper = make([]int, size, size)

	for i := range qf.Mapper {
		qf.Mapper[i] = i
	}

	return &qf
}

func (qf *qf) Union(lhs, rhs int) {
	if qf.Connected(lhs, rhs) {
		return
	}

	key := qf.Mapper[lhs]
	oldKey := qf.Mapper[rhs]

	for i, v := range qf.Mapper {
		if v == oldKey {
			qf.Mapper[i] = key
		}
	}

}

func (qf *qf) Connected(lhs, rhs int) bool {
	return qf.Mapper[lhs] == qf.Mapper[rhs]
}
