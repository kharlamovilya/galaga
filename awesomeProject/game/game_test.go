package game

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type (
	mockSpaceship struct{}
)

func (s mockSpaceship) Position() (Vector2D, error) {
	return Vector2D{}, fmt.Errorf("spaceship Position Error")
}

func (s mockSpaceship) SetPosition(newPosition Vector2D) error {
	return fmt.Errorf("spaceship SetPosition Error")
}

func (s mockSpaceship) Velocity() (Vector2D, error) {
	return Vector2D{}, fmt.Errorf("spaceship Velocity Error")
}

func (s mockSpaceship) Move(velocity Vector2D) error {
	return fmt.Errorf("spaceship Move error")
}

func TestSpaceship_Move(t *testing.T) {
	t.Log("Given the need to test Spaceship's movement.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			velocity := Vector2D{-7, 3}
			initialPosition := Vector2D{12, 5}
			spaceship := NewSpaceship(initialPosition)
			processMove(&spaceship, velocity)
			require.Equal(t, Vector2D{5, 8}, spaceship.position)
		}

		testID++
		t.Logf("\tTest %d:\tWhen unreadable position argument.", testID)
		{
			s := mockSpaceship{}
			_, err := s.Position()
			require.Error(t, err)
		}

		testID++
		t.Logf("\tTest %d:\tWhen unreadable velocity argument.", testID)
		{
			s := mockSpaceship{}
			_, err := s.Velocity()
			require.Error(t, err)
		}

		testID++
		t.Logf("\tTest %d:\tWhen unchangeable position argument.", testID)
		{
			s := mockSpaceship{}
			err := s.SetPosition(Vector2D{})
			require.Error(t, err)
		}
	}
}

func (s mockSpaceship) Angle() (int, error) {
	return 0, fmt.Errorf("spaceship Angle error")
}

func (s mockSpaceship) SetAngle(newAngle int) error {
	return fmt.Errorf("spaceship Angle error")
}

func (s mockSpaceship) RotationAngle() (int, error) {
	return 0, fmt.Errorf("spaceship Angle error")
}

func (s mockSpaceship) Rotate(newAngle int) error {
	return fmt.Errorf("spaceShip Rotate error")
}

func TestSpaceship_Rotate(t *testing.T) {
	t.Log("Given the need to test Spaceship's rotation.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			rotationAngle := 90
			spaceship := NewSpaceship(Vector2D{})
			processRotate(&spaceship, rotationAngle)
			angle, _ := spaceship.Angle()
			require.Equal(t, 90, angle)
		}

		testID++
		t.Logf("\tTest %d:\tWhen unreadable angle.", testID)
		{
			spaceship := mockSpaceship{}
			_, err := spaceship.Angle()
			require.Error(t, err)
		}

		testID++
		t.Logf("\tTest %d:\tWhen unreadable rotation angle.", testID)
		{
			spaceship := mockSpaceship{}
			_, err := spaceship.RotationAngle()
			require.Error(t, err)
		}

		testID++
		t.Logf("\tTest %d:\tWhen unchangeable angle.", testID)
		{
			newAngle := -00000
			spaceship := mockSpaceship{}
			err := spaceship.SetAngle(newAngle)
			require.Error(t, err)
		}
	}
}
