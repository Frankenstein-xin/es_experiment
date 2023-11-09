package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type RawBagStruct struct {
	Id                   int64          `json:"Id"`
	TenantId             int64          `json:"tenant_id"`
	ProjectId            string         `json:"project_id"`
	TaskId               int64          `json:"task_id"`
	VehicleId            string         `json:"vehicle_id"`
	Channel              string         `json:"channel"`
	TosRegion            string         `json:"tos_region"`
	TosBucket            string         `json:"tos_bucket"`
	ObjectName           string         `json:"object_name"`
	FileId               string         `json:"file_id"`
	FileName             string         `json:"file_name"`
	RealFileName         string         `json:"real_file_name"`
	Size                 int64          `json:"size"`
	UploadedSize         int64          `json:"uploaded_size"`
	Status               string         `json:"status"`
	ComplianceStatus     string         `json:"compliance_status"`
	ProtocolId           string         `json:"protocol_id"`
	Md5                  string         `json:"md5"`
	StartedAt            string         `json:"started_at" gorm:"column:started_at"`
	EndedAt              string         `json:"ended_at" gorm:"column:ended_at"`
	Mtime                time.Time      `json:"mtime"`
	ReadyTime            time.Time      `json:"ready_time"`
	CreatedAt            time.Time      `gorm:"column:created_at;<-:false" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"<-:update;autoUpdateTime" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"<-;column:deleted_at" json:"-"` // 需要给<-权限
	FileDeletedStatus    int64          `json:"file_deleted_status"`
	FileType             string         `json:"file_type" gorm:"column:file_type"`
	Compliance           int64          `json:"compliance" gorm:"compliance"`
	Version              int64          `json:"version" gorm:"column:version"`
	UserDefinedAttribute datatypes.JSON `json:"user_defined_attribute" gorm:"user_defined_attribute"`
	ComplianceTime       *time.Time     `json:"compliance_time" gorm:"compliance_time"`
	Encrypted            int64          `json:"encrypted" gorm:"encrypted"`
	BatchId              string         `json:"batch_id" `
	VehicleCalibrationId string         `json:"vehicle_calibration_id"`
	UploadStartTime      *time.Time     `json:"upload_start_time"`
	UploadEndTime        *time.Time     `json:"upload_end_time"`
	ReplayStatus         int64          `json:"replay_status"`
}

func (RawBagStruct) TableName() string {
	return "raw_bag"
}
