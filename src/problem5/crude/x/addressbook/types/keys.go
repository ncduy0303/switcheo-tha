package types

const (
	// ModuleName defines the module name
	ModuleName = "addressbook"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_addressbook"
)

var (
	ParamsKey = []byte("p_addressbook")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ContactKey      = "Contact/value/"
	ContactCountKey = "Contact/count/"
)
