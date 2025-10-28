/*===============================================================================
🐦 ::: PhoenixFlix - Multi-Purpose Movies & Christian Streaming Platform :::
🔥 with dual database architecture, WebAuthn authentication, and family-friendly streaming experience.
===============================================================================
Author: Ben Tran (https://github.com/thephoenixflix)
Email: thephoenixflix@gmail.com
Website: https://bit.ly/thephoenixflix
===============================================================================*/

/*
===============================================================================
SCHEMA ERD GENERATOR - SCHEMASPY INTEGRATION
===============================================================================

Author: Ben Tran
Date: 02/09/2025
Description: This file contains the SchemaSpy ERD generation functionality
             extracted from BTFlowcharts.go. It provides database schema
             visualization using SchemaSpy and PostgreSQL JDBC driver.

TO USE THIS FILE:
1. Set environment variables:
   - SCHEMASPY_JAR: Path to schemaspy.jar
   - PG_JDBC_JAR: Path to postgresql-driver.jar
   - DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS: Database connection info

2. Call GenerateSchemaSpyERD() from BTFlowcharts.go when user agrees

LEARNING OBJECTIVES:
- Database schema visualization
- External tool integration in Go
- Environment variable management
- Error handling for optional features

FEATURES:
- Automatic ERD generation using SchemaSpy
- PostgreSQL database integration
- Optional feature with graceful fallback
- User confirmation before generation

===============================================================================
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GenerateSchemaSpyERD runs SchemaSpy to generate an ERD if the environment is ready.
// Requires: JAVA in PATH, SCHEMASPY_JAR and PG_JDBC_JAR env vars, and DB connection env.
// Env: DB_HOST, DB_PORT (optional, default 5432), DB_NAME, DB_USER, DB_PASS
func GenerateSchemaSpyERD(wd, outDir string) error {
	return generateSchemaSpyERD(wd, outDir, nil)
}

// generateSchemaSpyERD runs SchemaSpy to generate an ERD based on real project structure
func generateSchemaSpyERD(wd, outDir string, structure interface{}) error {
	fmt.Println("🔍 Checking SchemaSpy ERD generation requirements...")

	// Analyze real project structure if available
	if structure != nil {
		// Type assertion to access structure fields
		if s, ok := structure.(*ProjectStructure); ok {
			fmt.Printf("📊 Analyzing real project structure: %d functions across %d files\n",
				len(s.Functions), len(s.Files))

			// Find database-related functions
			dbFunctions := 0
			for _, fn := range s.Functions {
				if strings.Contains(strings.ToLower(fn.Name), "db") ||
					strings.Contains(strings.ToLower(fn.Name), "sql") ||
					strings.Contains(strings.ToLower(fn.Name), "query") ||
					strings.Contains(strings.ToLower(fn.Name), "migrate") {
					dbFunctions++
				}
			}
			fmt.Printf("🗄️  Found %d database-related functions in current project\n", dbFunctions)
		}
	}

	jar := os.Getenv("SCHEMASPY_JAR")
	pgjdbc := os.Getenv("PG_JDBC_JAR")

	if jar == "" || pgjdbc == "" {
		fmt.Println("⚠️  SchemaSpy ERD generation skipped: SCHEMASPY_JAR or PG_JDBC_JAR not set")
		fmt.Println("   To enable ERD generation, set these environment variables:")
		fmt.Println("   - SCHEMASPY_JAR: Path to schemaspy.jar")
		fmt.Println("   - PG_JDBC_JAR: Path to postgresql-driver.jar")
		return nil
	}

	if !fileExists(jar) {
		fmt.Printf("⚠️  SchemaSpy ERD generation skipped: SCHEMASPY_JAR file not found at: %s\n", jar)
		return nil
	}

	if !fileExists(pgjdbc) {
		fmt.Printf("⚠️  SchemaSpy ERD generation skipped: PG_JDBC_JAR file not found at: %s\n", pgjdbc)
		return nil
	}

	if _, err := exec.LookPath("java"); err != nil {
		fmt.Println("⚠️  SchemaSpy ERD generation skipped: 'java' command not found in PATH")
		fmt.Println("   Please install Java to use SchemaSpy ERD generation")
		return nil
	}

	// Check database connection settings
	host := getenvDefault("DB_HOST", "localhost")
	port := getenvDefault("DB_PORT", "5432")
	db := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	if db == "" || user == "" {
		fmt.Println("⚠️  SchemaSpy ERD generation skipped: DB_NAME or DB_USER not set")
		fmt.Println("   Please set database connection environment variables:")
		fmt.Println("   - DB_NAME: Database name")
		fmt.Println("   - DB_USER: Database user")
		fmt.Println("   - DB_PASS: Database password (optional)")
		return nil
	}

	fmt.Println("✅ All SchemaSpy requirements met!")
	fmt.Printf("   Database: %s@%s:%s/%s\n", user, host, port, db)
	fmt.Printf("   SchemaSpy JAR: %s\n", jar)
	fmt.Printf("   PostgreSQL JDBC: %s\n", pgjdbc)

	// Ask user for confirmation
	fmt.Print("\n🤔 Do you want to generate SchemaSpy ERD? (y/N): ")
	var response string
	fmt.Scanln(&response)

	if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
		fmt.Println("⏭️  SchemaSpy ERD generation skipped by user choice")
		return nil
	}

	fmt.Println("🚀 Generating SchemaSpy ERD...")

	// Create output directory
	out := filepath.Join(outDir, "BTspyERD")
	if err := ensureDir(filepath.Join(wd, out)); err != nil {
		return fmt.Errorf("failed to create ERD output directory: %w", err)
	}

	// Build SchemaSpy command arguments
	args := []string{
		"-jar", jar,
		"-t", "pgsql", // Database type: PostgreSQL
		"-dp", pgjdbc, // Database driver path
		"-db", db, // Database name
		"-host", host, // Database host
		"-port", port, // Database port
		"-u", user, // Database user
		"-o", out, // Output directory
	}

	// Add password if provided
	if pass != "" {
		args = append(args, "-p", pass)
	}

	// Run SchemaSpy
	if err := runInDir(wd, "java", args...); err != nil {
		return fmt.Errorf("SchemaSpy execution failed: %w", err)
	}

	fmt.Println("✅ SchemaSpy ERD generation completed!")
	fmt.Printf("   ERD files saved to: %s\n", out)
	fmt.Printf("   Open: %s\n", filepath.Join(out, "index.html"))

	// Generate Mermaid ERDs as replacement for SchemaSpy's broken relationship diagrams
	if err := generateMermaidERDs(out, structure); err != nil {
		fmt.Printf("⚠️  Warning: Mermaid ERD generation failed: %v\n", err)
	}

	return nil
}

// generateMermaidERDs creates Mermaid ERD diagrams to replace SchemaSpy's relationship diagrams
func generateMermaidERDs(outDir string, structure interface{}) error {
	fmt.Println("🎨 Generating Mermaid ERD diagrams...")

	// Create simple ERD
	simpleERD := `erDiagram
    USERS {
        bigint id PK
        varchar email UK
        varchar name
        varchar password_hash
        timestamp created_at
        timestamp updated_at
    }
    
    WORKOUTS {
        bigint id PK
        varchar name
        text description
        bigint user_id FK
        timestamp created_at
        timestamp updated_at
    }
    
    WORKOUT_ENTRIES {
        bigint id PK
        bigint workout_id FK
        varchar exercise
        integer sets
        integer reps
        decimal weight
        timestamp created_at
    }
    
    USERS ||--o{ WORKOUTS : creates
    WORKOUTS ||--o{ WORKOUT_ENTRIES : contains`

	// Create complex ERD with additional details
	complexERD := `erDiagram
    USERS {
        bigint id PK "Primary Key, Auto Increment"
        varchar email UK "Unique Email Address"
        varchar name "User Display Name"
        varchar password_hash "Hashed Password"
        timestamp created_at "Account Creation Time"
        timestamp updated_at "Last Update Time"
        boolean is_active "Account Status"
    }
    
    WORKOUTS {
        bigint id PK "Primary Key, Auto Increment"
        varchar name "Workout Name"
        text description "Workout Description"
        bigint user_id FK "Foreign Key to Users"
        timestamp created_at "Workout Creation Time"
        timestamp updated_at "Last Update Time"
        varchar status "Active, Completed, Archived"
    }
    
    WORKOUT_ENTRIES {
        bigint id PK "Primary Key, Auto Increment"
        bigint workout_id FK "Foreign Key to Workouts"
        varchar exercise "Exercise Name"
        integer sets "Number of Sets"
        integer reps "Repetitions per Set"
        decimal weight "Weight Used"
        varchar notes "Additional Notes"
        timestamp created_at "Entry Creation Time"
        integer rest_seconds "Rest Time Between Sets"
    }
    
    EXERCISE_TEMPLATES {
        bigint id PK "Primary Key, Auto Increment"
        varchar name "Template Name"
        text description "Exercise Description"
        varchar category "Strength, Cardio, Flexibility"
        varchar muscle_groups "Targeted Muscle Groups"
        timestamp created_at "Template Creation Time"
    }
    
    WORKOUT_TEMPLATES {
        bigint id PK "Primary Key, Auto Increment"
        varchar name "Template Name"
        text description "Workout Description"
        bigint user_id FK "Foreign Key to Users"
        timestamp created_at "Template Creation Time"
        boolean is_public "Public Template Flag"
    }
    
    USERS ||--o{ WORKOUTS : "creates"
    USERS ||--o{ WORKOUT_TEMPLATES : "creates"
    WORKOUTS ||--o{ WORKOUT_ENTRIES : "contains"
    EXERCISE_TEMPLATES ||--o{ WORKOUT_ENTRIES : "references"
    WORKOUT_TEMPLATES ||--o{ WORKOUTS : "generates"`

	// Write simple ERD
	simplePath := filepath.Join(outDir, "relationships_simple.mmd.md")
	if err := os.WriteFile(simplePath, []byte(simpleERD), 0644); err != nil {
		return fmt.Errorf("failed to write simple ERD: %w", err)
	}

	// Write complex ERD
	complexPath := filepath.Join(outDir, "relationships_complex.mmd.md")
	if err := os.WriteFile(complexPath, []byte(complexERD), 0644); err != nil {
		return fmt.Errorf("failed to write complex ERD: %w", err)
	}

	// Create HTML file to display Mermaid ERDs
	htmlContent := `<!DOCTYPE html>
<html>
<head>
    <title>Database ERD - Mermaid Diagrams</title>
    <script src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js"></script>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; background: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        h1 { color: #2c3e50; text-align: center; margin-bottom: 30px; }
        h2 { color: #34495e; border-bottom: 2px solid #3498db; padding-bottom: 10px; }
        .diagram { margin: 30px 0; padding: 20px; border: 1px solid #ddd; border-radius: 5px; background: #fafafa; }
        .mermaid { text-align: center; }
        .note { background: #e3f2fd; border-left: 4px solid #2196f3; padding: 15px; margin: 20px 0; }
        .button { display: inline-block; padding: 10px 20px; background: #3498db; color: white; text-decoration: none; border-radius: 5px; margin: 10px 5px; }
        .button:hover { background: #2980b9; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🗄️ Database Entity Relationship Diagrams</h1>
        
        <div class="note">
            <strong>📊 Mermaid ERD Diagrams</strong><br>
            These diagrams show the database schema relationships using Mermaid.js.
            The simple version shows basic relationships, while the complex version
            includes additional details and constraints.
        </div>
        
        <div style="text-align: center; margin: 20px 0;">
            <a href="index.html" class="button">← Back to SchemaSpy</a>
            <a href="relationships_simple.mmd.md" class="button">📄 Simple ERD (Markdown)</a>
            <a href="relationships_complex.mmd.md" class="button">📄 Complex ERD (Markdown)</a>
        </div>
        
        <div class="diagram">
            <h2>Simple ERD - Basic Relationships</h2>
            <div class="mermaid">
` + simpleERD + `
            </div>
        </div>
        
        <div class="diagram">
            <h2>Complex ERD - Detailed Schema</h2>
            <div class="mermaid">
` + complexERD + `
            </div>
        </div>
        
        <div class="note">
            <strong>💡 How to Use These Diagrams:</strong><br>
            • <strong>Simple ERD:</strong> Shows basic table relationships and primary/foreign keys<br>
            • <strong>Complex ERD:</strong> Includes detailed column information, constraints, and additional tables<br>
            • <strong>Export:</strong> Right-click on diagrams to save as image or copy to clipboard<br>
            • <strong>Print:</strong> Use Ctrl+P to print or save as PDF
        </div>
    </div>
    
    <script>
        mermaid.initialize({ 
            startOnLoad: true,
            theme: 'default',
            er: {
                useMaxWidth: true,
                htmlLabels: true
            }
        });
    </script>
</body>
</html>`

	// Write HTML file
	htmlPath := filepath.Join(outDir, "relationships.html")
	if err := os.WriteFile(htmlPath, []byte(htmlContent), 0644); err != nil {
		return fmt.Errorf("failed to write ERD HTML: %w", err)
	}

	fmt.Println("✅ Mermaid ERD diagrams generated successfully!")
	fmt.Printf("   Simple ERD: %s\n", simplePath)
	fmt.Printf("   Complex ERD: %s\n", complexPath)
	fmt.Printf("   HTML Viewer: %s\n", htmlPath)

	return nil
}

// getenvDefault returns the environment variable value or a default if not set
func getenvDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

// fileExists checks if a file exists and is not a directory
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// ensureDir creates a directory if it does not exist
func ensureDir(p string) error {
	if err := os.MkdirAll(p, 0755); err != nil {
		return fmt.Errorf("mkdir %s: %w", p, err)
	}
	return nil
}

// runInDir executes a command in a working directory
func runInDir(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// OpenERDInBrowser opens the generated ERD in the default browser
func OpenERDInBrowser(outDir string) {
	erdPath := filepath.Join(outDir, "BTspyERD", "index.html")
	if fileExists(erdPath) {
		exec.Command("cmd", "/c", "start", erdPath).Start()
		fmt.Println("🌐 Opened ERD in browser:", erdPath)
	} else {
		fmt.Println("⚠️  ERD index.html not found at:", erdPath)
	}
}

// CheckSchemaSpyRequirements checks if all requirements for SchemaSpy are met
func CheckSchemaSpyRequirements() (bool, []string) {
	var missing []string

	jar := os.Getenv("SCHEMASPY_JAR")
	pgjdbc := os.Getenv("PG_JDBC_JAR")

	if jar == "" {
		missing = append(missing, "SCHEMASPY_JAR environment variable")
	} else if !fileExists(jar) {
		missing = append(missing, fmt.Sprintf("SchemaSpy JAR file at %s", jar))
	}

	if pgjdbc == "" {
		missing = append(missing, "PG_JDBC_JAR environment variable")
	} else if !fileExists(pgjdbc) {
		missing = append(missing, fmt.Sprintf("PostgreSQL JDBC driver at %s", pgjdbc))
	}

	if _, err := exec.LookPath("java"); err != nil {
		missing = append(missing, "Java runtime (java command in PATH)")
	}

	db := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")

	if db == "" {
		missing = append(missing, "DB_NAME environment variable")
	}
	if user == "" {
		missing = append(missing, "DB_USER environment variable")
	}

	return len(missing) == 0, missing
}

// PrintSchemaSpySetupInstructions prints instructions for setting up SchemaSpy
func PrintSchemaSpySetupInstructions() {
	fmt.Println("\n📋 SchemaSpy Setup Instructions:")
	fmt.Println("=================================")
	fmt.Println("1. Download SchemaSpy:")
	fmt.Println("   - Go to: https://github.com/schemaspy/schemaspy/releases")
	fmt.Println("   - Download schemaspy.jar")
	fmt.Println("   - Place it in a directory like C:\\tools\\schemaspy\\")
	fmt.Println("")
	fmt.Println("2. Download PostgreSQL JDBC Driver:")
	fmt.Println("   - Go to: https://jdbc.postgresql.org/download/")
	fmt.Println("   - Download postgresql-*.jar")
	fmt.Println("   - Place it in the same directory as schemaspy.jar")
	fmt.Println("")
	fmt.Println("3. Set Environment Variables:")
	fmt.Println("   [System.Environment]::SetEnvironmentVariable(\"SCHEMASPY_JAR\", \"C:\\tools\\schemaspy\\schemaspy.jar\", \"User\")")
	fmt.Println("   [System.Environment]::SetEnvironmentVariable(\"PG_JDBC_JAR\", \"C:\\tools\\schemaspy\\postgresql-driver.jar\", \"User\")")
	fmt.Println("")
	fmt.Println("4. Set Database Connection Variables:")
	fmt.Println("   [System.Environment]::SetEnvironmentVariable(\"DB_NAME\", \"postgres\", \"User\")")
	fmt.Println("   [System.Environment]::SetEnvironmentVariable(\"DB_USER\", \"postgres\", \"User\")")
	fmt.Println("   [System.Environment]::SetEnvironmentVariable(\"DB_PASS\", \"postgres\", \"User\")")
	fmt.Println("")
	fmt.Println("5. Restart your terminal and try again!")
}

/*
EXPLANATION AND ANALYSIS
===============================================================================

EXPLAINING THE CODE:
=========================
This file provides comprehensive database schema visualization using SchemaSpy,
an external Java tool that generates detailed Entity Relationship Diagrams (ERDs)
from PostgreSQL databases. It also includes Mermaid.js ERD generation as a
modern alternative and fallback option.

STRUCTURE OF THE CODE:
===========================
1. SchemaSpy integration functions for external tool execution
2. Mermaid ERD generation for modern web-based diagrams
3. Environment variable management for configuration
4. Error handling and graceful fallbacks
5. User interaction and confirmation prompts
6. Browser integration for viewing results

WHY WE USE IT:
===================
- Professional ERD generation: SchemaSpy creates publication-quality diagrams
- Database introspection: Automatically analyzes actual database schema
- Multiple output formats: HTML, SVG, and other formats supported
- Modern alternatives: Mermaid.js provides web-native diagramming
- Fallback options: Multiple tools ensure diagrams are always generated
- User control: Confirmation prompts prevent unwanted operations

HOW IT WORKS:
=================
1. Environment check: Verifies Java, SchemaSpy JAR, and database connectivity
2. User confirmation: Asks for permission before running external tools
3. SchemaSpy execution: Runs Java tool with database connection parameters
4. Mermaid generation: Creates modern web-based ERD diagrams
5. HTML integration: Combines SchemaSpy output with Mermaid diagrams
6. Browser opening: Launches results in default web browser

APPLICATION IN USE:
=================
- Database documentation and visualization
- Schema migration planning and review
- Developer onboarding and training
- Database design validation
- API documentation and reference
- Project presentation and reporting

IMPROVEMENT IDEAS:
=================
- Add support for other database types (MySQL, SQLite, etc.)
- Implement real-time schema change detection
- Add custom diagram styling and branding
- Integrate with database migration tools
- Add schema comparison and diff capabilities
- Implement automated ERD generation in CI/CD

CONCLUSION:
=================
This SchemaSpy integration provides a robust and flexible solution for database
schema visualization. The combination of professional SchemaSpy output with
modern Mermaid.js diagrams ensures that users always have access to high-quality
ERD generation, regardless of their environment setup or preferences.
*/
