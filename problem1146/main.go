package problem1146

type SnapshotArray struct {
	data    []map[int]int
	version int
}

func Constructor(length int) SnapshotArray {
	sa := SnapshotArray{
		data:    make([]map[int]int, length),
		version: 0,
	}

	for k := range sa.data {
		sa.data[k] = make(map[int]int, 0)
	}

	return sa
}

func (s *SnapshotArray) Set(index int, val int) {
	s.data[index][s.version] = val
}

func (s *SnapshotArray) Snap() int {
	s.version++
	return s.version - 1
}

func (s *SnapshotArray) Get(index int, snap_id int) int {
	for i := 0; i <= s.version; i++ {
		if v, ok := s.data[index][snap_id-i]; ok {
			return v
		}
	}

	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
