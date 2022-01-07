package models

type State int

const (
	State_LobbyDisconnected State = iota
	State_LobbyConnecting
	State_LobbyConnected
	State_WorldConnecting
	State_WorldConnected
)
