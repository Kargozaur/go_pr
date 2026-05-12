package phasher

type IHasher interface {
	Hash(password string) (string, error)
	VerifyPassword(password, passwordHash string) bool
}
