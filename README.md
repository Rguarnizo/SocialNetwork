# Go Microservices SocialMedia

Arquitectura y codigo para prueba tecnica en Golang, prototipo RedSocial.

## Servicios

- auth-service: autenticación y JWT
- profile-service: perfiles de usuario
- post-service: publicaciones y likes

## Despliegue y ejecución.

Ejercutar

```
docker compose up --build
```

Para probar el flujo end to end, ejecutar para Windows(PowerShell):

```
.\FlujoE2E.ps1
```

Para Linux:

```
chmod +x FlujoE2E.sh
./FlujoE2E.sh
```

Para ver la demostración del flujo da click en el siguiente enlace

[Video De Demostración Flujo end to end](https://drive.google.com/file/d/1WW-pFxto6aWS8-6_3z2UCkF9ZCP_uvdd/view?usp=sharing)

## Rutas principales

### Auth Service

GET /login

### Profile Service

GET /profile/me

### Post Service

POST /posts
GET /posts
POST /posts/{id}/like

## Estructura de Carpetas y Archivos.

Social Media tiene los 3 microservicios y un paquete transversar de Autenticación/Autorizacion, cada microservicio tiene su Dockerfile para despliegue.

Cada microservicio tiene su handler para el manejo de peticiones entrantes, su router para definir los Path de las peticiones, su repositorio para definir las operaciones a la base de datos, y su servicio para abstraer la funcionalidad de la aplicación.
