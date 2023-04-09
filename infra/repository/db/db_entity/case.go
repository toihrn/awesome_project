package db_entity

import "time"

type CaseEntity struct {
	ID                 int64     `gorm:"column:id;primary_key"` //  pk
	Name               string    `gorm:"column:name"`           // 案件名字
	CriminalId         int64     `gorm:"column:criminal_id"`    //  罪犯id
	CriminalAddress    string    `gorm:"column:criminal_address"`
	CriminalName       string    `gorm:"criminal_name"`                     // 罪犯名
	CriminalPhone      string    `gorm:"column:criminal_phone"`             //联系电话
	CriminalGender     string    `gorm:"column:criminal_gender"`            //性别
	Status             int32     `gorm:"column:status"`                     //案件状态
	CriminalPictureUrl string    `gorm:"column:criminal_picture_url"`       //
	FaceToken          string    `gorm:"column:face_token"`                 // baidu的FaceToken
	CreateTime         time.Time `gorm:"column:create_time;autoCreateTime"` //  系统创建时间（s）
	UpdateTime         time.Time `gorm:"column:update_time;autoUpdateTime"` //  系统更新时间（s）
	ExtraInfo          string    `gorm:"column:extra_info"`                 //扩展字段(json)
}

func (c *CaseEntity) TableName() string {
	return "t_case"
}

/*
ASSISTANT
可以使用 struct tags 标记的 gorm 构建 Go 结构体定义以处理数据库表。

以下是 t_case 表对应的 Go 结构体定义：

type TCase struct {
	ID                 int64     `gorm:"column:id;primaryKey;autoIncrement"`               // pk
	Name               string    `gorm:"column:name"`                                      // 案件名字
	CriminalID         int64     `gorm:"column:criminal_id"`                               // 罪犯id
	CriminalAddress    string    `gorm:"column:criminal_address"`                          // 罪犯地址
	CriminalName       string    `gorm:"column:criminal_name"`                             // 罪犯名
	CriminalPhone      string    `gorm:"column:criminal_phone"`                            // 联系电话
	CriminalGender     string    `gorm:"column:criminal_gender"`                           // 性别
	Status             int32     `gorm:"column:status"`                                    // 案件状态
	CriminalPictureURL string    `gorm:"column:criminal_picture_url"`                      // 罪犯图片 URL
	CreateTime         time.Time `gorm:"column:create_time;autoCreateTime"`                // 创建时间
	UpdateTime         time.Time `gorm:"column:update_time;autoUpdateTime"`                // 更新时间
	ExtraInfo          string    `gorm:"column:extra_info"`                                 // 扩展字段(json)
}

*/
