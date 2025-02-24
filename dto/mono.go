package dto

type MonoDTO struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Nim    string
	Nama   string
	Phone  string
	Alamat string
	Email  string
}
