/*
===============================================================================
BT PROJECT DIAGRAMS GENERATOR - COMPREHENSIVE CHART SYSTEM
===============================================================================

Author: Ben Tran
Date: 02/09/2025
Description:

Package main provides a comprehensive CLI tool to generate all project diagrams
and charts. This is the main orchestrator that coordinates all three chart
systems: SchemaERD, OGdiagrams, and StructureDiagrams.

Build: this file is compiled only with the 'flowcharts' build tag.
Usage: go run -tags flowcharts BTProjectDiagrams.go -interactive -out flowcharts

CHART SYSTEMS:
- BTProjectScanner: Automatic project scanning and function discovery
- SchemaERD: Database Entity-Relationship Diagrams
- OGdiagrams: Original educational diagrams (sequence, workout store, connections, brain)
- StructureDiagrams: Project structure analysis (development, execution, dependencies, building guide)
*/
package main

/* How to run
1) go mod tidy
2) go run -tags flowcharts BTProjectDiagrams.go -out flowcharts
3) go run

First-time setup (only once, if not already installed):
cd "C:\Users\Admin\Documents\GODEV\GO Courses\CompleteGO\Ex3Goose"
go mod tidy
go install github.com/ofabry/go-callvis@latest
go install github.com/loov/goda@latest
winget install --id Graphviz.Graphviz -e --silent --accept-package-agreements --accept-source-agreements

Generate flowcharts for this project (every time you want fresh graphs):
Generate charts (only compiles the tool):
 wcharts.go -out flowcharts

cd "C:\Users\Admin\Documents\GODEV\GO Courses\CompleteGO\Ex9"
go run -tags flowcharts BTProjectDiagrams.go -out flowcharts

Run the app (normal run):
cd "C:\Users\Admin\Documents\GODEV\GO Courses\CompleteGO\Ex3Goose"
go run .

Outputs:
flowcharts\graph.svg (functions call graph)
flowcharts\pkg-deps.svg (package dependencies)
This uses the Ex3Goose folder only for this example; future projects can run the same two commands in their own directories.

>>>>>> Alternative
Move the tool to tools/flowcharts/main.go
go run ./tools/flowcharts -out flowcharts


DETAIL FLOWCHART
cd "C:\Users\Admin\Documents\GODEV\GO Courses\CompleteGO\Ex10"
$env:PATH += ";C:\Program Files\Go\bin;C:\Users\Admin\go\bin;C:\Program Files\Graphviz\bin"
$env:PLANTUML_JAR = "C:\Users\Admin\plantuml.jar"
$env:DB_NAME = "postgres"
$env:DB_USER = "postgres"
$env:DB_PASS = "postgres"
go run -tags flowcharts BenTran_Project_builder/BTProjectDiagrams.go BenTran_Project_builder/AIAd_diagrams.go BenTran_Project_builder/Theory_diagrams.go BenTran_Project_builder/Existing_diagrams.go BenTran_Project_builder/SchemaERD.go BenTran_Project_builder/SVGCharts.go BenTran_Project_builder/Theory2Reality.go -interactive -out BTFlowcharts -root .

*/

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// main
// What: Parses flags and invokes BTFlowcharts.
// Why: Provides a simple CLI entrypoint to generate graphs.
// How: Reads -out and -root, then calls BTFlowcharts and exits on error.

// FlowchartOptions configures what to include in generated graphs.
type FlowchartOptions struct {
	NoStdlib      bool   // exclude stdlib from function call graph
	Group         string // grouping for go-callvis (e.g., "pkg,type")
	Focus         string // optional focus package/function for go-callvis
	Ignore        string // optional ignore regex for go-callvis
	IncludeTests  bool   // include tests in go-callvis graph
	GenerateUML   bool   // generate PlantUML class diagram if goplantuml is available
	Comprehensive bool   // also generate expanded charts under ComprehensiveCharts
}

func main() {
	outDir := flag.String("out", "BTFlowcharts", "output directory for generated graphs")
	root := flag.String("root", "", "project root (defaults to current working directory)")
	// Detail configuration flags
	noStd := flag.Bool("nostd", true, "exclude Go stdlib from function graph")
	group := flag.String("group", "pkg,type", "grouping for function graph (e.g., pkg,type)")
	focus := flag.String("focus", "", "optional focus (pkg or func regex) for function graph")
	ignore := flag.String("ignore", "", "optional ignore regex for function graph")
	tests := flag.Bool("tests", true, "include test files in function graph")
	uml := flag.Bool("uml", true, "generate PlantUML class diagram if goplantuml is installed")
	comprehensive := flag.Bool("comprehensive", true, "also generate expanded charts under ComprehensiveCharts")
	interactive := flag.Bool("interactive", false, "run in interactive mode with menu")
	flag.Parse()
	opts := FlowchartOptions{
		NoStdlib:      *noStd,
		Group:         *group,
		Focus:         *focus,
		Ignore:        *ignore,
		IncludeTests:  *tests,
		GenerateUML:   *uml,
		Comprehensive: *comprehensive,
	}

	if *interactive {
		runInteractiveMode(*root, *outDir, opts)
	} else {
		if err := BTFlowcharts(*root, *outDir, opts); err != nil {
			log.Fatalf("flowchart generation failed: %v", err)
		}
	}

}

