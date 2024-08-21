package usecase

import (
	"os"
	"rest-api-udemy/model"
	"rest-api-udemy/repository"
	"rest-api-udemy/validator"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// NOTE: ユースケースのインターフェース(ユーザーユースケースはこのインターフェースを実装する)
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	// NOTE: JWTトークンを返すのでstring型を返す
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// NOTE: このコンストラクタ-でリポジトリのインターフェースをユースケースを渡すことで、リポジトリ(のインターフェース)からユースケースに依存関係を注入(依存の向き: リポジトリ -> リポジトリのインターフェース -> ユースケース)
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	// NOTE: 構造体のインスタンスを作成し、そのポインタを返す
	return &userUsecase{ur, uv}
}

// NOTE: Interfaceを満たすようにメソッドを実装
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	encryption_complexity := 10
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), encryption_complexity)
	if err != nil {
		// NOTE: エラーの場合はゼロ値のUserResponseとエラーを返す
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	// NOTE: CreateUserで作成したユーザーの情報を　レスポンス用の構造体に変換して返す
	resUser := model.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}

	// NOTE: DBに保存されているハッシュ化されたパスワードとユーザーが入力した平文のパスワードを比較
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// NOTE: JWTトークンを生成して返す
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}
