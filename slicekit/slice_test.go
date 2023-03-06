package slicekit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIncludes(t *testing.T) {
	var s = []string{"abc", "efg", "xyz"}

	fmt.Println(Includes(s, "abc"))
	fmt.Println(Includes(s, "opq"))
}

func TestIncludesBy(t *testing.T) {
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
}

func TestIncludesSubSlice(t *testing.T) {
	var s = []string{"abc", "123", "xyz", "189", "pll"}
	var sub = []string{"123", "189", "xyz"}

	ok := IncludesSubSlice(s, sub)

	fmt.Println(ok)

	ok = IncludesSubSlice(s, []string{"1", "2", "3"})

	fmt.Println(ok)
}

func TestFilter(t *testing.T) {

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

}

func TestMap(t *testing.T) {
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

	fmt.Println(courseList)

	fmt.Println(list1)

}

func TestChunk(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}

	result1 := Chunk(arr, 1)
	result2 := Chunk(arr, 2)
	result3 := Chunk(arr, 3)
	result4 := Chunk(arr, 4)
	result5 := Chunk(arr, 5)
	result6 := Chunk(arr, 6)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
}

func TestCompact(t *testing.T) {

	var s1 = []int{0, 1, 2}
	var s2 = []string{"0", "1", ""}
	var s3 = []bool{false, true, true}

	r1 := Compact(s1)
	r2 := Compact(s2)
	r3 := Compact(s3)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)

	var list = [3]int{1, 2, 3}

	typeOfList := reflect.TypeOf(list)
	fmt.Println(typeOfList.Kind())

	typeOfS1 := reflect.TypeOf(s1)
	fmt.Println(typeOfS1.Kind())
}

func TestDifferenceWith(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	isDouble := func(v1, v2 int) bool {
		return v2 == 2*v1
	}

	result := DifferenceWith(slice1, slice2, isDouble)

	fmt.Println(result)
}

func TestEqual(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 3, 2}

	result1 := Equal(slice1, slice2)
	result2 := Equal(slice1, slice3)

	fmt.Println(result1)
	fmt.Println(result2)
}

func TestEqualWith(t *testing.T) {
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
