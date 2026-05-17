package thasher

type IHasher interface {
	Hash(string) string
}
