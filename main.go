package main

import (
	"fmt"
	"io/ioutil"

	"./glTF_"
	//"github.com/sturfeeinc/glTF/model"
)

func main() {
	bs, err := ioutil.ReadFile("./lucifer/Lucifer_finish.gltf")
	if err != nil {

	}
	var t glTF_.GLTF

	t.Compiler(bs)

	fmt.Println(t.Buffers[0])

	return
}
