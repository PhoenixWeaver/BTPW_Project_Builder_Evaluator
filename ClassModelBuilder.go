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
CLASS MODEL BUILDER - STEP-BY-STEP PROJECT CREATION TEACHING FUNCTIONS
===============================================================================

Author: AI Teaching Assistant (Generated Content)
Date: 17/09/2025
Description: This file contains comprehensive teaching functions that break down
the entire Go API project creation process step by step, line by line,
function by function, file by file, folder by folder.

These functions help students understand:
1. Complete project structure and organization
2. Step-by-step development workflow
3. File-by-file creation sequence
4. Function-by-function implementation
5. Folder-by-folder organization
6. Dependencies and relationships between components

TO USE THIS FILE:
1. Call ClassModelBuilder_WriteCompleteProjectGuide() for full teaching guide
2. Call ClassModelBuilder_WriteStepByStepWorkflow() for development workflow
3. Call ClassModelBuilder_WriteFileCreationSequence() for file creation order
4. Call ClassModelBuilder_WriteFunctionImplementationGuide() for function details
5. Call ClassModelBuilder_WriteFolderStructureGuide() for directory organization

FEATURES:
- Complete project teaching methodology
- Step-by-step implementation guidance
- File and function creation sequences
- Dependency relationship mapping
- Student-friendly explanations and examples

