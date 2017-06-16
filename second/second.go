package second

import "fmt"

type TestType struct {
	TestString string
	TestInt    int
}
func Print() {
	fmt.Print("Let it rock a second time\n")
}

func (o *TestType) Print() {
	fmt.Println("My string: ",o.TestString)
	fmt.Println("My int:",o.TestInt)
}

func (o *TestType) Loop() {
	for i:=0; i<=10;i++  {
		fmt.Print(i)
	}
	fmt.Println()

	// rune = char :)
	array := [] string {"a","b","c"}

	for i,v := range array{
		fmt.Printf("%v=%v // ",i,v)
	}

	// We don't care about value
	for i,_  := range array{
		fmt.Printf("%v // ",i)
	}
	fmt.Println()
}