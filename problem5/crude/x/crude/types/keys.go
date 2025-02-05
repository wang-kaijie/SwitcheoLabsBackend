package types

const (
	// ModuleName defines the module name
	ModuleName = "crude"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_crude"

	// UserKey is used to uniquely identify users within the system.
	// It will be used as the beginning of the key for each post, followed by their unique ID
	UserKey = "User/value/"

	// This key will be used to keep track of the ID of the latest user added to the store.
	UserCountKey = "User/count/"
)

var (
	ParamsKey = []byte("p_crude")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
