// Code generated by goyacc - DO NOT EDIT.

package easylang

import __yyfmt__ "fmt"

var _context = newContext()

type yySymType struct {
	yys          int
	value        float64
	str          string
	expr         expression
	descriptions []string
	arguments    []expression
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault  = 57370
	yyEofCode  = 57344
	AND        = 57358
	COLONEQUAL = 57351
	COMMA      = 57354
	DIVIDE     = 57368
	EQ         = 57359
	EQUALS     = 57349
	GE         = 57362
	GT         = 57361
	ID         = 57346
	LE         = 57364
	LPAREN     = 57352
	LT         = 57363
	MINUS      = 57365
	NE         = 57360
	NOT        = 57356
	NUM        = 57347
	OR         = 57357
	PARAMEQUAL = 57350
	PLUS       = 57366
	RPAREN     = 57353
	SEMI       = 57355
	STRING     = 57348
	TIMES      = 57367
	UNARY      = 57369
	yyErrCode  = 57345

	yyMaxDepth = 200
	yyTabOfs   = -42
)

var (
	yyXLAT = map[int]int{
		57365: 0,  // MINUS (50x)
		57354: 1,  // COMMA (40x)
		57355: 2,  // SEMI (36x)
		57357: 3,  // OR (34x)
		57352: 4,  // LPAREN (32x)
		57347: 5,  // NUM (32x)
		57353: 6,  // RPAREN (32x)
		57346: 7,  // ID (31x)
		57356: 8,  // NOT (29x)
		57348: 9,  // STRING (29x)
		57358: 10, // AND (28x)
		57359: 11, // EQ (26x)
		57360: 12, // NE (26x)
		57362: 13, // GE (24x)
		57361: 14, // GT (24x)
		57364: 15, // LE (24x)
		57363: 16, // LT (24x)
		57366: 17, // PLUS (21x)
		57380: 18, // postfix_expression (21x)
		57381: 19, // primary_expression (21x)
		57386: 20, // unary_expression (21x)
		57379: 21, // multiplicative_expression (17x)
		57368: 22, // DIVIDE (16x)
		57367: 23, // TIMES (16x)
		57371: 24, // additive_expression (15x)
		57382: 25, // relational_expression (11x)
		57344: 26, // $end (10x)
		57373: 27, // equality_expression (9x)
		57378: 28, // logical_and_expression (8x)
		57374: 29, // expression (7x)
		57377: 30, // graph_description_list (3x)
		57385: 31, // statement_suffix (3x)
		57376: 32, // graph_description (2x)
		57383: 33, // statement (2x)
		57372: 34, // argument_expression_list (1x)
		57351: 35, // COLONEQUAL (1x)
		57349: 36, // EQUALS (1x)
		57375: 37, // formula (1x)
		57350: 38, // PARAMEQUAL (1x)
		57384: 39, // statement_list (1x)
		57370: 40, // $default (0x)
		57345: 41, // error (0x)
		57369: 42, // UNARY (0x)
	}

	yySymNames = []string{
		"MINUS",
		"COMMA",
		"SEMI",
		"OR",
		"LPAREN",
		"NUM",
		"RPAREN",
		"ID",
		"NOT",
		"STRING",
		"AND",
		"EQ",
		"NE",
		"GE",
		"GT",
		"LE",
		"LT",
		"PLUS",
		"postfix_expression",
		"primary_expression",
		"unary_expression",
		"multiplicative_expression",
		"DIVIDE",
		"TIMES",
		"additive_expression",
		"relational_expression",
		"$end",
		"equality_expression",
		"logical_and_expression",
		"expression",
		"graph_description_list",
		"statement_suffix",
		"graph_description",
		"statement",
		"argument_expression_list",
		"COLONEQUAL",
		"EQUALS",
		"formula",
		"PARAMEQUAL",
		"statement_list",
		"$default",
		"error",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {37, 1},
		2:  {39, 1},
		3:  {39, 2},
		4:  {33, 4},
		5:  {33, 4},
		6:  {33, 10},
		7:  {33, 2},
		8:  {31, 1},
		9:  {31, 2},
		10: {30, 2},
		11: {30, 3},
		12: {32, 1},
		13: {19, 1},
		14: {19, 1},
		15: {19, 1},
		16: {19, 3},
		17: {18, 1},
		18: {18, 4},
		19: {34, 1},
		20: {34, 3},
		21: {20, 1},
		22: {20, 2},
		23: {20, 2},
		24: {21, 1},
		25: {21, 3},
		26: {21, 3},
		27: {24, 1},
		28: {24, 3},
		29: {24, 3},
		30: {25, 1},
		31: {25, 3},
		32: {25, 3},
		33: {25, 3},
		34: {25, 3},
		35: {27, 1},
		36: {27, 3},
		37: {27, 3},
		38: {28, 1},
		39: {28, 3},
		40: {29, 1},
		41: {29, 3},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [79][]uint8{
		// 0
		{54, 4: 50, 48, 7: 46, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 27: 59, 60, 47, 33: 45, 37: 43, 39: 44},
		{26: 42},
		{54, 4: 50, 48, 7: 46, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 41, 59, 60, 47, 33: 120},
		{40, 4: 40, 40, 7: 40, 40, 40, 26: 40},
		{29, 29, 29, 29, 84, 10: 29, 29, 29, 29, 29, 29, 29, 29, 22: 29, 29, 35: 106, 105, 38: 107},
		// 5
		{1: 99, 97, 87, 30: 98, 96},
		{28, 28, 28, 28, 6: 28, 10: 28, 28, 28, 28, 28, 28, 28, 28, 22: 28, 28},
		{27, 27, 27, 27, 6: 27, 10: 27, 27, 27, 27, 27, 27, 27, 27, 22: 27, 27},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 27: 59, 60, 94},
		{25, 25, 25, 25, 6: 25, 10: 25, 25, 25, 25, 25, 25, 25, 25, 22: 25, 25},
		// 10
		{21, 21, 21, 21, 6: 21, 10: 21, 21, 21, 21, 21, 21, 21, 21, 22: 21, 21},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 93},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 92},
		{18, 18, 18, 18, 6: 18, 10: 18, 18, 18, 18, 18, 18, 18, 18, 22: 18, 18},
		{15, 15, 15, 15, 6: 15, 10: 15, 15, 15, 15, 15, 15, 15, 15, 22: 76, 75},
		// 15
		{73, 12, 12, 12, 6: 12, 10: 12, 12, 12, 12, 12, 12, 12, 72},
		{1: 7, 7, 7, 6: 7, 10: 7, 7, 7, 70, 68, 69, 67},
		{1: 4, 4, 4, 6: 4, 10: 4, 64, 65},
		{1: 2, 2, 2, 6: 2, 10: 61},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 27: 63},
		// 20
		{29, 29, 29, 29, 84, 6: 29, 10: 29, 29, 29, 29, 29, 29, 29, 29, 22: 29, 29},
		{1: 3, 3, 3, 6: 3, 10: 3, 64, 65},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 83},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 66},
		{1: 5, 5, 5, 6: 5, 10: 5, 5, 5, 70, 68, 69, 67},
		// 25
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 82},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 81},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 80},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 71},
		{73, 8, 8, 8, 6: 8, 10: 8, 8, 8, 8, 8, 8, 8, 72},
		// 30
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 79},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 74},
		{13, 13, 13, 13, 6: 13, 10: 13, 13, 13, 13, 13, 13, 13, 13, 22: 76, 75},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 78},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 77},
		// 35
		{16, 16, 16, 16, 6: 16, 10: 16, 16, 16, 16, 16, 16, 16, 16, 22: 16, 16},
		{17, 17, 17, 17, 6: 17, 10: 17, 17, 17, 17, 17, 17, 17, 17, 22: 17, 17},
		{14, 14, 14, 14, 6: 14, 10: 14, 14, 14, 14, 14, 14, 14, 14, 22: 76, 75},
		{73, 9, 9, 9, 6: 9, 10: 9, 9, 9, 9, 9, 9, 9, 72},
		{73, 10, 10, 10, 6: 10, 10: 10, 10, 10, 10, 10, 10, 10, 72},
		// 40
		{73, 11, 11, 11, 6: 11, 10: 11, 11, 11, 11, 11, 11, 11, 72},
		{1: 6, 6, 6, 6: 6, 10: 6, 6, 6, 70, 68, 69, 67},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 27: 59, 60, 86, 34: 85},
		{1: 90, 6: 89},
		{1: 23, 3: 87, 6: 23},
		// 45
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 27: 59, 88},
		{1: 1, 1, 1, 6: 1, 10: 61},
		{24, 24, 24, 24, 6: 24, 10: 24, 24, 24, 24, 24, 24, 24, 24, 22: 24, 24},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 27: 59, 60, 91},
		{1: 22, 3: 87, 6: 22},
		// 50
		{19, 19, 19, 19, 6: 19, 10: 19, 19, 19, 19, 19, 19, 19, 19, 22: 19, 19},
		{20, 20, 20, 20, 6: 20, 10: 20, 20, 20, 20, 20, 20, 20, 20, 22: 20, 20},
		{3: 87, 6: 95},
		{26, 26, 26, 26, 6: 26, 10: 26, 26, 26, 26, 26, 26, 26, 26, 22: 26, 26},
		{35, 4: 35, 35, 7: 35, 35, 35, 26: 35},
		// 55
		{34, 4: 34, 34, 7: 34, 34, 34, 26: 34},
		{1: 103, 102},
		{7: 101, 32: 100},
		{1: 32, 32},
		{1: 30, 30},
		// 60
		{33, 4: 33, 33, 7: 33, 33, 33, 26: 33},
		{7: 101, 32: 104},
		{1: 31, 31},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 27: 59, 60, 118},
		{54, 4: 50, 48, 7: 62, 53, 49, 18: 52, 51, 55, 56, 24: 57, 58, 27: 59, 60, 116},
		// 65
		{4: 108},
		{5: 109},
		{1: 110},
		{5: 111},
		{1: 112},
		// 70
		{5: 113},
		{6: 114},
		{2: 115},
		{36, 4: 36, 36, 7: 36, 36, 36, 26: 36},
		{1: 99, 97, 87, 30: 98, 117},
		// 75
		{37, 4: 37, 37, 7: 37, 37, 37, 26: 37},
		{1: 99, 97, 87, 30: 98, 119},
		{38, 4: 38, 38, 7: 38, 38, 38, 26: 38},
		{39, 4: 39, 39, 7: 39, 39, 39, 26: 39},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("'%c'", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 41

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 4:
		{
			yyVAL.expr = AssignmentExpression(_context, yyS[yypt-3].str, yyS[yypt-1].expr)
			_context.addOutput(yyVAL.expr.VarName(), yyS[yypt-0].descriptions, 0, 0)
		}
	case 5:
		{
			yyVAL.expr = AssignmentExpression(_context, yyS[yypt-3].str, yyS[yypt-1].expr)
			_context.addNotOutputVar(yyVAL.expr.VarName(), yyS[yypt-0].descriptions, 0, 0)
		}
	case 6:
		{
			yyVAL.expr = ParamExpression(_context, yyS[yypt-9].str, yyS[yypt-6].value, yyS[yypt-4].value, yyS[yypt-2].value)
		}
	case 7:
		{
			varName := _context.newAnonymousVarName()
			yyVAL.expr = AssignmentExpression(_context, varName, yyS[yypt-1].expr)
			_context.addOutput(varName, yyS[yypt-0].descriptions, 0, 0)
		}
	case 9:
		{
			yyVAL.descriptions = yyS[yypt-1].descriptions
		}
	case 10:
		{
			if !IsValidDescription(yyS[yypt-0].str) {
				lexer, _ := yylex.(*yylexer)
				_context.addError(BadGraphDescError(lexer.lineno, lexer.column, yyS[yypt-0].str))
			}
			yyVAL.descriptions = append(yyVAL.descriptions, yyS[yypt-0].str)
		}
	case 11:
		{
			if !IsValidDescription(yyS[yypt-0].str) {
				lexer, _ := yylex.(*yylexer)
				_context.addError(BadGraphDescError(lexer.lineno, lexer.column, yyS[yypt-0].str))
			}
			yyVAL.descriptions = append(yyS[yypt-2].descriptions, yyS[yypt-0].str)
		}
	case 12:
		{
			yyVAL.str = yyS[yypt-0].str
		}
	case 13:
		{
			expr := _context.defined(yyS[yypt-0].str)
			if expr == nil {
				expr = _context.definedParam(yyS[yypt-0].str)
			}
			if expr != nil {
			} else if funcName, ok := noArgFuncMap[yyS[yypt-0].str]; ok {
				expr = FunctionExpression(_context, funcName, nil)
			} else {
				lexer, _ := yylex.(*yylexer)
				_context.addError(UndefinedVarError(lexer.lineno, lexer.column, yyS[yypt-0].str))
				expr = ErrorExpression(_context, yyS[yypt-0].str)
			}
			yyVAL.expr = expr
		}
	case 14:
		{
			yyVAL.expr = ConstantExpression(_context, yyS[yypt-0].value)
		}
	case 16:
		{
			yyVAL.expr = yyS[yypt-1].expr
		}
	case 17:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 18:
		{
			if _, ok := funcMap[yyS[yypt-3].str]; !ok {
				lexer, _ := yylex.(*yylexer)
				_context.addError(UndefinedFunctionError(lexer.lineno, lexer.column, yyS[yypt-3].str))
				yyVAL.expr = ErrorExpression(_context, yyS[yypt-3].str)
			} else {
				yyVAL.expr = FunctionExpression(_context, yyS[yypt-3].str, yyS[yypt-1].arguments)
			}
		}
	case 19:
		{
			yyVAL.arguments = []expression{yyS[yypt-0].expr}
		}
	case 20:
		{
			yyVAL.arguments = append(yyVAL.arguments, yyS[yypt-0].expr)
		}
	case 21:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 22:
		{
			yyVAL.expr = UnaryExpression(_context, yyS[yypt-1].str, yyS[yypt-0].expr)
		}
	case 23:
		{
			yyVAL.expr = UnaryExpression(_context, yyS[yypt-1].str, yyS[yypt-0].expr)
		}
	case 24:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 25:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 26:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 27:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 28:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 29:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 30:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 31:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 32:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 33:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 34:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 35:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 36:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 37:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 38:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 39:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 40:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 41:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
