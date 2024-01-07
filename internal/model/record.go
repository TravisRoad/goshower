package model

import (
	"time"

	"github.com/TravisRoad/goshower/global"
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	Source      global.Source `json:"source" gorm:"type:tinyint;comment:来源;uniqueIndex:idx_source"`
	Type        global.Type   `json:"type" gorm:"type:tinyint;comment:类型"`
	Link        string        `json:"link" gorm:"type:varchar(128);comment:链接"`
	MediaID     int           `json:"media_id" gorm:"type:int;comment:媒体ID;uniqueIndex:idx_source"`
	Title       string        `json:"title" gorm:"type:varchar(128);comment:标题"`
	TitleCn     string        `json:"title_cn" gorm:"type:varchar(128);comment:标题中文"`
	Summary     string        `json:"summary" gorm:"type:text;comment:概述"`
	PublishData time.Time     `json:"publish_date" gorm:"comment:发布时间"`
	Nsfw        bool          `json:"nsfw" gorm:"type:tinyint(1);comment:是否为限制级"`
	Platform    string        `json:"platform" gorm:"type:varchar(128);comment:发布平台"`
	ImageLarge  string        `json:"image_large" gorm:"type:varchar(128);comment:图片large地址"`
	ImageCommon string        `json:"image_common" gorm:"type:varchar(128);comment:图片common地址"`
	ImageMedium string        `json:"image_medium" gorm:"type:varchar(128);comment:图片medium地址"`
	Eps         int           `json:"eps" gorm:"comment:总集数"`
	RatingScore float32       `json:"rating_score" gorm:"comment:评分总数"`
	// RatingDetail string        `json:"rating_detail" gorm:"type:json;comment:评分详情"`
	// ImageSmall   string  `json:"image_small" gorm:"type:varchar(128);comment:图片small地址"`
	// ImageGrid    string  `json:"image_grid" gorm:"type:varchar(128);comment:图片grid地址"`
}

func (Media) TableName() string {
	return "media"
}

type SubjectRecord struct {
	gorm.Model
	UserID  uint `json:"user_id" gorm:"comment:用户ID;index:idx_user_media"`
	MediaID int  `json:"media_id" gorm:"type:bigint;comment:媒体ID;index:idx_user_media"`
	Action  int  `json:"action" gorm:"type:tinyint;comment:动作类型"`
}

func (SubjectRecord) TableName() string {
	return "subject_record"
}

type EpRecord struct {
	gorm.Model
	UserID  uint `json:"user_id" gorm:"comment:用户ID;index:idx_user_id"`
	Action  int  `json:"action" gorm:"type:tinyint;comment:动作类型"`
	Ep      int  `json:"ep" gorm:"comment:目标集数"`
	MediaID int  `json:"media_id" gorm:"type:int;comment:媒体ID"`
}

func (EpRecord) TableName() string {
	return "ep_record"
}

type EpRecordDetail struct {
	gorm.Model
	UserID  uint   `json:"user_id" gorm:"comment:用户ID;index:idx_user_id"`
	Detail  []byte `json:"detail" gorm:"type:blob;comment:详情"`
	MediaID int    `json:"media_id" gorm:"type:int;comment:媒体ID"`
}

func (EpRecordDetail) TableName() string {
	return "ep_record_detail"
}
