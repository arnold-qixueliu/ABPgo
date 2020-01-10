// main_test
package main

import (
	"fmt"
	"testing"
)

func Test_dd(t *testing.T) {
	t.Log("ddd777")

	for k, v := range mapStruct {
		fmt.Println(k + ":" + v.Name())
	}

}
