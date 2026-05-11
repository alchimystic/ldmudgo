package mudlib_sys

type FunctionFlag int

const (
	NAME_INHERITED     FunctionFlag = 0x80000000
	TYPE_MOD_STATIC    FunctionFlag = 0x40000000
	TYPE_MOD_NO_MASK   FunctionFlag = 0x20000000
	TYPE_MOD_PRIVATE   FunctionFlag = 0x10000000
	TYPE_MOD_PUBLIC    FunctionFlag = 0x08000000
	TYPE_MOD_VARARGS   FunctionFlag = 0x04000000
	TYPE_MOD_VIRTUAL   FunctionFlag = 0x02000000
	TYPE_MOD_PROTECTED FunctionFlag = 0x01000000
	TYPE_MOD_XVARARGS  FunctionFlag = 0x00800000
	TYPE_MOD_NOSAVE    FunctionFlag = TYPE_MOD_STATIC
	NAME_CROSS_DEFINED FunctionFlag = 0x00080000
	NAME_HIDDEN        FunctionFlag = 0x00020000
	NAME_PROTOTYPE     FunctionFlag = 0x00010000
	NAME_UNDEFINED     FunctionFlag = 0x00008000
	NAME_TYPES_LOST    FunctionFlag = 0x00004000
)

type ReturnFunctionFlag int

const (
	RETURN_FUNCTION_NAME    ReturnFunctionFlag = 0x01
	RETURN_FUNCTION_FLAGS   ReturnFunctionFlag = 0x02
	RETURN_FUNCTION_TYPE    ReturnFunctionFlag = 0x04
	RETURN_FUNCTION_LPCTYPE ReturnFunctionFlag = 0x20
	RETURN_FUNCTION_NUMARG  ReturnFunctionFlag = 0x08
	RETURN_FUNCTION_ARGTYPE ReturnFunctionFlag = 0x10
	RETURN_VARIABLE_VALUE   ReturnFunctionFlag = 0x08
	RETURN_FUNCTION_MASK    ReturnFunctionFlag = 0x2f
	RETURN_VARIABLE_MASK    ReturnFunctionFlag = 0x2f
)

type FunctionExistsFlag int

const (
	FEXISTS_PROGNAME FunctionExistsFlag = (0)
	FEXISTS_FILENAME FunctionExistsFlag = (1)
	FEXISTS_LINENO   FunctionExistsFlag = (2)
	FEXISTS_NUMARG   FunctionExistsFlag = (3)
	FEXISTS_TYPE     FunctionExistsFlag = (4)
	FEXISTS_FLAGS    FunctionExistsFlag = (5)
	FEXISTS_ALL      FunctionExistsFlag = (3)
)
