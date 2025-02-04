package core

//RaftValue collect a map of Raft index and the list of ETCD members associated
type RaftValue map[uint64][]string

// IsBetween returns true if value is in between min and max, min and max included
func IsBetween(value, min, max uint64) bool {
	if value >= min && value <= max {
		return true
	}
	return false
}
