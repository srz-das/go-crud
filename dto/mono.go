package dto

type MonoDTO struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement"`
	Nim    string
	Nama   string
	Phone  string
	Alamat string
	Email  string
}
