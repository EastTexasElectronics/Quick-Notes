// file_utils.go
package utils

import (
    "os"
)

// SaveNoteToFile saves the note to a file in the specified format.
func SaveNoteToFile(filePath, content string) error {
    return os.WriteFile(filePath, []byte(content), 0644)
}
