//go:build flowcharts
// +build flowcharts

/*
===============================================================================
THEORY TO REALITY - INTELLIGENT PROJECT PROGRESS ANALYSIS
===============================================================================

Author: Ben Tran
Date: 02/09/2025
Description: This file contains intelligent analysis functions that compare
             your real project implementation against the instructor's theory
             model. It provides gap analysis, progress tracking, and next-step
             recommendations based on actual code scanning.

TO USE THIS FILE:
1. Call Theory2Reality_WriteAllAnalysis() to generate all analysis diagrams
2. Each function generates specific progress and gap analysis
3. Diagrams are saved as Theory2Reality_*.mmd.md files for easy identification

THEORY TO REALITY OBJECTIVES:
- Compare real project against instructor's theory model
- Show progress tracking and gap analysis
- Provide next-step recommendations
- Intelligent project status assessment

FEATURES:
- Theory2Reality_progress_analysis.mmd.md - Your actual progress vs theory
- Theory2Reality_gap_analysis.mmd.md - What you still need to implement
- Theory2Reality_next_steps.mmd.md - Recommended next actions
- Theory2Reality_implementation_status.mmd.md - Detailed status breakdown

===============================================================================
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Theory2Reality_WriteAllAnalysis generates all theory-to-reality analysis diagrams
func Theory2Reality_WriteAllAnalysis(outDir string, structure *ProjectStructure) error {
	fmt.Println("🔍 Generating Theory to Reality Analysis...")

	// Generate progress analysis
	if err := Theory2Reality_WriteProgressAnalysis(outDir, structure); err != nil {
		return fmt.Errorf("failed to write progress analysis: %w", err)
	}

	// Generate gap analysis
	if err := Theory2Reality_WriteGapAnalysis(outDir, structure); err != nil {
		return fmt.Errorf("failed to write gap analysis: %w", err)
	}

	// Generate next steps analysis
	if err := Theory2Reality_WriteNextStepsAnalysis(outDir, structure); err != nil {
		return fmt.Errorf("failed to write next steps analysis: %w", err)
	}

	// Generate implementation status
	if err := Theory2Reality_WriteImplementationStatus(outDir, structure); err != nil {
		return fmt.Errorf("failed to write implementation status: %w", err)
	}

	fmt.Println("�� Theory to Reality analysis generated successfully!")
	return nil
}

// Theory2Reality_WriteProgressAnalysis creates a diagram showing your actual progress vs theory
func Theory2Reality_WriteProgressAnalysis(outDir string, structure *ProjectStructure) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph RealProject[\"📊 YOUR REAL PROJECT STATUS\"]\n"

	// Analyze what you've actually implemented
	hasServer := hasBasicServer(structure)
	hasDatabase := hasDatabaseLayer(structure)
	hasCRUD := hasCRUDOperations(structure)
	hasTesting := hasTestingSetup(structure)
	hasAuth := hasAuthentication(structure)
	hasMiddleware := hasMiddlewareLayer(structure)

	// Show what's completed
	if hasServer {
		content += "        REAL1[\"✅ Phase 1: Project Scaffolding<br/>📁 Project structure<br/>🌐 HTTP server<br/>🛣️ Basic routing\"]\n"
	}
	if hasDatabase {
		content += "        REAL2[\"✅ Phase 2: Data Layer<br/>��️ Database setup<br/>📋 Migrations<br/>💾 Data models\"]\n"
	}
	if hasCRUD {
		content += "        REAL3[\"✅ Phase 3: API CRUD Routes<br/>➕ Create operations<br/>🔍 Read operations<br/>✏️ Update operations<br/>🗑️ Delete operations\"]\n"
	}
	if hasTesting {
		content += "        REAL4[\"✅ Phase 4: Testing<br/>🧪 Test setup<br/>✅ Success tests<br/>❌ Error tests\"]\n"
	}
	if hasAuth {
		content += "        REAL5[\"✅ Phase 5: Authentication<br/>👤 User management<br/>🔐 Password security<br/>🎫 JWT tokens\"]\n"
	}
	if hasMiddleware {
		content += "        REAL6[\"✅ Phase 6: Middleware<br/>🛡️ Route protection<br/>🔐 Authorization<br/>✅ User permissions\"]\n"
	}

	content += "    end\n\n" +
		"    subgraph TheoryModel[\"�� INSTRUCTOR'S THEORY MODEL\"]\n" +
		"        THEORY1[\"Phase 1: Project Scaffolding<br/>42m 33s total time\"]\n" +
		"        THEORY2[\"Phase 2: Data Layer<br/>1h 35s total time\"]\n" +
		"        THEORY3[\"Phase 3: API CRUD Routes<br/>1h 24m 15s total time\"]\n" +
		"        THEORY4[\"Phase 4: Testing<br/>38m 20s total time\"]\n" +
		"        THEORY5[\"Phase 5: Authentication<br/>1h 20m 4s total time\"]\n" +
		"        THEORY6[\"Phase 6: Middleware<br/>58m 44s total time\"]\n" +
		"    end\n\n" +
		"    subgraph Progress[\"�� PROGRESS SUMMARY\"]\n"

	// Calculate progress percentage
	completedPhases := 0
	totalPhases := 6

	if hasServer {
		completedPhases++
	}
	if hasDatabase {
		completedPhases++
	}
	if hasCRUD {
		completedPhases++
	}
	if hasTesting {
		completedPhases++
	}
	if hasAuth {
		completedPhases++
	}
	if hasMiddleware {
		completedPhases++
	}

	progressPercent := (completedPhases * 100) / totalPhases

	content += fmt.Sprintf("        PROG1[\"📊 Overall Progress: %d%%<br/>✅ Completed: %d/6 phases<br/>🔄 Remaining: %d phases\"]\n",
		progressPercent, completedPhases, totalPhases-completedPhases)

	content += "    end\n\n" +
		"    %% Connections\n" +
		"    RealProject --> TheoryModel\n" +
		"    TheoryModel --> Progress\n" +
		"    RealProject --> Progress\n" +
		"```\n"

	path := filepath.Join(outDir, "Theory2Reality_progress_analysis.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// Theory2Reality_WriteGapAnalysis creates a diagram showing what you still need to implement
func Theory2Reality_WriteGapAnalysis(outDir string, structure *ProjectStructure) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Completed[\"✅ COMPLETED IMPLEMENTATIONS\"]\n"

	// Show what's completed
	hasServer := hasBasicServer(structure)
	hasDatabase := hasDatabaseLayer(structure)
	hasCRUD := hasCRUDOperations(structure)
	hasTesting := hasTestingSetup(structure)
	hasAuth := hasAuthentication(structure)
	hasMiddleware := hasMiddlewareLayer(structure)

	if hasServer {
		content += "        COMP1[\"✅ Project Scaffolding<br/>• Go project structure<br/>• HTTP server<br/>• Basic routing\"]\n"
	}
	if hasDatabase {
		content += "        COMP2[\"✅ Data Layer<br/>• Database setup<br/>• Migrations<br/>• Data models\"]\n"
	}
	if hasCRUD {
		content += "        COMP3[\"✅ CRUD Operations<br/>• Create handlers<br/>• Read handlers<br/>• Update handlers<br/>• Delete handlers\"]\n"
	}
	if hasTesting {
		content += "        COMP4[\"✅ Testing Setup<br/>• Test database<br/>• Unit tests<br/>• Test execution\"]\n"
	}
	if hasAuth {
		content += "        COMP5[\"✅ Authentication<br/>• User management<br/>• Password security<br/>• JWT tokens\"]\n"
	}
	if hasMiddleware {
		content += "        COMP6[\"✅ Middleware<br/>• Route protection<br/>• Authorization<br/>• User permissions\"]\n"
	}

	content += "    end\n\n" +
		"    subgraph Missing[\"🔄 MISSING IMPLEMENTATIONS\"]\n"

	// Show what's missing
	if !hasServer {
		content += "        MISS1[\"🔄 Project Scaffolding<br/>• Create Go project<br/>• HTTP server setup<br/>• Basic routing\"]\n"
	}
	if !hasDatabase {
		content += "        MISS2[\"�� Data Layer<br/>• Docker database<br/>• Migrations<br/>• Data models\"]\n"
	}
	if !hasCRUD {
		content += "        MISS3[\"�� CRUD Operations<br/>• Create handlers<br/>• Read handlers<br/>• Update handlers<br/>• Delete handlers\"]\n"
	}
	if !hasTesting {
		content += "        MISS4[\"🔄 Testing Setup<br/>• Test database<br/>• Unit tests<br/>• Test execution\"]\n"
	}
	if !hasAuth {
		content += "        MISS5[\"�� Authentication<br/>• User management<br/>• Password security<br/>• JWT tokens\"]\n"
	}
	if !hasMiddleware {
		content += "        MISS6[\"🔄 Middleware<br/>• Route protection<br/>• Authorization<br/>• User permissions\"]\n"
	}

	content += "    end\n\n" +
		"    subgraph Priority[\"🎯 IMPLEMENTATION PRIORITY\"]\n" +
		"        PRIO1[\"🔥 HIGH PRIORITY<br/>Next logical step<br/>based on dependencies\"]\n" +
		"        PRIO2[\"⚡ MEDIUM PRIORITY<br/>Can be implemented<br/>in parallel\"]\n" +
		"        PRIO3[\"�� LOW PRIORITY<br/>Future enhancements\"]\n" +
		"    end\n\n" +
		"    %% Connections\n" +
		"    Completed --> Missing\n" +
		"    Missing --> Priority\n" +
		"```\n"

	path := filepath.Join(outDir, "Theory2Reality_gap_analysis.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// Theory2Reality_WriteNextStepsAnalysis creates a diagram showing recommended next actions
func Theory2Reality_WriteNextStepsAnalysis(outDir string, structure *ProjectStructure) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Current[\"📍 CURRENT STATUS\"]\n"

	// Determine current phase
	hasServer := hasBasicServer(structure)
	hasDatabase := hasDatabaseLayer(structure)
	hasCRUD := hasCRUDOperations(structure)
	hasTesting := hasTestingSetup(structure)
	hasAuth := hasAuthentication(structure)
	hasMiddleware := hasMiddlewareLayer(structure)

	if hasMiddleware {
		content += "        CURR[\"🎉 PROJECT COMPLETE!<br/>All phases implemented<br/>Ready for production\"]\n"
	} else if hasAuth {
		content += "        CURR[\"🔄 Phase 6: Middleware<br/>Implement route protection<br/>and authorization\"]\n"
	} else if hasTesting {
		content += "        CURR[\"🔄 Phase 5: Authentication<br/>Implement user management<br/>and JWT tokens\"]\n"
	} else if hasCRUD {
		content += "        CURR[\"�� Phase 4: Testing<br/>Set up test database<br/>and write unit tests\"]\n"
	} else if hasDatabase {
		content += "        CURR[\"🔄 Phase 3: CRUD Operations<br/>Implement API handlers<br/>for all operations\"]\n"
	} else if hasServer {
		content += "        CURR[\"🔄 Phase 2: Data Layer<br/>Set up database<br/>and migrations\"]\n"
	} else {
		content += "        CURR[\"🔄 Phase 1: Project Scaffolding<br/>Create basic project<br/>structure and server\"]\n"
	}

	content += "    end\n\n" +
		"    subgraph NextSteps[\"🎯 RECOMMENDED NEXT STEPS\"]\n"

	// Provide specific next steps based on current status
	if !hasServer {
		content += "        NEXT1[\"1. Create Go Project<br/>�� go mod init workout-api<br/>�� mkdir internal/{app,api,store}\"]\n" +
			"        NEXT2[\"2. HTTP Server<br/>🌐 Basic server setup<br/>🌐 Listen on port 8080\"]\n" +
			"        NEXT3[\"3. Basic Routing<br/>🛣️ Add Chi router<br/>🛣️ Health check endpoint\"]\n"
	} else if !hasDatabase {
		content += "        NEXT1[\"1. Docker Database<br/>�� Create docker-compose.yml<br/>🐳 PostgreSQL service\"]\n" +
			"        NEXT2[\"2. Database Driver<br/>🔌 Add pgx dependency<br/>🔌 Connection setup\"]\n" +
			"        NEXT3[\"3. Migrations<br/>📋 Add Goose dependency<br/>📋 First migration file\"]\n"
	} else if !hasCRUD {
		content += "        NEXT1[\"1. Create Handler<br/>➕ POST endpoint<br/>➕ Data creation\"]\n" +
			"        NEXT2[\"2. Read Handler<br/>�� GET by ID<br/>�� Data retrieval\"]\n" +
			"        NEXT3[\"3. Update Handler<br/>✏️ PUT/PATCH<br/>✏️ Data modification\"]\n" +
			"        NEXT4[\"4. Delete Handler<br/>��️ DELETE<br/>🗑️ Data removal\"]\n"
	} else if !hasTesting {
		content += "        NEXT1[\"1. Test Database<br/>🗄️ Separate test DB<br/>🗄️ Test environment\"]\n" +
			"        NEXT2[\"2. Unit Tests<br/>🧪 Test functions<br/>�� Success scenarios\"]\n" +
			"        NEXT3[\"3. Error Tests<br/>❌ Error scenarios<br/>❌ Edge cases\"]\n"
	} else if !hasAuth {
		content += "        NEXT1[\"1. User Model<br/>�� User struct<br/>👤 User database\"]\n" +
			"        NEXT2[\"2. Password Security<br/>🔒 Password hashing<br/>🔒 Secure storage\"]\n" +
			"        NEXT3[\"3. JWT Tokens<br/>🎫 Token generation<br/>�� Token validation\"]\n"
	} else if !hasMiddleware {
		content += "        NEXT1[\"1. Auth Middleware<br/>🔐 Token validation<br/>�� User context\"]\n" +
			"        NEXT2[\"2. Route Protection<br/>🛡️ Protected endpoints<br/>🛡️ Access control\"]\n" +
			"        NEXT3[\"3. Ownership Validation<br/>✅ Resource ownership<br/>✅ User permissions\"]\n"
	}

	content += "    end\n\n" +
		"    subgraph Resources[\"�� LEARNING RESOURCES\"]\n" +
		"        RES1[\"�� Instructor's Model<br/>Follow the exact progression<br/>from IntructorProjectBuilderModel.txt\"]\n" +
		"        RES2[\"📖 Go Documentation<br/>Official Go docs<br/>for specific implementations\"]\n" +
		"        RES3[\"🔍 Code Examples<br/>Look at existing functions<br/>for patterns and structure\"]\n" +
		"    end\n\n" +
		"    %% Connections\n" +
		"    Current --> NextSteps\n" +
		"    NextSteps --> Resources\n" +
		"```\n"

	path := filepath.Join(outDir, "Theory2Reality_next_steps.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// Theory2Reality_WriteImplementationStatus creates a detailed status breakdown
func Theory2Reality_WriteImplementationStatus(outDir string, structure *ProjectStructure) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Status[\"�� DETAILED IMPLEMENTATION STATUS\"]\n"

	// Detailed analysis of each phase
	hasServer := hasBasicServer(structure)
	hasDatabase := hasDatabaseLayer(structure)
	hasCRUD := hasCRUDOperations(structure)
	hasTesting := hasTestingSetup(structure)
	hasAuth := hasAuthentication(structure)
	hasMiddleware := hasMiddlewareLayer(structure)

	// Phase 1: Project Scaffolding
	if hasServer {
		content += "        STAT1[\"✅ Phase 1: Project Scaffolding<br/>📁 Project structure: ✅<br/>�� HTTP server: ✅<br/>��️ Basic routing: ✅<br/>⚙️ Configuration: ✅\"]\n"
	} else {
		content += "        STAT1[\"🔄 Phase 1: Project Scaffolding<br/>📁 Project structure: ❌<br/>�� HTTP server: ❌<br/>🛣️ Basic routing: ❌<br/>⚙️ Configuration: ❌\"]\n"
	}

	// Phase 2: Data Layer
	if hasDatabase {
		content += "        STAT2[\"✅ Phase 2: Data Layer<br/>🐳 Docker database: ✅<br/>🔌 Database driver: ✅<br/>📋 Migrations: ✅<br/>�� Data models: ✅\"]\n"
	} else {
		content += "        STAT2[\"🔄 Phase 2: Data Layer<br/>🐳 Docker database: ❌<br/>🔌 Database driver: ❌<br/>📋 Migrations: ❌<br/>💾 Data models: ❌\"]\n"
	}

	// Phase 3: CRUD Operations
	if hasCRUD {
		content += "        STAT3[\"✅ Phase 3: CRUD Operations<br/>➕ Create operations: ✅<br/>🔍 Read operations: ✅<br/>✏️ Update operations: ✅<br/>🗑️ Delete operations: ✅\"]\n"
	} else {
		content += "        STAT3[\"🔄 Phase 3: CRUD Operations<br/>➕ Create operations: ❌<br/>🔍 Read operations: ❌<br/>✏️ Update operations: ❌<br/>🗑️ Delete operations: ❌\"]\n"
	}

	// Phase 4: Testing
	if hasTesting {
		content += "        STAT4[\"✅ Phase 4: Testing<br/>🗄️ Test database: ✅<br/>�� Unit tests: ✅<br/>✅ Success tests: ✅<br/>❌ Error tests: ✅\"]\n"
	} else {
		content += "        STAT4[\"�� Phase 4: Testing<br/>🗄️ Test database: ❌<br/>�� Unit tests: ❌<br/>✅ Success tests: ❌<br/>❌ Error tests: ❌\"]\n"
	}

	// Phase 5: Authentication
	if hasAuth {
		content += "        STAT5[\"✅ Phase 5: Authentication<br/>👤 User management: ✅<br/>🔐 Password security: ✅<br/>🎫 JWT tokens: ✅<br/>�� Auth endpoints: ✅\"]\n"
	} else {
		content += "        STAT5[\"🔄 Phase 5: Authentication<br/>👤 User management: ❌<br/>🔐 Password security: ❌<br/>�� JWT tokens: ❌<br/>�� Auth endpoints: ❌\"]\n"
	}

	// Phase 6: Middleware
	if hasMiddleware {
		content += "        STAT6[\"✅ Phase 6: Middleware<br/>🔐 Auth middleware: ✅<br/>🛡️ Route protection: ✅<br/>✅ User permissions: ✅<br/>📝 Context management: ✅\"]\n"
	} else {
		content += "        STAT6[\"🔄 Phase 6: Middleware<br/>🔐 Auth middleware: ❌<br/>🛡️ Route protection: ❌<br/>✅ User permissions: ❌<br/>📝 Context management: ❌\"]\n"
	}

	content += "    end\n\n" +
		"    subgraph Summary[\"📈 IMPLEMENTATION SUMMARY\"]\n"

	// Calculate detailed statistics
	completedPhases := 0
	totalPhases := 6

	if hasServer {
		completedPhases++
	}
	if hasDatabase {
		completedPhases++
	}
	if hasCRUD {
		completedPhases++
	}
	if hasTesting {
		completedPhases++
	}
	if hasAuth {
		completedPhases++
	}
	if hasMiddleware {
		completedPhases++
	}

	progressPercent := (completedPhases * 100) / totalPhases

	content += fmt.Sprintf("        SUM1[\"📊 Overall Progress: %d%%<br/>✅ Completed Phases: %d/6<br/>�� Remaining Phases: %d<br/>📁 Total Functions: %d<br/>📄 Total Files: %d\"]\n",
		progressPercent, completedPhases, totalPhases-completedPhases, len(structure.Functions), len(structure.Files))

	content += "    end\n\n" +
		"    %% Connections\n" +
		"    Status --> Summary\n" +
		"```\n"

	path := filepath.Join(outDir, "Theory2Reality_implementation_status.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// Helper functions to analyze project structure
func hasBasicServer(structure *ProjectStructure) bool {
	for _, fn := range structure.Functions {
		if strings.Contains(strings.ToLower(fn.Name), "main") &&
			strings.Contains(strings.ToLower(fn.Name), "server") {
			return true
		}
	}
	return false
}

func hasDatabaseLayer(structure *ProjectStructure) bool {
	for _, fn := range structure.Functions {
		if strings.Contains(strings.ToLower(fn.Name), "db") ||
			strings.Contains(strings.ToLower(fn.Name), "database") ||
			strings.Contains(strings.ToLower(fn.Name), "migrate") {
			return true
		}
	}
	return false
}

func hasCRUDOperations(structure *ProjectStructure) bool {
	crudCount := 0
	for _, fn := range structure.Functions {
		if strings.Contains(strings.ToLower(fn.Name), "create") ||
			strings.Contains(strings.ToLower(fn.Name), "read") ||
			strings.Contains(strings.ToLower(fn.Name), "update") ||
			strings.Contains(strings.ToLower(fn.Name), "delete") {
			crudCount++
		}
	}
	return crudCount >= 3 // Need at least 3 CRUD operations
}

func hasTestingSetup(structure *ProjectStructure) bool {
	for _, fn := range structure.Functions {
		if strings.Contains(strings.ToLower(fn.Name), "test") {
			return true
		}
	}
	return false
}

func hasAuthentication(structure *ProjectStructure) bool {
	for _, fn := range structure.Functions {
		if strings.Contains(strings.ToLower(fn.Name), "auth") ||
			strings.Contains(strings.ToLower(fn.Name), "token") ||
			strings.Contains(strings.ToLower(fn.Name), "jwt") ||
			strings.Contains(strings.ToLower(fn.Name), "login") {
			return true
		}
	}
	return false
}

func hasMiddlewareLayer(structure *ProjectStructure) bool {
	for _, fn := range structure.Functions {
		if strings.Contains(strings.ToLower(fn.Name), "middleware") ||
			strings.Contains(strings.ToLower(fn.Name), "auth") {
			return true
		}
	}
	return false
}
