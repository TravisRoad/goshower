package service

import (
	"errors"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
	"gorm.io/gorm"
)

type RecordService struct{}

func (s *RecordService) RecordSubject(id int, uid uint, action global.Status) error {
	record := model.SubjectRecord{}
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.SubjectRecord{}).Where("user_id = ? AND media_id = ?", uid, id).First(&record).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			record.UserID = uid
			record.MediaID = id
		}
		record.Action = int(action)
		if err := tx.Save(&record).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *RecordService) RecordEp(id int, uid uint, action global.Status, ep int) error {
	return nil
}
