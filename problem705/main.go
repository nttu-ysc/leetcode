package problem705

type MyHashSet struct {
	el map[int]bool
}

func Constructor() MyHashSet {
	return MyHashSet{el: make(map[int]bool)}
}

func (hs *MyHashSet) Add(key int) {
	hs.el[key] = true
}

func (hs *MyHashSet) Remove(key int) {
	hs.el[key] = false
}

func (hs *MyHashSet) Contains(key int) bool {
	v, ok := hs.el[key]
	if !ok {
		return false
	}

	return v
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
