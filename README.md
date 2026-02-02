# code-verse - Coding Typing Test

A Monkeytype-style typing test specifically designed for coding practice. Users type real code snippets and are scored on speed, accuracy, and syntax discipline.

## Architecture

### Monorepo Structure
```
├── apps/
│   └── web/                 # Next.js frontend application
├── services/
│   ├── api/                 # Go HTTP API server
│   ├── snippet-engine/      # Go package for snippet management
│   └── stats-engine/        # Go package for scoring and statistics
├── packages/
│   └── shared/              # Shared TypeScript types and contracts
├── docker/                  # Docker configuration
└── go.work                  # Go workspace configuration
```

### Technology Stack

**Backend:**
- **Go with Fiber**: Chosen for its excellent performance, simple API design, and built-in middleware support. Fiber provides a more ergonomic developer experience compared to net/http while maintaining high performance.
- **Embedded code snippets**: Real code snippets are stored as embedded assets for better deployment and consistency.
- **Stateless design**: All session state is managed client-side with server-side validation.

**Frontend:**
- **Next.js with React**: Provides excellent developer experience, built-in optimization, and server-side rendering capabilities.
- **Custom typing interface**: Manual keyboard capture with visual feedback, avoiding textarea limitations for precise character tracking.
- **Real-time feedback**: Immediate visual indication of correct/incorrect characters, current position, and progress.

## Core Features

### Supported Languages
- **Python**: Focused on clean syntax, proper indentation, and Pythonic patterns
- **JavaScript**: Modern ES6+ syntax, async patterns, and common frameworks
- **Go**: Idiomatic Go code with proper error handling and concurrency

### Typing Mechanics
- **Character-by-character typing**: Each keystroke is tracked and validated
- **Significant whitespace**: Indentation and spacing are part of the typing challenge
- **Backspace allowed**: Users can correct mistakes, but errors are tracked
- **Paste protection**: Clipboard input is blocked to ensure legitimate typing
- **Completion requirement**: Full snippet must match exactly to complete the test

### Scoring System
- **Characters Per Minute (CPM)**: Raw typing speed measurement
- **Accuracy Percentage**: Ratio of correct characters to total characters typed
- **Error Count**: Total number of incorrect characters typed
- **Completion Time**: Total time taken to complete the snippet
- **Syntax Discipline Penalties**: Deductions for whitespace errors, indentation mistakes, and symbol errors

The scoring algorithm weights syntax discipline heavily since proper formatting is crucial in coding.

## Development Setup

### Prerequisites
- Go 1.21+
- Node.js 18+
- Docker (optional)

### Local Development

1. **Clone and setup:**
```bash
git clone <repository>
cd code-verse
```

2. **Backend setup:**
```bash
# Initialize Go workspace
go work sync
cd services/api
go mod tidy
go run main.go
```

3. **Frontend setup:**
```bash
cd apps/web
npm install
npm run dev
```

4. **Docker setup (alternative):**
```bash
docker-compose up --build
```

### API Endpoints

- `GET /api/languages` - Get available programming languages
- `GET /api/snippets?language=python&difficulty=easy` - Get code snippets
- `POST /api/sessions/start` - Start a new typing session
- `POST /api/sessions/finish` - Complete session and get score

### Future Architecture Considerations

The current design supports future extensions:
- **User accounts**: JWT-based authentication can be added
- **Leaderboards**: PostgreSQL integration for persistent storage
- **Multiplayer races**: WebSocket support for real-time multiplayer
- **Code execution**: Sandboxed execution environment integration

## Project Philosophy

This implementation prioritizes:
- **Correctness**: All code is functional and testable
- **Maintainability**: Clear separation of concerns and readable code
- **Performance**: Efficient algorithms and minimal dependencies
- **Scalability**: Architecture supports future feature additions