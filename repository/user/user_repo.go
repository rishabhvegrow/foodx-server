package user

import (
    "os"
	"time"

	"gorm.io/gorm"
	"github.com/golang-jwt/jwt"

    "foodx-server/domain"
)

var secret = os.Getenv("JWT_SECRET_KEY")
var jwtSecretKey = []byte(secret)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (user_repo *UserRepository) GetUsers()(*[]domain.User, error){
    db := user_repo.db
	var users []domain.User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
	return &users, nil
}

func (user_repo *UserRepository) GetUser(userID any)(*domain.User, error){
    db := user_repo.db
    var user domain.User
    if err := db.First(&user, userID).Error; err != nil {
        return nil, err
    }

	return &user, nil
}

func (user_repo *UserRepository) CraeteUser(user domain.User)(*domain.User, error){
    db := user_repo.db
    if err := db.Create(&user).Error; err != nil {
        return nil, err
    }
	return &user, nil
}

func (user_repo *UserRepository) DeleteUser(userID any)(error){
    db := user_repo.db
	var user domain.User
    if err:= db.Delete(&user, userID).Error; err != nil {
        return err
    }
	return nil
}

func (user_repo *UserRepository) UserLogin(email string, password string) (*domain.UserLoginResponse, error) {
    db := user_repo.db
    var user domain.User
    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }

    token, err := generateToken(user.ID)
    if err != nil {
        return nil, err
    }

    response := &domain.UserLoginResponse{
        Token: token,
        Email: user.Email,
        Name:  user.Name,
    }

    return response, nil
}

func generateToken(userID uint) (string, error) {
    tokenClaims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
    tokenString, err := token.SignedString(jwtSecretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
