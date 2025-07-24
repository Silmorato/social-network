# social-network

Una red social simplificada construida con Go. Permite a los usuarios publicar tweets, seguir a otros y visualizar una línea de tiempo personalizada.

## 🚀 Tecnologías

- **Go**: Backend principal.
- **GORM + PostgreSQL**: Persistencia de datos.
- **Testify**: Tests unitarios.
- **Arquitectura limpia (Clean Architecture)**: Separación por capas y responsabilidades.

---

## 🧱 Arquitectura

La estructura está organizada en capas siguiendo los principios de Clean Architecture:

- `cmd/api`: Punto de entrada de la aplicación. Define el `main.go`, el `router` y las dependencias.
- `internal/adapters/database`: Implementaciones de repositorios (User, Tweet, Follow) y conexión a la base de datos.
- `internal/adapters/http`: Handlers HTTP expuestos como endpoints REST.
- `internal/domain`: Entidades principales del negocio: Tweet, Follow, User.
- `internal/dto`: Requests y responses para la API.
- `internal/errors`: Definición de errores personalizados.
- `internal/ports`: Interfaces que definen contratos entre capas.
- `internal/services`: Lógica de negocio.
- `scripts`: Scripts útiles como el de inicialización de la base de datos.

### 🖼️ Diagrama de arquitectura

![Diagrama de Arquitectura](./docs/images/architecture-diagram.png)]

## 📬 Endpoints principales

| Método | Endpoint       | Descripción                     |
|--------|----------------|---------------------------------|
| POST   | `/tweets`      | Publicar un nuevo tweet         |
| GET    | `/timeline`    | Obtener la línea de tiempo      |
| POST   | `/follow`      | Seguir a otro usuario           |

---

## 🧪 Tests

Este proyecto incluye tests unitarios utilizando Testify.

```bash
go test ./... -v