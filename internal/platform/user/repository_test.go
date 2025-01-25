package user

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByEmail_Success(t *testing.T) {
	// Crear el mock de la base de datos
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("un error '%s' no esperado al abrir la base de datos", err)
	}
	defer db.Close()

	// Crear el repositorio con el mock de la base de datos
	repo := NewRepository(db)

	// Configurar el comportamiento esperado del mock
	rows := mockDB.NewRows([]string{"id", "name", "age", "email", "password"}).
		AddRow("123", "Esteban", 25, "estebanpoly@gmail.com", "password123")

	mockDB.ExpectPrepare(QueryByEmail).
		ExpectQuery().
		WithArgs("estebanpoly@gmail.com").
		WillReturnRows(rows)

	// Llamar al método bajo prueba
	user, err := repo.GetUserByEmail("estebanpoly@gmail.com")

	// Verificar los resultados
	assert.NoError(t, err)
	assert.Equal(t, "estebanpoly@gmail.com", user.Email)
	assert.Equal(t, "123", user.ID)
}

func TestGetUserByEmail_Error(t *testing.T) {
	// Crear el mock de la base de datos
	db,mockDB , err := sqlmock.New()
	if err != nil {
		t.Fatalf("un error '%s' no esperado al abrir la base de datos", err)
	}
	defer db.Close()

	// Crear el repositorio con el mock de la base de datos
	repo := NewRepository(db)

	// Configurar el comportamiento esperado del mock
	mockDB.ExpectPrepare(QueryByEmail).
		ExpectQuery().
		WithArgs("estebanpoly@gmail.com").
		WillReturnError(sql.ErrNoRows)

	// Llamar al método bajo prueba
	_, err = repo.GetUserByEmail("estebanpoly@gmail.com")

	// Verificar los resultados
	assert.Error(t, err)
}

func TestGetUseByID_Succes(t *testing.T) {
	db, mockDB, err := sqlmock.New()	
	if err != nil {
		t.Fatalf("un error '%s' no esperado al abrir la base de datos", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	// Configurar el comportamiento esperado del mock
	rows := mockDB.NewRows([]string{"id", "name", "age", "email"}).
		AddRow("123", "Esteban", 25, "estebanpoly@gmail.com")

	mockDB.ExpectPrepare(queryGetByID).
		ExpectQuery().
		WithArgs("123").
		WillReturnRows(rows)

	user, err := repo.GetByID("123")

	assert.NoError(t, err)
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "Esteban", user.Name)
	assert.Equal(t, int8(25), user.Age)
	assert.Equal(t, "estebanpoly@gmail.com", user.Email)

}

func TestGetUserByID_Error(t *testing.T) {
	db,mockDB,err := sqlmock.New()
	if err != nil {
		t.Fatalf("un error '%s' no esperado al abrir la base de datos", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mockDB.ExpectPrepare(queryGetByID).
		ExpectQuery().
		WithArgs("123").
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetByID("123")

	assert.Error(t, err)
}

func TestSave_Success(t *testing.T) {
	db , mockDB , err := sqlmock.New()
	if err != nil {
		t.Fatalf("un error '%s' no esperado al abrir la base de datos", err)
	}
	defer db.Close()
	repo := NewRepository(db)

	expectedUser := domain.User{
		ID:       "123",
		Name:     "test",
		Age:      25,
		Email:    "estebanpoly@test",
		Password: "12345",
	}

	mockDB.ExpectPrepare(regexp.QuoteMeta(querySave)).
		ExpectExec().
		WithArgs(expectedUser.ID, expectedUser.Name, int8(expectedUser.Age), expectedUser.Email, expectedUser.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Save(expectedUser)
	if err != nil {
		t.Logf("Error: %v", err)
	}
	assert.NoError(t, err)
}

func TestSave_Error(t *testing.T) {
	db, mockDB, err := sqlmock.New() 
	if err != nil {
		t.Fatalf("un error '%s' no esperado al abrir la base de datos", err)
	}
	defer db.Close()
	repo := NewRepository(db)

	expectedUser := domain.User{
		ID:       "123",
		Name:     "test",
		Age:      25,
		Email:    "estebanpoly@test",
		Password: "12345",
	}	

	mockDB.ExpectPrepare(regexp.QuoteMeta(querySave)).
		ExpectExec().
		WithArgs(expectedUser.ID, expectedUser.Name, int8(expectedUser.Age), expectedUser.Email, expectedUser.Password).
		WillReturnError(sql.ErrConnDone)	

	err = repo.Save(expectedUser)
	if err != nil {
		t.Logf("Error: %v", err)
	}
	assert.Error(t, err)
}	
