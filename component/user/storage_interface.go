package user

type IStorage interface {
	GetUserById(id string) MUser
}

var storageInstance IStorage

func SetStorage(storage IStorage) {
	storageInstance = storage
}

func getStorage() IStorage {
	return storageInstance
}