// runInteractiveMode provides an interactive menu for chart generation
func runInteractiveMode(root, outDir string, opts FlowchartOptions) {
	fmt.Println("🎯 BT Project Diagrams - Interactive Mode")
	fmt.Println("==========================================")

	for {
		fmt.Println("\n📋 Available Chart Systems:")
		fmt.Println("1. Regenerate HTML Charts (default)")
		fmt.Println("2. Generate All Charts")
		fmt.Println("3. Project Scanner (Dynamic Reports)")
		fmt.Println("4. AI Advisor Diagrams (Project Recreation Guidance)")
		fmt.Println("5. Theory Model Diagrams (Educational Diagrams)")
		fmt.Println("6. SVG ComGo Deteail Model Diagrams (Instructor + AI)")
		fmt.Println("7. Schema ERD (Database Diagrams)")
		fmt.Println("8. Existing Diagrams (Current Project State Analysis)")
		fmt.Println("9. Theory to Reality Analysis (Implementation Progress)")
		fmt.Println("10. Model to Reality Analysis (Implementation Progress)")
		fmt.Println("11. AI Advisor Function Creation & Execution Order Diagrams")
		fmt.Println("12. Class Model Builder Teaching Guides")
		fmt.Println("99. 🔍 Project Status Evaluation & Assessment")
		fmt.Println("0. Exit")

		fmt.Print("\n🎯 Choose an option (1-12, 99) or press Enter to Regenerate HTML Charts: ")

		var choice string
		fmt.Scanln(&choice)

		// Default to "1" if empty input
		if choice == "" {
			choice = "1"
		}

		switch choice {
		case "1":
			fmt.Println("\n📋 Regenerating HTML Charts...")
			viewAllCurrentCharts(root, outDir)

		case "2":
			fmt.Println("\n🚀 Generating ALL charts (Schema ERD, Existing, Theory to Reality, Model to Reality)...")

			// Generate core charts first
			if err := BTFlowcharts(root, outDir, opts); err != nil {
				fmt.Printf("❌ Error generating core charts: %v\n", err)
			} else {
				fmt.Println("✅ Core charts generated successfully!")
			}

			// Generate Schema ERD (option 7)
			fmt.Println("\n🗄️ Generating Schema ERD...")
			// First scan the project to get current functions
			structure, err := Existing_scanProject(root)
			if err != nil {
				fmt.Printf("❌ Error scanning project: %v\n", err)
			} else {
				fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

				if err := generateSchemaSpyERD(root, outDir, structure); err != nil {
					fmt.Printf("❌ Error generating Schema ERD: %v\n", err)
				} else {
					fmt.Println("✅ Schema ERD generated successfully!")
				}
			}

			// Generate Existing Diagrams (option 8)
			fmt.Println("\n📊 Generating Existing Diagrams (Current Project State Analysis)...")
			if err := generateExistingDiagrams(root, outDir); err != nil {
				fmt.Printf("❌ Error generating existing diagrams: %v\n", err)
			} else {
				fmt.Println("✅ Existing diagrams generated successfully!")
			}

			// Generate Theory to Reality Analysis (option 9)
			fmt.Println("\n🔍 Generating Theory to Reality Analysis...")
			// First scan the project to get current functions
			structure, err = Existing_scanProject(root)
			if err != nil {
				fmt.Printf("❌ Error scanning project: %v\n", err)
			} else {
				fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

				if err := Theory2Reality_WriteAllAnalysis(outDir, structure); err != nil {
					fmt.Printf("❌ Error generating theory to reality analysis: %v\n", err)
				} else {
					fmt.Println("✅ Theory to reality analysis generated successfully!")
				}
			}

			// Generate Model to Reality Analysis (option 10)
			fmt.Println("\n🔍 Generating Model to Reality Analysis...")
			// First scan the project to get current functions
			structure, err = Existing_scanProject(root)
			if err != nil {
				fmt.Printf("❌ Error scanning project: %v\n", err)
			} else {
				fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

				if err := Theory2Reality_WriteAllAnalysis(outDir, structure); err != nil {
					fmt.Printf("❌ Error generating model to reality analysis: %v\n", err)
				} else {
					fmt.Println("✅ Model to reality analysis generated successfully!")
				}
			}

			fmt.Println("\n🎉 ALL charts generated successfully!")

		case "3":
			fmt.Println("\n🔍 Generating Project Scanner reports...")
			if err := generateScannerReports(root, outDir); err != nil {
				fmt.Printf("❌ Error generating scanner reports: %v\n", err)
			} else {
				fmt.Println("✅ Project scanner reports generated successfully!")
			}
		case "4":
			fmt.Println("\n🤖 Generating AI Advisor Diagrams (Project Recreation Guidance)...")
			if err := generateAIAdvisorDiagrams(outDir); err != nil {
				fmt.Printf("❌ Error generating AI advisor diagrams: %v\n", err)
			} else {
				fmt.Println("✅ AI advisor diagrams generated successfully!")
			}
		case "5":
			fmt.Println("\n🎓 Generating Theory Diagrams (Educational)...")
			if err := generateTheoryDiagrams(root, outDir); err != nil {
				fmt.Printf("❌ Error generating theory diagrams: %v\n", err)
			} else {
				fmt.Println("✅ Theory diagrams generated successfully!")
			}
		case "6":
			fmt.Println("\n🌐 Generating SVG Charts...")
			// First scan the project to get current functions
			structure, err := Existing_scanProject(root)
			if err != nil {
				fmt.Printf("❌ Error scanning project: %v\n", err)
				break
			}
			fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

			//NOTE - Omitted SVG Charts for now
			fmt.Printf("❌ Omitted SVG charts for Model to Reality Analysis: %v\n", err)
			// if err := generateSVGCharts(root, outDir, opts, structure); err != nil {
			// 	fmt.Printf("❌ Error generating SVG charts: %v\n", err)
			// } else {
			// 	fmt.Println("✅ SVG charts generated successfully!")
			// }
		case "7":
			fmt.Println("\n🗄️ Generating Schema ERD...")
			// First scan the project to get current functions
			structure, err := Existing_scanProject(root)
			if err != nil {
				fmt.Printf("❌ Error scanning project: %v\n", err)
				break
			}
			fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

			if err := generateSchemaSpyERD(root, outDir, structure); err != nil {
				fmt.Printf("❌ Error generating Schema ERD: %v\n", err)
			} else {
				fmt.Println("✅ Schema ERD generated successfully!")
			}
		case "8":
			fmt.Println("\n📊 Generating Existing Diagrams (Current Project State Analysis)...")
			if err := generateExistingDiagrams(root, outDir); err != nil {
				fmt.Printf("❌ Error generating existing diagrams: %v\n", err)
			} else {
				fmt.Println("✅ Existing diagrams generated successfully!")
			}
		case "9":
			fmt.Println("\n🔍 Generating Theory to Reality Analysis...")
			// First scan the project to get current functions
			structure, err := Existing_scanProject(root)
			if err != nil {
				fmt.Printf("❌ Error scanning project: %v\n", err)
				break
			}
			fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

			if err := Theory2Reality_WriteAllAnalysis(outDir, structure); err != nil {
				fmt.Printf("❌ Error generating theory to reality analysis: %v\n", err)
			} else {
				fmt.Println("✅ Theory to reality analysis generated successfully!")
			}
		case "11":
			fmt.Println("\n📊 Generating AI Advisor Function Creation & Execution Order Diagrams...")
			if err := AIAdCreate_Exe_WriteAllFunctionDiagrams(outDir); err != nil {
				fmt.Printf("❌ Error generating function diagrams: %v\n", err)
			} else {
				fmt.Println("✅ AI Advisor function creation and execution order diagrams generated successfully!")
			}
		case "12":
			fmt.Println("\n📚 Generating Class Model Builder Teaching Guides...")
			if err := ClassModelBuilder_WriteAllTeachingGuides(outDir); err != nil {
				fmt.Printf("❌ Error generating teaching guides: %v\n", err)
			} else {
				fmt.Println("✅ Class Model Builder teaching guides generated successfully!")
			}
		case "99":
			fmt.Println("\n🔍 Starting Project Status Evaluation & Assessment...")
			if err := ProjectEvaluator_WriteAllEvaluations(outDir); err != nil {
				fmt.Printf("❌ Error generating project evaluation: %v\n", err)
			} else {
				fmt.Println("✅ Project evaluation completed successfully!")
			}
		case "0":
			fmt.Println("\n👋 Goodbye!")
			return

		default:
			fmt.Printf("❌ Invalid choice: %s. Please choose 1-11 or 0.\n", choice)
		}

		// Ask if user wants to continue
		fmt.Print("\n🔄 Generate more charts? (y/N): ")
		var continueChoice string
		fmt.Scanln(&continueChoice)

		if continueChoice != "y" && continueChoice != "Y" && continueChoice != "yes" && continueChoice != "Yes" {
			fmt.Println("\n👋 Goodbye!")
			return
		}
	}
}

