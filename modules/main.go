package main

import (
	"fmt"
	"mymodule/helper"
	"mymodule/helper/rest"
)

/*
fonksiyonlarımızı küçük harfle başlatırsak dışarıdan erişim sağlanamaz.
helper.Helper1() şeklinde erişim sağlanabilir.
ama helper.helper1() şeklinde erişim sağlanamaz.
package içindeki fonksiyonlarımızı dışarıdan erişim sağlamak için büyük harfle başlatmalıyız.
package scope'unda erişim sağlamak için küçük harfle başlatmalıyız.
*/

func main() {
	fmt.Println("Hello, World!")
	helper.Helper1()
	rest.Rest1()
}
