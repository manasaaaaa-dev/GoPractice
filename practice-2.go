// Section 3 — Practice Program 1 — Rectangle Area & ScaleRect (same concept, different struct)
// Section 4 — Practice Program 2 — Bank Account Deposit (real-world style example)

package main

import"fmt"

type Rectangle struct {
	L,B float64
}

func AreaPassByValue (r Rectangle, change float64)  float64 {
	r.L = r.L * change 
	r.B = r.B * change 
	return r.L*r.B
}

func AreaPassByAdd (r *Rectangle , change float64 ) float64 {
	r.L = r.L * change 
	r.B = r.B * change 
	return r.L*r.B
}

func main () {
	r := Rectangle { L:3, B:4}
	//here if we pass the change factor in the 
	fmt.Println("Before:", AreaPassByValue(r,2)) 
	fmt.Println("values of L and b are  — L:", r.L, "B:", r.B)
    AreaPassByAdd(&r, 2)
	fmt.Println("After:", AreaPassByValue(r,2) )
	
}