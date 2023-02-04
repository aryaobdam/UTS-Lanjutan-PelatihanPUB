package handler

import (
	"UTS-LANJUTAN-ARYA/interfaces"
	"UTS-LANJUTAN-ARYA/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bimbelHandler struct {
	bimbelUc interfaces.BimbelUsecase
}

func BimbelRoute(bimbelUc interfaces.BimbelUsecase, r *gin.RouterGroup) {
	uc := bimbelHandler{
		bimbelUc,
	}

	v2 := r.Group("bimbel")
	v2.GET("", uc.GetStudent)
	v2.GET(":id_student", uc.GetRequest)
	v2.POST("", uc.PostStudent)
	v2.PUT(":id_student", uc.UpdateStudent)
	v2.DELETE(":id_student", uc.DeleteStudent)

	v2.GET("/prog", uc.GetProgram)
	v2.POST("/prog", uc.PostProgram)
	v2.PUT("/prog/:id_program", uc.UpdateProgram)
	v2.DELETE("/prog/:id_program", uc.DeleteProgram)

	v2.GET("/teach", uc.GetTeacher)
	v2.POST("/teach", uc.PostTeacher)
	v2.PUT("/teach/:id_teacher", uc.UpdateTeacher)
	v2.DELETE("/teach/:id_teacher", uc.DeleteTeacher)

	v2.POST("/sd", uc.PostProgramSD)
	v2.PUT("/sd/:student_id", uc.UpdateProgramSD)
	v2.DELETE("/sd/:id_sd", uc.DeleteProgramSD)

	v2.POST("/smp", uc.PostProgramSMP)
	v2.PUT("/smp/:student_id", uc.UpdateProgramSMP)
	v2.DELETE("/smp/:student_id", uc.UpdateProgramSMP)

	v2.POST("/sma", uc.PostProgramSMAIpa)
	v2.PUT("/sma/:student_id", uc.UpdateProgramSMAIpa)
	v2.DELETE("/sma/:student_id", uc.DeleteProgramSMAIpa)

	v2.POST("/smas", uc.PostProgramSMAIps)
	v2.PUT("/smas/:student_id", uc.UpdateProgramSMAIps)
	v2.DELETE("smas/:student_id", uc.DeleteProgramSMAIps)

}

func (bimbelHandler *bimbelHandler) GetStudent(c *gin.Context) {

	result := bimbelHandler.bimbelUc.GetStudent()
	if result == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:   result,
		Status: "succes",
	})
}

func (bimbelHandler *bimbelHandler) GetRequest(c *gin.Context) {
	result, resultsd, resultsmp, resultsma, resultsmas := bimbelHandler.bimbelUc.GetRequest(c)
	if result == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "data not found",
		})
		c.JSON(http.StatusBadRequest, response.Response{
			Status:  "errors",
			Message: "data not found",
		})
		return
	} else if resultsd != nil {
		c.JSON(http.StatusOK, response.Response{
			Status: "succes",
			Data:   result,
			Sd:     resultsd,
		})
		return
	} else if resultsmp != nil {
		c.JSON(http.StatusOK, response.Response{
			Status: "succes",
			Data:   result,
			Smp:    resultsmp,
		})
		return
	} else if resultsma != nil {
		c.JSON(http.StatusOK, response.Response{
			Status: "succes",
			Data:   result,
			Sma:    resultsma,
		})
		return
	} else if resultsmas != nil {
		c.JSON(http.StatusOK, response.Response{
			Status: "succes",
			Data:   result,
			Smas:   resultsmas,
		})
		return
	}
}

func (bimbelHandler *bimbelHandler) PostStudent(c *gin.Context) {
	err := bimbelHandler.bimbelUc.CreateStudent(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "Succes",
		Message: "Succes",
	})
}

func (bimbelHandler *bimbelHandler) UpdateStudent(c *gin.Context) {
	err := bimbelHandler.bimbelUc.UpdateStudent(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) DeleteStudent(c *gin.Context) {
	err := bimbelHandler.bimbelUc.DeleteStudent(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})

}

///////////////////////// P R O G R A M ////////////////////////////

func (bimbelHandler *bimbelHandler) PostProgram(c *gin.Context) {
	err := bimbelHandler.bimbelUc.CreateProgram(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) GetProgram(c *gin.Context) {
	result := bimbelHandler.bimbelUc.GetProgram()
	if result == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) UpdateProgram(c *gin.Context) {
	err := bimbelHandler.bimbelUc.UpdateProgram(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) DeleteProgram(c *gin.Context) {
	err := bimbelHandler.bimbelUc.DeleteProgram(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})

}

/////////////////////////// T E A C H E R /////////////////////////

func (bimbelHandler *bimbelHandler) PostTeacher(c *gin.Context) {
	err := bimbelHandler.bimbelUc.CreateTeacher(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) GetTeacher(c *gin.Context) {
	result := bimbelHandler.bimbelUc.GetTeacher()
	if result == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) UpdateTeacher(c *gin.Context) {
	err := bimbelHandler.bimbelUc.UpdateTeacher(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) DeleteTeacher(c *gin.Context) {
	err := bimbelHandler.bimbelUc.DeleteTeacher(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})

}

// /////////// P R O G   SD ///////////////////////////////////////
func (bimbelHandler *bimbelHandler) PostProgramSD(c *gin.Context) {
	err := bimbelHandler.bimbelUc.CreateProgramSD(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) UpdateProgramSD(c *gin.Context) {
	err := bimbelHandler.bimbelUc.UpdateProgramSD(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) DeleteProgramSD(c *gin.Context) {
	err := bimbelHandler.bimbelUc.DeleteProgramSD(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})

}

// /////////// P R O G   S M P ///////////////////////////////////////
func (bimbelHandler *bimbelHandler) PostProgramSMP(c *gin.Context) {
	err := bimbelHandler.bimbelUc.CreateProgramSMP(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) UpdateProgramSMP(c *gin.Context) {
	err := bimbelHandler.bimbelUc.UpdateProgramSMP(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) DeleteProgramSMP(c *gin.Context) {
	err := bimbelHandler.bimbelUc.DeleteProgramSMP(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})

}

// /////////// P R O G   S M A  I P A ///////////////////////////////////////
func (bimbelHandler *bimbelHandler) PostProgramSMAIpa(c *gin.Context) {
	err := bimbelHandler.bimbelUc.CreateProgramSMAIpa(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) UpdateProgramSMAIpa(c *gin.Context) {
	err := bimbelHandler.bimbelUc.UpdateProgramSMAIpa(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) DeleteProgramSMAIpa(c *gin.Context) {
	err := bimbelHandler.bimbelUc.DeleteProgramSMAIpa(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})

}

// /////////// P R O G   S M A  I P S ///////////////////////////////////////
func (bimbelHandler *bimbelHandler) PostProgramSMAIps(c *gin.Context) {
	err := bimbelHandler.bimbelUc.CreateProgramSMAIps(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) UpdateProgramSMAIps(c *gin.Context) {
	err := bimbelHandler.bimbelUc.UpdateProgramSMAIps(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (bimbelHandler *bimbelHandler) DeleteProgramSMAIps(c *gin.Context) {
	err := bimbelHandler.bimbelUc.DeleteTeacher(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "succes",
		Message: "succes",
	})

}
