package sub

import (
	"rsc.io/quote"
	"fmt"
	"github.com/gmcom2011/utility-go/sub_sub"
)

func Sub() {
    fmt.Println("Hello SubSub", sub_sub.SubSub())
    fmt.Println("Hello Sub",quote.Go())
}
