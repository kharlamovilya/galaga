package game

type spaceship struct {
	position Vector2D
	angle    int
}

func NewSpaceship(position Vector2D) spaceship {
	return spaceship{position, 0}
}

func (s *spaceship) Position() (Vector2D, error) {
	return s.position, nil
}

func (s *spaceship) SetPosition(newPosition Vector2D) error {
	s.position = newPosition
	return nil
}

func (s *spaceship) Velocity() (Vector2D, error) {
	return Vector2D{0, 0}, nil
}

func (s *spaceship) Move(velocity Vector2D) error {
	s.position = s.position.Plus(velocity)
	return nil
}

func (s *spaceship) Angle() (int, error) {
	return s.angle, nil
}

func (s *spaceship) SetAngle(newAngle int) error {
	s.angle = newAngle
	return nil
}

func (s *spaceship) RotationAngle() (int, error) {
	return 0, nil
}

func (s *spaceship) Rotate(newAngle int) error {
	s.SetAngle((newAngle + s.angle) % 360)
	return nil
}
