package internal

import (
	"github.com/decanus/bureka/dht/state"
	"github.com/decanus/bureka/pb"
)

type sofer interface {
	CreateGroup(credentials, id []byte)
}

type Receiver struct {
	s sofer
}

func New(s sofer) *Receiver {
	return &Receiver{s: s}
}

func (r *Receiver) Deliver(msg *pb.Message) {
	panic("implement me")
}

func (r *Receiver) Forward(msg *pb.Message, target state.Peer) bool {
	panic("implement me")
}

func (r *Receiver) Heartbeat(id state.Peer) {
	panic("implement me")
}
