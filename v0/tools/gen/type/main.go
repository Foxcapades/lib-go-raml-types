package main

import (
	"os"
	"strings"
	"text/template"
	"time"
)

const (
	iFilePrefix = "v0/pkg/raml/x-gen-type-"
	mFilePrefix = "v0/internal/raml/x-gen-type-"
	fileSuffix  = ".go"
)

var now = time.Now().Format(time.RFC3339Nano)

var types = []extTypeProps{
	{Name: "Any",          Type: "any",           Time: now, Base: true},
	{Name: "Include",      Type: "include",       Time: now, Base: true},
	{Name: "Nil",          Type: "nil",           Time: now, Base: true},

	{Name: "Array",        Type: "array",         Time: now, DefType: "[]interface{}", DefTypeName: "[]interface{}", EnumType: "interface{}"},
	{Name: "Bool",         Type: "bool",          Time: now, DefType: "bool",          DefTypeName: "Bool",          EnumType: "bool",        DefIsOpt: true},
	{Name: "Custom",       Type: "custom",        Time: now, DefType: "interface{}",   DefTypeName: "Untyped",       EnumType: "interface{}", DefIsOpt: true},
	{Name: "DateOnly",     Type: "date-only",     Time: now, DefType: "string",        DefTypeName: "String",        EnumType: "string",      DefIsOpt: true},
	{Name: "TimeOnly",     Type: "time-only",     Time: now, DefType: "string",        DefTypeName: "String",        EnumType: "string",      DefIsOpt: true},
	{Name: "DatetimeOnly", Type: "datetime-only", Time: now, DefType: "string",        DefTypeName: "String",        EnumType: "string",      DefIsOpt: true},
	{Name: "Datetime",     Type: "datetime",      Time: now, DefType: "string",        DefTypeName: "String",        EnumType: "string",      DefIsOpt: true},
	{Name: "File",         Type: "file",          Time: now, DefType: "interface{}",   DefTypeName: "Untyped",       EnumType: "interface{}", DefIsOpt: true},
	{Name: "Integer",      Type: "integer",       Time: now, DefType: "int64",         DefTypeName: "Int64",         EnumType: "int64",       DefIsOpt: true},
	{Name: "Number",       Type: "number",        Time: now, DefType: "float64",       DefTypeName: "Float64",       EnumType: "float64",     DefIsOpt: true},
	{Name: "Object",       Type: "object",        Time: now, DefType: "interface{}",   DefTypeName: "Untyped",       EnumType: "interface{}", DefIsOpt: true},
	{Name: "String",       Type: "string",        Time: now, DefType: "string",        DefTypeName: "String",        EnumType: "string",      DefIsOpt: true},
	{Name: "Union",        Type: "union",         Time: now, DefType: "interface{}",   DefTypeName: "Untyped",       EnumType: "interface{}", DefIsOpt: true},
}

func main() {
	iTpl := template.Must(template.ParseGlob("v0/gen/type-i/*"))
	mTpl := template.Must(template.ParseGlob("v0/gen/type-m/*"))

	for i := range types {
		iFile, err := os.Create(iFilePrefix + types[i].Type + fileSuffix)
		check(err)
		mFile, err := os.Create(mFilePrefix + types[i].Type + fileSuffix)
		check(err)

		check(iTpl.ExecuteTemplate(iFile, "root", types[i]))
		check(mTpl.ExecuteTemplate(mFile, "root", types[i]))

		iFile.Close()
		mFile.Close()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type extTypeProps struct {
	Name string
	Type string
	Time string

	DefType     string
	DefTypeName string
	DefIsOpt    bool

	EnumType string
	Base     bool
}

func (e extTypeProps) IsDefPtr() bool {
	return !(strings.HasPrefix(e.DefType, "[]") ||
		strings.HasPrefix(e.DefType, "map"))
}
