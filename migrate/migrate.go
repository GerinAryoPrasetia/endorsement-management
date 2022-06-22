package migrate

import (
	"project/config"
	m_creator "project/features/creator/data"
)

func AutoMigrate(){
	config.DB.AutoMigrate(
		&m_creator.Creator{}, 
		&m_creator.Category{},
		&m_creator.SocialMedia{},
		&m_creator.FeaturedContent{},
	)
}