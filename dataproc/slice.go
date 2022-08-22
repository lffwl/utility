package dataproc

import "192.168.0.209/wl/utility/define"

// SliceUnion 并集
func SliceUnion[T string](slice1, slice2 []T) []T {
	m := make(map[T]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// SliceIntersect 交集
func SliceIntersect[T string](slice1, slice2 []T) []T {
	m := make(map[T]int)
	nn := make([]T, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// SliceDifference 差集
func SliceDifference[T string](slice1, slice2 []T) []T {
	m := make(map[T]int)
	nn := make([]T, 0)
	inter := SliceIntersect[T](slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// Distinct 去重
func Distinct[T string](slc []T) []T {
	var result []T
	tempMap := map[T]define.Void{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = define.Void{}
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
