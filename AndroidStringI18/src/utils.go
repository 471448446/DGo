package src

import (
	"encoding/xml"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

func UTF82GBK(src string) (string, error) {
	reader := transform.NewReader(strings.NewReader(src), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return src, err
	} else {
		return string(buf), nil
	}
}
func GBK2UTF8(src string) (string, error) {
	reader := transform.NewReader(strings.NewReader(src), unicode.UTF8.NewDecoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return src, err
	} else {
		return string(buf), nil
	}
}

func StartElement(name string, attrs [][] string) xml.StartElement {
	var attr [] xml.Attr
	for _, e := range attrs {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: e[0]}, Value: e[1]})
	}
	return xml.StartElement{
		Name: xml.Name{Local: name},
		Attr: attr,
	}
}
