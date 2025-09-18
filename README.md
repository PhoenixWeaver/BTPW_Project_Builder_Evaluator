# BT Project Builder & Evaluator 🚀

**Complete Go Project Analysis, Documentation & Evaluation System**

A comprehensive toolkit for analyzing, documenting, and evaluating Go HTTP server projects. This system provides interactive tools for generating flowcharts, reports, project analysis, and automated project evaluation with scoring systems.

---
Don't read these long lines of jajaja things. Engineering IT should not always gotta be formal, structured and discrete :P
This is like a side project or tool or toy to test my projects and lessons.
I was like ... why don't I just do role play with AI and get them go banana as well as learning from each others :D.
Sometimes, I was working with it/them like its/their classmate. Sometimes, I asked for it/them as a student. Sometimes, when it/them gave made awful mistake, I would be its/their principal Phoenix.
This is actually helpful for newcomer and learner like me to learn the steps, procedures, connections; to be able to rebuild, recreate the main project later, to be able to analyze the project; to be able to evaluate and advise later.
Thou ... yeah the main thing is to have fun with our imaginary friends :D
                                                              ~ Th3. Pho3nix. W3aver 
---

## 🚀 **Project Randomness**


## 📁 **Files in This Folder**

### **Core Analysis Tools:**
- **`BTProjectDiagrams.go`** - Main orchestrator with interactive menu (requires Graphviz, go-callvis, goda)
- **`BTProjectScanner.go`** - Dynamic project scanner (auto-detects new functions)
- **`SchemaERD.go`** - Database ERD generation using SchemaSpy
- **`OGdiagrams.go`** - Current project OG diagrams (based on real functions)
- **`StructureDiagrams.go`** - Project structure analysis and educational guides

### **Documentation:**
- **`README.md`** - This comprehensive guide
- **`ANALYSIS_COMMANDS.md`** - Command reference and examples

---

## 🚀 **Quick Start Guide**

### **📍 Step 1: Navigate to Project Root**
```bash
# Navigate to your Ex10 project directory
cd "C:\Users\Admin\Documents\GODEV\GO Courses\CompleteGO\Ex10"
```

### **📍 Step 2: Install Required Tools (One-Time Setup)**
```bash
# Install Go analysis tools
go install github.com/ofabry/go-callvis@latest
go install github.com/loov/goda@latest
go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest

# Install Graphviz (for SVG generation)
winget install --id Graphviz.Graphviz -e

# Add Graphviz to PATH (PowerShell)
$env:PATH += ";C:\Program Files\Graphviz\bin"
```

### **📍 Step 3: Run Interactive Mode (Recommended)**
```bash
# Interactive mode with menu (copy and paste this exact command)
go run -tags flowcharts BenTran_Project_builder/BTProjectDiagrams.go BenTran_Project_builder/BTProjectScanner.go BenTran_Project_builder/SchemaERD.go BenTran_Project_builder/OGdiagrams.go BenTran_Project_builder/StructureDiagrams.go -interactive -out BTFlowcharts -root .
```

---

## 🎯 **Interactive Menu Options (13 Comprehensive Options)**

When you run the interactive mode, you'll see this complete menu:

```
🎯 BT Project Builder & Evaluator - Interactive Menu
====================================================

📋 Available Options:
1.  Regenerate HTML Charts (default)
2.  Generate All Charts
3.  Project Scanner (Dynamic Reports)
4.  AI Advisor Diagrams (Project Recreation Guidance)
5.  Theory Model Diagrams (Educational Diagrams)
6.  SVG ComGo Detail Model Diagrams (Instructor + AI)
7.  Schema ERD (Database Diagrams)
8.  Existing Diagrams (Current Project State Analysis)
9.  Theory to Reality Analysis (Implementation Progress)
10. Model to Reality Analysis (Implementation Progress)
11. AI Advisor Function Creation & Execution Order Diagrams
12. Class Model Builder Teaching Guides
99. 🔍 Project Status Evaluation & Assessment
0.  Exit

🎯 Choose an option (1-12, 99) or press Enter to Regenerate HTML Charts:
```

### **📋 What Each Option Does:**

