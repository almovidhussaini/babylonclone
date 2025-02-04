package types

const (
	// ModuleName defines the module name
	ModuleName = "babylonclone"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_babylonclone"
)

var (
	ParamsKey = []byte("p_babylonclone")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
