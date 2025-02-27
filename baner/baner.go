package baner

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func ImageFunc(ByneryImage string, Idage string,folder string) string{

	bynaryData, err := base64.StdEncoding.DecodeString(ByneryImage)
	if err != nil {
		log.Printf("Error image : %v", err)
	}

	errWriteFile := ioutil.WriteFile("./Statics/"+folder+"/"+Idage, bynaryData, 0644)
	if errWriteFile != nil {
		log.Printf("Error write: %v", errWriteFile)
	}
	return fmt.Sprintf(Idage)
}


func PDF_FUNC(ByneryImage string, Idage string,folder string) string{

	data := []byte(ByneryImage)

	err := ioutil.WriteFile("./Statics/"+folder+"/"+Idage, data, 0600)

	if err != nil {
		fmt.Println("Error")
	}
	fil_data, filerr := ioutil.ReadFile("./Statics/"+folder+"/"+Idage)

	if filerr != nil {
		fmt.Println("Can  not")
	}

	fmt.Println(string(fil_data))


	return fmt.Sprintf(Idage)

}


// func PDF_Func(ByneryImage string, Idage string,folder string) string{
	
// 	bynaryData, _ := base64.StdEncoding.DecodeString(ByneryImage)
// 	err:=ioutil.WriteFile("./Statics/"+folder+"/"+Idage, bynaryData,0600)

// 	if err !=nil{
// 		fmt.Println("Error")
// 	}
// 	fil_data,filerr:=ioutil.ReadFile("./Statics/"+folder+"/"+Idage)

// 	if filerr!=nil{
// 		fmt.Println("Can  not ")
// 	}
// 	fmt.Println(string(fil_data))

// 	return Idage
// }





