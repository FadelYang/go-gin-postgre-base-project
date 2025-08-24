# Go + Gin + PostgreSQL Base Project
This is simple Go, Gin, and PostgreSQL setup, the folder structure is group by each modules and implement "Service Repository Pattern" with dependency injection.

# 🏗️ Project Structure
```
project-root/
├─ common/              # shared common data types (structs) used in multiple files
├─ config/              # init config for tools like database, env, etc.
├─ modules/
│   └─ examples/
│       ├─ controller/
│       ├─ dto/
│       ├─ model/
│       ├─ providers/
│       ├─ repository/
│       ├─ routes/
│       └─ services/
├─ providers/           # global provider for dependency injection
├─ routes/              # main router setup
├─ tools/               # shared tools or logic used in multiple files
└─ main.go              # app entrypoint
```