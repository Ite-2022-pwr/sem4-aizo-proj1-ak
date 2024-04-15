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

## Sposób użycia

```
Sposób użycia:
$ ./aizo1 <cmd> [options]

        verify                          weryfikacja poprawności zaimplementowanych algorytmów sortowania
        generate <typ danych> <rozmiar> generowanie plików z losowymi danymi wejściowymi
        run [options]                   uruchamianie wybranego algorytmu sortowania
        interactive                     tryb interaktywny (menu)
        help                            wyświetlanie tej wiadomości

Przykłady:
$ ./aizo1 generate float32 100
$ ./aizo1 run -alg=quick -type=string -generate=100000 -pivot=losowy
$ ./aizo1 run -alg=quick -type=string -pivot=prawy -input-file=data/input/input_int_1000000.txt
```

### Opcja `run`

```
Sposób użycia: ./aizo1 run [options]

  -alg string
        algorytm sortowania (insertion, heap, shell, quick). [WYMAGANY]
  -gap string
        jakiej metody obliczania odstępu w sortowaniu Shella użyć: shell, franklazarus, hibbard. (default "shell")
  -generate int
        ile liczb do wygenerowania. (default 1000)
  -help
        wyświetlanie tej wiadomości
  -input-file string
        czy pobrać dane wejściowe z pliku i jakiego (zamiast generate).
  -pivot string
        jakiego pivota do sortowania szybkiego użyć: lewy, prawy, środkowy, losowy. (default "lewy")
  -save
        czy zapisać wygenerowane dane do pliku.
  -type string
        typ danych (int, int32, int64, float32, float64, byte, string). [WYMAGANY]
```