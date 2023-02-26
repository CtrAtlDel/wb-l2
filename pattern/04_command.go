package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Command interface
type Command interface {
	Execute() string
}

// Command 1 implementation
type MethodOn struct {
	reciver *Reciver
}

func (m *MethodOn) Execute() string {
	return m.reciver.On()
}

// Command 2 implementation
type MethodOff struct {
	reciver *Reciver
}

func (m *MethodOff) Execute() string {
	return m.reciver.Off()
}

// Struct and methods for Reciver implmentation
type Reciver struct{} //Reciver implementation (Приемние)

func (r *Reciver) On() string {
	return "Swich On"
}

func (r *Reciver) Off() string {
	return "Swich Off"
}

// Invoker (a thing wich will be invoke our commands)
type Invoker struct { // Implemet Command interface
	commands []Command
}

func (i *Invoker) CreateCommands(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) Execute() string{
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}