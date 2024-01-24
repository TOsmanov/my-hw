package printer

import (
	"github.com/TOsmanov/my-hw/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) string {
	var str string
	for i := 0; i < len(staff); i++ {
		str += staff[i].String()
	}
	return str
}
