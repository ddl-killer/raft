package raft

import "time"

const (
	TimeoutUnit                      = time.Millisecond
	HeartbeatTimeout                 = 50
	ElectionTimeoutMin               = 150
	ElectionTimeoutMax               = 300
	MockUnreliableRpcFailureDuration = 120
	MockUnreliableRpcDelayMin        = 60
	MockUnreliableRpcDelayMax        = 70
	MockUnreliableRpcLatencyMin      = 1
	MockUnreliableRpcLatencyMax      = 5
)

const (
	MockUnreliableRpcFailureRate = 0.1
	MockUnreliableRpcDelayRate   = (1 - MockUnreliableRpcFailureRate) * 0.2
)

type CMState int

const (
	Follower CMState = iota
	Candidate
	Leader
	Dead
)

func (s CMState) String() string {
	switch s {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	case Dead:
		return "Dead"
	default:
		panic("invalid state")
	}
}
