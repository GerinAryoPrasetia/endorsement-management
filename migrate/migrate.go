package migrate

import (
	"project/config"
	m_creator "project/features/creator/data"
)

func AutoMigrate(){
	config.DB.AutoMigrate(
		&m_creator.Creator{},
	)
}