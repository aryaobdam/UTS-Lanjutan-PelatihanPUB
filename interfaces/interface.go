package interfaces

import (
	"UTS-LANJUTAN-ARYA/models"

	"github.com/gin-gonic/gin"
)

type BimbelRepository interface {
	GetStudent() []*models.Tamp
	GetTeacher() []*models.Teacher
	GetProgram() []*models.Program
	GetRequest(id int) (*models.Tamp, *models.ProgramSD, *models.ProgramSMP, *models.ProgramSMAIpa, *models.ProgramSMAIps)
	CreateStudent(*models.Student) error
	CreateTeacher(*models.Teacher) error
	CreateProgram(*models.Program) error
	CreateProgramSD(*models.ProgramSD) error
	CreateProgramSMP(*models.ProgramSMP) error
	CreateProgramSMAIpa(*models.ProgramSMAIpa) error
	CreateProgramSMAIps(*models.ProgramSMAIps) error
	UpdateStudent(*models.Student) error
	UpdateTeacher(*models.Teacher) error
	UpdateProgram(*models.Program) error
	UpdateProgramSD(*models.ProgramSD) error
	UpdateProgramSMP(*models.ProgramSMP)error
	UpdateProgramSMAIpa(*models.ProgramSMAIpa)error
	UpdateProgramSMAIps(*models.ProgramSMAIps)error
	DeleteStudent(id int) error
	DeleteTeacher(id int) error
	DeleteProgram(id int) error
	DeleteProgramSD(id int) error
	DeleteProgramSMP(id int) error
	DeleteProgramSMAIpa(id int)error
	DeleteProgramSMAIps(id int)error
}

type BimbelUsecase interface {
	GetStudent() []*models.Tamp
	GetTeacher() []*models.Teacher
	GetProgram() []*models.Program
	GetRequest(*gin.Context) (*models.Tamp, *models.ProgramSD, *models.ProgramSMP, *models.ProgramSMAIpa, *models.ProgramSMAIps)
	CreateStudent(*gin.Context) error
	CreateTeacher(*gin.Context) error
	CreateProgram(*gin.Context) error
	CreateProgramSD(*gin.Context) error
	CreateProgramSMP(*gin.Context) error
	CreateProgramSMAIpa(*gin.Context) error
	CreateProgramSMAIps(*gin.Context) error
	UpdateStudent(*gin.Context) error
	UpdateTeacher(*gin.Context) error
	UpdateProgram(*gin.Context) error
	UpdateProgramSD(*gin.Context) error
	UpdateProgramSMP(*gin.Context) error
	UpdateProgramSMAIpa(*gin.Context) error
	UpdateProgramSMAIps(*gin.Context) error
	DeleteStudent(*gin.Context) error
	DeleteTeacher(*gin.Context) error
	DeleteProgram(*gin.Context) error
	DeleteProgramSD(*gin.Context) error
	DeleteProgramSMP(*gin.Context) error
	DeleteProgramSMAIpa(*gin.Context) error
	DeleteProgramSMAIps(*gin.Context) error

}
