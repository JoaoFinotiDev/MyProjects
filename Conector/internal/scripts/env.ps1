# scripts/env.ps1
# Uso:
#   .\scripts\env.ps1
#   go run .\cmd\gateway
# ou:
#   . .\scripts\env.ps1   # (dot-source) mantém vars na mesma sessão

$ErrorActionPreference = "Stop"

#MONGO
$env:MONGO_URI = "mongodb+srv://impvaultauto:WZUlqXaVHkOLtVal@vault0.ripbvuv.mongodb.net/?appName=Vault0"
$env:MONGO_DB  = "imp"
$env:PORT      = "8080"

#APP
$env:APP_NAME="fire-gopher"
$env:APP_ENV="dev"
$env:APP_PORT="8080"

#FIREBIRD
$env:FIREBIRD_HOST="localhost"
$env:FIREBIRD_PORT="3050"
$env:FIREBIRD_DB="c:\asr\master\banco\master.fdb"
$env:FIREBIRD_USER="sysdba"
$env:FIREBIRD_PASSWORD="masterkey"
$env:FIREBIRD_CHARSET="UTF8"


Write-Host "✅ Ambiente carregado:"
Write-Host "  MONGO_URI = $($env:MONGO_URI)"
Write-Host "  MONGO_DB  = $($env:MONGO_DB)"
Write-Host "  PORT      = $($env:PORT)"
Write-Host "  FIREBIRD_HOST = $($env:FIREBIRD_HOST)"
Write-Host "  FIREBIRD_PORT = $($env:FIREBIRD_PORT)"
Write-Host "  FIREBIRD_DB = $($env:FIREBIRD_DB)"
Write-Host "  FIREBIRD_USER = $($env:FIREBIRD_USER)"
Write-Host "  FIREBIRD_PASSWORD = $($env:FIREBIRD_PASSWORD)"
Write-Host "  FIREBIRD_CHARSET = $($env:FIREBIRD_CHARSET)"