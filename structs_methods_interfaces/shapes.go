package main

import "math"

type Retangulo struct {
    Largura  float64
    Altura float64
}

func (r Retangulo) Area() float64  {
    return r.Largura * r.Altura
}

type Circulo struct {
    Raio float64
}

func (c Circulo) Area() float64  {
    return math.Pi * c.Raio * c.Raio
}

type Triangulo struct {
    Base   float64
    Altura float64
}


func (t Triangulo) Area() float64  {
    return t.Base * t.Altura / 2
}


func Perimetro(rectangle Retangulo) float64 {
    return 2 * (rectangle.Largura + rectangle.Altura)
}

type Forma interface {
    Area() float64
}
