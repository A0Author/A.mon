/////IMPORT/////
package main
import("fmt";"os";"path/filepath";"io/ioutil";"math/rand";"time")
/////MAIN/////
func main(){
/////ACCESSING DNA/////
	NOWDIRECTORY,_:=os.Executable() 
	os.Chdir(filepath.Dir(NOWDIRECTORY))
        ANCESTORBODY,_:=ioutil.ReadFile("A.mon.exe")
	fmt.Println("ACTIVATING GENETIC INJECTOR!")
	fmt.Println("ANCESTOR BODY SIZE:\n",len(ANCESTORBODY))

	rand.Seed(time.Now().UnixNano())

	for X := 1; X <= 100; X++ {
		ANCESTORBODY=append(ANCESTORBODY,[]byte{byte(rand.Intn(255))}...)
		}


        ioutil.WriteFile("A.mon.exe", []byte(ANCESTORBODY), 0777)

	fmt.Println("GENETIC INJECTION FINISHED!")
	fmt.Println("ANCESTOR BODY SIZE:\n",len(ANCESTORBODY))
        fmt.Scanln()
}