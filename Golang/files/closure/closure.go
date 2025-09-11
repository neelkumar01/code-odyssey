package closure

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}


/*
- counter func has local var "count"
- if only counter func there,
	- it is called
	- after execution, count variable destroys
- with closure func inside
	- count gets saved with it !
	- when closure func executes, count then destroys
*/