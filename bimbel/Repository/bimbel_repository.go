package repository

import (
	"UTS-LANJUTAN-ARYA/models"

	"gorm.io/gorm"
)

type BimbelRepository struct {
	db *gorm.DB
}

func NewBimbelRepo(db *gorm.DB) *BimbelRepository {
	return &BimbelRepository{db}
}

func (bimbelRepo *BimbelRepository) GetStudent() []*models.Tamp {
	var result []*models.Tamp
	data := bimbelRepo.db.Model(&models.Student{}).
		Select("students.id_student, students.name_student, students.name_school, students.class, programs.name_program").
		Joins("inner join programs on programs.id_program = students.program_id").
		Scan(&result)
	if data.Error != nil {
		return nil
	}

	return result
}

func (bimbelRepo *BimbelRepository) GetRequest(id int) (*models.Tamp, *models.ProgramSD, *models.ProgramSMP, *models.ProgramSMAIpa, *models.ProgramSMAIps) {
	err := bimbelRepo.db.First(&models.Student{}, "id_student = ?", id).Error
	if err != nil {
		return nil, nil, nil, nil, nil
	}

	var result *models.Tamp
	bimbelRepo.db.Select("students.id_student, students.name_student, students.name_school, students.class, programs.name_program").
		Joins("inner join programs on programs.id_program = students.program_id").
		Where("id_student = ?", id).Find(&models.Student{}).Scan(&result)

	var resultsd *models.ProgramSD
	bimbelRepo.db.Find(&resultsd).Where("student_id = ?", id).Scan(&resultsd)

	var resultsmp *models.ProgramSMP
	bimbelRepo.db.Find(&resultsmp).Where("student_id = ?", id).Scan(&resultsmp)

	var resultsma *models.ProgramSMAIpa
	bimbelRepo.db.Find(&resultsma).Where("student_id = ?", id).Scan(&resultsma)

	var resultsmas *models.ProgramSMAIps
	bimbelRepo.db.Find(&resultsmas).Where("student_id = ?", id).Scan(&resultsmas)

	return result, resultsd, resultsmp, resultsma, resultsmas
}

func (bimbelRepo *BimbelRepository) CreateStudent(data *models.Student) error {
	err := bimbelRepo.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) UpdateStudent(data *models.Student) error {
	err := bimbelRepo.db.First(&models.Student{}, "id_student = ?", data.IDStudent).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Model(&models.Student{}).Where("id_student = ?", data.IDStudent).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) DeleteStudent(id int) error {
	err := bimbelRepo.db.First(&models.Student{}, "id_student = ?", id).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Delete(&models.Student{}, bimbelRepo.db.Where("id_student = ?", id)).Error
	if err != nil {
		return err
	}
	return nil
}

///////////////////////// T E A C H E R ///////////////////////////////////

func (bimbelRepo *BimbelRepository) GetTeacher() []*models.Teacher {
	var result []*models.Teacher
	data := bimbelRepo.db.Find(&result)
	if data.Error != nil {
		return nil
	}

	return result
}

func (bimbelRepo *BimbelRepository) CreateTeacher(data *models.Teacher) error {
	err := bimbelRepo.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) UpdateTeacher(data *models.Teacher) error {
	err := bimbelRepo.db.First(&models.Teacher{}, "id_teacher = ?", data.IDTeacher).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Model(&models.Teacher{}).Where("id_teacher = ?", data.IDTeacher).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) DeleteTeacher(id int) error {
	err := bimbelRepo.db.First(&models.Teacher{}, "id_Teacher = ?", id).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Delete(&models.Teacher{}, bimbelRepo.db.Where("id_Teacher = ?", id)).Error
	if err != nil {
		return err
	}
	return nil
}

//////////////////////////// P R O G R A M //////////////////////////

