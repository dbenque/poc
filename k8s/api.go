package k8s

import (
	"math/rand"
	"poc/k8s/runtime"
)

type Pod struct {
	Fieldstr     string
	Fieldint     int
	Fieldruntime k8sRuntime.RuntimeStruct // A field that we don't want
}

func NewPod() Pod {
	return Pod{getRandomStr(), rand.Intn(10), k8sRuntime.NewRuntimeStruct()}
}

var (
	strs = [...]string{
		"shoppingPod",
		"bookingPod",
		"mngPod",
	}
)

func getRandomStr() string {

	return strs[rand.Intn(len(strs))]

}
