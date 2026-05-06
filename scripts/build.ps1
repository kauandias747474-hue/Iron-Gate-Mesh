Write-Host "--- IronGate: Compilando Visuals ---" -ForegroundColor Cyan


$rootDir = Get-Item "$PSScriptRoot\.."
Push-Location $rootDir.FullName


go build -o irongate-visuals.exe .

if ($LASTEXITCODE -eq 0) {
    Write-Host " Sucesso! Gerado: irongate-visuals.exe" -ForegroundColor Green
} else {
    Write-Host " Erro na compilação." -ForegroundColor Red
}

Pop-Location
