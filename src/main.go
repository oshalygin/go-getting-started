package main

func main() {
	var myInt int
	myInt = 42

	println(myInt)

	var myFloat32 float32
	myFloat32 = 33.22

	println(myFloat32)

	myString := "Hello Go!"

	println(myString)

	myComplex := complex(2, 3)

	println(myComplex)
	println(real(myComplex), imag(myComplex))

}
