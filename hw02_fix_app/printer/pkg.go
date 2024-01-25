package printer

import (
	"fmt"
	"io"

	"github.com/TOsmanov/my-hw/hw02_fix_app/types"
)

func PrintStaff(w io.Writer, staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Fprintln(w, staff[i].String())
	}
}
