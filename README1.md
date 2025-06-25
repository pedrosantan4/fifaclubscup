# 🏆 FIFA Club World Cup Simulator — by pedrosantan4 ⚽

Bem-vindo ao simulador mais insano do **Mundial de Clubes FIFA** feito em **Go (Golang)**!  
Aqui você vê a magia dos clubes do mundo inteiro disputando a glória eterna — tudo isso com **Clean Architecture**, **paralelismo com goroutines**, e um toque de realismo que vai fazer você gritar “É TETRA!” 🇧🇷

---

## 🚀 O que esse projeto faz?

- 🛠️ **Simula uma Copa do Mundo de Clubes** com 32 times divididos em 8 grupos
- ⚽ **Partidas com lógica baseada na força dos times**
- 🔄 Fases de **grupos**, **oitavas**, **quartas**, **semi**, **final** e **3º lugar**
- 🧠 Usa **goroutines e channels** para paralelizar jogos em cada fase
- 📦 Persiste o histórico de campeões em um `.json`
- 📊 Exibe todos os campeões anteriores com placares finais e top 3

---

## 🧱 Arquitetura

Organizado com os princípios de **Clean Code**, **KISS**, **SOLID** e **DRY**:

```
fifaclubscup/
├── cmd/                 # Executável principal (main.go)
├── domain/              # Regras e entidades do negócio (Team, Match, Tournament)
├── internal/            # Orquestração do torneio (TournamentManager)
├── infra/               # Persistência (histórico .json)
│   └── data/history.json
└── go.mod
```

---

## ⚙️ Como rodar

> Pré-requisitos: Go instalado (1.20+)

```bash
git clone https://github.com/seu-usuario/fifaclubscup.git
cd fifaclubscup
go run ./cmd
```

---

## ✨ Exemplo de saída

```text
🌍 Iniciando o Mundial de Clubes FIFA...

══════════════════════════════════════════════════

🏁 Apresentação dos Grupos

Grupo A:
   • Porto (Portugal)
   • Santos (Brazil)
   • River Plate (Argentina)
   • Boca Juniors (Argentina)

📊 Classificação dos Grupos

════════════════════════════════════════════════════════════

📊 Classificação do Grupo A
───────────────────────────────────────────────────────
Pos  Clube                  Pts  SG  GM  GS
───────────────────────────────────────────────────────
 1º  Porto                   10    2    8    6
 2º  Santos                   9    2    6    4
 3º  Boca Juniors             7   -1    4    5
 4º  River Plate              6   -3    2    5
───────────────────────────────────────────────────────
...

Grande Final
   Santos x Club León

Disputa 3º Lugar
   Celtic x Ajax

══════════════════════════════════════════════════

🎬 Resumo do Torneio

🏆 FIFA Club World Cup - 2025
🥇 Campeão: Santos
   Santos x Club León
══════════════════════════════════════════════════

```

---

## 🧪 Em breve (roadmap)

- [ ] Web API com leaderboard ao vivo
- [ ] Ranking global dos clubes

---

## 🤖 Tecnologias

- **Go** — linguagem rápida, concisa e concorrente
- **JSON** — persistência leve e humana
- **Goroutines/Channels** — concorrência real para simular vários jogos ao mesmo tempo
- **Arquitetura limpa** — domínios separados, fácil de expandir

---

## 🧠 Conceitos aplicados

- SOLID & Clean Code
- Clean Architecture
- Goroutines & Channels
- Parsing e persistência com JSON
- Design modular e testável

sbc

---

## 🧔 Autor

**Pedro Santana**  
Data Engineer • Dev Golang • Torcedor fiel do bom design de software  
[LinkedIn](https://linkedin.com/in/pedrosantan4)

---

## ⚽ Bora simular?

> “FIFA aprovaria esse projeto se tivesse coragem.” — ChatGPT