package main


func main(){

	//Array 
	 Array := [10]int{}

	for  i := 0 ; i<10 ; i++ {
		Array[i] = i*10
	}

	println("printing Array values")
	for i := range Array{
		println(Array[i])
	}

	//Slice
	Slice := []int{}

	for  i := 0 ; i<10 ; i++ {
		Slice = append(Slice, i*10)
	}

    println("Printing Slice values")
    for _,v := range Array{
		println(v)
	}
}