package fastvars

import (
	"errors"
	"fmt"
	"io"

	"github.com/valyala/fasttemplate"
)

const startTag = "#{"
const endTag = "}"

//FastVars map
type FastVars struct {
	variables map[string]interface{}
	tplCache  map[string]fasttemplate.Template
}

//NewFastVars create default FastVars map
func NewFastVars() (FastVars, error) {
	return NewFastVarsDict(map[string]interface{}{})
}

//NewFastVarsDict create FastVars initialize it with given map
func NewFastVarsDict(vars map[string]interface{}) (FastVars, error) {
	var res FastVars
	if vars == nil {
		return res, errors.New("vars is empty")
	}
	res = FastVars{variables: vars, tplCache: map[string]fasttemplate.Template{}}
	return res, nil
}

//Process transform given template
func (fv *FastVars) Process(tpl string) (string, error) {
	t := fasttemplate.New(tpl, startTag, endTag)
	s := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		val, _ := fv.Get(tag)
		str := fmt.Sprintf("%v", val)
		return w.Write([]byte(str))
	})
	return s, nil
}

//Get returns value by key
func (fv *FastVars) Get(key string) (interface{}, error) {
	var res interface{}
	res = fv.variables[key]
	if res != nil {
		switch res.(type) {
		case string:
			return fv.Process(res.(string))
		default:
			return res, nil
		}
	}
	return res, errors.New("no such variable")
}
//Append add vars to fv.variables map
func (fv *FastVars) Append(vars map[string]interface{}) {
	for key, value := range vars {
		fv.variables[key] = value
	}
}
