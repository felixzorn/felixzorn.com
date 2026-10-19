# deinname.dev

Mein Portfolio als **Inselreise**: Ein Boot fährt zu fünf Inseln, eine kleine Figur geht an Land, und jede Insel ist ein Kapitel (Über mich, Projekte, Erfahrung, Skills, Kontakt). Licht und Stimmung folgen der Uhrzeit des Besuchers.

Ausgeliefert wird alles von einem einzigen Go-Binary, die Seite ist per `go:embed` eingebaut.

## Sprache

Standard ist Englisch. Über den Umschalter `EN | DE` oben rechts wechselst du auf Deutsch. Die Wahl steht im Cookie `lang`.

## Starten

```sh
go run .
# → http://localhost:8080
```

Voraussetzung: Go 1.22 oder neuer.

## Routen

| Route | Antwort |
|---|---|
| `GET /` | die Inselreise (`web/index.html`) |
| `GET /healthz` | `ok` plus Uptime |

## Projektstruktur

```
.
├── main.go              Go-Server (liefert web/ aus, Healthcheck)
├── go.mod
├── Dockerfile           Image auf Basis von scratch
└── web/
    ├── index.html       Inselreise (three.js)
    └── vendor/
        └── three.min.js three.js r128 (MIT-Lizenz), lokal statt CDN
```

## Inselreise bedienen

- **Steuern:** WASD oder Pfeiltasten. Klick oder Tippen aufs Wasser fährt automatisch dorthin.
- **Anlegen:** Klick auf ein Holzschild, die Seekarte oder die Navigation unten. Fährt man selbst an einen Steg heran, erscheint „⚓ An Land gehen“ (oder Leertaste/E/Enter).
- **Ablegen:** „⚓ Ablegen“ im Panel, Esc, eine Steuertaste oder Klick aufs Wasser.
- **Tageszeit:** folgt der Ortszeit des Besuchers. „Tageszeit ändern“ schaltet die Stimmungen zur Vorschau durch.
- **Klassische Ansicht:** alle Inhalte als schlichte Seite, z. B. für Recruiter oder ohne WebGL.

## Inhalte anpassen

Alles ist noch mit Beispielinhalten gefüllt.

1. **`web/index.html`**, Objekt `CONTENT` ganz oben im Script: je Sprache (`de` und `en`) Name, Rolle, Intro, Texte, Projekte, Erfahrung, Skills.
   E-Mail, Links und die Farben deiner Figur stehen in `SHARED` und gelten für beide Sprachen (Haut, Haare, Pulli, Hose unter `SHARED.me`).
2. **`web/index.html`**, Wörterbuch `I18N`: Oberflächentexte, Inselnamen, Seitentitel und Beschreibung (`title`, `descr`) je Sprache.
3. Zum Schluss nach `Dein Name`, `deinname`, `example.com`, `Beispiel` und `Muster` suchen, damit nichts übrig bleibt. Dabei beide Sprachen prüfen.

Die GitHub- und LinkedIn-Links zeigen aktuell nur auf die Startseiten, dort deine echten Profil-URLs eintragen.

## Vor dem Livegang: Google Fonts

Die Seite lädt Schriften von Google Fonts. In Deutschland gilt das Einbinden von Google Fonts ohne Einwilligung als datenschutzrechtlich problematisch (IP-Adresse geht an Google).
Besser: die Schriften selbst hosten, z. B. mit [google-webfonts-helper](https://gwfh.mranftl.com/) herunterladen, nach `web/fonts/` legen und den `<link>` zu Google durch eigene `@font-face`-Regeln ersetzen.
Verwendet werden Gabarito, Figtree, JetBrains Mono und Kalam.

## Bauen und deployen

Einzelnes Binary:

```sh
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o deinname-dev .
ADDR=:8080 ./deinname-dev
```

Für einen Linux-Server von macOS oder Windows aus:

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o deinname-dev .
```

Docker:

```sh
docker build -t deinname-dev .
docker run -p 8080:8080 deinname-dev
```

HTTPS übernimmt am einfachsten ein Reverse Proxy davor, zum Beispiel Caddy:

```
deinname.dev {
    reverse_proxy localhost:8080
}
```

## Ideen für später

- Figur persönlicher machen (Brille, Bart, Mütze, Haustier an Bord)
- Schilder an den Projekthäusern, Kran bewegt Kisten
- Springende Fische, Wetter, eine Nebelinsel als Easter Egg
