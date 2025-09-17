//go:build flowcharts
// +build flowcharts

/*
===============================================================================
EXISTING DIAGRAMS - CURRENT PROJECT STATE ANALYSIS
===============================================================================

Author: Ben Tran
Date: 02/09/2025
Description: This file contains functions that analyze the current state of
             the project by scanning actual files and functions. These functions
             generate diagrams showing what actually exists in the current
             project, real function locations, and current project statistics.

TO USE THIS FILE:
1. Call Existing_scanProject() to analyze current project state
2. Each function generates analysis of actual project content
3. Diagrams are saved as Existing_*.mmd.md files for easy identification

EXISTING OBJECTIVES:
- Current project state analysis
- Real function discovery and mapping
- Actual project statistics
- Current vs. ideal comparison

FEATURES:
- Existing_function_inventory.md - Current function inventory
- Existing_dynamic_development_sequence.mmd.md - Current development sequence
- Existing_project_status_report.md - Current project statistics
- Existing_current_application_brain.mmd.md - Current application brain

===============================================================================
*/

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// FunctionInfo represents a discovered function
type FunctionInfo struct {
	Name     string
	File     string
	Package  string
	Line     int
	IsMethod bool
	Receiver string
	Purpose  string
}

// ProjectStructure represents the discovered project structure
type ProjectStructure struct {
	Functions []FunctionInfo
	Files     []string
	Packages  map[string][]string
}

// Existing_scanProject scans the project directory for Go files and extracts function information
func Existing_scanProject(rootDir string) (*ProjectStructure, error) {

	structure := &ProjectStructure{
		Functions: []FunctionInfo{},
		Files:     []string{},
		Packages:  make(map[string][]string),
	}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories and vendor
		if info.IsDir() {
			// Don't skip the root directory itself
			if path == rootDir {
				return nil
			}
			if strings.HasPrefix(info.Name(), ".") || info.Name() == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		// Only process .go files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Skip test files for now (can be made configurable)
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		// Extract functions from this file
		functions, err := Existing_extractFunctions(path)
		if err != nil {
			return err
		}

		structure.Functions = append(structure.Functions, functions...)
		structure.Files = append(structure.Files, path)

		// Group by package
		if len(functions) > 0 {
			pkg := functions[0].Package
			structure.Packages[pkg] = append(structure.Packages[pkg], path)
		}

		return nil
	})

	return structure, err
}

// Existing_extractFunctions extracts function information from a Go file
func Existing_extractFunctions(filePath string) ([]FunctionInfo, error) {
	var functions []FunctionInfo

	// Parse the file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// Extract package name
	packageName := node.Name.Name

	// Walk the AST to find functions
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			// Include all functions (both exported and unexported)
			// This gives a complete picture of the project structure

			funcInfo := FunctionInfo{
				Name:     x.Name.Name,
				File:     filePath,
				Package:  packageName,
				Line:     fset.Position(x.Pos()).Line,
				IsMethod: x.Recv != nil,
				Purpose:  Existing_getSimplePurpose(FunctionInfo{Name: x.Name.Name, File: filePath}),
			}

			// Extract receiver for methods
			if x.Recv != nil && len(x.Recv.List) > 0 {
				if ident, ok := x.Recv.List[0].Type.(*ast.Ident); ok {
					funcInfo.Receiver = ident.Name
				}
			}

			functions = append(functions, funcInfo)
		}
		return true
	})

	return functions, nil
}

// Existing_generateUpdatedReports generates updated flowcharts and documentation
func Existing_generateUpdatedReports(outDir string, structure *ProjectStructure) error {
	// Generate function inventory
	if err := Existing_generateFunctionInventory(outDir, structure); err != nil {
		return err
	}

	// Generate updated development sequence
	if err := Existing_generateDynamicDevelopmentSequence(outDir, structure); err != nil {
		return err
	}

	// Generate project status report
	if err := Existing_generateProjectStatusReport(outDir, structure); err != nil {
		return err
	}

	// Generate application brain diagram
	if err := Existing_WriteApplicationBrainDiagram(outDir, structure); err != nil {
		return err
	}

	// Generate store connections diagram
	if err := Existing_WriteStoreConnectionsDiagram(outDir, structure); err != nil {
		return err
	}

	return nil
}

