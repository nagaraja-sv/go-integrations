package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Dispensing is to
type Dispensing struct {
	Product []struct {
		Name            string `json:"Name"`
		ID              string `json:"ID"`
		ProductRecordID string `json:"ProductRecordId"`
		Description     string `json:"Description"`
		StdQty          string `json:"StdQty"`
		RawMaterials    []struct {
			Process []struct {
				ProcessName    string `json:"ProcessName"`
				ProcessNameSeq string `json:"ProcessNameSeq#"`
				Materials      []struct {
					MaterailName string `json:"MaterailName"`
					MaterailID   string `json:"MaterailID"`
					Spec         string `json:"Spec"`
					StdQty       string `json:"StdQty "`
				} `json:"Materials"`
			} `json:"Process"`
		} `json:"RawMaterials"`
		PackingMaterials []struct {
			MaterailName string `json:"MaterailName"`
			StdQty       string `json:"StdQty "`
		} `json:"PackingMaterials"`
	} `json:"Product"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("procurement.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened procurement.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Dispensing
	var kinDispensing Dispensing

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'kinDispensing' which we defined above
	json.Unmarshal(byteValue, &kinDispensing)

	// we iterate through every user within our Dispensing array
	for i := 0; i < len(kinDispensing.Product); i++ {

		fmt.Println("Product Name " + kinDispensing.Product[0].Name)
		fmt.Println("Product StdQty " + kinDispensing.Product[0].StdQty)
		fmt.Println("Processname " + kinDispensing.Product[0].RawMaterials[0].Process[0].ProcessName)
		fmt.Println("Material " + kinDispensing.Product[0].RawMaterials[0].Process[0].Materials[0].MaterailName)
		fmt.Println("Material StdQuty " + kinDispensing.Product[0].RawMaterials[0].Process[0].Materials[0].StdQty)

		//requriedqunity := batch * MaterialStdQuty / stdy

		if kinDispensing.Product[i].Name == "CLOPIDOGREL PELLETS 50 %  w/w" {

			stdy := kinDispensing.Product[i].StdQty
			Processname := kinDispensing.Product[i].RawMaterials[i].Process[i].ProcessName

			fmt.Println("Processname " + Processname)
			fmt.Println("Product StdQty " + stdy)

			if kinDispensing.Product[i].RawMaterials[i].Process[i].Materials[i].MaterailName == "Clopidogrel Bisulphate" {

				MaterialStdQuty := kinDispensing.Product[i].RawMaterials[i].Process[i].Materials[i].StdQty
				fmt.Println("MaterialStdQuty " + MaterialStdQuty)
			}

		}
	}

}
