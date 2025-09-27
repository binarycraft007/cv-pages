package main

import (
	"fmt"
	"html/template"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct to hold the data from config.yaml
type Config struct {
	Name        string       `yaml:"name"`
	Synopsis    Synopsis     `yaml:"synopsis"`
	Description string       `yaml:"description"`
	Skills      []string     `yaml:"skills"`
	Projects    []Project    `yaml:"projects"`
	Experience  []Experience `yaml:"experience"`
	Education   []Education  `yaml:"education"`
	SeeAlso     []Link       `yaml:"see_also"`
}

// Synopsis struct for contact information
type Synopsis struct {
	Email    string `yaml:"email"`
	Linkedin string `yaml:"linkedin"`
	Github   string `yaml:"github"`
}

// Project struct for project details
type Project struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	URL         string `yaml:"url"`
}

// Experience struct for job history
type Experience struct {
	Company     string `yaml:"company"`
	Title       string `yaml:"title"`
	Dates       string `yaml:"dates"`
	Description string `yaml:"description"`
}

// Education struct for academic background
type Education struct {
	Degree      string `yaml:"degree"`
	Institution string `yaml:"institution"`
	Dates       string `yaml:"dates"`
}

// Link struct for other online profiles
type Link struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func main() {
	// Read the config.yaml file
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("Error reading config.yaml: %v\n", err)
		return
	}

	// Unmarshal the YAML data into our Config struct
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Error unmarshaling YAML: %v\n", err)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html.tmpl")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the public directory if it doesn't exist
	if _, err := os.Stat("public"); os.IsNotExist(err) {
		os.Mkdir("public", 0755)
	}

	// Create the index.html file
	htmlFile, err := os.Create("public/index.html")
	if err != nil {
		fmt.Printf("Error creating index.html: %v\n", err)
		return
	}
	defer htmlFile.Close()

	// Execute the template with the config data and write to index.html
	err = tmpl.Execute(htmlFile, config)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	// Copy the style.css file to the public directory
	cssInput, err := os.ReadFile("css/style.css")
	if err != nil {
		fmt.Printf("Error reading style.css: %v\n", err)
		return
	}

	err = os.WriteFile("public/style.css", cssInput, 0644)
	if err != nil {
		fmt.Printf("Error writing style.css: %v\n", err)
		return
	}

	fmt.Println("Static site generated successfully in the 'public' directory.")
}
