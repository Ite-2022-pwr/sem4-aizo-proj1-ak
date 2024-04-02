# AZO - zadanie projektowe nr 1
# Badanie efektywności wybranych algorytmów sortowania ze względu na złożoność obliczeniową

Autor: [Artur Kręgiel](https://github.com/arkregiel)

Prowadzący: [dr inż. Zbigniew Buchalski](https://wit.pwr.edu.pl/wydzial/struktura-organizacyjna/pracownicy/zbigniew-buchalski)

## Opis projektu

Opis projektu znajduje się na [tej stronie](http://dariusz.banasiak.staff.iiar.pwr.wroc.pl/azo/AZO_lista1.pdf).

Wybrano wariant trzeci:

```
5.0 – sortowanie przez wstawianie, przez kopcowanie, Shella i szybkie - program w wersji obiektowej,
dodatkowo badamy wpływ typu danych np. int i float, wpływ wyboru odstępów dla algorytmu
Shella (2 różne wzory tworzące algorytmy o różnych złożonościach) oraz sposób wyboru pivota
(skrajny lewy, skrajny prawy, środkowy oraz losowy) dla algorytmu szybkiego.
```

Projekt, za zgodą prowadzącego, został zaimplementowany w języku programowania [Go](https://go.dev/).

## Kompilacja i uruchamianie

```
$ go build -o aizo1
$ ./aizo1
```

Program można również uruchomić bez wyraźnego budowania projektu:

```
$ go run .
```

