package serverTypes

import (
	"math/rand"
	"poc/k8s"
	"poc/server"
)

type Silo struct {
	serverAPI.SiloBase
	Spec k8s.Pod
}

func NewSilo() Silo {
	return Silo{serverAPI.SiloBase{getRandomStr(), rand.Intn(10)}, k8s.NewPod()}
}

var (
	strs = [...]string{
		"dario",
		"nico",
		"david",
	}
)

func getRandomStr() string {

	return strs[rand.Intn(len(strs))]

}
