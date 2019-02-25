// 文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "stressed"
	strSpllit := strings.Split(str, "")
	var buffer bytes.Buffer

	for i := len(strSpllit) - 1; i >= 0; i-- {
		buffer.WriteString(strSpllit[i])
	}
	fmt.Println(buffer.String())
}
