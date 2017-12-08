package main

import (
	"fmt"
	"github.com/LibiChai/mypass/lib"
	"os"
)

func main(){


	argCount := len(os.Args[1:])

	if(argCount != 3 || (os.Args[1] != "-e") && (os.Args[1] != "-d")){
		fmt.Printf("Usage: mypass -e|-d content key \n \n" +
			"   -e   encrypted content use key \n" +
			"   -d   deciphering content use key \n")
	}
	var output string;
	var err error
	if(os.Args[1] == "-e"){
		output,err = encode(os.Args[2],os.Args[3])
	}

	if(os.Args[1] == "-d"){
		output,err = decode(os.Args[2],os.Args[3])
	}

	if(err != nil){
		fmt.Println(err.Error())
		return
	}

	fmt.Println(output)



}


func encode(content ,key string) (string,error){

	tmp ,err := lib.AesEncrypt([]byte(content),[]byte(key))

	if(err != nil){
		return "",err
	}

	return lib.EncodeSegment(tmp),nil

}

func decode(secretContent,key string)(string,error){

	byteSecretContent,err := lib.DecodeSegment(secretContent)

	if(err != nil){
		return "",err
	}
	byteContent ,err := lib.AesDecrypt(byteSecretContent,[]byte(key))

	if(err != nil){
		return "",err
	}

	return string(byteContent),nil
}

