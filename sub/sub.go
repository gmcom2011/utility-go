package sub

import (
	"rsc.io/quote"
	"fmt"
	"github.com/gmcom2011/utility-go/sub/sub_sub"
)

func Sub() {
	sub_sub.SubSub()
    fmt.Println("Hello SubSub")
    fmt.Println("Hello Sub",quote.Go())
}
