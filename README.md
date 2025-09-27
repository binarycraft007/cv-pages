# CV Pages

This project is a simple static site generator for creating a personal CV or resume page. It uses Go to parse a YAML configuration file and render an HTML page from a template. The styling is done with a simple CSS file, and the project includes a GitHub Actions workflow for automatic deployment to GitHub Pages.

## Features

- **Simple Configuration**: Easily update your CV by editing a single `config.yaml` file.
- **Go Templating**: Uses Go's built-in `html/template` package for rendering the site.
- **Minimalist Design**: A clean, single-column layout with a dark theme.
- **Automated Deployment**: Includes a GitHub Actions workflow to build and deploy the site to GitHub Pages automatically.

## Getting Started

### Prerequisites

- Go (version 1.21 or later)
- A GitHub account

### Instructions

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/binarycraft007/cv-pages.git
    cd cv-pages
    ```

2.  **Update the configuration:**
    Edit the `config.yaml` file to include your personal information, skills, projects, experience, and education.

3.  **Build the site:**
    Run the following command to generate the static files in a `public` directory:
    ```bash
    go run main.go
    ```

4.  **Preview the site:**
    Open the `public/index.html` file in your browser to see the generated CV.

## Configuration

The `config.yaml` file is the single source of truth for your CV. It has the following structure:

```yaml
name: "Your Name"
synopsis:
  email: "your.email@example.com"
  linkedin: "https://linkedin.com/in/your-profile"
  github: "https://github.com/your-username"
description: "A brief description of yourself."
skills:
  - "Skill 1"
  - "Skill 2"
projects:
  - title: "Project Title"
    description: "Project description."
    url: "https://github.com/your-username/project"
experience:
  - company: "Company Name"
    title: "Your Title"
    dates: "Year - Year"
    description: "Description of your role and responsibilities."
education:
  - degree: "Your Degree"
    institution: "University Name"
    dates: "Year - Year"
```

## Deployment

This project is set up for automatic deployment to GitHub Pages. The `.github/workflows/deploy.yml` file defines a GitHub Actions workflow that:

1.  **Triggers** on every push to the `main` branch.
2.  **Checks out** the code.
3.  **Sets up Go**.
4.  **Builds** the static site by running `go run main.go`.
5.  **Deploys** the contents of the `public` directory to your GitHub Pages site.

To enable this, you need to:

1.  Push your code to a GitHub repository.
2.  In your repository settings, go to **Pages**.
3.  Under **Build and deployment**, select **GitHub Actions** as the source.
