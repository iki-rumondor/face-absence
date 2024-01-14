package utils

import "path"

type StorageURLGenerator struct {
	BaseURL string
}
	
func (g *StorageURLGenerator) GenerateStorageURL(filePath string) string {
	return path.Join(g.BaseURL, filePath)
}
