# getting_started_TEST

## 2. Static code analysis of the triangle program

### Install Metrics software in your IDE
Jeg har installeret [gocyclo](https://github.com/fzipp/gocyclo) til at måle cyclomatic complexities i Go. Go og VS Code understøtter ikke på samme måde code metrics som Java/NetBeans.

### Check coding standards in your Triangle program

Go har generelt en række [kodestandarder](https://golang.org/doc/effective_go.html), som skal overholdes, og Go kan ikke compiles hvis disse standarder ikke overholdes. Dvs at programmet simpelthen slet ikke kan køres, hvis man har oprettet en variabel og ikke bruger den. Det burde ikke have betydning for hvorvidt programmet kan køre, for det har ingen indflydelse på noget som helst, men Go opretholder alligevel disse kodestandarder utrolig striks, og man er derfor tvunget til at overholde dem. 

![](https://github.com/xNoga/getting_started_TEST/blob/master/img/Screen%20Shot%202018-02-17%20at%2012.14.18.png)

Linket oven over indeholder et hav af standarder som også bruges i Go. F.eks. er der rigtig mange kald i Go, der kræver at man opretter en variabel med resultatet og en variable med error, hvis en error opstår. Man er derfor nødt til at tjekke om man fik en error, fordi man som nævnt ikke må oprette variabler, der ikke bliver brugt. Det er generelle kodestandarder som Go kræver man overholder. 

![](https://github.com/xNoga/getting_started_TEST/blob/master/img/Screen%20Shot%202018-02-17%20at%2012.25.09.png)

Der findes et hav af kodestandarder, der skal overholdes, så jeg nævner ikke dem alle her. Men de kan ses i linket øverst i afsnittet.

### Calculate central metrics in your triangle program

Som nævnt kan jeg ikke helt gøre dette i Go. Der findes dog en hjemmeside kaldet [goreportcard.com](https://goreportcard.com) som kan måle ens kode på nogle andre parametre og give en general status for ens program. For min kode kan resultatet ses [her.](https://goreportcard.com/report/github.com/xNoga/getting_started_TEST)

### Find out what CC variation your metrics tool uses

Jeg har brugt [gocyclo](https://github.com/fzipp/gocyclo) til at måle CC med i min Go-applikation. Gocyclo er et lille program der måler CC ud fra følgende parametre:
```
1 is the base complexity of a function
+1 for each 'if', 'for', 'case', '&&' or '||'
```

Når jeg kører gocyclo får jeg følgende resultat: 

![](https://github.com/xNoga/getting_started_TEST/blob/master/img/Screen%20Shot%202018-02-17%20at%2012.50.46.png)

Hvis vi tager exempel i den første linje:
6 indikerer mængden af compleksitet. "main" refererer til hvilken package der er tale om. "findTriangleName" er navnet på den givne funktion, og "main.go:37:1" indikerer hvor henne i filen der er tale om.
Som det kan ses på resultatet er vores program ikke særlig komplekts. 

### Possibly refactor your code based on static testing results.
Jeg har ikke refactored noget som følge af CC results, men jeg har refaktoreret koden så den understøtter "Unit" (Go håndterer tests lidt anderldes, men konceptet er det samme) tests.

### Write test cases in xUnit tool.
Se koden. Brug kommandoen "go test" for at køre tests (fra en terminal der befinder sig på roden af projektet selvfølgelig).
Jeg havde desværre en del problemer med at sætte Unit-tests op i Go, da det gav mig problemer med inputparametre af en eller anden årsag. Der skulle derfor gerne have været et par tests mere, men grundet tidspres må jeg nøjes med de tre tests der er. Den første test tester om  input gør som forventet, og test nr. 2 tester om en trekant er valid og om trekantens navn derefter er som vi forventer. 

## Peer Review Checklist

Checklisten kommer med generelle råd til hvordan opnår en mere effektiv Code Review. Code Review er generelt ikke noget vi har gjort os så meget i på uddannelse, men kilden giver rigtig fine perspektiver på, hvorfor det er vigtigt at bruge en række praksisser mht Code Review. Jeg synes specielt punkterne, der omhandler at man ikke skal kigge på code review i mere end 60-90 minutter. I forlængelse heraf skriver de også man ikke skal kigge på mere end 200 linjer kode i "one sitting" og at man gerne skal tage 20 minutters pause mellem hver review settings. Jeg synes det er gode råd, da man nemt kan komme til at læse forkert/ikke koncentere sig ordentligt, når man har kigget på andres kode i længere tid. Jeg har i hvert fald selv flere gange mistet koncentrationen, når jeg har skulle sætte mig ind i andres kode, og gerne prøve at review'e det. 
De skriver også det er vigtigt at defects er fixed og ikke bare fundet. Man kan som programmør nogle gange godt støde på situationer, hvor man ved at et givent flow i applikationen giver en fejl, men man gør ikke noget ved det lige med det samme, fordi man måske føler der er andre vigtigere ting at tage sig til. Jeg synes det er en god pointe, at man husker at rette disse defects med det samme. 

## Review code that mysteriously fails its unit tests

## Coding Standard Document
Jeg har personligt ikke en række standarder, som jeg mener er de rigtige/forkerte. Jeg foretækker i stedet at man som team bare blive enige om, HVILKE standarder man bruger, således at hele team'et bruger de samme standarder, best practices og code conventions. Det er i hvert fald vigtigt at blive enige om naming conventions - om man f.eks. bruger camelcase eller lowercase når man navngiver. Der er selvfølgelig også bitte små ting som hvordan man bruger tabs og spaces.
Det kan også være en fordel efter min mening, at man f.eks. også bliver enige om, hvordan man bliver enige om opbygningen af routes i et REST-API, så alle ved, uden at skulle spørge, hvad et givent kald gør. 
Det er generelt bare vigtigt at man som team bliver enige om, hvordan koden skal se ud, og hvordan evt. et almindelig funktionskald skal se ud.
Det er yderligere også vigtigt at gøre brug af comments i sin kode, så alle kan læse sig frem til, hvad et givent stykke kode gør. Så længe man gør brug af comments kan man også i højere grad afvige fra en aftalt form, da resten af team'et forhåbentlig stadig kan læse sig frem til, hvad koden gør. 

## Highlights from lecture by Gitte Ottosen, Gapgemini-Sogeti

# Running the program
First of all you should have Go installed on your computer. If you have not installed Go a guide can be found [here.](https://golang.org/doc/install)
To run the program simply pull this project to a folder. From there you just type a following command in your terminal at the root of the project: 

```
go run main.go 1 2 3
> Scalene triangle
```
```
go run main.go 1 1 2
> Isosceles triangle
```

```
go run main.go 2 2 2 
> Equilateral triangle
```

It is important to use a space between each integer. Only type numbers. The program will tell you if you have used a wrong input.

## Test cases

I have left out *Created By, Date of creation, Executed By and Date of execution* as I have created and executed all test on the same day; February 10, 2018.

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC001 | The objective of this test is to verify that a triangle with **three** sides of equal length will return 'equilateral triangle' | The user should obtain three equal integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'equilateral triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid and the program returned 'equilateral triangle' <br><br> 2. If the input was invalid the program will return 'invalid input' | Success |

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC002 | The objective of this test is to verify that a triangle with **three** sides of equal length will return 'equilateral triangle' | The user should obtain three equal integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'equilateral triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid but the program **did not** return 'equilateral triangle' <br><br> 2. If the input was invalid the program will return 'invalid input' | Fail |

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC003 | The objective of this test is to verify that a triangle with **no** sides of equal length will return 'scalene triangle' | The user should obtain three different integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'scalene triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid and the program returned 'scalene triangle' <br><br> 2. If the input was invalid the program will return 'invalid input' | Success |

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC004 | The objective of this test is to verify that a triangle with **no** sides of equal length will return 'scalene triangle' | The user should obtain three different integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'scalene triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was not valid and the program returned 'invalid input' | Fail |

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC005 | The objective of this test is to verify that a triangle with **two** sides of equal length will return 'isosceles triangle' | The user should obtain three integers and two of the integers should be equal | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'isosceles triangle' if two integers are equal <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid and the program returned 'isosceles triangle' <br><br> 2. If the input was invalid the program will return 'invalid input' | Success |
