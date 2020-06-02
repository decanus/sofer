package sofer

import (
	"github.com/decanus/bureka/dht"
	"github.com/decanus/bureka/dht/state"
	"github.com/decanus/bureka/pb"
)


// Sofer implements the SCRIBE protocol and interacts with the Pastry DHT.
type Sofer struct {
	dht *dht.DHT
}

func (s *Sofer) Deliver(msg *pb.Message) {
	panic("implement me")
}

func (s *Sofer) Forward(msg *pb.Message, target state.Peer) bool {
	panic("implement me")
}

func (s *Sofer) Heartbeat(id state.Peer) {
	panic("implement me")
}