// Existing_generateFunctionInventory creates a comprehensive inventory of all functions
func Existing_generateFunctionInventory(outDir string, structure *ProjectStructure) error {
	var content strings.Builder

	content.WriteString("# Existing Function Inventory - Auto-Generated\n\n")
	content.WriteString("This document provides a comprehensive inventory of all functions currently existing in the project.\n\n")

	// Group functions by package
	packageGroups := Existing_categorizeFunctions(structure.Functions)

	for pkg, functions := range packageGroups {
		content.WriteString(fmt.Sprintf("## Package: %s\n\n", pkg))
		content.WriteString(fmt.Sprintf("**Files:** %d  |  **Functions:** %d\n\n", len(structure.Packages[pkg]), len(functions)))

		// Sort functions by name
		sort.Slice(functions, func(i, j int) bool {
			return functions[i].Name < functions[j].Name
		})

		for _, fn := range functions {
			content.WriteString(fmt.Sprintf("- **%s**", fn.Name))
			if fn.IsMethod {
				content.WriteString(fmt.Sprintf(" (method on %s)", fn.Receiver))
			}
			content.WriteString(fmt.Sprintf(" - %s\n", fn.Purpose))
			content.WriteString(fmt.Sprintf("  - File: `%s` (line %d)\n", fn.File, fn.Line))
		}
		content.WriteString("\n")
	}

	content.WriteString(fmt.Sprintf("## Summary\n\n"))
	content.WriteString(fmt.Sprintf("- **Total Functions:** %d\n", len(structure.Functions)))
	content.WriteString(fmt.Sprintf("- **Total Files:** %d\n", len(structure.Files)))
	content.WriteString(fmt.Sprintf("- **Total Packages:** %d\n", len(structure.Packages)))

	path := filepath.Join(outDir, "Existing_function_inventory.md")
	return os.WriteFile(path, []byte(content.String()), 0644)
}