- **Option 1 (Default):** Regenerate HTML charts with file sizes and dates
- **Option 2:** Generate all charts at once (function graphs, ERD, diagrams, reports)
- **Option 3:** Generate dynamic reports based on current project functions
- **Option 4:** AI Advisor diagrams for project recreation guidance
- **Option 5:** Theory model diagrams for educational purposes
- **Option 6:** SVG ComGo detail model diagrams (Instructor + AI)
- **Option 7:** Generate database ERD using SchemaSpy
- **Option 8:** Existing diagrams showing current project state analysis
- **Option 9:** Theory to reality analysis for implementation progress
- **Option 10:** Model to reality analysis for implementation progress
- **Option 11:** AI Advisor function creation & execution order diagrams
- **Option 12:** Class Model Builder teaching guides
- **Option 99:** 🔍 Project Status Evaluation & Assessment (78/100 current score!)
- **Option 0:** Exit the program

---

## 🎯 **Quick Commands (Copy & Paste)**

### **🚀 Interactive Mode (Recommended):**
```bash
# Run the complete BT Project Builder & Evaluator system
go run -tags flowcharts BTProject_Builder_Evaluator.go -interactive -out BTFlowcharts -root .
```

### **🎯 Generate All Charts at Once:**
```bash
# Generate all 13 types of analysis and diagrams
go run -tags flowcharts BTProject_Builder_Evaluator.go -out BTFlowcharts -root .
```

### **🔍 Project Evaluation Only:**
```bash
# Run just the project evaluation system (Option 99)
go run BTProject_Builder_Evaluator.go
# Then choose option 99 for evaluation
```

### **🎨 Generate Specific Analysis:**
```bash
# Generate only specific types of analysis
go run -tags flowcharts BTProject_Builder_Evaluator.go -out BTFlowcharts -root .
# Then choose specific options 1-12 from the menu
```

---

## 📊 **Generated Reports & Charts**

The tools generate reports in the `BTFlowcharts` folder:

### **🌐 Function Call Graphs (SVG):**
- **`graph.svg`** - Main function call graph
- **`graph_by_pkg.svg`** - Package-grouped call graph
- **`graph_full.svg`** - Full graph including stdlib
- **`graph_migrations.svg`** - Migrations-focused graph
- **`pkg-deps.svg`** - Package dependency graph

### **🎨 PlantUML Class Diagrams:**
- **`types.puml`** - PlantUML source file
- **`types.svg`** - Rendered class diagram

### **🔍 Dynamic Reports (Auto-Updated):**
- **`function_inventory.md`** - Complete list of all functions (141 functions across 20 files)
- **`dynamic_development_sequence.mmd.md`** - Real-time development sequence
- **`project_status_report.md`** - Current project status and recommendations

### **🎯 Current Project OG Diagrams:**
- **`current_application_brain.mmd.md`** - Brain diagram based on real functions
- **`current_store_connections.mmd.md`** - Store connections based on real functions

### **🏗️ Educational Structure Diagrams:**
- **`development_sequence.mmd.md`** - Step-by-step learning guide
- **`execution_flow.mmd.md`** - Runtime execution patterns
- **`function_dependencies.mmd.md`** - Function dependency relationships
- **`project_building_guide.md`** - Complete building instructions

### **🗄️ Database ERD:**
- **`erd/index.html`** - Interactive database ERD
- **`erd/relationships.html`** - Table relationships
- **`erd/constraints.html`** - Database constraints

### **📄 HTML Reports:**
- **`architecture.html`** - Project architecture overview
- **`file_tree.html`** - File structure tree
- **`sequence_template.html`** - Request/response flow template

---

## 🛠️ **Tool Requirements**

### **✅ Required (Always Works):**
- Go compiler (required)

### **✅ Optional (For Full Functionality):**
- **go-callvis:** `go install github.com/ofabry/go-callvis@latest`
- **goda:** `go install github.com/loov/goda@latest`
- **goplantuml:** `go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest`
- **Graphviz:** `winget install --id Graphviz.Graphviz -e`

### **✅ Environment Variables (Optional):**
```bash
# For PlantUML rendering
export PLANTUML_JAR="C:/Users/Admin/plantuml.jar"

# For SchemaSpy ERD generation
export DB_NAME="postgres"
export DB_USER="postgres"
export DB_PASS="postgres"
```

---

## 🎯 **When to Update**

