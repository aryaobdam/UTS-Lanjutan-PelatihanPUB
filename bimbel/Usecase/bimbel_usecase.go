package usecase

import (
	"UTS-LANJUTAN-ARYA/interfaces"
	"UTS-LANJUTAN-ARYA/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BimbelUsecase struct {
	bimbelRepo interfaces.BimbelRepository
}

func NewBimbelUC(BimbelRepo interfaces.BimbelRepository) *BimbelUsecase {
	return &BimbelUsecase{
		BimbelRepo,
	}

}

func (bimbelUc *BimbelUsecase) GetProgram() []*models.Program {
	result := bimbelUc.bimbelRepo.GetProgram()
	if result == nil {
		return nil
	}

	return result
}

func (bimbelUc *BimbelUsecase) GetStudent() []*models.Tamp {
	result := bimbelUc.bimbelRepo.GetStudent()
	if result == nil {
		return nil
	}

	return result

}

func (bimbelUc *BimbelUsecase) GetTeacher() []*models.Teacher {
	result := bimbelUc.bimbelRepo.GetTeacher()
	if result == nil {
		return nil
	}

	return result

}

func (bimbelUc *BimbelUsecase) GetRequest(c *gin.Context) (*models.Tamp, *models.ProgramSD, *models.ProgramSMP, *models.ProgramSMAIpa, *models.ProgramSMAIps){
	ID, err := strconv.Atoi(c.Param("id_student"))
	if err != nil || ID <= 0 {
		return nil, nil, nil, nil, nil
	} 

	result, resultsd, resultsmp, resultsma, resultsmas := bimbelUc.bimbelRepo.GetRequest(ID)
	if result == nil {
		return nil, nil, nil, nil, nil
	}else if resultsd == nil{
		return result, nil, resultsmp, resultsma, resultsmas
	}else if resultsmp == nil{
		return result, resultsd, nil, resultsma, resultsmas
	}else if resultsma == nil{
		return result, resultsd, resultsmp, nil, resultsmas
	}else if resultsmas == nil{
		return result, resultsd, resultsmp, resultsma, nil
	}

	return result, resultsd, resultsmp, resultsma, resultsmas
}

////////////////////////// C R E A T E /////////////////////////////

func (bimbelUc *BimbelUsecase) CreateProgram(c *gin.Context) error {
	var result models.Program
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.CreateProgram(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) CreateStudent(c *gin.Context) error {
	var result models.Student
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.CreateStudent(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) CreateTeacher(c *gin.Context) error {
	var result models.Teacher
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.CreateTeacher(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) CreateProgramSD(c *gin.Context) error {
	var result models.ProgramSD
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.CreateProgramSD(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) CreateProgramSMP(c *gin.Context) error {
	var result models.ProgramSMP
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.CreateProgramSMP(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) CreateProgramSMAIpa(c *gin.Context) error {
	var result models.ProgramSMAIpa
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.CreateProgramSMAIpa(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) CreateProgramSMAIps(c *gin.Context) error {
	var result models.ProgramSMAIps
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.CreateProgramSMAIps(&result)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////// U P D A T E //////////////////////////////

func (bimbelUc *BimbelUsecase) UpdateStudent(c *gin.Context) error {
	var result models.Student
	ID, err := strconv.Atoi(c.Param("id_student"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.UpdateStudent(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) UpdateTeacher(c *gin.Context) error {
	var result models.Teacher
	ID, err := strconv.Atoi(c.Param("id_teacher"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.UpdateTeacher(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) UpdateProgram(c *gin.Context) error {
	var result models.Program
	ID, err := strconv.Atoi(c.Param("id_program"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.UpdateProgram(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) UpdateProgramSD(c *gin.Context) error {
	var result models.ProgramSD
	ID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.UpdateProgramSD(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) UpdateProgramSMP(c *gin.Context) error {
	var result models.ProgramSMP
	ID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.UpdateProgramSMP(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) UpdateProgramSMAIpa(c *gin.Context) error {
	var result models.ProgramSMAIpa
	ID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.UpdateProgramSMAIpa(&result)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) UpdateProgramSMAIps(c *gin.Context) error {
	var result models.ProgramSMAIps
	ID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = bimbelUc.bimbelRepo.UpdateProgramSMAIps(&result)
	if err != nil {
		return err
	}

	return nil
}

////////////////////// D E L E T E ////////////////////////////////

func (bimbelUc *BimbelUsecase) DeleteStudent(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id_student"))
	if err != nil || ID <= 0 {
		return err
	}

	err = bimbelUc.bimbelRepo.DeleteStudent(ID)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) DeleteTeacher(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id_teacher"))
	if err != nil || ID <= 0 {
		return err
	}

	err = bimbelUc.bimbelRepo.DeleteTeacher(ID)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) DeleteProgram(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id_program"))
	if err != nil || ID <= 0 {
		return err
	}

	err = bimbelUc.bimbelRepo.DeleteProgram(ID)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) DeleteProgramSD(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = bimbelUc.bimbelRepo.DeleteProgramSD(ID)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) DeleteProgramSMP(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = bimbelUc.bimbelRepo.DeleteProgramSMP(ID)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) DeleteProgramSMAIpa(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = bimbelUc.bimbelRepo.DeleteProgramSMAIpa(ID)
	if err != nil {
		return err
	}

	return nil
}

func (bimbelUc *BimbelUsecase) DeleteProgramSMAIps(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = bimbelUc.bimbelRepo.DeleteProgramSMAIps(ID)
	if err != nil {
		return err
	}

	return nil
}
