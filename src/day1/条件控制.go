package main

func main() {
	tiaojian()
}
func tiaojian()  {
	a := "b"
	switch {
	case true:
		a = "A"
	case false:
		a = "B"

	}
	switch a {
	case "A":
		print("adf")
		fallthrough
	case "B":
		print("asdf")
		fallthrough
	default:
		print("哈哈")
	}


}


