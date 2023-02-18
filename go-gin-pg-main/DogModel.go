package main

type Dog struct {
	tableName struct{} `pg:"dogs"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsSpot    bool     `json:"is_spot" pg:"is_spot"`
	Color     string   `json:"color"  pg:"color"`
}

// FindAllDogs Получить список песиков.
func FindAllDogs() []Dog {
	var dogs []Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dogs).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dogs
}

// CreateDog Создать песика.
func CreateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// FindDogById Получить песика по id.
func FindDogById(id string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dog).
		Where("id = ?", id).
		First()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// DeleteDogById Удалить песика по id.
func DeleteDogById(id string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).
		Where("id = ?", id).
		Delete()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

func UpdateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	oldDog := FindDogById(dog.ID)

	oldDog.Name = dog.Name
	oldDog.IsSpot = dog.IsSpot
	oldDog.Color = dog.Color

	_, err := pgConnect.Model(&oldDog).
		Where("name = ?", oldDog.Name).
		Where("id = ?", oldDog.ID).
		Where("color = ?", oldDog.Color).
		Where("is_spot = ?", oldDog.IsSpot).
		Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return oldDog
}
