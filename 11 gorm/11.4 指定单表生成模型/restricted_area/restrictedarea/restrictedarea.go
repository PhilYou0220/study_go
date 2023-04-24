
package restrictedarea

type RestrictedArea struct{
    ID uint 
    AreaID int 
    Name string 
    TargetType bool `gorm:"default:0"`
    AreaPoints string 
    CenterPoints string 
    DateType bool `gorm:"default:0"`
    Deleted bool `gorm:"default:0"`
    PassType bool `gorm:"default:0"`
    ConfigSpeed float64 `gorm:"default:0.0"`
    CreatedBy int `gorm:"default:0"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    ConditionType int8 `gorm:"default:0"`
    HasWhiteCar int8 `gorm:"default:0"`
    StopLimit int `gorm:"default:0"`
}
