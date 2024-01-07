package customHTTP

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")

	// Path lengkap ke file
	filePath := filepath.Join("internal/assets/temp", filename)

	// Cek apakah file ada
	_, err := os.Stat(filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Set header untuk download
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	// Buka file
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// Salin isi file ke response writer
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file content"})
		return
	}
}
