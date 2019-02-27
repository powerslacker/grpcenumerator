package generate

import (
	"bytes"
	"strconv"
	"strings"
)

/*
String transforms a list of strings into a protobuf `enum`. The `suffix`
will be appeneded to each item in the given `list`. If no `suffix` is desired,
use an empty string. The `name` given will be the name of the enum type
produced.
*/
func String(name string, list []string, suffix string) string {
	buf := new(bytes.Buffer)
	buf.WriteString("enum " + name + " {\n")
	if len(suffix) > 0 {
		suffix = "_" + suffix
	}
	for i, v := range list {
		v += suffix
		v = strings.ToUpper(strings.Join(strings.Split(v, " "), "_"))
		line := "\t"
		line += strings.ToUpper(v)
		line += " = "
		line += strconv.Itoa(i)
		line += ";"
		line += "\n"
		buf.WriteString(line)
	}
	buf.WriteString("}")
	return buf.String()
}
