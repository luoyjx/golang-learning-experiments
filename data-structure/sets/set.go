package go_sets

type StringSet struct {
	m map[string]struct{}
}

func (ss *StringSet) Add(s string) bool {
	if _, exists := ss.m[s]; exists {
		return false
	}

	ss.m[s] = struct{}{}

	return true
}

func (ss *StringSet) Contains(s string) bool {
	_, exists := ss.m[s]
	return exists
}

func (ss *StringSet) Clear() {
	ss.m = make(map[string]struct{})
}

func (ss *StringSet) Remove(s string) bool {
	if _, exists := ss.m[s]; exists {
		delete(ss.m, s)
		return true
	}

	return false
}




