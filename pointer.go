package main

func main(){
    a :=15

    mul(a)
	println("without pointer ",a)

	mul1(&a) //passing the address
	println("with pointer ",a)

}

func mul( v int){
	v=v*10
}

func mul1( v *int){
	*v = *v*10
}