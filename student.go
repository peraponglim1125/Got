package got
import "time"

type Student struct {	

	ID        int
	FirstName string
	LastName  string
	BirthDate time.Time
	Email     string
	Grade     string
	// Additional fields can be added as needed
}	