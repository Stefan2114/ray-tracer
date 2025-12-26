package mpi

import (
	mpi "github.com/sbromberger/gompi"
)

var World *mpi.Communicator

func MpiInit() {
	mpi.Start(false)
	World = mpi.NewCommunicator(nil)
}

func MpiStop() {
	mpi.Stop()
}
