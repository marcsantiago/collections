package main

import (
	"fmt"
	"strings"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/collections/namedtuple"
)

var exampleData = `
ID\tMARKS\tNAME\tCLASS\n   
1\t97\tRaymond\t7\n         
2\t50\tSteven\t4\n         
3\t91\tAdrian\t9\n         
4\t72\tStewart\t5\n         
5\t80\tPeter\t6`

func main() {
	rows := strings.Split(exampleData, "\\n")
	var nt namedtuple.NamedTuple
	for i, row := range rows {
		cols := strings.Split(row, "\\t")
		if i == 0 {
			nt = namedtuple.New(cols...)
			continue
		}
		nt = nt.Add(collections.StringValues(cols).Data())
	}

	var avgMarks float64
	for i := 0; i < nt.Len(); i++ {
		avgMarks += nt.Get(i, "MARKS").Float64()
		fmt.Println(nt.Get(i, "NAME").String(), nt.Get(i, "MARKS").Float64())
	}
	avgMarks = avgMarks / float64(nt.Len())
	fmt.Println("AVG Marks", avgMarks)
}
