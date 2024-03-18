# Dokumentation des Go-Projekts

## Kubectl commands



```bash
kubectl apply -f auth-service-pod.yaml
```
```bash
kubectl get pod product-service   
```

IPs:

```bash
kubectl get service
```

```bash
kubctl apply - f manifest/deployment.yaml
```

Delete deployments

```bash
kubectl delete deployments.app nginx-deployment
```

## Auth Paket

Das `auth` Paket bietet Funktionen für die Authentifizierung von Benutzern.

### AuthLoginHandler

```go
func AuthLoginHandler(w http.ResponseWriter, r *http.Request)
```

AuthLoginHandler authentifiziert einen Benutzer und generiert ein Token.

#### Parameter

- `w http.ResponseWriter`: Die Antwortwriter-Schnittstelle für die HTTP-Antwort.
- `r *http.Request`: Ein HTTP-Anforderungsobjekt.

#### Rückgabewerte

Keine.

### AuthLogoutHandler

```go
func AuthLogoutHandler(w http.ResponseWriter, r *http.Request)
```

AuthLogoutHandler meldet einen Benutzer ab.

#### Parameter

- `w http.ResponseWriter`: Die Antwortwriter-Schnittstelle für die HTTP-Antwort.
- `r *http.Request`: Ein HTTP-Anforderungsobjekt.

#### Rückgabewerte

Keine.

## Internal Paket

Das `internal` Paket enthält interne Funktionen für die Anwendung.

### Product Struktur

```go
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}
```

Product repräsentiert ein Produkt mit einer eindeutigen ID, einem Namen und einem Preis.

### ProductListHandler

```go
func ProductListHandler(w http.ResponseWriter, r *http.Request)
```

ProductListHandler gibt eine Liste aller Produkte zurück.

#### Parameter

- `w http.ResponseWriter`: Die Antwortwriter-Schnittstelle für die HTTP-Antwort.
- `r *http.Request`: Ein HTTP-Anforderungsobjekt.

#### Rückgabewerte

Keine.

### ProductDetailHandler

```go
func ProductDetailHandler(w http.ResponseWriter, r *http.Request)
```

ProductDetailHandler gibt die Details eines bestimmten Produkts zurück.

#### Parameter

- `w http.ResponseWriter`: Die Antwortwriter-Schnittstelle für die HTTP-Antwort.
- `r *http.Request`: Ein HTTP-Anforderungsobjekt.

#### Rückgabewerte

Keine.

### CheckoutPlaceOrderHandler

```go
func CheckoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request)
```

CheckoutPlaceOrderHandler verarbeitet die Bestellung eines Benutzers.

#### Parameter

- `w http.ResponseWriter`: Die Antwortwriter-Schnittstelle für die HTTP-Antwort.
- `r *http.Request`: Ein HTTP-Anforderungsobjekt.

#### Rückgabewerte

Keine.

## Pkg Paket

Das `pkg` Paket enthält Hilfsfunktionen und -strukturen.

### CreateToken

```go
func CreateToken(username string) (string, error)
```

CreateToken generiert ein JWT-Token für den angegebenen Benutzernamen.

#### Parameter

- `username string`: Der Benutzername für den das Token erstellt werden soll.

#### Rückgabewerte

- `string`: Das generierte JWT-Token.
- `error`: Ein Fehler, falls das Token nicht erstellt werden konnte.

### VerifyToken

```go
func VerifyToken(tokenString string) bool
```

VerifyToken überprüft, ob das gegebene JWT-Token gültig ist.

#### Parameter

- `tokenString string`: Das zu überprüfende JWT-Token.

#### Rückgabewerte

- `bool`: true, wenn das Token gültig ist, sonst false.

## Dockerfile

This Dockerfile is used to build a Docker image for running the Go application. It consists of two stages: the build stage and the production stage.

### Build Stage

- **Base Image**: This stage starts from a base image containing the Go runtime (`golang:1.22.0`).
- **Working Directory**: Sets the current working directory inside the container to `/app`.
- **Copy Dependencies**: Copies the `go.mod` and `go.sum` files into the working directory.
- **Download Dependencies**: Downloads all dependencies using `go mod download`.
- **Copy Source Code**: Copies the source code from the current directory into the working directory.
- **Build Application**: Builds the Go application using `CGO_ENABLED=0 go build -o main cmd/main.go`.

### Production Stage

- **Base Image**: This stage starts from a lightweight Alpine Linux image (`alpine:3.19.1`).
- **Copy Executable**: Copies the built executable from the build stage into the production stage.
- **Expose Port**: Exposes port 8080 to allow external connections to the application.
- **Entry Point**: Defines the entry point for the container as `./main`, which runs the built executable.

This Dockerfile enables easy deployment and distribution of the Go application in a containerized environment.
