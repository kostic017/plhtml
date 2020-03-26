Konačno ljudi koji "programiraju" u HTML-u, te sebe nazivaju "programerima" mogu to i da budu.

Sav kod napisan u programskom jeziku HTML (ili proceduralnom jeziku HTML, skraćeno PL/HTML po uzoru na PL/SQL)
predstavlja sintaksički i semantički validan HTML kod, što je bio glavni cilj prilikom dizajniranja jezika. Sintaksička
ispravnost programa je proveravana korišećnjem [W3C validatora](https://validator.w3.org/#validate-by-input). Programi
su semantički validni u smislu da su nazivi naredbi birani među nazivima HTML elemenata tako da naziv ili opis elementa
ukazuje na značenje naredbe (`var`, `input`, `output`, isl).

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

Za implementaciju je korišćen Go programski jezik bez third-party alata, odnosno skener i parser generatora kao što su
flex i bison. Skener nije implementiran uz pomoć regularnih izraza, već je korišćen algoritam opisan na stranici
[Implementing Lexers and Parsers](http://www.cse.chalmers.se/edu/year/2015/course/DAT150/lectures/proglang-04.html).
Na sličan način je implementiran i [skener za Go](https://github.com/golang/go/blob/master/src/go/scanner/scanner.go).
Osnovna ideja je da skener čita karakter po karakter i da se nakon svakog formiranog tokena na osnovu prvog sledećeg
karaktera `c` određuje koja vrsta tokena je sledeća:

  - ako je `c` cifra, čita se dalje sve dok su cifre, pa se formira ceo broj
  - ako je `c` slovo, čita se dalje sve dok su slova ili cifre, pa se formira identifikator
  - ako je `c` navodnik, čita se dalje sve dok se ne naiđe do drugog navodnika, pa se formira string
  - ako je `c` razmak, ignoriše se i čita se sledeći karakter
  
U rečima "čita se dalje sve dok" uočavamo pravilo maximal munch.

Parser je implementiran metodom rekurzivnog spusta, gramatika je data u fajlu [grammar.txt](grammar.txt).

Semantička analiza je implementirana nekim delom, proverava se da li su sve promenljive deklarisane pre korišćenja i da
li se ne deklariše više promeljivih sa istim nazivom u istom opsegu, uz korišćenje [špageti steka](https://www.geeksforgeeks.org/g-fact-87/).

Interpretiranje je takođe implementirano nekim delom, ali u dovoljnoj meri da se programi u */tests/examples* folderu
izvršavaju uspešno.

Nakon što izbildujete izvršni fajl sa `go build`, interpretacija se pokreće kroz komandnu liniju

```bat
plhtml <putanja_do_fajla.html>
```
