package server

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func Test_GetNewsContent(t *testing.T) {
	publishTime := time.Now().Add(-24 * time.Hour)
	//publishTime := time.Now()
	g := Gocn{}
	_, contents := g.GetNewsContent(publishTime)
	fmt.Println(contents)
	fmt.Println(strings.Join(contents, ""))
}
