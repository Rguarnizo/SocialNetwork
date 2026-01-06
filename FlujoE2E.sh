#!/bin/bash

set -e

AUTH_URL="http://localhost:8080"
POST_URL="http://localhost:9090"
PROFILE_URL="http://localhost:7070"

EMAIL="test@test.com"
PASSWORD="123456"

echo "🔐 1. Login..."

LOGIN_RESPONSE=$(curl -s -X POST \
  "$AUTH_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email":"test@test.com",
    "password":"123456"
  }')

echo "$LOGIN_RESPONSE"
TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')

if [ "$TOKEN" == "null" ] || [ -z "$TOKEN" ]; then
  echo "❌ Login failed"
  echo "$LOGIN_RESPONSE"
  exit 1
fi

echo "✅ Token obtenido"
echo

echo "📝 2. Crear publicación..."

CREATE_POST_RESPONSE=$(curl -s -X POST \
  "$POST_URL/posts" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Esta es una publicación creada desde el script"
  }')

echo "Respuesta CreatePost:"
echo "$CREATE_POST_RESPONSE"
echo

echo "📥 3. Obtener publicaciones..."

GET_POSTS_RESPONSE=$(curl -s -X GET \
  "$POST_URL/posts" \
  -H "Authorization: Bearer $TOKEN")

echo "Posts:"
echo "$GET_POSTS_RESPONSE"
echo

echo "👤 4. Obtener perfil del usuario..."

PROFILE_RESPONSE=$(curl -s -X GET \
  "$PROFILE_URL/profile/me" \
  -H "Authorization: Bearer $TOKEN")

echo "Perfil:"
echo "$PROFILE_RESPONSE"
echo

echo "🎉 Flujo completo ejecutado correctamente"
