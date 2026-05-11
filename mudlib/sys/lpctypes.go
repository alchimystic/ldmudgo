package mudlib_sys

type CompileTimeType int

const (
	TYPE_UNKNOWN      CompileTimeType = 0
	TYPE_NUMBER       CompileTimeType = 1
	TYPE_STRING       CompileTimeType = 2
	TYPE_VOID         CompileTimeType = 3
	TYPE_OBJECT       CompileTimeType = 4
	TYPE_MAPPING      CompileTimeType = 5
	TYPE_FLOAT        CompileTimeType = 6
	TYPE_ANY          CompileTimeType = 7
	TYPE_CLOSURE      CompileTimeType = 8
	TYPE_SYMBOL       CompileTimeType = 9
	TYPE_QUOTED_ARRAY CompileTimeType = 10
	TYPE_STRUCT       CompileTimeType = 11
	TYPE_BYTES        CompileTimeType = 12
	TYPE_LWOBJECT     CompileTimeType = 13
	TYPE_COROUTINE    CompileTimeType = 14
	TYPE_LPCTYPE      CompileTimeType = 15
	TYPE_MOD_POINTER  CompileTimeType = 0x0040
)

type RuntimeType int

const (
	T_INVALID      RuntimeType = 0x0
	T_LVALUE       RuntimeType = 0x1
	T_NUMBER       RuntimeType = 0x2
	T_STRING       RuntimeType = 0x3
	T_POINTER      RuntimeType = 0x4
	T_OBJECT       RuntimeType = 0x5
	T_MAPPING      RuntimeType = 0x6
	T_FLOAT        RuntimeType = 0x7
	T_CLOSURE      RuntimeType = 0x8
	T_SYMBOL       RuntimeType = 0x9
	T_QUOTED_ARRAY RuntimeType = 0xa
	T_STRUCT       RuntimeType = 0xb
	T_BYTES        RuntimeType = 0xc
	T_LWOBJECT     RuntimeType = 0xd
	T_COROUTINE    RuntimeType = 0xe
	T_PYTHON       RuntimeType = 0xf
	T_LPCTYPE      RuntimeType = 0x10
)

type ClosureType int

const (
	CLOSURE_LFUN           ClosureType = 0
	CLOSURE_IDENTIFIER     ClosureType = 2
	CLOSURE_PRELIMINARY    ClosureType = 3
	CLOSURE_BOUND_LAMBDA   ClosureType = 4
	CLOSURE_LAMBDA         ClosureType = 5
	CLOSURE_UNBOUND_LAMBDA ClosureType = 6
	CLOSURE_OPERATOR       ClosureType = (0xe800)
	CLOSURE_EFUN           ClosureType = (0xf000)
	CLOSURE_SIMUL_EFUN     ClosureType = (0xf800)
)

func ClosureIsLFun(x ClosureType) bool {
	//return x == ((x) & ~1) == 0
	return x == CLOSURE_LFUN
}

func ClosureIsIdentifier(x ClosureType) bool {
	return x == CLOSURE_IDENTIFIER
}

func ClosureIsBoundLambda(x ClosureType) bool {
	return x == CLOSURE_BOUND_LAMBDA
}

func ClosureIsLambda(x ClosureType) bool {
	return x == CLOSURE_LAMBDA
}

func ClosureIsUnboundLambda(x ClosureType) bool {
	return x == CLOSURE_UNBOUND_LAMBDA
}

func ClosureIsSimulExtFun(x ClosureType) bool {
	return x == CLOSURE_SIMUL_EFUN
}

func ClosureIsExtFun(x ClosureType) bool {
	return x == CLOSURE_EFUN
}

func ClosureIsOperator(x ClosureType) bool {
	return x == CLOSURE_OPERATOR
}