func (bimbelRepo *BimbelRepository) GetProgram() []*models.Program {
	var result []*models.Program
	data := bimbelRepo.db.Find(&result)
	if data.Error != nil {
		return nil
	}

	return result
}

func (bimbelRepo *BimbelRepository) CreateProgram(data *models.Program) error {
	err := bimbelRepo.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) UpdateProgram(data *models.Program) error {
	err := bimbelRepo.db.First(&models.Program{}, "id_Program = ?", data.IDProgram).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Model(&models.Program{}).Where("id_Program = ?", data.IDProgram).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) DeleteProgram(id int) error {
	err := bimbelRepo.db.First(&models.Program{}, "id_Program = ?", id).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Delete(&models.Program{}, bimbelRepo.db.Where("id_Program = ?", id)).Error
	if err != nil {
		return err
	}
	return nil
}

// ///// P R O G R A M   S D ////////////////////////////////////////
func (bimbelRepo *BimbelRepository) CreateProgramSD(data *models.ProgramSD) error {
	err := bimbelRepo.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) UpdateProgramSD(data *models.ProgramSD) error {
	err := bimbelRepo.db.Find(&models.ProgramSD{}, "student_id = ?", data.StudentID).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Model(&models.ProgramSD{}).Where("student_id = ?", data.StudentID).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) DeleteProgramSD(id int) error {
	err := bimbelRepo.db.First(&models.ProgramSD{}, "student_id = ?", id).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Delete(&models.ProgramSD{}, bimbelRepo.db.Where("student_id = ?", id)).Error
	if err != nil {
		return err
	}
	return nil
}

// //////// P R O G   S M P ////////////////////////////////////////
func (bimbelRepo *BimbelRepository) CreateProgramSMP(data *models.ProgramSMP) error {
	err := bimbelRepo.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) UpdateProgramSMP(data *models.ProgramSMP) error {
	err := bimbelRepo.db.Find(&models.ProgramSMP{}, "student_id = ?", data.StudentID).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Model(&models.ProgramSMP{}).Where("student_id = ?", data.StudentID).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) DeleteProgramSMP(id int) error {
	err := bimbelRepo.db.First(&models.ProgramSMP{}, "student_id = ?", id).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Delete(&models.ProgramSMP{}, bimbelRepo.db.Where("student_id = ?", id)).Error
	if err != nil {
		return err
	}
	return nil
}

// //////// P R O G   S M A ////////////////////////////////////////
func (bimbelRepo *BimbelRepository) CreateProgramSMAIpa(data *models.ProgramSMAIpa) error {
	err := bimbelRepo.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) UpdateProgramSMAIpa(data *models.ProgramSMAIpa) error {
	err := bimbelRepo.db.Find(&models.ProgramSMAIpa{}, "student_id = ?", data.StudentID).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Model(&models.ProgramSMAIpa{}).Where("student_id = ?", data.StudentID).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) DeleteProgramSMAIpa(id int) error {
	err := bimbelRepo.db.First(&models.ProgramSMAIpa{}, "student_id = ?", id).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Delete(&models.ProgramSMAIpa{}, bimbelRepo.db.Where("student_id = ?", id)).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) CreateProgramSMAIps(data *models.ProgramSMAIps) error {
	err := bimbelRepo.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) UpdateProgramSMAIps(data *models.ProgramSMAIps) error {
	err := bimbelRepo.db.Find(&models.ProgramSMAIps{}, "student_id = ?", data.StudentID).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Model(&models.ProgramSMAIps{}).Where("student_id = ?", data.StudentID).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (bimbelRepo *BimbelRepository) DeleteProgramSMAIps(id int) error {
	err := bimbelRepo.db.First(&models.ProgramSMAIps{}, "student_id = ?", id).Error
	if err != nil {
		return err
	}

	err = bimbelRepo.db.Delete(&models.ProgramSMAIps{}, bimbelRepo.db.Where("student_id = ?", id)).Error
	if err != nil {
		return err
	}
	return nil
}