// generateScannerReports runs only the project scanner functionality
func generateScannerReports(root, outDir string) error {
	fmt.Println("🔍 Scanning project for functions...")
	structure, err := Existing_scanProject(root)
	if err != nil {
		return fmt.Errorf("project scanning failed: %w", err)
	}

	fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

	fmt.Println("📝 Generating dynamic reports...")
	return Existing_generateUpdatedReports(outDir, structure)
}

// generateCurrentProjectOGDiagrams runs the current project OG diagrams functionality
func generateCurrentProjectOGDiagrams(root, outDir string) error {
	fmt.Println("🎯 Generating current project OG diagrams...")

	// First scan the project to get current functions
	structure, err := Existing_scanProject(root)
	if err != nil {
		return fmt.Errorf("project scanning failed: %w", err)
	}

	fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

	// Generate current project OG diagrams
	// Theory diagrams are now reserved for future instructor+AI models
	fmt.Println("🎓 Theory diagrams are reserved for future instructor+AI models")
	return nil
}

// generateSVGCharts runs only the SVG chart generation functionality
// generateSVGCharts has been moved to SVGCharts.go for real project analysis

// generateSchemaERD runs only the Schema ERD functionality
func generateSchemaERD(root, outDir string) error {
	fmt.Println("🗄️ Generating Schema ERD...")
	return GenerateSchemaSpyERD(root, outDir)
}

// generateStructureDiagrams runs only the structure diagrams functionality
func generateStructureDiagrams(outDir string) error {
	fmt.Println("🏗️ Generating structure diagrams...")
	return AIAd_WriteAllStructureDiagrams(outDir)
}

// generateTheoryDiagrams runs the theory diagrams functionality
func generateTheoryDiagrams(root, outDir string) error {
	fmt.Println("🎓 Generating Lesson Model Diagrams (Instructor's Teaching Progression)...")

	// Generate lesson-based diagrams following instructor's progression
	return LessonModel_WriteAllLessonDiagrams(outDir)
}

// generateAIAdvisorDiagrams runs the AI advisor diagrams functionality
func generateAIAdvisorDiagrams(outDir string) error {
	fmt.Println("🤖 Generating AI advisor diagrams...")
	return AIAd_WriteAllStructureDiagrams(outDir)
}

