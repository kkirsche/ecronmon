package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Hosts is an array of Host objects
type Hosts []Host

// Host represents a network host
type Host struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	IP        string     `json:"ip_address"`
}

// GetHosts is used to retrieve a list of host objects from the database
func GetHosts(db *gorm.DB) Hosts {
	var hosts Hosts
	db.Find(&hosts)
	return hosts
}

// GetHost is used to retrieve a single host object from the database
func GetHost(db *gorm.DB, id int) Host {
	var host Host
	db.Find(&host, id)
	return host
}

// CreateHost is used to create a new host object in the database
func CreateHost(db *gorm.DB, host Host) Host {
	db.Save(&host)
	return host
}

// UpdateHost is used to update an existing host object in the database
func UpdateHost(db *gorm.DB, host Host) {
	db.Save(&host)
}

// DeleteHost is used to delete a host object in the database
func DeleteHost(db *gorm.DB, host Host) {
	db.Delete(&host)
}
