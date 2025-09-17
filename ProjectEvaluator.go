//go:build flowcharts
// +build flowcharts

/*
===============================================================================
PROJECT EVALUATOR - COMPREHENSIVE PROJECT STATUS ASSESSMENT SYSTEM
===============================================================================

Author: Professor Model Builder (AI Teaching Assistant)
Date: 17/09/2025
Description: This file contains comprehensive project evaluation functions that
analyze the current state of a Go API project and provide detailed feedback
on progress, quality, structure, and next steps.

FEATURES:
- Current project status analysis
- Progress tracking with completion percentages
- Quality assessment with sub-scores
- Structure and architecture evaluation
- Error and mistake detection
- Intelligent advice system
- Comprehensive evaluation reports
- Final scoring and ratings

TO USE THIS FILE:
1. Call ProjectEvaluator_WriteComprehensiveAssessment() for full evaluation
2. Individual evaluation functions can be called for specific aspects
3. Reports are saved as ProjectEvaluator_*.mmd.md files

===============================================================================
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ProjectStatus represents the current status of the project
type ProjectStatus struct {
	CurrentPhase      string
	CompletionPercent int
	NextStep          string
	QualityScore      int
	StructureScore    int
	ErrorCount        int
	WarningCount      int
	AdviceList        []string
	SubScores         map[string]int
	FinalScore        int
	Rating            string
}

// ProjectEvaluator_WriteComprehensiveAssessment creates a complete project evaluation
func ProjectEvaluator_WriteComprehensiveAssessment(outDir string) error {
	fmt.Println("🔍 Starting Comprehensive Project Evaluation...")

	// Analyze current project status
	status := ProjectEvaluator_AnalyzeProjectStatus()

	// Generate comprehensive assessment report
	content := ProjectEvaluator_GenerateAssessmentReport(status)

	path := filepath.Join(outDir, "ProjectEvaluator_comprehensive_assessment.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// ProjectEvaluator_AnalyzeProjectStatus analyzes the current project state
func ProjectEvaluator_AnalyzeProjectStatus() ProjectStatus {
	status := ProjectStatus{
		SubScores:  make(map[string]int),
		AdviceList: []string{},
	}

	// Find the actual project directory (not the diagram generator)
	projectRoot := ProjectEvaluator_FindProjectRoot()

	// Analyze project structure
	status.StructureScore = ProjectEvaluator_AnalyzeStructure(projectRoot)

	// Analyze code quality
	status.QualityScore = ProjectEvaluator_AnalyzeCodeQuality(projectRoot)

	// Determine current phase and progress
	status.CurrentPhase, status.CompletionPercent = ProjectEvaluator_DetermineProgress(projectRoot)

	// Identify next steps
	status.NextStep = ProjectEvaluator_IdentifyNextStep(status.CurrentPhase)

	// Count errors and warnings
	status.ErrorCount, status.WarningCount = ProjectEvaluator_CountIssues(projectRoot)

	// Generate advice
	status.AdviceList = ProjectEvaluator_GenerateAdvice(status)

	// Calculate sub-scores
	status.SubScores = ProjectEvaluator_CalculateSubScores(status, projectRoot)

	// Calculate final score and rating
	status.FinalScore, status.Rating = ProjectEvaluator_CalculateFinalScore(status)

	return status
}

// ProjectEvaluator_FindProjectRoot finds the actual Go project directory
func ProjectEvaluator_FindProjectRoot() string {
	// Start from current directory and look for go.mod
	currentDir, _ := os.Getwd()

	// Check current directory first
	if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
		return currentDir
	}

	// Look in parent directories
	parentDir := filepath.Dir(currentDir)
	if _, err := os.Stat(filepath.Join(parentDir, "go.mod")); err == nil {
		return parentDir
	}

	// Look in grandparent directory
	grandparentDir := filepath.Dir(parentDir)
	if _, err := os.Stat(filepath.Join(grandparentDir, "go.mod")); err == nil {
		return grandparentDir
	}

	// If no go.mod found, assume current directory but look for Go files
	if files, err := filepath.Glob("*.go"); err == nil && len(files) > 0 {
		return currentDir
	}

	// Default to current directory
	return currentDir
}

// ProjectEvaluator_AnalyzeStructure analyzes project structure and organization
func ProjectEvaluator_AnalyzeStructure(projectRoot string) int {
	score := 0

	// Check for essential directories
	essentialDirs := []string{"internal", "internal/app", "internal/api", "internal/store", "internal/database", "internal/middleware"}
	for _, dir := range essentialDirs {
		if _, err := os.Stat(filepath.Join(projectRoot, dir)); err == nil {
			score += 15
		}
	}

	// Check for essential files
	essentialFiles := []string{"main.go", "go.mod", "docker-compose.yml"}
	for _, file := range essentialFiles {
		if _, err := os.Stat(filepath.Join(projectRoot, file)); err == nil {
			score += 10
		}
	}

	// Check for test files
	if files, err := filepath.Glob(filepath.Join(projectRoot, "*_test.go")); err == nil && len(files) > 0 {
		score += 10
	}

	return min(score, 100)
}

// ProjectEvaluator_AnalyzeCodeQuality analyzes code quality and best practices
func ProjectEvaluator_AnalyzeCodeQuality(projectRoot string) int {
	score := 0

	// Check for proper error handling
	if ProjectEvaluator_HasErrorHandling(projectRoot) {
		score += 25
	}

	// Check for proper logging
	if ProjectEvaluator_HasLogging(projectRoot) {
		score += 20
	}

	// Check for proper documentation
	if ProjectEvaluator_HasDocumentation(projectRoot) {
		score += 15
	}

	// Check for proper testing
	if ProjectEvaluator_HasTesting(projectRoot) {
		score += 20
	}

	// Check for proper configuration
	if ProjectEvaluator_HasConfiguration(projectRoot) {
		score += 20
	}

	return min(score, 100)
}

// ProjectEvaluator_DetermineProgress determines current phase and completion percentage
func ProjectEvaluator_DetermineProgress(projectRoot string) (string, int) {
	phase := "Project Initialization"
	completion := 0

	// Check project foundation
	if _, err := os.Stat(filepath.Join(projectRoot, "main.go")); err == nil {
		completion += 10
		phase = "Application Foundation"
	}

	if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
		completion += 5
	}

	// Check internal structure
	if _, err := os.Stat(filepath.Join(projectRoot, "internal")); err == nil {
		completion += 10
		phase = "Internal Structure"
	}

	// Check application layer
	if _, err := os.Stat(filepath.Join(projectRoot, "internal/app")); err == nil {
		completion += 10
		phase = "Application Layer"
	}

	// Check API layer
	if _, err := os.Stat(filepath.Join(projectRoot, "internal/api")); err == nil {
		completion += 15
		phase = "API Layer"
	}

	// Check database layer
	if _, err := os.Stat(filepath.Join(projectRoot, "internal/database")); err == nil {
		completion += 15
		phase = "Database Layer"
	}

	// Check store layer
	if _, err := os.Stat(filepath.Join(projectRoot, "internal/store")); err == nil {
		completion += 15
		phase = "Store Layer"
	}

	// Check middleware
	if _, err := os.Stat(filepath.Join(projectRoot, "internal/middleware")); err == nil {
		completion += 10
		phase = "Authentication & Middleware"
	}

	// Check testing
	if files, err := filepath.Glob(filepath.Join(projectRoot, "*_test.go")); err == nil && len(files) > 0 {
		completion += 10
		phase = "Testing & Deployment"
	}

	return phase, min(completion, 100)
}

// ProjectEvaluator_IdentifyNextStep identifies the next logical step
func ProjectEvaluator_IdentifyNextStep(currentPhase string) string {
	switch currentPhase {
	case "Project Initialization":
		return "Create main.go and initialize Go module"
	case "Application Foundation":
		return "Create internal directory structure"
	case "Internal Structure":
		return "Implement application layer with app.go"
	case "Application Layer":
		return "Create API handlers in internal/api"
	case "API Layer":
		return "Set up database connection and migrations"
	case "Database Layer":
		return "Implement store layer for data access"
	case "Store Layer":
		return "Add authentication and middleware"
	case "Authentication & Middleware":
		return "Write comprehensive tests"
	case "Testing & Deployment":
		return "Optimize and deploy the application"
	default:
		return "Review and refactor existing code"
	}
}

// ProjectEvaluator_CountIssues counts errors and warnings in the project
func ProjectEvaluator_CountIssues(projectRoot string) (int, int) {
	errors := 0
	warnings := 0

	// Check for common issues
	if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err != nil {
		errors++
	}

	if _, err := os.Stat(filepath.Join(projectRoot, "main.go")); err != nil {
		errors++
	}

	// Check for missing essential directories
	essentialDirs := []string{"internal", "internal/app", "internal/api"}
	for _, dir := range essentialDirs {
		if _, err := os.Stat(filepath.Join(projectRoot, dir)); err != nil {
			warnings++
		}
	}

	// Check for missing test files
	if files, err := filepath.Glob(filepath.Join(projectRoot, "*_test.go")); err != nil || len(files) == 0 {
		warnings++
	}

	return errors, warnings
}

// ProjectEvaluator_GenerateAdvice generates intelligent advice based on current status
func ProjectEvaluator_GenerateAdvice(status ProjectStatus) []string {
	advice := []string{}

	// Structure advice
	if status.StructureScore < 70 {
		advice = append(advice, "🔧 Improve project structure by creating missing directories")
	}

	// Quality advice
	if status.QualityScore < 70 {
		advice = append(advice, "📝 Add proper error handling and logging")
	}

	// Progress advice
	if status.CompletionPercent < 50 {
		advice = append(advice, "🚀 Focus on completing the core application layers first")
	} else if status.CompletionPercent < 80 {
		advice = append(advice, "🔐 Implement authentication and middleware")
	} else {
		advice = append(advice, "🧪 Add comprehensive testing and documentation")
	}

	// Error advice
	if status.ErrorCount > 0 {
		advice = append(advice, "❌ Fix critical errors before proceeding")
	}

	if status.WarningCount > 2 {
		advice = append(advice, "⚠️ Address warnings to improve code quality")
	}

	return advice
}

// ProjectEvaluator_CalculateSubScores calculates detailed sub-scores
func ProjectEvaluator_CalculateSubScores(status ProjectStatus, projectRoot string) map[string]int {
	scores := make(map[string]int)

	scores["Structure"] = status.StructureScore
	scores["Code Quality"] = status.QualityScore
	scores["Progress"] = status.CompletionPercent
	scores["Error Handling"] = ProjectEvaluator_ScoreErrorHandling(projectRoot)
	scores["Testing"] = ProjectEvaluator_ScoreTesting(projectRoot)
	scores["Documentation"] = ProjectEvaluator_ScoreDocumentation(projectRoot)
	scores["Configuration"] = ProjectEvaluator_ScoreConfiguration(projectRoot)

	return scores
}

// ProjectEvaluator_CalculateFinalScore calculates the final overall score
func ProjectEvaluator_CalculateFinalScore(status ProjectStatus) (int, string) {
	// More generous weighted average of sub-scores
	totalScore := 0
	weights := map[string]int{
		"Structure":      25,
		"Code Quality":   30,
		"Progress":       25,
		"Error Handling": 10,
		"Testing":        5,
		"Documentation":  3,
		"Configuration":  2,
	}

	for category, score := range status.SubScores {
		if weight, exists := weights[category]; exists {
			totalScore += (score * weight) / 100
		}
	}

	// More lenient penalty system
	totalScore -= status.ErrorCount * 5   // Reduced from 10
	totalScore -= status.WarningCount * 1 // Reduced from 2

	// Add bonus for having Go files (even if not perfect structure)
	if status.QualityScore > 0 {
		totalScore += 10 // Bonus for having some code
	}

	totalScore = max(0, min(totalScore, 100))

	// More generous rating system
	var rating string
	switch {
	case totalScore >= 85:
		rating = "🌟 EXCELLENT"
	case totalScore >= 75:
		rating = "⭐ VERY GOOD"
	case totalScore >= 65:
		rating = "👍 GOOD"
	case totalScore >= 50:
		rating = "📈 FAIR"
	case totalScore >= 30:
		rating = "⚠️ NEEDS IMPROVEMENT"
	default:
		rating = "🚨 REQUIRES ATTENTION"
	}

	return totalScore, rating
}

// Helper functions for detailed analysis
func ProjectEvaluator_HasErrorHandling(projectRoot string) bool {
	// Check for error handling patterns in Go files
	if files, err := filepath.Glob(filepath.Join(projectRoot, "*.go")); err == nil {
		for _, file := range files {
			if content, err := os.ReadFile(file); err == nil {
				if strings.Contains(string(content), "if err != nil") {
					return true
				}
			}
		}
	}
	return false
}

func ProjectEvaluator_HasLogging(projectRoot string) bool {
	// Check for logging patterns
	if files, err := filepath.Glob(filepath.Join(projectRoot, "*.go")); err == nil {
		for _, file := range files {
			if content, err := os.ReadFile(file); err == nil {
				if strings.Contains(string(content), "log.") || strings.Contains(string(content), "Logger") {
					return true
				}
			}
		}
	}
	return false
}

func ProjectEvaluator_HasDocumentation(projectRoot string) bool {
	// Check for documentation
	if files, err := filepath.Glob(filepath.Join(projectRoot, "*.go")); err == nil {
		for _, file := range files {
			if content, err := os.ReadFile(file); err == nil {
				if strings.Contains(string(content), "//") && len(strings.Split(string(content), "//")) > 5 {
					return true
				}
			}
		}
	}
	return false
}

func ProjectEvaluator_HasTesting(projectRoot string) bool {
	// Check for test files
	if files, err := filepath.Glob(filepath.Join(projectRoot, "*_test.go")); err == nil {
		return len(files) > 0
	}
	return false
}

func ProjectEvaluator_HasConfiguration(projectRoot string) bool {
	// Check for configuration files
	configFiles := []string{"docker-compose.yml", ".env", "config.json", "config.yaml"}
	for _, file := range configFiles {
		if _, err := os.Stat(filepath.Join(projectRoot, file)); err == nil {
			return true
		}
	}
	return false
}

func ProjectEvaluator_ScoreErrorHandling(projectRoot string) int {
	if ProjectEvaluator_HasErrorHandling(projectRoot) {
		return 80
	}
	return 20
}

func ProjectEvaluator_ScoreTesting(projectRoot string) int {
	if ProjectEvaluator_HasTesting(projectRoot) {
		return 85
	}
	return 15
}

func ProjectEvaluator_ScoreDocumentation(projectRoot string) int {
	if ProjectEvaluator_HasDocumentation(projectRoot) {
		return 70
	}
	return 30
}

func ProjectEvaluator_ScoreConfiguration(projectRoot string) int {
	if ProjectEvaluator_HasConfiguration(projectRoot) {
		return 75
	}
	return 25
}

// ProjectEvaluator_GenerateAssessmentReport generates the comprehensive assessment report
func ProjectEvaluator_GenerateAssessmentReport(status ProjectStatus) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Assessment[\"🔍 COMPREHENSIVE PROJECT EVALUATION REPORT\"]\n" +
		"        %% Header\n" +
		"        subgraph Header[\"📊 PROJECT STATUS OVERVIEW\"]\n" +
		fmt.Sprintf("            H1[\"📅 Evaluation Date: %s<br/>🎯 Current Phase: %s<br/>📈 Completion: %d%%<br/>🏆 Final Score: %d/100<br/>⭐ Rating: %s\"]\n", timestamp, status.CurrentPhase, status.CompletionPercent, status.FinalScore, status.Rating) +
		"        end\n\n" +

		"        %% Progress Analysis\n" +
		"        subgraph Progress[\"📈 PROGRESS ANALYSIS\"]\n" +
		fmt.Sprintf("            P1[\"🎯 Current Phase: %s<br/>📊 Completion Percentage: %d%%<br/>⏭️ Next Step: %s<br/>🔍 Status: %s\"]\n", status.CurrentPhase, status.CompletionPercent, status.NextStep, status.Rating) +
		"        end\n\n" +

		"        %% Quality Assessment\n" +
		"        subgraph Quality[\"🏆 QUALITY ASSESSMENT\"]\n" +
		fmt.Sprintf("            Q1[\"🏗️ Structure Score: %d/100<br/>📝 Code Quality: %d/100<br/>❌ Errors: %d<br/>⚠️ Warnings: %d\"]\n", status.StructureScore, status.QualityScore, status.ErrorCount, status.WarningCount) +
		"        end\n\n" +

		"        %% Detailed Sub-Scores\n" +
		"        subgraph SubScores[\"📊 DETAILED SUB-SCORES\"]\n"

	// Add sub-scores
	for category, score := range status.SubScores {
		content += fmt.Sprintf("            S%d[\"%s: %d/100\"]\n", len(status.SubScores), category, score)
	}

	content += "        end\n\n" +

		"        %% Advice Section\n" +
		"        subgraph Advice[\"💡 INTELLIGENT ADVICE & RECOMMENDATIONS\"]\n"

	// Add advice items
	for i, advice := range status.AdviceList {
		content += fmt.Sprintf("            A%d[\"%s\"]\n", i+1, advice)
	}

	content += "        end\n\n" +

		"        %% Final Assessment\n" +
		"        subgraph Final[\"🎯 FINAL ASSESSMENT\"]\n" +
		fmt.Sprintf("            F1[\"🏆 Overall Score: %d/100<br/>⭐ Rating: %s<br/>📈 Progress: %d%% Complete<br/>🎯 Focus: %s\"]\n", status.FinalScore, status.Rating, status.CompletionPercent, status.NextStep) +
		"        end\n" +
		"    end\n\n" +

		"    %% Connections\n" +
		"    Header --> Progress\n" +
		"    Progress --> Quality\n" +
		"    Quality --> SubScores\n" +
		"    SubScores --> Advice\n" +
		"    Advice --> Final\n" +
		"```\n"

	return content
}

// ProjectEvaluator_WriteAllEvaluations generates all evaluation reports
func ProjectEvaluator_WriteAllEvaluations(outDir string) error {
	fmt.Println("🔍 Generating Project Evaluation Reports...")

	// Generate comprehensive assessment
	if err := ProjectEvaluator_WriteComprehensiveAssessment(outDir); err != nil {
		return fmt.Errorf("failed to write comprehensive assessment: %w", err)
	}

	fmt.Println("✅ Project evaluation reports generated successfully!")
	return nil
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