// generateExistingDiagrams runs the existing diagrams functionality
func generateExistingDiagrams(root, outDir string) error {
	fmt.Println("📊 Generating existing diagrams...")

	// First scan the project to get current functions
	structure, err := Existing_scanProject(root)
	if err != nil {
		return fmt.Errorf("project scanning failed: %w", err)
	}

	fmt.Printf("✅ Found %d functions across %d files\n", len(structure.Functions), len(structure.Files))

	// Generate existing diagrams based on discovered functions
	if err := Existing_generateUpdatedReports(outDir, structure); err != nil {
		return err
	}

	// Generate architecture and file tree diagrams
	if err := Existing_WriteArchitectureDiagram(root, outDir); err != nil {
		return fmt.Errorf("architecture diagram failed: %w", err)
	}

	// Generate both simplified and full function dependency diagrams
	if err := Existing_WriteFunctionDependencyDiagram(root, outDir, 1); err != nil {
		return fmt.Errorf("simplified function dependency diagram failed: %w", err)
	}
	if err := Existing_WriteFunctionDependencyDiagram(root, outDir, 2); err != nil {
		return fmt.Errorf("full function dependency diagram failed: %w", err)
	}

	return nil
}

// viewAllCurrentCharts displays all currently available charts
func viewAllCurrentCharts(root, outDir string) {
	fmt.Println("🔄 Regenerating HTML charts from existing .mmd.md files...")

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Printf("❌ Error creating output directory: %v\n", err)
		return
	}

	// Convert all .mmd.md files to HTML
	htmlFilesCreated := 0
	err := filepath.Walk(outDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".mmd.md") {
			// Convert to HTML
			if err := convertMermaidFileToHTML(path); err != nil {
				fmt.Printf("⚠️  Warning: Could not convert %s to HTML: %v\n", filepath.Base(path), err)
			} else {
				htmlFile := strings.Replace(path, ".mmd.md", ".html", 1)
				fmt.Printf("✅ Created: %s\n", filepath.Base(htmlFile))
				htmlFilesCreated++
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("❌ Error scanning for .mmd.md files: %v\n", err)
		return
	}

	if htmlFilesCreated == 0 {
		fmt.Println("❌ No .mmd.md files found to convert to HTML.")
		fmt.Println("   Please generate charts first using options 2-7.")
		return
	}

	fmt.Printf("\n✅ Successfully created %d HTML files!\n", htmlFilesCreated)

	// Ask if user wants to open the HTML files
	fmt.Print("\n🌐 Open HTML charts in browser? (y/N): ")
	var openChoice string
	fmt.Scanln(&openChoice)

	if openChoice == "y" || openChoice == "Y" || openChoice == "yes" || openChoice == "Yes" {
		fmt.Println("🌐 Opening HTML charts in browser...")
		openChartsInBrowser(outDir)
	} else {
		fmt.Println("📁 HTML files created. You can open them manually from the BTFlowcharts folder.")
	}
}

