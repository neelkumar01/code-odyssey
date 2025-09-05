package maps
import "fmt"

// var m map[keyType]valueType
var m map[string]int					   // nil, no storage
var p = make(map[string]int)		// non-nil, empty

func Map() {
	a := m["myKey"]		// returns 0
	fmt.Println(a)
	fmt.Println(m["myKey"])         //  0

	m = p 	// both point to same hash table

	m["myKey"]++
	fmt.Println(m["myKey"])		// 1

	fmt.Println(p["myKey"])		// 1

	
}