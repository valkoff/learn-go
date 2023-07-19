package reflection

func walk(x interface{}, fn func(input string)) {
	fn("I have no idea what I'm doing")
}
