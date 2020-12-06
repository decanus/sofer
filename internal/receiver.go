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

	// @todo if we are forwarding a join message, and the origin is from self, we set the target as the parent of the group.

	panic("implement me")
}

func (r *Receiver) Heartbeat(id state.Peer) {
	panic("implement me")
}
