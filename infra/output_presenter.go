package infra

import (
    "html/template"
    "os"
    "fifaclubscup/domain"
)

type OutputPresenter struct {
    HTMLPath string
}

func (p *OutputPresenter) Present(t domain.Tournament) error {
    // p.printToConsole(t)
    if p.HTMLPath != "" {
        return p.renderHTML(t)
    }
    return nil
}

// func (p *OutputPresenter) printToConsole(t domain.Tournament) {
//     fmt.Println("=== FIFA World Cup Simulation ===")
//     fmt.Println("\n-- Group Stage --")
//     for _, group := range t.Groups {
//         fmt.Printf("Group %s:\n", group.Name)
//         for _, team := range group.Teams {
//             fmt.Printf("  %s\n", team.Name)
//         }
//         fmt.Println()
//     }

//     fmt.Println("-- Knockout Stage --")
//     for _, m := range t.Finals {
//         fmt.Println(m.Summary())
//     }

//     fmt.Printf("\nChampion: %s\n", t.Champion.Name)
// }

func (p *OutputPresenter) renderHTML(t domain.Tournament) error {
    const tmpl = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>FIFA World Cup Results</title>
  </head>
  <body>
    <h1>FIFA World Cup Simulation</h1>
    <h2>Group Stage</h2>
    {{range .Groups}}
      <h3>Group {{.Name}}</h3>
      <ul>
        {{range .Teams}}<li>{{.Name}}</li>{{end}}
      </ul>
    {{end}}
    <h2>Knockout Stage</h2>
    <ul>
      {{range .Finals}}
        <li>{{summary .}}</li>
      {{end}}
    </ul>
    <h2>Champion</h2>
    <p>{{.Champion.Name}}</p>
  </body>
</html>`

    tpls, err := template.New("report").Funcs(template.FuncMap{
        "summary": func(m *domain.Match) string { return m.Summary() },
    }).Parse(tmpl)
    if err != nil {
        return err
    }

    file, err := os.Create(p.HTMLPath)
    if err != nil {
        return err
    }
    defer file.Close()

    return tpls.Execute(file, t)
}


