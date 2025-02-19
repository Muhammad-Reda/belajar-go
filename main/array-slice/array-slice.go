package arrayslice

import "fmt"

func ArraySlice() {
	underlying := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice := underlying[2:5]

	fmt.Printf("%v", slice)
	fmt.Printf("\n%d", len(slice))
	fmt.Printf("\n%d", cap(slice))

	fmt.Println("\nAfter underlying changed")
	underlying[3] = 99
	slice = append(slice, 100)
	slice = append(slice, 100)
	slice = append(slice, 100)
	slice = append(slice, 100)
	slice = append(slice, 100) // make new underlying array, cause unsufficient cap of old underlying array
	slice = append(slice, 200)
	slice[7] = 300
	underlying[7] = 150
	fmt.Printf("\n%v", underlying)
	fmt.Printf("\n%v", slice)
}

func SliceMakeCopy() {
	sliceOri := make([]int, 3)
	sliceOri[0] = 1
	sliceOri[1] = 2
	sliceOri[2] = 3

	sliceCopy1 := make([]int, len(sliceOri))
	sliceCopyElement := copy(sliceCopy1, sliceOri)
	fmt.Printf("SiceCopy1: %v\n", sliceCopy1)
	fmt.Printf("SliceCopyElement: %v\n", sliceCopyElement)

	sliceCopy1[0] = 5
	sliceOri[1] = 5
	fmt.Printf("SliceOri: %v\n", sliceOri)
	fmt.Printf("sliceCopy1: %v\n", sliceCopy1)

	sliceCopy2 := make([]int, 2)
	sliceCopyElement2 := copy(sliceCopy2, sliceOri)
	fmt.Printf("SiceCopy1: %v\n", sliceCopy2)
	fmt.Printf("SliceCopyElement: %v\n", sliceCopyElement2)

	sliceCopy3 := make([]int, 5)
	sliceCopyElement3 := copy(sliceCopy3, sliceOri)
	fmt.Printf("SiceCopy1: %v\n", sliceCopy3)
	fmt.Printf("SliceCopyElement: %v\n", sliceCopyElement3)

}

func NewSlice() {
	// 1. Copy partial slice
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, 3)

	// Copy 3 elemen pertama
	copy(dest, source[:3])
	fmt.Printf("Copy partial: %v\n\n", dest)

	// 2. Copy dengan overlap
	numbers := []int{1, 2, 3, 4, 5}
	// Copy elemen dari index 2
	copy(numbers, numbers[2:])
	fmt.Printf("Copy dengan overlap: %v\n\n", numbers)

	// 3. Copy ke slice kosong
	emptyDest := make([]int, 0)
	copied := copy(emptyDest, source)
	fmt.Printf("Copy ke slice kosong: %v\n", emptyDest)
	fmt.Printf("Jumlah yang dicopy: %d\n\n", copied)

	// 4. Copy dari slice nil
	var nilSource []int
	normalDest := make([]int, 3)
	copiedFromNil := copy(normalDest, nilSource)
	fmt.Printf("Copy dari nil source: %v\n", normalDest)
	fmt.Printf("Jumlah yang dicopy: %d\n", copiedFromNil)
}
