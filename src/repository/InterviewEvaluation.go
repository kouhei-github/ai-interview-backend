package repository

import "gorm.io/gorm"

type InterviewEvaluation struct {
	gorm.Model
	InterviewID  uint       `gorm:"type:int;not null" json:"interview_id"`
	ScoreByAI    float32    `gorm:"not null" json:"score_by_ai"`
	ScoreByHuman float32    `gorm:"not null" json:"score_by_human"`
	Comment      string     `gorm:"type:longtext" json:"comment"`
	Interview    *Interview `gorm:"foreignkey:InterviewID"`
}

func (receiver *InterviewEvaluation) Save() error {
	return db.Create(receiver).Error
}

func (receiver *InterviewEvaluation) FindById(id uint) (*InterviewEvaluation, error) {
	/**
	このクエリは、複数のInterviewEvaluationレコードを取得し、それぞれに関連付けられたInterviewとApplicantの情報をロードします。
	Preload関数の中に無名関数を指定することで、特定のカラムのみを取得しています。
	この方法で、GORMは多対一のリレーションシップを通じて関連するエンティティのデータを取得し、
	その結果をGoの構造体にマップします。結果を複数の変数に分けて格納する代わりに、一連の入れ子になったオブジェクトの形で全ての情報を一つの変数に格納します。これにより、ソースコードの簡潔さと処理の効率性の両方を実現できます。
	*/
	var interviewEvaluation InterviewEvaluation
	database := db.Preload("Interview", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "InterviewType", "InterviewStatus", "ApplicantID") // ApplicantIDがないとApplicantモデルからデータを取得できない
	}).Preload("Interview.Applicant", func(db *gorm.DB) *gorm.DB {
		return db.Select("name", "Email", "Phone", "Resume", "AppliedJob", "ID")
	}).Where("id = ?", id).First(&interviewEvaluation)
	return &interviewEvaluation, database.Error
}
