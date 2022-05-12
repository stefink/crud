package main

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func (p *Person) Insert() uint {
	db := Init()
	defer db.Commit()
	person := Person{Name: p.Name, Age: p.Age}
	db.Create(&person)
	fmt.Printf(p.Name + " Inserted ")
	fmt.Println("with ID", person.ID)
	return person.ID
}

func Delete(id string) {
	db := Init()
	defer db.Commit()
	fmt.Println(db.Delete(&Person{}, id).RowsAffected, " Rows Deleted")
}

func Update(id int, new_person *Person) {
	db := Init()
	defer db.Commit()
	var person Person
	if err := db.First(&person, id).Error; err != nil {
		fmt.Println("Element not found!")
		os.Exit(1)
	}
	db.Model(&person).Updates(new_person)
	fmt.Println("Value was updated!")
}

func GetElementByID(id int) (person Person) {
	db := Init()
	defer db.Commit()
	if err := db.First(&person, id).Error; err != nil {
		fmt.Println("Element not found!")
		os.Exit(1)
	}
	return person
}

func ListAllElements() {
	db := Init()
	defer db.Commit()
	var persons []Person
	db.Find(&persons)
	if len(persons) == 0 {
		fmt.Println("No records found!")
		os.Exit(1)
	}
	for _, person := range persons {
		fmt.Println("ID: ", person.ID, "\tName: ", person.Name, "\tAge: ", person.Age)
	}
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Person{})
	return db
}
