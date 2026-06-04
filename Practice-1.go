type meaning struct {
	word string

var m map[string] meaning

func main() {
	m = make(map[string]meaning)
	m["cat"] = meaning{
		"an animal belonging to the feline species",
	}
	fmt.Println(m["cat"])
}





