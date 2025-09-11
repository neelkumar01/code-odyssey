package structs
import "fmt"

type myCollection struct{
	count int
}

func Give(x myCollection){
	x.count--
}
func Sold(x *myCollection) {
	x.count--
}

func PassingStruct() {
	var card = myCollection{
		count: 24,
	}
	fmt.Println(card.count)
	Sold(&card)
	fmt.Println(card.count)
}
/*
Give() only takes struct value but don't modify original

Sold() takes struct pointer and modifies original struct
*/