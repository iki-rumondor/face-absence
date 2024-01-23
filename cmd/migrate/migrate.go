package migrate

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/registry"
	"gorm.io/gorm"
)

func ReadTerminal(db *gorm.DB, args []string) {
	switch {
	case args[1] == "fresh":
		if err := freshDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	case args[1] == "seed":
		if err := seederData(db); err != nil {
			log.Fatal(err.Error())
		}
	default:
		fmt.Println("Hello")
	}
}

func freshDatabase(db *gorm.DB) error {
	for _, model := range registry.RegisterModels() {
		if err := db.Migrator().DropTable(model.Model); err != nil {
			return err
		}
	}
	for _, model := range registry.RegisterModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}

	db.Create(&domain.User{
		Nama:     "Admin",
		Username: "admin",
		Password: "123456",
	})
	db.Create(&domain.Admin{
		Uuid:   uuid.NewString(),
		UserID: 1,
	})

	return nil
}

func seederData(db *gorm.DB) error {

	userData := domain.User{
		Nama:     "John Doe",
		Username: "johndoess",
		Password: "123",
	}

	if err := db.Create(&userData).Error; err != nil {
		return err
	}

	teacher := domain.Teacher{
		Uuid:          uuid.NewString(),
		Nuptk:         "1234567890",
		StatusPegawai: "AKTIF",
		Nip:           "98765432122",
		JK:            "Laki-laki",
		TempatLahir:   "Jakarta",
		TanggalLahir:  "1990-01-01",
		NoHp:          "081234567890",
		Jabatan:       "Guru",
		TotalJtm:      "40",
		Alamat:        "Jl. Contoh No. 123",
		UserID:        userData.ID,
	}

	if err := db.Create(&teacher).Error; err != nil {
		return err
	}

	class := domain.Class{
		Uuid:      uuid.NewString(),
		Name:      "VII-A",
		TeacherID: teacher.ID,
	}

	if err := db.Create(&class).Error; err != nil {
		return err
	}

	subject := domain.Subject{
		Uuid:      uuid.NewString(),
		Name:      "Fisika",
		TeacherID: teacher.ID,
	}

	if err := db.Create(&subject).Error; err != nil {
		return err
	}

	userStudent := domain.User{
		Nama:     "John Doe",
		Username: "student",
		Password: "123",
	}

	if err := db.Create(&userStudent).Error; err != nil {
		return err
	}

	student := domain.Student{
		Uuid:         uuid.NewString(),
		NIS:          "12345678",
		JK:           "LAKI-LAKI",
		TempatLahir:  "Gorontalo",
		TanggalLahir: "2000-01-01",
		Alamat:       "Alamat Siswa",
		UserID:       userStudent.ID,
		ClassID:      class.ID,
	}

	if err := db.Create(&student).Error; err != nil {
		return err
	}

	if err := db.Create(&domain.SchoolYear{Name: "2023"}).Error; err != nil {
		return err
	}

	return nil
}
