package request

type WorkEntry struct {
	ID        uint   `gorm:"primary_key"`
	StartTime string `gorm:"not null"`
	EndTime   string `gorm:"not null"`
	BreakTime string
}
