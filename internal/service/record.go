package service

import (
	"errors"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
	"gorm.io/gorm"
)

type RecordService struct{}

func (s *RecordService) RecordSubject(id int, src global.Source, uid uint, action global.Status) error {
	record := model.SubjectRecord{}
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		var cnt int64
		if err := tx.Model(&model.Media{}).Where("media_id = ? AND source = ?", id, src).Count(&cnt).Error; err != nil {
			return err
		}
		// there is no such media
		if cnt == 0 {
			mediaer, err := getMediaer(src)
			if err != nil {
				return err
			}
			detail, err := mediaer.MediaDetail(id)
			if err != nil {
				return err
			}
			if err := tx.Model(&model.Media{}).Create(detail).Error; err != nil {
				return err
			}
		}
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

func (s *RecordService) RecordEp(id int, src global.Source, uid uint, action global.Status, ep int) error {
	return nil
}
