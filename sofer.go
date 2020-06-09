package sofer

import (
	"sync"

	"github.com/decanus/bureka/dht"
	"github.com/decanus/bureka/dht/state"

	"github.com/decanus/sofer/internal"
)

type Credentials []byte

type GroupID string

type group struct {
	parent   state.Peer
	children []state.Peer

	credentials Credentials
}

// GroupCredentials is a class that allows for various credential schemes
type GroupCredentials interface {
	IsAuthorized(group, credentials []byte) bool
}

// Sofer implements the SCRIBE protocol and interacts with the Pastry DHT.
type Sofer struct {
	sync.RWMutex

	groupCredentials GroupCredentials
	groups           map[GroupID]*group

	receiver *internal.Receiver
	dht      *dht.DHT
}

func New(dht *dht.DHT, credentials GroupCredentials) *Sofer {
	s := &Sofer{
		dht:              dht,
		groupCredentials: credentials,
	}

	r := internal.New(s)

	dht.AddApplication("@TODO", r)

	return s
}

// CreateGroup creates a group with the specific access credentials.
func (s *Sofer) CreateGroup(credentials []byte, id []byte) {
	s.Lock()
	defer s.Unlock()

	gid := GroupID(id)

	_, ok := s.groups[gid]
	if ok {
		return
	}

	s.groups[GroupID(id)] = &group{
		parent:      nil,
		children:    make([]state.Peer, 0),
		credentials: credentials,
	}
}
