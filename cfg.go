//line cfg.y:2
package confish

import __yyfmt__ "fmt"

//line cfg.y:2
//line cfg.y:6
type cfgSymType struct {
	yys int
	val string
}

const SECTION_NAME = 57346
const ATTR_NAME = 57347
const SCALAR = 57348

var cfgToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"':'",
	"','",
	"'{'",
	"'}'",
	"'['",
	"']'",
	"SECTION_NAME",
	"ATTR_NAME",
	"SCALAR",
}
var cfgStatenames = [...]string{}

const cfgEofCode = 1
const cfgErrCode = 2
const cfgMaxDepth = 200

//line cfg.y:112

//line yacctab:1
var cfgExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const cfgNprod = 26
const cfgPrivate = 57344

var cfgTokenNames []string
var cfgStates []string

const cfgLast = 36

var cfgAct = [...]int{

	21, 19, 22, 18, 3, 10, 25, 15, 3, 30,
	24, 29, 11, 2, 5, 34, 31, 13, 26, 23,
	9, 20, 17, 16, 27, 14, 28, 12, 8, 7,
	6, 32, 33, 4, 1, 35,
}
var cfgPact = [...]int{

	-1000, -2, -1000, -1000, 8, -1000, -6, 5, -1000, -1000,
	-1000, -1000, 13, -5, -1000, -1000, -1000, -1000, -10, -1000,
	1, -1000, -1000, -10, -1000, -10, 4, 12, -1000, -1000,
	-10, -10, 11, -1000, -10, -1000,
}
var cfgPgo = [...]int{

	0, 34, 13, 33, 30, 29, 28, 27, 25, 23,
	22, 21, 0, 19, 18,
}
var cfgR1 = [...]int{

	0, 1, 1, 3, 5, 2, 4, 4, 4, 7,
	6, 8, 8, 8, 9, 11, 11, 11, 11, 13,
	10, 14, 14, 14, 14, 12,
}
var cfgR2 = [...]int{

	0, 0, 2, 0, 0, 6, 0, 2, 2, 0,
	4, 1, 1, 1, 3, 0, 1, 3, 2, 0,
	4, 0, 3, 5, 2, 1,
}
var cfgChk = [...]int{

	-1000, -1, -2, 10, -3, 6, -4, -5, -6, -2,
	11, 7, -7, 4, -8, 12, -9, -10, 8, 6,
	-11, -12, 12, -13, 9, 5, -14, -12, -12, 7,
	5, 4, -12, -12, 4, -12,
}
var cfgDef = [...]int{

	1, -2, 2, 3, 0, 6, 4, 0, 7, 8,
	9, 5, 0, 0, 10, 11, 12, 13, 15, 19,
	0, 16, 25, 21, 14, 18, 0, 0, 17, 20,
	24, 0, 0, 22, 0, 23,
}
var cfgTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 8, 3, 9, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 6, 3, 7,
}
var cfgTok2 = [...]int{

	2, 3, 10, 11, 12,
}
var cfgTok3 = [...]int{
	0,
}

var cfgErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	cfgDebug        = 0
	cfgErrorVerbose = false
)

type cfgLexer interface {
	Lex(lval *cfgSymType) int
	Error(s string)
}

type cfgParser interface {
	Parse(cfgLexer) int
	Lookahead() int
}

type cfgParserImpl struct {
	lookahead func() int
}

func (p *cfgParserImpl) Lookahead() int {
	return p.lookahead()
}

