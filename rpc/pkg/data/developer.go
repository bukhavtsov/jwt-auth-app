package data

import (
	"github.com/jinzhu/gorm"
)

type DeveloperData interface {
	Create(developer *Developer) (int64, error)
	Read(id int64) (*Developer, error)
	ReadAll() ([]*Developer, error)
	Update(developer *Developer) (*Developer, error)
	Delete(id int64) error
}

type Developer struct {
	Id           int64  `gorm:"column:id" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	Age          int64  `gorm:"column:age" json:"age"`
	PrimarySkill string `gorm:"column:primary_skill" json:"primary_skill"`
}

type developerData struct {
	db *gorm.DB
}

func NewDeveloperData(db *gorm.DB) *developerData {
	return &developerData{db}
}

func (d *developerData) Read(id int64) (*Developer, error) {
	developer := Developer{}
	if err := d.db.Where("id = ?", id).Find(&developer).Error; err != nil {
		return nil, err
	}
	return &developer, nil
}

func (d *developerData) ReadAll() ([]*Developer, error) {
	developers := make([]*Developer, 0)
	if err := d.db.Find(&developers).Error; err != nil {
		return []*Developer{}, err
	}
	return developers, nil
}

func (d *developerData) Create(developer *Developer) (int64, error) {
	if err := d.db.Create(developer).Error; err != nil {
		return -1, err
	}
	return developer.Id, nil
}

func (d *developerData) Update(developer *Developer) (*Developer, error) {
	if err := d.db.Save(&developer).Error; err != nil {
		return nil, err
	}
	return developer, nil
}

func (d *developerData) Delete(id int64) error {
	if err := d.db.Where("id = ?", id).Delete(&Developer{}).Error; err != nil {
		return err
	}
	return nil
}
