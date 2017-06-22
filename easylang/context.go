package easylang

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type context interface {
	define(varName string, expr expression)
	defineParam(varName string, expr expression)
	isDefined(varName string) bool
	isParamDefined(varName string) bool
}

type Context struct {
	sequence int

	params   []string
	paramMap map[string]expression

	definedVars   []string
	definedVarMap map[string]expression

	outputVars       []string
	outputLineWidth  []int
	outputColors     []string
	outputLineStyles []string
	outputOtherDesc  []string

	// TODO: Handle errors
}

func newContext() *Context {
	return &Context{
		paramMap:      map[string]expression{},
		definedVarMap: map[string]expression{},
	}
}

func (this *Context) newAnonymousVarName() string {
	ret := fmt.Sprintf("__anonymous_%d", this.sequence)
	this.sequence++
	return ret
}

func (this *Context) define(varName string, expr expression) {
	varName = strings.ToLower(varName)
	if expr, ok := this.definedVarMap[varName]; ok {
		expr.IncrementRefCount()
		return
	}
	this.definedVarMap[varName] = expr
	this.definedVars = append(this.definedVars, varName)
	expr.IncrementRefCount()
}

func (this *Context) defineParam(varName string, expr expression) {
	varName = strings.ToLower(varName)
	if _, ok := this.paramMap[varName]; ok {
		return
	}
	this.paramMap[varName] = expr
	this.params = append(this.params, varName)

	this.definedVarMap[varName] = expr
	this.definedVars = append(this.definedVars, varName)
	expr.IncrementRefCount()
}

func (this Context) isDefined(varName string) bool {
	_, ok := this.definedVarMap[varName]
	return ok
}

func (this Context) isParamDefined(varName string) bool {
	_, ok := this.paramMap[varName]
	return ok
}

func (this *Context) addOutput(varName string, descriptions []string, line, column int) {
	varName = strings.ToLower(varName)
	for _, v := range this.outputVars {
		if v == varName {
			panic("duplicate output variable")
		}
	}
	this.outputVars = append(this.outputVars, varName)
}

func (this *Context) defined(varName string) expression {
	expr, _ := this.definedVarMap[varName]
	return expr
}

func (this *Context) definedParam(varName string) expression {
	expr, _ := this.paramMap[varName]
	return expr
}

func (this *Context) definedCodes(indent string) string {
	lines := make([]string, len(this.definedVarMap))
	i := 0
	for _, varName := range this.definedVars {
		expr, ok := this.definedVarMap[varName]
		if !ok {
			continue
		}
		lines[i] = fmt.Sprintf("%so.%s = %s", indent, varName, expr.Codes())
		i++
	}
	return strings.Join(lines, "\n")
}

func (this *Context) updateLastValueCodes(indent string) string {
	lines := []string{}
	for _, varName := range this.definedVars {
		expr, ok := this.definedVarMap[varName]
		if !ok {
			continue
		}
		switch expr.(type) {
		case *constantexpr:
		case *assignexpr:
		case *paramexpr:
		default:
			lines = append(lines, fmt.Sprintf("%so.%s.updateLastValue()", indent, varName))
		}
	}
	return strings.Join(lines, "\n")
}

func (this *Context) refValuesCodes() string {
	items := make([]string, len(this.outputVars))
	for i, varName := range this.outputVars {
		items[i] = fmt.Sprintf("o.%s", varName)
	}
	return strings.Join(items, ", ")
}

func (this *Context) getCodes(indent string) string {
	lines := make([]string, len(this.outputVars))
	for i, varName := range this.outputVars {
		lines[i] = fmt.Sprintf("%so.%s.Get(index),", indent, varName)
	}
	return strings.Join(lines, "\n")
}

func (this *Context) paramCodes() string {
	sa := make([]string, len(this.params)+1)
	sa[0] = "data"
	for i, p := range this.params {
		sa[i+1] = p
	}
	return strings.Join(sa, ", ")
}

func (this *Context) removeUnusedParams() {
	for _, expr := range this.paramMap {
		pExpr, _ := expr.(*paramexpr)
		if pExpr.operand.RefCount() == 1 {
			if _, ok := this.definedVarMap[pExpr.operand.VarName()]; ok {
				delete(this.definedVarMap, pExpr.operand.VarName())
			}
		}
	}
}

func (this *Context) generateCode(name string, filePath string) error {
	name = strings.ToUpper(name)

	this.removeUnusedParams()

	const indent = "    "
	code := fmt.Sprintf(`-----------------------------------------------------------
-- GENERATED BY EASYLANG COMPILER.
-- !!!! DON'T MODIFY IT!!!!!!
-----------------------------------------------------------

%sClass = {}


function %sClass:new(%s)
    o = {}
    setmetatable(o, self)
    self.__index = self

    o.data = data
%s

    o.ref_values = {%s}
    return o
end

function %sClass:updateLastValue()
%s
end

function %sClass:Len()
    return self.data.Len()
end


function %sClass:Get(index)
    return {
%s
    }
end

FormulaClass = %sClass
	`, name,
		name,
		this.paramCodes(),
		this.definedCodes(indent),
		this.refValuesCodes(),
		name,
		this.updateLastValueCodes(indent),
		name,
		name,
		this.getCodes("        "),
		name)

	return ioutil.WriteFile(filePath, []byte(code), 0666)
}