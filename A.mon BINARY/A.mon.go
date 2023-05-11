/////IMPORT/////
package main
import("time";"os";"os/exec";"path/filepath";"io/ioutil";"math/rand";"strings";"fmt")
/////MAIN/////
func main(){
	///fmt.Println("I AM ALIVE!")
/////HATCH/////
	DNALOCATION:=2465280
	HABITATAREA:="HABITAT"
	MUTATIONRARITY:=10
	REPRODUCTIONRATE:=70

	///fmt.Println("HATCHED!")
/////ACCESSING DNA/////
	///fmt.Println("SENSING ENVIROMENT!")
	///fmt.Println("SENSING CURRENT DIRECTORY!")
	NOWDIRECTORY,_:=os.Executable() 
	os.Chdir(filepath.Dir(NOWDIRECTORY))
        ANCESTORBODY,_:=ioutil.ReadFile("A.mon.exe")
        ///fmt.Println("ANCESTOR BODY SIZE:",len(ANCESTORBODY))
	///fmt.Println("ACQUIRING ANCESTOR DNA!")	
	ANCESTORDNA:=ANCESTORBODY[DNALOCATION:]
	///fmt.Println("ANCESTOR DNA SIZE:\n",len(ANCESTORDNA))

/////ACTION/////
	///fmt.Println("ACTS UPON THE WORLD!")

	ACTIONEXECUTABLE:="ACTION.exe"
	ioutil.WriteFile(ACTIONEXECUTABLE, []byte(ANCESTORDNA), 0777)
	exec.Command("rundll32.exe", "url.dll,FileProtocolHandler",filepath.Dir(NOWDIRECTORY)+"\\"+ACTIONEXECUTABLE).Start()
	///fmt.Println("IT IS DONE!")

	fmt.Println(string(ANCESTORDNA))
/////SENSORY//////
	time.Sleep(time.Duration(REPRODUCTIONRATE)*time.Second)
	///fmt.Println("RECOGNIZING NEARBY HABITATS!")
	NOWHABITAT,_:=os.Getwd()
	UPHABITAT:=filepath.Dir(NOWHABITAT)
	FILELIST,_:=ioutil.ReadDir(NOWHABITAT)
	MAYBEHABITATS:=make([]string,len(FILELIST))
	MAYBEHABITATS[0]=UPHABITAT
	HABITATINDEX:=1
	for _,FILE:=range FILELIST{
		if FILE.IsDir(){
			MAYBEHABITATS[HABITATINDEX]=NOWHABITAT+"\\"+FILE.Name()
			HABITATINDEX=HABITATINDEX+1
		}
	}
	HABITATS:=make([]string,HABITATINDEX)
	copy(HABITATS,MAYBEHABITATS)
	rand.Shuffle(len(HABITATS),func(X,Y int){HABITATS[X],HABITATS[Y]=HABITATS[Y],HABITATS[X]})
	///for HABITATNUMBER,HABITATNAME:=range HABITATS{
		///fmt.Println("HABITAT",HABITATNUMBER,HABITATNAME)
	///}
/////MUTATION/////
	for HABITATINDEX,_:=range HABITATS{
		time.Sleep(time.Duration(1)*time.Second)
		if !strings.Contains(HABITATS[HABITATINDEX], HABITATAREA[:]){
			continue
		}
		///fmt.Println("STARTING MUTATION!")
		rand.Seed(time.Now().UnixNano())
		///fmt.Println("MUTATING OFFSPRING DNA!")
		OFFSPRINGDNA:=ANCESTORDNA
/////ALTERATION/////
		for ALT,_:=range ANCESTORDNA{
			if rand.Intn(len(ANCESTORDNA)*MUTATIONRARITY)==0{ 
				///fmt.Println("ALTERING GENE NUMBER:",ALT)	
				OFFSPRINGDNA[ALT]=byte(rand.Intn(255))
			}		
		}				
/////DELETION/////		
		D:=0
		for DEL,_:=range OFFSPRINGDNA{
			if rand.Intn(len(ANCESTORDNA)*MUTATIONRARITY)==0{
				///fmt.Println("DELETING GENE NUMBER:",DEL)
				OFFSPRINGDNA=append(OFFSPRINGDNA[:DEL-D],OFFSPRINGDNA[(DEL-D)+1:]...)							
				D=D+1				
			}	
		}
		D=0
/////CREATION/////
		C:=0
		for CRE,_:=range OFFSPRINGDNA{
			if rand.Intn(len(ANCESTORDNA)*MUTATIONRARITY)==0{

				CRE=CRE+C
			        PARTONE:=make([]byte,len(OFFSPRINGDNA[:CRE]))
				copy(PARTONE,OFFSPRINGDNA[:CRE])
			        PARTTWO:=make([]byte,len(OFFSPRINGDNA[CRE:]))
				copy(PARTTWO,OFFSPRINGDNA[CRE:])
				NEWGENE:=[]byte{byte(rand.Intn(255))}

				ONEANDGENE:=append(PARTONE,NEWGENE...)
				OFFSPRINGDNA=append(ONEANDGENE,PARTTWO...)

				CRE=CRE+1

				///fmt.Println("CREATING NEW GENE BEFORE GENE NUMBER:",CRE-C-1)		
				C=C+1	
			}				
		}
		C=0
		if rand.Intn(len(ANCESTORDNA)*MUTATIONRARITY)==0{
			OFFSPRINGDNA=append(OFFSPRINGDNA,[]byte{byte(rand.Intn(255))}...)
			///fmt.Println("CREATING NEW GENE IN THE END OF OFFSPRING'S DNA")
		}
                ///fmt.Println("OFFSPRING DNA SIZE:\n",len(OFFSPRINGDNA))
		fmt.Println(string(OFFSPRINGDNA))
		OFFSPRINGDNA=append(ANCESTORBODY[:DNALOCATION],[]byte(OFFSPRINGDNA)...)
		fmt.Println(len(OFFSPRINGDNA))
		///fmt.Println("CONNECTING OFFSPRING'S DNA AND BODY!")
/////REPRODUCTION/////
		///fmt.Println("STARTING REPRODUCTION!")
	        ///fmt.Println("SELECTING RANDOM OFFSPRING HABITAT!")
		SELECTEDHABITAT:=HABITATS[HABITATINDEX]
	        ///fmt.Println("SELECTED HABITAT:",HABITATINDEX,SELECTEDHABITAT)
		ERROR:=ioutil.WriteFile(SELECTEDHABITAT+"\\"+"A.mon.exe", []byte(OFFSPRINGDNA), 0777)
		if ERROR==nil{
			///fmt.Println("OFFSPRING BORN IN SELECTED HABITAT!")

			exec.Command("rundll32.exe", "url.dll,FileProtocolHandler",SELECTEDHABITAT+"\\"+"A.mon.exe").Start()
		}else{
			///fmt.Println("ERROR!",ERROR)
			}
	}
/////DEATH(?)/////
	os.Exit(0)
}
