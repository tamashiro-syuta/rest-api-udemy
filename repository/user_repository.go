package repository

import (
	"rest-api-udemy/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	// NOTE: ユーザーのポインタを引数に取り、エラーを返す
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

// NOTE: このコンストラクタ-でDBのインスタンスをリポジトリに渡すことで、DBからリポジトリに依存関係を注入(依存の向き: リポジトリ -> DB)
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// NOTE: DBからemailが一致するユーザーを取得し、引数で受け取ったuserのポインタのデータを取得したユーザーのデータで更新
	// NOTE: エラーがあればそのまま返す
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	// NOTE: DBにユーザーを作成し、作成に成功した場合、引数で受け取ったuserのポインタのデータを作成したユーザーのデータで更新
	// NOTE: エラーがあればそのまま返す
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
