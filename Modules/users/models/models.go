package models

import (
	"time"
)

// Users represents the core user entity
type Users struct {
	ID         string     `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` 
	Email      string     `json:"email" gorm:"uniqueIndex;not null"`                        
	Phone      string     `json:"phone" gorm:"uniqueIndex"`                                  
	Username   string     `json:"username" gorm:"uniqueIndex;not null"`                      
	Password   string     `json:"-" gorm:"not null"`                                        
	AuthType   string     `json:"auth_type" gorm:"not null;default:'email'"`                 
	IsVerified bool       `json:"is_verified" gorm:"default:false"`                          
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`                          
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`                          
	DeletedAt  *time.Time `json:"deleted_at" gorm:"index"`                                  

	// Association with UserDetails
	UserDetails UserDetails `json:"user_details" gorm:"foreignKey:UserID"` // One-to-One relationship with UserDetails
}

// UserDetails represents additional details for a user
type UserDetails struct {
	ID          string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` 
	UserID      string    `json:"user_id" gorm:"uniqueIndex;not null"`                       
	FirstName   string    `json:"first_name" gorm:"not null"`                                
	LastName    string    `json:"last_name" gorm:"not null"`                                 
	ProfilePic  string    `json:"profile_pic"`                                               
	Gender      string    `json:"gender"`                                                    
	DateOfBirth time.Time `json:"date_of_birth"`                                             
	AboutMe     string    `json:"about_me" gorm:"type:text"`                                 
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`                          
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`                          
	
	// Association with Users
	User *Users `json:"user" gorm:"foreignKey:UserID"` // Belongs-to relationship with Users (using a pointer)
}