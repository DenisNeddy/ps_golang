package api_test

import (
	"os"
	"project/download-json/api"
	"project/download-json/bins"
	"project/download-json/storage"
	"testing"

	"github.com/joho/godotenv"
)

func setupTest(t *testing.T) *storage.Storage {
	// Загузка .env
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Не удалось загрузить .env файл")
	}

	// Создать пустой Storage
	store := storage.NewStorage([]bins.Bin{})

	// Сохранить в тестовый файл
	err = store.Save("test_data.json")
	if err != nil {
		t.Fatal("Не удалось создать тестовый файл")
	}

	return store
}

func teardownTest(t *testing.T, binIds []string) {
	// Удаляем все созданные bins

	for _, id := range binIds {
		api.DeleteBin(id)
	}

	// Удалить тестовый файл
	os.Remove("test_data.json")
	os.Remove("test_update.json")
}

func createTestBin(t *testing.T, store *storage.Storage, name string) string {
	err := api.CreateBin2(store, "test_data.json", false, name)

	if err != nil {
		t.Fatalf("Не удалось создать тестовый bin: %v", err)
	}

	return store.Bins[len(store.Bins)-1].Id
}

func TestCreateBin(t *testing.T) {
	// Setup

	store := setupTest(t)
	createdBinIds := []string{}
	defer teardownTest(t, createdBinIds)

	// Arrange

	testName := "test_create_bin"
	initialCount := len(store.Bins)

	// Act

	err := api.CreateBin2(store, "test_data.json", false, testName)

	// Assert

	if err != nil {
		t.Errorf("Ожидалось nil, получено: %v", err)
	}

	if len(store.Bins) != initialCount+1 {
		t.Errorf("Ожидалось %d bins, получено %d", initialCount+1, len(store.Bins))
	}

	lastBin := store.Bins[len(store.Bins)-1]

	if lastBin.Name != testName {
		t.Errorf("Ожидалось имя '%s', получено '%s'", testName, lastBin.Name)
	}

	// Сохранить ID Для cleanup

	createdBinIds = append(createdBinIds, lastBin.Id)

}
