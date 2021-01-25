package seed

import (
	"awesomeProject/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

var users = []models.User{
	models.User{
		Nickname: "Grinaldi Wisnu",
		Email:    "grinaldifoc@gmail.com",
		Password: "grinaldi",
	},
	models.User{
		Nickname: "Nasha Salsabila Destya Ananta",
		Email:    "ashblda21@gmail.com",
		Password: "nasha",
	},
}

var posts = []models.Post{
	{
		Title:    "Title 1",
		Content:  "Hello world 1",
		AuthorID: uint32(1),
	},
	{
		Title:    "Title 2",
		Content:  "Hello world 2",
		AuthorID: uint32(1),
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot drop table : %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
