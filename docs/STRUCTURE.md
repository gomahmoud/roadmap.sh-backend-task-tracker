task-tracker/
├── docs/                     # Documentation directory
│   ├── STRUCTURE.md          # Application directory structure
│   ├── USAGE.md              # CLI commands, flags, and usage examples
│   └── REQUIRMENTS.md        # Application's overview, scope, functional requirements etc.
├── main.go                   # Application entry point (compiles to 'task-cli' binary)
├── cli.go                    # Command-line flags and console rendering
├── manager.go                # Business logic and status management
├── task.go                   # Task struct definition
└── storage.go                # Interface and JSON persistence implementation
├── go.mod                    # Module definition (e.g., github.com/gomahmoud/roadmapsh-backend-task-tracker)
└── README.md                 # Project overview and installation guide
