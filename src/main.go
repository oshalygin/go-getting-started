package main

func main() {

	if foo := 2; foo == 4 {
		println("foobar")
	} else {
		println("more foobars")
	}

	baz := 2
	switch baz {
	case 1:
		{
			println("bazzing")
		}
	case 2:
		{
			println("lol baz")
		}
	}

	for iterator := 0; iterator < 5; iterator++ {
		println(iterator)
	}

	j := 0
	for {
		j++
		println(j)
		if j == 5 {
			break
		}
	}

	slice := []string{"foo", "bar", "baz"}

	for index, value := range slice {
		println(index, value)
	}

	dictionary := make(map[string]string)

	dictionary["first"] = "foo"
	dictionary["second"] = "bar"
	dictionary["third"] = "baz"

	for key, value := range dictionary {
		println(key, value)
	}
}
