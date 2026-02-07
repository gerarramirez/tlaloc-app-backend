# Tlaloc Security Service

API Gateway de seguridad para gestión de autenticación y autorización con JWT moderno.

## Características

- ✅ Login con JWT (Access Token + Refresh Token)
- ✅ Refresh Token service  
- ✅ Token validation para microservicios
- ✅ Logout service
- ✅ Middleware de autenticación y autorización por roles
- ✅ Hashing de contraseñas con Argon2
- ✅ Seguridad con estándares modernos
- ✅ Base de datos PostgreSQL con GORM
- ✅ Estructura consistente con otros microservicios

## Endpoints

### Autenticación Pública
- `POST /auth/login` - Login de usuario
- `POST /auth/refresh` - Refresh token
- `POST /auth/validate` - Validar token (para microservicios)

### Autenticación Requerida
- `POST /auth/logout` - Logout de usuario
- `GET /auth/me` - Obtener información del usuario actual

### Admin only
- `GET /admin/test` - Endpoint de prueba para administradores

## Configuración

Variables de entorno en `config.env`:
```
DB_HOST=localhost
DB_USER=postgres  
DB_PASSWORD=tu_password
DB_NAME=tlaloc_security
DB_PORT=5432
JWT_SECRET=secreto-min-32-caracteres-para-jwt
JWT_REFRESH_SECRET=secreto-min-32-caracteres-para-refresh
SERVER_PORT=8081
```

## Ejemplos de Uso

### Login
```bash
curl -X POST http://localhost:8081/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@tlaloc.com",
    "password": "admin123456"
  }'
```

### Validar Token (para microservicios)
```bash
curl -X POST http://localhost:8081/auth/validate \
  -H "Content-Type: application/json" \
  -d '{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }'
```

### Refresh Token
```bash
curl -X POST http://localhost:8081/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }'
```

## Medidas de Seguridad Implementadas

1. **JWT moderno** con `github.com/golang-jwt/jwt/v5`
2. **Argon2ID** para hashing de contraseñas (el algoritmo más seguro recomendado)
3. **Separación de claves** para Access Token y Refresh Token
4. **Tiempos de expiración** diferentes (15 min access, 7 días refresh)
5. **Blacklist** de tokens mediante logout
6. **Validación de claims** en todos los tokens
7. **Middleware** para proteger rutas
8. **CORS configurado** para producción
9. **Rol-based access control** (RBAC)
10. **Input validation** en todos los endpoints

## Usuarios de Prueba

- **Admin**: `admin@tlaloc.com` / `admin123456`
- **User**: `user@tlaloc.com` / `user123456`

## Instalación y Ejecución

```bash
cd tlaloc-security-service
go mod tidy
go run main.go
```

El servicio se iniciará en `http://localhost:8081`