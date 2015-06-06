package k8sRuntime

import "math/rand"

type RuntimeStruct struct {
	R1 string
	R2 int
}

func NewRuntimeStruct() RuntimeStruct {
	return RuntimeStruct{getRandomStr(), rand.Intn(100)}
}

var (
	strs = [...]string{
		"banana",
		"orange",
		"apple",
	}
)

func getRandomStr() string {

	return strs[rand.Intn(len(strs))]

}
