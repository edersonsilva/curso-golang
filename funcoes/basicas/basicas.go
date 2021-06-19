package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func F4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func F5() (string, string) {
	return "Retorno 1", "Reotrno 2 "
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), F4("Param3", "Param4")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := F5()
	fmt.Println("F5:", r51, r52)

}
