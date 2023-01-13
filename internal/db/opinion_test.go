//go:build integration
// +build integration

package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/untaldouglas/go-rest-api/internal/opinion"
	"testing"
)

func TestOpinionDatabase(t *testing.T) {
	t.Run("test create opinion", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		opi, err := db.PostOpinion(context.Background(), opinion.Opinion{
			Asunto:    "asunto desde testing",
			Contenido: "Contenido desde testing",
			Autor:     "Autor desde testing",
		})
		assert.NoError(t, err)

		newOpi, err := db.GetOpinion(context.Background(), opi.ID)
		assert.NoError(t, err)
		assert.Equal(t, "asunto desde testing", newOpi.Asunto)

	})

	t.Run("test borrar opinion", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		opi, err := db.PostOpinion(context.Background(), opinion.Opinion{
			Asunto:    "asunto a borrar desde testing",
			Contenido: "Contenido a borrar desde testing",
			Autor:     "Autor a borrar desde testing",
		})
		assert.NoError(t, err)
		err = db.DeleteOpinion(context.Background(), opi.ID)
		assert.NoError(t, err)

		_, err = db.GetOpinion(context.Background(), opi.ID)
		assert.Error(t, err)

	})
}
