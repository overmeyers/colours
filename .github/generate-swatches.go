package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Colour struct {
	Name  string `json:"name"`
	Hex   string `json:"hex"`
	RGB   string `json:"rgb"`
	HSL   string `json:"hsl"`
	Ratio string `json:"ratio"`
	Note  string `json:"note,omitempty"`
}

type ColourData struct {
	Colours []Colour `json:"colours"`
}

func parseRGB(rgb string) (uint8, uint8, uint8, error) {
	parts := strings.Fields(rgb)
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid RGB format: %s", rgb)
	}

	r, err := strconv.ParseUint(parts[0], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid red value: %w", err)
	}
	g, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid green value: %w", err)
	}
	b, err := strconv.ParseUint(parts[2], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid blue value: %w", err)
	}

	return uint8(r), uint8(g), uint8(b), nil
}

func generateSwatchImage(rgb string, width, height int, outputPath string) error {
	r, g, b, err := parseRGB(rgb)
	if err != nil {
		return fmt.Errorf("invalid RGB colour %s: %w", rgb, err)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	col := color.RGBA{R: r, G: g, B: b, A: 255}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, col)
		}
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create image file: %w", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	return nil
}

func generateSwatchTable(colours []Colour) string {
	var sb strings.Builder

	// Table header
	sb.WriteString("| swatch                                             | name       | rgb           | hex      | hsl              | ratio to base |\n")
	sb.WriteString("|:--------------------------------------------------:|:-----------|---------------|----------|------------------|--------------:|\n")

	// Table rows
	for _, colour := range colours {
		swatchPath := fmt.Sprintf("assets/swatches/%s.png", strings.ToLower(colour.Name))
		swatch := fmt.Sprintf("![#%s](%s)", colour.Hex, swatchPath)

		ratio := colour.Ratio
		if colour.Note != "" {
			ratio = fmt.Sprintf("%s[^%s]", colour.Ratio, colour.Note)
		}

		sb.WriteString(fmt.Sprintf("| %s | %-10s | `%-13s` | `%s` | `%-16s` | %13s |\n",
			swatch,
			colour.Name,
			colour.RGB,
			colour.Hex,
			colour.HSL,
			ratio,
		))
	}

	return sb.String()
}

func main() {
	jsonData, err := os.ReadFile("colours.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading colours.json: %v\n", err)
		os.Exit(1)
	}

	var colourData ColourData
	if err := json.Unmarshal(jsonData, &colourData); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing colours.json: %v\n", err)
		os.Exit(1)
	}

	assetsDir := filepath.Join("assets", "swatches")
	if err := os.MkdirAll(assetsDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating assets directory: %v\n", err)
		os.Exit(1)
	}

	for _, colour := range colourData.Colours {
		outputPath := filepath.Join(assetsDir, fmt.Sprintf("%s.png", strings.ToLower(colour.Name)))
		if err := generateSwatchImage(colour.RGB, 30, 30, outputPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating swatch for %s: %v\n", colour.Name, err)
			os.Exit(1)
		}
	}
	fmt.Printf("Generated %d swatch images in %s\n", len(colourData.Colours), assetsDir)

	readme, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading README.md: %v\n", err)
		os.Exit(1)
	}

	newTable := generateSwatchTable(colourData.Colours)

	tableStartPattern := regexp.MustCompile(`\| swatch\s+\| name\s+\| rgb\s+\| hex\s+\| hsl\s+\| ratio to base \|`)
	match := tableStartPattern.FindIndex(readme)

	if match == nil {
		fmt.Fprintln(os.Stderr, "Could not find the colour table in README.md")
		os.Exit(1)
	}

	tableStartIndex := match[0]

	afterTableStart := readme[tableStartIndex:]
	tableEndPattern := regexp.MustCompile(`\n\n`)
	endMatch := tableEndPattern.FindIndex(afterTableStart)

	if endMatch == nil {
		fmt.Fprintln(os.Stderr, "Could not find the end of the colour table")
		os.Exit(1)
	}

	tableEndIndex := tableStartIndex + endMatch[0]

	var newReadme strings.Builder
	newReadme.Write(readme[:tableStartIndex])
	newReadme.WriteString(newTable)
	newReadme.Write(readme[tableEndIndex:])

	if err := os.WriteFile("README.md", []byte(newReadme.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing README.md: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("README.md updated successfully with colour swatches from colours.json")
	fmt.Printf("Total: %d colours processed\n", len(colourData.Colours))
}
