package service

import (
	"e4cm/internal/model"
	"log"
)

func Migrate(csvPath, jsonPath, dbPath string) {
	idMap, err := ReadCSV(csvPath)
	if err != nil {
		log.Fatalf("read csv: %v", err)
	}

	twikooComments, err := ReadJSON(jsonPath)
	if err != nil {
		log.Fatalf("read json: %v", err)
	}

	db, err := OpenDB(dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer CloseDB(db)

	var owner model.User
	err = db.First(&owner, "is_owner = ?", true).Error
	if err != nil {
		log.Fatalf("get owner: %v", err)
	}

	comments := make([]model.Comment, 0)
	for _, twikooComment := range twikooComments {
		oldID := twikooComment.GetOldID()
		if _, ok := idMap[oldID]; !ok {
			log.Printf("old id %s not found in id map", oldID)
			continue
		}

		comment := twikooComment.ToComment(idMap[oldID])

		if twikooComment.Master {
			comment.UserID = &owner.ID
			comment.Source = model.SourceSystem
			comment.Nickname = owner.Username
			comment.Email = ""
		}

		comments = append(comments, comment)
	}

	err = db.Create(&comments).Error
	if err != nil {
		log.Fatalf("create comments: %v", err)
	}

	log.Printf("total %d comments", len(twikooComments))
	log.Printf("successfully migrated %d comments", len(comments))

}
