package main

import "fmt"

type dog struct{
    weight int
}

func (d dog) getWeight() int{
    return d.weight
}

func (d dog) String() string{
    return "dog"
}

func (d dog) feedingRate() int {
    return 2
}

func (d dog) monthlyFoodWeight() int{
    return d.weight * d.feedingRate()
}

type cat struct{
    weight int
}

func (c cat) getWeight() int{
    return c.weight
}

func (c cat) String() string{
    return "cat"
}

func (c cat) feedingRate() int {
    return 7
}

func (c cat) monthlyFoodWeight() int{
    return c.weight * c.feedingRate()
}

type cow struct{
    weight int
}

func (m cow) getWeight() int{
    return m.weight
}

func (m cow) String() string{
    return "cow"
}

func (m cow) feedingRate() int {
    return 25
}

func (m cow) monthlyFoodWeight() int{
    return m.weight * m.feedingRate()
}

type animal interface{
    fmt.Stringer
    getWeight() int
    feedingRate() int
    monthlyFoodWeight() int
}

func farmData(farmAnimals []animal) int {

    var totalFarmFood int

    for _, val := range farmAnimals{
            typeAnimal := val.String()
            weightAnimal := val.getWeight()
            monthlyFoodWeight := val.monthlyFoodWeight()
            totalFarmFood += monthlyFoodWeight
            fmt.Printf("Type of animal is %s, weight equals %d kg, required quantity of food for this animal per month is %d kg \n", typeAnimal, weightAnimal, monthlyFoodWeight)
        }

    fmt.Printf("Total food for all farm animals equals per month: %d kg\n", totalFarmFood)

    return totalFarmFood
}

func main(){

    //create animals
    farmAnimals := []animal{
        dog{10}, cat{5}, cow{200}, cow{150}, cat{5}, cat{7}, dog{20}, dog{35}, cow{300}, cow{250}, cow{145},
    }

    farmData(farmAnimals)
}