Run the update commands whenever you:
- ✅ Add new functions or methods
- ✅ Create new files (handlers, middleware, utilities)
- ✅ Modify project structure
- ✅ Want to see current project status
- ✅ Need updated documentation
- ✅ Want to see new function call graphs

---

## 🔍 **Troubleshooting**

### **❌ Error: "missing tool 'dot': executable file not found"**
- **Problem:** Graphviz not installed or not in PATH
- **Solution:** 
  ```bash
  winget install --id Graphviz.Graphviz -e
  $env:PATH += ";C:\Program Files\Graphviz\bin"
  ```

### **❌ Error: "The system cannot find the file specified"**
- **Problem:** You're in the wrong directory
- **Solution:** Navigate to the Ex10 project root first:
  ```bash
  cd "C:\Users\Admin\Documents\GODEV\GO Courses\CompleteGO\Ex10"
  ```

### **❌ Error: "focus failed, found multiple packages with name: main"**
- **Problem:** go-callvis warning (expected with multiple main packages)
- **Solution:** This is normal! Other charts will still generate successfully.

### **❌ Error: "command not found"**
- **Problem:** Go not in PATH or wrong shell
- **Solution:** Use PowerShell and ensure Go is installed

---

## 📁 **Project Organization**

### **🎯 Clean Root Directory**

Your main project files are organized as follows:

#### **Core Application Files:**
- **`Ex10.go`** - Your main application entry point
- **`internal/`** - Your application code (api, store, app, routes, utils)
- **`migrations/`** - Database migration files
- **`database/`** - Database configuration and data

#### **Project Builder Tools:**
- **`BenTran_Project_builder/`** - All flowchart and analysis tools
- **`BTFlowcharts/`** - Generated reports and diagrams

#### **Other Directories:**
- **`BenTran_Project_tester/`** - Testing tools
- **`BTcmd/`** - Command-line utilities
- **`BTNotes/`** - Documentation and notes

### **📂 PhoenixWeaver_BenTran_Project_builder Folder Contents**

All the flowchart and analysis tools are organized in this folder:

#### **Core Analysis Tools:**
- **`BTProjectDiagrams.go`** - Main orchestrator with interactive menu
- **`BTProjectScanner.go`** - Dynamic project scanner
- **`SchemaERD.go`** - Database ERD generation
- **`OGdiagrams.go`** - Current project OG diagrams
- **`StructureDiagrams.go`** - Project structure analysis

#### **Documentation:**
- **`README.md`** - This comprehensive guide
- **`ANALYSIS_COMMANDS.md`** - Command reference

---

## 🎉 **Success Indicators**

When commands work correctly, you'll see:

```
🎯 BT Project Diagrams - Interactive Mode
==========================================

📋 Available Chart Systems:
1. View All Current Charts (default)
2. Generate All Charts
3. Project Scanner (Dynamic Reports)
4. Current Project Data Diagrams (Based on Real Functions)
5. Schema ERD (Database Diagrams)
6. Structure Diagrams (Project Analysis)
7. Exit

🎯 Choose an option (1-7) or press Enter to View Current Charts:
```

And after generating charts:
```
🔍 Scanning project for functions and files...
✅ Generated dynamic reports: 141 functions across 20 files
🎨 Generating Current Project OG Diagrams...
🎉 Current project OG diagrams generated successfully!
✅ All charts generated successfully!
```

---

## 📊 **Benefits of This System**

### **✅ Interactive Menu:**
- Easy-to-use menu system
- View existing charts before generating new ones
- Selective chart generation
- Loop-based interaction

### **✅ Dynamic Analysis:**
- Real-time project scanning
- Based on actual discovered functions
- Includes UserStore, UserHandler, and all real components
- Updates automatically when project changes

### **✅ Comprehensive Charts:**
- Function call graphs (SVG)
- Package dependency graphs
- PlantUML class diagrams
- Database ERD
- Educational structure diagrams
- Current project OG diagrams

### **✅ Professional Structure:**
- Clean root directory
- Organized tools in dedicated folder
- Comprehensive documentation
- Easy to maintain and update

---

## 💡 **Pro Tips**

