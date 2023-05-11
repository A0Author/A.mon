package main
import("fmt";"os/exec")
func main(){
	fmt.Println("I AM ALIVE!")
	exec.Command("I AM ALIVE").Start()
}