package slicekit

import (
	"fmt"
	"strconv"
)

func ExampleIncludes() {
	var s = []string{"abc", "efg", "xyz"}

	fmt.Println(Includes(s, "abc"))
	fmt.Println(Includes(s, "opq"))

	// Output:
	// true
	// false
}

func ExampleIncludesBy() {
	type Carouse struct {
		Name  string
		Price int
	}
	var s = []Carouse{{Name: "python", Price: 99}}

	ok := IncludesBy(s, func(item Carouse) bool {
		return item.Name == "python"
	})
	fmt.Println(ok)
	ok = IncludesBy(s, func(item Carouse) bool {
		return item.Name == "go"
	})
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleIncludesSubSlice() {
	var s = []string{"abc", "123", "xyz", "189", "pll"}
	var sub = []string{"123", "189", "xyz"}

	ok := IncludesSubSlice(s, sub)

	fmt.Println(ok)

	ok = IncludesSubSlice(s, []string{"1", "2", "3"})

	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleEach() {
	slice := []int{1, 2, 3, 4, 5, 6}

	Each(slice, func(idx int, item int) {
		fmt.Println(item, idx)
	})

	// Output:
	//1 0
	//2 1
	//3 2
	//4 3
	//5 4
	//6 5
}

func ExampleChunk() {
	arr := []string{"a", "b", "c", "d", "e"}

	result1 := Chunk(arr, 1)
	result2 := Chunk(arr, 2)
	result3 := Chunk(arr, 3)
	result4 := Chunk(arr, 4)
	result5 := Chunk(arr, 5)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// [[a] [b] [c] [d] [e]]
	// [[a b] [c d] [e]]
	// [[a b c] [d e]]
	// [[a b c d] [e]]
	// [[a b c d e]]

}

func ExampleFilter() {

	type Carouse struct {
		Name  string
		Price int
	}

	var courseList = []Carouse{
		{Name: "go", Price: 99}, {Name: "python", Price: 50},
		{Name: "c#", Price: 89}, {Name: "javascript", Price: 29},
		{Name: "java", Price: 129}, {Name: "c/c++", Price: 159},
	}

	var needs = []string{"go", "python", "javascript"}

	list1 := Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		func(idx int, item int) bool {
			return item >= 3
		})
	list2 := Filter(courseList,
		func(idx int, item Carouse) bool {
			return item.Price > 50
		})

	list3 := Filter(courseList, func(idx int, item Carouse) bool {
		return Includes(needs, item.Name)
	})

	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)

	// Output:
	//[3 4 5 6 7 8 9]
	// [{go 99} {c# 89} {java 129} {c/c++ 159}]
	// [{go 99} {python 50} {javascript 29}]

}

func ExampleFilterMap() {
	s1 := []int{1, 2, 3}

	getEvenNumStr := func(i, num int) (string, bool) {
		if num%2 == 0 {
			return strconv.FormatInt(int64(num), 10), true
		}
		return "", false
	}

	result1 := FilterMap(s1, getEvenNumStr)

	fmt.Println(result1)

	// Output:
	// [2]
}

func ExampleConcat() {
	s1, s2 := []int{1, 2, 3}, []int{10, 12, 38}

	result1 := Concat(s1, s2)

	fmt.Println(result1)

	// Output:
	// [1 2 3 10 12 38]
}

func ExampleMap() {
	type Carouse struct {
		Name  string
		Price int
	}

	var courseList = []Carouse{
		{Name: "go", Price: 99}, {Name: "python", Price: 50},
		{Name: "c#", Price: 89}, {Name: "javascript", Price: 29},
		{Name: "java", Price: 129}, {Name: "c/c++", Price: 159},
	}

	list1 := Map(courseList, func(idx int, item Carouse) Carouse {
		item.Price = item.Price * 2

		return item
	})

	fmt.Println(list1)

	// Output:
	// [{go 198} {python 100} {c# 178} {javascript 58} {java 258} {c/c++ 318}]

}

func ExampleCompact() {

	var s1 = []int{0, 1, 2}
	var s2 = []string{"0", "1", ""}
	var s3 = []bool{false, true, true}

	r1 := Compact(s1)
	r2 := Compact(s2)
	r3 := Compact(s3)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)

	// Output:
	// [1 2]
	// [0 1]
	// [true true]
}

func ExampleDifference() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5}

	result1 := Difference(slice1, slice2)
	fmt.Println(result1)

	// Output:
	// [1 2 3]
}

func ExampleDifferenceWith() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	isDouble := func(v1, v2 int) bool {
		return v2 == 2*v1
	}

	result := DifferenceWith(slice1, slice2, isDouble)

	fmt.Println(result)

	// Output:
	// [1 5]
}

func ExampleEqual() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 3, 2}

	result1 := Equal(slice1, slice2)
	result2 := Equal(slice1, slice3)

	fmt.Println(result1)
	fmt.Println(result2)
}

func ExampleEqualWith() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{2, 4, 6}

	isDouble := func(a, b int) bool {
		return b == a*2
	}

	result1 := EqualWith(slice1, slice2, isDouble)

	fmt.Println(result1)

	// Output:
	// true
}

func ExampleEvery() {
	s1, s2 := []int{1, 2, 3, 4, 5, 6, 7}, []int{2, 4, 6, 8, 10}

	result1 := Every(s1, func(idx int, item int) bool {
		return item < 10
	})
	result2 := Every(s2, func(idx int, item int) bool {
		return item%2 == 0
	})
	result3 := Every(s2, func(idx int, item int) bool {
		return item+1 < 10
	})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	//true
	//true
	//false
}

func ExampleNone() {
	s1, s2 := []int{1, 2, 3, 4, 5, 6, 7}, []int{2, 4, 6, 8, 10}

	result1 := None(s1, func(idx int, item int) bool {
		return item == 1
	})

	result2 := None(s2, func(idx int, item int) bool {
		return item%2 == 0
	})

	result3 := None(s1, func(idx int, item int) bool {
		return item == 10
	})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output
	// false
	// false
	// true
}

func ExampleSome() {
	s1, s2 := []int{1, 2, 3, 4, 5, 6, 7}, []int{2, 4, 6, 8, 10}

	result1 := Some(s1, func(idx int, item int) bool {
		return item == 1
	})

	result2 := Some(s2, func(idx int, item int) bool {
		return item%2 == 0
	})

	result3 := Some(s1, func(idx int, item int) bool {
		return item == 10
	})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output
	// true
	// true
	// false
}

func ExampleCount() {
	nums := []int{1, 2, 3, 3, 4}

	result1 := Count(nums, 1)
	result2 := Count(nums, 3)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 1
	// 2
}

func ExampleCountBy() {
	nums := []int{1, 2, 3, 4, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result := CountBy(nums, isEven)

	fmt.Println(result)

	// Output:
	// 2
}

func ExampleFind() {
	nums := []int{1, 2, 3, 4, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result, ok := Find(nums, isEven)

	fmt.Println(*result)
	fmt.Println(ok)

	// Output:
	// 2
	// true
}

func ExampleFindLast() {
	nums := []int{1, 2, 3, 4, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result, ok := FindLast(nums, isEven)

	fmt.Println(*result)
	fmt.Println(ok)

	// Output:
	// 4
	// true
}

func ExampleFill() {
	list := Fill(5, 10)
	fmt.Println(list)

	// Output:
	// [5 5 5 5 5 5 5 5 5 5]
}
