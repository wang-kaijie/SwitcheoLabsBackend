package types

const (
	// ModuleName defines the module name
	ModuleName = "userblock"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_userblock"
)

var (
	ParamsKey = []byte("p_userblock")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