===============================================================================
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// ClassModelBuilder_WriteCompleteProjectGuide creates a comprehensive teaching guide with fixed syntax
func ClassModelBuilder_WriteCompleteProjectGuide(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph TeachingGuide[\"🎓 COMPLETE PROJECT TEACHING GUIDE\"]\n" +
		"        %% Phase 1: Project Foundation\n" +
		"        subgraph Phase1[\"🏗️ PHASE 1: PROJECT FOUNDATION (15 minutes)\"]\n" +
		"            T1[\"1.1 Create Project Directory<br/>📍 mkdir phoenixflix<br/>🎯 Set up workspace\"]\n" +
		"            T2[\"1.2 Initialize Go Module<br/>📍 go mod init github.com/author/phoenixflix<br/>🎯 Create module\"]\n" +
		"            T3[\"1.3 Create Main Entry Point<br/>📍 touch main.go<br/>🎯 Application entry\"]\n" +
		"            T4[\"1.4 Create Internal Structure<br/>📍 mkdir internal<br/>🎯 Organize packages\"]\n" +
		"        end\n\n" +
		"        %% Phase 2: Application Layer\n" +
		"        subgraph Phase2[\"🏗️ PHASE 2: APPLICATION LAYER (20 minutes)\"]\n" +
		"            T5[\"2.1 Create App Package<br/>📍 mkdir internal/app<br/>🎯 Application logic\"]\n" +
		"            T6[\"2.2 Create App Struct<br/>📍 internal/app/app.go<br/>🎯 Application container\"]\n" +
		"            T7[\"2.3 Create Logger<br/>📍 log.New with timestamps<br/>🎯 Structured logging\"]\n" +
		"            T8[\"2.4 Create Constructor<br/>📍 NewApplication function<br/>🎯 App initialization\"]\n" +
		"        end\n\n" +
		"        %% Phase 3: HTTP Server\n" +
		"        subgraph Phase3[\"🌐 PHASE 3: HTTP SERVER (15 minutes)\"]\n" +
		"            T9[\"3.1 Create HTTP Server<br/>📍 http.Server struct<br/>🎯 Server configuration\"]\n" +
		"            T10[\"3.2 Add Timeouts<br/>📍 ReadTimeout, WriteTimeout<br/>🎯 Server performance\"]\n" +
		"            T11[\"3.3 Create Health Check<br/>📍 HealthCheck handler<br/>🎯 Server monitoring\"]\n" +
		"            T12[\"3.4 Add Command Line Flags<br/>📍 flag package<br/>🎯 Configurable port\"]\n" +
		"        end\n\n" +
		"        %% Phase 4: Routing System\n" +
		"        subgraph Phase4[\"🛣️ PHASE 4: ROUTING SYSTEM (20 minutes)\"]\n" +
		"            T13[\"4.1 Install Chi Router<br/>📍 go get chi/v5<br/>🎯 HTTP routing\"]\n" +
		"            T14[\"4.2 Create Routes Package<br/>📍 mkdir internal/routes<br/>🎯 Route organization\"]\n" +
		"            T15[\"4.3 Create SetupRoutes Function<br/>📍 internal/routes/routes.go<br/>🎯 Route configuration\"]\n" +
		"            T16[\"4.4 Connect Routes to Server<br/>📍 server.Handler = routes<br/>🎯 Route integration\"]\n" +
		"        end\n\n" +
		"        %% Phase 5: API Layer\n" +
		"        subgraph Phase5[\"🌐 PHASE 5: API LAYER (30 minutes)\"]\n" +
		"            T17[\"5.1 Create API Package<br/>📍 mkdir internal/api<br/>🎯 API handlers\"]\n" +
		"            T18[\"5.2 Create Workout Handler<br/>📍 internal/api/workout_handler.go<br/>🎯 CRUD operations\"]\n" +
		"            T19[\"5.3 Create Handler Methods<br/>📍 HandleGetWorkoutByID, HandleCreateWorkout<br/>🎯 HTTP endpoints\"]\n" +
		"            T20[\"5.4 Add Handler to App<br/>📍 app.WorkoutHandler<br/>🎯 Handler integration\"]\n" +
		"        end\n\n" +
		"        %% Phase 6: Database Layer\n" +
		"        subgraph Phase6[\"🗄️ PHASE 6: DATABASE LAYER (45 minutes)\"]\n" +
		"            T21[\"6.1 Create Docker Compose<br/>📍 docker-compose.yml<br/>🎯 PostgreSQL container\"]\n" +
		"            T22[\"6.2 Install pgx Driver<br/>📍 go get pgx/v5<br/>🎯 Database connection\"]\n" +
		"            T23[\"6.3 Create Database Package<br/>📍 mkdir internal/database<br/>🎯 DB management\"]\n" +
		"            T24[\"6.4 Create Connection Function<br/>📍 OpenDatabase function<br/>🎯 DB connection\"]\n" +
		"            T25[\"6.5 Create Migration System<br/>📍 Migrate function<br/>🎯 Schema management\"]\n" +
		"        end\n\n" +
		"        %% Phase 7: Store Layer\n" +
		"        subgraph Phase7[\"💾 PHASE 7: STORE LAYER (40 minutes)\"]\n" +
		"            T26[\"7.1 Create Store Package<br/>📍 mkdir internal/store<br/>🎯 Data access\"]\n" +
		"            T27[\"7.2 Create Workout Store<br/>📍 internal/store/workout_store.go<br/>🎯 CRUD operations\"]\n" +
		"            T28[\"7.3 Implement CRUD Methods<br/>📍 Create, Read, Update, Delete<br/>🎯 Data operations\"]\n" +
		"            T29[\"7.4 Connect Store to Handler<br/>📍 handler uses store<br/>🎯 Data flow\"]\n" +
		"        end\n\n" +
		"        %% Phase 8: Authentication\n" +
		"        subgraph Phase8[\"🔐 PHASE 8: AUTHENTICATION (50 minutes)\"]\n" +
		"            T30[\"8.1 Create User Store<br/>📍 internal/store/user_store.go<br/>🎯 User management\"]\n" +
		"            T31[\"8.2 Create Token Store<br/>📍 internal/store/token_store.go<br/>🎯 JWT tokens\"]\n" +
		"            T32[\"8.3 Create Middleware Package<br/>📍 mkdir internal/middleware<br/>🎯 Request processing\"]\n" +
		"            T33[\"8.4 Implement Auth Middleware<br/>📍 AuthMiddleware function<br/>🎯 Request authentication\"]\n" +
		"            T34[\"8.5 Add JWT Validation<br/>📍 Token validation logic<br/>🎯 Security\"]\n" +
		"        end\n\n" +
		"        %% Phase 9: Testing & Deployment\n" +
		"        subgraph Phase9[\"🧪 PHASE 9: TESTING & DEPLOYMENT (30 minutes)\"]\n" +
		"            T35[\"9.1 Create Test Files<br/>📍 test files<br/>🎯 Unit testing\"]\n" +
		"            T36[\"9.2 Write Integration Tests<br/>📍 API endpoint tests<br/>🎯 Integration testing\"]\n" +
		"            T37[\"9.3 Add Error Handling<br/>📍 Proper error responses<br/>🎯 Error management\"]\n" +
		"            T38[\"9.4 Final Testing<br/>📍 curl commands<br/>🎯 End-to-end testing\"]\n" +
		"        end\n" +
		"    end\n\n" +
		"    %% Teaching sequence connections\n" +
		"    Phase1 --> Phase2\n" +
		"    Phase2 --> Phase3\n" +
		"    Phase3 --> Phase4\n" +
		"    Phase4 --> Phase5\n" +
		"    Phase5 --> Phase6\n" +
		"    Phase6 --> Phase7\n" +
		"    Phase7 --> Phase8\n" +
		"    Phase8 --> Phase9\n\n" +
		"    %% Internal phase connections\n" +
		"    T1 --> T2 --> T3 --> T4\n" +
		"    T5 --> T6 --> T7 --> T8\n" +
		"    T9 --> T10 --> T11 --> T12\n" +
		"    T13 --> T14 --> T15 --> T16\n" +
		"    T17 --> T18 --> T19 --> T20\n" +
		"    T21 --> T22 --> T23 --> T24 --> T25\n" +
		"    T26 --> T27 --> T28 --> T29\n" +
		"    T30 --> T31 --> T32 --> T33 --> T34\n" +
		"    T35 --> T36 --> T37 --> T38\n" +
		"```\n"

	path := filepath.Join(outDir, "ClassModelBuilder_complete_project_guide.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// ClassModelBuilder_WriteStepByStepWorkflow creates a detailed development workflow with fixed syntax
func ClassModelBuilder_WriteStepByStepWorkflow(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Workflow[\"📋 STEP-BY-STEP DEVELOPMENT WORKFLOW\"]\n" +
		"        %% Step 1: Project Setup\n" +
		"        subgraph Step1[\"📁 STEP 1: PROJECT SETUP\"]\n" +
		"            W1[\"1. Create Project Directory<br/>📍 mkdir phoenixflix<br/>🎯 Initialize workspace\"]\n" +
		"            W2[\"2. Initialize Go Module<br/>📍 go mod init project<br/>🎯 Create module file\"]\n" +
		"            W3[\"3. Create Basic Structure<br/>📍 touch main.go<br/>🎯 Entry point\"]\n" +
		"            W4[\"4. Create Internal Directory<br/>📍 mkdir internal<br/>🎯 Package organization\"]\n" +
		"        end\n\n" +
		"        %% Step 2: Application Foundation\n" +
		"        subgraph Step2[\"🏗️ STEP 2: APPLICATION FOUNDATION\"]\n" +
		"            W5[\"5. Create App Package<br/>📍 mkdir internal/app<br/>🎯 Application logic\"]\n" +
		"            W6[\"6. Create App Struct<br/>📍 type Application struct<br/>🎯 App container\"]\n" +
		"            W7[\"7. Create Constructor<br/>📍 func NewApplication<br/>🎯 App initialization\"]\n" +
		"            W8[\"8. Update Main Function<br/>📍 app, err := NewApplication<br/>🎯 App startup\"]\n" +
		"        end\n\n" +
		"        %% Step 3: HTTP Server\n" +
		"        subgraph Step3[\"🌐 STEP 3: HTTP SERVER\"]\n" +
		"            W9[\"9. Create HTTP Server<br/>📍 http.Server struct<br/>🎯 Server config\"]\n" +
		"            W10[\"10. Add Health Check<br/>📍 func HealthCheck<br/>🎯 Server monitoring\"]\n" +
		"            W11[\"11. Add Command Flags<br/>📍 flag package<br/>🎯 Configurable port\"]\n" +
		"            W12[\"12. Start Server<br/>📍 server.ListenAndServe<br/>🎯 Server startup\"]\n" +
		"        end\n\n" +
		"        %% Step 4: Routing System\n" +
		"        subgraph Step4[\"🛣️ STEP 4: ROUTING SYSTEM\"]\n" +
		"            W13[\"13. Install Chi Router<br/>📍 go get chi/v5<br/>🎯 HTTP routing\"]\n" +
		"            W14[\"14. Create Routes Package<br/>📍 mkdir internal/routes<br/>🎯 Route organization\"]\n" +
		"            W15[\"15. Create SetupRoutes<br/>📍 func SetupRoutes<br/>🎯 Route config\"]\n" +
		"            W16[\"16. Connect to Server<br/>📍 server.Handler = routes<br/>🎯 Route integration\"]\n" +
		"        end\n\n" +
		"        %% Step 5: API Handlers\n" +
		"        subgraph Step5[\"🌐 STEP 5: API HANDLERS\"]\n" +
		"            W17[\"17. Create API Package<br/>📍 mkdir internal/api<br/>🎯 API handlers\"]\n" +
		"            W18[\"18. Create Workout Handler<br/>📍 type WorkoutHandler<br/>🎯 CRUD operations\"]\n" +
		"            W19[\"19. Add Handler Methods<br/>📍 HandleGetWorkoutByID<br/>🎯 HTTP endpoints\"]\n" +
		"            W20[\"20. Add to App Struct<br/>📍 app.WorkoutHandler<br/>🎯 Handler integration\"]\n" +
		"        end\n\n" +
		"        %% Step 6: Database Setup\n" +
		"        subgraph Step6[\"🗄️ STEP 6: DATABASE SETUP\"]\n" +
		"            W21[\"21. Create Docker Compose<br/>📍 docker-compose.yml<br/>🎯 PostgreSQL container\"]\n" +
		"            W22[\"22. Install pgx Driver<br/>📍 go get pgx/v5<br/>🎯 Database driver\"]\n" +
		"            W23[\"23. Create Database Package<br/>📍 mkdir internal/database<br/>🎯 DB management\"]\n" +
		"            W24[\"24. Create Connection<br/>📍 func OpenDatabase<br/>🎯 DB connection\"]\n" +
		"        end\n\n" +
		"        %% Step 7: Store Layer\n" +
		"        subgraph Step7[\"💾 STEP 7: STORE LAYER\"]\n" +
		"            W25[\"25. Create Store Package<br/>📍 mkdir internal/store<br/>🎯 Data access\"]\n" +
		"            W26[\"26. Create Workout Store<br/>📍 type WorkoutStore<br/>🎯 Data operations\"]\n" +
		"            W27[\"27. Implement CRUD<br/>📍 Create, Read, Update, Delete<br/>🎯 Data management\"]\n" +
		"            W28[\"28. Connect to Handler<br/>📍 handler uses store<br/>🎯 Data flow\"]\n" +
		"        end\n\n" +
		"        %% Step 8: Authentication\n" +
		"        subgraph Step8[\"🔐 STEP 8: AUTHENTICATION\"]\n" +
		"            W29[\"29. Create User Store<br/>📍 internal/store/user_store.go<br/>🎯 User management\"]\n" +
		"            W30[\"30. Create Token Store<br/>📍 internal/store/token_store.go<br/>🎯 JWT tokens\"]\n" +
		"            W31[\"31. Create Middleware<br/>📍 mkdir internal/middleware<br/>🎯 Request processing\"]\n" +
		"            W32[\"32. Implement Auth<br/>📍 AuthMiddleware function<br/>🎯 Request authentication\"]\n" +
		"        end\n\n" +
		"        %% Step 9: Testing & Polish\n" +
		"        subgraph Step9[\"🧪 STEP 9: TESTING & POLISH\"]\n" +
		"            W33[\"33. Create Tests<br/>📍 test files<br/>🎯 Unit testing\"]\n" +
		"            W34[\"34. Test Endpoints<br/>📍 curl commands<br/>🎯 Integration testing\"]\n" +
		"            W35[\"35. Add Error Handling<br/>📍 Proper error responses<br/>🎯 Error management\"]\n" +
		"            W36[\"36. Final Testing<br/>📍 Complete API testing<br/>🎯 End-to-end validation\"]\n" +
		"        end\n" +
		"    end\n\n" +
		"    %% Workflow sequence connections\n" +
		"    Step1 --> Step2\n" +
		"    Step2 --> Step3\n" +
		"    Step3 --> Step4\n" +
		"    Step4 --> Step5\n" +
		"    Step5 --> Step6\n" +
		"    Step6 --> Step7\n" +
		"    Step7 --> Step8\n" +
		"    Step8 --> Step9\n\n" +
		"    %% Internal step connections\n" +
		"    W1 --> W2 --> W3 --> W4\n" +
		"    W5 --> W6 --> W7 --> W8\n" +
		"    W9 --> W10 --> W11 --> W12\n" +
		"    W13 --> W14 --> W15 --> W16\n" +
		"    W17 --> W18 --> W19 --> W20\n" +
		"    W21 --> W22 --> W23 --> W24\n" +
		"    W25 --> W26 --> W27 --> W28\n" +
		"    W29 --> W30 --> W31 --> W32\n" +
		"    W33 --> W34 --> W35 --> W36\n" +
		"```\n"

	path := filepath.Join(outDir, "ClassModelBuilder_step_by_step_workflow.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// ClassModelBuilder_WriteFileCreationSequence creates a file-by-file creation guide
func ClassModelBuilder_WriteFileCreationSequence(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph FileSequence[\"📁 FILE-BY-FILE CREATION SEQUENCE\"]\n" +
		"        %% Phase 1: Project Foundation Files\n" +
		"        subgraph Files1[\"🏗️ PHASE 1: FOUNDATION FILES\"]\n" +
		"            F1[\"1. main.go<br/>📍 Project root<br/>🎯 Application entry point<br/>📝 Package main, func main()\"]\n" +
		"            F2[\"2. go.mod<br/>📍 Project root<br/>🎯 Module definition<br/>📝 go mod init command\"]\n" +
		"            F3[\"3. .gitignore<br/>📍 Project root<br/>🎯 Version control<br/>📝 Ignore database files\"]\n" +
		"        end\n\n" +
		"        %% Phase 2: Application Layer Files\n" +
		"        subgraph Files2[\"🏗️ PHASE 2: APPLICATION LAYER FILES\"]\n" +
		"            F4[\"4. internal/app/app.go<br/>📍 internal/app/<br/>🎯 Application struct<br/>📝 type Application struct\"]\n" +
		"            F5[\"5. internal/routes/routes.go<br/>📍 internal/routes/<br/>🎯 Route configuration<br/>📝 func SetupRoutes()\"]\n" +
		"        end\n\n" +
		"        %% Phase 3: API Layer Files\n" +
		"        subgraph Files3[\"🌐 PHASE 3: API LAYER FILES\"]\n" +
		"            F6[\"6. internal/api/workout_handler.go<br/>📍 internal/api/<br/>🎯 HTTP handlers<br/>📝 type WorkoutHandler struct\"]\n" +
		"            F7[\"7. internal/api/user_handler.go<br/>📍 internal/api/<br/>🎯 User endpoints<br/>📝 type UserHandler struct\"]\n" +
		"            F8[\"8. internal/api/token_handler.go<br/>📍 internal/api/<br/>🎯 Token endpoints<br/>📝 type TokenHandler struct\"]\n" +
		"        end\n\n" +
		"        %% Phase 4: Database Layer Files\n" +
		"        subgraph Files4[\"🗄️ PHASE 4: DATABASE LAYER FILES\"]\n" +
		"            F9[\"9. docker-compose.yml<br/>📍 Project root<br/>🎯 PostgreSQL container<br/>📝 Docker configuration\"]\n" +
		"            F10[\"10. internal/database/database.go<br/>📍 internal/database/<br/>🎯 DB connection<br/>📝 func OpenDatabase()\"]\n" +
		"            F11[\"11. internal/database/migrate.go<br/>📍 internal/database/<br/>🎯 Schema migration<br/>📝 func Migrate()\"]\n" +
		"        end\n\n" +
		"        %% Phase 5: Store Layer Files\n" +
		"        subgraph Files5[\"💾 PHASE 5: STORE LAYER FILES\"]\n" +
		"            F12[\"12. internal/store/workout_store.go<br/>📍 internal/store/<br/>🎯 Workout CRUD<br/>📝 type WorkoutStore struct\"]\n" +
		"            F13[\"13. internal/store/user_store.go<br/>📍 internal/store/<br/>🎯 User CRUD<br/>📝 type UserStore struct\"]\n" +
		"            F14[\"14. internal/store/token_store.go<br/>📍 internal/store/<br/>🎯 Token CRUD<br/>📝 type TokenStore struct\"]\n" +
		"        end\n\n" +
		"        %% Phase 6: Middleware Files\n" +
		"        subgraph Files6[\"🛡️ PHASE 6: MIDDLEWARE FILES\"]\n" +
		"            F15[\"15. internal/middleware/auth.go<br/>📍 internal/middleware/<br/>🎯 Authentication<br/>📝 func AuthMiddleware()\"]\n" +
		"            F16[\"16. internal/middleware/cors.go<br/>📍 internal/middleware/<br/>🎯 CORS handling<br/>📝 func CORSMiddleware()\"]\n" +
		"            F17[\"17. internal/middleware/ownership.go<br/>📍 internal/middleware/<br/>🎯 Ownership validation<br/>📝 func ValidateOwnership()\"]\n" +
		"        end\n\n" +
		"        %% Phase 7: Test Files\n" +
		"        subgraph Files7[\"🧪 PHASE 7: TEST FILES\"]\n" +
		"            F18[\"18. internal/api/workout_handler_test.go<br/>📍 internal/api/<br/>🎯 Handler tests<br/>📝 func TestWorkoutHandler()\"]\n" +
		"            F19[\"19. internal/store/workout_store_test.go<br/>📍 internal/store/<br/>🎯 Store tests<br/>📝 func TestWorkoutStore()\"]\n" +
		"            F20[\"20. main_test.go<br/>📍 Project root<br/>🎯 Integration tests<br/>📝 func TestMain()\"]\n" +
		"        end\n" +
		"    end\n\n" +
		"    %% File creation sequence connections\n" +
		"    Files1 --> Files2\n" +
		"    Files2 --> Files3\n" +
		"    Files3 --> Files4\n" +
		"    Files4 --> Files5\n" +
		"    Files5 --> Files6\n" +
		"    Files6 --> Files7\n\n" +
		"    %% Internal file connections\n" +
		"    F1 --> F2 --> F3\n" +
		"    F4 --> F5\n" +
		"    F6 --> F7 --> F8\n" +
		"    F9 --> F10 --> F11\n" +
		"    F12 --> F13 --> F14\n" +
		"    F15 --> F16 --> F17\n" +
		"    F18 --> F19 --> F20\n" +
		"```\n"

	path := filepath.Join(outDir, "ClassModelBuilder_file_creation_sequence.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// ClassModelBuilder_WriteFunctionImplementationGuide creates a function-by-function guide
func ClassModelBuilder_WriteFunctionImplementationGuide(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph FunctionGuide[\"⚙️ FUNCTION-BY-FUNCTION IMPLEMENTATION GUIDE\"]\n" +
		"        %% Phase 1: Core Functions\n" +
		"        subgraph Funcs1[\"🏗️ PHASE 1: CORE FUNCTIONS\"]\n" +
		"            FN1[\"1. func main()<br/>📍 main.go<br/>🎯 Application entry point<br/>📝 Initialize app, start server\"]\n" +
		"            FN2[\"2. func NewApplication()<br/>📍 internal/app/app.go<br/>🎯 App constructor<br/>📝 Create logger, return app\"]\n" +
		"            FN3[\"3. func HealthCheck()<br/>📍 internal/app/app.go<br/>🎯 Health endpoint<br/>📝 Return server status\"]\n" +
		"        end\n\n" +
		"        %% Phase 2: Routing Functions\n" +
		"        subgraph Funcs2[\"🛣️ PHASE 2: ROUTING FUNCTIONS\"]\n" +
		"            FN4[\"4. func SetupRoutes()<br/>📍 internal/routes/routes.go<br/>🎯 Route configuration<br/>📝 Create chi router, define routes\"]\n" +
		"            FN5[\"5. func NewWorkoutHandler()<br/>📍 internal/api/workout_handler.go<br/>🎯 Handler constructor<br/>📝 Create handler instance\"]\n" +
		"        end\n\n" +
		"        %% Phase 3: API Handler Functions\n" +
		"        subgraph Funcs3[\"🌐 PHASE 3: API HANDLER FUNCTIONS\"]\n" +
		"            FN6[\"6. func HandleGetWorkoutByID()<br/>📍 internal/api/workout_handler.go<br/>🎯 Get workout endpoint<br/>📝 Extract ID, call store, return data\"]\n" +
		"            FN7[\"7. func HandleCreateWorkout()<br/>📍 internal/api/workout_handler.go<br/>🎯 Create workout endpoint<br/>📝 Parse JSON, validate, call store\"]\n" +
		"            FN8[\"8. func HandleUpdateWorkout()<br/>📍 internal/api/workout_handler.go<br/>🎯 Update workout endpoint<br/>📝 Parse JSON, update store\"]\n" +
		"            FN9[\"9. func HandleDeleteWorkout()<br/>📍 internal/api/workout_handler.go<br/>🎯 Delete workout endpoint<br/>📝 Extract ID, delete from store\"]\n" +
		"        end\n\n" +
		"        %% Phase 4: Database Functions\n" +
		"        subgraph Funcs4[\"🗄️ PHASE 4: DATABASE FUNCTIONS\"]\n" +
		"            FN10[\"10. func OpenDatabase()<br/>📍 internal/database/database.go<br/>🎯 DB connection<br/>📝 Connect to PostgreSQL\"]\n" +
		"            FN11[\"11. func Migrate()<br/>📍 internal/database/migrate.go<br/>🎯 Schema migration<br/>📝 Create tables, indexes\"]\n" +
		"        end\n\n" +
		"        %% Phase 5: Store Functions\n" +
		"        subgraph Funcs5[\"💾 PHASE 5: STORE FUNCTIONS\"]\n" +
		"            FN12[\"12. func NewWorkoutStore()<br/>📍 internal/store/workout_store.go<br/>🎯 Store constructor<br/>📝 Create store instance\"]\n" +
		"            FN13[\"13. func CreateWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Create operation<br/>📝 INSERT INTO workouts\"]\n" +
		"            FN14[\"14. func GetWorkoutByID()<br/>📍 internal/store/workout_store.go<br/>🎯 Read operation<br/>📝 SELECT FROM workouts\"]\n" +
		"            FN15[\"15. func UpdateWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Update operation<br/>📝 UPDATE workouts SET\"]\n" +
		"            FN16[\"16. func DeleteWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Delete operation<br/>📝 DELETE FROM workouts\"]\n" +
		"        end\n\n" +
		"        %% Phase 6: Authentication Functions\n" +
		"        subgraph Funcs6[\"🔐 PHASE 6: AUTHENTICATION FUNCTIONS\"]\n" +
		"            FN17[\"17. func NewUserStore()<br/>📍 internal/store/user_store.go<br/>🎯 User store constructor<br/>📝 Create user store\"]\n" +
		"            FN18[\"18. func CreateUser()<br/>📍 internal/store/user_store.go<br/>🎯 User creation<br/>📝 Hash password, insert user\"]\n" +
		"            FN19[\"19. func GetUserByEmail()<br/>📍 internal/store/user_store.go<br/>🎯 User lookup<br/>📝 SELECT user by email\"]\n" +
		"            FN20[\"20. func NewTokenStore()<br/>📍 internal/store/token_store.go<br/>🎯 Token store constructor<br/>📝 Create token store\"]\n" +
		"            FN21[\"21. func CreateToken()<br/>📍 internal/store/token_store.go<br/>🎯 Token creation<br/>📝 Generate JWT token\"]\n" +
		"            FN22[\"22. func ValidateToken()<br/>📍 internal/store/token_store.go<br/>🎯 Token validation<br/>📝 Verify JWT signature\"]\n" +
		"        end\n\n" +
		"        %% Phase 7: Middleware Functions\n" +
		"        subgraph Funcs7[\"🛡️ PHASE 7: MIDDLEWARE FUNCTIONS\"]\n" +
		"            FN23[\"23. func AuthMiddleware()<br/>📍 internal/middleware/auth.go<br/>🎯 Authentication middleware<br/>📝 Validate JWT token\"]\n" +
		"            FN24[\"24. func CORSMiddleware()<br/>📍 internal/middleware/cors.go<br/>🎯 CORS handling<br/>📝 Set CORS headers\"]\n" +
		"            FN25[\"25. func ValidateOwnership()<br/>📍 internal/middleware/ownership.go<br/>🎯 Ownership validation<br/>📝 Check user ownership\"]\n" +
		"        end\n" +
		"    end\n\n" +
		"    %% Function implementation sequence connections\n" +
		"    Funcs1 --> Funcs2\n" +
		"    Funcs2 --> Funcs3\n" +
		"    Funcs3 --> Funcs4\n" +
		"    Funcs4 --> Funcs5\n" +
		"    Funcs5 --> Funcs6\n" +
		"    Funcs6 --> Funcs7\n\n" +
		"    %% Internal function connections\n" +
		"    FN1 --> FN2 --> FN3\n" +
		"    FN4 --> FN5\n" +
		"    FN6 --> FN7 --> FN8 --> FN9\n" +
		"    FN10 --> FN11\n" +
		"    FN12 --> FN13 --> FN14 --> FN15 --> FN16\n" +
		"    FN17 --> FN18 --> FN19 --> FN20 --> FN21 --> FN22\n" +
		"    FN23 --> FN24 --> FN25\n" +
		"```\n"

	path := filepath.Join(outDir, "ClassModelBuilder_function_implementation_guide.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// ClassModelBuilder_WriteFolderStructureGuide creates a folder-by-folder organization guide
func ClassModelBuilder_WriteFolderStructureGuide(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph FolderGuide[\"📁 FOLDER-BY-FOLDER ORGANIZATION GUIDE\"]\n" +
		"        %% Root Level\n" +
		"        subgraph Root[\"🏠 ROOT LEVEL\"]\n" +
		"            R1[\"phoenixflix/<br/>📍 Project root directory<br/>🎯 Main project folder<br/>📝 Contains: main.go, go.mod, docker-compose.yml\"]\n" +
		"        end\n\n" +
		"        %% Internal Package\n" +
		"        subgraph Internal[\"📦 INTERNAL PACKAGE\"]\n" +
		"            I1[\"internal/<br/>📍 Private package directory<br/>🎯 Application-specific code<br/>📝 Contains: app, api, store, database, middleware\"]\n" +
		"        end\n\n" +
		"        %% Application Layer\n" +
		"        subgraph AppLayer[\"🏗️ APPLICATION LAYER\"]\n" +
		"            A1[\"internal/app/<br/>📍 Application logic<br/>🎯 Core application struct<br/>📝 Contains: app.go (Application struct, NewApplication)\"]\n" +
		"            A2[\"internal/routes/<br/>📍 Route configuration<br/>🎯 HTTP routing setup<br/>📝 Contains: routes.go (SetupRoutes function)\"]\n" +
		"        end\n\n" +
		"        %% API Layer\n" +
		"        subgraph APILayer[\"🌐 API LAYER\"]\n" +
		"            API1[\"internal/api/<br/>📍 HTTP handlers<br/>🎯 API endpoint handlers<br/>📝 Contains: workout_handler.go, user_handler.go, token_handler.go\"]\n" +
		"        end\n\n" +
		"        %% Database Layer\n" +
		"        subgraph DatabaseLayer[\"🗄️ DATABASE LAYER\"]\n" +
		"            D1[\"internal/database/<br/>📍 Database management<br/>🎯 DB connection and migration<br/>📝 Contains: database.go, migrate.go\"]\n" +
		"        end\n\n" +
		"        %% Store Layer\n" +
		"        subgraph StoreLayer[\"💾 STORE LAYER\"]\n" +
		"            S1[\"internal/store/<br/>📍 Data access layer<br/>🎯 CRUD operations<br/>📝 Contains: workout_store.go, user_store.go, token_store.go\"]\n" +
		"        end\n\n" +
		"        %% Middleware Layer\n" +
		"        subgraph MiddlewareLayer[\"🛡️ MIDDLEWARE LAYER\"]\n" +
		"            M1[\"internal/middleware/<br/>📍 Request processing<br/>🎯 Authentication and validation<br/>📝 Contains: auth.go, cors.go, ownership.go\"]\n" +
		"        end\n\n" +
		"        %% Test Files\n" +
		"        subgraph TestFiles[\"🧪 TEST FILES\"]\n" +
		"            T1[\"*_test.go files<br/>📍 Throughout packages<br/>🎯 Unit and integration tests<br/>📝 Contains: handler tests, store tests, integration tests\"]\n" +
		"        end\n\n" +
		"        %% Configuration Files\n" +
		"        subgraph ConfigFiles[\"⚙️ CONFIGURATION FILES\"]\n" +
		"            C1[\"go.mod<br/>📍 Project root<br/>🎯 Module definition<br/>📝 Dependencies and module name\"]\n" +
		"            C2[\"go.sum<br/>📍 Project root<br/>🎯 Dependency checksums<br/>📝 Exact versions and hashes\"]\n" +
		"            C3[\".gitignore<br/>📍 Project root<br/>🎯 Version control<br/>📝 Ignore database and build files\"]\n" +
		"            C4[\"docker-compose.yml<br/>📍 Project root<br/>🎯 Database container<br/>📝 PostgreSQL configuration\"]\n" +
		"        end\n" +
		"    end\n\n" +
		"    %% Folder organization connections\n" +
		"    Root --> Internal\n" +
		"    Internal --> AppLayer\n" +
		"    Internal --> APILayer\n" +
		"    Internal --> DatabaseLayer\n" +
		"    Internal --> StoreLayer\n" +
		"    Internal --> MiddlewareLayer\n" +
		"    Root --> TestFiles\n" +
		"    Root --> ConfigFiles\n\n" +
		"    %% Internal layer connections\n" +
		"    AppLayer --> APILayer\n" +
		"    APILayer --> StoreLayer\n" +
		"    StoreLayer --> DatabaseLayer\n" +
		"    APILayer --> MiddlewareLayer\n\n" +
		"    %% Configuration connections\n" +
		"    C1 --> C2\n" +
		"    C3 --> C4\n" +
		"```\n"

	path := filepath.Join(outDir, "ClassModelBuilder_folder_structure_guide.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// ClassModelBuilder_WriteAllTeachingGuides generates all teaching guides
func ClassModelBuilder_WriteAllTeachingGuides(outDir string) error {
	fmt.Println("📚 Generating Class Model Builder Teaching Guides...")

	// Generate complete project guide (fixed syntax)
	if err := ClassModelBuilder_WriteCompleteProjectGuide(outDir); err != nil {
		return fmt.Errorf("failed to write complete project guide: %w", err)
	}

	// Generate step-by-step workflow (fixed syntax)
	if err := ClassModelBuilder_WriteStepByStepWorkflow(outDir); err != nil {
		return fmt.Errorf("failed to write step-by-step workflow: %w", err)
	}

	// Generate file creation sequence
	if err := ClassModelBuilder_WriteFileCreationSequence(outDir); err != nil {
		return fmt.Errorf("failed to write file creation sequence: %w", err)
	}

	// Generate function implementation guide
	if err := ClassModelBuilder_WriteFunctionImplementationGuide(outDir); err != nil {
		return fmt.Errorf("failed to write function implementation guide: %w", err)
	}

	// Generate folder structure guide
	if err := ClassModelBuilder_WriteFolderStructureGuide(outDir); err != nil {
		return fmt.Errorf("failed to write folder structure guide: %w", err)
	}

	fmt.Println("✅ Class Model Builder teaching guides generated successfully!")
	return nil
}
