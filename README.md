# Dokumentation des Go-Projekts

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
