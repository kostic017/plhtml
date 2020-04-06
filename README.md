Konačno osobe koje "programiraju" u HTML-u i sebe nazivaju "programerima" mogu i da budu programeri.

Sav kod napisan u programskom jeziku HTML (skraćeno PL/HTML) predstavlja sintaksički i semantički validan HTML kod,
što je bio glavni cilj prilikom dizajniranja jezika. Sintaksička ispravnost programa je proveravana korišećnjem
[W3C validatora](https://validator.w3.org/#validate-by-input). Programi su semantički validni u smislu da su nazivi
naredbi birani među nazivima HTML elemenata tako da naziv ili opis elementa ukazuje na značenje naredbe
(`var`, `input`, `output`, isl).

Sledeći program ispisuje prvih `n` Fibonačijevih brojeva.

```html
<!doctype html>
<html lang="en">
    <head>
        <!-- Program to Display Fibonacci Series -->
        <title>`Fibonacci numbers`</title>
    </head>
    <body>

        <main>

            <var class="integer">a</var>
            <var class="integer">b</var>
            <var class="integer">c</var>
            <var class="integer">i</var>
            <var class="integer">n</var>

            <data value="0">a</data>
            <data value="1">b</data>
            <data value="0">c</data>
            <data value="1">i</data>

            <output>`n: `</output>
            <input name="n">

            <div data-while="i &leq; n">
                <data value="a + b">c</data>
                <output>c + ` `</output>
                <data value="b">a</data>
                <data value="c">b</data>
                <data value="i + 1">i</data>
            </div>

        </main>

    </body>
</html>
```

# Specifikacija

  * Ključne reči: `doctype`, `lang`, `html`, `head`, `title`,`body`, `main`, `var`, `class`, `output`, `input`, `name`, `data`, `value`, `div`, `if`, `while`
  * Ugrađeni tipovi: `integer`, `real`, `boolean`, `string`
  * Aritmetički operatori: `+`, `-`, `*`, `/`, `%`, `(`, `)`
  * Logički operatori: `&and;`, `&or;`, `!`
  * Operatori poređenja: `&lt;`, `&gt;`, `&leq;`, `&geq;`, `&equals;`, `&ne;`
  * Specijalni karakteri: `\\`, `\t`, `\n`

# Implementacija
Za implementaciju je korišćen Go programski jezik bez third-party alata (skener i parser generatora).

Skener nije implementiran uz pomoć regularnih izraza/konačnih automata, već je korišćen algoritam opisan na stranici
[Implementing Lexers and Parsers](http://www.cse.chalmers.se/edu/year/2015/course/DAT150/lectures/proglang-04.html).
Osnovna ideja je da skener čita kod karakter po karakter i da se nakon svakog formiranog tokena na osnovu prvog sledećeg
karaktera `c` određuje koja vrsta tokena je sledeća:

  - ako je `c` cifra, čita se dalje sve dok su cifre, pa se formira ceo broj
  - ako je `c` slovo, čita se dalje sve dok su slova ili cifre, pa se formira identifikator
  - ako je `c` navodnik, čita se dalje sve dok se ne naiđe na prvi sledeći navodnik, pa se formira string
  - ako je `c` razmak, ignoriše se i čita se sledeći karakter

U rečima "čita se dalje sve dok" uočavamo pravilo maximal munch.

Parser je implementiran metodom rekurzivnog spusta. Pregled gramatike je dat u fajlu [grammar.txt](grammar.txt).

U toku semantičke analize se proverava:
  - korektnost tipova
  - da li su promenljive deklarisane pre korišćenja
  - da li se ne deklariše više promeljivih sa istim nazivom u istom opsegu

# Ograničenja
Interpretiranje je implementirano nekim delom, ali u dovoljnoj meri da se programi u */tests* folderu izvršavaju
uspešno. Tokom interpretacije ignorišu se opsezi i tipovi. Samim tim feature-i kao što su inicijalizacija nulom i
učitavanje iz konzole funkcionišu samo za cele brojeve.

# Korišćenje
  1. Preuzmite [Go kompajler](https://golang.org/dl/).
  2. Izbildujte izvršni fajl sa `go build`.
  3. Pokrenite interpretaciju kroz komandnu liniju.

```bat
plhtml <path_to_source> [<path_to_input>]
```

# Reference
Kombinovao sam materijale sa različitih izvora, poput:
  - [CS143 Compilers](https://web.stanford.edu/class/archive/cs/cs143/cs143.1128/)
  - [Implementing Lexers and Parsers](http://www.cse.chalmers.se/edu/year/2015/course/DAT150/lectures/proglang-04.html)
  - [A Simple Recursive Descent Parser](http://math.hws.edu/javanotes/c9/s5.html)
  - [Let’s Build A Simple Interpreter](https://ruslanspivak.com/lsbasi-part1/)

Gledao sam i kako su implementirani kompajleri za neke poznate jezike:
  - [Go](https://github.com/golang/go/blob/master/src/go)
  - [TypeScript](https://github.com/microsoft/TypeScript/tree/master/src/compiler)
