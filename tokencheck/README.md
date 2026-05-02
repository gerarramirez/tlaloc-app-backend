# tlaloc-tokencheck-lib

🔐 **Librería Centralizada de Validación de Tokens JWT para Microservicios Tlaloc**

## 📋 **Propósito**

Evitar la duplicación de código de autenticación en múltiples microservicios, proporcionando validación JWT consistente y mantenible desde un único punto.

---

## 🚀 **Características**

### **🔐 Validación JWT Segura**
- ✅ Parsing y validación de tokens JWT con `golang-jwt/jwt/v5`
- ✅ Verificación de firma y expiración
- ✅ Soporte para tokens de prueba en modo testing
- ✅ Validación de issuer para seguridad adicional

### **🛡️ Middleware Simplificado**
- ✅ `RequireAuth(jwtSecret)` - Autenticación básica
- ✅ `RequireRole(role)` - Verificación de roles específicos
- ✅ `RequireAdmin()` - Acceso solo para administradores
- ✅ Helper functions para obtener datos del usuario

### **🔄 Metadata del Token**
- ✅ Información completa del token disponible en contexto
- ✅ Tiempo de expiración y metadata adicional
- ✅ Soporte para testing con bypass de validación

---

## 📦 **Instalación**

```bash
# En tu proyecto principal o microservicio
go get github.com/tlaloc-dev/tlaloc-tokencheck-lib@v1.0.0
```

---

## 🎯 **Uso Básico**

### **En tu microservicio:**
```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/tlaloc-dev/tlaloc-tokencheck-lib/tokencheck"
)

func main() {
    e := echo.New()

    // Middleware de autenticación
    e.Use(tokencheck.RequireAuth("tu-jwt-secret-aqui"))
    
    // Rutas protegidas
    protected := e.Group("/api")
    protected.GET("/profile", handleProfile)
    protected.POST("/data", handleCreateData)
    
    // Rutas de admin
    admin := e.Group("/admin")
    admin.Use(tokencheck.RequireAdmin())
    admin.GET("/users", handleListUsers)
    
    e.Logger.Fatal(e.Start(":8080"))
}
```

### **Con Configuración Personalizada:**
```go
config := tokencheck.NewJWTConfig("tu-jwt-secret")
config.SkipValidation = false // Para testing
config.RequiredForTest = true
config.TestUserID = 1
config.TestEmail = "test@example.com"
config.TestRole = "admin"

e.Use(tokencheck.RequireAuthWithConfig(config))
```

---

## 🔧 **API Reference**

### **Middleware Principales**
```go
// Autenticación básica
e.Use(tokencheck.RequireAuth(jwtSecret))

// Con rol específico
e.Use(tokencheck.RequireRole("user_manager"))
e.Use(tokencheck.RequireRole("financial_analyst"))

// Solo administradores
e.Use(tokencheck.RequireAdmin())

// Con configuración personalizada
config := tokencheck.NewJWTConfig(jwtSecret)
e.Use(tokencheck.ValidateToken(config))
```

### **Helper Functions**
```go
// En tus handlers
func handleProfile(c echo.Context) error {
    userID, err := tokencheck.GetUserID(c)
    if err != nil {
        return c.JSON(401, "Unauthorized")
    }
    
    email, err := tokencheck.GetUserEmail(c)
    role, err := tokencheck.GetUserRole(c)
    
    return c.JSON(200, map[string]interface{}{
        "user_id": userID,
        "email":   email,
        "role":    role,
    })
}

// Verificar si está autenticado
if tokencheck.IsAuthenticated(c) {
    // Usuario válido
}

// Obtener metadata del token
if metadata, err := tokencheck.GetTokenInfo(c); err == nil {
    expiresIn := metadata["expires_in"].(int64)
    if expiresIn < 300 { // 5 minutos
        // Renovar token próximamente
    }
}
```

### **Manejo de Testing**
```go
// Generar token de prueba
testToken, err := tokencheck.GenerateTestToken(config)

// Middleware que permite testing
config.RequiredForTest = true
e.Use(tokencheck.RequireAuthWithConfig(config))

// En tu middleware, verificar si es testing mode
if tokencheck.IsTestMode(c) {
    // Permitir acceso sin validación completa
}
```

---

## 🔍 **Ejemplos de Implementación**

