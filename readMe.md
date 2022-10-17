# Golang Validation

## Baked-in Validation
https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme

### Jenis Tag Validation
- Fields
- Networks
- Strings
- Format
- Comparisons
- Other

### Multiple Tag Validation
```
"required,alpha"
```

### Tag Parameter
biasa digunakan jika ada tag validate yang mengharuskan ada isinya.
```
"min=5,max=25"
```

### Validate Struct
dengan menambahkan reflection tag di struct fieldnya dengan tag validate.
```
type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=10,max=255"`
}
```

### Validation Errors
validator package akan mengembalikan data error. jika (error == nil) artinya semua data valid dan sebaliknya.
package ini memiliki implementasi error yaitu `ValidationErrors` (alias untuk []FieldError).

### Validate Cross Field
```
type RegisterRequest struct {
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=5,max=255"`
	ConfirmPassword string `validate:"required,min=5,max=255,eqfield=Password"`
}
```

### Validate Nested Struct
secara otomatis validator package melakukan validasi terhadap field struct sesuai tag validatenya.
```
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
```

### Validate Collection
Tidak seperti tipe data struct, jika kita memiliki field dengan data collection seperti Array, Slice, atau Map, secara default validator package tidak akan melakukan validasi terhadap data didalam collection tersebut. namun, jika kita ingin melakukan validasi maka kita bisa menambahkan tag `dive`.
```
type RegisterCollectionRequest struct {
	Email           string    `validate:"required,email"`
	Password        string    `validate:"required,min=5,max=255"`
	ConfirmPassword string    `validate:"required,min=5,max=255,eqfield=Password"`
	Addresses       []Address `validate:"dive"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
}
```

### Validate Basic Collection
Bagaimana jika data collectionnya bukan struct?, misal []string. Pada kasus ini kita bisa tambahkan tag validate langsung setelah dive

```
type RegisterBasicCollectionRequest struct {
	Email           string    `validate:"required,email"`
	Password        string    `validate:"required,min=5,max=255"`
	ConfirmPassword string    `validate:"required,min=5,max=255,eqfield=Password"`
	Addresses       []Address `validate:"dive"`
	Hobbies         []string  `validate:"dive,required,min=5,max=25"`
}
```

### Validate Map
Selain collection array atau slice, kita juga bisa melakukan validate field Map.
```
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
```
### Validate Basic Map
```
type RegisterBasicMapRequest struct {
	Email           string            `validate:"required,email"`
	Password        string            `validate:"required,min=5,max=255"`
	ConfirmPassword string            `validate:"required,min=5,max=255,eqfield=Password"`
	Addresses       []Address         `validate:"dive"`
	Hobbies         []string          `validate:"dive,required,min=5,max=25"`
	Schools         map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	Wallet          map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
}
```

### Alias Tag
pada beberapa kasus, kadang kita sering menggunakan beberapa tag validate yang sama untuk field yang berbeda. validator package memiliki fitur untuk menambahkan alias, yaitu nama tag baru untuk tag lain.
```
validate.RegisterAlias("varchar", "required,min=5,max=25")

type Seller struct {
	Username string `validate:"varchar,lowercase"`
	Store    string `validate:"varchar"`
	Name     string `validate:"varchar"`
}
```

### Custom Validate
untuk yang tag tidak ada di Baked-in Validation kita bisa custom.
```
type Users struct {
	Username string `validate:"username"`
}

validate.RegisterValidation("username", validUsername)

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
```

### Custom Validate Parameter
```
type Profile struct {
	PIN string `validate:"pin=6"`
}

validate.RegisterValidation("pin", validPIN)

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
```

### OR Rule
Pada beberapa kasus, kadang kita ingin membuat kondisi or pada validate. contoh sebuah field boleh email ataupun nomor telepon.
bisa menggunakan tanda pipe (|, = OR). secara default tanda (, = AND).

```
type LoginPipeRequest struct {
	EmailPhone string `validate:"required,email|numeric"`
	Password   string `validate:"required,min=5,max=255"`
}
```

### Custom Validate Cross Field
```
field level = field.GetStructFieldOK2()
```

### Struct Level Validation
Kadang ada kasus melakukan validate butuh kombinasi lebih dari dua field. validator package mendukung pembuatan validate di level struct, namun kita perlu membuat validate function menggunakan parameter struct level.

```
validate.RegisterStructValidation()
```