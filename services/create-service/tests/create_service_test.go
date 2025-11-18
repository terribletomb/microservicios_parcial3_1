package tests

import (
	"context"
	"testing"

	"create-service/models"
	"create-service/services"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type fakeRepo struct{}

func (f *fakeRepo) Insertar(ctx context.Context, libro *models.Libro) (*models.Libro, error) {
	libro.ID = primitive.NewObjectID()
	return libro, nil
}

// Añadimos los métodos que faltan para satisfacer la interfaz LibroRepositorio.
// No necesitan hacer nada para esta prueba específica.
func (f *fakeRepo) BuscarPorID(ctx context.Context, id string) (*models.Libro, error) {
	return nil, nil
}
func (f *fakeRepo) BuscarTodos(ctx context.Context) ([]*models.Libro, error) {
	return nil, nil
}
func (f *fakeRepo) Actualizar(ctx context.Context, id string, datos bson.M) (*models.Libro, error) {
	return nil, nil
}
func (f *fakeRepo) Eliminar(ctx context.Context, id string) error {
	return nil
}

func TestCrearLibro(t *testing.T) {
	t.Run("debe crear un libro exitosamente", func(t *testing.T) {
		s := services.NuevoServicioCrear(&fakeRepo{})
		libro := &models.Libro{
			Titulo:    "El Principito",
			Autor:     "Antoine de Saint-Exupéry",
			Anio:      1943,
			Paginas:   96,
			Editorial: "Anagrama",
		}
		creado, err := s.CrearLibro(context.Background(), libro)
		assert.NoError(t, err)
		assert.NotNil(t, creado)
		assert.NotEmpty(t, creado.ID)
	})
}
