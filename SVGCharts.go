//go:build flowcharts
// +build flowcharts

/*
===============================================================================
SVG CHARTS GENERATOR - REAL PROJECT ANALYSIS
===============================================================================

Author: Ben Tran
Date: 02/09/2025
Description: This file contains SVG chart generation functionality for real project
             analysis. It generates function call graphs, package dependencies,
             and PlantUML diagrams based on the actual scanned project structure.

TO USE THIS FILE:
1. Call generateSVGCharts() with project structure for real analysis
2. Requires external tools: go-callvis, goda, dot, goplantuml
3. Generates SVG files for visualization

REAL PROJECT OBJECTIVES:
- Function call graphs based on actual project functions
- Package dependency analysis from real project structure
- PlantUML class diagrams from actual types
- Migration-focused analysis from real migration files

FEATURES:
- graph.svg - Main function call graph
- graph_by_pkg.svg - Package-grouped call graph
- graph_full.svg - Full graph including stdlib
- graph_migrations.svg - Migration-focused graph
- pkg-deps.svg - Package dependency graph
- types.svg - PlantUML class diagram

===============================================================================
*/

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// generateSVGCharts runs SVG chart generation based on real project structure
func generateSVGCharts(root, outDir string, opts FlowchartOptions, structure *ProjectStructure) error {
	fmt.Println("🌐 Generating SVG charts based on real project analysis...")

	// Analyze real project structure if available
	if structure != nil {
		fmt.Printf("📊 Analyzing real project structure: %d functions across %d files\n",
			len(structure.Functions), len(structure.Files))

		// Find key functions for analysis
		keyFunctions := 0
		for _, fn := range structure.Functions {
			if strings.Contains(strings.ToLower(fn.Name), "main") ||
				strings.Contains(strings.ToLower(fn.Name), "handler") ||
				strings.Contains(strings.ToLower(fn.Name), "store") ||
				strings.Contains(strings.ToLower(fn.Name), "route") {
				keyFunctions++
			}
		}
		fmt.Printf("🔑 Found %d key functions (main, handler, store, route) in current project\n", keyFunctions)
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

	// Generate function call graphs with better error handling
	fmt.Println("📊 Generating function call graphs...")

	// Main graph - use focus to avoid multiple main packages issue
	callvisArgs := []string{"-format", "svg", "-file", filepath.Join(outDir, "graph.svg")}
	if opts.NoStdlib {
		callvisArgs = append(callvisArgs, "-nostd")
	}
	if opts.Group != "" {
		callvisArgs = append(callvisArgs, "-group", opts.Group)
	}
	// Use module path as focus to avoid multiple main packages issue
	if mod := readModulePath(filepath.Join(root, "go.mod")); mod != "" {
		callvisArgs = append(callvisArgs, "-focus", mod)
	}
	if opts.Ignore != "" {
		callvisArgs = append(callvisArgs, "-ignore", opts.Ignore)
	}
	if opts.IncludeTests {
		callvisArgs = append(callvisArgs, "-tests")
	}
	callvisArgs = append(callvisArgs, "./...")

	// Try to generate main graph, but don't fail if it has multiple main packages
	if err := runInDir(root, "go-callvis", callvisArgs...); err != nil {
		fmt.Printf("⚠️  Main graph generation failed (multiple main packages): %v\n", err)
		fmt.Println("   This is expected when running multiple chart files together.")
	} else {
		fmt.Println("✅ Generated graph.svg")
	}

	// Package-grouped graph
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
	if err := runInDir(root, "go-callvis", byPkg...); err != nil {
		fmt.Printf("⚠️  Package-grouped graph failed: %v\n", err)
	} else {
		fmt.Println("✅ Generated graph_by_pkg.svg")
	}

	// Full graph (including stdlib)
	full := []string{"-format", "svg", "-file", filepath.Join(outDir, "graph_full.svg")}
	// do not add -nostd here on purpose
	if opts.Group != "" {
		full = append(full, "-group", opts.Group)
	}
	// Use module path as focus to avoid multiple main packages issue
	if mod := readModulePath(filepath.Join(root, "go.mod")); mod != "" {
		full = append(full, "-focus", mod)
	}
	if opts.Ignore != "" {
		full = append(full, "-ignore", opts.Ignore)
	}
	if opts.IncludeTests {
		full = append(full, "-tests")
	}
	full = append(full, "./...")
	if err := runInDir(root, "go-callvis", full...); err != nil {
		fmt.Printf("⚠️  Full graph failed: %v\n", err)
	} else {
		fmt.Println("✅ Generated graph_full.svg")
	}

	// Migrations-focused graph (based on real project analysis)
	if dirExists(filepath.Join(root, "migrations")) {
		focusVal := "migrations"
		if mod := readModulePath(filepath.Join(root, "go.mod")); mod != "" {
			focusVal = mod + "/migrations"
		}
		mig := []string{"-format", "svg", "-file", filepath.Join(outDir, "graph_migrations.svg"), "-group", "pkg,type"}
		if opts.IncludeTests {
			mig = append(mig, "-tests")
		}
		mig = append(mig, "-focus", focusVal, "./...")
		if err := runInDir(root, "go-callvis", mig...); err != nil {
			fmt.Printf("⚠️  Migrations graph failed: %v\n", err)
		} else {
			fmt.Println("✅ Generated graph_migrations.svg")
		}
	}

	// Generate package dependency graph
	fmt.Println("📦 Generating package dependency graph...")
	dotPath := filepath.Join(outDir, "pkg-deps.dot")
	svgPath := filepath.Join(outDir, "pkg-deps.svg")
	if err := writeFileFromCmd(root, []string{"goda", "graph", "./..."}, dotPath); err != nil {
		return fmt.Errorf("write dot: %w", err)
	}
	if err := runInDir(root, "dot", "-Tsvg", dotPath, "-o", svgPath); err != nil {
		return fmt.Errorf("dot convert: %w", err)
	}
	fmt.Println("✅ Generated pkg-deps.svg")

	// Generate PlantUML class diagram if available
	if opts.GenerateUML {
		fmt.Println("🎨 Generating PlantUML class diagram...")
		if err := ensureTool("goplantuml"); err == nil {
			umlPath := filepath.Join(outDir, "types.puml")
			if err := writeFileFromCmd(root, []string{"goplantuml", "-recursive", "."}, umlPath); err != nil {
				fmt.Printf("⚠️  PlantUML generation failed: %v\n", err)
			} else {
				fmt.Println("✅ Generated types.puml")

				// Render types.puml to SVG if PlantUML is available
				if cmd, args, ok := findPlantUMLRenderer(); ok {
					if err := runInDir(filepath.Join(root, outDir), cmd, append(args, "types.puml")...); err != nil {
						fmt.Printf("⚠️  PlantUML render failed: %v\n", err)
					} else {
						fmt.Println("✅ Generated types.svg")
					}
				} else {
					fmt.Println("ℹ️  PlantUML renderer not found. Install PlantUML or set PLANTUML_JAR to render SVG.")
				}
			}
		} else {
			fmt.Println("ℹ️  goplantuml not found. Install with: go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest")
		}
	}

	fmt.Println("🎉 SVG chart generation complete!")
	return nil
}