// Existing_generateDynamicDevelopmentSequence creates an updated development sequence based on discovered functions
func Existing_generateDynamicDevelopmentSequence(outDir string, structure *ProjectStructure) error {
	content := `# Existing Dynamic Development Sequence - Auto-Generated

This diagram shows the **order in which functions should be created** based on the current project structure.
Understanding this helps you know **where to start** when building similar projects.

` + "```mermaid\n" + `
flowchart TD
`

	// Group functions by phase
	phaseGroups := make(map[string][]FunctionInfo)
	for _, fn := range structure.Functions {
		phase := Existing_determinePhase(fn)
		phaseGroups[phase] = append(phaseGroups[phase], fn)
	}

	// Define phase order
	phases := []string{"Foundation", "Data Layer", "Store Layer", "Application Layer", "API Layer", "Routing Layer", "Main App"}

	phaseNum := 1
	for _, phase := range phases {
		if functions, exists := phaseGroups[phase]; exists {
			content += fmt.Sprintf("    subgraph Phase%d[\"🚀 PHASE %d: %s\"]\n", phaseNum, phaseNum, phase)
			for i, fn := range functions {
				content += fmt.Sprintf("        F%d[\"%d. %s<br/>📍 %s<br/>🎯 %s\"]\n",
					phaseNum*100+i, phaseNum*100+i, fn.Name, fn.File, fn.Purpose)
			}
			content += "    end\n\n"
			phaseNum++
		}
	}

	content += "```\n"

	path := filepath.Join(outDir, "Existing_dynamic_development_sequence.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// Existing_generateProjectStatusReport creates a comprehensive status report
func Existing_generateProjectStatusReport(outDir string, structure *ProjectStructure) error {
	var content strings.Builder

	content.WriteString("# Existing Project Status Report - Auto-Generated\n\n")
	content.WriteString(fmt.Sprintf("**Generated:** %s\n\n", "2025-09-14"))

	content.WriteString("## 📊 Current Project Statistics\n\n")
	content.WriteString(fmt.Sprintf("- **Total Functions:** %d\n", len(structure.Functions)))
	content.WriteString(fmt.Sprintf("- **Total Files:** %d\n", len(structure.Files)))
	content.WriteString(fmt.Sprintf("- **Total Packages:** %d\n", len(structure.Packages)))

	content.WriteString("\n## 📁 Current Package Breakdown\n\n")
	for pkg, files := range structure.Packages {
		content.WriteString(fmt.Sprintf("- **%s:** %d files\n", pkg, len(files)))
	}

	content.WriteString("\n## 🎯 Current Development Phases\n\n")
	phaseGroups := make(map[string][]FunctionInfo)
	for _, fn := range structure.Functions {
		phase := Existing_determinePhase(fn)
		phaseGroups[phase] = append(phaseGroups[phase], fn)
	}

	for phase, functions := range phaseGroups {
		content.WriteString(fmt.Sprintf("- **%s:** %d functions\n", phase, len(functions)))
	}

	path := filepath.Join(outDir, "Existing_project_status_report.md")
	return os.WriteFile(path, []byte(content.String()), 0644)
}

// Existing_categorizeFunctions groups functions by package
func Existing_categorizeFunctions(functions []FunctionInfo) map[string][]FunctionInfo {
	groups := make(map[string][]FunctionInfo)
	for _, fn := range functions {
		groups[fn.Package] = append(groups[fn.Package], fn)
	}
	return groups
}

// Existing_determinePhase determines which development phase a function belongs to
func Existing_determinePhase(fn FunctionInfo) string {
	fileName := filepath.Base(fn.File)

	// Foundation functions
	if strings.Contains(fileName, "database") || strings.Contains(fn.Name, "Open") || strings.Contains(fn.Name, "Migrate") {
		return "Foundation"
	}

	// Data layer
	if strings.Contains(fileName, "store") && !strings.Contains(fileName, "test") {
		return "Store Layer"
	}

	// Application layer
	if strings.Contains(fileName, "app") {
		return "Application Layer"
	}

	// API layer
	if strings.Contains(fileName, "handler") {
		return "API Layer"
	}

	// Routing layer
	if strings.Contains(fileName, "routes") {
		return "Routing Layer"
	}

	// Main app
	if strings.Contains(fileName, "main") || strings.Contains(fileName, "Ex10") {
		return "Main App"
	}

	return "Data Layer"
}

// Existing_getSimplePurpose provides a simple purpose description for a function
func Existing_getSimplePurpose(fn FunctionInfo) string {
	name := strings.ToLower(fn.Name)

	if strings.Contains(name, "create") {
		return "Creates new data"
	}
	if strings.Contains(name, "get") || strings.Contains(name, "find") {
		return "Retrieves data"
	}
	if strings.Contains(name, "update") {
		return "Updates existing data"
	}
	if strings.Contains(name, "delete") {
		return "Removes data"
	}
	if strings.Contains(name, "new") {
		return "Factory function"
	}
	if strings.Contains(name, "handle") {
		return "HTTP request handler"
	}
	if strings.Contains(name, "validate") {
		return "Input validation"
	}
	if strings.Contains(name, "setup") {
		return "Configuration function"
	}
	if strings.Contains(name, "open") {
		return "Opens connections"
	}
	if strings.Contains(name, "migrate") {
		return "Database migration"
	}

	return "General function"
}

// Existing_WriteApplicationBrainDiagram creates a brain diagram based on actual discovered functions
func Existing_WriteApplicationBrainDiagram(outDir string, structure *ProjectStructure) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Brain[\"🧠 THE BRAIN OF THE APPLICATION (Current Project)\"]\n"

	// Find the main application function
	var appFunction *FunctionInfo
	for _, fn := range structure.Functions {
		if fn.Name == "NewApplication" && strings.Contains(fn.File, "app.go") {
			appFunction = &fn
			break
		}
	}

	if appFunction != nil {
		content += fmt.Sprintf("        NEWAPP[\"🎯 %s() Function<br/>📍 Location: %s:%d<br/>🏆 MOST IMPORTANT FUNCTION<br/>- Creates all dependencies<br/>- Orchestrates initialization<br/>- Manages application lifecycle<br/>- Central coordination point\"]\n",
			appFunction.Name, appFunction.File, appFunction.Line)
	} else {
		content += "        NEWAPP[\"🎯 NewApplication() Function<br/>📍 Location: internal/app/app.go<br/>🏆 MOST IMPORTANT FUNCTION<br/>- Creates all dependencies<br/>- Orchestrates initialization<br/>- Manages application lifecycle<br/>- Central coordination point\"]\n"
	}

	content += "    end\n\n"

	// Find actual stores and handlers
	content += "    subgraph CriticalDeps[\"🔴 CRITICAL DEPENDENCIES (From Current Project)\"]\n"

	// Find stores
	for _, fn := range structure.Functions {
		if strings.Contains(fn.Name, "Store") && strings.Contains(fn.File, "store") {
			content += fmt.Sprintf("        %s[\"%s<br/>📍 %s<br/>🎯 %s\"]\n",
				strings.ToUpper(strings.ReplaceAll(fn.Name, "New", "")), fn.Name, fn.File, fn.Purpose)
		}
	}

	// Find handlers
	for _, fn := range structure.Functions {
		if strings.Contains(fn.Name, "Handler") && strings.Contains(fn.File, "handler") {
			content += fmt.Sprintf("        %s[\"%s<br/>📍 %s<br/>🎯 %s\"]\n",
				strings.ToUpper(strings.ReplaceAll(fn.Name, "New", "")), fn.Name, fn.File, fn.Purpose)
		}
	}

	content += "    end\n\n"

	// Find main function
	content += "    subgraph Usage[\"📱 HOW THE BRAIN IS USED (Current Project)\"]\n"
	for _, fn := range structure.Functions {
		if fn.Name == "main" && strings.Contains(fn.File, "Ex10.go") {
			content += fmt.Sprintf("        MAIN[\"Ex10.go (main)<br/>📍 %s:%d\"]\n", fn.File, fn.Line)
			break
		}
	}

	content += "        ROUTES[\"routes.SetupRoutes()\"]\n"
	content += "        SERVER[\"HTTP Server\"]\n"
	content += "    end\n\n"

	content += "    %% Brain connections\n"
	content += "    NEWAPP -->|\"Creates and manages\"| CriticalDeps\n"
	content += "    MAIN -->|\"Uses\"| NEWAPP\n"
	content += "    ROUTES -->|\"Routes to\"| CriticalDeps\n"
	content += "    SERVER -->|\"Serves\"| ROUTES\n"
	content += "```\n"

	path := filepath.Join(outDir, "Existing_application_brain.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// Existing_WriteStoreConnectionsDiagram creates a connections diagram based on actual discovered functions
func Existing_WriteStoreConnectionsDiagram(outDir string, structure *ProjectStructure) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph External[\"External Dependencies (Current Project)\"]\n" +
		"        DB[(PostgreSQL Database)]\n" +
		"        DOCKER[Docker Container<br/>workoutDB]\n" +
		"        MIGRATIONS[Migration Files<br/>00001_users.sql<br/>00002_workouts.sql<br/>00003_workout_entries.sql]\n" +
		"    end\n\n"

	// Find actual store files
	content += "    subgraph StoreLayer[\"Store Layer (Current Project)\"]\n"
	for _, fn := range structure.Functions {
		if strings.Contains(fn.File, "store") && !strings.Contains(fn.File, "test") {
			fileName := filepath.Base(fn.File)
			content += fmt.Sprintf("        %s[\"%s<br/>🎯 %s<br/>- %s\"]\n",
				strings.ToUpper(strings.ReplaceAll(fileName, ".go", "")), fileName, fn.Purpose, fn.Name)
		}
	}
	content += "    end\n\n"

	// Find actual handler files
	content += "    subgraph APILayer[\"API Layer (Current Project)\"]\n"
	for _, fn := range structure.Functions {
		if strings.Contains(fn.File, "handler") {
			fileName := filepath.Base(fn.File)
			content += fmt.Sprintf("        %s[\"%s<br/>🎯 %s<br/>- %s\"]\n",
				strings.ToUpper(strings.ReplaceAll(fileName, ".go", "")), fileName, fn.Purpose, fn.Name)
		}
	}
	content += "    end\n\n"

	// Find main app
	content += "    subgraph MainApp[\"Main App (Current Project)\"]\n"
	for _, fn := range structure.Functions {
		if fn.Name == "main" && strings.Contains(fn.File, "Ex10.go") {
			content += fmt.Sprintf("        EX10[\"Ex10.go<br/>📍 %s:%d<br/>🎯 %s\"]\n", fn.File, fn.Line, fn.Purpose)
			break
		}
	}
	content += "    end\n\n"

	content += "    %% Critical Connections (Current Project)\n"
	content += "    StoreLayer -->|\"🔴 CRITICAL<br/>Uses database connection\"| DB\n"
	content += "    APILayer -->|\"🔴 CRITICAL<br/>Uses store interfaces\"| StoreLayer\n"
	content += "    EX10 -->|\"Creates and initializes\"| StoreLayer\n"
	content += "    EX10 -->|\"Creates and initializes\"| APILayer\n"
	content += "    DOCKER -->|\"Hosts\"| DB\n"
	content += "    MIGRATIONS -->|\"Creates tables\"| DB\n"
	content += "```\n"

	path := filepath.Join(outDir, "Existing_store_connections.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// Existing_WriteArchitectureDiagram creates an architecture diagram based on the current project structure
func Existing_WriteArchitectureDiagram(wd, outDir string) error {
	var b strings.Builder
	b.WriteString("```mermaid\n")
	b.WriteString("flowchart TD\n")

	// High-level flow: API → App → Store → DB
	b.WriteString("  Client((Client)) --> API[\"API (routes + handlers)\"]\n")
	b.WriteString("  API --> App[\"App (internal/app.Application)\"]\n")
	b.WriteString("  App --> Store[\"Store (internal/store)\"]\n")
	b.WriteString("  Store --> DB[(PostgreSQL)]\n")

	// Optional context nodes
	if fileExists(filepath.Join(wd, "docker-compose.yml")) {
		b.WriteString("  Docker[/docker-compose.yml/] --> DB\n")
	}
	if fileExists(filepath.Join(wd, "migrations")) {
		b.WriteString("  Goose[/migrations/] --> DB\n")
	}

	// Subgraphs for clarity (visible grouping only)
	b.WriteString("  subgraph API_Layer[API Layer]\n")
	b.WriteString("    API_ROUTES[\"internal/routes\"]\n")
	b.WriteString("    API_HANDLERS[\"internal/api/*\"]\n")
	b.WriteString("  end\n")
	b.WriteString("  API_ROUTES --> API_HANDLERS\n")
	b.WriteString("  API_HANDLERS --> App\n")

	b.WriteString("  subgraph App_Layer[Application Layer]\n")
	b.WriteString("    APP_STRUCT[\"app.Application\"]\n")
	b.WriteString("  end\n")
	b.WriteString("  API --> APP_STRUCT\n")
	b.WriteString("  APP_STRUCT --> Store\n")

	b.WriteString("  subgraph Store_Layer[Data Access Layer]\n")
	b.WriteString("    STORE_IFACE[\"store interfaces\"]\n")
	b.WriteString("    STORE_IMPL[\"store implementations (e.g., PG)\"]\n")
	b.WriteString("  end\n")
	b.WriteString("  Store --> STORE_IFACE --> STORE_IMPL --> DB\n")

	b.WriteString("```\n")
	path := filepath.Join(outDir, "Existing_architecture.mmd.md")
	return os.WriteFile(path, []byte(b.String()), 0644)
}

// Existing_WriteFileTreeDiagram scans common project directories and writes a simple Mermaid tree
// func Existing_WriteFileTreeDiagram(wd, outDir string) error {
// 	var b strings.Builder
// 	b.WriteString("```mermaid\n")
// 	b.WriteString("graph TD\n")
// 	b.WriteString("  ROOT[\"Project Root\"]\n")

// 	// List of top-level directories to include
// 	include := []string{"internal", "migrations", "database", "cmd"}
// 	for _, name := range include {
// 		p := filepath.Join(wd, name)
// 		if dirExists(p) {
// 			// Use node IDs instead of labels with special characters
// 			dirID := strings.ReplaceAll(name, "/", "_")
// 			b.WriteString("  ROOT --> " + dirID + "[\"" + name + "\"]\n")

// 			// Add one level of children
// 			entries, err := os.ReadDir(p)
// 			if err == nil {
// 				for _, e := range entries {
// 					label := e.Name()
// 					// Create safe node ID by replacing special characters
// 					childID := strings.ReplaceAll(name+"_"+label, "/", "_")
// 					childID = strings.ReplaceAll(childID, ".", "_")
// 					childID = strings.ReplaceAll(childID, "-", "_")
// 					b.WriteString("  " + dirID + " --> " + childID + "[\"" + label + "\"]\n")
// 				}
// 			}
// 		}
// 	}
// 	b.WriteString("```\n")
// 	path := filepath.Join(outDir, "Existing_file_tree.mmd.md")
// 	return os.WriteFile(path, []byte(b.String()), 0644)
// }

// Existing_WriteFunctionDependencyDiagram analyzes actual project functions and creates a dependency diagram
// mode: 1 = simplified (exclude BT folders), 2 = full (all functions)
func Existing_WriteFunctionDependencyDiagram(wd, outDir string, mode int) error {
	// Fix: Scan the parent directory (actual project) and focus on internal folder
	projectRoot := filepath.Dir(wd)                       // Go up one level to the actual project
	internalDir := filepath.Join(projectRoot, "internal") // Focus on internal directory

	// First try to scan the internal directory specifically
	structure, err := Existing_scanProject(internalDir)
	if err != nil {
		// If internal directory scan fails, fall back to full project scan
		structure, err = Existing_scanProject(projectRoot)
		if err != nil {
			return fmt.Errorf("failed to scan project: %w", err)
		}
	}

	// Filter functions based on mode - Focus on internal directory structure
	var filteredFunctions []FunctionInfo
	if mode == 1 {
		// Simplified mode: Focus on internal directory and exclude BT folders and testing/analysis functions
		for _, fn := range structure.Functions {
			filePath := strings.ToLower(fn.File)
			funcName := strings.ToLower(fn.Name)

			// Skip BT folders and testing/analysis functions
			if strings.Contains(filePath, "bt") ||
				strings.Contains(filePath, "test") ||
				strings.Contains(filePath, "analysis") ||
				strings.Contains(filePath, "debug") ||
				strings.Contains(filePath, "tester") ||
				strings.Contains(funcName, "test") ||
				strings.Contains(funcName, "debug") ||
				strings.Contains(funcName, "analysis") {
				continue
			}

			// Prioritize functions from internal directory structure
			if strings.Contains(filePath, "internal") ||
				strings.Contains(filePath, "store") ||
				strings.Contains(filePath, "api") ||
				strings.Contains(filePath, "app") ||
				strings.Contains(filePath, "database") ||
				strings.Contains(filePath, "tokens") ||
				strings.Contains(filePath, "middleware") {
				filteredFunctions = append(filteredFunctions, fn)
			}
		}
	} else {
		// Full mode: include all functions
		filteredFunctions = structure.Functions
	}

	var b strings.Builder
	b.WriteString("```mermaid\n")
	b.WriteString("flowchart TB\n")
	b.WriteString("    %% Generated from actual project analysis - VERTICAL LAYOUT\n")
	if mode == 1 {
		b.WriteString("    %% SIMPLIFIED MODE - Core functions only (excludes BT folders and testing)\n")
	} else {
		b.WriteString("    %% FULL MODE - All functions in project\n")
	}
	b.WriteString("    %% Total functions found: " + fmt.Sprintf("%d", len(structure.Functions)) + "\n")
	b.WriteString("    %% Functions included: " + fmt.Sprintf("%d", len(filteredFunctions)) + "\n\n")

	// Add FIXED high-resolution styling and configuration for better HTML visibility
	b.WriteString("    %% FIXED High-resolution configuration for HTML visibility\n")
	b.WriteString("    classDef mainClass fill:#ffebee,stroke:#d32f2f,stroke-width:4px,color:#000,font-size:16px,font-weight:bold\n")
	b.WriteString("    classDef databaseClass fill:#e3f2fd,stroke:#0277bd,stroke-width:3px,color:#000,font-size:14px,font-weight:bold\n")
	b.WriteString("    classDef storeClass fill:#f3e5f5,stroke:#7b1fa2,stroke-width:3px,color:#000,font-size:14px,font-weight:bold\n")
	b.WriteString("    classDef tokenClass fill:#fff3e0,stroke:#f57c00,stroke-width:3px,color:#000,font-size:14px,font-weight:bold\n")
	b.WriteString("    classDef middlewareClass fill:#fff8e1,stroke:#f57c00,stroke-width:3px,color:#000,font-size:14px,font-weight:bold\n")
	b.WriteString("    classDef apiClass fill:#fce4ec,stroke:#c2185b,stroke-width:3px,color:#000,font-size:14px,font-weight:bold\n")
	b.WriteString("    classDef appClass fill:#e8f5e8,stroke:#388e3c,stroke-width:3px,color:#000,font-size:14px,font-weight:bold\n")
	b.WriteString("    classDef otherClass fill:#fafafa,stroke:#616161,stroke-width:3px,color:#000,font-size:14px,font-weight:bold\n")
	b.WriteString("\n")

	// Group functions by internal directory structure
	appFuncs := []FunctionInfo{}
	storeFuncs := []FunctionInfo{}
	apiFuncs := []FunctionInfo{}
	databaseFuncs := []FunctionInfo{}
	tokenFuncs := []FunctionInfo{}
	middlewareFuncs := []FunctionInfo{}
	mainFuncs := []FunctionInfo{}
	otherFuncs := []FunctionInfo{}

	for _, fn := range filteredFunctions {
		filePath := strings.ToLower(fn.File)
		funcName := strings.ToLower(fn.Name)

		// Categorize by internal directory structure
		if strings.Contains(filePath, "internal/app") || strings.Contains(filePath, "app") ||
			strings.Contains(funcName, "newapplication") || strings.Contains(funcName, "application") {
			appFuncs = append(appFuncs, fn)
		} else if strings.Contains(filePath, "internal/store") || strings.Contains(filePath, "store") ||
			strings.Contains(funcName, "store") || strings.Contains(funcName, "create") ||
			strings.Contains(funcName, "get") || strings.Contains(funcName, "update") ||
			strings.Contains(funcName, "delete") {
			storeFuncs = append(storeFuncs, fn)
		} else if strings.Contains(filePath, "internal/api") || strings.Contains(filePath, "api") ||
			strings.Contains(funcName, "handle") || strings.Contains(funcName, "handler") {
			apiFuncs = append(apiFuncs, fn)
		} else if strings.Contains(filePath, "internal/database") || strings.Contains(filePath, "database") ||
			strings.Contains(funcName, "open") || strings.Contains(funcName, "migrate") ||
			strings.Contains(funcName, "database") {
			databaseFuncs = append(databaseFuncs, fn)
		} else if strings.Contains(filePath, "internal/tokens") || strings.Contains(filePath, "tokens") ||
			strings.Contains(funcName, "token") || strings.Contains(funcName, "jwt") {
			tokenFuncs = append(tokenFuncs, fn)
		} else if strings.Contains(filePath, "internal/middleware") || strings.Contains(filePath, "middleware") ||
			strings.Contains(funcName, "middleware") || strings.Contains(funcName, "auth") ||
			strings.Contains(funcName, "validate") {
			middlewareFuncs = append(middlewareFuncs, fn)
		} else if strings.Contains(filePath, "main") || funcName == "main" {
			mainFuncs = append(mainFuncs, fn)
		} else {
			otherFuncs = append(otherFuncs, fn)
		}
	}

	// Write main functions first
	if len(mainFuncs) > 0 {
		b.WriteString("    subgraph MainApp[\"🚀 MAIN APPLICATION (Entry Point)\"]\n")
		for _, fn := range mainFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Write database functions
	if len(databaseFuncs) > 0 {
		b.WriteString("    subgraph Database[\"🗄️ DATABASE LAYER (internal/database)\"]\n")
		for _, fn := range databaseFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Write store functions
	if len(storeFuncs) > 0 {
		b.WriteString("    subgraph Store[\"💾 STORE LAYER (internal/store)\"]\n")
		for _, fn := range storeFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Write token functions
	if len(tokenFuncs) > 0 {
		b.WriteString("    subgraph Tokens[\"🔑 TOKEN LAYER (internal/tokens)\"]\n")
		for _, fn := range tokenFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Write middleware functions
	if len(middlewareFuncs) > 0 {
		b.WriteString("    subgraph Middleware[\"🛡️ MIDDLEWARE LAYER (internal/middleware)\"]\n")
		for _, fn := range middlewareFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Write API functions
	if len(apiFuncs) > 0 {
		b.WriteString("    subgraph API[\"🌐 API LAYER (internal/api)\"]\n")
		for _, fn := range apiFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Write app functions
	if len(appFuncs) > 0 {
		b.WriteString("    subgraph App[\"🏗️ APPLICATION LAYER (internal/app)\"]\n")
		for _, fn := range appFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Write other functions
	if len(otherFuncs) > 0 {
		b.WriteString("    subgraph Other[\"📦 OTHER FUNCTIONS\"]\n")
		for _, fn := range otherFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Write main functions
	if len(mainFuncs) > 0 {
		b.WriteString("    subgraph MainApp[\"🚀 MAIN FUNCTIONS (Build Last)\"]\n")
		for _, fn := range mainFuncs {
			nodeID := strings.ReplaceAll(fn.Name, ".", "_")
			nodeID = strings.ReplaceAll(nodeID, "-", "_")
			shortPurpose := fn.Purpose
			if len(shortPurpose) > 35 {
				shortPurpose = shortPurpose[:32] + "..."
			}
			b.WriteString(fmt.Sprintf("        %s[\"%s()<br/>📁 %s<br/>%s\"]\n",
				nodeID, fn.Name, filepath.Base(fn.File), shortPurpose))
		}
		b.WriteString("    end\n\n")
	}

	// Add comprehensive dependency relationships based on actual project analysis
	b.WriteString("    %% ENHANCED dependency patterns (based on actual project analysis)\n")

	// Create a map of function names to node IDs for easier lookup
	funcMap := make(map[string]string)
	for _, fn := range filteredFunctions {
		nodeID := strings.ReplaceAll(fn.Name, ".", "_")
		nodeID = strings.ReplaceAll(nodeID, "-", "_")
		funcMap[strings.ToLower(fn.Name)] = nodeID
	}

	// Enhanced dependency analysis based on actual project structure
	for _, fn := range filteredFunctions {
		funcName := strings.ToLower(fn.Name)
		fileName := strings.ToLower(filepath.Base(fn.File))
		nodeID := funcMap[funcName]

		// 1. Main function dependencies
		if funcName == "main" {
			// Main typically calls NewApplication
			if newNodeID, exists := funcMap["newapplication"]; exists {
				b.WriteString(fmt.Sprintf("    %s --> %s\n", nodeID, newNodeID))
			}
		}

		// 2. Application constructor dependencies
		if strings.Contains(funcName, "newapplication") {
			// NewApplication typically creates stores and handlers
			for _, otherFn := range filteredFunctions {
				otherName := strings.ToLower(otherFn.Name)
				if (strings.Contains(otherName, "new") && strings.Contains(otherName, "store")) ||
					(strings.Contains(otherName, "new") && strings.Contains(otherName, "handler")) {
					otherNodeID := funcMap[otherName]
					b.WriteString(fmt.Sprintf("    %s --> %s\n", nodeID, otherNodeID))
				}
			}
		}

		// 3. Handler dependencies on stores
		if strings.Contains(funcName, "handler") && strings.Contains(funcName, "new") {
			// Extract the resource type (e.g., "workout" from "NewWorkoutHandler")
			resourceType := strings.Replace(funcName, "new", "", 1)
			resourceType = strings.Replace(resourceType, "handler", "", 1)

			// Find corresponding store
			for _, otherFn := range filteredFunctions {
				otherName := strings.ToLower(otherFn.Name)
				if strings.Contains(otherName, "new") && strings.Contains(otherName, "store") &&
					strings.Contains(otherName, resourceType) {
					otherNodeID := funcMap[otherName]
					b.WriteString(fmt.Sprintf("    %s --> %s\n", otherNodeID, nodeID))
				}
			}
		}

		// 4. Store dependencies on database
		if strings.Contains(funcName, "store") && strings.Contains(funcName, "new") {
			// Store constructors depend on database connection
			for _, otherFn := range filteredFunctions {
				otherName := strings.ToLower(otherFn.Name)
				if strings.Contains(otherName, "open") || strings.Contains(otherName, "connect") ||
					strings.Contains(otherName, "database") {
					otherNodeID := funcMap[otherName]
					b.WriteString(fmt.Sprintf("    %s --> %s\n", otherNodeID, nodeID))
				}
			}
		}

		// 5. Database connection dependencies
		if strings.Contains(funcName, "open") && strings.Contains(fileName, "database") {
			// Database connection typically depends on migration
			for _, otherFn := range filteredFunctions {
				otherName := strings.ToLower(otherFn.Name)
				if strings.Contains(otherName, "migrate") {
					otherNodeID := funcMap[otherName]
					b.WriteString(fmt.Sprintf("    %s --> %s\n", otherNodeID, nodeID))
				}
			}
		}

		// 6. Route setup dependencies
		if strings.Contains(funcName, "setup") && strings.Contains(funcName, "route") {
			// Route setup depends on handlers
			for _, otherFn := range filteredFunctions {
				otherName := strings.ToLower(otherFn.Name)
				if strings.Contains(otherName, "handler") && strings.Contains(otherName, "new") {
					otherNodeID := funcMap[otherName]
					b.WriteString(fmt.Sprintf("    %s --> %s\n", otherNodeID, nodeID))
				}
			}
		}

		// 7. Middleware dependencies
		if strings.Contains(funcName, "middleware") {
			// Middleware typically depends on authentication stores
			for _, otherFn := range filteredFunctions {
				otherName := strings.ToLower(otherFn.Name)
				if strings.Contains(otherName, "token") || strings.Contains(otherName, "user") {
					otherNodeID := funcMap[otherName]
					b.WriteString(fmt.Sprintf("    %s --> %s\n", otherNodeID, nodeID))
				}
			}
		}
	}

	// Apply CSS classes to function nodes for better styling
	b.WriteString("    %% Apply styling classes\n")
	for _, fn := range filteredFunctions {
		nodeID := strings.ReplaceAll(fn.Name, ".", "_")
		nodeID = strings.ReplaceAll(nodeID, "-", "_")
		filePath := strings.ToLower(fn.File)
		funcName := strings.ToLower(fn.Name)

		// Determine class based on internal directory structure
		var className string
		if strings.Contains(filePath, "main") || funcName == "main" {
			className = "mainClass"
		} else if strings.Contains(filePath, "internal/database") || strings.Contains(filePath, "database") {
			className = "databaseClass"
		} else if strings.Contains(filePath, "internal/store") || strings.Contains(filePath, "store") {
			className = "storeClass"
		} else if strings.Contains(filePath, "internal/tokens") || strings.Contains(filePath, "tokens") {
			className = "tokenClass"
		} else if strings.Contains(filePath, "internal/middleware") || strings.Contains(filePath, "middleware") {
			className = "middlewareClass"
		} else if strings.Contains(filePath, "internal/api") || strings.Contains(filePath, "api") {
			className = "apiClass"
		} else if strings.Contains(filePath, "internal/app") || strings.Contains(filePath, "app") {
			className = "appClass"
		} else {
			className = "otherClass"
		}

		b.WriteString(fmt.Sprintf("    class %s %s\n", nodeID, className))
	}

	b.WriteString("```\n")

	// Write to file
	var filename string
	if mode == 1 {
		filename = "Existing_function_dependencies_simplified.mmd.md"
	} else {
		filename = "Existing_function_dependencies_full.mmd.md"
	}
	path := filepath.Join(outDir, filename)
	return os.WriteFile(path, []byte(b.String()), 0644)
}

// Helper functions for file/directory existence checks are defined in BTProjectDiagrams.go