### **1. Microservicio Básico**
```go
package main

import (
    "github.com/labstack/echo/v4"
    tokencheck "github.com/tlaloc-dev/tlaloc-tokencheck-lib/tokencheck"
)

func main() {
    e := echo.New()
    
    // Autenticación para todas las rutas API
    api := e.Group("/api")
    api.Use(tokencheck.RequireAuth("your-super-secret-jwt-key"))
    
    // Rutas por rol
    userRoutes := api.Group("/user")
    userRoutes.GET("/profile", getUserProfile)
    
    adminRoutes := api.Group("/admin")
    adminRoutes.Use(tokencheck.RequireRole("admin"))
    adminRoutes.GET("/dashboard", getAdminDashboard)
    
    e.Logger.Fatal(e.Start(":8080"))
}

func getUserProfile(c echo.Context) error {
    userID, _ := tokencheck.GetUserID(c)
    return c.JSON(200, map[string]interface{}{
        "user_id": userID,
        "message": "User profile data",
    })
}
```

### **2. Configuración con Testing**
```go
package main

import (
    "github.com/labstack/echo/v4"
    tokencheck "github.com/tlaloc-dev/tlaloc-tokencheck-lib/tokencheck"
    "os"
)

func main() {
    jwtSecret := os.Getenv("JWT_SECRET")
    
    // Configuración con soporte para testing
    config := tokencheck.NewJWTConfig(jwtSecret)
    
    // En modo desarrollo, permitir testing sin validación
    if os.Getenv("APP_ENV") == "development" {
        config.SkipValidation = true
        config.RequiredForTest = true
        config.TestUserID = 1
        config.TestEmail = "dev@test.com"
        config.TestRole = "admin"
    }
    
    e := echo.New()
    e.Use(tokencheck.RequireAuthWithConfig(config))
    
    // Resto de la configuración...
}
```

---

## 🧪 **Testing**

### **Test con Token Manual**
```bash
# 1. Generar token de prueba
curl -X POST http://localhost:8080/auth/tokencheck/test \
  -H "Content-Type: application/json" \
  -d '{"jwt_secret":"your-secret"}'
```

### **Test con el Middleware**
```bash
# 2. Usar middleware en endpoint protegido
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer <token-valido>"
```

### **Test de Errores Esperados**
```bash
# Token ausente
curl -X GET http://localhost:8080/api/profile
# Expected: 401 Unauthorized

# Token malformado
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: InvalidToken"
# Expected: 400 Bad Request

# Token expirado
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer <token-expirado>"
# Expected: 401 Unauthorized
```

---

## 🔄 **Beneficios Sobre AuthMiddleware Copiado**

### **✅ Actualización Centralizada**
- Cambias la validación en **un solo lugar** y se actualiza en todos tus servicios automáticamente
- Sin necesidad de actualizar 3 archivos diferentes

### **✅ Sin Código Duplicado**
- Eliminas inconsistencias entre servicios
- Menos riesgo de errores de configuración

### **✅ Mantenimiento Simplificado**
- Un punto único para bugs de autenticación
- Validación consistente en todos los microservicios

### **✅ Testing Unificado**
- El mismo código de testing para todos los servicios
- Facilita debugging y desarrollo

---

## 🔐 **Seguridad Implementada**

- ✅ **Validación robusta**: JWT estándar con verificación de firma y expiración
- ✅ **Prevención de ataques**: Constant-time comparison para tokens
- ✅ **Logging estructurado**: Información consistente de errores
- ✅ **Modo testing**: Soporte para desarrollo con bypass de validación
- ✅ **CORS Headers**: Headers consistentes para comunicación cross-origin

---

## 📚 **Versionamiento**

- **v1.0.0** - Versión inicial con validación JWT básica
- **Roadmap**: Soporte para refresh tokens, validación de device fingerprinting, rate limiting

---

## 🤝 **Soporte**

Para issues, preguntas o sugerencias:
- GitHub Repository: `github.com/tlaloc-dev/tlaloc-tokencheck-lib`
- Issues: Crear un issue en el repositorio principal del proyecto
- Comunidad: Contribuciones welcome bajo el código de conducta del proyecto

---

*Esta librería fue diseñada para resolver el problema de la duplicación de código de autenticación entre múltiples microservicios Tlaloc.*