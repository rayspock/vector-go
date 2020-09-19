package core

//Vector ... X, Y, Z value in three dimensional space
type Vector struct {
	X, Y, Z float64
}

//Dot ... Vector A Dot Vector B
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

//Cross ... Vector A Cross Vector B
func (a Vector) Cross(b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}
