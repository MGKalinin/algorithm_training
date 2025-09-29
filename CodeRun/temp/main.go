package main

import (
	"bufio"
	"fmt"
	"os"
)

// Граф из списка маршрутов
func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int // n stop, m route
	fmt.Fscan(in, &n, &m)

	fmt.Println(n, m)

	//anya := make([]int, n)
	//boris := make([]int, m)
	//
	//for i := range n {
	//	fmt.Fscan(in, &anya[i])
	//}
	//
	//for i := range m {
	//	fmt.Fscan(in, &boris[i])
	//}
	//fmt.Fprintln(out, len(notUniq))
	//for _, v := range notUniq {
	//	fmt.Fprint(out, v, " ")
	//}
	//fmt.Fprint(out, "\n")
	//
	//fmt.Fprintln(out, len(uniqAnya))
	//for _, v := range uniqAnya {
	//	fmt.Fprint(out, v, " ")
	//}
	//fmt.Fprint(out, "\n")
	//

}
