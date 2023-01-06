package game

type IRotatable interface {
	Angle() (int, error)
	SetAngle(newAngle int) error
	RotationAngle() (int, error)
	Rotate(int) error
}

type rotate struct {
	rotatable IRotatable
}

func NewRotate(rotatable IRotatable) *rotate {
	return &rotate{rotatable}
}

func (r *rotate) Execute() {
	currentAngle, _ := r.rotatable.Angle()
	addingAngle, _ := r.rotatable.RotationAngle()
	r.rotatable.SetAngle(currentAngle + addingAngle)
}
