package main

import (
	"encoding/json"
	"fmt"
)

type Snapshot struct {
	SaveInterval   int64 `default:"10"`
	DeleteInterval int64 `default:"500" json:"delete_interval"`
}

type ListResponse struct {
	List []int `json:"list"`
}

func main() {
	s := Snapshot{}
	j := `{"saveinterval":250,"delete_interval":500}`

	err := json.Unmarshal([]byte(j), &s)

	if err != nil {
		panic(err)
	}

	fmt.Println(s)

	lp := ListResponse{nil}

	if len(lp.List) == 0 {
		lp.List = make([]int, 0)
	}

	bs, _ := json.Marshal(lp)
	fmt.Println(string(bs))

}
