package main
import "fmt"

func main() {

	//length means how many elements are present in the slice
	//capacity means the elements that the slice can accomodate 
	s:=[]int{2, 3, 5, 7, 11, 13}
	print(s)
	//len=6 cap=6 [2 3 5 7 11 13]


	s = s[:0]
	print(s)
	//len=0 cap=6 []
	s = s[:4]
	print(s)
	//len=4 cap=6 [2 3 5 7]

	s = s[2:]
	print(s)
	//len=2 cap=4 [5 7]

	func print(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
}