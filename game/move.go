package game

type IMovable interface {
	Position() (Vector2D, error)
	SetPosition(Vector2D) error
	Velocity() (Vector2D, error)
	Move() error
}

type move struct {
	movable IMovable
}

func NewMove(movable IMovable) *move {
	return &move{movable}
}

func (m *move) Execute() error {
	pos, _ := m.movable.Position()
	vel, _ := m.movable.Velocity()
	m.movable.SetPosition(pos.Plus(vel))
	return nil
}
