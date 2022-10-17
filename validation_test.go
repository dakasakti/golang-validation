package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

// instansiasi
var (
	validate = validator.New()
)

func TestValidationVariable(t *testing.T) {
	name := ""
	err := validate.Var(name, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariable(t *testing.T) {
	password := "rahasia"
	confirmPassword := "beda"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationMultipleTag(t *testing.T) {
	name := "dakasakti99"

	err := validate.Var(name, "required,alpha")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTagParameter(t *testing.T) {
	name := "daka"

	err := validate.Var(name, "required,alpha,min=5,max=25")
	if err != nil {
		fmt.Println(err.Error())
	}
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=10,max=255"`
}

func TestValidationStruct(t *testing.T) {
	login := LoginRequest{
		Email:    "dakasakti.id@gmail.com",
		Password: "admin",
	}

	err := validate.Struct(login)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type RegisterRequest struct {
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=5,max=255"`
	ConfirmPassword string `validate:"required,min=5,max=255,eqfield=Password"`
}

func TestValidationCrossField(t *testing.T) {
	register := RegisterRequest{
		Email:           "dakasakti.id@gmail.com",
		Password:        "password",
		ConfirmPassword: "beda",
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type RegisterNestedRequest struct {
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=5,max=255"`
	ConfirmPassword string `validate:"required,min=5,max=255,eqfield=Password"`
	Address         Address
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
}

func TestValidationNestedStruct(t *testing.T) {
	register := RegisterNestedRequest{
		Email:           "dakasakti.id@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type RegisterCollectionRequest struct {
	Email           string    `validate:"required,email"`
	Password        string    `validate:"required,min=5,max=255"`
	ConfirmPassword string    `validate:"required,min=5,max=255,eqfield=Password"`
	Addresses       []Address `validate:"dive"`
}

func TestValidationCollectionStruct(t *testing.T) {
	register := RegisterCollectionRequest{
		Email:           "dakasakti.id@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
		Addresses: []Address{
			{
				Street: "Jl. Ahmad Yani",
				City:   "",
			},
			{
				Street: "",
				City:   "Palembang",
			},
		},
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type RegisterBasicCollectionRequest struct {
	Email           string    `validate:"required,email"`
	Password        string    `validate:"required,min=5,max=255"`
	ConfirmPassword string    `validate:"required,min=5,max=255,eqfield=Password"`
	Addresses       []Address `validate:"dive"`
	Hobbies         []string  `validate:"dive,required,min=5,max=25"`
}

func TestValidationBasicCollectionStruct(t *testing.T) {
	register := RegisterBasicCollectionRequest{
		Email:           "dakasakti.id@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
		Addresses: []Address{
			{
				Street: "Jl. Ahmad Yani",
				City:   "Palembang",
			},
			{
				Street: "Jl. Sudirman",
				City:   "Palembang",
			},
		},
		Hobbies: []string{
			"Design", "Editing", "Apa",
		},
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type School struct {
	Name string `validate:"required"`
}

type RegisterMapRequest struct {
	Email           string            `validate:"required,email"`
	Password        string            `validate:"required,min=5,max=255"`
	ConfirmPassword string            `validate:"required,min=5,max=255,eqfield=Password"`
	Addresses       []Address         `validate:"dive"`
	Hobbies         []string          `validate:"dive,required,min=5,max=25"`
	Schools         map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
}

func TestValidationMapStruct(t *testing.T) {
	register := RegisterMapRequest{
		Email:           "dakasakti.id@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
		Addresses: []Address{
			{
				Street: "Jl. Ahmad Yani",
				City:   "Palembang",
			},
			{
				Street: "Jl. Sudirman",
				City:   "Palembang",
			},
		},
		Hobbies: []string{
			"Design", "Editing",
		},
		Schools: map[string]School{
			"SMA": {Name: ""},
		},
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type RegisterBasicMapRequest struct {
	Email           string            `validate:"required,email"`
	Password        string            `validate:"required,min=5,max=255"`
	ConfirmPassword string            `validate:"required,min=5,max=255,eqfield=Password"`
	Addresses       []Address         `validate:"dive"`
	Hobbies         []string          `validate:"dive,required,min=5,max=25"`
	Schools         map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	Wallet          map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
}

func TestValidationBasicMapStruct(t *testing.T) {
	register := RegisterBasicMapRequest{
		Email:           "dakasakti.id@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
		Addresses: []Address{
			{
				Street: "Jl. Ahmad Yani",
				City:   "Palembang",
			},
			{
				Street: "Jl. Sudirman",
				City:   "Palembang",
			},
		},
		Hobbies: []string{
			"Design", "Editing",
		},
		Schools: map[string]School{
			"SMA": {Name: "SMKN 1 Palembang"},
		},
		Wallet: map[string]int{
			"BCA": 0,
		},
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type Seller struct {
	Username string `validate:"varchar,lowercase"`
	Store    string `validate:"varchar"`
	Name     string `validate:"varchar"`
}

func TestValidationTagAlias(t *testing.T) {
	validate.RegisterAlias("varchar", "required,min=5,max=25")

	register := Seller{
		Username: "admin",
		Store:    "Kita",
		Name:     "Dakasakti",
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func validUsername(field validator.FieldLevel) bool {
	value := field.Field().String()

	if value != strings.ToLower(value) {
		return false
	}

	if len(value) < 5 {
		return false
	}

	return true
}

type Users struct {
	Username string `validate:"username"`
}

func TestValidationCustom(t *testing.T) {
	validate.RegisterValidation("username", validUsername)

	register := Users{
		Username: "apa",
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func validPIN(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		fmt.Println(err.Error())
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

type Profile struct {
	PIN string `validate:"pin=6"`
}

func TestValidationCustomParameter(t *testing.T) {
	validate.RegisterValidation("pin", validPIN)

	register := Profile{
		PIN: "12345",
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type LoginPipeRequest struct {
	EmailPhone string `validate:"required,email|numeric"`
	Password   string `validate:"required,min=5,max=255"`
}

func TestValidationOrRule(t *testing.T) {

	register := LoginPipeRequest{
		EmailPhone: "dakasakti.id@gmail.com",
		Password:   "admin",
	}

	err := validate.Struct(register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}
