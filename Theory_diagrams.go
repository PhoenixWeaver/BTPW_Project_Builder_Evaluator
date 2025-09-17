//go:build flowcharts
// +build flowcharts

/*
===============================================================================
LESSON MODEL DIAGRAMS - INSTRUCTOR'S STEP-BY-STEP TEACHING PROGRESSION
===============================================================================

Author: Ben Tran
Date: 02/09/2025
Description: This file contains lesson-based diagram functions that follow the
             instructor's teaching progression from IntructorProjectBuilderModel.txt.
             These functions generate diagrams showing the exact step-by-step
             approach the instructor used to build the project, making it easy
             to understand the learning progression and build order.

TO USE THIS FILE:
1. Call LessonModel_WriteAllLessonDiagrams() to generate all lesson diagrams
2. Each function generates specific lesson-based guidance
3. Diagrams are saved as LessonModel_*.mmd.md files for easy identification

LESSON MODEL OBJECTIVES:
- Follow instructor's exact teaching progression
- Show step-by-step learning order
- Demonstrate proper build sequence
- Provide educational scaffolding

FEATURES:
- LessonModel_instructor_progression.mmd.md - Complete instructor teaching flow
- LessonModel_build_sequence.mmd.md - Step-by-step build order
- LessonModel_learning_phases.mmd.md - Learning phases and milestones
- LessonModel_project_scaffolding.mmd.md - How to scaffold from scratch

===============================================================================
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// LessonModel_WriteAllLessonDiagrams generates all lesson-based diagrams following instructor's progression
func LessonModel_WriteAllLessonDiagrams(outDir string) error {
	fmt.Println("🎓 Generating Lesson Model Diagrams (Instructor's Teaching Progression)...")

	// Generate instructor progression diagram
	if err := LessonModel_WriteInstructorProgressionDiagram(outDir); err != nil {
		return fmt.Errorf("failed to write instructor progression diagram: %w", err)
	}

	// Generate build sequence diagram
	if err := LessonModel_WriteBuildSequenceDiagram(outDir); err != nil {
		return fmt.Errorf("failed to write build sequence diagram: %w", err)
	}

	// Generate learning phases diagram
	if err := LessonModel_WriteLearningPhasesDiagram(outDir); err != nil {
		return fmt.Errorf("failed to write learning phases diagram: %w", err)
	}

	// Generate project scaffolding diagram
	if err := LessonModel_WriteProjectScaffoldingDiagram(outDir); err != nil {
		return fmt.Errorf("failed to write project scaffolding diagram: %w", err)
	}

	fmt.Println("�� Lesson Model diagrams generated successfully!")
	return nil
}

// Theory_WriteProjectOGDiagrams now calls the lesson model diagrams
func Theory_WriteProjectOGDiagrams(outDir string, structure *ProjectStructure) error {
	fmt.Println("🎓 Generating Lesson Model Diagrams based on Instructor's Teaching Progression...")
	return LessonModel_WriteAllLessonDiagrams(outDir)
}

// LessonModel_WriteInstructorProgressionDiagram creates a diagram showing the instructor's complete teaching progression
func LessonModel_WriteInstructorProgressionDiagram(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Phase1[\"��️ PHASE 1: PROJECT SCAFFOLDING (42m 33s)\"]\n" +
		"        P1A[\"Creating the Go Project<br/>01:56:28 - 02:07:25<br/>📁 Initialize project structure\"]\n" +
		"        P1B[\"Creating an HTTP Server<br/>02:07:26 - 02:14:10<br/>🌐 Basic server setup\"]\n" +
		"        P1C[\"Parsing Command-Line Flags<br/>02:14:11 - 02:18:47<br/>⚙️ Configuration management\"]\n" +
		"        P1D[\"Chi Router<br/>02:18:48 - 02:27:06<br/>🛣️ HTTP routing setup\"]\n" +
		"        P1E[\"API Route Handlers<br/>02:27:07 - 02:39:05<br/>📡 Basic API endpoints\"]\n" +
		"    end\n\n" +
		"    subgraph Phase2[\"🗄️ PHASE 2: DATA LAYER (1h 35s)\"]\n" +
		"        P2A[\"Postgres Database Docker Container<br/>02:39:06 - 02:46:17<br/>🐳 Database setup\"]\n" +
		"        P2B[\"pgx Driver for PostgreSQL<br/>02:46:18 - 02:55:48<br/>🔌 Database connection\"]\n" +
		"        P2C[\"SQL Migrations with Goose<br/>02:55:49 - 03:13:08<br/>�� Schema management\"]\n" +
		"        P2D[\"Running Goose Migrations<br/>03:13:09 - 03:20:39<br/>▶️ Apply migrations\"]\n" +
		"        P2E[\"Defining Data Types in Store<br/>03:20:40 - 03:29:12<br/>📊 Data models\"]\n" +
		"        P2F[\"CreateWorkout Query<br/>03:29:13 - 03:39:46<br/>💾 First database operation\"]\n" +
		"    end\n\n" +
		"    subgraph Phase3[\"�� PHASE 3: API CRUD ROUTES (1h 24m 15s)\"]\n" +
		"        P3A[\"CreateWorkout Handler<br/>03:39:47 - 03:49:30<br/>➕ Create functionality\"]\n" +
		"        P3B[\"Testing CreateWorkout Endpoint with cURL<br/>03:49:31 - 03:55:26<br/>🧪 API testing\"]\n" +
		"        P3C[\"Getting Workouts By ID<br/>03:55:27 - 04:05:25<br/>🔍 Read functionality\"]\n" +
		"        P3D[\"Updating Workouts<br/>04:05:26 - 04:17:06<br/>✏️ Update functionality\"]\n" +
		"        P3E[\"Handlers for Getting & Updating Workouts<br/>04:17:07 - 04:33:34<br/>🔄 Complete CRUD\"]\n" +
		"        P3F[\"Deleting Workouts<br/>04:33:35 - 04:42:53<br/>��️ Delete functionality\"]\n" +
		"        P3G[\"JSON Response Writer Refactor<br/>04:42:54 - 04:49:04<br/>♻️ Code improvement\"]\n" +
		"        P3H[\"Logging & JSON Error Responses<br/>04:49:05 - 05:04:09<br/>📝 Error handling\"]\n" +
		"    end\n\n" +
		"    subgraph Phase4[\"🧪 PHASE 4: TESTING GO APPLICATIONS (38m 20s)\"]\n" +
		"        P4A[\"Using a Testing Database<br/>05:04:10 - 05:12:26<br/>��️ Test environment\"]\n" +
		"        P4B[\"Connecting to the Test Database<br/>05:12:27 - 05:19:50<br/>�� Test connections\"]\n" +
		"        P4C[\"Testing CreateWorkout Success<br/>05:19:51 - 05:26:21<br/>✅ Success tests\"]\n" +
		"        P4D[\"Testing CreateWorkout Errors<br/>05:26:22 - 05:35:34<br/>❌ Error tests\"]\n" +
		"        P4E[\"Running Tests in Go<br/>05:35:35 - 05:42:34<br/>🏃 Test execution\"]\n" +
		"    end\n\n" +
		"    subgraph Phase5[\"�� PHASE 5: AUTHENTICATION (1h 20m 4s)\"]\n" +
		"        P5A[\"Managing User Data<br/>05:42:35 - 05:48:25<br/>�� User management\"]\n" +
		"        P5B[\"User SQL Queries<br/>05:48:26 - 05:56:36<br/>💾 User database ops\"]\n" +
		"        P5C[\"Validating User Data<br/>05:56:37 - 06:04:55<br/>✅ Data validation\"]\n" +
		"        P5D[\"Register User API<br/>06:04:56 - 06:09:55<br/>�� User registration\"]\n" +
		"        P5E[\"Hashing & Storing User Passwords<br/>06:09:56 - 06:21:57<br/>🔒 Password security\"]\n" +
		"        P5F[\"Token Authentication & OAuth 2.0<br/>06:21:58 - 06:29:04<br/>🎫 Token system\"]\n" +
		"        P5G[\"Creating a Tokens Table<br/>06:29:05 - 06:33:31<br/>🗄️ Token storage\"]\n" +
		"        P5H[\"Generating JSON Web Tokens<br/>06:33:32 - 06:49:19<br/>🔑 JWT creation\"]\n" +
		"        P5I[\"Token API Handlers<br/>06:49:20 - 07:00:13<br/>📡 Token endpoints\"]\n" +
		"        P5J[\"Testing the Authentication Routes<br/>07:00:14 - 07:02:48<br/>�� Auth testing\"]\n" +
		"    end\n\n" +
		"    subgraph Phase6[\"🛡️ PHASE 6: MIDDLEWARE (58m 44s)\"]\n" +
		"        P6A[\"Getting User Tokens<br/>07:02:49 - 07:10:57<br/>🎫 Token retrieval\"]\n" +
		"        P6B[\"Modifying Request Context<br/>07:10:58 - 07:18:27<br/>📝 Context management\"]\n" +
		"        P6C[\"Authentication Middleware<br/>07:18:28 - 07:25:24<br/>🔐 Auth middleware\"]\n" +
		"        P6D[\"Protecting Routes with Middleware<br/>07:25:25 - 07:36:31<br/>🛡️ Route protection\"]\n" +
		"        P6E[\"Adding User ID Migration<br/>07:36:32 - 07:42:19<br/>📋 Schema update\"]\n" +
		"        P6F[\"Validating User Workout Ownership<br/>07:42:20 - 07:53:16<br/>✅ Ownership validation\"]\n" +
		"        P6G[\"Testing API Endpoints<br/>07:53:17 - 08:01:39<br/>🧪 Final testing\"]\n" +
		"    end\n\n" +
		"    subgraph Phase7[\"🎯 PHASE 7: WRAPPING UP (1m 54s)\"]\n" +
		"        P7A[\"Wrapping Up<br/>Final review and completion\"]\n" +
		"    end\n\n" +
		"    %% Phase connections\n" +
		"    Phase1 --> Phase2\n" +
		"    Phase2 --> Phase3\n" +
		"    Phase3 --> Phase4\n" +
		"    Phase4 --> Phase5\n" +
		"    Phase5 --> Phase6\n" +
		"    Phase6 --> Phase7\n\n" +
		"    %% Internal phase connections\n" +
		"    P1A --> P1B --> P1C --> P1D --> P1E\n" +
		"    P2A --> P2B --> P2C --> P2D --> P2E --> P2F\n" +
		"    P3A --> P3B --> P3C --> P3D --> P3E --> P3F --> P3G --> P3H\n" +
		"    P4A --> P4B --> P4C --> P4D --> P4E\n" +
		"    P5A --> P5B --> P5C --> P5D --> P5E --> P5F --> P5G --> P5H --> P5I --> P5J\n" +
		"    P6A --> P6B --> P6C --> P6D --> P6E --> P6F --> P6G\n" +
		"    P6G --> P7A\n" +
		"```\n"

	path := filepath.Join(outDir, "LessonModel_instructor_progression.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// LessonModel_WriteBuildSequenceDiagram creates a diagram showing the optimal build sequence
func LessonModel_WriteBuildSequenceDiagram(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Start[\"🚀 START HERE: Project Foundation\"]\n" +
		"        S1[\"1. Create Go Project<br/>�� go mod init<br/>📁 Basic structure\"]\n" +
		"        S2[\"2. HTTP Server<br/>�� Basic server<br/>�� Listen on port\"]\n" +
		"        S3[\"3. Command Line Flags<br/>⚙️ Configuration<br/>⚙️ Port, host settings\"]\n" +
		"        S4[\"4. Chi Router<br/>��️ HTTP routing<br/>��️ Route definitions\"]\n" +
		"        S5[\"5. Basic API Handlers<br/>�� Hello world endpoints<br/>📡 Health check\"]\n" +
		"    end\n\n" +
		"    subgraph DataLayer[\"🗄️ DATA LAYER: Database Foundation\"]\n" +
		"        D1[\"6. Docker Database<br/>🐳 PostgreSQL container<br/>🐳 docker-compose.yml\"]\n" +
		"        D2[\"7. Database Driver<br/>�� pgx driver<br/>�� Connection setup\"]\n" +
		"        D3[\"8. Migration System<br/>�� Goose migrations<br/>📋 Schema versioning\"]\n" +
		"        D4[\"9. Run Migrations<br/>▶️ Apply schema<br/>▶️ Create tables\"]\n" +
		"        D5[\"10. Data Types<br/>�� Go structs<br/>�� Database models\"]\n" +
		"        D6[\"11. First Query<br/>💾 Create operation<br/>💾 Database interaction\"]\n" +
		"    end\n\n" +
		"    subgraph CRUD[\"�� CRUD OPERATIONS: Core Functionality\"]\n" +
		"        C1[\"12. Create Handler<br/>➕ POST endpoint<br/>➕ Data creation\"]\n" +
		"        C2[\"13. Test Create<br/>🧪 cURL testing<br/>�� API validation\"]\n" +
		"        C3[\"14. Read Handler<br/>🔍 GET by ID<br/>�� Data retrieval\"]\n" +
		"        C4[\"15. Update Handler<br/>✏️ PUT/PATCH<br/>✏️ Data modification\"]\n" +
		"        C5[\"16. Delete Handler<br/>🗑️ DELETE<br/>🗑️ Data removal\"]\n" +
		"        C6[\"17. Response Refactor<br/>♻️ JSON responses<br/>♻️ Error handling\"]\n" +
		"    end\n\n" +
		"    subgraph Testing[\"🧪 TESTING: Quality Assurance\"]\n" +
		"        T1[\"18. Test Database<br/>🗄️ Separate test DB<br/>🗄️ Test environment\"]\n" +
		"        T2[\"19. Test Connection<br/>�� Test DB setup<br/>🔗 Test isolation\"]\n" +
		"        T3[\"20. Success Tests<br/>✅ Happy path tests<br/>✅ Expected behavior\"]\n" +
		"        T4[\"21. Error Tests<br/>❌ Error scenarios<br/>❌ Edge cases\"]\n" +
		"        T5[\"22. Test Execution<br/>🏃 go test<br/> Automated testing\"]\n" +
		"    end\n\n" +
		"    subgraph Auth[\"🔐 AUTHENTICATION: Security Layer\"]\n" +
		"        A1[\"23. User Management<br/>�� User data model<br/>👤 User operations\"]\n" +
		"        A2[\"24. User Queries<br/>💾 User database ops<br/>💾 User CRUD\"]\n" +
		"        A3[\"25. Data Validation<br/>✅ Input validation<br/>✅ Data integrity\"]\n" +
		"        A4[\"26. Registration API<br/>�� User signup<br/>📝 Account creation\"]\n" +
		"        A5[\"27. Password Security<br/>�� Password hashing<br/>�� Secure storage\"]\n" +
		"        A6[\"28. Token System<br/>�� JWT tokens<br/>🎫 Authentication\"]\n" +
		"        A7[\"29. Token Storage<br/>🗄️ Token database<br/>��️ Token management\"]\n" +
		"        A8[\"30. Token Handlers<br/>�� Login endpoints<br/>📡 Token generation\"]\n" +
		"    end\n\n" +
		"    subgraph Middleware[\"🛡️ MIDDLEWARE: Request Processing\"]\n" +
		"        M1[\"31. Token Retrieval<br/>🎫 Extract tokens<br/>🎫 Request parsing\"]\n" +
		"        M2[\"32. Context Management<br/>�� Request context<br/>📝 User context\"]\n" +
		"        M3[\"33. Auth Middleware<br/>�� Authentication<br/>🔐 Token validation\"]\n" +
		"        M4[\"34. Route Protection<br/>🛡️ Protected endpoints<br/>🛡️ Access control\"]\n" +
		"        M5[\"35. Ownership Validation<br/>✅ Resource ownership<br/>✅ User permissions\"]\n" +
		"        M6[\"36. Final Testing<br/>�� Complete testing<br/>🧪 End-to-end validation\"]\n" +
		"    end\n\n" +
		"    %% Build sequence connections\n" +
		"    Start --> DataLayer\n" +
		"    DataLayer --> CRUD\n" +
		"    CRUD --> Testing\n" +
		"    Testing --> Auth\n" +
		"    Auth --> Middleware\n\n" +
		"    %% Internal connections\n" +
		"    S1 --> S2 --> S3 --> S4 --> S5\n" +
		"    D1 --> D2 --> D3 --> D4 --> D5 --> D6\n" +
		"    C1 --> C2 --> C3 --> C4 --> C5 --> C6\n" +
		"    T1 --> T2 --> T3 --> T4 --> T5\n" +
		"    A1 --> A2 --> A3 --> A4 --> A5 --> A6 --> A7 --> A8\n" +
		"    M1 --> M2 --> M3 --> M4 --> M5 --> M6\n" +
		"```\n"

	path := filepath.Join(outDir, "LessonModel_build_sequence.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// LessonModel_WriteLearningPhasesDiagram creates a diagram showing the learning phases and milestones
func LessonModel_WriteLearningPhasesDiagram(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Phase1[\"�� LEARNING PHASE 1: FOUNDATION (42m 33s)\"]\n" +
		"        L1A[\"📚 Concepts Learned:<br/>• Go project structure<br/>• HTTP server basics<br/>• Command-line configuration<br/>• HTTP routing with Chi<br/>• Basic API design\"]\n" +
		"        L1B[\"�� Milestone:<br/>Working HTTP server<br/>with basic routing\"]\n" +
		"        L1C[\"🔧 Skills Developed:<br/>• Project initialization<br/>• Server configuration<br/>• Route handling<br/>• API endpoint creation\"]\n" +
		"    end\n\n" +
		"    subgraph Phase2[\"🎯 LEARNING PHASE 2: DATA PERSISTENCE (1h 35s)\"]\n" +
		"        L2A[\"📚 Concepts Learned:<br/>• Docker containerization<br/>• PostgreSQL database<br/>• Database drivers (pgx)<br/>• Migration systems (Goose)<br/>• Data modeling<br/>• SQL operations\"]\n" +
		"        L2B[\"🎯 Milestone:<br/>Database-connected<br/>application with<br/>data persistence\"]\n" +
		"        L2C[\"🔧 Skills Developed:<br/>• Database setup<br/>• Migration management<br/>• Data type definition<br/>• CRUD operations\"]\n" +
		"    end\n\n" +
		"    subgraph Phase3[\"🎯 LEARNING PHASE 3: API DEVELOPMENT (1h 24m 15s)\"]\n" +
		"        L3A[\"📚 Concepts Learned:<br/>• RESTful API design<br/>• HTTP methods (GET, POST, PUT, DELETE)<br/>• JSON serialization<br/>• Error handling<br/>• API testing with cURL<br/>• Response formatting\"]\n" +
		"        L3B[\"🎯 Milestone:<br/>Complete CRUD API<br/>with proper error handling\"]\n" +
		"        L3C[\"🔧 Skills Developed:<br/>• API endpoint creation<br/>• HTTP method handling<br/>• JSON processing<br/>• Error response design<br/>• API testing\"]\n" +
		"    end\n\n" +
		"    subgraph Phase4[\"�� LEARNING PHASE 4: TESTING (38m 20s)\"]\n" +
		"        L4A[\"📚 Concepts Learned:<br/>• Test-driven development<br/>• Test database setup<br/>• Unit testing in Go<br/>• Test isolation<br/>• Success and error testing<br/>• Test execution\"]\n" +
		"        L4B[\"🎯 Milestone:<br/>Fully tested application<br/>with automated tests\"]\n" +
		"        L4C[\"🔧 Skills Developed:<br/>• Test writing<br/>• Test environment setup<br/>• Test execution<br/>• Quality assurance\"]\n" +
		"    end\n\n" +
		"    subgraph Phase5[\"�� LEARNING PHASE 5: SECURITY (1h 20m 4s)\"]\n" +
		"        L5A[\"📚 Concepts Learned:<br/>• User authentication<br/>• Password security<br/>• JWT tokens<br/>• OAuth 2.0 concepts<br/>• Token-based auth<br/>• User management\"]\n" +
		"        L5B[\"🎯 Milestone:<br/>Secure application<br/>with user authentication\"]\n" +
		"        L5C[\"🔧 Skills Developed:<br/>• Authentication systems<br/>• Password hashing<br/>• Token management<br/>• Security best practices\"]\n" +
		"    end\n\n" +
		"    subgraph Phase6[\"�� LEARNING PHASE 6: MIDDLEWARE (58m 44s)\"]\n" +
		"        L6A[\"📚 Concepts Learned:<br/>• Middleware patterns<br/>• Request processing<br/>• Context management<br/>• Route protection<br/>• Authorization<br/>• Resource ownership\"]\n" +
		"        L6B[\"🎯 Milestone:<br/>Production-ready<br/>application with<br/>complete security\"]\n" +
		"        L6C[\"�� Skills Developed:<br/>• Middleware development<br/>• Request interception<br/>• Access control<br/>• Authorization logic\"]\n" +
		"    end\n\n" +
		"    %% Learning progression\n" +
		"    Phase1 --> Phase2\n" +
		"    Phase2 --> Phase3\n" +
		"    Phase3 --> Phase4\n" +
		"    Phase4 --> Phase5\n" +
		"    Phase5 --> Phase6\n\n" +
		"    %% Internal connections\n" +
		"    L1A --> L1B --> L1C\n" +
		"    L2A --> L2B --> L2C\n" +
		"    L3A --> L3B --> L3C\n" +
		"    L4A --> L4B --> L4C\n" +
		"    L5A --> L5B --> L5C\n" +
		"    L6A --> L6B --> L6C\n" +
		"```\n"

	path := filepath.Join(outDir, "LessonModel_learning_phases.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// LessonModel_WriteProjectScaffoldingDiagram creates a diagram showing how to scaffold the project from scratch
func LessonModel_WriteProjectScaffoldingDiagram(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph Setup[\"�� PROJECT SCAFFOLDING: Start from Zero\"]\n" +
		"        S1[\"Step 1: Initialize Project<br/>📁 mkdir workout-api<br/>📁 cd workout-api<br/>📁 go mod init workout-api\"]\n" +
		"        S2[\"Step 2: Create Basic Structure<br/>📁 mkdir internal<br/>📁 mkdir internal/app<br/>📁 mkdir internal/api<br/>📁 mkdir internal/store\"]\n" +
		"        S3[\"Step 3: Main Application<br/>📄 Create main.go<br/>📄 Basic HTTP server<br/>📄 Listen on port 8080\"]\n" +
		"        S4[\"Step 4: Configuration<br/>📄 Add command-line flags<br/>📄 Port configuration<br/>�� Environment variables\"]\n" +
		"        S5[\"Step 5: Router Setup<br/>📄 Add Chi router<br/>�� Basic routes<br/>📄 Health check endpoint\"]\n" +
		"    end\n\n" +
		"    subgraph Database[\"🗄️ DATABASE SCAFFOLDING: Data Layer\"]\n" +
		"        D1[\"Step 6: Docker Setup<br/>📄 Create docker-compose.yml<br/>�� PostgreSQL service<br/>📄 Database configuration\"]\n" +
		"        D2[\"Step 7: Database Driver<br/>📄 Add pgx dependency<br/>📄 Connection setup<br/>📄 Database interface\"]\n" +
		"        D3[\"Step 8: Migrations<br/>📁 Create migrations folder<br/>�� Add Goose dependency<br/>📄 First migration file\"]\n" +
		"        D4[\"Step 9: Data Models<br/>�� Define Go structs<br/>�� Database types<br/>📄 JSON tags\"]\n" +
		"        D5[\"Step 10: Store Layer<br/>�� Create store interface<br/>�� Implement store methods<br/>📄 Database operations\"]\n" +
		"    end\n\n" +
		"    subgraph API[\"📡 API SCAFFOLDING: Endpoints\"]\n" +
		"        A1[\"Step 11: API Handlers<br/>📄 Create handler functions<br/>�� HTTP method handling<br/>�� Request/response processing\"]\n" +
		"        A2[\"Step 12: Routes<br/>📄 Define API routes<br/>�� Route handlers<br/>📄 Middleware setup\"]\n" +
		"        A3[\"Step 13: JSON Processing<br/>📄 Request parsing<br/>�� Response formatting<br/>�� Error handling\"]\n" +
		"        A4[\"Step 14: Testing<br/>�� cURL commands<br/>📄 API testing<br/>📄 Endpoint validation\"]\n" +
		"    end\n\n" +
		"    subgraph Testing[\"�� TESTING SCAFFOLDING: Quality Assurance\"]\n" +
		"        T1[\"Step 15: Test Setup<br/>📄 Test database<br/>📄 Test configuration<br/>📄 Test isolation\"]\n" +
		"        T2[\"Step 16: Unit Tests<br/>📄 Test functions<br/>📄 Success scenarios<br/>📄 Error scenarios\"]\n" +
		"        T3[\"Step 17: Test Execution<br/>📄 go test commands<br/>�� Test coverage<br/>�� Continuous testing\"]\n" +
		"    end\n\n" +
		"    subgraph Security[\"�� SECURITY SCAFFOLDING: Authentication\"]\n" +
		"        Sec1[\"Step 18: User Model<br/>�� User struct<br/>�� User database<br/>📄 User operations\"]\n" +
		"        Sec2[\"Step 19: Authentication<br/>�� Password hashing<br/>📄 JWT tokens<br/>📄 Login system\"]\n" +
		"        Sec3[\"Step 20: Authorization<br/>📄 Middleware<br/>📄 Route protection<br/>📄 User permissions\"]\n" +
		"    end\n\n" +
		"    %% Scaffolding flow\n" +
		"    Setup --> Database\n" +
		"    Database --> API\n" +
		"    API --> Testing\n" +
		"    Testing --> Security\n\n" +
		"    %% Internal connections\n" +
		"    S1 --> S2 --> S3 --> S4 --> S5\n" +
		"    D1 --> D2 --> D3 --> D4 --> D5\n" +
		"    A1 --> A2 --> A3 --> A4\n" +
		"    T1 --> T2 --> T3\n" +
		"    Sec1 --> Sec2 --> Sec3\n" +
		"```\n"

	path := filepath.Join(outDir, "LessonModel_project_scaffolding.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}
