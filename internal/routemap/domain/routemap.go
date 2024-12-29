package domain

import "github.com/google/uuid"

type TerminalType uint8

const (
	TerminalAirplane TerminalType = 1 << iota
	TerminalBus
	TerminalShip
	TerminalTrain
)

type TransportType uint8

const (
	TransportAirplane TransportType = 1 << iota
	TransportBus
	TransportShip
	TransportTrain
)

type (
	TerminalUUID = uuid.UUID
	RouteUUID    = uuid.UUID
)

type Terminal struct {
	ID      TerminalUUID
	Name    string
	Type    TerminalType
	Country string
	State   string
	City    string
}

type Route struct {
	ID            RouteUUID
	Source        Terminal
	Destination   Terminal
	RouteNumber   uint
	TransportType TransportType
	Distance      float64
}
