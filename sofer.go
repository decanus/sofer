package sofer

import (
	"sync"

	"github.com/decanus/bureka/dht"
	"github.com/decanus/bureka/dht/state"

	"github.com/decanus/sofer/internal"
)

// @TODO: One of the problems I currently see is credential handling.
// For a node to join the group it seems as though they need credentials in the paper.
// I would argue that this is a security vulnerability, the tree should be buildable without
// needing to have the credentials. The credentials however should be needed when disseminating messages.
// This means that the passed credentials should for example be a list of pubkeys that can send
// Or some ZK proof scheme as mentioned in the ZK GROUP paper from Signal.
// Credentials are therefore not credentials but verification data for when credentials are used.
// So to send a message, if it is a list of pubkeys some unique signature should be used, this is probably not 100% safe.

type Credentials []byte

type groupID string

type group struct {
	parent      state.Peer
	children    []state.Peer
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
	groups           map[groupID]*group

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

// JoinGroup joins a specific group.
func (s *Sofer) JoinGroup(id []byte, peer state.Peer) {
	gid := groupID(id)

	g := s.groups[gid]
	if g == nil {
		//g := &group{
		//	parent:      nil,
		//	children:    make([]state.Peer, 0),
		//	credentials: nil,
		//}
		//
		//s.groups[gid] = g
		// @todo create the group
	}

	g.children = append(g.children, peer)
}

// CreateGroup creates a group with the specific access credentials.
func (s *Sofer) CreateGroup(credentials []byte, id []byte) {
	s.Lock()
	defer s.Unlock()

	gid := groupID(id)
	_, ok := s.groups[gid]
	if ok {
		return
	}

	s.groups[gid] = &group{
		parent:      nil,
		children:    make([]state.Peer, 0),
		credentials: credentials,
	}
}