func cfgNewParser() cfgParser {
	p := &cfgParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const cfgFlag = -1000

func cfgTokname(c int) string {
	if c >= 1 && c-1 < len(cfgToknames) {
		if cfgToknames[c-1] != "" {
			return cfgToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func cfgStatname(s int) string {
	if s >= 0 && s < len(cfgStatenames) {
		if cfgStatenames[s] != "" {
			return cfgStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func cfgErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !cfgErrorVerbose {
		return "syntax error"
	}

	for _, e := range cfgErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + cfgTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := cfgPact[state]
	for tok := TOKSTART; tok-1 < len(cfgToknames); tok++ {
		if n := base + tok; n >= 0 && n < cfgLast && cfgChk[cfgAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if cfgDef[state] == -2 {
		i := 0
		for cfgExca[i] != -1 || cfgExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; cfgExca[i] >= 0; i += 2 {
			tok := cfgExca[i]
			if tok < TOKSTART || cfgExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if cfgExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += cfgTokname(tok)
	}
	return res
}

func cfglex1(lex cfgLexer, lval *cfgSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = cfgTok1[0]
		goto out
	}
	if char < len(cfgTok1) {
		token = cfgTok1[char]
		goto out
	}
	if char >= cfgPrivate {
		if char < cfgPrivate+len(cfgTok2) {
			token = cfgTok2[char-cfgPrivate]
			goto out
		}
	}
	for i := 0; i < len(cfgTok3); i += 2 {
		token = cfgTok3[i+0]
		if token == char {
			token = cfgTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = cfgTok2[1] /* unknown char */
	}
	if cfgDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", cfgTokname(token), uint(char))
	}
	return char, token
}

func cfgParse(cfglex cfgLexer) int {
	return cfgNewParser().Parse(cfglex)
}

func (cfgrcvr *cfgParserImpl) Parse(cfglex cfgLexer) int {
	var cfgn int
	var cfglval cfgSymType
	var cfgVAL cfgSymType
	var cfgDollar []cfgSymType
	_ = cfgDollar // silence set and not used
	cfgS := make([]cfgSymType, cfgMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	cfgstate := 0
	cfgchar := -1
	cfgtoken := -1 // cfgchar translated into internal numbering
	cfgrcvr.lookahead = func() int { return cfgchar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		cfgstate = -1
		cfgchar = -1
		cfgtoken = -1
	}()
	cfgp := -1
	goto cfgstack

ret0:
	return 0

ret1:
	return 1

cfgstack:
	/* put a state and value onto the stack */
	if cfgDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", cfgTokname(cfgtoken), cfgStatname(cfgstate))
	}

	cfgp++
	if cfgp >= len(cfgS) {
		nyys := make([]cfgSymType, len(cfgS)*2)
		copy(nyys, cfgS)
		cfgS = nyys
	}
	cfgS[cfgp] = cfgVAL
	cfgS[cfgp].yys = cfgstate

cfgnewstate:
	cfgn = cfgPact[cfgstate]
	if cfgn <= cfgFlag {
		goto cfgdefault /* simple state */
	}
	if cfgchar < 0 {
		cfgchar, cfgtoken = cfglex1(cfglex, &cfglval)
	}
	cfgn += cfgtoken
	if cfgn < 0 || cfgn >= cfgLast {
		goto cfgdefault
	}
	cfgn = cfgAct[cfgn]
	if cfgChk[cfgn] == cfgtoken { /* valid shift */
		cfgchar = -1
		cfgtoken = -1
		cfgVAL = cfglval
		cfgstate = cfgn
		if Errflag > 0 {
			Errflag--
		}
		goto cfgstack
	}

cfgdefault:
	/* default state action */
	cfgn = cfgDef[cfgstate]
	if cfgn == -2 {
		if cfgchar < 0 {
			cfgchar, cfgtoken = cfglex1(cfglex, &cfglval)
		}

		/* look through exception table */
		xi := 0
		for {
			if cfgExca[xi+0] == -1 && cfgExca[xi+1] == cfgstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			cfgn = cfgExca[xi+0]
			if cfgn < 0 || cfgn == cfgtoken {
				break
			}
		}
		cfgn = cfgExca[xi+1]
		if cfgn < 0 {
			goto ret0
		}
	}
	if cfgn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			cfglex.Error(cfgErrorMessage(cfgstate, cfgtoken))
			Nerrs++
			if cfgDebug >= 1 {
				__yyfmt__.Printf("%s", cfgStatname(cfgstate))
				__yyfmt__.Printf(" saw %s\n", cfgTokname(cfgtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for cfgp >= 0 {
				cfgn = cfgPact[cfgS[cfgp].yys] + cfgErrCode
				if cfgn >= 0 && cfgn < cfgLast {
					cfgstate = cfgAct[cfgn] /* simulate a shift of "error" */
					if cfgChk[cfgstate] == cfgErrCode {
						goto cfgstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if cfgDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", cfgS[cfgp].yys)
				}
				cfgp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if cfgDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", cfgTokname(cfgtoken))
			}
			if cfgtoken == cfgEofCode {
				goto ret1
			}
			cfgchar = -1
			cfgtoken = -1
			goto cfgnewstate /* try again in the same state */
		}
	}

	/* reduction by production cfgn */
	if cfgDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", cfgn, cfgStatname(cfgstate))
	}

	cfgnt := cfgn
	cfgpt := cfgp
	_ = cfgpt // guard against "declared and not used"

	cfgp -= cfgR2[cfgn]
	// cfgp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if cfgp+1 >= len(cfgS) {
		nyys := make([]cfgSymType, len(cfgS)*2)
		copy(nyys, cfgS)
		cfgS = nyys
	}
	cfgVAL = cfgS[cfgp+1]

	/* consult goto table to find next state */
	cfgn = cfgR1[cfgn]
	cfgg := cfgPgo[cfgn]
	cfgj := cfgg + cfgS[cfgp].yys + 1

	if cfgj >= cfgLast {
		cfgstate = cfgAct[cfgg]
	} else {
		cfgstate = cfgAct[cfgj]
		if cfgChk[cfgstate] != -cfgn {
			cfgstate = cfgAct[cfgg]
		}
	}
	// dummy call; replaced with literal code
	switch cfgnt {

	case 3:
		cfgDollar = cfgS[cfgpt-1 : cfgpt+1]
		//line cfg.y:21
		{
			var a interface{} = cfglex
			a.(*cfgLex).Push(0, cfgDollar[1].val)
		}
	case 4:
		cfgDollar = cfgS[cfgpt-4 : cfgpt+1]
		//line cfg.y:28
		{
			var a interface{} = cfglex
			a.(*cfgLex).Pop()
		}
	case 9:
		cfgDollar = cfgS[cfgpt-1 : cfgpt+1]
		//line cfg.y:43
		{
			var a interface{} = cfglex
			a.(*cfgLex).Push(3, cfgDollar[1].val)
		}
	case 10:
		cfgDollar = cfgS[cfgpt-4 : cfgpt+1]
		//line cfg.y:50
		{
			var a interface{} = cfglex
			a.(*cfgLex).Pop()
		}
	case 11:
		cfgDollar = cfgS[cfgpt-1 : cfgpt+1]
		//line cfg.y:57
		{
			var a interface{} = cfglex
			a.(*cfgLex).AttrVal(cfgDollar[1].val)
		}
	case 16:
		cfgDollar = cfgS[cfgpt-1 : cfgpt+1]
		//line cfg.y:72
		{
			var a interface{} = cfglex
			a.(*cfgLex).Append(cfgDollar[1].val)
		}
	case 17:
		cfgDollar = cfgS[cfgpt-3 : cfgpt+1]
		//line cfg.y:76
		{
			var a interface{} = cfglex
			a.(*cfgLex).Append(cfgDollar[3].val)
		}
	case 18:
		cfgDollar = cfgS[cfgpt-2 : cfgpt+1]
		//line cfg.y:80
		{
		}
	case 19:
		cfgDollar = cfgS[cfgpt-1 : cfgpt+1]
		//line cfg.y:84
		{
			var a interface{} = cfglex
			a.(*cfgLex).SetInMap(true)
		}
	case 20:
		cfgDollar = cfgS[cfgpt-4 : cfgpt+1]
		//line cfg.y:89
		{
			var a interface{} = cfglex
			a.(*cfgLex).SetInMap(false)
		}
	case 22:
		cfgDollar = cfgS[cfgpt-3 : cfgpt+1]
		//line cfg.y:97
		{
			var a interface{} = cfglex
			a.(*cfgLex).Put(cfgDollar[1].val, cfgDollar[3].val)
		}
	case 23:
		cfgDollar = cfgS[cfgpt-5 : cfgpt+1]
		//line cfg.y:101
		{
			var a interface{} = cfglex
			a.(*cfgLex).Put(cfgDollar[3].val, cfgDollar[5].val)
		}
	case 24:
		cfgDollar = cfgS[cfgpt-2 : cfgpt+1]
		//line cfg.y:105
		{
		}
	}
	goto cfgstack /* stack new state and value */
}
