package jwt

type JwtSecret struct {
	phrase string
}

func NewJwtSecret(phrase string) JwtSecret {
	return JwtSecret{phrase}
}
