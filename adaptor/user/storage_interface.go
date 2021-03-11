package user

type MUser struct{}

type IStorage interface {
	GetUserById(id string) MUser
}

var storageInstance IStorage

func SetStorage(storage IStorage) {
	storageInstance = storage
}

func GetStorage() IStorage {
	return storageInstance
}
