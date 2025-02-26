package domain

type Mono struct {
	ID     uint64 `gorm:"primaryKey"`
	Nim    string // Student ID
	Nama   string
	Phone  string
	Alamat string // Address
	Email  string
}
