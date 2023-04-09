package model

import (
	"github.com/bytedance/sonic"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
)

type FaceTag struct {
	Age              int `json:"age"`
	Gender           int `json:"gender"`
	BodyType         int `json:"body_type"`
	FaceShape        int `json:"face_shape"`
	HairStyle        int `json:"hair_style"`
	HairColor        int `json:"hair_color"`
	FacialHair       int `json:"facial_hair"`
	EyebrowStyle     int `json:"eyebrow_style"`
	EyebrowDirection int `json:"eyebrow_direction"`
	EyeShape         int `json:"eye_shape"`
	EyeDistance      int `json:"eye_distance"`
	DoubleEyelid     int `json:"double_eyelid"`
	EyeBag           int `json:"eye_bag"`
	NoseShape        int `json:"nose_shape"`
	NoseEyeDistance  int `json:"nose_eye_distance"`
	NasolabialFold   int `json:"nasolabial_fold"` // 法令纹
	Wrinkles         int `json:"wrinkles"`
	MouthSize        int `json:"mouth_size"`
	LipThickness     int `json:"lip_thickness"`
	FacialExpression int `json:"facial_expression"`
}

func (t *FaceTag) ToMap() map[string]int {
	str := xjson.ToString(t)
	resMap := map[string]int{}
	err := sonic.UnmarshalString(str, resMap)
	if err != nil {
		return nil
	}
	return resMap
}