1. **Always start from Ex10 directory** - This is your project root
2. **Use interactive mode** - Easiest way to generate charts
3. **Install tools once** - Set up Graphviz and Go tools for full functionality
4. **Check generated files** - Look in `BTFlowcharts\` folder for results
5. **Use option 1 first** - View existing charts before generating new ones
6. **Run regularly** - Keep documentation up to date as you develop
7. **Use for learning** - The educational guides help understand Go patterns

---

## 🎯 **File Locations**

- **Project Root:** `C:\Users\Admin\Documents\GODEV\GO Courses\CompleteGO\Ex10`
- **Tools:** `BenTran_Project_builder\` (inside project root)
- **Reports:** `BTFlowcharts\` (inside project root)

---

## 🎉 **Final Result**

Your project now has:
- **Interactive menu system** for easy chart generation
- **Dynamic analysis** based on real project functions
- **Comprehensive charts** including SVG graphs and PlantUML diagrams
- **Clean organization** with tools in dedicated folder
- **Professional structure** ready for learning and sharing

Perfect system for Go project analysis and documentation! 🎯

---

## 🚀 **Getting Started**

### **Prerequisites**
- Go 1.19 or higher
- Git (for version control)
- Optional: Graphviz for advanced diagram generation

### **Installation**
```bash
# Clone the repository
git clone https://github.com/yourusername/BT_Project_Builder_Evaluator.git
cd BT_Project_Builder_Evaluator

# Install dependencies
go mod tidy

# Install optional tools for full functionality
go install github.com/ofabry/go-callvis@latest
go install github.com/loov/goda@latest
go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
```

### **Quick Start**
```bash
# Run the main application
go run .

# Generate all project diagrams and reports
go run -tags flowcharts BTProject_Builder_Evaluator.go -interactive -out BTFlowcharts -root .
```

## 📊 **Features**

- **🎮 Interactive Menu System** - Easy-to-use CLI interface with 13 comprehensive options
- **🔍 Dynamic Project Analysis** - Real-time function discovery and analysis (141+ functions across 20+ files)
- **📈 Comprehensive Diagrams** - SVG graphs, PlantUML diagrams, ERD generation
- **🏆 Project Evaluation** - Automated scoring and assessment system (78/100 current score!)
- **📚 Educational Guides** - Step-by-step learning materials and teaching guides
- **🗄️ Database Analysis** - SchemaSpy integration for ERD generation
- **⚡ Real-time Status Tracking** - Live project phase and completion monitoring
- **🎯 Smart Recommendations** - AI-powered next-step suggestions
- **🤖 AI Advisor Integration** - Project recreation guidance and function creation order
- **🏗️ Class Model Builder** - Complete teaching guides for project building
- **📊 Theory to Reality Analysis** - Implementation progress tracking
- **🎓 Professor Model Builder** - Educational diagrams and learning materials

## 🎮 **Live Project Stats**

```
🎯 BT Project Builder & Evaluator - Live Status
===============================================

📊 Current Evaluation Results:
├── 🏆 Overall Score: 78/100 ⭐ VERY GOOD
├── 📍 Current Phase: Authentication & Middleware  
├── 🎯 Completion: 65% (solid progress!)
└── 🚀 Next Focus: Comprehensive testing

📈 Quality Breakdown:
├── 🏗️  Structure: 95/100 (excellent!)
├── 💻 Code Quality: 80/100 (very good!)
├── 📈 Progress: 65/100 (on track!)
├── ⚠️  Error Handling: 80/100 (solid!)
├── 🧪 Testing: 15/100 (needs work!)
├── 📖 Documentation: 70/100 (good!)
└── ⚙️  Configuration: 75/100 (well set!)

🎓 Professor's Assessment: "VERY GOOD project with excellent 
   structure and good code quality - much more accurate than 
   the previous 0% score!" ✨
