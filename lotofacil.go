package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maximoPorJogo int = 18
const minimoPorJogo int = 15
const numerosValidos string = "01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25"

var dezenasSorteadas = "20 01 05 03 15 06 07 08 25 21 23 22 18 09 10"

type jogo struct {
	dezenas    []string
	nroJogos   int
	nroAcertos int
	mensagem   string
	valido     bool
}

func (j *jogo) setDezena(d []string) {
	j.dezenas = d
	j.valido = true
}

func (j *jogo) validaMaxPorJogo() {
	if j.valido == true {
		if len(j.dezenas) > maximoPorJogo {
			j.mensagem = fmt.Sprintf("Jogo invalido, excedeu 18 números. Jogo com: %d números !", len(j.dezenas))
			j.nroJogos = len(j.dezenas)
			j.valido = false
		} else {
			j.mensagem = ""
			j.nroJogos = 0
		}
	}
}

func (j *jogo) validaMinPorJogo() {
	if j.valido == true {
		if len(j.dezenas) < minimoPorJogo {
			j.mensagem = fmt.Sprintf("Jogo invalido, menor do que 15 números. Jogo com: %d números !", len(j.dezenas))
			j.nroJogos = len(j.dezenas)
			j.valido = false
		} else {
			j.mensagem = ""
			j.nroJogos = 0
		}
	}
}

func (j *jogo) verificaNumerosValidos() {
	if j.valido == true {
		for _, i := range j.dezenas {
			if strings.Contains(numerosValidos, i) == false {
				j.mensagem = fmt.Sprintf("Jogo invalido, número %s fora da faixa valida 1..25 !", i)
				j.nroJogos = len(j.dezenas)
				j.valido = false
				return
			}
		}
		j.mensagem = ""
		j.nroJogos = len(j.dezenas)
	}
}

func (j *jogo) verificaAcertos() {
	if j.valido == true {
		acertos := 0
		for _, i := range j.dezenas {
			if strings.Contains(dezenasSorteadas, i) == true {
				acertos++
			}
		}
		if acertos == 15 {
			j.nroAcertos = acertos
			j.mensagem = fmt.Sprintf("!!! VOCÊ GANHOU !!!")
		} else {
			j.nroAcertos = acertos
			j.mensagem = fmt.Sprintf("%d números Jogados e %d acertos !", j.nroJogos, acertos)
		}
	}
}

func inicia() map[int]jogo {
	nroJogo := 0
	jogos := make(map[int]jogo)
	var j jogo

	file, err := os.Open("jogos.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Lendo arquivo texto com os jogos e alimentando variavel do tipo MAP
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		j.setDezena(strings.Fields(scanner.Text()))
		j.validaMaxPorJogo()
		j.validaMinPorJogo()
		j.verificaNumerosValidos()
		j.verificaAcertos()
		nroJogo++
		jogos[nroJogo] = j
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return jogos
}

func main() {

	result := inicia()
	//Imprimindo resultado dos jogos
	fmt.Printf("\nDezenas sorteadas: %s\n", dezenasSorteadas)
	for i, v := range result {
		fmt.Printf("\nJogo: %d --> %s\n   Dezenas: %s\n", i, v.mensagem, v.dezenas)
	}

}
