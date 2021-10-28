package controller

import (
	"github.com/nitaxxix/sa-64-final/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateUser(c *gin.Context) {

	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /user/dentist

func GetUserDentist(c *gin.Context) {

	var user entity.User

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ? ", id).Scan(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /user/dentistass
func GetUserDentistass(c *gin.Context) {

	var user entity.User

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ? ", id).Scan(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /user/nurse
func GetUserNurse(c *gin.Context) {

	var user entity.User

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ? ", id).Scan(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /user/pharmacist
func GetUserPharmacist(c *gin.Context) {

	var user entity.User

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ? ", id).Scan(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /user/financial
func GetUserFinancial(c *gin.Context) {

	var user entity.User

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /users

func ListUser(c *gin.Context) {

	var users []entity.User

	if err := entity.DB().Raw("SELECT * FROM users").Scan(&users).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": users})

}
