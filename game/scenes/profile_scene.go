package scenes

import (
	"fmt"
	"image/color"

	"github.com/allanjose001/go-battleship/game/components"
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/allanjose001/go-battleship/game/state"
	"github.com/allanjose001/go-battleship/internal/service"
	"github.com/hajimehoshi/ebiten/v2"
)

// ProfileScene representa a tela de perfil do jogador.
type ProfileScene struct {
	state   *state.GameState
	profile *service.Profile
	root    *components.Container // O container pai que envolve toda a cena.
}

func NewProfileScene(s *state.GameState) *ProfileScene {
	// Recupera os dados do jogador do serviço
	p, _ := service.FindProfile("malub")
	
	// Define o tamanho lógico da tela (coordenadas internas)
	sw, sh := float32(1280), float32(720)
	screenSize := basic.Size{W: sw, H: sh}

	// --- COLUNA PRINCIPAL CENTRALIZADA ---
	// Coluna principal que centraliza verticalmente
	mainColumn := components.NewColumn(
		basic.Point{},
		40, 
		basic.Size{W: sw, H: sh},
		basic.Center,
		basic.Center, 
		[]components.Widget{},
	)

	// --- SEÇÃO 1: TÍTULO ---
	titleDiv := components.NewContainer(
		basic.Point{X: 0, Y: 0},
		basic.Size{W: 100, H: 80},
		0, nil,
		basic.Center, basic.Center,
		components.NewText(basic.Point{}, "PERFIL DE JOGADOR", colors.White, 42),
		nil,
	)

	// --- SEÇÃO 2: STATUS (Cards de Informação) ---
	statusDiv := components.NewContainer(
		basic.Point{X: 0, Y: 100},
		basic.Size{W: 750, H: 100},
		0, nil,
		basic.Center, basic.Center,
		components.NewRow(
			basic.Point{}, 40,
			basic.Size{W: 800, H: 150},
			basic.Center, basic.Center,
			[]components.Widget{
				createStatCard("Partidas", fmt.Sprintf("%d", p.GamesPlayed), 200, 120),
				createStatCard("Vitórias", fmt.Sprintf("%d", 10), 200, 120),
				createStatCard("Taxa de vitória", fmt.Sprintf("%d", 30), 200, 120),
			},
		),
		nil,
	)

	// --- SEÇÃO 3: MURAL DE MEDALHAS ---
	// Container para agrupar título do mural + grid de medalhas
	muralContainer := components.NewContainer(
		basic.Point{},
		basic.Size{W: sw, H: 250},
		0, nil,
		basic.Center, basic.Center,
		nil,
		nil,
	)

	// Título do Mural
	muralLabelDiv := components.NewContainer(
		basic.Point{},
		basic.Size{W: sw, H: 50},
		0, nil,
		basic.Center, basic.Start,
		components.NewText(basic.Point{}, "MURAL DE MEDALHAS", colors.White, 28),
		nil,
	)

	// Mock das medalhas - Agora com mais medalhas e grid responsivo
	medalData := []struct {
		Icon, Title, Desc string
	}{
		{"🥇", "VETERANO", "10+ Partidas"},
		{"🎯", "SNIPER", "90% Precisão"},
		{"⚡", "VELOZ", "Vitória em <5min"},
		{"🛡️", "IMPENETRÁVEL", "0 acertos sofridos"},
	}

	// Grid de medalhas com 2 colunas
	medalGrid := components.NewColumn(
		basic.Point{},
		20,
		basic.Size{W: 900, H: 200},
		basic.Center,
		basic.Center,
		createMedalRows(medalData, 900),
	)

	// Coluna interna do mural (título + grid)
	muralInnerColumn := components.NewColumn(
		basic.Point{},
		15,
		basic.Size{W: sw, H: 250},
		basic.Center,
		basic.Center,
		[]components.Widget{muralLabelDiv, medalGrid},
	)
	muralContainer.Child = muralInnerColumn

	// --- SEÇÃO 4: BOTÃO RETORNAR ---
	buttonDiv := components.NewContainer(
		basic.Point{},
		basic.Size{W: sw, H: 100},
		0, nil,
		basic.Center, basic.Center,
		components.NewButton(
			basic.Point{}, basic.Size{W: 240, H: 60},
			"RETORNAR", color.RGBA{45, 67, 103, 255}, colors.White,
			func(b *components.Button) { fmt.Println("Retornar ao menu") },
		),
		nil,
	)

	// Adiciona todos os componentes à coluna principal com espaçamento generoso
	mainColumn.Children = []components.Widget{
		titleDiv,
		statusDiv,
		muralContainer,
		buttonDiv,
	}

	// Recalcula o alinhamento da coluna principal
	mainColumn = components.NewColumn(
		basic.Point{},
		50, 
		screenSize,
		basic.Center, 
		basic.Center, 
		mainColumn.Children,
	)

	// O root é o container principal que desenha o fundo escuro da tela
	root := components.NewContainer(
		basic.Point{}, screenSize, 0,
		color.RGBA{10, 25, 40, 255}, 
		basic.Center,            
		basic.Center,            
		mainColumn, nil,
	)

	return &ProfileScene{state: s, profile: p, root: root}
}

