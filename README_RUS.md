# Day 01 - Go Boot camp

## Содержание

1. [Глава I](#chapter-i) \
    1.1. [Основные правила](#general-rules)
2. [Глава II](#chapter-ii) \
    2.1. [Правила дня](#rules-of-the-day)
3. [Глава III](#chapter-iii) \
    3.1. [Введение](#intro)
4. [Глава IV](#chapter-iv) \
    4.1. [Упражнение 00: Чтение](#exercise-00-reading)
5. [Глава V](#chapter-v) \
    5.1. [Упражнение 01: Оценка ущерба](#exercise-01-assessing-damage)
6. [Глава VI](#chapter-vi) \
    6.1. [Упражнение 02: Афтепати](#exercise-02-afterparty)


<h2 id="chapter-i" >Глава I</h2>
<h2 id="general-rules" >Основные правила</h2>

<h2 id="chapter-i" >Глава I</h2>
<h2 id="general-rules" >Основные правила</h2>

* Твоя программа не должна закрываться неожиданно (выдавая ошибку при корректном вводе). Если это произойдет, твой проект будет считаться неработаспособным и получит 0 во время оценки.
* Мы рекомендуем тебе писать тесты для твоего проекта, даже если если они и не оцениваются. Это даст тебе возможность легко тестировать твою работу и работу твоих пиров. Ты убедишься что тесты очень полезны, во время защиты. Во время защиты ты свободен использовать свои тесты и/или тесты пира которого ты проверяешь.
* Отправляй свою работу в нужный git репозиторий. Работа будет оцениваться только из git репозитория.
* Если твой код использует сторонние зависимости, следует использовать [Go Modules](https://go.dev/blog/using-go-modules) для управления ими.

<h2 id="chapter-ii" >Глава II</h2>
<h2 id="rules-of-the-day" >Правила дня</h2>

* Пиши код только в `*.go` файлах и (в случае стронних зависимостей) `go.mod` + `go.sum`
* Твой код для этого задания должен собираться с использовния простого `go build`

<h2 id="chapter-iii" >Глава III</h2>
<h2 id="intro" >Введение</h2>

В мире программирования много популярных форматов данных и в го в частности. Но это очень похоже что ты встретишься с одним из них на своем пути - XML или JSON. Множество и множество API которые используют JSON и/или XML для представления структурированной информации.

И...иногда некоторые пекари используют их для хранения рецептов. Так старая известная пекарня в Вилларибе использовала только старый добрый XML для хранения списка рецептов тортов. Ес мы возьмем и посмотрим на кусочек этой базы данных это будет выглядеть типа такого:

```xml
<recipes>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>40 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>3</itemcount>
                <itemunit>cups</itemunit>
            </item>
            <item>
                <itemname>Vanilla extract</itemname>
                <itemcount>1.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>7</itemcount>
                <itemunit></itemunit> <!-- itemunit может быть пустым  -->
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
                <itemunit>pieces</itemunit>
            </item>
            <!-- Здесь может быть больше ингридиентов  -->
        </ingredients>
    </cake>
    <cake>
        <name>Blueberry Muffin Cake</name>
        <stovetime>30 min</stovetime>
        <ingredients>
            <item>
                <itemname>Baking powder</itemname>
                <itemcount>3</itemcount>
                <itemunit>teaspoons</itemunit>
            </item>
            <item>
                <itemname>Brown sugar</itemname>
                <itemcount>0.5</itemcount>
                <itemunit>cup</itemunit>
            </item>
            <item>
                <itemname>Blueberries</itemname>
                <itemcount>1</itemcount>
                <itemunit>cup</itemunit>
            </item>
            <!-- Здесь может быть больше ингридиентов  -->
        </ingredients>
    </cake>
    <!-- Здесь может быть больше тортов  -->
</recipes>
```

Жизнь была прекрасна и проста пока владелец пекарни не заметил что в соседней деревне Виллабажо сейчас живет грязный предатель который решил украсть его рецепты. Для того что бы это провернуть он использовал другой тип хранения данных и это утаило его от правосудия! 

```json
{
  "cake": [
    {
      "name": "Red Velvet Strawberry Cake",
      "time": "45 min",
      "ingredients": [
        {
          "ingredient_name": "Flour",
          "ingredient_count": "2",
          "ingredient_unit": "mugs"
        },
        {
          "ingredient_name": "Strawberries",
          "ingredient_count": "8"  // ingredient_unit больше тут нет!
        },
        {
          "ingredient_name": "Coffee Beans",
          "ingredient_count": "2.5",
          "ingredient_unit": "tablespoons"
        },
        {
          "ingredient_name": "Cinnamon",
          "ingredient_count": "1"
        }
      ]
    },
    {
      "name": "Moonshine Muffin",
      "time": "30 min",
      "ingredients": [
        {
          "ingredient_name": "Brown sugar",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        },
        {
          "ingredient_name": "Blueberries",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        }
      ]
    }
  ]
}
```

Он не мог не заметить что вор не только своровал его рецепты, но и также он изменил некоторые из них. Некоторые ингридиенты отсутствуют, количества изменены, части переименованы. Так что он подготовил свою месть!


<h2 id="chapter-iv" >Глава IV</h2>
<h3 id="ex00">Упражнение 00: Чтение</h3>

Перво-наперво, он стал изучать как читать из базы данных. Владелец уже имел CLI, так он решил что чтение файла будет простым, так обе из них будут работать (файлы могут различаться по расширению для простоты) 

`~$ ./readDB -f original_database.xml`
`~$ ./readDB -f stolen_database.json`


Не только это, он также решил что чтение обеих файлов не будет сложным производить через один интерфейс, который будет назван `DBReader`. Это означает что чтение различных форматов реализовано через разные *имплементации* одного интерфейса `DBReader`, который будет выдавать одни и те же типы объектов в результате, без разницы прочтен файл был из оригинальной базы данных или украденной. Верно, его идея это выбирать подходящую имплементацию на основе расширения файла.

И тебе нужно будет помочь ему с этим. Подумай какие типы объектов в этой базе данных и как их можно представить в коде. Затем напиши интерфейс `DBReader` и две реализации его, одну для чтения JSON, другую для чтения XML. Обе из них должны возвращать объект одного типа в качестве результата. 

To check that his idea works, make the code print JSON version of the database when it's reading from XML and vice versa. Both XML and JSON fields should be indented with 4 spaces ("pretty-printing").

Для проверки что это идея работает сделай код для вывода на экран JSON версии базы данных XML и наоборот (лат. `vice versa`). Оба формата XML и JSON должны иметь отступ полей в 4 пробела (так называемый приятный вывод "pretty-printing").

<h2 id="chapter-v" >Глава V</h2>
<h3 id="ex01">Exercise 01: Оценка ущерба</h3>

Okay, so now the owner decided to compare the databases. You've seen that the stolen database has modified versions of the same recipes, meaning there are several possible cases:

Оки, так теперь владелец решил сравнить базы данных. Ты увидел что сворованная база данных имеет измененные версии одних и тех же рецептов, имеется в виду что они похожие в таких случах:

1) New cake is added or old one removed
2) Cooking time is different for the same cake
3) New ingredient is added or removed for the same cake. *Important:* the order of ingredients doesn't matter. Only the names are.
4) The count of units for the same ingredient has changed.
5) The unit itself for measuring the ingredient has changed.
6) Ingredient unit is missing or added

1) Новый торт добавлен или старый один удален.
2) Время приготовления различное для одного и того же торта
3) Новые ингридиенты добавлены или удалены из одного же торта.
4) Количество частей для одного и того же ингридиента было поменяно.
5) Изменилась сама единица измерения ингредиента
6) Единица ингридиента была потеряна или добавлена

Быстрым просмотром базы данных, владелец так же уведомил что никто не удосужился переименовать торты или ингридиенты (возможно, только добавлены некоторые новые), так что ты можешь предположить что имена в обеих базах данных одинаковые.

Твое приложение должно работать как-то так:

`~$ ./compareDB --old original_database.xml --new stolen_database.json`

Оно должно работать с обеими форматами (JSON и XML) для орининальной И новой базой данных, переиспользуя код из Упражнения 00.

Вывод должен выглядить таким образом (некоторые случаи объяснены ниже):

```
ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee beans" for cake  "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake  "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake  "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"
```

<h2 id="chapter-vi" >Глава VI</h2>
<h3 id="ex02">Exercise 02: Афтепати</h3>

Копая вглубь базы данных, владелец пекарни Виллариба внезапно осознал - этот парень просто сокровище. Некоторые рецепты были улучшены значительно в сравнении со старой версией и эти новые идеи действительно творческие. Он пробился в Виллабажои нашел парня, который как сперва казалось украл его самое драгоценное наследие.

...The same evening in the tavern two old bakers were hugging, drinking and laughing so hard that it was heard in both villages. They became best friends during the last couple of hours, and each of them was really happy to finally find the person who could blabber all night about cakes! Also turns out, the guy did't steal the database, he was just trying to guess by the taste and tried to improvise around a bit.

After this whole mess they both agreed to use your code, but asked you to do one last task for them. They were quite impressed by how you've managed to do the comparison between the databases, so they've also asked you to do the same thing with their server filesystem backups, so neither bakery would run into any technical issues in the future.

So, your program should take two filesystem dumps.

`~$ ./compareFS --old snapshot1.txt --new snapshot2.txt`

They are both plain text files, unsorted, and each of them includes a filepath on every like, like this:

```
/etc/stove/config.xml
/Users/baker/recipes/database.xml
/Users/baker/recipes/database_version3.yaml
/var/log/orders.log
/Users/baker/pokemon.avi
```

Your tool should output the very similar thing to a previous code (without CHANGED case though):

```
ADDED /etc/systemd/system/very_important/stash_location.jpg
REMOVED /var/log/browser_history.txt
```

There is one issue though - the files can be really big, so you can assume both of them won't fit into RAM on the same time. There are two possible ways to overcome this - either to compress the file in memory somehow, or just read one of them and then avoid reading the other. **Side note:** this is actually a very popular interview question.


