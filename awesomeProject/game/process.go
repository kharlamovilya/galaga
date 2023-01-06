package game

func processMove(movingObj IMovable, velocity Vector2D) error {
	err := movingObj.Move(velocity)
	return err
}

func processRotate(rotatingObj IRotatable, newAngle int) error {
	err := rotatingObj.Rotate(newAngle)
	return err
}
