package main

type PizzaServicer interface {
	getPrice() int
}

type VeggeMania struct {
	price int
}

func (v *VeggeMania) getPrice() int {
	return v.price
}

type TomatoTopping struct {
	toppingPrice int
	pizza        PizzaServicer
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + t.toppingPrice
}

type CheeseTopping struct {
	toppingPrice int
	pizza        PizzaServicer
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + c.toppingPrice
}
