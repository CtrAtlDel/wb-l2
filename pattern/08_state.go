package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type IMobileAlert interface {
	Alert() string
}

type MobileAlert struct {
	state IMobileAlert
}

func (m *MobileAlert) Alert() string {
	return m.state.Alert()
}

func CreateMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MAVibration{}}
}

func (m *MobileAlert) SetState(s IMobileAlert) {
	m.state = s
}

type MAVibration struct{} // implement IMobileAlert

func (m *MAVibration) Alert() string {
	return "Bzhzhzhhzhzhzhhzhzhzhhzhz"
}

type MASound struct{} // implement IMobileAlert

func (m *MASound) Alert() string {
	return "* Some sound * "
}
