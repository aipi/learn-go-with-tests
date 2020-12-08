package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

const ultimaPalavra = "Vai!"
const inicioContagem = 3
const escrita = "escrita"
const pausa = "pausa"

func Contagem(saida io.Writer, sleeper Sleeper) {
    for i := inicioContagem; i > 0; i-- {
        sleeper.Sleep()
        fmt.Fprintln(saida, i)
    }

    sleeper.Sleep()
    fmt.Fprint(saida, ultimaPalavra)
}

type SpyContagemOperacoes struct {
    Chamadas []string
}

func (s *SpyContagemOperacoes) Sleep() {
    s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
    s.Chamadas = append(s.Chamadas, escrita)
    return
}

type SleeperConfiguravel struct {
    duracao time.Duration
    sleep   func(time.Duration)
}

func (s *SleeperConfiguravel) Sleep() {
    s.sleep(s.duracao)
}

type TempoSpy struct {
    duracaoPausa time.Duration
}

func (t *TempoSpy) Sleep(duracao time.Duration) {
    t.duracaoPausa = duracao
}

func main() {
    sleeper := &SleeperConfiguravel{10 * time.Second, time.Sleep}
    Contagem(os.Stdout, sleeper)
}

type Sleeper interface {
    Sleep()
}

type SleeperPadrao struct {}

func (d *SleeperPadrao) Sleep() {
    time.Sleep(1 * time.Second)
}
