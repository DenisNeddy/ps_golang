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

func TestUpdateBin(t *testing.T) {
	// Setup
	store := setupTest(t)

	createdBinIds := []string{}
	defer teardownTest(t, createdBinIds)

	// Arrange - создаем тестовый bin

	testName := "test_update_bin_original"
	binId := createTestBin(t, store, testName)

	createdBinIds = append(createdBinIds, binId)

	// Добавляем еще один ишт для обновления

	updatedName := "test_update_bin_updated"
	newBin := bins.NewBin(binId, false, updatedName)
	store.Bins = []bins.Bin{*newBin}
	store.List = [][2]string{{binId, updatedName}}

	// Сохраняем в отдельный файл для обновления

	err := store.Save("test_update.json")

	if err != nil {
		t.Fatalf("Не удалось сохранить файл для обновления: %v", err)
	}

	// Act - обновляем bin на сервере

	updatedStore, err := api.ChangeBin("test_update.json", binId)

	// Assert

	if err != nil {
		t.Errorf("Ожидалось nil, получено: %v", err)
	}

	if updatedStore == nil {
		t.Fatal("Обновленный store не должен быть nil")
	}

	if len(updatedStore.Bins) == 0 {
		t.Fatal("Обновленный store должен содеражать bins")
	}

	found := false

	for _, bin := range updatedStore.Bins {
		if bin.Id == binId && bin.Name == updatedName {
			found = true
			break
		}
	}

	if !found {
		t.Error("Bin с ID '%s' и именем '%s' не найден после обновления", binId, updatedName)
	}
}

func TestGetBin(t *testing.T) {
	// Setup

	store := setupTest(t)
	createdBinIds := []string{}

	defer teardownTest(t, createdBinIds)

	// Arrange - создаем тестовый bin
	testName := "test_get_bin"
	binId := createTestBin(t, store, testName)
	createdBinIds = append(createdBinIds, binId)

	// Act - получение bin с сервера

	retrievedStore, err := api.GetBin(binId)

	if err != nil {
		t.Errorf("Ожидалось nil, получено: %v", err)
	}

	if retrievedStore == nil {
		t.Fatal("Полученный store не должен быть nil")
	}

	if len(retrievedStore.Storage.Bins) == 0 {
		t.Fatal("Полученный store должен содержать bins")
	}

	// Проверяем что получили правильный bin
	found := false

	for _, bin := range retrievedStore.Storage.Bins {
		if bin.Id == binId && bin.Name == testName {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Bin с ID '%s' и именем '%s' не найден в полученных данных", binId, testName)
	}
}

// Test на удаление bin

func TestDeleteBin(t *testing.T) {
	// Setup

	store := setupTest(t)
	createdBinIds := []string{}
	defer teardownTest(t, createdBinIds)

	// Arrange - создаем тестовый bin
	testName := "test_delete_bin"
	binId := createTestBin(t, store, testName)
	createdBinIds = append(createdBinIds, binId)

	initialCount := len(store.Bins)

	// Act - удаляем bin
	err := api.DeleteBin(binId)

	// Assert
	if err != nil {
		t.Errorf("Ожидалось nil, получено: %v", err)
	}

	// Перезагружаем store из файла, чтобы проверить удаление

	err = store.Load("test_data.json")

	if err != nil {
		t.Fatalf("Не удалось загрузить store после удаления: %v", err)
	}

	// Проверяем, что bin удален из локального хранилища
	if len(store.Bins) >= initialCount {
		t.Errorf("Ожидалось меньше %d bins после удаления, получено %d", initialCount, len(store.Bins))
	}

	// Проверяем, что удаленный bin отсутстует

	for _, bin := range store.Bins {
		if bin.Id == binId {
			t.Errorf("Bin с ID '%s' все еще присутствует после удаления", binId)
		}
	}

	// Удаляем ID из списка для cleanup(уже удален)

	for i, id := range createdBinIds {
		if id == binId {
			createdBinIds = append(createdBinIds[:i], createdBinIds[i+1:]...)
			break
		}
	}
}
