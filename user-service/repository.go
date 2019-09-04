package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/jobin212/shippy/user-service/proto/user"
)

type repository interface {
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
	Create(iser *pb.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
}
