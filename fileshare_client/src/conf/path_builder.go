package conf

import (
	"strings"
)

type PathBuilder struct {
	lastIsSeparator bool
	builder         strings.Builder
}

func isSeparator(ch byte) bool {
	return ch == '/' || ch == '\\'
}

func (obj *PathBuilder) Append(pathStr string) *PathBuilder {
	count := len(pathStr)
	if count == 0 {
		return obj
	}
	defer func() {
		if count == 0 {
			return
		}
		obj.lastIsSeparator = isSeparator(pathStr[count-1])
	}()
	if obj.builder.Len() == 0 {
		obj.builder.WriteString(pathStr)
		return obj
	}
	if obj.lastIsSeparator {
		if isSeparator(pathStr[0]) {
			obj.builder.WriteString(pathStr[1:])
		} else {
			obj.builder.WriteString(pathStr)
		}
	} else {
		if isSeparator(pathStr[0]) {
			obj.builder.WriteString(pathStr)
		} else {
			obj.builder.WriteString("/")
			obj.builder.WriteString(pathStr)
		}
	}
	return obj
}

func (obj *PathBuilder) Clear() *PathBuilder {
	obj.builder.Reset()
	return obj
}

func (obj *PathBuilder) FullPath() string {
	return obj.builder.String()
}
