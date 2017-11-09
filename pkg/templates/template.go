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
	jww "github.com/spf13/jwalterweatherman"
)

type FileTemplate struct {
	TemplateName string
	Vars         []string
	License      string
	Data         []byte
}

type AnsTempl interface {
	render() string
}

func (ft FileTemplate) propagateData() {
	a, err := Asset("templates/" + ft.TemplateName)
	if err != nil {
		jww.FATAL.Printf("templates: unable to find template %s", ft.TemplateName)
	}
	ft.Data = a
}

// New returns a new instance of the FileTemplate name
func New(n string) FileTemplate {
	return FileTemplate{
		TemplateName: n,
	}
}