package main

/*
#include <stdio.h>
int Num(void){
	return 555;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("start")
	fmt.Println(C.Num())
}
