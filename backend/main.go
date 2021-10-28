package main

import (
	"github.com/nitaxxix/sa-64-final/controller"

	"github.com/nitaxxix/sa-64-final/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// Appoint
	r.GET("/appoints", controller.ListAppoint)
	r.POST("/appoint", controller.CreateAppoint)

	// Insurance
	r.GET("/insrs", controller.ListInsurance)

	r.POST("/insr", controller.CreateInsurance)

	// Job
	r.GET("/jobs", controller.ListJob)

	r.POST("/job", controller.CreateJob)

	// MedicalProduct
	r.GET("/medical_products", controller.ListMedicalProduct)
	r.POST("/medical_product", controller.CreateMedicalProduct)

	// MedRecord
	r.GET("/api/MedRec", controller.ListMedRecord)
	r.POST("/api/submit", controller.CreateMedRecord)

	// Patient
	r.GET("/patients", controller.ListPatient)
	r.POST("/patient", controller.CreatePatient)

	// RemedyType
	r.GET("/remedy_types", controller.ListRemedyType)
	r.POST("/remedy_type", controller.CreateRemedyType)

	// Role
	r.GET("/roles", controller.ListRole)
	r.POST("/role", controller.CreateRole)

	// Screening
	r.GET("/screenings", controller.ListScreening)
	r.POST("/screening", controller.CreateScreening)

	// Sex
	r.GET("/sexs", controller.ListSex)
	r.POST("/sex", controller.CreateSex)

	// Treatment
	r.POST("/treatmentRecord", controller.CreateTreatment)

	// User
	r.GET("/users", controller.ListUser)
	r.GET("/user/dentist/:id", controller.GetUserDentist)
	r.GET("/user/dentistass", controller.GetUserDentistass)
	r.GET("/user/nurse", controller.GetUserNurse)
	r.GET("/user/pharmacist", controller.GetUserPharmacist)
	r.GET("/user/financial", controller.GetUserFinancial)
	r.POST("/user", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
