package main
  
import "fmt"
func main(){
	for i, v:=range "Hey World!"{
		fmt.Printf("the index of %c is %d, it's bits value is %v",v,i,v);
	}
}