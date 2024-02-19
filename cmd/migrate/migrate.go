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
	case args[1] == "migrate":
		if err := migrateDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	case args[1] == "truncate":
		if err := truncateTable(db, args[2]); err != nil {
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
		Password: "secretADMIN01.",
	})
	db.Create(&domain.Admin{
		Uuid:   uuid.NewString(),
		JK:     "LAKI-LAKI",
		UserID: 1,
	})

	return nil
}
func migrateDatabase(db *gorm.DB) error {
	for _, model := range registry.RegisterModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}

	return nil
}

func truncateTable(db *gorm.DB, table string) error{
	return db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)).Error
}

func seederData(db *gorm.DB) error {

	user1 := domain.User{
		Nama:     "John Doe",
		Username: "johndoess",
		Password: "123",
	}

	user2 := domain.User{
		Nama:     "John Doe",
		Username: "doejhon",
		Password: "123",
	}

	if err := db.Create(&user1).Error; err != nil {
		return err
	}

	if err := db.Create(&user2).Error; err != nil {
		return err
	}

	teacher1 := domain.Teacher{
		Uuid:          "teacher1",
		Nuptk:         "1234567890",
		StatusPegawai: "AKTIF",
		Nip:           "987654321223",
		JK:            "LAKI-LAKI",
		TempatLahir:   "Jakarta",
		TanggalLahir:  "1990-01-01",
		NoHp:          "081234567890",
		Jabatan:       "Guru",
		TotalJtm:      "40",
		Alamat:        "Jl. Contoh No. 123",
		UserID:        user1.ID,
	}

	teacher2 := domain.Teacher{
		Uuid:          "teacher2",
		Nuptk:         "1234567890",
		StatusPegawai: "AKTIF",
		Nip:           "98765432122",
		JK:            "LAKI-LAKI",
		TempatLahir:   "Jakarta",
		TanggalLahir:  "1990-01-01",
		NoHp:          "081234567890",
		Jabatan:       "Guru",
		TotalJtm:      "40",
		Alamat:        "Jl. Contoh No. 123",
		UserID:        user2.ID,
	}

	if err := db.Create(&teacher1).Error; err != nil {
		return err
	}

	class := domain.Class{
		Uuid:      "class1",
		Name:      "VII-A",
		TeacherID: teacher1.ID,
	}

	if err := db.Create(&class).Error; err != nil {
		return err
	}

	subject := domain.Subject{
		Uuid: "subject1",
		Name: "Fisika",
		Teachers: []domain.Teacher{
			teacher1,
			teacher2,
		},
	}

	if err := db.Create(&subject).Error; err != nil {
		return err
	}

	student := domain.Student{
		Nama:         "Siswa",
		Uuid:         "student1",
		NIS:          "12345678",
		JK:           "LAKI-LAKI",
		TempatLahir:  "Gorontalo",
		TanggalLahir: "2000-01-01",
		Alamat:       "Alamat Siswa",
		TanggalMasuk: "26-01-2000",
		ClassID:      class.ID,
	}

	if err := db.Create(&student).Error; err != nil {
		return err
	}

	sy := domain.SchoolYear{Uuid: "sy1", Name: "2023"}

	if err := db.Create(&sy).Error; err != nil {
		return err
	}

	schedule := domain.Schedule{
		Uuid:         "schedule1",
		Day:          "SENIN",
		Start:        "09:00",
		End:          "10:00",
		ClassID:      class.ID,
		SubjectID:    subject.ID,
		SchoolYearID: sy.ID,
	}

	if err := db.Create(&schedule).Error; err != nil {
		return err
	}

	if err := db.Create(&domain.Absence{
		Uuid:       "absence1",
		Status:     "HADIR",
		StudentID:  student.ID,
		ScheduleID: schedule.ID,
	}).Error; err != nil {
		return err
	}

	return nil
}
