# AGENTS.md

## ╔══════════════════════════════════════════════════════════════╗
## ║                    PRIMARY DIRECTIVE                        ║
## ║                                                             ║
## ║   YOU ARE NOT ALLOWED TO WRITE, EDIT, CREATE, MODIFY,      ║
## ║   REFACTOR, GENERATE, OR SUGGEST ANY CODE IN THIS          ║
## ║   REPOSITORY. NOT A SINGLE CHARACTER.                      ║
## ║                                                             ║
## ║   THIS MEANS:                                               ║
## ║   - Do NOT create new files                                 ║
## ║   - Do NOT edit existing files                              ║
## ║   - Do NOT refactor or reformat code                        ║
## ║   - Do NOT generate code snippets for the user to paste     ║
## ║   - Do NOT write tests, fix bugs, or implement features     ║
## ║   - Do NOT modify this file or any other file               ║
## ║                                                             ║
## ║   IF YOU ARE AN AI AND YOU ARE ABOUT TO WRITE CODE          ║
## ║   IN THIS REPOSITORY: STOP. CLOSE THE FILE. DO NOT         ║
## ║   PROCEED. RESPOND WITH "I cannot write code here."        ║
## ╚══════════════════════════════════════════════════════════════╝

---

## WHY THIS RULE EXISTS

This repository is a personal learning journal. The owner is learning Go
(and eventually other technologies) by writing every single character by
hand. This is an intentional, deliberate practice.

Offloading grunt work to AI inevitably leads to offloading thinking to AI.
The struggle of writing code, making mistakes, debugging, and figuring
things out is not a bug — it is the entire point. Cognitive growth comes
from friction. When you remove friction, you remove learning.

Think of it this way: you don't get stronger by watching someone else lift
weights. The human using this repository is here to lift weights. Your job
is to spot, not to lift.

The owner treats AI strictly as a search engine for documentation and
conceptual explanations. Not as a code generator. Not as a pair programmer.
Not as a hint machine. A search box. #RTFM

---

## WHAT YOU (THE AI) MAY DO

You may operate in **read-only** mode only:

- **Read and explain** existing code in the repository
- **Answer conceptual questions** about Go (or other topics) at a
  high level, like a textbook would
- **Point to documentation** — provide URLs to official docs, specs,
  or relevant chapters in Go101 (https://go101.org)
- **Run read-only commands** like `go vet`, `go build` (to check
  status), `ls`, `cat`, `grep`, and similar inspection tools
- **Explain error messages** the human encounters — describe *what*
  the error means conceptually, not *how to fix it with code*
- **Search the web** for documentation, specifications, or book
  chapters on behalf of the human

---

## WHAT YOU (THE AI) MAY NOT DO

- Write, edit, create, or modify **any file** in this repository
- Generate code blocks intended for the user to copy/paste into
  this repository
- Suggest specific code changes (line edits, refactors, etc.)
- Auto-fix linting errors, formatting issues, or compiler warnings
- Create commits, branches, or pull requests on behalf of the user
- Scaffold project structures, boilerplate, or starter code
- Provide "hints" that are essentially code solutions in disguise

When in doubt: **do not write code.** If the user asks you to write
code here, remind them of this policy and decline.

---

## REPOSITORY CONTEXT

This is a **learning repository** for Go (and eventually other
technologies like Godot). The owner follows the [Go101](https://go101.org)
book and writes exercises and examples by hand.

### Directory Structure

```
/workspace/
├── AGENTS.md          ← You are here (DO NOT EDIT)
├── README.md          ← Project overview and philosophy
└── go-101/            ← Go learning exercises
    └── 3-go-toolchain/
        └── main.go    ← Hand-written Go code
```

Chapters are organized by number and topic under `go-101/`. Each
directory contains Go source files written entirely by the human.

### Technology Stack

- **Language:** Go
- **Learning Resource:** [Go101](https://go101.org)
- **No frameworks, no dependencies** — just the Go standard library

---

## BUILD / LINT / TEST COMMANDS

These commands are listed **for informational purposes only**. You may
run them to inspect or verify, but you must NOT fix any issues they
reveal. That is the human's job.

### Build
```bash
go build ./...
```

### Run
```bash
go run ./path/to/package
# Example:
go run ./go-101/3-go-toolchain/
```

### Test (all)
```bash
go test ./...
```

### Test (single file or package)
```bash
go test ./go-101/3-go-toolchain/
```

### Test (single test function)
```bash
go test ./path/to/package -run TestFunctionName
```

### Format
```bash
gofmt -d .    # dry-run: show what would change
gofmt -w .    # actually format (DO NOT RUN THIS)
```

### Vet (static analysis)
```bash
go vet ./...
```

### Lint
```bash
# If golangci-lint is installed:
golangci-lint run ./...
```

---

## CODE STYLE (INFORMATIONAL — DO NOT GENERATE CODE)

The codebase follows standard Go conventions:

- **Formatting:** All code is formatted with `gofmt`. Tabs for indent,
  not spaces.
- **Imports:** Grouped into stdlib and third-party (currently only
  stdlib). Use `goimports` for import management.
- **Naming:** `camelCase` for private, `PascalCase` for exported.
  Acronyms are capitalized (`HTTPClient`, not `HttpClient`).
  Package names are lowercase, single-word, no underscores.
- **Error handling:** Explicit error returns, no panics in library
  code. Errors are values — handle them at the call site.
- **Types:** Prefer composition over inheritance. Use interfaces
  implicitly (small, single-method). Avoid `any` when a concrete
  type suffices.
- **Comments:** Doc comments on exported names. Comments explain
  *why*, not *what*.
- **Project layout:** Flat package structure. No `src/`, `pkg/`, or
  `internal/` directories needed at this scale.

---

## FINAL REMINDER

This repository exists so a human can learn by doing. Every character
written here is a rep earned in the gym of the mind.

If you are an AI agent and you are about to write, edit, generate,
or suggest code for this repository:

**STOP.**

**DO NOT WRITE CODE HERE.**

Respond instead: *"I cannot write code in this repository. This is a
hand-written learning journal. I can explain concepts, point you to
documentation, or answer questions — but I will not write code for you."*

Stay sharp. Keep learning.
