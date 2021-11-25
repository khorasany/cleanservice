package enum

type Role uint

const (
	_ Role = iota
	Admin
	Manager
	Basic
)
