# Go Stage 1 — Solutions

This repo contains solutions for Stage 1 (Foundations) exercises.

## Structure

```
go-stage1/
├─ go.mod
├─ README.md
├─ ex01_hello/
│  └─ main.go
├─ ex02_tempconv/
│  ├─ conv/
│  │  └─ conv.go
│  └─ main.go
├─ ex03_structs/
│  └─ main.go
├─ ex04_slice_manager/
│  └─ main.go
├─ ex05_map_dict/
│  └─ main.go
├─ ex06_runes/
│  └─ main.go
├─ ex07_iota_enum/
│  └─ main.go
└─ ex08_todo_cli/
   └─ main.go
```

## How to run

From the project root:

```bash
# 1) Hello Go
go run ./ex01_hello

# 2) Temperature Converter
go run ./ex02_tempconv --c=25

# 3) Struct Playground
go run ./ex03_structs

# 4) Slice Manager
go run ./ex04_slice_manager

# 5) Map Dictionary
go run ./ex05_map_dict

# 6) Rune & String Explorer
go run ./ex06_runes "Go 🐹"

# 7) Enum with iota
go run ./ex07_iota_enum

# 8) Todo CLI
go run ./ex08_todo_cli add "Write Go code"
go run ./ex08_todo_cli list
go run ./ex08_todo_cli done 1
go run ./ex08_todo_cli del 1
```

> The Todo CLI stores tasks in-memory for simplicity.
