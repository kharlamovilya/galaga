package game

type spaceship struct {
	position Vector2D
	velocity Vector2D
	angle    int
	fuel     int
}

func NewSpaceship(position, velocity Vector2D, angle, fuel int) spaceship {
	return spaceship{position: position, fuel: fuel}
}

// Движение

func (s *spaceship) Position() (Vector2D, error) {
	return s.position, nil
}

func (s *spaceship) SetPosition(newPosition Vector2D) error {
	s.position = newPosition
	return nil
}

func (s *spaceship) Velocity() (Vector2D, error) {
	return s.velocity, nil
}

func (s *spaceship) Move() error {
	vel, _ := s.Velocity()
	s.position = s.position.Plus(vel)
	return nil
}

// Поворот

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

// Сжигание топлива

func (s *spaceship) Fuel() (int, error) {
	return s.fuel, nil
}

func (s *spaceship) SetFuel(newFuelLevel int) error {
	s.fuel = newFuelLevel
	return nil
}
