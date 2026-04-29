package config

type Auth struct {
	Secret []byte
}

func LoadAuth() *Auth {
	return &Auth{
		Secret: []byte(KeyAuthSecret.GetValueDefault("auth-secret")),
	}
}
