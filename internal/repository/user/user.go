package user

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/rocksus/go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db          *gorm.DB
	gcm         cipher.AEAD
	time        uint32
	memory      uint32
	parallelism uint8
	keyLen      uint32
	secret      string
}

func GetRepository(db *gorm.DB, secret string, time, memory, keyLen uint32, parallelism uint8) (Repository, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &userRepo{
		db:          db,
		gcm:         gcm,
		time:        time,
		memory:      memory,
		parallelism: parallelism,
		keyLen:      keyLen,
		secret:      secret,
	}, nil
}

func (ur *userRepo) RegisterUser(userData model.User) (model.User, error) {
	if err := ur.db.Create(&userData).Error; err != nil {
		return model.User{}, err
	}

	return userData, nil
}
func (ur *userRepo) CheckRegistered(username string) (bool, error) {
	var userData model.User

	if err := ur.db.Where(model.User{Username: username}).First(&userData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}

	return userData.ID != "", nil
}