func createMedalRows(data []struct{ Icon, Title, Desc string }, containerWidth float32) []components.Widget {
	var rows []components.Widget
	
	// Tamanho fixo para os cards de medalha
	cardWidth := float32(400)
	cardHeight := float32(90)
	
	// Itera sobre as medalhas em grupos de 2
	for i := 0; i < len(data); i += 2 {
		var rowWidgets []components.Widget
		
		// Adiciona até 2 medalhas por linha
		for j := 0; j < 2 && (i+j) < len(data); j++ {
			m := data[i+j]
			medalCard := createMedalCard(m.Icon, m.Title, m.Desc, cardWidth, cardHeight)
			rowWidgets = append(rowWidgets, medalCard)
		}
		
		// Cria uma linha com as medalhas centralizadas
		row := components.NewRow(
			basic.Point{},
			30,
			basic.Size{W: containerWidth, H: cardHeight + 10},
			basic.Center,
			basic.Center,
			rowWidgets,
		)
		
		rows = append(rows, row)
	}
	
	return rows
}

// createStatCard encapsula a lógica de criar um balão de estatística centralizado.
func createStatCard(label, value string, w, h float32) *components.Container {
	labelTxt := components.NewText(basic.Point{}, label, colors.Black, 20)

	valueTxt := components.NewText(basic.Point{}, value, colors.Black, 25)	

	content := components.NewColumn(
		basic.Point{X: 0, Y: 0}, 20,
		basic.Size{W: 50, H: 75},
		basic.Center,
		basic.Center,
		[]components.Widget{labelTxt, valueTxt},
	)
	
	// Container branco com sombra suave
	return components.NewContainer(
		basic.Point{},
		basic.Size{W: w, H: h},
		15,
		color.RGBA{255, 255, 255, 255},
		basic.Center,
		basic.Center,
		content,
		nil,
	)
}

// createMedalCard encapsula a lógica de criar um balão de medalha (ícone + texto).
func createMedalCard(icon, title, desc string, w, h float32) *components.Container {
	// Ícone maior e mais visível
	iconTxt := components.NewText(basic.Point{}, icon, color.RGBA{255, 200, 0, 255}, 28)
	
	// Título em negrito
	titleTxt := components.NewText(basic.Point{}, title, color.RGBA{40, 40, 50, 255}, 16)
	
	// Descrição
	descTxt := components.NewText(basic.Point{}, desc, color.RGBA{100, 100, 110, 255}, 12)

	textCol := components.NewColumn(
		basic.Point{}, 4,
		basic.Size{W: w * 0.6, H: h},
		basic.Center,
		basic.Start,
		[]components.Widget{titleTxt, descTxt},
	)
	
	content := components.NewRow(
		basic.Point{}, 15,
		basic.Size{W: w - 30, H: h},
		basic.Center,
		basic.Center,
		[]components.Widget{iconTxt, textCol},
	)
	
	return components.NewContainer(
		basic.Point{},
		basic.Size{W: w, H: h},
		12,
		color.RGBA{255, 255, 255, 255},
		basic.Center,
		basic.Center,
		content,
		nil,
	)
}

// Implementações do contrato Scene
func (s *ProfileScene) OnEnter(prev Scene, size basic.Size) {
	// Atualiza os dados do perfil ao entrar na cena
	p, _ := service.FindProfile("malub")
	s.profile = p
}

func (s *ProfileScene) OnExit(next Scene) {}

// Update propaga a atualização de baixo para cima na árvore de componentes.
func (s *ProfileScene) Update() error {
	s.root.Update(basic.Point{X: 0, Y: 0})
	return nil
}

// Draw renderiza recursivamente toda a cena.
func (s *ProfileScene) Draw(screen *ebiten.Image) {
	s.root.Draw(screen)
}