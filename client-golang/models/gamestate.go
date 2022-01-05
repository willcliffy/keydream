package models

type State int

const (
	State_Disconnected State = iota
	State_LobbyConnected
	State_WorldConnecting
	State_WorldConnected
)
