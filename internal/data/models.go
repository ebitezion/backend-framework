package data

import (
	"database/sql"
	"errors"
)

var (
	// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
	// looking up a movie that doesn't exist in our database.
	ErrRecordNotFound = errors.New("record not found")

	// Define a custom ErrEditConflict error. We'll return this from our Update() method
	// when there is a data race.
	ErrEditConflict = errors.New("edit conflict")
)

// Create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	//AModel MyModel
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the intitialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		//	AModel: MyModel{DB: db},
	}
}
