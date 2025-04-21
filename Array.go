// Using VAR keyword we create array
// Syntax -  1.Var array_name[length]Type
//			 2. Var array_name[index]= element


package main

import "fmt"

func main() {
	arr := [4]string{"mirza", "jasmine", "chanchal", "roshii"}
	fmt.Println("Elemnts of Array")
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

}



package main

import "fmt"

func main() {
	arr := [3][3]string{{"chanchal", "roshii", "mehrunisha"}, {"gulaabjamub", "rasgulla", "kajukatli"}, {"Ninjahattori", "Sinchan", "Ben10"}}
	fmt.Println("Element of array")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(arr[x][y])
		}
	}
}

package main

import "fmt"

func main() {
	var myarr [2]int

	fmt.Println("Element of array:", myarr)
	//Bydefault it will be Element of array: [0 0]
}



package main

import "fmt"

func main() {
	arr1 := [3]int{7, 8, 6}
	arr2 := [...]int{0, 0, 0, 0, 7, 8, 6}
	arr3 := [3]int{7, 8, 6}

	fmt.Println("length of array1:", len(arr1))
	fmt.Println("length of array2:", len(arr2))
	fmt.Println("length of array3:", len(arr3))

}



package main

import "fmt"

func main() {
	myarr := [...]int{0, 0, 0, 0, 7, 8, 6}

	for x := 0; x < len(myarr); x++ {
		fmt.Printf("%d\n", myarr[x])
	}
}



package main

import "fmt"

func main() {

	my_array := [...]int{1, 0, 0, 0, 7, 8, 6}
	fmt.Println("Original array(Before):", my_array)
	new_array := my_array

	fmt.Println("New array(before):", new_array)
	new_array[0] = 0
	fmt.Println("New array(After):", new_array)

	fmt.Println("Original array(After):", my_array)
}



package main

import "fmt"

func main() {
	arr1 := [3]int{7, 8, 6}
	arr2 := [...]int{7, 8, 6}
	arr3 := [3]int{7, 8, 5}
	fmt.Println("length of array1:", len(arr1))
	fmt.Println("length of array2:", len(arr2))
	fmt.Println("length of array3:", len(arr3))
	fmt.Println(arr1 == arr2)
	fmt.Println(arr2 == arr3)
	fmt.Println(arr1 == arr3)

}
