package confish

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

// The parser expects the lexer to return 0 on Eof. Give it a name for clarity.
const Eof = 0

func init() {
	cfgDebug = 1
	cfgErrorVerbose = true
}

// The parser uses the type <prefix>Lex as a lexer. It must provide the methods
// Lex(*<prefix>SymType) int and Error(string).
type cfgLex struct {
	text   []byte
	peek   rune
	retVal *rune
	stack  *confStack
	inMap  bool
}

func (l *cfgLex) SetInMap(inMap bool) {
	l.inMap = inMap
}

func (l *cfgLex) Push(t int, name string) {
	name = strings.Trim(name, " \t")

	item := l.stack.Peek()
	elem := item.Elem.(reflect.Value)

	e := indirectVal(elem)
	et := e.Type()

	if et.Kind() == reflect.Struct {
		for i := 0; i < et.NumField(); i++ {
			field := et.Field(i)
			tag := field.Tag.Get("cfg-attr")
			if tag == name {
				elemField := e.Field(i)
				f := indirectVal(elemField)
				ft := f.Type()

				switch ft.Kind() {
				case reflect.String, reflect.Bool:
					l.stack.Push(elemField)
					return
				case reflect.Int, reflect.Int32, reflect.Int64:
					l.stack.Push(elemField)
					return
				case reflect.Map:
					if f.IsNil() {
						f.Set(reflect.MakeMap(ft))
					}
					l.stack.Push(elemField)
				case reflect.Slice:
					if f.IsNil() {
						f.Set(reflect.MakeSlice(ft, 0, 10))
					}

					fieldType := elemField.Type()
					fet := indirectType(fieldType.Elem())
					switch fet.Kind() {
					case reflect.Struct:
						child := reflect.New(fet)
						elemField.Set(reflect.Append(elemField, child))
						l.stack.Push(child)
					default:
						l.stack.Push(elemField)
					}
				default:
					l.stack.Push(elemField)
				}
			}
		}
	}
}

func (l *cfgLex) Pop() {
	_, nonEmpty := l.stack.Pop()
	if !nonEmpty {
		panic("Error to pop stack: ")
	}
}

func (l *cfgLex) AttrVal(val string) {
	item := l.stack.Peek()
	elem := item.Elem.(reflect.Value)
	elemType := elem.Type()
	switch elemType.Kind() {
	case reflect.String:
		elem.SetString(val)
	case reflect.Int, reflect.Int32, reflect.Int64:
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		elem.SetInt(int64(intVal))
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(err)
		}
		elem.SetFloat(floatVal)
	case reflect.Bool:
		elem.SetBool(strings.ToLower(val) == "true")
	default:
		panic(fmt.Sprintf("Unknown attr type: %+v for value %s.", elemType.Kind(), val))
	}
}

func (l *cfgLex) Append(val string) {
	item := l.stack.Peek()
	elem := item.Elem.(reflect.Value)
	elem.Set(reflect.Append(elem, reflect.ValueOf(val)))
}

func (l *cfgLex) Put(key, val string) {
	item := l.stack.Peek()
	elem := item.Elem.(reflect.Value)
	elem.SetMapIndex(reflect.ValueOf(convert(key, elem.Type().Key())),
		reflect.ValueOf(convert(val, elem.Type().Elem())))
}

// The parser calls this method to get each new token.
func (l *cfgLex) Lex(yylval *cfgSymType) int {
	yylval.val = ""
	for {
		c := l.next()
		switch c {
		case Eof:
			return Eof
		case ':', ',', '{', '}', '[', ']':
			return int(c)
		case ' ', '\t', '\n', '\r':
			// Ignore white spaces.
		case '#':
			l.ignoreCommentLine(l.next())
		case '"':
			l.matchString(l.next(), yylval)
			return SCALAR
		default:
			return l.matchAhead(c, yylval)
		}
	}
}

// Return the next rune for the lexer.
func (l *cfgLex) next() rune {
	if l.retVal != nil {
		val := *l.retVal
		l.retVal = nil
		return val
	}

	if l.peek != Eof {
		r := l.peek
		l.peek = Eof
		return r
	}
	if len(l.text) == 0 {
		return Eof
	}
	c, size := utf8.DecodeRune(l.text)
	l.text = l.text[size:]
	if c == utf8.RuneError && size == 1 {
		log.Print("invalid utf8")
		return l.next()
	}
	return c
}

func (l *cfgLex) unreadRune(c rune) {
	l.retVal = &c
}

// Lex a comment.
func (l *cfgLex) ignoreCommentLine(c rune) {
	for {
		switch c {
		case '\n':
			l.unreadRune(c)
			return
		}

		c = l.next()
	}
}

// Lex a string
func (l *cfgLex) matchString(c rune, yylval *cfgSymType) {
	var buf bytes.Buffer
	var lastRune rune = Eof
	for {
		switch c {
		case Eof:
			panic("failed to match a string literal")
		case '"':
			if lastRune != '\\' {
				yylval.val = buf.String()
				return
			}
		}

		lastRune = c
		buf.WriteRune(c)
		c = l.next()
	}
}

// Lex a scalar
func (l *cfgLex) matchAhead(c rune, yylval *cfgSymType) int {
	var buf bytes.Buffer
	for {
		switch c {
		case Eof:
			if buf.Len() > 0 {
				yylval.val = buf.String()
				l.unreadRune(c)
				return SCALAR
			}
			return Eof
		case '{':
			yylval.val = buf.String()
			l.unreadRune(c)
			return SECTION_NAME
		case ':':
			yylval.val = buf.String()
			l.unreadRune(c)
			if l.inMap {
				return SCALAR
			} else {
				return ATTR_NAME
			}
		case ',', '}', ']':
			yylval.val = buf.String()
			l.unreadRune(c)
			return SCALAR
		case '\n', '\r':
			yylval.val = buf.String()
			return SCALAR
		case '#':
			l.ignoreCommentLine(l.next())
			c = l.next()
			continue
		case ' ', '\t':
			// Ignore white spaces.
		}

		buf.WriteRune(c)
		c = l.next()
	}
	return Eof
}

func (l *cfgLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func newCfgLex(content []byte, cfgVar interface{}) *cfgLex {
	stack := newConfStack()

	stack.Push(reflect.ValueOf(cfgVar).Elem())

	return &cfgLex{
		text:  content,
		stack: stack,
	}
}

func indirectVal(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			newVal := reflect.New(val.Type().Elem())
			val.Set(newVal)
			return newVal.Elem()
		}
		val = val.Elem()
	}
	return val
}

func indirectType(val reflect.Type) reflect.Type {
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func convert(val string, toType reflect.Type) interface{} {
	switch toType.Kind() {
	case reflect.String:
		return val
	case reflect.Int:
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return intVal
	case reflect.Int32:
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return int32(intVal)
	case reflect.Int64:
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return int64(intVal)
	case reflect.Float32:
		floatVal, err := strconv.ParseFloat(val, 32)
		if err != nil {
			panic(err)
		}
		return float32(floatVal)
	case reflect.Float64:
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(err)
		}
		return floatVal
	case reflect.Bool:
		return strings.ToLower(val) == "true"
	default:
		panic(fmt.Sprintf("Unknown value type: %+v for value %s.", toType.Kind(), val))
	}
}
