param([string]$Target)

switch ($Target) {
    "run"   { go run ./cmd/server }
    "build" { go build -o bin/server.exe ./cmd/server }
    "test"  { go test ./internal/api/... }
    default { 
        Write-Host "Usage: .\make.ps1 [run|build|test]"
        Write-Host "  run   - go run ./cmd/server"
        Write-Host "  build - go build -o bin/server.exe ./cmd/server" 
        Write-Host "  test  - go test ./internal/api/..."
    }
}