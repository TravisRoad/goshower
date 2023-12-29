package model

import (
	"time"

	"github.com/TravisRoad/goshower/global"
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	Source      int         `json:"source" gorm:"type:tinyint;comment:来源;uniqueIndex:idx_source"`
	Type        global.Type `json:"type" gorm:"type:tinyint;comment:类型;uniqueIndex:idx_source"`
	MediaID     int64       `json:"media_id" gorm:"type:int;comment:媒体ID"`
	Title       string      `json:"title" gorm:"type:varchar(128);comment:标题"`
	TitleCn     string      `json:"title_cn" gorm:"type:varchar(128);comment:标题中文"`
	Summary     string      `json:"summary" gorm:"type:text;comment:概述"`
	PublishData time.Time   `json:"publish_date" gorm:"comment:发布时间"`
	Status      int         `json:"status" gorm:"type:tinyint;comment:状态"`
	StatusText  string      `json:"status_text" gorm:"type:varchar(32);comment:状态文本"`
}

type AnimeMedia struct {
	Media
	Nsfw         bool    `json:"nsfw" gorm:"type:tinyint(1);comment:是否为限制级"`
	Platform     string  `json:"platform" gorm:"type:varchar(128);comment:发布平台"`
	ImageLarge   string  `json:"image_large" gorm:"type:varchar(128);comment:图片large地址"`
	ImageCommon  string  `json:"image_common" gorm:"type:varchar(128);comment:图片common地址"`
	ImageMedium  string  `json:"image_medium" gorm:"type:varchar(128);comment:图片medium地址"`
	Eps          int     `json:"eps" gorm:"comment:总集数"`
	RatingScore  float32 `json:"rating_score" gorm:"comment:评分总数"`
	RatingDetail string  `json:"rating_detail" gorm:"type:json;comment:评分详情"`
	// ImageSmall   string  `json:"image_small" gorm:"type:varchar(128);comment:图片small地址"`
	// ImageGrid    string  `json:"image_grid" gorm:"type:varchar(128);comment:图片grid地址"`
}

func (AnimeMedia) TableName() string {
	return "anime_media"
}

type Record struct {
	gorm.Model
	UserID  uint `json:"user_id" gorm:"type:bigint;comment:用户ID;index:idx_user_id"`
	Action  int  `json:"action" gorm:"type:tinyint;comment:动作类型"`
	MediaID uint `json:"media_id" gorm:"type:bigint;comment:媒体ID"`
}

type AnimeRecord struct {
	Record
	MediaAction bool `json:"media_action" gorm:"comment:媒体动作，否则为面向单集的动作"`
	TargetEp    int  `json:"target_ep" gorm:"comment:目标集数"`
}

func (AnimeRecord) TableName() string {
	return "anime_record"
}
