package main

import (
	"fmt"
)

type World struct {
	Areas []Area
}

type Area struct {
	Name        string
	Description string
	Enemies     []Enemy
	Missions    []Mission
}

type Character struct {
	Name    string
	attack  int
	defense int
	Level   int
	role    string
	Health  int
}

type Enemy struct {
	Name   string
	Level  int
	Health int
	Damage int
}

type Mission struct {
	Name        string
	Description string
	Targets     map[string]int // Mapa de enemigos a eliminar y la cantidad requerida
	Rewards     []Reward
}

type Reward struct {
	Type  string
	Value int
}

// explorarar un área
func (player *Character) Explore(area Area) {
	fmt.Printf("%s está explorando el área %s...\n", player.Name, area.Name)
	// encontrar objetos, enfrentar enemigos aleatorios, etc.
}

// luchar contra un enemigo
func (player *Character) Fight(enemy Enemy) {
	fmt.Printf("%s está luchando contra el enemigo %s...\n", player.Name, enemy.Name)
	// calcular daños, actualizaciones de salud, etc.
}

// completar una misión
func (player *Character) CompleteMission(mission Mission) {
	fmt.Printf("%s está completando la misión %s...\n", player.Name, mission.Name)
	//  objetivos de la misión y otorgar recompensas.
}

// Función principal
func main() {
	//  mundo de ejemplo con una área y un personaje
	world := World{
		Areas: []Area{
			{
				Name:        "Bosque",
				Description: "Un oscuro y misterioso bosque lleno de peligros.",
				Enemies: []Enemy{
					{
						Name:   "Lobo",
						Level:  3,
						Health: 30,
						Damage: 5,
					},
				},
				Missions: []Mission{
					{
						Name:        "Cazar Lobos",
						Description: "Elimina a 5 lobos del bosque.",
						Targets: map[string]int{
							"Lobo": 5,
						},
						Rewards: []Reward{
							{
								Type:  "Experiencia",
								Value: 100,
							},
						},
					},
				},
			},
		},
	}

	player := Character{
		Name:   "Cazador",
		Level:  1,
		Health: 100,
	}

	fmt.Println("Bienvenido al mundo de aventuras!")
	fmt.Println("Estás en el área:", world.Areas[0].Name)
	fmt.Println("Descripción del área:", world.Areas[0].Description)
	fmt.Println("Personaje:", player.Name)
	fmt.Println("Nivel:", player.Level)

	// acciones que el jugador puede realizar
	player.Explore(world.Areas[0])
	player.Fight(world.Areas[0].Enemies[0])
	player.CompleteMission(world.Areas[0].Missions[0])
}
