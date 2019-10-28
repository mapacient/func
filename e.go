package e

import (
	"container/heap"
//	"fmt"
)

func e(sorted [][]int, j int) int {
	n := len(sorted)
	k := 0
	if n == 0 {
		return invalid
	}
	total := 0
	min := make([]*Array, 0, n)
	max := make([]*Array, 0, n)
	for i := 0; i < n; i++ {
		if len(sorted[i]) > 0 {
			a := &Array{sl: sorted[i],
				cur:       -1,
				l:         -1,
				r:         len(sorted[i]),
				index_min: i,
				index_max: i,
			}
			min = append(min, a)
			max = append(max, a)
			total += len(sorted[i])
		}
	}
	if j > total || j < 1 {
		return invalid
	}
	min_h := Min_A(min)
	max_h := Max_A(max)
	heap.Init(&min_h)
	heap.Init(&max_h)
	min_right, max_left := 0, 0
	max_array_left, min_array_right := max_h[0], min_h[0]
	cur_array := min_h[0]
	for limit := 0; limit < 59; limit++ {
		for i:=0;i<len(min_h);i++{
			//why:=min_h[i]
			//fmt.Print("why string is ")

			//fmt.Print(why.sl,why.Value()," k is ",k,"\n")
		}
		//fmt.Print("\n\n")
		if k == j && min_right >= max_left {
			if max_array_left.cur==-1{
				return min_right
			}

			return max_left
		}
		//I need move right
		//I think k == j but max_left>min_right
		if k < j || (k == j && max_left > min_right) {
			//max left array need move right
			if max_left < min_right {
				cur_array = max_array_left
			} else {
				//min right array need move right
				cur_array = min_array_right
			}
			//fmt.Print(cur_array.Value(), " k is :", k, " is still less than need\n")
			if cur_array.cur >= len(cur_array.sl)-1 {
				//fmt.Print("rrrrmoved")
				if cur_array.cur == len(cur_array.sl)-1 {
					//k++
				}
				heap.Remove(&min_h, cur_array.index_min)
				heap.Remove(&max_h, cur_array.index_max)
				max_array_left = max_h[0]
				max_left = max_array_left.Value()
				min_array_right = min_h[0]
				min_right = min_array_right.Value1()
				continue
			}
			c := cur_array.cur
			cur_array.Right()
			k = k + cur_array.cur - c
			heap.Fix(&max_h, cur_array.index_max)
			heap.Fix(&min_h, cur_array.index_min)
			max_array_left = max_h[0]
			max_left = max_array_left.Value()
			min_array_right = min_h[0]
			min_right = min_array_right.Value1()
		} else if k > j {
			//max left need move left
				cur_array = max_array_left
			if cur_array.cur <= 0 {
				heap.Remove(&max_h, cur_array.index_max)
				heap.Remove(&min_h, cur_array.index_min)
				if cur_array.cur == 0 {
					k--
				}
				max_array_left = max_h[0]
				max_left = max_array_left.Value()
				min_array_right = min_h[0]
				min_right = min_array_right.Value1()
				continue
			}
			c := cur_array.cur
			cur_array.Left()
			k = k - (c - cur_array.cur)
			heap.Fix(&max_h, cur_array.index_max)
			heap.Fix(&min_h, cur_array.index_min)
			max_array_left = max_h[0]
			max_left = max_array_left.Value()
			min_array_right = min_h[0]
			min_right = min_array_right.Value1()
		}
	}
	return 0
}
func e1() {
	sl := make([]int, 8)
	for i := 0; i < 7; i++ {
		sl[i] = i
	}
	a := &Array{sl: sl, cur: -1, l: 0, r: len(sl) - 1}
	b := &Array{sl: sl, cur: -1, l: 0, r: len(sl) - 1}
	a.Right()
	b.Left()
}
func (a *Max_A) Len() int {
	return len(*a)
}
func (a *Max_A) Less(i, j int) bool {
	return (*a)[i].Value() > (*a)[j].Value()
}
func (a *Max_A) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	(*a)[i].index_max = i
	(*a)[j].index_max = j
}
func (a *Max_A) Push(itf interface{}) {
	arr := itf.(*Array)
	*a = append(*a, arr)
}
func (a *Max_A) Pop() interface{} {
	n := len(*a)
	if n == 0 {
		return nil
	}
	r := (*a)[n-1]
	(*a) = (*a)[:n-1]
	return r
}
func (a *Min_A) Push(itf interface{}) {
	arr := itf.(*Array)
	*a = append(*a, arr)
}
func (a *Min_A) Pop() interface{} {
	n := len(*a)
	if n == 0 {
		return nil
	}
	r := (*a)[n-1]
	(*a) = (*a)[:n-1]
	return r
}
func (a *Min_A) Len() int {
	return len(*a)
}
func (a *Min_A) Less(i, j int) bool {
	return (*a)[i].Value1() < (*a)[j].Value1()
}
func (a *Min_A) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	(*a)[i].index_min = i
	(*a)[j].index_min = j
}
func (a *Array) Index(i int) int {
	if i < 0 {
		return -1 << 63
	}
	if i >= len(a.sl) {
		return (1<<63) - 1
	}
	return a.sl[i]
}
func (a *Array) Value1() int {
	return a.Index(a.cur + 1)
}
func (a *Array) Value() int {
	return a.Index(a.cur)
}

type Max_A []*Array
type Min_A []*Array

func (a *Array) Right() {
	d := (a.r - a.cur + 1) >> 1
	if a.cur!=-1{
		a.l = a.cur
	}
	a.cur = a.cur + d
	//fmt.Print("now at ", a.Value(), "\n")
}
func (a *Array) Left() {
	d := (a.cur - a.l + 1) >> 1
	if a.cur!=-1{
		a.r = a.cur
	}
	a.cur = a.cur - d
	//fmt.Print("now at ", a.Value(), "\n")
}

const (
	invalid = -1
)

type Array struct {
	sl                   []int
	cur                  int
	l, r                 int
	index_max, index_min int
}
