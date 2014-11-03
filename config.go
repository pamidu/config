package main

import (
	"fmt"
	"bufio"
    "os"
    "strings"
)

type config struct{

	 m map[string]string 
}

//Config.txt file load for modifications 
func (c *config) loadfile(){
	 configFile:="Config.txt"
	if c.m==nil{
		c.m=make(map[string]string)
	}
	//file open 
	file, err := os.Open(configFile)
	if file != nil {
    	fmt.Println(file)
	}else{
		fmt.Println("File cannot Open",file,err)
	}
	//read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//split key and value 
		if scanner.Text()!=""{
       		stringSlice := strings.Split(scanner.Text(),"-")
    		c.m[stringSlice[0]]=stringSlice[1]
		}
	}
}

//add new key and value if key exist value will update 
func (c *config) Add(Key, Value string) {
	if c.m==nil{
		c.m=make(map[string]string)
	}
	c.m[Key]=Value
}


//return value to given key 
func (c *config) Get(Key string) string{
	if c.m==nil{
		c.m=make(map[string]string)
	}
	result:=c.m[Key]
	if(result==""){
		fmt.Println("Key Not Fuond ")
	}else{
		fmt.Println(c.m[Key])
	}
	return c.m[Key]
}

//delete data relevent to given key from config.txt 
func (c *config) Delete(Key string){
	if c.m==nil{
		c.m=make(map[string]string)
	}
	delete(c.m,Key)
	
}

//write modified data to config.txt 
func (c *config) writetoFile(){
	configFile:="Config.txt"
	
	if c.m==nil{
		c.m=make(map[string]string)
	}

	file, err := os.Open(configFile)
	
	if err != nil {
    	fmt.Println(err,file)
	}
  	if _, err := os.Stat(configFile); os.IsNotExist(err) {
   		_, err := os.Create(configFile)
   		if err==nil{
   			fmt.Printf("%s file  created ... \n", configFile)
   		}else{
   			fmt.Printf("file cannot create please check file location ")    	
   		}
	}

	file1, err := os.OpenFile(configFile,os.O_WRONLY,0600)
    if err != nil {
        // panic(err)
       fmt.Printf("Appended into file not success please check again \n")
    }
    defer file.Close()
	//fmt.Println(c.m,"before write to file 23")
	for k,v :=range c.m{
		//fmt.Println(k,v,"\n")
	    if _, err = file1.WriteString(k+"-"+v+"\n"); err != nil {
    	  	panic(err)
    	}
	}
}

func main(){
	c:=config{}
	c.loadfile()
	value1:=c.Get("key2")
	if(value1==""){
		fmt.Println("Key Not Fuond ")
	}
	c.Add("key1","test value Add to  key1")
	c.Add("key2","test value Add to  key2")
	c.Add("key3","test value change in key 3")
	c.Delete("key2")
	c.Delete("key1")
	c.Add("key3","test value Add to  key33")
	c.Add("key4","test value Add to  key4")
	c.Add("key5","test value Add to  key5")
	c.Add("key6","test value Add to  key6")
	value:=c.Get("key4")
	if(value==""){
		fmt.Println("Key Not Fuond main")
	}
	c.writetoFile()
	
}