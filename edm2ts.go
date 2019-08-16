package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Make sure we have at least two args and second isn't empty.  Second arg is URL; first is command executed
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Printf("Usage: %s %s [output file]\n", os.Args[0], "https://api.somesite.com/odata/$metadata")
		os.Exit(1)
	}

	// Default output file name if none provided
	outputFile := "entities.ts"

	//Check that we have a URL to request from
	metadataUrl := os.Args[1]

	//If there are more than two args we assume the third is the output filename.  Path is relative to execution if not provided
	if len(os.Args) > 2 && os.Args[2] != "" {
		outputFile = os.Args[2]
	}

	//Retrieve EDM metadata from odata service
	metadataResponse, err := http.Get(metadataUrl)
	errorCheck(err)

	defer metadataResponse.Body.Close()

	// Convert file to byte[] so it can be unmarshal'd
	byteValue, err := ioutil.ReadAll(metadataResponse.Body)
	errorCheck(err)

	var edmx Edmx
	var outputString string

	//Unmarshal retrieved edm data into our edmx classes
	xmlError := xml.Unmarshal(byteValue, &edmx)
	errorCheck(xmlError)

	//Loop through dataservices->schemas->entity types to get to the properties for each entity
	for _, dataService := range edmx.DataServices {
		for _, schema := range dataService.Schemas {
			for _, entityType := range schema.EntityTypes {

				outputString += "export class " + entityType.Name + " {\n"

				for _, entityData := range entityType.Properties {
					var dataType = "string"
					switch entityData.Type {
					case "Edm.Int32":
						dataType = "number"
					case "Edm.Int64":
						dataType = "number"
					case "Edm.Decimal":
						dataType = "number"
					case "Edm.Boolean":
						dataType = "boolean"
					}

					outputString += "\t" + entityData.Name + ": " + dataType + ";\n"
				}

				outputString += "}" + "\n\n"
			}
		}
	}

	//fmt.Print(outputString)

	//Write the string buffer to the output file.  Perm 0644 is owner read/write and all read
	fileErr := ioutil.WriteFile(outputFile, []byte(outputString), 0644)
	errorCheck(fileErr)

}
