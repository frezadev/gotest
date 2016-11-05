package models

type BrandAnalysis struct {
	ID        int64  `db:"id"`
	BrandName string `db:"brand_name"`
	TotalUser string `db:"total_user"`
}
