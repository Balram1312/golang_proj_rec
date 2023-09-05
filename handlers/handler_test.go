package handlers
import(
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"fmt"
	"bytes"
)

type MockApp struct {
	mock.Mock
	*MyApp
}

// ExistingUser is a mock implementation of the ExistingUser method
// func (m *MockApp) ExistingUser(username string) bool {
// 	args := m.Called(username)
// 	return args.Bool(0)
// }

func TestPostEmployees(t *testing.T){
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	   })
	db, _ := gorm.Open(dialector, &gorm.Config{})
	
	app := &MyApp{DB:db}
	
	//test cases
	t.Run("User exists", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"ID","Username", "Password"}).AddRow(1, "balram","balram1312").AddRow(2,"suresh","balram1312")
		mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
		result := app.ExistingUser("balram")
		assert.Equal(t, result,true)
	})

	t.Run("User does not exist", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"ID","Username", "Password"}).AddRow(1, "balram","balram1312").AddRow(2,"suresh","balram1312")
		mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

		result := app.ExistingUser("john")
		assert.Equal(t, result,false)
	})

	t.Run("User exists with different password", func(t *testing.T) {
		// Define the expected rows for the SQL query
		expectedRows := sqlmock.NewRows([]string{"ID", "Username", "Password"}).
			AddRow(1, "balram", "balram1312").
			AddRow(2, "suresh", "differentpassword")

		// Expect the SQL query and return the expected rows
		mock.ExpectQuery(`SELECT`).WillReturnRows(expectedRows)

		result := app.ExistingUser("suresh")
		assert.False(t, result)
	})

	t.Run("No rows returned from the database", func(t *testing.T) {
		// Expect the SQL query to be executed and return no rows
		mock.ExpectQuery(`SELECT`).WillReturnRows(sqlmock.NewRows([]string{}))

		result := app.ExistingUser("newuser")
		assert.False(t, result)
	})

	t.Run("Database query fails", func(t *testing.T) {
		// Expect the SQL query to fail
		mock.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("database error"))
		
		result := app.ExistingUser("suresh")
		assert.False(t, result)
	})
}

func TestPostEmployeesHandler(t *testing.T) {
	// Initialize a new Gin router
	router := gin.Default()

	// Create a new SQL mock
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()

	// Initialize the GORM database
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	// Create an instance of MockApp
	app := &MockApp{
		MyApp: &MyApp{
			DB: db, // Use a mock database or set it as needed
		},
	}

	// Set up expectations for the ExistingUser method
	// app.On("ExistingUser", "suresh").Return(true)
	// app.On("ExistingUser", "newuser").Return(true)
	// // Set up the router with the post_employees handler
	router.POST("/post_employees", app.post_employees)

	t.Run("User already exists", func(t *testing.T) {
		// Create a request to the endpoint
		reqBody := []byte(`{"Username": "suresh", "Password": "password123"}`)
		req, _ := http.NewRequest("POST", "/post_employees", bytes.NewBuffer(reqBody))
	
		// Create a response recorder to capture the response
		w := httptest.NewRecorder()
	
		// Serve the request to the router
		router.ServeHTTP(w, req)
	
		// Check the response status code and body
		assert.Equal(t, 400, w.Code)
	
		// Add more assertions if needed
	})
	
	t.Run("User does not exist", func(t *testing.T) {
		// Create a request to the endpoint
		reqBody := []byte(`{"Username": "newuser", "Password": "password123"}`)
		req, _ := http.NewRequest("POST", "/post_employees", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		// Create a response recorder to capture the response
		w := httptest.NewRecorder()
	
		// Serve the request to the router
		router.ServeHTTP(w, req)
	
		// Check the response status code and body
		assert.Equal(t, 201, w.Code)
	
		// Add more assertions if needed
	})
	
}
