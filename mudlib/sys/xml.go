package mudlib_sys

type XmlTagIdx int

const (
	XML_TAG_NAME       XmlTagIdx = 0 /* Name of the current tag */
	XML_TAG_ATTRIBUTES XmlTagIdx = 1 /* Atttributes of the current tag */
	XML_TAG_CONTENTS   XmlTagIdx = 2 /* Contents of the current tag */
	XML_TAG_SIZE       int       = int(XML_TAG_CONTENTS) + 1
)
