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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	jww "github.com/spf13/jwalterweatherman"
)

const (
	fileMode = 0770
)

type FileTemplate struct {
	Path     string
	Template *template.Template
	Data     map[string]string
}

type FileTemplates []FileTemplate

func NewFileTemplates() FileTemplates {

	f := make(FileTemplates, 0)

	for _, t := range templateAssets() {
		f = append(f, NewFileTemplate(t))
	}
	return f
}

func NewFileTemplate(p string) FileTemplate {

	path := p[len("templates"):len(p)]

	d, err := Asset(p)
	if err != nil {
		jww.ERROR.Printf("Templates:NewFileTemplate: unable to retrieve data for template: %s : %s", p, err)
	}
	tmpl, _ := template.New(path).Parse(string(d))

	ft := FileTemplate{
		Path:     path,
		Template: tmpl,
	}

	return ft
}

func templateAssets() []string {

	var a = make([]string, 0)
	for _, t := range AssetNames() {
		if strings.Contains(t, "templates") {
			a = append(a, t)
		}
	}
	return a
}

func (ft FileTemplate) Write(d string) error {
	f := d + ft.Path

	if _, err := os.Stat(filepath.Dir(f)); err != nil {
		err := os.MkdirAll(filepath.Dir(f), os.FileMode(fileMode))
		if err != nil {
			return err
		}
	}

	iow, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0770)
	if err != nil {
		fmt.Println(err)
		return err
	}

	ft.Template.Execute(iow, ft.Data)

	return nil
}

func (ft *FileTemplates) AddData(d map[string]string) {

	fts := *ft
	fts2 := make(FileTemplates, 0)

	for _, t := range fts {
		t.Data = d
		fts2 = append(fts2, t)
	}

	*ft = fts2
}
