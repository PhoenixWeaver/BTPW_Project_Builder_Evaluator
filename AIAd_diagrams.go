/*===============================================================================
🐦 ::: PhoenixFlix - Multi-Purpose Movies & Christian Streaming Platform :::
🔥 with dual database architecture, WebAuthn authentication, and family-friendly streaming experience.
===============================================================================
Author: Ben Tran (https://github.com/thephoenixflix)
Email: thephoenixflix@gmail.com
Website: https://bit.ly/thephoenixflix
===============================================================================*/

/*===============================================================================
AI ADVISOR DIAGRAMS - STEP-BY-STEP PROJECT RECREATION GUIDE
===============================================================================

Author: Ben Tran
Date: 02/09/2025
Description: This file contains AI Advisor functions that provide step-by-step
             guidance for recreating projects from scratch. These functions
             generate diagrams showing build order, dependencies, and development
             sequence to help developers understand where to start and what to
             build next.

TO USE THIS FILE:
1. Call AIAd_WriteAllStructureDiagrams() to generate all AI advice diagrams
2. Each function generates specific guidance for project recreation
3. Diagrams are saved as AIAd_*.mmd.md files for easy identification

AI ADVISOR OBJECTIVES:
- Step-by-step project recreation guidance
- Build order recommendations based on dependencies
- Development sequence understanding
- Complete project building instructions

FEATURES:
- AIAd_development_sequence.mmd.md - How functions were created
- AIAd_execution_flow.mmd.md - How functions execute at runtime
- AIAd_function_dependencies.mmd.md - What to build first
- AIAd_project_building_guide.md - Complete step-by-step guide

===============================================================================
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// AIAd_WriteFunctionFlowAnalysis generates comprehensive AI advisor function flow analysis diagrams.
// This is the main orchestrator function that calls all individual AI advisor diagram functions.
func AIAd_WriteFunctionFlowAnalysis(outDir string) error {
	fmt.Println("🎯 Generating AI Advisor Function Flow Analysis...")

	// Generate development sequence diagram
	if err := AIAd_WriteDevelopmentSequenceDiagram(outDir); err != nil {
		return fmt.Errorf("failed to write AI advisor development sequence diagram: %w", err)
	}
	fmt.Println("✅ Generated AIAd_development_sequence.mmd.md")

	// Generate execution flow diagram
	if err := AIAd_WriteExecutionFlowDiagram(outDir); err != nil {
		return fmt.Errorf("failed to write AI advisor execution flow diagram: %w", err)
	}
	fmt.Println("✅ Generated AIAd_execution_flow.mmd.md")

	// Generate function dependency diagram
	if err := AIAd_WriteFunctionDependencyDiagram(outDir); err != nil {
		return fmt.Errorf("failed to write AI advisor function dependency diagram: %w", err)
	}
	fmt.Println("✅ Generated AIAd_function_dependencies.mmd.md")

	// Generate project building guide
	if err := AIAd_WriteProjectBuildingGuide(outDir); err != nil {
		return fmt.Errorf("failed to write AI advisor project building guide: %w", err)
	}
	fmt.Println("✅ Generated AIAd_project_building_guide.md")

	fmt.Println("🎉 AI Advisor Function Flow Analysis Complete!")
	return nil
}

// AIAd_WriteDevelopmentSequenceDiagram writes a Mermaid diagram showing the order functions were created during development.
// This diagram helps understand the development sequence and where to start when building similar projects.
func AIAd_WriteDevelopmentSequenceDiagram(outDir string) error {
	content := `# AI Advisor: Development Sequence - How Functions Were Created

This diagram shows the **order in which functions were created** during development.
Understanding this helps you know **where to start** when building similar projects.

` + "```mermaid\n" + `
flowchart TD
    subgraph Phase1["🚀 PHASE 1: Foundation (Start Here)"]
        F1["1. main()<br/>📍 Ex10.go<br/>🎯 Entry point<br/>Creates application"]
        F2["2. NewApplication()<br/>📍 internal/app/app.go<br/>🎯 Application factory<br/>Orchestrates everything"]
        F3["3. Open()<br/>📍 internal/store/database.go<br/>🎯 Database connection<br/>Foundation for data"]
    end
    
    subgraph Phase2["🏗️ PHASE 2: Data Layer"]
        F4["4. MigrateFS()<br/>📍 internal/store/database.go<br/>🎯 Database migrations<br/>Creates tables"]
        F5["5. User struct<br/>📍 internal/store/user_store.go<br/>🎯 Data model<br/>Defines user structure"]
        F6["6. password.Set()<br/>📍 internal/store/user_store.go<br/>🎯 Password hashing<br/>Security foundation"]
        F7["7. password.Matches()<br/>📍 internal/store/user_store.go<br/>🎯 Password verification<br/>Authentication logic"]
    end
    
    subgraph Phase3["🏪 PHASE 3: Store Layer"]
        F8["8. NewPostgresUserStore()<br/>📍 internal/store/user_store.go<br/>🎯 User store factory<br/>Data access pattern"]
        F9["9. CreateUser()<br/>📍 internal/store/user_store.go<br/>🎯 User creation<br/>Database operations"]
        F10["10. GetUserByUsername()<br/>📍 internal/store/user_store.go<br/>🎯 User retrieval<br/>Authentication support"]
        F11["11. Workout struct<br/>📍 internal/store/workout_store.go<br/>🎯 Workout model<br/>Complex data structure"]
        F12["12. WorkoutEntry struct<br/>📍 internal/store/workout_store.go<br/>🎯 Entry model<br/>Related data structure"]
    end
    
    subgraph Phase4["🏪 PHASE 4: Complex Store Operations"]
        F13["13. NewPostgresWorkoutStore()<br/>📍 internal/store/workout_store.go<br/>🎯 Workout store factory<br/>Complex data access"]
        F14["14. CreateWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Workout creation<br/>Transaction management"]
        F15["15. GetWorkoutByID()<br/>📍 internal/store/workout_store.go<br/>🎯 Workout retrieval<br/>Complex queries"]
        F16["16. UpdateWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Workout updates<br/>Data modification"]
        F17["17. DeleteWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Workout deletion<br/>Data cleanup"]
    end
    
    subgraph Phase5["🌐 PHASE 5: API Layer"]
        F18["18. NewUserHandler()<br/>📍 internal/api/user_handler.go<br/>🎯 User handler factory<br/>HTTP layer foundation"]
        F19["19. validateRegisterRequest()<br/>📍 internal/api/user_handler.go<br/>🎯 Input validation<br/>Data safety"]
        F20["20. HandleRegisterUser()<br/>📍 internal/api/user_handler.go<br/>🎯 User registration<br/>HTTP endpoint"]
        F21["21. NewWorkoutHandler()<br/>📍 internal/api/workout_handler.go<br/>🎯 Workout handler factory<br/>HTTP layer"]
        F22["22. HandleGetWorkoutByID()<br/>📍 internal/api/workout_handler.go<br/>🎯 Get workout endpoint<br/>HTTP GET"]
    end
    
    subgraph Phase6["🌐 PHASE 6: Complex API Operations"]
        F23["23. HandleCreateWorkout()<br/>📍 internal/api/workout_handler.go<br/>🎯 Create workout endpoint<br/>HTTP POST"]
        F24["24. HandleUpdateWorkoutByID()<br/>📍 internal/api/workout_handler.go<br/>🎯 Update workout endpoint<br/>HTTP PUT"]
        F25["25. HandleDeleteWorkoutByID()<br/>📍 internal/api/workout_handler.go<br/>🎯 Delete workout endpoint<br/>HTTP DELETE"]
    end
    
    subgraph Phase7["🛣️ PHASE 7: Routing & Utilities"]
        F26["26. SetupRoutes()<br/>📍 internal/routes/routes.go<br/>🎯 Route configuration<br/>URL mapping"]
        F27["27. WriteJSON()<br/>📍 internal/utils/utils.go<br/>🎯 JSON response utility<br/>Consistent responses"]
        F28["28. ReadIDParam()<br/>📍 internal/utils/utils.go<br/>🎯 Parameter extraction<br/>URL parameter handling"]
    end
    
    subgraph Phase8["🛡️ PHASE 8: Middleware & Security (Future)"]
        F29["29. AuthMiddleware()<br/>📍 internal/middleware/auth.go<br/>🎯 Authentication middleware<br/>JWT token validation"]
        F30["30. LoggingMiddleware()<br/>📍 internal/middleware/logging.go<br/>🎯 Request logging<br/>HTTP request/response logging"]
        F31["31. CORSMiddleware()<br/>📍 internal/middleware/cors.go<br/>🎯 CORS handling<br/>Cross-origin requests"]
        F32["32. RateLimitMiddleware()<br/>📍 internal/middleware/ratelimit.go<br/>🎯 Rate limiting<br/>Prevent abuse"]
    end
    
    %% Development flow
    F1 --> F2
    F2 --> F3
    F3 --> F4
    F4 --> F5
    F5 --> F6
    F6 --> F7
    F7 --> F8
    F8 --> F9
    F9 --> F10
    F10 --> F11
    F11 --> F12
    F12 --> F13
    F13 --> F14
    F14 --> F15
    F15 --> F16
    F16 --> F17
    F17 --> F18
    F18 --> F19
    F19 --> F20
    F20 --> F21
    F21 --> F22
    F22 --> F23
    F23 --> F24
    F24 --> F25
    F25 --> F26
    F26 --> F27
    F27 --> F28
    F28 --> F29
    F29 --> F30
    F30 --> F31
    F31 --> F32
    
` + "```\n"

	path := filepath.Join(outDir, "AIAd_development_sequence.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// AIAd_WriteExecutionFlowDiagram writes a Mermaid diagram showing the order functions execute at runtime.
// This diagram helps understand how the application works step by step during execution.
func AIAd_WriteExecutionFlowDiagram(outDir string) error {
	content := `# AI Advisor: Execution Flow - How Functions Execute at Runtime

This diagram shows the **order in which functions execute** when the application runs.
Understanding this helps you know **how the application works** step by step.

` + "```mermaid\n" + `
flowchart TD
    subgraph Startup["🚀 APPLICATION STARTUP"]
        E1["1. main()<br/>📍 Ex10.go:15<br/>🎯 Program entry point<br/>▶️ Creates application instance"]
        E2["2. NewApplication()<br/>📍 internal/app/app.go:134<br/>🎯 Application factory<br/>▶️ Orchestrates all initialization"]
        E3["3. Open()<br/>📍 internal/store/database.go:117<br/>🎯 Database connection<br/>▶️ Connects to PostgreSQL"]
        E4["4. MigrateFS()<br/>📍 internal/store/database.go:198<br/>🎯 Run migrations<br/>▶️ Creates/updates database schema"]
        E5["5. NewPostgresUserStore()<br/>📍 internal/store/user_store.go:84<br/>🎯 Create user store<br/>▶️ Initializes user data access"]
        E6["6. NewPostgresWorkoutStore()<br/>📍 internal/store/workout_store.go:86<br/>🎯 Create workout store<br/>▶️ Initializes workout data access"]
        E7["7. NewUserHandler()<br/>📍 internal/api/user_handler.go:86<br/>🎯 Create user handler<br/>▶️ Initializes user HTTP handling"]
        E8["8. NewWorkoutHandler()<br/>📍 internal/api/workout_handler.go:264<br/>🎯 Create workout handler<br/>▶️ Initializes workout HTTP handling"]
        E9["9. SetupRoutes()<br/>📍 internal/routes/routes.go<br/>🎯 Configure routes<br/>▶️ Maps URLs to handlers"]
        E10["10. http.ListenAndServe()<br/>📍 Ex10.go:25<br/>🎯 Start HTTP server<br/>▶️ Begins accepting requests"]
    end
    
    subgraph UserRegistration["👤 USER REGISTRATION REQUEST"]
        UR1["HTTP POST /users/register<br/>🌐 Client request"]
        UR2["HandleRegisterUser()<br/>📍 internal/api/user_handler.go:156<br/>🎯 Process registration<br/>▶️ Handles HTTP request"]
        UR3["validateRegisterRequest()<br/>📍 internal/api/user_handler.go:99<br/>🎯 Validate input<br/>▶️ Checks data integrity"]
        UR4["password.Set()<br/>📍 internal/store/user_store.go:72<br/>🎯 Hash password<br/>▶️ Secures password"]
        UR5["CreateUser()<br/>📍 internal/store/user_store.go:167<br/>🎯 Save user<br/>▶️ Stores in database"]
        UR6["WriteJSON()<br/>📍 internal/utils/utils.go<br/>🎯 Send response<br/>▶️ Returns JSON to client"]
    end
    
    subgraph WorkoutCreation["🏋️ WORKOUT CREATION REQUEST"]
        WC1["HTTP POST /workouts<br/>🌐 Client request"]
        WC2["HandleCreateWorkout()<br/>📍 internal/api/workout_handler.go:419<br/>🎯 Process creation<br/>▶️ Handles HTTP request"]
        WC3["CreateWorkout()<br/>📍 internal/store/workout_store.go:125<br/>🎯 Save workout<br/>▶️ Database transaction"]
        WC4["WriteJSON()<br/>📍 internal/utils/utils.go<br/>🎯 Send response<br/>▶️ Returns created workout"]
    end
    
    subgraph WorkoutRetrieval["📖 WORKOUT RETRIEVAL REQUEST"]
        WR1["HTTP GET /workouts/{id}<br/>🌐 Client request"]
        WR2["HandleGetWorkoutByID()<br/>📍 internal/api/workout_handler.go:305<br/>🎯 Process retrieval<br/>▶️ Handles HTTP request"]
        WR3["ReadIDParam()<br/>📍 internal/utils/utils.go<br/>🎯 Extract ID<br/>▶️ Parses URL parameter"]
        WR4["GetWorkoutByID()<br/>📍 internal/store/workout_store.go:185<br/>🎯 Fetch workout<br/>▶️ Database query"]
        WR5["WriteJSON()<br/>📍 internal/utils/utils.go<br/>🎯 Send response<br/>▶️ Returns workout data"]
    end
    
    %% Startup flow
    E1 --> E2
    E2 --> E3
    E3 --> E4
    E4 --> E5
    E5 --> E6
    E6 --> E7
    E7 --> E8
    E8 --> E9
    E9 --> E10
    
    %% User registration flow
    E10 --> UR1
    UR1 --> UR2
    UR2 --> UR3
    UR3 --> UR4
    UR4 --> UR5
    UR5 --> UR6
    
    %% Workout creation flow
    E10 --> WC1
    WC1 --> WC2
    WC2 --> WC3
    WC3 --> WC4
    
    %% Workout retrieval flow
    E10 --> WR1
    WR1 --> WR2
    WR2 --> WR3
    WR3 --> WR4
    WR4 --> WR5
    
` + "```\n"

	path := filepath.Join(outDir, "AIAd_execution_flow.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// AIAd_WriteFunctionDependencyDiagram writes a diagram showing which functions depend on which other functions.
// This diagram helps understand what to build first and the dependency relationships.
func AIAd_WriteFunctionDependencyDiagram(outDir string) error {
	content := "```mermaid\n" + `
flowchart TD
    subgraph Foundation["🏗️ FOUNDATION FUNCTIONS (Build First)"]
        DB_OPEN["Open()<br/>📍 database.go<br/>🎯 No dependencies<br/>Pure database connection"]
        DB_MIGRATE["MigrateFS()<br/>📍 database.go<br/>🎯 Depends: Open()<br/>Needs database connection"]
        PASS_SET["password.Set()<br/>📍 user_store.go<br/>🎯 No dependencies<br/>Pure bcrypt hashing"]
        PASS_MATCH["password.Matches()<br/>📍 user_store.go<br/>🎯 No dependencies<br/>Pure bcrypt verification"]
    end
    
    subgraph DataAccess["💾 DATA ACCESS FUNCTIONS (Build Second)"]
        USER_STORE_NEW["NewPostgresUserStore()<br/>📍 user_store.go<br/>🎯 Depends: *sql.DB<br/>Store factory"]
        USER_CREATE["CreateUser()<br/>📍 user_store.go<br/>🎯 Depends: password.Set()<br/>Database operations"]
        USER_GET["GetUserByUsername()<br/>📍 user_store.go<br/>🎯 Depends: *sql.DB<br/>Database queries"]
        
        WORKOUT_STORE_NEW["NewPostgresWorkoutStore()<br/>📍 workout_store.go<br/>🎯 Depends: *sql.DB<br/>Store factory"]
        WORKOUT_CREATE["CreateWorkout()<br/>📍 workout_store.go<br/>🎯 Depends: *sql.DB<br/>Transaction management"]
        WORKOUT_GET["GetWorkoutByID()<br/>📍 workout_store.go<br/>🎯 Depends: *sql.DB<br/>Complex queries"]
        WORKOUT_UPDATE["UpdateWorkout()<br/>📍 workout_store.go<br/>🎯 Depends: *sql.DB<br/>Data modification"]
        WORKOUT_DELETE["DeleteWorkout()<br/>📍 workout_store.go<br/>🎯 Depends: *sql.DB<br/>Data cleanup"]
    end
    
    subgraph Application["🎯 APPLICATION FUNCTIONS (Build Third)"]
        APP_NEW["NewApplication()<br/>📍 app.go<br/>🎯 Depends: ALL stores<br/>🏆 MOST COMPLEX<br/>Orchestrates everything"]
    end
    
    subgraph Utilities["🔧 UTILITY FUNCTIONS (Build Fourth)"]
        WRITE_JSON["WriteJSON()<br/>📍 utils.go<br/>🎯 No dependencies<br/>Pure JSON encoding"]
        READ_ID["ReadIDParam()<br/>📍 utils.go<br/>🎯 No dependencies<br/>Pure parameter parsing"]
    end
    
    subgraph Routing["🛣️ ROUTING FUNCTIONS (Build Fifth)"]
        SETUP_ROUTES["SetupRoutes()<br/>📍 routes.go<br/>🎯 Depends: ALL handlers<br/>URL mapping"]
    end
    
    subgraph APIHandlers["🌐 API HANDLER FUNCTIONS (Build Sixth)"]
        USER_HANDLER_NEW["NewUserHandler()<br/>📍 user_handler.go<br/>🎯 Depends: UserStore, Logger<br/>Handler factory"]
        USER_VALIDATE["validateRegisterRequest()<br/>📍 user_handler.go<br/>🎯 No dependencies<br/>Pure validation"]
        USER_REGISTER["HandleRegisterUser()<br/>📍 user_handler.go<br/>🎯 Depends: validate, CreateUser, WriteJSON<br/>HTTP endpoint"]
        
        WORKOUT_HANDLER_NEW["NewWorkoutHandler()<br/>📍 workout_handler.go<br/>🎯 Depends: WorkoutStore, Logger<br/>Handler factory"]
        WORKOUT_GET_HANDLER["HandleGetWorkoutByID()<br/>📍 workout_handler.go<br/>🎯 Depends: ReadIDParam, GetWorkoutByID, WriteJSON<br/>HTTP GET"]
        WORKOUT_CREATE_HANDLER["HandleCreateWorkout()<br/>📍 workout_handler.go<br/>🎯 Depends: CreateWorkout, WriteJSON<br/>HTTP POST"]
        WORKOUT_UPDATE_HANDLER["HandleUpdateWorkoutByID()<br/>📍 workout_handler.go<br/>🎯 Depends: ReadIDParam, GetWorkoutByID, UpdateWorkout, WriteJSON<br/>HTTP PUT"]
        WORKOUT_DELETE_HANDLER["HandleDeleteWorkoutByID()<br/>📍 workout_handler.go<br/>🎯 Depends: DeleteWorkout<br/>HTTP DELETE"]
    end
    
    subgraph MainApp["🚀 MAIN FUNCTIONS (Build Last)"]
        MAIN["main()<br/>📍 Ex10.go<br/>🎯 Depends: NewApplication, SetupRoutes<br/>Program entry"]
    end
    
    %% Foundation dependencies
    DB_MIGRATE --> DB_OPEN
    
    %% Data access dependencies
    USER_STORE_NEW --> DB_OPEN
    USER_CREATE --> PASS_SET
    USER_CREATE --> USER_STORE_NEW
    USER_GET --> USER_STORE_NEW
    
    WORKOUT_STORE_NEW --> DB_OPEN
    WORKOUT_CREATE --> WORKOUT_STORE_NEW
    WORKOUT_GET --> WORKOUT_STORE_NEW
    WORKOUT_UPDATE --> WORKOUT_STORE_NEW
    WORKOUT_DELETE --> WORKOUT_STORE_NEW
    
    %% Application dependencies
    APP_NEW --> DB_OPEN
    APP_NEW --> DB_MIGRATE
    APP_NEW --> USER_STORE_NEW
    APP_NEW --> WORKOUT_STORE_NEW
    APP_NEW --> USER_HANDLER_NEW
    APP_NEW --> WORKOUT_HANDLER_NEW
    
    %% Handler dependencies
    USER_HANDLER_NEW --> USER_STORE_NEW
    USER_REGISTER --> USER_VALIDATE
    USER_REGISTER --> USER_CREATE
    USER_REGISTER --> WRITE_JSON
    
    WORKOUT_HANDLER_NEW --> WORKOUT_STORE_NEW
    WORKOUT_GET_HANDLER --> READ_ID
    WORKOUT_GET_HANDLER --> WORKOUT_GET
    WORKOUT_GET_HANDLER --> WRITE_JSON
    WORKOUT_CREATE_HANDLER --> WORKOUT_CREATE
    WORKOUT_CREATE_HANDLER --> WRITE_JSON
    WORKOUT_UPDATE_HANDLER --> READ_ID
    WORKOUT_UPDATE_HANDLER --> WORKOUT_GET
    WORKOUT_UPDATE_HANDLER --> WORKOUT_UPDATE
    WORKOUT_UPDATE_HANDLER --> WRITE_JSON
    WORKOUT_DELETE_HANDLER --> WORKOUT_DELETE
    
    %% Routing dependencies
    SETUP_ROUTES --> USER_HANDLER_NEW
    SETUP_ROUTES --> WORKOUT_HANDLER_NEW
    
    %% Main dependencies
    MAIN --> APP_NEW
    MAIN --> SETUP_ROUTES
    
` + "```\n"

	path := filepath.Join(outDir, "AIAd_function_dependencies.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// AIAd_WriteProjectBuildingGuide writes a comprehensive guide for building the project from scratch.
// This guide provides step-by-step instructions for recreating the project.
func AIAd_WriteProjectBuildingGuide(outDir string) error {
	content := `# AI Advisor: How to Build This Project from Scratch

This guide shows you **exactly where to start** and **what to build next** when creating this project.

## 🎯 BUILD ORDER (Follow This Sequence)

### STEP 1: Project Foundation (Start Here!)
` + "```bash\n" + `
# 1. Create project structure
mkdir your-project && cd your-project
go mod init your-project

# 2. Create directory structure
mkdir -p internal/{store,api,app,routes,utils}
mkdir -p migrations
mkdir -p cmd/server
` + "```\n" + `

### STEP 2: Database Layer (Build This First)
**Why first?** Everything depends on data storage.

#### 2.1 Create database.go
- ✅ **Function:** ` + "`Open()`" + ` - Database connection
- ✅ **Function:** ` + "`MigrateFS()`" + ` - Schema migrations
- 📍 **Location:** ` + "`internal/store/database.go`" + `

#### 2.2 Create user_store.go
- ✅ **Struct:** ` + "`User`" + ` - User data model
- ✅ **Struct:** ` + "`password`" + ` - Password handling
- ✅ **Function:** ` + "`password.Set()`" + ` - Hash passwords
- ✅ **Function:** ` + "`password.Matches()`" + ` - Verify passwords
- ✅ **Function:** ` + "`NewPostgresUserStore()`" + ` - Store factory
- ✅ **Function:** ` + "`CreateUser()`" + ` - Save users
- ✅ **Function:** ` + "`GetUserByUsername()`" + ` - Find users
- 📍 **Location:** ` + "`internal/store/user_store.go`" + `

#### 2.3 Create workout_store.go
- ✅ **Struct:** ` + "`Workout`" + ` - Workout data model
- ✅ **Struct:** ` + "`WorkoutEntry`" + ` - Exercise entry model
- ✅ **Function:** ` + "`NewPostgresWorkoutStore()`" + ` - Store factory
- ✅ **Function:** ` + "`CreateWorkout()`" + ` - Save workouts (with transactions!)
- ✅ **Function:** ` + "`GetWorkoutByID()`" + ` - Find workouts
- ✅ **Function:** ` + "`UpdateWorkout()`" + ` - Modify workouts
- ✅ **Function:** ` + "`DeleteWorkout()`" + ` - Remove workouts
- 📍 **Location:** ` + "`internal/store/workout_store.go`" + `

### STEP 3: Application Layer (Build This Second)
**Why second?** Orchestrates all dependencies.

#### 3.1 Create app.go
- ✅ **Struct:** ` + "`Application`" + ` - Dependency container
- ✅ **Function:** ` + "`NewApplication()`" + ` - **MOST IMPORTANT FUNCTION**
- 📍 **Location:** ` + "`internal/app/app.go`" + `

### STEP 4: Utility Layer (Build This Third)
**Why third?** Handlers need utilities.

#### 4.1 Create utils.go
- ✅ **Function:** ` + "`WriteJSON()`" + ` - JSON responses
- ✅ **Function:** ` + "`ReadIDParam()`" + ` - URL parameters
- 📍 **Location:** ` + "`internal/utils/utils.go`" + `

### STEP 5: API Layer (Build This Fourth)
**Why fourth?** Needs stores and utilities.

#### 5.1 Create user_handler.go
- ✅ **Struct:** ` + "`UserHandler`" + ` - HTTP handler
- ✅ **Function:** ` + "`NewUserHandler()`" + ` - Handler factory
- ✅ **Function:** ` + "`validateRegisterRequest()`" + ` - Input validation
- ✅ **Function:** ` + "`HandleRegisterUser()`" + ` - Registration endpoint
- 📍 **Location:** ` + "`internal/api/user_handler.go`" + `

#### 5.2 Create workout_handler.go
- ✅ **Struct:** ` + "`WorkoutHandler`" + ` - HTTP handler
- ✅ **Function:** ` + "`NewWorkoutHandler()`" + ` - Handler factory
- ✅ **Function:** ` + "`HandleGetWorkoutByID()`" + ` - GET endpoint
- ✅ **Function:** ` + "`HandleCreateWorkout()`" + ` - POST endpoint
- ✅ **Function:** ` + "`HandleUpdateWorkoutByID()`" + ` - PUT endpoint
- ✅ **Function:** ` + "`HandleDeleteWorkoutByID()`" + ` - DELETE endpoint
- 📍 **Location:** ` + "`internal/api/workout_handler.go`" + `

### STEP 6: Routing Layer (Build This Fifth)
**Why fifth?** Connects URLs to handlers.

#### 6.1 Create routes.go
- ✅ **Function:** ` + "`SetupRoutes()`" + ` - Route configuration
- 📍 **Location:** ` + "`internal/routes/routes.go`" + `

### STEP 7: Main Application (Build This Last)
**Why last?** Ties everything together.

#### 7.1 Create main.go
- ✅ **Function:** ` + "`main()`" + ` - Program entry point
- 📍 **Location:** ` + "`Ex10.go`" + ` (or ` + "`cmd/server/main.go`" + `)

## 🧠 KEY INSIGHTS FOR BUILDING

### 1. **Dependency Order Matters**
- Database → Store → App → Utils → Handlers → Routes → Main
- Each layer depends on the previous layers

### 2. **Start with Data Models**
- ` + "`User`" + ` struct first (simpler)
- ` + "`Workout`" + ` struct second (more complex)
- ` + "`WorkoutEntry`" + ` struct third (related data)

### 3. **Build Incrementally**
- Create one function at a time
- Test each function before moving to the next
- Don't build everything at once

### 4. **Follow the Pattern**
For each new feature:
1. **Model** (struct) → **Store** (data access) → **Handler** (HTTP) → **Route** (URL)

## 🚀 QUICK START CHECKLIST

- [ ] **Step 1:** Create project structure
- [ ] **Step 2:** Build ` + "`database.go`" + ` with ` + "`Open()`" + ` and ` + "`MigrateFS()`" + `
- [ ] **Step 3:** Build ` + "`User`" + ` struct and ` + "`password`" + ` handling
- [ ] **Step 4:** Build ` + "`CreateUser()`" + ` and ` + "`GetUserByUsername()`" + `
- [ ] **Step 5:** Build ` + "`Workout`" + ` and ` + "`WorkoutEntry`" + ` structs
- [ ] **Step 6:** Build ` + "`CreateWorkout()`" + ` with transactions
- [ ] **Step 7:** Build ` + "`NewApplication()`" + ` function
- [ ] **Step 8:** Build utility functions (` + "`WriteJSON`" + `, ` + "`ReadIDParam`" + `)
- [ ] **Step 9:** Build user and workout handlers
- [ ] **Step 10:** Build ` + "`SetupRoutes()`" + ` function
- [ ] **Step 11:** Build ` + "`main()`" + ` function
- [ ] **Step 12:** Test the complete application

## 💡 PRO TIPS

1. **Always start with the simplest version** - Add complexity gradually
2. **Test each layer independently** - Don't wait until the end
3. **Follow the dependency flow** - Lower layers first, higher layers last
4. **Use interfaces for testability** - Makes development easier
5. **Copy the patterns** - Once you understand one handler, others follow the same pattern

This order ensures you never get stuck because you're always building on solid foundations! 🎯
`

	path := filepath.Join(outDir, "AIAd_project_building_guide.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// AIAd_WriteAllStructureDiagrams generates all AI advisor structure analysis diagrams
func AIAd_WriteAllStructureDiagrams(outDir string) error {
	fmt.Println("📊 Generating AI advisor function flow analysis...")
	if err := AIAd_WriteFunctionFlowAnalysis(outDir); err != nil {
		return fmt.Errorf("AI advisor function flow analysis failed: %w", err)
	}

	fmt.Println("📊 Generating AI advisor development sequence diagram...")
	if err := AIAd_WriteDevelopmentSequenceDiagram(outDir); err != nil {
		return fmt.Errorf("AI advisor development sequence diagram failed: %w", err)
	}

	fmt.Println("📊 Generating AI advisor execution flow diagram...")
	if err := AIAd_WriteExecutionFlowDiagram(outDir); err != nil {
		return fmt.Errorf("AI advisor execution flow diagram failed: %w", err)
	}

	fmt.Println("📊 Generating AI advisor function dependency diagram...")
	if err := AIAd_WriteFunctionDependencyDiagram(outDir); err != nil {
		return fmt.Errorf("AI advisor function dependency diagram failed: %w", err)
	}

	fmt.Println("📊 Generating AI advisor project building guide...")
	if err := AIAd_WriteProjectBuildingGuide(outDir); err != nil {
		return fmt.Errorf("AI advisor project building guide failed: %w", err)
	}

	return nil
}
