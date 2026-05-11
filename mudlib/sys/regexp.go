package mudlib_sys

type REOptionFlag int

const (
	RE_GLOBAL         REOptionFlag = 0x0001 /* Apply RE globally (if possible) */
	RE_EXCOMPATIBLE   REOptionFlag = 0x0002 /* RE is compatible with ex */
	RE_CASELESS       REOptionFlag = 0x0004
	RE_MULTILINE      REOptionFlag = 0x0008
	RE_DOTALL         REOptionFlag = 0x0010
	RE_EXTENDED       REOptionFlag = 0x0020
	RE_ANCHORED       REOptionFlag = 0x0040
	RE_DOLLAR_ENDONLY REOptionFlag = 0x0080
	RE_NOTBOL         REOptionFlag = 0x0100
	RE_NOTEOL         REOptionFlag = 0x0200
	RE_UNGREEDY       REOptionFlag = 0x0400
	RE_NOTEMPTY       REOptionFlag = 0x0800
	RE_MATCH_SUBS     REOptionFlag = 0x1000 /* Return matched subexpressions */
	RE_OMIT_DELIM     REOptionFlag = 0x1000 /* Omit the delimiters */
	RE_TRADITIONAL    REOptionFlag = 0x04000000
	RE_PCRE           REOptionFlag = 0x02000000
	RE_PACKAGE_MASK   int          = int(RE_TRADITIONAL) | int(RE_PCRE)
)
