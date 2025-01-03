package domain

import "github.com/google/uuid"

type TerminalType uint8

const (
	TerminalAirplane TerminalType = 1 + iota
	TerminalBus
	TerminalShip
	TerminalTrain
)

func (t TerminalType) IsValid() bool {
	return t >= TerminalAirplane && t <= TerminalTrain
}

type TransportType uint8

const (
	TransportAirplane TransportType = 1 + iota
	TransportBus
	TransportShip
	TransportTrain
)

func (t TransportType) IsValid() bool {
	return t >= TransportAirplane && t <= TransportTrain
}

type (
	TerminalUUID = uuid.UUID
	RouteUUID    = uuid.UUID
)

func NilUUID() uuid.UUID {
	return uuid.Nil
}

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

type TerminalFilter struct {
	ID   TerminalUUID
	Name string
	Type TerminalType
}

type RouteFilter struct {
	ID            RouteUUID
	Source        Terminal
	Destination   Terminal
	TransportType TransportType
}
