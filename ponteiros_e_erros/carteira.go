package main

import "fmt"
import "errors"

type Stringer interface {
        String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
    saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
    c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
    return c.saldo
}

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (c *Carteira) Retirar(quantidade Bitcoin) error {

    if quantidade > c.saldo {
        return ErroSaldoInsuficiente
    }

    c.saldo -= quantidade
    return nil
}
