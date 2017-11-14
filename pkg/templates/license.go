// MIT License

// Copyright (c) 2017 Wian Vos

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package templates

import (
	"bytes"
	"errors"
	"path/filepath"
	"strings"
	"text/template"

	jww "github.com/spf13/jwalterweatherman"
)

type License struct {
	Template *template.Template
	Name     string
	Data     map[string]string
}

type Licenses []License

func NewLicenses() Licenses {

	ls := make(Licenses, 0)

	for _, l := range licenseAssets() {
		ls = append(ls, NewLicense(l))
	}
	return ls
}

func NewLicense(p string) License {

	name := strings.TrimSuffix(filepath.Base(p), filepath.Ext(filepath.Base(p)))
	d, err := Asset(p)

	if err != nil {
		jww.ERROR.Printf("Templates:NewLicense: unable to retrieve data for license: %s : %s", name, err)
	}

	tmpl, _ := template.New(name).Parse(string(d))

	l := License{
		Name:     name,
		Template: tmpl,
	}

	return l
}

func licenseAssets() []string {

	var a = make([]string, 0)
	for _, t := range AssetNames() {
		if strings.Contains(t, "licenses") {
			a = append(a, t)
		}
	}
	return a
}

func (l *Licenses) AddData(d map[string]string) {

	ls := *l
	ls2 := make(Licenses, 0)

	for _, lic := range ls {
		lic.Data = d
		ls2 = append(ls2, lic)
	}

	*l = ls2
}

func (ls Licenses) GetLicense(n string) (string, error) {

	var o string
	for _, l := range ls {
		if l.Name == n {
			return l.String()
		}
	}

	return o, errors.New("No such license")
}

func (l License) String() (string, error) {

	var err error

	if err != nil {
		return "", err
	}

	var lb bytes.Buffer
	if err = l.Template.Execute(&lb, l.Data); err != nil {
		return "", err
	}

	result := lb.String()
	return result, err
}