// convertMermaidFileToHTML converts a single .mmd.md file to HTML
func convertMermaidFileToHTML(filePath string) error {
	// Read the .mmd.md file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Extract Mermaid content from the file
	var mermaidContent strings.Builder
	lines := strings.Split(string(content), "\n")
	inMermaidBlock := false

	for _, line := range lines {
		if strings.TrimSpace(line) == "```mermaid" {
			inMermaidBlock = true
			continue
		}
		if strings.TrimSpace(line) == "```" && inMermaidBlock {
			inMermaidBlock = false
			continue
		}
		if inMermaidBlock {
			mermaidContent.WriteString(line + "\n")
		}
	}

	// Create HTML file with Mermaid.js
	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <script src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js"></script>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 10px; }
        .mermaid { text-align: center; }
    </style>
</head>
<body>
    <div class="mermaid">
%s
    </div>
    <script>mermaid.initialize({startOnLoad:true});</script>
</body>
</html>`, mermaidContent.String())

	// Write HTML file
	htmlFile := strings.Replace(filePath, ".mmd.md", ".html", 1)
	return os.WriteFile(htmlFile, []byte(htmlContent), 0644)
}

// openChartsInBrowser opens all available HTML charts in the browser
func openChartsInBrowser(outDir string) {
	// List of HTML files to open
	htmlFiles := []string{
		"function_dependencies.html",
		"development_sequence.html",
		"execution_flow.html",
		"architecture.html",
		"file_tree.html",
		"sequence_template.html",
		"workout_store.html",
		"workout_store_connections.html",
		"application_brain.html",
		"current_application_brain.html",
		"current_store_connections.html",
		"dynamic_development_sequence.html",
		"BTspyERD/index.html",
	}

	// List of SVG files to open
	svgFiles := []string{
		"graph.svg",
		"graph_by_pkg.svg",
		"graph_full.svg",
		"graph_migrations.svg",
		"pkg-deps.svg",
		"types.svg",
	}

	openedCount := 0

	// Open HTML files
	for _, file := range htmlFiles {
		filePath := filepath.Join(outDir, file)
		if fileExists(filePath) {
			exec.Command("cmd", "/c", "start", filePath).Start()
			fmt.Printf("🌐 Opened %s\n", filepath.Base(file))
			openedCount++
		}
	}

	// Open SVG files
	for _, file := range svgFiles {
		filePath := filepath.Join(outDir, file)
		if fileExists(filePath) {
			exec.Command("cmd", "/c", "start", filePath).Start()
			fmt.Printf("🌐 Opened %s\n", filepath.Base(file))
			openedCount++
		}
	}

	if openedCount > 0 {
		fmt.Printf("✅ Opened %d charts in browser!\n", openedCount)
	} else {
		fmt.Println("❌ No charts found to open.")
	}
}

// BTFlowcharts
// What: Orchestrates generation of function, package, and optional UML graphs.
// Why: Single entry point to keep graphs up-to-date for large Go projects.
// How: Verifies required tools, creates output dir, runs go-callvis and goda+dot; optionally goplantuml.
func BTFlowcharts(projectRoot, outDir string, opts FlowchartOptions) error {
	// Determine working dir: prefer provided root, else CWD; then resolve module root (go.mod)
	wd := projectRoot
	if wd == "" {
		var err error
		wd, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("getwd: %w", err)
		}
	}
	if mr, ok := findModuleRoot(wd); ok {
		wd = mr
	}
	if err := ensureDir(filepath.Join(wd, outDir)); err != nil {
		return err
	}

	// Ensure tools exist
	if err := ensureTool("go-callvis"); err != nil {
		return wrapInstallHint(err, "go install github.com/ofabry/go-callvis@latest")
	}
	if err := ensureTool("goda"); err != nil {
		return wrapInstallHint(err, "go install github.com/loov/goda@latest")
	}
	if err := ensureTool("dot"); err != nil {
		return wrapInstallHint(err, "winget install --id Graphviz.Graphviz -e")
	}

	// Generate function call graph (graph.svg)
	callvisArgs := []string{"-format", "svg", "-file", filepath.Join(outDir, "graph.svg")}
	if opts.NoStdlib {
		callvisArgs = append(callvisArgs, "-nostd")
	}
	if opts.Group != "" {
		callvisArgs = append(callvisArgs, "-group", opts.Group)
	}
	if opts.Focus != "" {
		callvisArgs = append(callvisArgs, "-focus", opts.Focus)
	}
	if opts.Ignore != "" {
		callvisArgs = append(callvisArgs, "-ignore", opts.Ignore)
	}
	if opts.IncludeTests {
		callvisArgs = append(callvisArgs, "-tests")
	}
	callvisArgs = append(callvisArgs, "./...")
	if err := runInDir(wd, "go-callvis", callvisArgs...); err != nil {
		fmt.Printf("⚠️  go-callvis failed (expected with multiple main packages): %v\n", err)
		fmt.Println("   This is normal when running multiple chart files together.")
		fmt.Println("   Other charts will still be generated successfully.")
	}

	// Extra 1: generate a package-grouped call graph (alternative perspective)
	byPkg := append([]string{}, callvisArgs...)
	for i := range byPkg {
		if byPkg[i] == filepath.Join(outDir, "graph.svg") {
			byPkg[i] = filepath.Join(outDir, "graph_by_pkg.svg")
		}
	}
	if idx := indexOf(byPkg, "-group"); idx >= 0 && idx+1 < len(byPkg) {
		byPkg[idx+1] = "pkg"
	} else {
		byPkg = append([]string{"-group", "pkg"}, byPkg...)
	}
	if err := runInDir(wd, "go-callvis", byPkg...); err != nil {
		fmt.Println("Note: pkg-grouped graph generation failed (continuing):", err)
	}

	// Extra 2: generate a full graph including stdlib to surface DB/sql edges
	full := []string{"-format", "svg", "-file", filepath.Join(outDir, "graph_full.svg")}
	// do not add -nostd here on purpose
	if opts.Group != "" {
		full = append(full, "-group", opts.Group)
	}
	if opts.Focus != "" {
		full = append(full, "-focus", opts.Focus)
	}
	if opts.Ignore != "" {
		full = append(full, "-ignore", opts.Ignore)
	}
	if opts.IncludeTests {
		full = append(full, "-tests")
	}
	full = append(full, "./...")
	if err := runInDir(wd, "go-callvis", full...); err != nil {
		fmt.Println("Note: full stdlib-inclusive graph generation failed (continuing):", err)
	}

	// Extra 3: if a migrations package exists, generate a focused graph to surface those edges
	if dirExists(filepath.Join(wd, "migrations")) {
		focusVal := "migrations"
		if mod := readModulePath(filepath.Join(wd, "go.mod")); mod != "" {
			focusVal = mod + "/migrations"
		}
		mig := []string{"-format", "svg", "-file", filepath.Join(outDir, "graph_migrations.svg"), "-group", "pkg,type"}
		if opts.IncludeTests {
			mig = append(mig, "-tests")
		}
		mig = append(mig, "-focus", focusVal, "./...")
		if err := runInDir(wd, "go-callvis", mig...); err != nil {
			fmt.Println("Note: migrations-focused graph generation failed (continuing):", err)
		}
	}

	// Generate package dependency graph (pkg-deps.dot -> .svg)
	dotPath := filepath.Join(outDir, "pkg-deps.dot")
	svgPath := filepath.Join(outDir, "pkg-deps.svg")
	// Note: We capture 'goda graph' output to a .dot file explicitly.
	// If you only need the file, the prior invocation can be skipped.
	// Pipe is not as portable; call `goda graph` to file via cmd redirection
	if err := writeFileFromCmd(wd, []string{"goda", "graph", "./..."}, dotPath); err != nil {
		return fmt.Errorf("write dot: %w", err)
	}
	if err := runInDir(wd, "dot", "-Tsvg", dotPath, "-o", svgPath); err != nil {
		return fmt.Errorf("dot convert: %w", err)
	}

	// Optionally generate a PlantUML class diagram of structs/interfaces if available.
	if opts.GenerateUML {
		if err := ensureTool("goplantuml"); err == nil {
			umlPath := filepath.Join(outDir, "types.puml")
			if err := writeFileFromCmd(wd, []string{"goplantuml", "-recursive", "."}, umlPath); err != nil {
				return fmt.Errorf("goplantuml: %w", err)
			}
			// Render types.puml to SVG if PlantUML (or plantuml.jar + java) is available.
			if cmd, args, ok := findPlantUMLRenderer(); ok {
				if err := runInDir(filepath.Join(wd, outDir), cmd, append(args, "types.puml")...); err != nil {
					fmt.Println("Note: PlantUML render failed (continuing):", err)
				}
			} else {
				fmt.Println("Note: types.puml generated; PlantUML not found on PATH. Install PlantUML or set PLANTUML_JAR to render SVG.")
				fmt.Println("Install hints: go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest ; winget install --id PlantUML.PlantUML -e or set $env:PLANTUML_JAR")
			}
		} else {
			fmt.Println("Note: skipping UML generation (goplantuml not found)")
		}
	}

	// Step 1: Scan project for functions and generate dynamic reports
	fmt.Println("🔍 Scanning project for functions and files...")
	structure, err := Existing_scanProject(wd)
	if err != nil {
		fmt.Printf("⚠️  Project scan failed: %v (continuing with static charts)\n", err)
	} else {
		// Generate dynamic reports based on discovered functions
		if err := Existing_generateUpdatedReports(outDir, structure); err != nil {
			fmt.Printf("⚠️  Dynamic reports failed: %v (continuing with static charts)\n", err)
		} else {
			fmt.Printf("✅ Generated dynamic reports: %d functions across %d files\n", len(structure.Functions), len(structure.Files))
		}
	}

	// Step 2: Generate static educational charts
	// Bonus: emit a lightweight Mermaid architecture diagram for higher-level relationships.
	_ = Existing_WriteArchitectureDiagram(wd, outDir)
	// Emit a Mermaid file/package tree for quick project overview.
	//_ = Existing_WriteFileTreeDiagram(wd, outDir)
	// Generate current project OG diagrams based on discovered functions
	// if structure != nil {
	// 	_ = Theory_WriteProjectOGDiagrams(outDir, structure)
	// }
	// Emit function flow analysis diagrams for learning and development guidance.
	_ = AIAd_WriteFunctionFlowAnalysis(outDir)
	// Optionally generate ERD via SchemaSpy if environment is configured and user agrees.
	//_ = GenerateSchemaSpyERD(wd, outDir)
	// SchemaSpy ERD generation moved to individual options to avoid duplicate prompts

	fmt.Printf("Generated:\n- %s\n- %s\n", filepath.Join(outDir, "types.svg"), svgPath)
	// return nil

	// Open the generated files - this is the new way to open the files
	fmt.Printf("Generated:\n- %s\n- %s\n", filepath.Join(outDir, "graph.svg"), svgPath)

	// Always open all charts at the end (required)
	openAllCharts(outDir)
	return nil
}

// ensureTool checks if a tool is present in PATH.
func ensureTool(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("missing tool %q: %w", name, err)
	}
	return nil
}

// wrapInstallHint adds a short install hint to an error (keeps original error wrapped).
func wrapInstallHint(err error, hint string) error {
	return fmt.Errorf("%w\nInstall hint: %s", err, hint)
}

// runInDir has been moved to SchemaERD.go to avoid conflicts

// writeFileFromCmd runs a command and writes its stdout to a file.
func writeFileFromCmd(dir string, cmdArgs []string, outPath string) error {
	if len(cmdArgs) == 0 {
		return errors.New("writeFileFromCmd: empty cmd")
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	if err := os.WriteFile(outPath, out, 0644); err != nil {
		return err
	}
	return nil
}

// ensureDir has been moved to SchemaERD.go to avoid conflicts

// findPlantUMLRenderer returns a command and args to render PlantUML to SVG.
// Prefers "plantuml" if present; otherwise uses "java -jar <plantuml.jar> -tsvg" if PLANTUML_JAR
// env var or common install paths are found.
func findPlantUMLRenderer() (string, []string, bool) {
	if _, err := exec.LookPath("plantuml"); err == nil {
		return "plantuml", []string{"-tsvg"}, true
	}
	jar := os.Getenv("PLANTUML_JAR")
	if jar == "" {
		// Get user home directory for cross-platform compatibility
		homeDir, _ := os.UserHomeDir()
		candidates := []string{
			`C:\\Program Files\\PlantUML\\plantuml.jar`,
			`C:\\Program Files (x86)\\PlantUML\\plantuml.jar`,
			`C:\\Program Files\\Common Files\\PlantUML\\plantuml.jar`,
		}
		// Add user home directory candidates
		if homeDir != "" {
			candidates = append(candidates,
				filepath.Join(homeDir, "plantuml.jar"),
				filepath.Join(homeDir, "Downloads", "plantuml.jar"),
				filepath.Join(homeDir, "Documents", "plantuml.jar"),
			)
		}
		for _, c := range candidates {
			if fileExists(c) {
				jar = c
				break
			}
		}
	}
	if jar != "" {
		if _, err := exec.LookPath("java"); err == nil {
			return "java", []string{"-jar", jar, "-tsvg"}, true
		}
	}
	return "", nil, false
}

// writeMermaidDiagram writes a simple architecture diagram to help visualize high-level flows.
// writeMermaidDiagram and writeMermaidFileTree have been moved to Existing_diagrams.go
// as Existing_WriteArchitectureDiagram and Existing_WriteFileTreeDiagram

// writeSequenceTemplate has been moved to OGdiagrams.go

// writeWorkoutStoreDiagram has been moved to OGdiagrams.go

// writeOGWorkoutStoreDiagram has been moved to OGdiagrams.go

// writeWorkoutStoreConnectionsDiagram has been moved to OGdiagrams.go

// writeApplicationBrainDiagram has been moved to OGdiagrams.go

// SchemaSpy ERD generation has been moved to SchemaERD.go
// This allows for better separation of concerns and user interaction

// indexOf returns the first index of target in slice or -1 if not found.
func indexOf(slice []string, target string) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// readModulePath returns the module path from go.mod if available.
func readModulePath(goModPath string) string {
	f, err := os.Open(goModPath)
	if err != nil {
		return ""
	}
	defer f.Close()
	buf := make([]byte, 256)
	n, _ := f.Read(buf)
	if n <= 0 {
		return ""
	}
	text := string(buf[:n])
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return ""
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// findModuleRoot walks upward from startDir to locate the directory containing go.mod.
// Returns the found directory and true if located; otherwise returns startDir and false.
func findModuleRoot(startDir string) (string, bool) {
	dir := startDir
	for {
		if fileExists(filepath.Join(dir, "go.mod")) {
			return dir, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return startDir, false
		}
		dir = parent
	}
}

// fileExists has been moved to SchemaERD.go to avoid conflicts

// writeFunctionFlowAnalysis has been moved to StructureDiagrams.go

// writeDevelopmentSequenceDiagram has been moved to StructureDiagrams.go

// writeExecutionFlowDiagram has been moved to StructureDiagrams.go

// writeFunctionDependencyDiagram has been moved to StructureDiagrams.go

// writeProjectBuildingGuide has been moved to StructureDiagrams.go

// openAllCharts opens all generated charts (required for BTFlowcharts)
func openAllCharts(outDir string) {
	// Open ERD using the new SchemaERD functionality
	OpenERDInBrowser(outDir)

	// Open SVG files
	svgFiles := []string{
		filepath.Join(outDir, "graph.svg"),
		filepath.Join(outDir, "graph_by_pkg.svg"),
		filepath.Join(outDir, "graph_full.svg"),
		filepath.Join(outDir, "graph_migrations.svg"),
		filepath.Join(outDir, "pkg-deps.svg"),
		filepath.Join(outDir, "types.svg"),
	}

	for _, svgFile := range svgFiles {
		if fileExists(svgFile) {
			exec.Command("cmd", "/c", "start", svgFile).Start()
			fmt.Printf("Opened %s\n", filepath.Base(svgFile))
		}
	}

	// Create and open HTML versions of Mermaid files
	createMermaidHTML(outDir)
}

func createMermaidHTML(outDir string) {
	mermaidFiles := []string{
		filepath.Join(outDir, "Existing_architecture.mmd.md"),
		filepath.Join(outDir, "Existing_function_dependencies_simplified.mmd.md"),
		filepath.Join(outDir, "Existing_function_dependencies_full.mmd.md"),
		filepath.Join(outDir, "Existing_application_brain.mmd.md"),
		filepath.Join(outDir, "Existing_store_connections.mmd.md"),
		filepath.Join(outDir, "AIAd_development_sequence.mmd.md"),
		filepath.Join(outDir, "AIAd_execution_flow.mmd.md"),
		filepath.Join(outDir, "AIAd_function_dependencies.mmd.md"),
		filepath.Join(outDir, "Existing_dynamic_development_sequence.mmd.md"),
		filepath.Join(outDir, "AIAdCreate_Exe_function_creation_order.mmd.md"),
		filepath.Join(outDir, "AIAdCreate_Exe_function_execution_order.mmd.md"),
		filepath.Join(outDir, "ClassModelBuilder_complete_project_guide.mmd.md"),
		filepath.Join(outDir, "ClassModelBuilder_step_by_step_workflow.mmd.md"),
		filepath.Join(outDir, "ClassModelBuilder_file_creation_sequence.mmd.md"),
		filepath.Join(outDir, "ClassModelBuilder_function_implementation_guide.mmd.md"),
		filepath.Join(outDir, "ClassModelBuilder_folder_structure_guide.mmd.md"),
		filepath.Join(outDir, "ProjectEvaluator_comprehensive_assessment.mmd.md"),
	}

	for _, file := range mermaidFiles {
		// Read the .mmd file content
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		// Extract Mermaid content from markdown code blocks
		contentStr := string(content)
		lines := strings.Split(contentStr, "\n")
		var mermaidContent strings.Builder

		inMermaidBlock := false
		for _, line := range lines {
			if strings.TrimSpace(line) == "```mermaid" {
				inMermaidBlock = true
				continue
			}
			if inMermaidBlock && strings.TrimSpace(line) == "```" {
				break
			}
			if inMermaidBlock {
				mermaidContent.WriteString(line + "\n")
			}
		}

		// Create HTML file with Mermaid.js and high-resolution settings
		htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Function Dependencies - High Resolution</title>
    <script src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js"></script>
    <style>
        body { 
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; 
            margin: 0; 
            padding: 20px; 
            background-color: #f8f9fa;
        }
        .container {
            max-width: 100%%;
            margin: 0 auto;
            background: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        .mermaid { 
            text-align: center; 
            font-size: 14px;
            line-height: 1.4;
        }
        h1 { 
            color: #2c3e50; 
            text-align: center;
            border-bottom: 3px solid #3498db;
            padding-bottom: 10px;
            margin-bottom: 30px;
        }
        .info {
            background-color: #e8f4f8;
            padding: 15px;
            border-radius: 5px;
            margin-bottom: 20px;
            border-left: 4px solid #3498db;
        }
        .info h3 {
            margin-top: 0;
            color: #2980b9;
        }
        /* High-resolution print styles */
        @media print {
            body { background: white; }
            .container { box-shadow: none; }
            .mermaid { 
                font-size: 12px;
                page-break-inside: avoid;
            }
        }
        /* High-resolution screen styles */
        @media screen {
            .mermaid { 
                font-size: 16px;
                zoom: 1.2;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🔗 Function Dependencies Diagram</h1>
        <div class="info">
            <h3>📊 High-Resolution View</h3>
            <p>This diagram shows the dependency relationships between functions in your project. 
            Use Ctrl+Plus to zoom in for better readability, or print to PDF for high-quality output.</p>
        </div>
        <div class="mermaid">
%s
        </div>
    </div>
    <script>
        mermaid.initialize({
            startOnLoad: true,
            theme: 'default',
            flowchart: {
                useMaxWidth: true,
                htmlLabels: true,
                curve: 'basis'
            },
            themeVariables: {
                fontSize: '16px',
                fontFamily: 'Segoe UI, Tahoma, Geneva, Verdana, sans-serif'
            }
        });
    </script>
</body>
</html>`, mermaidContent.String())

		// Write HTML file
		htmlFile := strings.Replace(file, ".mmd.md", ".html", 1)
		os.WriteFile(htmlFile, []byte(htmlContent), 0644)

		// Open HTML file in browser
		exec.Command("cmd", "/c", "start", htmlFile).Start()
		fmt.Printf("Created and opened %s\n", filepath.Base(htmlFile))
	}
}

/*

# 1) Try to locate plantuml.jar
$jar = Get-ChildItem -Path "C:\Program Files","C:\Program Files (x86)","$env:Admin" -Filter plantuml*.jar -Recurse -ErrorAction SilentlyContinue | Select-Object -First 1 -Expand FullName

# 2) If not found, DOWNLOAD the jar locally ++++++++++++++++++++++++++++++++++++++++
if (-not $jar) {
  $dest = "$env:USERPROFILE\plantuml.jar"
  Invoke-WebRequest -Uri "https://github.com/plantuml/plantuml/releases/latest/download/plantuml.jar" -OutFile $dest
  $jar = $dest
}

# 3) Set env and verify Java
$env:PLANTUML_JAR = $jar
java -jar "$env:PLANTUML_JAR" -version

$env:PLANTUML_JAR = "C:\Users\Admin\plantuml.jar"

   [System.Environment]::SetEnvironmentVariable("PLANTUML_JAR", "C:\Users\Admin\plantuml.jar", "User")

#4) FOR SCHEMASPY_JAR
4) SchemaSpy
# Set the path to the SchemaSpy JAR file
[System.Environment]::SetEnvironmentVariable("SCHEMASPY_JAR", "C:\tools\schemaspy\schemaspy.jar", "User")

# Set the path to the PostgreSQL JDBC Driver JAR file
[System.Environment]::SetEnvironmentVariable("PG_JDBC_JAR", "C:\tools\schemaspy\postgresql-driver.jar", "User")

Write-Host "Environment variables SCHEMASPY_JAR and PG_JDBC_JAR have been set."
Write-Host "IMPORTANT: Please close and reopen your terminal for the changes to take effect."

5) Poweshell
$env:DB_NAME="postgres"
$env:DB_USER="postgres"
$env:DB_PASS="postgres" # Use the password from your docker-compose.yml

FINAL: DETAIL FLOWCHART
cd "C:\Users\Admin\Documents\GODEV\GO Courses\CompleteGO\Ex10"
$env:PATH="$env:PATH;C:\Program Files\Go\bin;C:\Users\Admin\go\bin;C:\Program Files\Graphviz\bin"
$env:PLANTUML_JAR = "C:\Users\Admin\plantuml.jar"
$env:DB_NAME="postgres"
$env:DB_USER="postgres"
$env:DB_PASS="postgres"
$env:SCHEMASPY_JAR="C:\tools\schemaspy\schemaspy.jar"
$env:PG_JDBC_JAR="C:\tools\schemaspy\postgresql-driver.jar"

# Interactive Mode (Recommended)
go run -tags flowcharts BenTran_Project_builder/BTProjectDiagrams.go BenTran_Project_builder/BTProjectScanner.go BenTran_Project_builder/SchemaERD.go BenTran_Project_builder/OGdiagrams.go BenTran_Project_builder/StructureDiagrams.go -interactive -out BTFlowcharts -root .
# All Charts at Once
go run -tags flowcharts BenTran_Project_builder/BTProjectDiagrams.go BenTran_Project_builder/BTProjectScanner.go BenTran_Project_builder/SchemaERD.go BenTran_Project_builder/OGdiagrams.go BenTran_Project_builder/StructureDiagrams.go -out BTFlowcharts

Current File Structure:
Function Distribution:
BTProjectDiagrams.go (14 functions):
main() - CLI entry point
BTFlowcharts() - Main orchestrator
ensureTool(), wrapInstallHint() - Tool management
writeFileFromCmd(), findPlantUMLRenderer() - Command execution
writeMermaidDiagram(), writeMermaidFileTree() - Mermaid generation
indexOf(), readModulePath(), dirExists(), findModuleRoot() - Utilities
openAllCharts(), createMermaidHTML() - File management
BTProjectScanner.go (9 functions):
Existing_scanProject() - Project scanning
extractFunctions() - Function extraction
generateUpdatedReports() - Report orchestration
generateFunctionInventory() - Function inventory
generateDynamicDevelopmentSequence() - Dynamic sequence
generateProjectStatusReport() - Status report
categorizeFunctions() - Function grouping
determinePhase() - Phase determination
getSimplePurpose() - Purpose analysis
Generated Files (All Dynamic):
Dynamic Reports (from Existing_diagrams.go):
Existing_function_inventory.md - Comprehensive function list
Existing_dynamic_development_sequence.mmd.md - Auto-generated development order
Existing_project_status_report.md - Current project statistics
Static Charts (from other files):
All the existing educational diagrams and charts
Usage:


*/
