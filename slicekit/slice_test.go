package slicekit

import (
	"fmt"
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
