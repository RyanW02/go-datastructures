package slicelist

type SliceList struct {
	slice []interface{}
}

func NewSliceList() SliceList {
	return SliceList{
		slice: make([]interface{}, 0),
	}
}

func (sl *SliceList) Add(value interface{}) {
	sl.slice = append(sl.slice, value)
}

func (sl *SliceList) AddAll(values []interface{}) {
	sl.slice = append(sl.slice, values...)
}

func (sl *SliceList) Get(index int) interface{} {
	return sl.slice[index]
}

func (sl *SliceList) Values() []interface{} {
	values := make([]interface{}, sl.Size())

	for i, value := range sl.slice {
		values[i] = value
	}

	return values
}

func (sl *SliceList) Clear() {
	sl.slice = make([]interface{}, 0)
}

func (sl *SliceList) Contains(value interface{}) bool {
	return sl.IndexOf(value) != -1
}

// returns -1 if list does not contain the element
func (sl *SliceList) IndexOf(value interface{}) int {
	for index, el := range sl.slice {
		if el == value {
			return index
		}
	}

	return -1
}

func (sl *SliceList) Size() int {
	return len(sl.slice)
}

func (sl *SliceList) IsEmpty() bool {
	return sl.Size() == 0
}

func (sl *SliceList) RemoveAt(index int) bool {
	if index >= len(sl.slice) {
		return false
	}

	sl.slice = append(sl.slice[:index], sl.slice[index+1:]...)
	return true
}

// n is the maximum amount of removals to max, or -1 for unlimited
func (sl *SliceList) Remove(value interface{}, n int) (removals int) {
	for i := 0; i < len(sl.slice) && n != removals; i++ {
		if sl.slice[i] == value {
			sl.RemoveAt(i)
			i--
			removals++
		}
	}

	return
}

func (sl *SliceList) Set(index int, value interface{}) {
	sl.slice[index] = value
}
