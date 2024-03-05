// Пакет keyboard позволяет принимать данные от пользователя и
// конвертировать их под необходимые типы данных
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Принимает данные от пользователя и возращает значение
// float64 с ошибкой (nil если ошибка отсутсвует)
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
