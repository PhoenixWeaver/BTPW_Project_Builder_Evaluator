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
AI ADVISOR CREATION & EXECUTION FUNCTIONS - FUNCTION CREATION AND EXECUTION ORDER DIAGRAMS
===============================================================================

Author: AI Advisor (Generated Content)
Date: 17/09/2025
Description: This file contains AI-generated functions that create flowcharts showing:
1. Function creation order (which function was created first > second > third > current)
2. Function execution order (which function should start first > second > third > current)

These diagrams help understand the development sequence and runtime execution flow
of the project functions based on internal directory structure.

TO USE THIS FILE:
1. Call AIAdCreate_Exe_WriteFunctionCreationOrder() for creation sequence
2. Call AIAdCreate_Exe_WriteFunctionExecutionOrder() for execution sequence
3. Diagrams are saved as AIAdCreate_Exe_*.mmd.md files

FEATURES:
- AIAdCreate_Exe_function_creation_order.mmd.md - Shows development sequence
- AIAdCreate_Exe_function_execution_order.mmd.md - Shows runtime execution flow

===============================================================================
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// AIAdCreate_Exe_WriteFunctionCreationOrder creates a flowchart showing the order functions were created
func AIAdCreate_Exe_WriteFunctionCreationOrder(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph CreationOrder[\"📅 FUNCTION CREATION ORDER (Development Sequence)\"]\n" +
		"        %% Phase 1: Foundation Functions (Created First)\n" +
		"        subgraph Phase1[\"🏗️ PHASE 1: FOUNDATION (Created First)\"]\n" +
		"            F1[\"1. main()<br/>📍 main.go<br/>🎯 Application entry point\"]\n" +
		"            F2[\"2. OpenDatabase()<br/>📍 internal/database/database.go<br/>🎯 Database connection\"]\n" +
		"            F3[\"3. Migrate()<br/>📍 internal/database/migrate.go<br/>🎯 Schema migration\"]\n" +
		"        end\n\n" +
		"        %% Phase 2: Store Layer Functions (Created Second)\n" +
		"        subgraph Phase2[\"💾 PHASE 2: STORE LAYER (Created Second)\"]\n" +
		"            F4[\"4. NewPostgresWorkoutStore()<br/>📍 internal/store/workout_store.go<br/>🎯 Store constructor\"]\n" +
		"            F5[\"5. CreateWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Create operation\"]\n" +
		"            F6[\"6. GetWorkoutByID()<br/>📍 internal/store/workout_store.go<br/>🎯 Read operation\"]\n" +
		"            F7[\"7. UpdateWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Update operation\"]\n" +
		"            F8[\"8. DeleteWorkout()<br/>📍 internal/store/workout_store.go<br/>🎯 Delete operation\"]\n" +
		"        end\n\n" +
		"        %% Phase 3: Token Layer Functions (Created Third)\n" +
		"        subgraph Phase3[\"🔑 PHASE 3: TOKEN LAYER (Created Third)\"]\n" +
		"            F9[\"9. NewTokenStore()<br/>📍 internal/tokens/token_store.go<br/>🎯 Token store constructor\"]\n" +
		"            F10[\"10. CreateToken()<br/>📍 internal/tokens/token_store.go<br/>🎯 Token creation\"]\n" +
		"            F11[\"11. ValidateToken()<br/>📍 internal/tokens/token_store.go<br/>🎯 Token validation\"]\n" +
		"            F12[\"12. RefreshToken()<br/>📍 internal/tokens/token_store.go<br/>🎯 Token refresh\"]\n" +
		"        end\n\n" +
		"        %% Phase 4: User Store Functions (Created Fourth)\n" +
		"        subgraph Phase4[\"👤 PHASE 4: USER STORE (Created Fourth)\"]\n" +
		"            F13[\"13. NewUserStore()<br/>📍 internal/store/user_store.go<br/>🎯 User store constructor\"]\n" +
		"            F14[\"14. CreateUser()<br/>📍 internal/store/user_store.go<br/>🎯 User creation\"]\n" +
		"            F15[\"15. GetUserByEmail()<br/>📍 internal/store/user_store.go<br/>🎯 User lookup\"]\n" +
		"            F16[\"16. UpdateUser()<br/>📍 internal/store/user_store.go<br/>🎯 User update\"]\n" +
		"        end\n\n" +
		"        %% Phase 5: Middleware Functions (Created Fifth)\n" +
		"        subgraph Phase5[\"🛡️ PHASE 5: MIDDLEWARE (Created Fifth)\"]\n" +
		"            F17[\"17. AuthMiddleware()<br/>📍 internal/middleware/auth.go<br/>🎯 Authentication middleware\"]\n" +
		"            F18[\"18. ValidateUserOwnership()<br/>📍 internal/middleware/ownership.go<br/>🎯 Ownership validation\"]\n" +
		"            F19[\"19. CORSMiddleware()<br/>📍 internal/middleware/cors.go<br/>🎯 CORS handling\"]\n" +
		"        end\n\n" +
		"        %% Phase 6: API Layer Functions (Created Sixth)\n" +
		"        subgraph Phase6[\"🌐 PHASE 6: API LAYER (Created Sixth)\"]\n" +
		"            F20[\"20. NewWorkoutHandler()<br/>📍 internal/api/workout_handler.go<br/>🎯 Handler constructor\"]\n" +
		"            F21[\"21. HandleCreateWorkout()<br/>📍 internal/api/workout_handler.go<br/>🎯 Create endpoint\"]\n" +
		"            F22[\"22. HandleGetWorkout()<br/>📍 internal/api/workout_handler.go<br/>🎯 Read endpoint\"]\n" +
		"            F23[\"23. HandleUpdateWorkout()<br/>📍 internal/api/workout_handler.go<br/>🎯 Update endpoint\"]\n" +
		"            F24[\"24. HandleDeleteWorkout()<br/>📍 internal/api/workout_handler.go<br/>🎯 Delete endpoint\"]\n" +
		"        end\n\n" +
		"        %% Phase 7: Application Layer (Created Last)\n" +
		"        subgraph Phase7[\"🚀 PHASE 7: APPLICATION LAYER (Created Last)\"]\n" +
		"            F25[\"25. NewApplication()<br/>📍 internal/app/app.go<br/>🎯 Application constructor\"]\n" +
		"            F26[\"26. SetupRoutes()<br/>📍 internal/app/routes.go<br/>🎯 Route configuration\"]\n" +
		"            F27[\"27. StartServer()<br/>📍 internal/app/server.go<br/>🎯 Server startup\"]\n" +
		"        end\n" +
		"    end\n\n" +
		"    %% Creation sequence connections\n" +
		"    Phase1 --> Phase2\n" +
		"    Phase2 --> Phase3\n" +
		"    Phase3 --> Phase4\n" +
		"    Phase4 --> Phase5\n" +
		"    Phase5 --> Phase6\n" +
		"    Phase6 --> Phase7\n\n" +
		"    %% Internal phase connections\n" +
		"    F1 --> F2 --> F3\n" +
		"    F4 --> F5 --> F6 --> F7 --> F8\n" +
		"    F9 --> F10 --> F11 --> F12\n" +
		"    F13 --> F14 --> F15 --> F16\n" +
		"    F17 --> F18 --> F19\n" +
		"    F20 --> F21 --> F22 --> F23 --> F24\n" +
		"    F25 --> F26 --> F27\n" +
		"```\n"

	path := filepath.Join(outDir, "AIAdCreate_Exe_function_creation_order.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// AIAdCreate_Exe_WriteFunctionExecutionOrder creates a flowchart showing the order functions should execute
func AIAdCreate_Exe_WriteFunctionExecutionOrder(outDir string) error {
	content := "```mermaid\n" +
		"flowchart TD\n" +
		"    subgraph ExecutionOrder[\"⚡ FUNCTION EXECUTION ORDER (Runtime Sequence)\"]\n" +
		"        %% Runtime Phase 1: Application Startup\n" +
		"        subgraph Runtime1[\"🚀 RUNTIME PHASE 1: APPLICATION STARTUP\"]\n" +
		"            E1[\"1. main()<br/>📍 main.go<br/>🎯 Application entry point\"]\n" +
		"            E2[\"2. OpenDatabase()<br/>📍 internal/database/database.go<br/>🎯 Establish DB connection\"]\n" +
		"            E3[\"3. Migrate()<br/>📍 internal/database/migrate.go<br/>🎯 Apply schema migrations\"]\n" +
		"        end\n\n" +
		"        %% Runtime Phase 2: Store Initialization\n" +
		"        subgraph Runtime2[\"💾 RUNTIME PHASE 2: STORE INITIALIZATION\"]\n" +
		"            E4[\"4. NewPostgresWorkoutStore()<br/>📍 internal/store/workout_store.go<br/>🎯 Initialize workout store\"]\n" +
		"            E5[\"5. NewUserStore()<br/>📍 internal/store/user_store.go<br/>🎯 Initialize user store\"]\n" +
		"            E6[\"6. NewTokenStore()<br/>📍 internal/tokens/token_store.go<br/>🎯 Initialize token store\"]\n" +
		"        end\n\n" +
		"        %% Runtime Phase 3: Middleware Setup\n" +
		"        subgraph Runtime3[\"🛡️ RUNTIME PHASE 3: MIDDLEWARE SETUP\"]\n" +
		"            E7[\"7. AuthMiddleware()<br/>📍 internal/middleware/auth.go<br/>🎯 Setup authentication\"]\n" +
		"            E8[\"8. CORSMiddleware()<br/>📍 internal/middleware/cors.go<br/>🎯 Setup CORS\"]\n" +
		"            E9[\"9. ValidateUserOwnership()<br/>📍 internal/middleware/ownership.go<br/>🎯 Setup ownership validation\"]\n" +
		"        end\n\n" +
		"        %% Runtime Phase 4: Handler Initialization\n" +
		"        subgraph Runtime4[\"🌐 RUNTIME PHASE 4: HANDLER INITIALIZATION\"]\n" +
		"            E10[\"10. NewWorkoutHandler()<br/>📍 internal/api/workout_handler.go<br/>🎯 Initialize workout handler\"]\n" +
		"            E11[\"11. NewUserHandler()<br/>📍 internal/api/user_handler.go<br/>🎯 Initialize user handler\"]\n" +
		"            E12[\"12. NewTokenHandler()<br/>📍 internal/api/token_handler.go<br/>🎯 Initialize token handler\"]\n" +
		"        end\n\n" +
		"        %% Runtime Phase 5: Application Layer Setup\n" +
		"        subgraph Runtime5[\"🏗️ RUNTIME PHASE 5: APPLICATION LAYER SETUP\"]\n" +
		"            E13[\"13. NewApplication()<br/>📍 internal/app/app.go<br/>🎯 Initialize application\"]\n" +
		"            E14[\"14. SetupRoutes()<br/>📍 internal/app/routes.go<br/>🎯 Configure HTTP routes\"]\n" +
		"        end\n\n" +
		"        %% Runtime Phase 6: Server Startup\n" +
		"        subgraph Runtime6[\"🌐 RUNTIME PHASE 6: SERVER STARTUP\"]\n" +
		"            E15[\"15. StartServer()<br/>📍 internal/app/server.go<br/>🎯 Begin serving requests\"]\n" +
		"        end\n\n" +
		"        %% Runtime Phase 7: Request Processing (Per Request)\n" +
		"        subgraph Runtime7[\"📡 RUNTIME PHASE 7: REQUEST PROCESSING (Per Request)\"]\n" +
		"            E16[\"16. CORSMiddleware()<br/>📍 internal/middleware/cors.go<br/>🎯 Handle CORS\"]\n" +
		"            E17[\"17. AuthMiddleware()<br/>📍 internal/middleware/auth.go<br/>🎯 Authenticate request\"]\n" +
		"            E18[\"18. Route Handler<br/>📍 internal/api/*.go<br/>🎯 Process request\"]\n" +
		"            E19[\"19. Store Operation<br/>📍 internal/store/*.go<br/>🎯 Database operation\"]\n" +
		"            E20[\"20. Response<br/>📍 internal/api/*.go<br/>🎯 Send response\"]\n" +
		"        end\n" +
		"    end\n\n" +
		"    %% Execution sequence connections\n" +
		"    Runtime1 --> Runtime2\n" +
		"    Runtime2 --> Runtime3\n" +
		"    Runtime3 --> Runtime4\n" +
		"    Runtime4 --> Runtime5\n" +
		"    Runtime5 --> Runtime6\n" +
		"    Runtime6 --> Runtime7\n\n" +
		"    %% Internal phase connections\n" +
		"    E1 --> E2 --> E3\n" +
		"    E4 --> E5 --> E6\n" +
		"    E7 --> E8 --> E9\n" +
		"    E10 --> E11 --> E12\n" +
		"    E13 --> E14\n" +
		"    E16 --> E17 --> E18 --> E19 --> E20\n\n" +
		"    %% Loop back for multiple requests\n" +
		"    E20 -.->|\"Next Request\"| E16\n" +
		"```\n"

	path := filepath.Join(outDir, "AIAdCreate_Exe_function_execution_order.mmd.md")
	return os.WriteFile(path, []byte(content), 0644)
}

// AIAdCreate_Exe_WriteAllFunctionDiagrams generates both creation and execution order diagrams
func AIAdCreate_Exe_WriteAllFunctionDiagrams(outDir string) error {
	fmt.Println("📊 Generating AI Advisor Function Creation and Execution Order Diagrams...")

	// Generate function creation order diagram
	if err := AIAdCreate_Exe_WriteFunctionCreationOrder(outDir); err != nil {
		return fmt.Errorf("failed to write function creation order diagram: %w", err)
	}

	// Generate function execution order diagram
	if err := AIAdCreate_Exe_WriteFunctionExecutionOrder(outDir); err != nil {
		return fmt.Errorf("failed to write function execution order diagram: %w", err)
	}

	fmt.Println("✅ AI Advisor function creation and execution order diagrams generated successfully!")
	return nil
}
