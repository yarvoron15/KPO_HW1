package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"vyshkat/internal/domain"
	"vyshkat/internal/service"
)

type CLI struct{ fleet *service.FleetService }

func NewCLI(fleet *service.FleetService) *CLI { return &CLI{fleet: fleet} }
func (c *CLI) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		c.menu()
		choice, ok := c.line(scanner, "Выберите пункт: ")
		if !ok {
			return
		}
		switch choice {
		case "1":
			c.addVehicle(scanner, "scooter")
		case "2":
			c.addVehicle(scanner, "electric-bike")
		case "3":
			c.addVehicle(scanner, "regular-bike")
		case "4":
			c.addHelmet(scanner)
		case "5":
			c.addStation(scanner)
		case "6":
			c.summary()
		case "7":
			c.beginnerReport()
		case "8":
			c.inventoryReport()
		case "9":
			c.demo()
		case "0":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Некорректный пункт меню.")
		}
	}
}
func (c *CLI) menu() {
	fmt.Println("\n=== ВышКат — флот ВШЭ ФКН ===")
	fmt.Println("1. Добавить электрический самокат")
	fmt.Println("2. Добавить электрический велосипед")
	fmt.Println("3. Добавить механический велосипед")
	fmt.Println("4. Добавить шлем")
	fmt.Println("5. Добавить зарядную станцию")
	fmt.Println("6. Отчёт: размер флота и энергопотребление")
	fmt.Println("7. Отчёт: транспорт для новичков (балл >= 6)")
	fmt.Println("8. Полный реестр инвентаря")
	fmt.Println("9. Загрузить демо-набор")
	fmt.Println("0. Выход")
}
func (c *CLI) line(s *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)
	if !s.Scan() {
		return "", false
	}
	return strings.TrimSpace(s.Text()), true
}
func (c *CLI) number(s *bufio.Scanner, prompt string) (float64, bool) {
	value, ok := c.line(s, prompt)
	if !ok {
		return 0, false
	}
	n, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", "."), 64)
	if err != nil || n < 0 {
		fmt.Println("Введите неотрицательное число.")
		return 0, false
	}
	return n, true
}
func (c *CLI) score(s *bufio.Scanner) (int, bool) {
	value, ok := c.line(s, "Оценка для новичка (1-10): ")
	if !ok {
		return 0, false
	}
	n, err := strconv.Atoi(value)
	if err != nil || n < 1 || n > 10 {
		fmt.Println("Оценка должна быть от 1 до 10.")
		return 0, false
	}
	return n, true
}
func (c *CLI) addVehicle(s *bufio.Scanner, kind string) {
	model, ok := c.line(s, "Модель: ")
	if !ok {
		return
	}
	number, ok := c.line(s, "Инвентарный номер: ")
	if !ok {
		return
	}
	score, ok := c.score(s)
	if !ok {
		return
	}
	var (
		v   domain.RentalVehicle
		err error
	)
	if kind == "regular-bike" {
		v, err = domain.NewRegularBike(number, model, score)
	} else {
		energy, valid := c.number(s, "Ежедневное потребление, кВт·ч: ")
		if !valid {
			return
		}
		if kind == "scooter" {
			v, err = domain.NewElectricScooter(number, model, energy, score)
		} else {
			v, err = domain.NewElectricBike(number, model, energy, score)
		}
	}
	if err != nil {
		fmt.Println("Ошибка валидации параметров:", err)
		return
	}
	accepted, reason := c.fleet.AddVehicle(v)
	if accepted {
		fmt.Println("Транспорт добавлен:", reason)
	} else {
		fmt.Println("Транспорт отклонён:", reason)
	}
}
func (c *CLI) addHelmet(s *bufio.Scanner) {
	number, ok := c.line(s, "Инвентарный номер: ")
	if !ok {
		return
	}
	size, ok := c.line(s, "Размер: ")
	if !ok {
		return
	}
	added, reason := c.fleet.AddItem(domain.NewHelmet(number, size))
	if added {
		fmt.Println("Шлем добавлен.")
	} else {
		fmt.Println("Шлем отклонён:", reason)
	}
}
func (c *CLI) addStation(s *bufio.Scanner) {
	number, ok := c.line(s, "Инвентарный номер: ")
	if !ok {
		return
	}
	energy, ok := c.number(s, "Ежедневное потребление, кВт·ч: ")
	if !ok {
		return
	}
	station, err := domain.NewChargingStation(number, energy)
	if err != nil {
		fmt.Println("Не удалось создать станцию:", err)
		return
	}
	added, reason := c.fleet.AddItem(station)
	if added {
		fmt.Println("Зарядная станция добавлена.")
	} else {
		fmt.Println("Зарядная станция отклонена:", reason)
	}
}
func (c *CLI) summary() {
	fmt.Printf("\nТранспортных средств: %d\nВсего позиций инвентаря: %d\nСуточное энергопотребление: %.2f кВт·ч\n", c.fleet.GetVehiclesCount(), c.fleet.GetItemsCount(), c.fleet.GetTotalDailyEnergy())
}
func (c *CLI) beginnerReport() {
	vehicles := c.fleet.GetVehiclesForBeginners()
	fmt.Println("\nТранспорт для новичков:")
	if len(vehicles) == 0 {
		fmt.Println("(нет подходящих транспортных средств)")
		return
	}
	fmt.Printf("%-28s %-16s %s\n", "Название", "Инв. номер", "Балл")
	for _, v := range vehicles {
		fmt.Printf("%-28s %-16s %d\n", v.GetName(), v.GetInventoryNumber(), v.GetBeginnerScore())
	}
}
func (c *CLI) inventoryReport() {
	fmt.Println("\nПолный реестр:")
	fmt.Printf("%-4s %-28s %s\n", "№", "Название", "Инвентарный номер")
	for i, item := range c.fleet.GetAllItems() {
		fmt.Printf("%-4d %-28s %s\n", i+1, item.GetName(), item.GetInventoryNumber())
	}
}
func (c *CLI) demo() {
	scooters := []struct {
		n, m string
		e    float64
		s    int
	}{{"SC-001", "Xiaomi Pro 2", 1.8, 8}, {"SC-002", "Ninebot Max", 2.1, 7}, {"SC-003", "Kugoo M4", 2.4, 5}}
	for _, d := range scooters {
		v, err := domain.NewElectricScooter(d.n, d.m, d.e, d.s)
		if err != nil {
			fmt.Println("Ошибка демо-набора:", err)
			continue
		}
		c.fleet.AddVehicle(v)
	}
	bikes := []struct {
		n, m string
		e    float64
		s    int
	}{{"EB-001", "Cube Touring", 3.2, 9}, {"EB-002", "Merida eSPRESSO", 2.8, 6}}
	for _, d := range bikes {
		v, err := domain.NewElectricBike(d.n, d.m, d.e, d.s)
		if err != nil {
			fmt.Println("Ошибка демо-набора:", err)
			continue
		}
		c.fleet.AddVehicle(v)
	}
	c.fleet.AddItem(domain.NewHelmet("H-001", "M"))
	c.fleet.AddItem(domain.NewHelmet("H-002", "L"))
	station, err := domain.NewChargingStation("CS-001", 4.5)
	if err == nil {
		c.fleet.AddItem(station)
	}
	fmt.Println("Демо-набор загружен: 3 самоката, 2 электровелосипеда, 2 шлема, 1 зарядная станция.")
}
