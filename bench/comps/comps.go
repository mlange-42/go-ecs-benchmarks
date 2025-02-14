package comps

import "github.com/akmonengine/volt"

const (
	PositionId = iota
	VelocityId
	C1Id
	C2Id
	C3Id
	C4Id
	C5Id
	C6Id
	C7Id
	C8Id
	C9Id
	C10Id
)

// Position component
type Position struct {
	X float64
	Y float64
}

func (c Position) GetComponentId() volt.ComponentId {
	return PositionId
}

// Velocity component
type Velocity struct {
	X float64
	Y float64
}

func (c Velocity) GetComponentId() volt.ComponentId {
	return VelocityId
}

type C1 struct {
	X float64
	Y float64
}

func (c C1) GetComponentId() volt.ComponentId {
	return C1Id
}

type C2 struct {
	X float64
	Y float64
}

func (c C2) GetComponentId() volt.ComponentId {
	return C2Id
}

type C3 struct {
	X float64
	Y float64
}

func (c C3) GetComponentId() volt.ComponentId {
	return C3Id
}

type C4 struct {
	X float64
	Y float64
}

func (c C4) GetComponentId() volt.ComponentId {
	return C4Id
}

type C5 struct {
	X float64
	Y float64
}

func (c C5) GetComponentId() volt.ComponentId {
	return C5Id
}

type C6 struct {
	X float64
	Y float64
}

func (c C6) GetComponentId() volt.ComponentId {
	return C6Id
}

type C7 struct {
	X float64
	Y float64
}

func (c C7) GetComponentId() volt.ComponentId {
	return C7Id
}

type C8 struct {
	X float64
	Y float64
}

func (c C8) GetComponentId() volt.ComponentId {
	return C8Id
}

type C9 struct {
	X float64
	Y float64
}

func (c C9) GetComponentId() volt.ComponentId {
	return C9Id
}

type C10 struct {
	X float64
	Y float64
}

func (c C10) GetComponentId() volt.ComponentId {
	return C10Id
}
