package confish

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

var spacePerIndent = 4

var indentSpaces map[int]string

func init() {
	indentSpaces = make(map[int]string, 10)
}

func indent(n int) string {
	if s, ok := indentSpaces[n]; ok {
		return s
	}
	s := strings.Repeat(" ", n*spacePerIndent)
	indentSpaces[n] = s
	return s
}

type confWriter struct {
	w           io.Writer
	indentLevel int
}

func (cw *confWriter) outputStruct(conf interface{}) error {
	v := reflect.ValueOf(conf)
	if v.Kind() != reflect.Struct {
		return errors.New("value is not a Struct type")
	}

	cw.indentLevel++

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		attr := t.Field(i).Tag.Get("cfg-attr")
		if attr == "" {
			// No cfg-attr tag attached, ignore this field.
			continue
		}

		fv := indirect(f.Interface())
		ft := reflect.TypeOf(fv)

		switch ft.Kind() {
		case reflect.Struct:
			cw.outputStruct(fv)
		case reflect.Slice:
			cw.outputSlice(attr, fv)
		case reflect.Map:
			cw.outputMap(attr, fv)
		case reflect.Int, reflect.Int32, reflect.Int64:
			fmt.Fprintf(cw.w, "%s%s: %d\n", indent(cw.indentLevel), attr, fv)
		case reflect.Float32, reflect.Float64:
			fmt.Fprintf(cw.w, "%s%s: %f\n", indent(cw.indentLevel), attr, fv)
		case reflect.Bool:
			fmt.Fprintf(cw.w, "%s%s: %t\n", indent(cw.indentLevel), attr, fv)
		case reflect.String:
			fmt.Fprintf(cw.w, "%s%s: \"%s\"\n", indent(cw.indentLevel), attr, fv)
		}

		fmt.Fprint(cw.w)
	}

	cw.indentLevel--
	return nil
}

func (cw *confWriter) outputSlice(attr string, conf interface{}) {
	v := reflect.ValueOf(conf)
	if v.Kind() != reflect.Slice {
		return
	}

	et := indirectType(v.Type().Elem())

	switch et.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Fprintf(cw.w, "%s%s: [\n", indent(cw.indentLevel), attr)
		for i := 0; i < v.Len(); i++ {
			fmt.Fprintf(cw.w, "%s%d,\n", indent(cw.indentLevel+1), v.Index(i).Interface())
		}
		fmt.Fprintf(cw.w, "%s]\n", indent(cw.indentLevel))
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(cw.w, "%s%s: [\n", indent(cw.indentLevel), attr)
		for i := 0; i < v.Len(); i++ {
			fmt.Fprintf(cw.w, "%s%f,\n", indent(cw.indentLevel+1), v.Index(i).Interface())
		}
		fmt.Fprintf(cw.w, "%s]\n", indent(cw.indentLevel))
	case reflect.String:
		fmt.Fprintf(cw.w, "%s%s: [\n", indent(cw.indentLevel), attr)
		for i := 0; i < v.Len(); i++ {
			fmt.Fprintf(cw.w, "%s\"%s\",\n", indent(cw.indentLevel+1), v.Index(i).Interface())
		}
		fmt.Fprintf(cw.w, "%s]\n", indent(cw.indentLevel))
	case reflect.Struct:
		for i := 0; i < v.Len(); i++ {
			fmt.Fprintf(cw.w, "%s%s {\n", indent(cw.indentLevel), attr)
			cw.outputStruct(indirect(v.Index(i).Interface()))
			fmt.Fprintf(cw.w, "%s}\n", indent(cw.indentLevel))
		}
	}
}

func (cw *confWriter) outputMap(attr string, conf interface{}) error {
	v := reflect.ValueOf(conf)
	if v.Kind() != reflect.Map {
		return errors.New("value is not a slice type")
	}

	fmt.Fprintf(cw.w, "%s%s: {\n", indent(cw.indentLevel), attr)
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)
		fmt.Fprintf(cw.w, "%s", indent(cw.indentLevel+1))
		fmt.Fprintf(cw.w, "%s: %s,\n", primitiveString(key), primitiveString(val))
	}
	fmt.Fprintf(cw.w, "%s}\n", indent(cw.indentLevel))
	return nil
}

func WriteFile(confFile string, confVar interface{}, sectionName string) error {
	f, err := os.Create(confFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return Write(f, confVar, sectionName)
}

func Write(w io.Writer, confVar interface{}, sectionName string) error {
	conf := indirect(confVar)
	cw := &confWriter{w, 0}
	fmt.Fprintf(w, "%s {\n", sectionName)
	cw.outputStruct(conf)
	fmt.Fprint(w, "}\n")
	return nil
}
