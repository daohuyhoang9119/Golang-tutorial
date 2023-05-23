package main

import "fmt"

func main() {
	// 	fmt.Println("Hello, World!")
	// 	fmt.Println("go" + "lang")
	// 	fmt.Println(1 + 1)
	// 	fmt.Println(true && false)
	// 	fmt.Println(true || false)
	// 	fmt.Println(!true)
	// 	fmt.Println("----------");
	// //	Variables
	// 	var a = "hello my name is Ryan"
	// 	fmt.Println(a)

	// 	var b = 0
	// 	fmt.Println(b)

	// 	var d = true
	//     fmt.Println(d)
	// 	var e int
	//     fmt.Println(e)

	// 	f := "apple"

	//     fmt.Println(f)
	// 	g:=""
	// 	fmt.Println(g)

	// Constant
	// const n = 0.314
	// var d = 4
	// fmt.Println(d);
	// var area float64 = float64(d)*n;
	// fmt.Println("The are of a round is:" + " ",area)

	//FOR LOOP
	// for i<=5{
	// 	fmt.Println("count number ", i)
	// 	i++
	// }

	// for j := 7; j <= 9; j++ {
	//     fmt.Println("this is a second for loop: ",j)
	// }
	// fmt.Println("this is a second for loop: ")
	// for n := 0; n <= 5; n++ {
    //     if n%2 == 0 {
    //         continue
    //     }
    //     fmt.Println(n)
    // }
	// if 7%2 == 0 {
    //     fmt.Println("7 is even")
    // } else {
    //     fmt.Println("7 is odd")
    // }

	// whatAmI :=func (i interface{})  {
	// 	switch t := i.(type) {
	// 	case bool:
    //         fmt.Println("I'm a bool")
    //     case int:
    //         fmt.Println("I'm an int")
    //     default:
    //         fmt.Printf("Don't know type %T\n", t)
    // 	}
	// }
	// whatAmI(true)
    // whatAmI(1)

	//----------ARRAY--------------
	var a [5]int
	fmt.Println("first empty arr:",a)
	a[4] = 23
	fmt.Println("after add value to arr:",a)
	fmt.Println(len(a))
	
	b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("array b:", b)
	
	//----------SLICE--------------
	scientists := []string{
        "Einstein",
        "Turing",
        "Lovelace",
        "Curie",
        "Franklin",
        "Hodgkin",
    }

    scientists = append(scientists, "Hawkins")

    fmt.Println(scientists)
}