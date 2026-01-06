$ErrorActionPreference = "Stop"

$AUTH_URL    = "http://localhost:8080"
$POST_URL    = "http://localhost:9090"
$PROFILE_URL = "http://localhost:7070"

$EMAIL    = "test@test.com"
$PASSWORD = "123456"

Write-Host "🔐 1. Login..."

$loginResponse = Invoke-RestMethod `
    -Method GET `
    -Uri "$AUTH_URL/auth/login" `
    -Body @{
        user     = $EMAIL
        password = $PASSWORD
    }

$TOKEN = $loginResponse.token

if (-not $TOKEN) {
    Write-Error "❌ Login failed"
    exit 1
}

Write-Host "✅ Token obtenido"
Write-Host ""

$headers = @{
    Authorization = "Bearer $TOKEN"
    "Content-Type" = "application/json"
}

Write-Host "📝 2. Crear publicación..."

$createPostBody = @{
    content = "Esta es una publicación creada desde PowerShell"
} | ConvertTo-Json

$createPostResponse = Invoke-RestMethod `
    -Method POST `
    -Uri "$POST_URL/post" `
    -Headers $headers `
    -Body $createPostBody

Write-Host "Respuesta CreatePost:"
$createPostResponse | ConvertTo-Json -Depth 5
Write-Host ""

Write-Host "📥 3. Obtener publicaciones..."

$getPostsResponse = Invoke-RestMethod `
    -Method GET `
    -Uri "$POST_URL/posts" `
    -Headers $headers

Write-Host "Posts:"
$getPostsResponse | ConvertTo-Json -Depth 5
Write-Host ""

Write-Host "👤 4. Obtener perfil del usuario..."

$profileRespo

