package main

import (
	"encoding/xml"
)

type Edmx struct {
	XMLName      xml.Name      `xml:"Edmx"`
	DataServices []DataService `xml:"DataServices"`
	Version      string        `xml:"Version,attr"`
}

type DataService struct {
	XMLName   xml.Name `xml:"DataServices"`
	Schemas   []Schema `xml:"Schema"`
	XMLNS     string   `xml:"xmlns,attr"`
	Namespace string   `xml:"Namespace,attr"`
}

type Schema struct {
	XMLName          xml.Name          `xml:"Schema"`
	EntityTypes      []EntityType      `xml:"EntityType"`
	EntityContainers []EntityContainer `xml:"EntityContainer"`
	XMLNS            string            `xml:"xmlns,attr"`
	Namespace        string            `xml:"Namespace"`
}

type EntityContainer struct {
	XMLName    xml.Name    `xml:"EntityContainer"`
	Name       string      `xml:"Name,attr"`
	EntitySets []EntitySet `xml:"EntitySet"`
}

type EntitySet struct {
	XMLName    xml.Name `xml:"EntitySet"`
	Name       string   `xml:"Name,attr"`
	EntityType string   `xml:"EntityType,attr"`
}

type EntityType struct {
	XMLName            xml.Name             `xml:"EntityType"`
	Key                EntityKey            `xml:"Key"`
	Name               string               `xml:"Name,attr"`
	Properties         []Property           `xml:"Property"`
	NavigationProperty []NavigationProperty `xml:"NavigationProperty"`
}

type EntityKey struct {
	XMLName      xml.Name      `xml:"Key"`
	PropertyRefs []PropertyRef `xml:"PropertyRef"`
}

type PropertyRef struct {
	XMLName xml.Name `xml:"PropertyRef"`
	Name    string   `xml:"Name,attr"`
}

type Property struct {
	XMLName  xml.Name `xml:"Property"`
	Name     string   `xml:"Name,attr"`
	Type     string   `xml:"Type,attr"`
	Nullable string   `xml:"Nullable,attr"`
}

type NavigationProperty struct {
	XMLName xml.Name `xml:"NavigationProperty"`
	Name    string   `xml:"Name,attr"`
	Type    string   `xml:"Type,attr"`
}
