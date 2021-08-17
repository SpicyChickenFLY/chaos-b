package dao

import (
	"time"

	"github.com/SpicyChickenFLY/chaos-b/model"
	"github.com/lingdor/stackerror"
	"gorm.io/gorm"
)

// ==================== Test ====================

// GetAllTests get all Tests
func GetAllTests(tx *gorm.DB, tests *model.Tests) error {
	result := tx.Where(&model.Test{Deleted: false}).Find(tests)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTestByID get Tests by ID
func GetTestByID(tx *gorm.DB, tests *model.Tests, testID int) error {
	result := tx.Where(&model.Test{ID: testID, Deleted: false}).First(&tests)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// CreateTest add Test
func CreateTest(tx *gorm.DB, test *model.Test) error {
	test.CreatedAt = time.Now()
	test.UpdatedAt = time.Now()
	if err := tx.Create(&test).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DeleteTest delete Test
func DeleteTest(tx *gorm.DB, testID int) error {
	if err := tx.Model(&model.Test{}).Where(
		"id=?", testID).Update("deleted", true).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// UpdateTest update Test
func UpdateTest(tx *gorm.DB, test *model.Test) error {
	test.UpdatedAt = time.Now()
	if err := tx.Model(test).Where(
		"id=?", test.ID).Update(
		"content", test.Content).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}