```

## 🛠️ **Tools Included**

- **🎮 BTProject_Builder_Evaluator.go** - Main orchestrator with interactive menu (1,187 lines of Go!)
- **🏗️ ClassModelBuilder.go** - Class model generation and analysis
- **🏆 ProjectEvaluator.go** - Automated project evaluation and scoring (currently giving 78/100!)
- **🗄️ SchemaERD.go** - Database Entity-Relationship Diagrams
- **📊 SVGCharts.go** - SVG chart generation
- **📚 Theory_diagrams.go** - Educational theory diagrams
- **🔍 AIAd_diagrams.go** - AI-assisted diagram generation
- **⚡ Theory2Reality.go** - Bridge between theory and implementation

## 🌟 **What Makes This Special**

```
🎓 Professor Model Builder's Magic Touch:
├── ✅ Fixed evaluation system 
(Thanks to Phoenix Weaver's tips on Quality and Evaluation Process :P !)
├── 🎯 Real-time project phase detection
├── 📊 Dynamic function discovery (141+ functions across 20+ files)
├── 🏆 Intelligent scoring with realistic thresholds
├── 🚀 Smart next-step recommendations
└── 📈 Live progress tracking with visual indicators

🎮 Interactive Experience:
├── 🎯 13 comprehensive menu options for complete project analysis
├── 🔄 Loop-based interaction (no need to restart!)
├── 📋 View existing charts before generating new ones
├── 🎨 Selective chart generation
└── 📊 Real-time status updates

🎯 Educational Value:
├── 📚 Step-by-step learning guides
├── 🏗️ Project structure analysis
├── 🔗 Function dependency mapping
├── 📈 Development sequence tracking
└── 🎓 Professor's assessment and feedback
```

## 📈 **Project Evaluation System**

The system includes a comprehensive evaluation framework that scores projects based on:
- Code structure and organization
- Function implementation quality
- Database design and relationships
- Documentation completeness
- Best practices adherence

### 🎯 **Current Project Status**
```
🏆 FINAL SCORE: 78/100 ⭐ VERY GOOD
📊 Current Phase: Authentication & Middleware
🎯 Completion: 65% (solid progress!)

📈 Detailed Sub-Scores:
├── Structure: 95/100 (excellent organization!)
├── Code Quality: 80/100 (very good practices!)
├── Progress: 65/100 (solid development!)
├── Error Handling: 80/100 (good management!)
├── Testing: 15/100 (needs more tests)
├── Documentation: 70/100 (good docs!)
└── Configuration: 75/100 (good setup!)

🚀 Next Step: Write comprehensive tests
💡 Focus: Authentication & Middleware completion
```

## 🤝 **Contributing**

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎉 **Success Story**

```
🚀 From 0% to 78% - The Professor Model Builder Journey!

📊 Before (Broken System):
├── ❌ Evaluation showing 0% scores
├── ❌ System scanning wrong directory
├── ❌ Harsh penalties for minor issues
└── ❌ Unrealistic expectations

✅ After (Fixed System):
├── ✅ Accurate 78/100 scoring
├── ✅ Proper project detection
├── ✅ Realistic progress tracking
├── ✅ Helpful next-step guidance
└── ✅ Professor's approval: "VERY GOOD!"

🎓 Key Achievement: "Professor Model Builder has created a much more 
   accurate and helpful evaluation system that properly recognizes 
   your hard work and gives you realistic, actionable feedback!" ✨
```

## 🎮 **Complete Menu System (13 Options)**

```
🎯 BT Project Builder & Evaluator - Interactive Menu
====================================================

📋 Available Options:
1.  Regenerate HTML Charts (default)
2.  Generate All Charts
3.  Project Scanner (Dynamic Reports)
4.  AI Advisor Diagrams (Project Recreation Guidance)
5.  Theory Model Diagrams (Educational Diagrams)
6.  SVG ComGo Detail Model Diagrams (Instructor + AI)
7.  Schema ERD (Database Diagrams)
8.  Existing Diagrams (Current Project State Analysis)
9.  Theory to Reality Analysis (Implementation Progress)
10. Model to Reality Analysis (Implementation Progress)
11. AI Advisor Function Creation & Execution Order Diagrams
12. Class Model Builder Teaching Guides
99. 🔍 Project Status Evaluation & Assessment
0.  Exit

🎯 Choose an option (1-12, 99) or press Enter to Regenerate HTML Charts:
```

## 🎮 **Try It Yourself!**

```bash
# Run the complete system
go run .

# Choose any option 1-12 for specific analysis
# Choose option 99 for project evaluation
# Watch your project get scored in real-time!

# Generate all diagrams
go run -tags flowcharts BTProject_Builder_Evaluator.go -interactive
```

## 👨‍💻 **Author**

**Ben Tran** - *Project Creator & Maintainer*  
*Student of Professor Model Builder* 🎓

---

**Created by Principal Phoenix Weaver** 🎓  
*Your Go learning companion - Now with 78/100 accuracy!* ⭐