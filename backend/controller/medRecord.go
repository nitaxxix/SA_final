package controller

import (
	"github.com/nitaxxix/sa-64-final/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)
// get /MedRec
func ListMedRecord(c *gin.Context) {
	var medRecord []entity.MedRecord
	if err := entity.DB().Preload("User").Preload("User.Role").Preload("MedicalProduct").Preload("TreatmentRecord").Preload("TreatmentRecord.ScreeningRecord").Preload("TreatmentRecord.ScreeningRecord.Patient").Raw("SELECT * FROM med_records").Find(&medRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medRecord})
}

// POST /api/submit
func CreateMedRecord(c *gin.Context) {
	var medRecord entity.MedRecord
	var treatment entity.Treatment
	var pharmacist entity.User
	var medicalProduct entity.MedicalProduct

	if err := c.ShouldBindJSON(&medRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา TreatmentRecord ด้วย id
	if err := entity.DB().Where("id = ?", medRecord.TreatmentID).First(&treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TreatmentRecord not found"})
		return
	}

	// ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", medRecord.UserPharmacistID).First(&pharmacist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dentist not found"})
		return
	}
	entity.DB().Joins("Role").Find(&pharmacist)

	if pharmacist.Role.Name != "Dentist" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only for dentsit"})
		return
	}

	// ค้นหา MedicalProduct ด้วย id
	if err := entity.DB().Where("id = ?", medRecord.MedicalProductID).First(&medicalProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalProduct not found"})
		return
	}

	// สร้าง
	mr := entity.MedRecord{
		Treatment:      treatment,
		UserPharmacist: pharmacist,
		MedicalProduct: medicalProduct,
		Amount:         medRecord.Amount,
	}

	// บันทึก
	if err := entity.DB().Create(&mr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mr})

}
