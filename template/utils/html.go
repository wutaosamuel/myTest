package utils

import (
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ReplaceHTML check replace template
func ReplaceHTML(template string, num int, pattern string) (string, error) {
	symbol := "{{{ " + strconv.Itoa(num) + " }}}"
	if !strings.Contains(template, symbol) {
		return "", Err("Template do not contain " + symbol)
	}
	return strings.Replace(template, symbol, pattern, 1), nil
}

// ReplacePattern replace {{{ num }}} in pattern by struct
func ReplacePattern(pattern string, body interface{}) (string, error) {
	var p string
	b := reflect.ValueOf(body)
	for b.Kind() == reflect.Ptr || b.Kind() == reflect.Interface {
		b = b.Elem()
	}
	if b.Kind() != reflect.Struct {
		return "", Err("ReplacePattern not valid type")
	}
	for i := 0; i < b.NumField(); i++ {
		if i == 0 {
			p, _ = ReplaceHTML(pattern, i+1, b.Field(i).String())
		}
		if i != 0 {
			p, _ = ReplaceHTML(p, i+1, b.Field(i).String())
		}
	}
	return p, nil
}

// AppendHTML append content
// the symbol is <!-- {{{ }}} -->
func AppendHTML(template, pattern string) (string, error) {
	symbol := "<!-- {{{ 0 }}} -->"
	if !strings.Contains(template, symbol) {
		return "", Err("Template do not contain " + symbol)
	}
	pattern = pattern + "\n" + symbol + "\n"
	return strings.Replace(template, symbol, pattern, 1), nil
}

// AppendObj append obj into html and return by string
func AppendObj(body interface{}, name, pattern string) (string, error) {
	// read html && pattern
	h, err := ReadHTML(name)
	if err != nil {
		return "", err
	}
	p, err := ReadHTML(pattern)
	if err != nil {
		return "", err
	}

	// replace pattern by obj
	rp, err := ReplacePattern(p, body)
	if err != nil {
		return "", err
	}
	// append into html
	html, err := AppendHTML(h, rp)
	if err != nil {
		return "", err
	}
	return html, nil
}

// AppendPage update html after append a obj
func AppendPage(body interface{}, name, pattern string) error {
	html, err := AppendObj(body, name, pattern)
	if err != nil {
		return err
	}
	return SaveHTML(name, html)
}

// DeleteHTML delete obj by pattern to string
func DeleteHTML(template, pattern string) (string, error) {
	if !strings.Contains(template, pattern) {
		return "", Err("Page do not contain the object")
	}
	return strings.Replace(template, pattern, "", 1), nil
}

// DeleteObj delete a struct obj by pattern to string
func DeleteObj(body interface{}, name, pattern string) (string, error) {
	h, err := ReadHTML(name)
	if err != nil {
		return "", err
	}
	p, err := ReadHTML(pattern)
	if err != nil {
		return "", err
	}

	rp, err := ReplacePattern(p, body)
	if err != nil {
		return "", err
	}

	html, err := DeleteHTML(h, rp)
	if err != nil {
		return "", err
	}
	return html, nil
}

// DeletePage delete a obj in html
func DeletePage(body interface{}, name, pattern string) error {
	html, err := DeleteObj(body, name, pattern)
	if err != nil {
		return err
	}
	return SaveHTML(name, html)
}

// SaveHTML save html file from string
func SaveHTML(name, html string) error {
	if err := ioutil.WriteFile(name, []byte(html), 0777); err != nil {
		return err
	}
	return nil
}

// ReadHTML read html file
func ReadHTML(name string) (string, error) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
