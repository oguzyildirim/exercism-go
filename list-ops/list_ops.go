// Package listops implement basic list operations.
package listops

// IntList collections of integers
type IntList []int

type binFunc func(x, y int) int
type predFunc func(x int) bool
type unaryFunc func(y int) int

// Foldr reduce each item into the accumulator from the left
func (il IntList) Foldr(fn binFunc, num int) int {
	for i := il.Length() - 1; i >= 0; i-- {
		num = fn(il[i], num)
	}
	return num
}

// Foldl reduce each item into the accumulator from the right
func (il IntList) Foldl(fn binFunc, num int) int {
	for _, i := range il {
		num = fn(num, i)
	}
	return num
}

// Length calculates lenght of integer list
func (il IntList) Length() int {
	var i int
	for range il {
		i++
	}
	return i
}

// Filter reduce the list of all items for which predicate(item) is True
func (il IntList) Filter(fn predFunc) IntList {
	list := make(IntList, 0)
	for _, i := range il {
		if fn(i) {
			list = append(list, i)
		}
	}
	return list
}

// Map maps the collection items with given
func (il IntList) Map(fn unaryFunc) IntList {
	result := make(IntList, 0)
	for _, val := range il {
		result = append(result, fn(val))
	}
	return result
}

// Append merge two lists together
func (il IntList) Append(list IntList) IntList {
	merged := make(IntList, il.Length()+list.Length())
	for i, num := range il {
		merged[i] = num
	}
	for i, num := range list {
		merged[il.Length()+i] = num
	}
	return merged
}

// Concat combine all items in all lists into one flattened list
func (il IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		il = il.Append(v)
	}
	return il
}

// Reverse returns a list with all the original items, but in reversed order
func (il IntList) Reverse() IntList {
	if il.Length() == 0 {
		return []int{}
	}
	result := make(IntList, 0)
	for i := il.Length(); i > 0; i-- {
		result = append(result, il[i-1])
	}
	return result
}
