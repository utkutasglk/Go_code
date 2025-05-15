package intermediate

import (
	"fmt"
	"time"
)


func main() {

	// Mon Jan 2 15:04:05 MST 2006
	layout := "2006-01-02T15:05:05Z07:00"
	str :="2024-07-04T14:30:18Z"

	t, err := time.Parse(layout,str)

	if err != nil {
		fmt.Println("Error parsing time: ",err)
		return
	}
	fmt.Println(t)

	str1 := "Jul 03, 2024 03:18 PM"
	layout1 := "Jan 02, 2006 03:04 PM"

	t1, err1 := time.Parse(layout1,str1)

	if err1 != nil {
		fmt.Println("Error parsing time: ",err1)
		return
	}
	fmt.Println(t1)

}