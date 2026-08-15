package metadata

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSidecar(t *testing.T) {
	dir := t.TempDir()
	image := filepath.Join(dir, "cover.jpg")
	if err := os.WriteFile(image, []byte("image"), 0600); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "metadata.yml")
	content := `title: My Audiobook
subtitle: Optional
release_date: "2024-01-15"
publisher_name: Publisher
publisher_summary: Summary
authors:
  - name: Author
narrators:
  - name: Narrator
product_images:
  "500": ./cover.jpg
category_ladders:
  - root: Fiction
    ladder:
      - id: fiction
        name: Fiction
`
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	product, err := LoadSidecar(path)
	if err != nil {
		t.Fatal(err)
	}
	if product.Title != "My Audiobook" || product.Authors[0].Name != "Author" || product.CategoryLadders[0].Ladder[0].ID != "fiction" {
		t.Fatalf("decoded product fields incorrectly: %+v", product)
	}
	if product.ProductImages.Image500 != image {
		t.Fatalf("image path = %q, want %q", product.ProductImages.Image500, image)
	}
}

func TestLoadSidecarMissingAndMalformed(t *testing.T) {
	product, err := LoadSidecar(filepath.Join(t.TempDir(), "metadata.yml"))
	if err != nil || product != nil {
		t.Fatalf("missing sidecar = (%v, %v), want (nil, nil)", product, err)
	}
	path := filepath.Join(t.TempDir(), "metadata.yml")
	if err := os.WriteFile(path, []byte("title: [unterminated"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadSidecar(path); err == nil {
		t.Fatal("malformed YAML should return an error")
	}
}
