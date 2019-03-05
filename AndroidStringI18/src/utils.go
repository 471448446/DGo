package src

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)


func PrintErrorIfExist(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func UTF82GBK(src string) (string, error) {
	reader := transform.NewReader(strings.NewReader(src), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return src, err
	} else {
		return string(buf), nil
	}
}

func CreateStartElement(name string, attrs [][] string) xml.StartElement {
	var attr [] xml.Attr
	for _, e := range attrs {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: e[0]}, Value: e[1]})
	}
	return xml.StartElement{
		Name: xml.Name{Local: name},
		Attr: attr,
	}
}
