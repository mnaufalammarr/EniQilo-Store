package services

import (
	"EniQilo/entities"
	"EniQilo/repositories"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type StaffService interface {
	Create(signupRequest entities.SignUpRequest) (entities.Staff, error)
	Login(loginRequest entities.SignInRequest) (string, error)
	FindByID(id int) (entities.Staff, error)
}

type staffService struct {
	staffRepository repositories.StaffRepository
	userRepository  repositories.UserRepository
}

func NewStaffService(staffRepository repositories.StaffRepository, userRepository repositories.UserRepository) *staffService {
	return &staffService{staffRepository, userRepository}
}

func (s *staffService) Create(signupRequest entities.SignUpRequest) (entities.Staff, error) {
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), 10)

	if err != nil {
		return entities.Staff{}, err
	}
	fmt.Println("req", signupRequest)

	user, err := s.userRepository.FindByPhone(signupRequest.Phone)
	fmt.Println("user", user)
	if err != nil {
		return entities.Staff{}, err
	}
	//fmt.Println(isPhoneExist)
	if user.Id != 0 && user.Role == true {
		fmt.Println("hitted error phone redudant")
		return entities.Staff{}, errors.New("Phone ALREADY EXIST")
	}
	userRequest := entities.User{
		Name:  signupRequest.Name,
		Phone: signupRequest.Phone,
		Role:  true,
	}
	fmt.Println("usReq", userRequest)

	newUser, err := s.userRepository.CreateUser(userRequest)
	fmt.Println("newUser", newUser)
	fmt.Println("newUserErr", err)

	staffRequest := entities.StaffRequast{
		UserId:   newUser,
		Password: string(hash),
	}

	fmt.Println("staRq", staffRequest)

	newStaff, err := s.staffRepository.Create(staffRequest)
	fmt.Println("newStaf", newStaff)
	fmt.Println("newStafErr", err)

	staffNew := entities.Staff{
		UserID: entities.User{
			Name:  userRequest.Name,
			Phone: userRequest.Phone,
			Role:  userRequest.Role,
		},
		Password: staffRequest.Password,
	}

	return staffNew, err
}

func (s *staffService) Login(loginRequest entities.SignInRequest) (string, error) {
	//get staff
	user, err := s.userRepository.FindByPhone(loginRequest.Phone)
	fmt.Println("userLogin", user.Id)

	if err != nil {
		fmt.Println("error", err)
		return "", err

	} else if user.Id == 0 {
		return "",
			errors.New("Invalid phone or password")
	}

	staff, err := s.staffRepository.FindByUserId(user.Id)
	fmt.Println("staffLogin", staff)

	if err != nil {
		if staff.Id == 0 {
			return "", errors.New("Invalid phone or password")
		}
		fmt.Println("error", err)
		return "", err

	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(loginRequest.Password))

	fmt.Println("errlogin", err.Error())
	if err != nil {
		return "", err
	}

	//sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": staff.Id,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	fmt.Println("tokeLogin", tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *staffService) FindByID(id int) (entities.Staff, error) {
	staff, err := s.staffRepository.FindById(id)
	if err != nil {
		return entities.Staff{}, err
	}

	// Jika staff tidak ditemukan, kembalikan error
	if staff.Id == 0 {
		return entities.Staff{}, errors.New("staff not found")
	}

	return staff, nil
}
