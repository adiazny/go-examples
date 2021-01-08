package main

import (
	"fmt"
	"sort"
)

type measures struct {
	Value int
	Type  int
	Unit  int
}

type measureGrps struct {
	Grpid    int
	Created  int64
	Category int
	Measures []measures
}

// Generate dummy structs to work with in the example
func createTestStructs() (m1, m3 measures, m2, m4 *measures) {
	// m1 and m3 has struct type measures
	m1.Type = 1
	m1.Value = 10110
	m1.Unit = -3

	m3 = measures{42355, 1, -4}

	// m2 and m4 are pointers to struct type measures
	m2 = new(measures)
	m2.Type = 1
	m2.Value = 12310
	m2.Unit = -2

	m4 = &measures{98355, 1, -5}

	// fmt.Printf("Measures Struct 1: %+v\n", m1)
	// fmt.Printf("Measures Struct 2: %+v\n", m2)
	// fmt.Printf("Measures Struct 3: %+v\n", m3)
	// fmt.Printf("Measures Struct 4: %+v\n", m4)
	return
}

func main() {

	m1, m3, m2, m4 := createTestStructs()

	m1Slice := make([]measures, 1)
	m1Slice[0] = m1

	m2Slice := make([]measures, 1)
	m2Slice[0] = *m2

	m3Slice := make([]measures, 1)
	m3Slice[0] = m3

	m4Slice := make([]measures, 1)
	m4Slice[0] = *m4

	mgrps1 := measureGrps{1, 1609451681, 1, m1Slice}
	mgrps2 := measureGrps{1, 1609441681, 1, m2Slice}
	mgrps3 := measureGrps{1, 1609431681, 1, m3Slice}
	mgrps4 := measureGrps{1, 1609421681, 1, m4Slice}

	// Make unsorted slice of MeasureGrps structs
	mgrpsSlice := make([]measureGrps, 4)
	mgrpsSlice[0] = mgrps4
	mgrpsSlice[1] = mgrps3
	mgrpsSlice[2] = mgrps2
	mgrpsSlice[3] = mgrps1

	fmt.Println("Unsorted Slice:", mgrpsSlice)
	fmt.Printf("%#v\n", mgrpsSlice)

	// Sort the slice by decending created fields
	// Also an example of a closure (anonymous function)
	sort.Slice(mgrpsSlice, func(i, j int) bool {
		return mgrpsSlice[i].Created > mgrpsSlice[j].Created
	})

	fmt.Println("Sorted Slice by Created Timestamp:", mgrpsSlice)

}
