%{
package confish

%}

%union{
	val string
}

%token ':' ',' '{' '}' '[' ']'
%token SECTION_NAME ATTR_NAME SCALAR

%%

sections
    :
    | sections section
    ;

section
    : SECTION_NAME {
        var a interface{} = cfglex
        a.(*cfgLex).Push(0, $1.val)
      }

      '{'

      section_body {
        var a interface{} = cfglex
        a.(*cfgLex).Pop()
      }

      '}'
    ;

section_body
    :
    | section_body attr_def
    | section_body section
    ;

attr_def
    : ATTR_NAME {
        var a interface{} = cfglex
        a.(*cfgLex).Push(3, $1.val)
      }

      ':'

      attr_value {
        var a interface{} = cfglex
        a.(*cfgLex).Pop()
      }
    ;

attr_value
    : SCALAR {
        var a interface{} = cfglex
        a.(*cfgLex).AttrVal($1.val)
      }

    | sequence
    | map
    ;

sequence
    : '[' sequence_items ']'
    ;

sequence_items
    :
    | scalar {
        var a interface{} = cfglex
        a.(*cfgLex).Append($1.val)
      }
    | sequence_items ',' scalar {
        var a interface{} = cfglex
        a.(*cfgLex).Append($3.val)
      }
    | sequence_items ',' {}
    ;

map
    : '{' {
        var a interface{} = cfglex
        a.(*cfgLex).SetInMap(true)
	  }
	  map_entries
	  '}' {
        var a interface{} = cfglex
        a.(*cfgLex).SetInMap(false)
	  }
    ;

map_entries
    :
    | scalar ':' scalar {
        var a interface{} = cfglex
        a.(*cfgLex).Put($1.val, $3.val)
      }
    | map_entries ',' scalar ':' scalar {
        var a interface{} = cfglex
        a.(*cfgLex).Put($3.val, $5.val)
      }
	| map_entries ',' {}
    ;

scalar
    : SCALAR
    ;

%%

