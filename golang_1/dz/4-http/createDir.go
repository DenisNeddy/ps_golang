// **Описание**: Создайте функцию, которая принимает название
// пакета и создает базовую структуру Go пакета с файлом
// пакета и тестовым файлом.
// **Входные данные**: packageName (строка) - название пакета
// для создания
// **Выходные данные**: error - ошибка создания пакета
// (nil если успешно)
// **Ограничения**: packageName не пустой, содержит только
// строчные буквы и подчеркивания
// **Примеры**:
// Input: packageName = "utils"
// Output: nil (создается папка utils с utils.go и utils_test.go)
//
// Input: packageName = "http_client"
// Output: nil (создается папка http_client с http_client.go
// и http_client_test.go)
import (
  "errors"
  "os"
  "regexp"
  "path/filepath"
  "fmt"
)
func createPackageStructure(packageName string) error {
    if packageName == "" {
      return errors.New("Name is empty")
    }
    validate := `^[a-z_]+$`
    matched, err := regexp.MatchString(validate, packageName)
    if err != nil {
      return err
    }

    if !matched {
      return errors.New("incorrect symbols in package name")
    }

    if _, err := os.Stat(packageName); err == nil {
      return errors.New("directory already exists")
    }

    err = os.MkdirAll(packageName, 0755)
    if err != nil {
      return err
    }
    mainFile := filepath.Join(packageName, packageName + ".go")
    mainFileContent := fmt.Sprintf("package %s\n", packageName)
    err = os.WriteFile(mainFile, []byte(mainFileContent), 0644)
    if err != nil {
      
      os.RemoveAll(packageName)
      return err
    }
    testFile := filepath.Join(packageName, packageName + "_test.go")
    testFileContent := fmt.Sprintf(`package %s

import "testing"

func Test%s(t *testing.T) {}
`, packageName, toCamelCase(packageName))
    err = os.WriteFile(testFile, []byte(testFileContent), 0644)
    if err != nil {
      os.RemoveAll(packageName)
     
      return err
    }
    return nil
}

func toCamelCase(name string) string {
    parts := strings.Split(name, "_")
    for i, part := range parts {
        if len(part) > 0 {
            parts[i] = strings.Title(part) // Пишет первую букву заглавной
        }
    }
    return strings.Join(parts, "")
}