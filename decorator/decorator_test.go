package decorator

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient of the pizza decorator it must return the text %s the expected text, not '%s'", pizzaResult, expectedText)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{&PizzaDecorator{}}
	onionResult, err := onion.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(onionResult, "onion") {
		t.Errorf("When calling the add ingredient of the onion decorator it "+"must return a text with the word 'onion', not '%s'", onionResult)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the meat decorator without"+"an IngredientAdd in its Ingredient field must returne an error,"+"not a string with '%s'", meatResult)
	}

	meat = &Meat{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(meatResult, "meat") {
		t.Errorf("When calling the add ingredient of the meat decorator it"+"must return a text with the word 'meat', not '%s'", meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	expectedText := "Pizza with the following ingredients: meat, onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and meat the returned "+"string must contain the text '%s' but '%s' didn't have it", expectedText, pizzaResult)
	}
	t.Log(pizzaResult)
}

func TestMyServer_ServeHTTP(t *testing.T) {
	http.Handle("/", &MyServer{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TestLoggerServer_ServeHTTP(t *testing.T) {
	http.Handle("/", &LoggerServer{
		LogWriter: os.Stdout,
		Handler:   &MyServer{},
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
