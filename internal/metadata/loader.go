package metadata

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"yt2abs/internal/types"
)

// LoadSidecar loads metadata.yml. A missing file is intentionally not an error.
// Local image paths are made relative to the sidecar's directory.
func LoadSidecar(path string) (*types.Product, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot access metadata file %q: %w", path, err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("metadata path %q is a directory", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read metadata file %q: %w", path, err)
	}
	var product types.Product
	if err := yaml.Unmarshal(data, &product); err != nil {
		return nil, fmt.Errorf("invalid YAML in metadata file %q: %w", path, err)
	}
	if product.ProductImages.Image500 != "" && !isRemote(product.ProductImages.Image500) && !filepath.IsAbs(product.ProductImages.Image500) {
		product.ProductImages.Image500 = filepath.Join(filepath.Dir(path), product.ProductImages.Image500)
	}
	return &product, nil
}

func isRemote(source string) bool {
	u, err := url.Parse(source)
	return err == nil && (strings.EqualFold(u.Scheme, "http") || strings.EqualFold(u.Scheme, "https"))
}
