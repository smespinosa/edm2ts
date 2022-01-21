package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	metadataUrl := os.Args[1] + "$metadata"

	//Losing $metadata from os.Args?  Make sure its provided or we end up with json describing endpoints only
	if !strings.Contains(metadataUrl, "$metadata") {
		if !strings.HasSuffix(metadataUrl, "/") {
			metadataUrl += "/"
		}
		metadataUrl += "$metadata"
	}

	//If there are more than two args we assume the third is the output filename.  Path is relative to execution if not provided
	if len(os.Args) > 2 && os.Args[2] != "" {
		outputFile = os.Args[2]
	}

	byteValue, err := httpCall(metadataUrl)
	errorCheck(err)

	var edmx Edmx
	var outputString string

	//Unmarshal retrieved edm data into our edmx objects
	xmlError := xml.Unmarshal(byteValue, &edmx)
	errorCheck(xmlError)

	// Loop through dataservices->schemas->entity types to get to the properties for each entity
	for _, dataService := range edmx.DataServices {
		for _, schema := range dataService.Schemas {
			for _, entityType := range schema.EntityTypes {
				outputString += "export interface " + entityType.Name + " {\n"

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
					case "Edm.DateTimeOffset":
						dataType = "Date"
					case "Edm.DateTime":
						dataType = "Date"
					}

					//if nullable is not specified it is implied nullable
					isNullable := true

					if len(entityData.Nullable) > 0 {
						isNullable, err = strconv.ParseBool(entityData.Nullable)
						errorCheck(err)
					}

					newProperty := "\t" + entityData.Name + ": " + dataType

					if isNullable {
						newProperty += " | null"
					}

					newProperty += ";\n"

					outputString += newProperty
				}

				for _, navigationData := range entityType.NavigationProperty {
					isCollection := strings.HasPrefix(navigationData.Type, "Collection")

					navigationPropertyExpanded := strings.Split(navigationData.Type, ".")
					navigationPropertyType := navigationPropertyExpanded[len(navigationPropertyExpanded)-1]

					if isCollection {
						navigationPropertyType = navigationPropertyType[0:len(navigationPropertyType)-1] + "[]"
					}

					newNavigrationProperty := "\t" + navigationData.Name + "?: " + navigationPropertyType + ";\n"

					outputString += newNavigrationProperty
				}

				outputString += "}" + "\n\n"
			}
		}
	}

	fileErr := ioutil.WriteFile(outputFile, []byte(outputString), 0644)
	errorCheck(fileErr)

}

func httpCall(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	byteValue, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
