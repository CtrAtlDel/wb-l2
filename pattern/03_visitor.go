package pattern

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Visitor_pattern
*/
type Visitor interface { // interface for visitors
	VisitAlbuquerque(*Albuquerque) string
	VisitIowa(*Iowa) string
	VisitTexas(*Texas) string
}

// Implementation for interface visitor
type Tourist struct{}

func (t *Tourist) VisitAlbuquerque(a *Albuquerque) string { return a.GetAlbuquerqueLozung() }

func (t *Tourist) VisitIowa(i *Iowa) string { return i.GetIowaLozung() }

func (t *Tourist) VisitTexas(txs *Texas) string { return txs.GetTexasLozung() }

// Place which can be visited by Visitor
type Place interface {
	Accept(v Visitor) string
}

// Places with can be visited by visitor
type City struct {
	places []Place
}

// Add place
func (c *City) AddPlace(p Place) {
	c.places = append(c.places, p)
}

// Create a string with place and their lozung
func (c *City) Accept(v Visitor) string {
	var res string
	for _, j := range c.places {
		res += j.Accept(v)
	}
	return res
}

// Structs with methods
type Albuquerque struct{}

func (a *Albuquerque) GetAlbuquerqueLozung() string {
	return "I am Albuquerque"
}

func (a *Albuquerque) Accept(v Visitor) string {
	return v.VisitAlbuquerque(a)
}

type Iowa struct{}

func (i *Iowa) GetIowaLozung() string {
	return "I am Iowa"
}

func (i *Iowa) Accept(v Visitor) string {
	return v.VisitIowa(i)
}

type Texas struct{}

func (t *Texas) GetTexasLozung() string {
	return "I am Texas"
}

func (t *Texas) Accept(v Visitor) string {
	return v.VisitTexas(t)
}
