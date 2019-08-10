**Громадная просьба\! Оставьте свои замечания в
[комментариях](User_comments:adriano32/Инженерное_ПО "wikilink"),
я сам с удовольствием внесу правки после ваших предложений, а потом уже
готовую статью можно будет перенести собственно L.O.R. Wiki**

При написании статьи [автор](User:adriano32 "wikilink")
<s>копипастил</s> вдохновлялся собственным опытом
использования некоторого из этого софта, частично
[LOR-FAQ-Scientific](LOR-FAQ-Scientific "wikilink") (в разделе
визуализации данных - на большее пока не сподвигся),
букмарками после просмотров [1](http://www.opencollector.org/)
и не только, новостями и обсуждениями на L.O.R.

По поводу структуры: когда статья будет наполнена максимально возможным
количеством софта, структуру заголовков приведу в соответствие к
статьям [Linux в школе](Linux_в_школе "wikilink") и [Программы
для дома и работы](Программы_для_дома_и_работы "wikilink") — то есть
подзаголовок-\>список.

## Инструменты для математических и научных расчётов

### Приложения

#### Scilab

[Scilab](http://www.scilab.org/) — свободный открытый кроссплатформенный
пакет для выполнения расчётов для образовательных, научных и инженерных
задач с собственным MATLAB-оподобным языком.

Содержит большое количество встроенных функций для вычислений,
визуализации, оптимизации, обработки сигналов, расчёта систем
управления итп. Также существуют ряд модулей для расширения
стандартного функционала, они поставляются отдельно, но
[доступны для загрузки](http://atoms.scilab.org/) на официальном
сайте.

Также в поставку Scilab'a входит графическая среда моделирования
[Xcos](http://www.scilab.org/products/xcos) на базе
[Scicos](http://www-rocq.inria.fr/scicos/), который тоже продолжает
развиваться, подробная документация с примерами, конвертор
MATLAB'овских M-файлов.

Кроме официальной документации по Scilab есть достаточно много
руководств и на
[английском](http://www.google.com/search?q=scilab+filetype%3Apdf&lr=lang_en),
и на [русском
языке](http://www.google.com/search?q=scilab+filetype%3Apdf&lr=lang_ru).
Вот лишь одно из них [из библиотеки ALT
Linux](http://docs.altlinux.org/books/2008/altlibrary-scilab-20090409.pdf)

#### Octave

[GNU Octave](http://www.gnu.org/software/octave/) — высокоуровневый
интерпретируемый язык для расчётов в учёбе и науке. Сходен с языком
MATLAB'a. Встроенные библиотеки и компоненты позволяют проводить
операции с матрицами и полиномами, численное дифференцирование и
интегрирование, оптимизацию, статистические расчёты, обработку сигналов,
строить 2D и 3D графики и другое.

Те, кому мало встроенных пакетов, могут доустановить недостающие из
[Octave-Forge](http://octave.sourceforge.net/).

Как и Scilab, Octave имеет превосходную
[документацию](http://www.gnu.org/software/octave/doc/interpreter/)
с примерами.

#### FreeMat

[FreeMat](http://freemat.sourceforge.net/) — ещё один пакет для научных
вычислений, который гоняется за славой MATLAB'a. В версии 4.0 заявлена
[95%](http://freemat.sourceforge.net/faq.html#4) совместимость по
количеству функций.

[Документация](http://freemat.sourceforge.net/help/index.html) доступна
на официальном сайте.

#### Sage

[Sage](http://www.sagemath.org/) — предоставляет основанный на
[Python](http://www.python.org/) интерфейс к [почти сотне популярных
программных пакетов](http://www.sagemath.org/links-components.html)
для математических расчётов, символьных вычислений, статистического
анализа, программирования, построения графиков и многого другого.

Работать с Sage можно в командной строке с использованием
[IPython](http://ipython.scipy.org/) или через интерфейс "Notebook" в
браузере благодаря [SageNB (The Sage Notebook
server)](http://nb.sagemath.org/). Благодаря последнему, Sage совсем не
обязательно устанавливать на свой компьютер или, например, на все
компьютеры в аудитории: можно установить Sage на один
высокопроизводительный компьютер в локальной сети или же
можно воспользоваться одним из тестовых Sage-серверов. Список
официальных публичных Sage-серверов доступен на [главной
странице проекта](http://www.sagemath.org/), один из публичных
Sage-серверов также работает в [Институте кибернетики им. Глушкова НАН
Украины при поддержке НаУКМА (Национального университета
"Киево-Могилянская академия")](http://sage.openopt.org/).

#### Cantor (KDE Edu)

[Cantor](http://edu.kde.org/cantor/) (KDE) — часть проекта KDE Edu,
фронтенд к Sage, Maxima, R или KAlgebra (на выбор). Предоставляет
интерфейс в виде рабочих листов.

#### SMath Studio

[SMath Studio](http://en.smath.info/) — символьная среда вычислений,
разрабатываемая петербуржцем Андреем Ивашовым и программистом из
Бреста Романом Стрильцом на [Mono](http://www.mono-project.com/).
Smath во многом похож на Mathcad и даже умеет полностью импортировать
.xmcd файлы, если те содержат уже имплементированные функции.

#### Jasymca

[Jasymca: Programmable Java
calculator](http://webuser.hs-furtwangen.de/~dersch/#Jasymca) —
Java-апплет с функционалом неплохой системы компьютерной алгебры
(CAS): арифметика, операции с матрицами, полиномами, символические
вычисления, интегрирование, дифференцирование, визуализация,
програмирование. Самый главный плюс, что версия 3.0 поставляется в
том числе под Android, а иметь такой программируемый калькулятор на
смарте очень даже неплохо.

#### Maxima

[Maxima](http://maxima.sourceforge.net/) — свободная кроссплатформенная
система компьютерной алгебры (CAS), написанная на [Common
Lisp](http://www.clisp.org/).

Имеет несколько графических интерфейсов, среди которых самыми ходовыми
являются Xmaxima и [wxMaxima](http://andrejv.github.com/wxmaxima/).
При этом с Maxima можно работать и в командной сроке, и при помощи
дополнительных расширений в буфере GNU Emacs или в браузере.

#### Mathomatic

[Mathomatic](http://www.mathomatic.org/) — кроссплатформенная система
компьютерной алгебры (CAS) с командным интерфейсом, написанная на C.
Среди возможностей: решение уравнений, дифференцирование, операции с
полиномами, операции с действительными и комлексными числами и
другое.

[Документация](http://www.mathomatic.org/math/doc/index.html) в
комплекте.

#### Singular

[Singular](http://www.singular.uni-kl.de/) — система компьютерной
алгебры (CAS), специализирующаяся на операциях над полиномами, а
именно их применении в коммутативной и некоммутативной алгебре,
алгебраической геометрии, теории особенностей (singularity
theory).

Функционал может быть расширен за счёт [дополнительных
библиотек](http://www.singular.uni-kl.de/Manual/latest/sing_775.htm#SEC851)
и линковки с [другими
приложениями](http://www.singular.uni-kl.de/index.php/third-party-software/13692.html).

[Документация](http://www.singular.uni-kl.de/index.php/singular-manual.html)
на официальном сайте.

#### Axiom

[Axiom](http://axiom-developer.org/) — свободная кроссплатформенная
система компьютерной алгебры (CAS). Код Axiom разрабатывается в
лучших традициях Literate Programming Дональда Эрвина Кнута.

#### OpenAxiom

[OpenAxiom: Scientific Computation System](http://www.open-axiom.org/) —
живой форк Axiom.

#### FriCAS

[FriCAS](http://fricas.sourceforge.net/) — ещё один живой форк Axiom.

#### ARIBAS

[ARIBAS](http://www.mathematik.uni-muenchen.de/~forster/sw/aribas.html)
— малость устаревший, но рабочий кроссплатформенный интерактивный
интерпретатор для арифметических операций с большими целыми числами
и числами с плавающей точкой с произвольной точностью с синтаксисом,
напоминающим Pascal/Modula, написанный на C. Поскольку часть кода
написана на ассемблере ARIBAS работает только на 32-битных системах.

#### REDUCE

[REDUCE](http://reduce-algebra.com/) — написанная на диалекте Lisp'a
Standard Lisp система для алгебраических расчётов. Умеет работать с
полиномами, матрицами, выполнять аналитическое дифференцирование и
интегрирование, факторизацию и полиномизацию и другое.

#### R

[R](http://www.r-project.org/) — это язык программирования и среда для
статистических расчётов и визуализации их результатов.

#### PSPP

[GNU PSPP](http://www.gnu.org/software/pspp/) — программа для
статистического анализа больших наборов данных. Имеет как
командный, так и графический интерфейс, а также встроенные функции
для обработки данных и визуализации.

#### gretl

[gretl](http://gretl.sourceforge.net/) —кроссплатформенный пакет для
эконометрического анализа, написанный на C. Поддерживает несколько
форматов входных файлов (XML, CSV, Excel, Gnumeric).

#### JMulTi

[JMulTi](http://www.jmulti.de/) — приложение на Java для анализа
временных рядов. Версия для Linux заброшена разработчиками в
силу ряда причин, однако всё ещё доступна для загрузки.

#### Stata

[Stata](http://www.stata.com/) — коммерческое ПО для статистического
анализа, доступна версия для Linux.

#### IBM SPSS Statistics

[IBM SPSS
Statistics](http://www.ibm.com/software/analytics/spss/products/statistics/)
— мощный коммерческий пакет от IBM для статистических расчётов. Есть
триальные версии некоторых компонентов, работающие в Linux (32-bit
only), например SPSS Statistics Desktop, доступные после регистрации.

### Библиотеки

#### GMP

[GMP](http://gmplib.org/) — свободная открытая библиотека для расчётов с
произвольной точностью над целыми числами, рациональными дробями и
числами с плавающей запятой. Написана с использованием языка
ассемблера и C. Считается одной из самых быстрых среди себе
равных, как при вычислениях с операндами малой разрядности, так и
при расчётах с большим количеством разрядов, благодаря
оптимизированному коду и использованию асимптотически
оптимальных алгоритмов.

Функции в GMP поделены на категории для удобства использования:

  - mpz - высокоуровневые арифметические и логические функции для
    операций над знаковыми целыми числами (около 140).
  - mpq - высокоуровневые арифметические функции для операций над
    рациональными дробями (около 35).
  - mpf - высокоуровневые арифметические функции для операций над
    числами с плавающей точкой (около 65) для вычислений с
    точностью не ниже двойной.
  - C++ классы для интерфейса к функциям категорий mpz, mpq, mpf.
  - mpn - низкоуровневые функции для операции с беззнаковыми целыми
    числами с менее удобным вызовом, но низкими накладными
    расходами. Используются функциями вышеперечисленных
    категорий.
  - mpfr - отдельно поддерживаемые и разрабатываемые высокоуровневые
    арифметические функции для операций над числами с плавающей
    точкой с высокоточным округлением.

#### GSL — GNU Scientific Library

[GSL - GNU Scientific Library](http://www.gnu.org/software/gsl/)
написана на C и содержит значительное количество функций - от
элементарных математических операций и операций с комплексными
числами до численных методов дифференцирования, интерполяции,
аппроксимации, решения дифференциальных уравнений,
wavelet-преобразования и многих других.

#### MPFR

[MPFR](http://www.mpfr.org/) — C-шная библиотека для расчётов с
произвольной точностью над числами с плавающей точкой с
высокоточным округлением.

#### MPFR++ и MPFI

[MPFR++ MPFI](http://perso.ens-lyon.fr/nathalie.revol/software.html) —
C++ интерфейс к MPFR и основанная на MPFR библиотека интерфальной
арифметики с произвольной точностью соответственно.

#### MPIR

[MPIR](http://www.mpir.org/) — библиотека для вычислений с произвольной
точностью, форк GMP. Главные особенности: поддержка сборки при помощи
Microsoft Visual Studio 2010 (32-bit and 64-bit) и полная интерфейсная
совместимость с GMP.

#### LAPACK

[LAPACK](http://www.netlib.org/lapack/) — библиотека на Fortran 90 для
решения различных систем линейных уравнений, поиска решений линейных
систем уравнений методом наименьших квадратов, задач на собственные
значения (краевые задачи), сингулярных задач. Особенность состоит в
оптимизации по сравнению с EISPACK или LINPACK в операциях с памятью и
при распараллеливании.

#### LinBox

[LinBox](http://www.linalg.org/) — C++ библиотека для точных
высокопроизводительных алгебраических расчётов с плотными,
разрежёнными и структурированными матрицами над целыми числами и
конечными полями.

#### NTL

[NTL](http://www.shoup.net/ntl/) — библиотека на С++, содержащая
структуры и алгоритмы для операций с знаковыми целыми числами
произвольной точности, а также векторами, матрицами и полиномами над
целыми числами и конечными полями.

#### BLAS

[BLAS (Basic Linear Algebra Subprograms)](http://www.netlib.org/blas/) —
библиотека для операций над векторами и матрицами на Fortran 77 (да-да).
Если вы не страдаете подобным некрофильством, то есть [версии BLAS от
производителей аппаратных
платформ](http://www.netlib.org/blas/faq.html#5). Также не стоит
забывать о более современных
[ATLAS](http://math-atlas.sourceforge.net/), [Goto
BLAS2](http://www.tacc.utexas.edu/tacc-projects/gotoblas2/),
поддерживающих многопоточность.

#### GotoBLAS2

[GotoBLAS2](http://www.tacc.utexas.edu/tacc-projects/gotoblas2/) и
GotoBLAS — вариации библиотеки BLAS от Texas Advanced Computing Center.
Сейчас уже не разрабатывается, но исходный код открыт под лицензией
BSD.

#### ATLAS

[Automatically Tuned Linear Algebra Software
(ATLAS)](http://math-atlas.sourceforge.net/) предоставляет интерфейсы на
C и Fortran77 к кроссплатформенной производительной реализации BLAS, а
также некоторым функциям из LAPACK.

#### Eigen

[Eigen](http://eigen.tuxfamily.org/) — библиотека для операций с
матрицами любых размеров, содержащих целые числа, числа с
плавающей запятой, комплексные числа.

#### ESSL

[ESSL](http://publib.boulder.ibm.com/infocenter/clresctr/vxrx/index.jsp?topic=/com.ibm.cluster.essl.doc/esslbooks.html)
— Engineering and Scientific Subroutine Library (ESSL) and Parallel ESSL
от IBM это коллекция специально спроектированных функций для инженерных
и научных расчётов на серверах и блейд-серверах с процессорами IBM
POWER™.

#### IML

[IML](http://www.cs.uwaterloo.ca/~astorjoh/iml.html) — библиотека
реализованных на C алгоритмов для расчёта точных решений плотных
систем целочисленных линейных уравнений. IML показано использовать
совместно с ATLAS/BLAS и GMP.

#### GLPK

[GLPK (GNU Linear Programming
Kit)](http://www.gnu.org/software/glpk/glpk.html) — набор библиотек на
ANSI C для решения задач масштабного линейного программирования
(large-scale linear programming), смешанного целочисленного
программирования и других. Поддерживает язык моделирования
GNU MathProg.

#### SciPy

[SciPy](http://www.scipy.org/) — одна из двух базовых для Python
библиотек для научных вычислений. Содержит модули для
статистических расчётов, оптимизации, численного
интегрирования, решения дифференциальных уравнений,
линейной алгебры, преобразования Фурье, обработки сигналов,
обработки изображений и прочих функций.

#### NumPy

[NumPy](http://numpy.org/) — одна из двух базовых для Python библиотек
для научных вычислений. Содержит функции для операций с многомерными
массивами и матрицами, а также стандартные функции линейной алгебры,
преобразования Фурье, генераторы случайных чисел и инструменты для
интеграции C/C++ и Fortran кода.

#### mpmath

[mpmath](http://code.google.com/p/mpmath/) — библиотека на Python для
арифметических операций произвольной точности над числами с плавающей
точкой. Местами использует асимптотически оптимальные алгоритмы, при
наличии [gmpy](http://code.google.com/p/gmpy/) ускоряет операции за
счёт использования GMP/MPIR вместо собственной реализации.

#### gmpy

[gmpy](http://code.google.com/p/gmpy/) — wrapper к GMP на Python.

#### Pynac

[Pynac](http://pynac.sagemath.org/) — аналог C++'ового GiNaC на Python
для операций с символьными выражениями.

#### SymPy

[SymPy](http://code.google.com/p/sympy/) — библиотека на Python для
символьной математики.

#### Pyrex

[Pyrex](http://www.cosc.canterbury.ac.nz/greg.ewing/python/Pyrex/) —
фактически язык, который представляет смесь C и Python, позволяя,
например, смешивать типы данных и собирать результат в C-шное
расширение для Python.

#### RPy

[RPy](http://rpy.sourceforge.net/) — биндинг к R для Python

#### CVXOPT

[CVXOPT](http://abel.ee.ucla.edu/cvxopt/) — пакет для выпуклой
оптимизации на Python.

#### FLINT

[FLINT: Fast Library for Number Theory](http://www.flintlib.org/) —
быстрая, оптимизированная в том числе для многоядерных и
многопроцессорных систем, библиотека функций на C для теории
чисел.

#### PolyBoRi

[PolyBoRi](http://polybori.sourceforge.net/) — C++ библиотека с
интерфейсом на Python предоставляющая высокоуровневые типы
данных для работы с булевыми многочленами и одночленами и др.,
основываясь на [CUDD: CU Decision Diagram
Package](http://vlsi.colorado.edu/~fabio/CUDD/) и
[M4RI(e)](http://m4ri.sagemath.org/).

#### GAP

[GAP - Groups, Algorithms, Programming](http://www.gap-system.org/)

#### Givaro

[Givaro](http://ljk.imag.fr/CASYS/LOGICIELS/givaro/) — C++ библиотека
для различных арифметических и алгебраических операций: от арифметики
с произвольной точностью над целыми числами и расчётов с дробями до
дейстий с векторами и матрицами. Основан на GMP.

#### GiNaC

[GiNaC](http://www.ginac.de/) — (GiNaC is Not a CAS), предназначенна для
неинтерактивных операций с символьными математическими выражениями.

#### CLN

[CLN](http://www.ginac.de/CLN/) — C++ библиотека для операций с
произвольной точностью над целыми, рациональными дробными,
комплексными числами, числами с плавающей точкой. Оборудована
сборщиком мусора, умножением Карацубы и Schönhage-Strassen.

#### jsMath

[jsMath](http://www.math.union.edu/~dpvc/jsMath/) — метод внедрения
математических записей в веб-страницы. В зависимости от наличия
шрифтов (лучше TeX шрифты, но можно и Unicode) выдаёт результат
немного отличающегося но приемлимого качества в большинстве
браузеров.

#### M4RI(e)

[M4RI(e)](http://m4ri.sagemath.org/) названа в честь "Method of the Four
Russians" multiplication algorithm by Gregory Bard. Представляет из себя
библиотеку для быстрой арифметики над плотными матрицами через F₂.

#### SWIG

[SWIG (Simplified Wrapper and Interface
Generator)](http://www.swig.org/) — этот программный продукт позволяет
подключать написанные на C/C++ программы к ряду других языков
программирования.

#### PARI/GP

[PARI/GP](http://pari.math.u-bordeaux.fr/) — C-шная библиотека и
интерактивная консоль, прдоставляющая функционал системы
компьютерной алгебры (CAS).

#### zn\_poly

[zn\_poly](http://cims.nyu.edu/~harvey/code/zn_poly/)

#### PDL

[PDL (Perl Data Language)](http://pdl.perl.org/) — основанный на Perl
язык, предназначенный для операций с N-мерными массивами данных.
Стремится стать заменой MATLAB или IDL, поставляется вместе со
встроенными интерактивными оболочками.

#### Tachyon

[Tachyon](http://jedi.ks.uiuc.edu/~johns/raytracer/) — Multiprocessor
Ray Tracing System (зачем я её сюда добавил?)

## Симуляторы и средства имитационного моделирования

### Общего назначения

#### Xcos (Scilab)

[Xcos](http://www.scilab.org/products/xcos) — аналог Simulink MATLAB'a в
Scilab, основан на Scicos, заменил его с версии Scilab 5.2, имеет
некоторые улучшения в плане эргономики, цикл разработки
синхронизирован с циклом разработки Scilab.

#### Scicos (Scilab)

[Scicos](http://www-rocq.inria.fr/scicos/) — аналог Simulink MATLAB'a в
Scilab.

### Электроника

#### Qucs

[Qucs](http://qucs.sourceforge.net/) — свободный симулятор электрических
цепей, пока что на Qt3.

#### QucsStudio

[QucsStudio](http://www.mydarc.de/DD6UM/QucsStudio/qucsstudio.html) —
свежий форк Qucs, правда пока только для Windows.

#### KTechlab

[KTechlab](http://sourceforge.net/apps/mediawiki/ktechlab/) — довольно
медленно, но верно развивающаяся по направлению к KDE4 и Qt4 программа
для симуляции электрических цепей и микроконтроллеров (пока что PIC от
MICROCHIP) в этих цепях (да, прямо как в Proteus).

#### SPICE

SPICE — это целый класс и не одно поколение численных симуляторов
аналоговых электрических цепей.

##### ngspice

[ngspice](http://ngspice.sourceforge.net/)

##### SPICE3F5

[SPICE3F5](http://embedded.eecs.berkeley.edu/pubs/downloads/spice/spice.html)

##### CIDER 1B1

[CIDER 1B1](http://www.eecs.berkeley.edu/IPRO/Software/Description/cider1b1.html)

##### XSPICE

[XSPICE](http://users.ece.gatech.edu/~mrichard/Xspice/)

#### QSapecNG

[QSapecNG](http://qsapecng.sourceforge.net/) — символьный симулятор
линейных электрических цепей, преемник SAPWIN (Symbolic Analysis
Program for Windows), переписанный на Boost C++ по модульному принципу с
графической оболочкой на Qt. В отличие от симуляторов семейства SPICE
(Simulation Program with Integrated Circuit Emphasis), использующих
численные методы, производит расчёты на базе аналитических
выражений, с использованием преобразования Лапласа и
операционного исчисления.

#### Logisim

[Logisim](http://ozark.hendrix.edu/~burch/logisim/) — достаточно простое
средство для симуляции цифровых электричечких схем, написанное на Java.
Годится для демонстрации при обучении основам цифровой схемотехники.

#### NuSMV

[NuSMV](http://nusmv.fbk.eu/) — приложение для символической проверки
моделей, базирующееся на BDD (Binary decision diagram).

#### DLSim 3

[DLSim 3](http://www.cs.oberlin.edu/~rms/dlsim.com/) — бесплатный для
образовательных целей, требуется регистрация
[1](http://www.cs.oberlin.edu/~rms/dlsim.com/download/DLSim3b305r5.zip),
[2](http://www.cs.oberlin.edu/~rms/dlsim.com/download_sdk/DLSim3b305r5-SDK.zip).

#### KLogic

[KLogic](http://www.a-rostin.de/) — медленно, но верно портируемое на
KDE4 приложение для симуляции цифровых цепей как и положено с функцией
построения временных диаграмм.

#### TKGate

[TKGate](http://www.tkgate.org/) — цифровой схемный симулятор с
интерфейсом на Tcl/Tk с поддержкой скриптов, различных моделей
логических уровней, настраиваемыми моделями задержек, точками останова,
построением временных диаграмм и многими другими вкусностями. Ждём
выхода версии 2.0.

#### Electric

[Electric](http://www.gnu.org/software/electric/) — среда разработки
принципиальных электрических и интегральных схем и их симуляции.
Написана на Java.

#### gpsim

[gpsim](http://gpsim.sourceforge.net/) — программный симулятор для
микроконтроллеров PIC от Microchip.

#### SimCAS

[SimCAS](http://sourceforge.net/projects/simcas/) — аналитический
симулятор аналоговых электрических схем.

### Языки описания аппаратуры (HDL)

#### Verilator

[Verilator](http://www.veripool.org/wiki/verilator/) — один из
быстрейших симуляторов Verilog HDL. Как указывает автор,
предназначен для больших проектов, где важна быстрая симуляция, а
особенно хорош для создания исполняемых моделей ЦПУ при проектировании
встраиваемых систем.

#### Icarus Verilog

[Icarus Verilog](http://iverilog.icarus.com/) — средство симулирования и
синтезирования для языка Verilog HDL.

#### ISE Design Suite (Xilinx)

[ISE Design Suite](http://www.xilinx.com/support/download/index.htm) —
проприетарная IDE для разработки на языках HDL для ПЛИС Xilinx.

#### Quartus II (Altera)

[Quartus
II](http://www.altera.com/products/software/quartus-ii/web-edition/qts-we-index.html)
— проприетарная IDE для разработки на языках HDL для ПЛИС Altera.

#### FreeHDL

[FreeHDL](http://freehdl.seul.org/) — часть Qucs, симулятор VHDL.

#### GHDL

[GHDL](http://ghdl.free.fr/) — симулятор VHDL.

#### zamiaCAD

[zamiaCAD](http://zamiacad.sourceforge.net/) — модульная среда
разработки на языках описания аппаратуры (HDL). Поддерживается
разработка на VHDL 2003, поддержку парсера и синтаксического дерева VHDL
2003 и Verilog 2005, имеется симулятор, средство построения временных
диаграмм.

#### VeriWell

[VeriWell](http://sourceforge.net/projects/veriwell/) — симулятор
Verilog HDL.

### Физические

#### Step (KDE Edu)

[Step](http://edu.kde.org/applications/all/step/) (KDE) — часть проекта
KDE Edu, умеет симулировать механику, пружины, гравитацию, кулоновские
силы и молекулярную динамику жидкостей и газов.

#### Algodoo

[Algodoo (Phun)](http://www.algodoo.com/) — замена симулятора Phun.
Доступна демка, по словам разработчиков, в Algodoo много улучшений
касательно физического движка, возможностей симулятора и визуализации.

## Системы автоматизированного проектирования

### Toped

[Toped](http://www.toped.org.uk/) — кроссплатформенный редактор макетов
интегральных схем, поддерживающий форматы GDS, OASIS и CIF.

### LayoutEditor

[LayoutEditor](http://www.layouteditor.net/) — приложение-редактор
проектов для производства MEMS/IC (Microelectromechanical
systems/Integrated circuit - микроэлектромеханических
систем/интегральных схем). Также может быть
использован для разработки Multi-Chip-Modules (MCM),
Chip-on-Board (COB), Low temperature co-fired ceramics (LTCC),
Monolithic Microwave Integrated Circuits (MMIC), printed circuit boards
(PCB), thick film technology. Имеет интегрированный SchematicEditor для
разработки принципиальных электрических схем.

### OpenSCAD

[OpenSCAD - The Programmers Solid 3D CAD
Modeller](http://www.openscad.org/) — неинтерактивная среда
проектирования сплошных 3D деталей. Не такая живописная и
красочная как [Blender](http://www.blender.org/), больше подходит для
технических задач. OpenSCAD использует Qt4, библиотеки
[OpenCSG](http://www.opencsg.org/) и [CGAL](http://www.cgal.org/). Для
создания 3D модели необходимо написать скрипт с описанием объекта и
затем запустить рендеринг. Работает одним из двух способов: либо
полностью описание объекта скриптом, либо DXF/STL/OFF файл
двумерного объекта и соответствующая обработка этого "сечения"
скриптом.

### gEDA

[gEDA](http://www.gpleda.org/) — GPL Electronic Design Automation,
комбайн, объединивший несколько отдельных программ для
проектирования электроники в общем цикле разработки. В gEDA
входят:

  - **gschem** — разработка принципиальных электрических схем (schematic
    capture)
  - **gnetlist** — работа с netlist'ами.
  - **gattrib** — работа с номенклатурой компонентов
  - **[gedasymbols.org](http://gedasymbols.org/)** — ресурс-репозиторий
    для обмена и расшаривания разработанных компонентов
  - **[gerbv](http://gerbv.sourceforge.net/)** — просмотрщик для Gerber
    RS-274X, файлов Excellon drill и CSV pick-and-place.
  - **[PCB](http://pcb.gpleda.org/)** — редактор печатных плат

### KiCad EDA Suite

[KiCad EDA Suite](http://kicad.sourceforge.net/) — програмный пакет для
автоматизированного проектирования принципиальных электрических схем и
печатных плат. В его состав входят:

  - **EESchema** - входящий в состав KiCad EDA Suite редактор
    принципиальных электрических схем. [Годное руководство по
    Eeschema](ftp://ftp.ntcsm.ru/pub/kicad/doc/odt/eeschema.odt) от
    KiCad Russian Team.

<!-- end list -->

  - **CvPCB** - входящий в состав KiCad EDA Suite редактор связей
    "компонент принципиальной схемы" - "корпус и геометрия выводов
    на печатной плате". [Годное руководство по
    CvPCB](ftp://ftp.ntcsm.ru/pub/kicad/doc/ru_ru/cvpcb.odt) от KiCad
    Russian Team.

<!-- end list -->

  - **PCBNew** - входящий в состав KiCad EDA Suite редактор печатных
    плат [Годное руководство по
    PCBNew](ftp://ftp.ntcsm.ru/pub/kicad/doc/ru_ru/pcbnew.pdf) от KiCad
    Russian Team.

<!-- end list -->

  - **GerbView** - входящий в состав KiCad EDA Suite просмотрщик Gerber
    файлов формата RS-274X. [Годное руководство по
    GerbView](ftp://ftp.ntcsm.ru/pub/kicad/doc/ru_ru/gerbview.odt) от
    KiCad Russian Team.

<!-- end list -->

  - **Bitmap2Component** - входящий в состав KiCad EDA Suite редактор
    изображений элементов.

### EAGLE

[EAGLE by CadSoft](http://www.cadsoft.de/) — коммерческое ПО для
разработки печатных плат, в том числе для Linux. EAGLE Light
Edition для некоммерческих организаций с ограничениями проектирования
одно/двухсторонней платы размером не более 100 x 80 мм, максимум
одним листом схемы и поддержкой лишь по почте и через форум
доступен бесплатно.

### atlc

[atlc - Arbitrary Transmission Line Calculator (for transmission lines
and directional couplers)](http://atlc.sourceforge.net/) — специфичная
среда, предназначенная для расчёта распределённых линий с довольно
оригинальным способом ввода данных. У кого читали математическую
физику на электрофаке, те, думаю, оценят возможности утилиты.

### QCAD

[QCAD](http://www.qcad.org/qcad.html) — кроссплатформенное приложение
для создания 2D-чертежей, имеющее как коммерческую платную версию,
так и свободную открытую community-версию. Свободно снабжается
большим количеством примеров, готовых чертежей деталей и
запчастей (более 4500) и документацией.

### FreeCAD

[FreeCAD](http://sourceforge.net/apps/mediawiki/free-cad/) — открытая
среда 3D проектирования, создания чертежей и моделей, с интерфейсом
на Qt и геометрическим ядром Open CASCADE и трёхмерным движком Coin
3D. Имеет модульную архитектуру, Python API, импорт/экспорт из/в STEP,
IGES, OBJ, DXF, SVG, U3D, STL.

### BRL-CAD

[BRL-CAD](http://brlcad.org/) — многофункциональная кроссплатформенная
система автоматизированного проектирования и моделирования объёмных
тел методами конструктивной блочной геометрии (CSG - Constructive
Solid Geometry) с более чем 20-летней историей.

### LibreCAD

[LibreCAD](http://www.librecad.org/) — ещё одна среда для разработки
2D-чертежей.

### Ayam

[Ayam](http://ayam.sourceforge.net/) — free 3D modelling environment for
the RenderMan interface.

### XTrackCAD

[XTrackCAD](http://www.xtrkcad.org/) — есть в мире OpenSource и такое -
система автоматизированного проектирования железнодорожных путей.

### Archimedes

[Archimedes](http://www.archimedes.org.br/) — свободная CAD написанная
на Java из Бразилии. В скриншотах почему-то приведены примеры
проектирования комнаты, хотя, лично проверил, умеет он не
только это.

### Sweet Home 3D

[Sweet Home 3D](http://www.sweethome3d.com/) — свободная САПР для
дизайна интерьеров на Java в 2D с 3D просмотром. В какой-то
степени покрывает функционал проприетарной
[PRO100](http://uk.ecrusoftware.com/pro100) для дизайна мебели и
интерьеров.

### OpenCASCADE

[OpenCASCADE](http://www.opencascade.org/) — 3D modeling & numerical
simulation. Требует регистрацию для загрузки. Напишите отзыв.

### VariCAD

[VariCAD](http://varicad.com/) — коммерческий 2D/3D CAD. Умеет экспорт
STEP (3D), STL (3D), IGES (3D and 2D), DWG (2D), DXF (2D) и импорт STEP
(3D), DWG (2D), DXF (2D), IGES (2D). Стоит относительно недорого. Есть
30-дневный триал.

### DraftSight

[DraftSight by Dassault Systèmes
(DS)](http://www.3ds.com/products/draftsight/) — бесплатный 2D CAD в
Beta стадии с официальной поддержкой 32-bit Ubuntu, Fedora, openSUSE и
Mandriva. Требует бесплатную активацию в течение 30 дней через интернет
и последующую реактивацию через 6 месяцев и затем каждый год. Открывает
DWG/DXF файлы и экспортирует в PDF.

### MEDUSA4 Personal

[MEDUSA4
Personal](http://www.cad-schroer.com/Software/MEDUSA4/CADFreeware/) —
бесплатный 2D/3D CAD для частного некоммерческого использования.
Требует бесплатную регистрацию и активацию каждые 6 месяцев. Имеет
достаточно интересную возможность получить платное разрешение для
коммерческого использования конкретного чертежа.

### Bricscad

[Bricscad от ООО Кадсофт (Волгоград)](http://www.brics-cad.ru/) —
коммерческий 2D/3D CAD с поддержкой формата DWG с версией для
Linux. Доступная 30-дневная trial версия.

### Siemens NX 7

[Siemens NX 7](http://www.plm.automation.siemens.com/ru_ru/products/nx/)
— коммерческий CAD/CAE/CAM от Siemens. Имеется версия для Linux/Unix.

### Autodesk® Software

[Autodesk®](http://www.autodesk.com/products-all) — как ни странно, но у
Autodesk® есть коммерческие продукты под Linux, кроме web-based, правда
немного и почти все они так или иначе связаны с 3D и рекомендуемая ОС,
естественно, RHEL или на худой конец Fedora. Связана такая печальная
ситуация с поддержкой Linux с тем, что Autodesk® не гнушается
использовать для своих мощных продуктов .NET и при этом писать
нехилые требования к оборудованию.

Перечень на июнь 2011 года (информация о поддержке Linux может
изменяться от версии к версии - могут как убрать, что более
вероятно, так и добавить, что менее вероятно):

  - [Autodesk® Maya® (64 bit)](http://www.autodesk.com/maya)
  - [Autodesk® Flare® (64 bit)](http://www.autodesk.com/flare)
  - [Autodesk® FBX® (64 bit)](http://www.autodesk.com/fbx)
  - [Autodesk® Mudbox™ (64 bit)](http://www.autodesk.com/mudbox)
  - [Autodesk® Softimage®](http://www.autodesk.com/softimage)
  - [Autodesk® Infrastructure Map Server (32 bit - только
    сервер)](http://www.autodesk.com/infrastructuremapserver)
  - [Autodesk® Moldflow® Insight®
    (64-Bit)](http://usa.autodesk.com/moldflow/)
  - [Autodesk® Simulation® (32, 64 bit
    recommended)](http://www.autodesk.com/algor).

## SCADA системы

Тут прямо скажем и выбор невелик, и названия какие-то однообразные. Как
заметил автор первого проекта, Роман Савоченко в одном из интервью:
"Тёзки". :)

### OpenSCADA

[OpenSCADA](http://oscada.org/ru/) — разработка Романа Савоченко из
Днепродзержинска представляет собой уже вполне надёжную SCADA
систему с модульной организацией и графическим интерфейсом,
разрабатываемым с использованием Qt. К LTS версии 0.7.0
реализована поддержка протоколов OPC\_UA, ModBUS, SNMP, DCON,
собственного протокола OpenSCADA на физических последовательных
интерфейсах (RS232, RS485, Modem ...) и Ethernet, а также баз
данных и ряда конкретного оборудования.

*FIXME: Приглашаем Романа Савоченко собственноручно исправить или
дополнить вышенаписанное*

### openSCADA

[openSCADA](http://openscada.org/) — эта SCADA система разрабатывается
на Java с тестовым клиентом на GTK разработчиками из Германии во главе
с Йенсом Райменом (Jens Reimann) и Юргеном Розе (Jürgen Rose). Проект
состоит из двух частей: Atlantis, содержащий имплементацию openSCADA
интерфейсов на Java к таким внешним системам как Siemens S7 PLC, OPC,
SNMP, реляционные базы данных и т. п., и Utgard - 100% pure JAVA OPC
Client API с поддержкой интерфейса OPC DA 2.0 для соединения с OPC
сервером.

## Системы визуализации данных

### Построение 2D/3D графиков/диаграмм

#### gnuplot

[gnuplot](http://www.gnuplot.info/) — кроссплатформенная консольная
утилита для построения 2D и 3D графиков.

Not so Frequently Asked Questions по gnuplot -
[2](http://t16web.lanl.gov/Kawano/gnuplot/index-e.html).

##### Графические фронтенды к gnuplot

[Qgfe](http://web.archive.org/web/20090609160201/http://www.xm1math.net/qgfe/)
(Qt3) - очень старый, но рабочий проект Дэвида Иши (David Ishee),
подойдёт для начального ознакомления с функционалом gnuplot.

[UniGNUPlot](http://unicalculus.sourceforge.net/) (Tcl/Tk) - проект
также заброшен, но само приложение всё ещё рабочее.

##### Интерфейсы и биндинги

[Gnuplot.py](http://gnuplot-py.sourceforge.net/) — интерфейс к gnuplot
для Python.

[JavaPlot](http://gnujavaplot.sourceforge.net/) — интерфейс к gnuplot
для Java.

Информация по интерфейсам к gnuplot на ANSI C, C++, Perl, Fortran
доступна собственно на сайте
[gnuplot](http://www.gnuplot.info/links.html) и на сайте [Николаса
Девиларада (Nicolas Devillard)](http://ndevilla.free.fr/gnuplot/),
разработчика интерфейса к gnuplot на ANSI C.

*FIXME: Собрать ответы на вопросы по gnuplot в отдельную статью
[gnuplot](gnuplot "wikilink")*

#### QtiPlot

[QtiPlot](http://soft.proindependent.com/qtiplot.html) (Qt4) -
кроссплатформенное приложение базирующееся на Qt4 и
библиотеках Qwt (класс QwtPlot) и QwtPlot3D для построения 2D
и 3D графиков с функциями аппроксимации, анализа построенных графиков,
интегрирования LaTeX, экспорта в различные графические форматы. По
функционалу близок к [Origin](http://www.originlab.com/). В платной
версии доступна поддержка скриптов на Python, что позволяет
использовать QtiPlot в связке с SciPy, NumPy, SymPy и т. д.

#### SciDAVis

[SciDAVis](http://scidavis.sourceforge.net/) — форк
[QtiPlot](http://soft.proindependent.com/qtiplot.html), базирующееся на
Qt4 приложение для построения 2D и 3D графиков, круговых и столбчатых
диаграмм. Данные для построений могут быть введены как интерактивно,
так и импортированы из ASCII файлов или вычислены при помощи
стандартных функций или функций Python. Как и многие
приложения этого класса, позволяет экспортировать результаты в
различные форматы, от растровых и векторных графических до EPS и PDF.
Проект сотрудничает с [LabPlot](http://labplot.sourceforge.net/),
работая над созданием общего бэкенда, оставаясь при этом
независимым от библиотек KDE.

#### LabPlot

[LabPlot](http://labplot.sourceforge.net/) (KDE) — приложение для
анализа и визуализации данных для KDE с большим перечнем
возможностей, среди которых построение 2D, 3D графиков,
поверхностей, экспорт и импорт как результатов построений, так и
данных и многое другое. По функционалу напоминает проприетарный пакет
[Origin](http://www.originlab.com/) и умеет импортировать файлы его
проектов.

#### Veusz

[Veusz](http://home.gna.org/veusz/) — основанный на Qt4 пакет для
визуализации. Написан на Python с использованием PyQt4 для
построений и графического интерфейса пользователя, и NumPy для
манипуляций с числовыми данными. Veusz разработан для получения
Postscript/PDF/SVG файлов качества, приемлимого для сопровождения
публикаций графическим материалом.

#### Kst

[Kst](http://kst-plot.kde.org/) — ПО для визуализации данных в KDE.

#### ROOT

[ROOT](http://root.cern.ch/) — разрабатываемое CERN на C++ приложение и
набор библиотек для визуализации данных. Содержит один из самых полных
наборов функций и методов, которые могут понадобиться для построения
результатов научных экспериментов, расчётов и анализа данных,
несмотря на свою узкую специализацию для задач физики.

#### GSEGrafix

[GSEGrafix](http://www.gnu.org/software/gsegrafix/) (GNOME) - приложение
для построения 2D и 3D графиков в GNOME. Графики строятся с
использованием сглаживания при помощи виджета GNOME Canvas.
Приложение может считывать данные для построения из ASCII файлов и
вызываться из внешних программ и скриптов
[Octave](http://www.gnu.org/software/octave/).

#### RLPlot

[RLPlot](http://rlplot.sourceforge.net/) — приложение для построения
графиков, диаграмм, гистрограмм и др. на основе табличных данных.
Имеет функцию экспорта в несколько растровых графических форматов,
векторный SVG и в EPS.

#### Asymptote

[Asymptote](http://asymptote.sourceforge.net/) — is a powerful
descriptive vector graphics language that provides a natural
coordinate-based framework for technical drawing. Labels and equations
are typeset with LaTeX, for high-quality PostScript output.

#### GLE

[GLE (Graphics Layout Engine)](http://glx.sourceforge.net/) — is a
graphics scripting language designed for creating publication quality
figures (e.g., a chart, plot, graph, or diagram).

#### OpenDX

[OpenDX](http://www.opendx.org/) — открытая свободная версия IBM
Visualization Data Explorer для 3D визуализации. После выхода версии
3.1.4B IBM открыл исходники VDE и пригласил к разработке сообщество.
Графический интерфейс основан на OSF/Motif и X Window System.

#### CLUCalc

[CLUCalc](http://www.clucalc.info/) 4.3.3 - software tool for 3D
visualizations and scientific calculations free for non-commercial use.

#### MayaVi и Mayavi2

[MayaVi](http://mayavi.sourceforge.net/)
[Mayavi2](http://code.enthought.com/projects/mayavi/) — две генерации
средства 3D визуализации на Python.

#### VisIt

[VisIt](https://wci.llnl.gov/codes/visit/) — инструмент для визуализации
данных, специализация на обработке больших объёмов данных, в том числе
скалярных и векторных полей, с использованием распараллеливания
вычислений, а также возможность реализации особенных алгоритмов
обработки и работы с нестандартными моделями данных. VisIt имеет
мощный графический пользовательский интерфейс, также имеются
интерфейсы на C++, Python и Java. Для сборки нужны библиотеки
[VTK 5.0.0i](http://www.vtk.org/),
[HDF4 4.2.5](http://www.hdfgroup.org/products/hdf4/),
[Xdmf 2.1.1](http://www.xdmf.org/) версий не ниже, чем указаны.

#### ParaView

[ParaView](http://www.paraview.org/) — кроссплатформенное приложение для
анализа и визуализации данных с упором на обработку экстремально больших
объёмов данных в системах с распределённой памятью. Тем не менее никто
не запрещает использовать его для небольших наборов данных.

#### VTK, VTKEdge, ITK

[VTK](http://www.vtk.org/), [VTKEdge](http://www.vtkedge.org/),
[ITK](http://www.itk.org/) — продукты компании Kitware для визуализации,
обработки изображений, 3D графики. Базовым является VTK, предоставляющий
C++ классы и интерфейсы на Tcl/Tk, Java и Python. VTKEdge и ITK
базируются на VTK и расширяют функционал.

#### Grace

[Grace](http://plasma-gate.weizmann.ac.il/Grace/) — WYSIWYG 2D plotting
tool for the X Window System and M\*tif.

#### GD

[GD](https://bitbucket.org/pierrejoye/gd-libgd/overview)
[1](http://www.boutell.com/gd/)- GD написан на C, однако биндинги
доступны для Perl, PHP, Python и др.

#### KmPlot (KDE Edu)

[KmPlot](http://edu.kde.org/applications/all/kmplot/) (KDE) — часть
проекта KDE Edu, позволяет строить графики функций, заданных
аналитически или параметрически, в прямоугольной или полярной
системах координат на плоскости. Умеет строить производные 1-го и
2-го порядков и интеграл от заданных функций. Есть функция экспорта в
один из растровых графических форматов (BMP and PNG), в векторный SVG
и в формате XML.

#### KAlgebra (KDE Edu)

[KAlgebra](http://edu.kde.org/kalgebra/) (KDE) — часть проекта KDE Edu,
позволяет строить 2D и 3D графики, заданные аналитически. Поддерживает
язык разметки MathML. Позволяет экспортировать графики в SVG и PNG
форматах.

### Геометрические построения

#### Kig (KDE Edu)

[Kig](http://edu.kde.org/kig/) (KDE) — часть проекта KDE Edu, позволяет
выполнять геометрически построения на плоскости. Является заменой
заброшенного проекта [KGeo](http://kgeo.sourceforge.net/).
Позволяет экспортировать построения в несколько растровых
графических форматов (BMP, JPEG, PNG), векторный SVG, формат
XFig, XML и даже LaTeX.

#### KSEG

[KSEG](http://www.mit.edu/~ibaran/kseg.html) (Qt3) - позволяет выполнять
геометрические построения на плоскости. Проект заброшен, хотя при этом
доступен во многих дистрибутивах и вполне работоспособен. Имеет
встроенный калькулятор с элементарными функциями. Может
экспортировать построения в один из нескольких растровых
графических форматов и векторный SVG.

#### Dr. Geo

[Dr. Geo](http://community.ofset.org/index.php/DrGeo)
([1](http://www.gnu.org/software/dr-geo/)) - удобное кроссплатформенное
приложение для выполнения геометрических построений на плоскости,
написанное на [Pharo
Smalltalk](http://www.pharo-project.org/home).

### Графы, блок-схемы, UML

#### Graphviz

[Graphviz](http://www.graphviz.org/) — удобное средство для построения
графов. Описание графа выполняется на языке DOT. Есть как консольное
приложение, так и простой графический интерфейс.

#### ArgoUML

[ArgoUML](http://argouml.tigris.org/) — среда UML моделирования с
поддержкой стандарта UML 1.4 написанное на Java.

#### Dia

[Dia](http://live.gnome.org/Dia) [1](http://projects.gnome.org/dia/)
(GTK+) - редактор блок-схем и диаграмм. Позиционируется как свободный
аналог Visio от Microsoft. Умеет экспортировать в большое количество
графических форматов, XML, поддерживает язык UML.

#### Umbrello

[Umbrello](http://uml.sourceforge.net/) — Umbrello UML Modeller is a
Unified Modelling Language diagram programme for KDE. Версия 2.\* в
KDE4.

#### ORA

[ORA](http://www.casos.cs.cmu.edu/projects/ora/) — dynamic meta-network
assessment and analysis tool developed by CASOS at Carnegie Mellon. It
contains hundreds of social network, dynamic network metrics, trail
metrics, procedures for grouping nodes, identifying local patterns,
comparing and contrasting networks, groups, and individuals from a
dynamic meta-network perspective.

#### Rocs (KDE Edu)

[Rocs](http://edu.kde.org/applications/all/rocs/) (KDE) — часть проекта
KDE Edu, среда разработки графов с поддержкой Qt Script.

#### Kivio

[Kivio](https://projects.kde.org/projects/koffice/repository/revisions/master/show/kivio)
— часть проекта [KOffice](http://www.koffice.org/kivio/). Ещё не доведён
до стадии релиза, потому не поставляется в стабильных архивах исходного
кода. Получить код можно, склонировав master-branch git-репозитория
KOffice или из установив из репозитория вашего дистрибутива.

#### OpenOffice.org Draw

[OpenOffice.org Draw](http://www.openoffice.org/product/draw.html) —
упомянут здесь исключительно для полноты картины.

#### LibreOffice Draw

[LibreOffice Draw](http://www.libreoffice.org/features/draw/) — упомянут
здесь исключительно для полноты картины.

#### yEd

[yEd](http://www.yworks.com/en/products_yed_about.html) — редактор
блок-схем с поддержкой UML на Java от компании
[yWorks](http://www.yworks.com/).

#### UMLet

[UMLet - UML Tool for Fast UML Diagrams](http://www.umlet.com/) — ещё
один редактор UML диаграмм на Java с простым интерфейсом и
возможностью экспорта в JPEG, SVG, EPS, PDF. Доступен как
плагин к Eclipse, так и отдельным приложением.

#### Violet

[Violet](http://violet.sourceforge.net/) — ещё один UML редактор на
Java.

#### NetworkX

[NetworkX](http://networkx.lanl.gov/) — Python package for the creation,
manipulation, and study of the structure, dynamics, and functions of
complex networks. (Взято [на
wiki.python.org](http://wiki.python.org/moin/PythonGraphApi))

#### Dunnart

[Dunnart Constraint-Based Diagram
Editor](http://www.csse.monash.edu.au/~mwybrow/dunnart/) — весьма
скромно названный редактор диаграмм, пока с закрытым кодом, но с
версией под Linux, посоветованный уважаемым
[AP](http://www.linux.org.ru/people/AP/profile) в [этой
теме](http://www.linux.org.ru/forum/talks/6432208).

#### Gaphas

[Gaphas](http://gaphor.sourceforge.net/) — Python-based diagramming
widget for GTK+.

#### Adaptagrams

[Adaptagrams](http://adaptagrams.sourceforge.net/) — библиотека для
применения в приложениях, связанных с адаптивным построением
диаграмм. Среди возможностей: инструменты для рисования,
автоматизированное создание структуры/макета
документа/диаграммы, отрисовка графов и макетов
диаграмм и др. Кстати, именно она используется в
[Dunnart](http://www.dunnart.org/),
[Inkscape](http://www.inkscape.org/),
[Graphviz](http://www.graphviz.org/),
[Arcadia](http://arcadiapathways.sourceforge.net/),
[Gaphas](http://gaphor.sourceforge.net/).

#### Gliffy

[Gliffy](http://www.gliffy.com/) — редактор диаграмм он-лайн.

### Диаграммы состояний конечных автоматов (FSM - Finite-state machine)

#### Qfsm

[Qfsm](http://qfsm.sourceforge.net/) (Qt)- графический редактор диаграмм
состояний конечных автоматов с интерактивной симуляцией, генерацией VHDL
кода, экспортом в растровые и векторные графические форматы, EPS, LaTeX,
HTML, обычный текст.

### Просмотр временных диаграмм

#### TimingAnalyzer

[TimingAnalyzer](http://www.timing-diagrams.com/) — написан на Java, но
выглядит очень красиво.

#### GTKWave

[GTKWave](http://gtkwave.sourceforge.net/) — GTK+ based wave viewer for
Unix and Win32 which reads LXT, LXT2, VZT, FST, and GHW files as well as
standard Verilog VCD/EVCD files and allows their viewing.

#### Dinotrace

[Dinotrace](http://www.veripool.org/wiki/dinotrace/) — Dinotrace is a
X-11 waveform viewer which understands Verilog Value Change Dumps,
ASCII, and other trace formats. It allows placing cursors, highlighting
signals, searching, printing, and other capabilities. Dinotrace comes
with a interface to Emacs which allows source code and log files to be
annotated with the values of signals.

#### Gwave

[Gwave](http://gwave.sourceforge.net/) — Gwave is a tool for viewing
analog data, such as the output of Spice simulations. Gwave can read
"raw" files from spice2G6, spice3F5 or ngspice, and a tabular ASCII
format suitable for use with GnuCAP or homegrown tools. It can also read
several binary and ascii files written by commercial spice-type
simulators such as hspice, tspice, and nanosim.

### Библиотеки

#### matplotlib

[matplotlib](http://matplotlib.sourceforge.net/) — библиотека языка
Python для построения разнообразных 2D графиков. Может использоваться
как в Python-скриптах, так и в интерактивной среде iPython. Может
быть расширена за счёт ряда тулкитов, в том числе mplot3d для 3D и
Natgrid для разреженных сеток.

#### MathGL

[MathGL](http://mathgl.sourceforge.net/) — библиотека для построения
более 55 видов 2D и 3D графиков, разрабатываемая [Алексеем
Абалакиным aka
abalakin](http://www.linux.org.ru/people/abalakin/profile). Имеет Qt,
FLTK, OpenGL интерфейсы, может быть вызвана из
C++/C/Fortran/Python/Octave и других программ и выполнять экспорт в
растровые и векторные графические форматы. Поддерживает собственный
скриптовый язык MGL.

*FIXME: [Алексей Абалакин aka
abalakin](http://www.linux.org.ru/people/abalakin/profile), исправьте
или добавьте, пожалуйста, что считаете нужным.*

##### Графические фронтенды

[UDAV](http://udav.sourceforge.net/) — кроссплатформенный графический
фронтенд к MathGL.

#### Qwt

[Qwt](http://qwt.sourceforge.net/) — библиотека классов и компонентов
графического интерфейса пользователя, удобных для построения графиков
в приложениях, базирующихся на Qt.

##### Интерфейсы и биндинги

[PyQwt](http://pyqwt.sourceforge.net/) — биндинги к Qwt для Python.

[Korundum/QtRuby](http://rubyforge.org/projects/korundum/) — проект, в
том числе содержащий биндинги Qwt для Ruby.

##### Дополнения

[QwtPolar](http://qwtpolar.sourceforge.net/) — дополнение к Qwt,
библиотека классов для отображения данных в полярной системе
координат.

#### QwtPlot3D

[QwtPlot3D](http://qwtplot3d.sourceforge.net/) — основанная на Qt и
OpenGL C++ библиотека для построение 3D графиков, поверхностей,
векторных полей.

#### DISLIN

[DISLIN](http://www.mps.mpg.de/dislin/) — надо проверить, она в
исходниках или только в бинарных пакетах, на сайте написано,
что free for non-commercial use, а значит не OpenSource.

#### CGAL

[CGAL - Computational Geometry Algorithms Library](http://www.cgal.org/)
— provides easy access to efficient and reliable geometric algorithms in
the form of a C++ library.

#### OpenCSG

[OpenCSG](http://www.opencsg.org/) — библиотека для рендеринга
изображений методами конструктивной блочной геометрии (CSG -
Constructive Solid Geometry) при помощи OpenGL.

#### GLEW

[GLEW - The OpenGL Extension Wrangler
Library](http://glew.sourceforge.net/) — a cross-platform open-source
C/C++ extension loading library. GLEW provides efficient run-time
mechanisms for determining which OpenGL extensions are supported on the
target platform.

## Среды разработки

### KDevelop

[KDevelop](http://kdevelop.org/)

### Lazarus

[Lazarus](http://www.lazarus.freepascal.org/)

### Eclipse

[Eclipse](http://eclipse.org/)

### NetBeans

[NetBeans](http://netbeans.org/)

### Geany

[Geany](http://www.geany.org/)

### GNU Emacs

[GNU Emacs](http://www.gnu.org/software/emacs/)

### Scintilla/SciTE

[Scintilla/SciTE](http://www.scintilla.org/)

### gVim

[gVim](http://www.vim.org/)

### Code::Blocks

[Code::Blocks](http://www.codeblocks.org/)

### Qt Creator IDE

[Qt Creator IDE](http://qt.nokia.com/products/developer-tools/)

### PyDev

[PyDev](http://pydev.org/) — Python IDE для Eclipse, может быть
использована для разработки Python, Jython и IronPython.

### Komodo IDE

[Komodo IDE](http://www.activestate.com/komodo-ide) — платная IDE для
Python, PHP, Ruby, JavaScript, Perl и Web Dev от ActiveState. Кто знает
чем она хуже или, может быть, лучше вышеупомянутых свободных -
добавьте.

## Проектирование баз данных и ER-модели данных

### Open System Architect (OSA)

[Open System Architect (OSA) by
CodeByDesign](http://www.codebydesign.com/) — свободное открытое ПО для
разработки и верификации ERD с встроенным SQL редактором и поддержкой
UML в стадии разработки.

### DBDesigner

[DBDesigner](http://www.fabforce.net/dbdesigner4/index.php) — среда для
проектирования структуры баз данных. Распространяется в бинарных
пакетах.

### DB Designer Fork

[DB Designer Fork](http://sourceforge.net/projects/dbdesigner-fork/) —
как ни странно, но это свободный форк вышеупомянутого DBDesigner от
fabFORCE, причём вполне живучий, если судить по дате последнего выпуска
и коммита. Поддерживает SQL скрипты БД Oracle, SQL Server, MySQL,
FireBird, SQLite и PostgreSQL.

### Visual Paradigm

[Visual Paradigm for UML](http://www.visual-paradigm.com/) — платный
инструмент для разработки ERD с использованием UML. Доступны
Enterprise Trial'ы и бесплатные Community Edition версии для
некоммерческого пользования.

### MySQL Workbench

[MySQL Workbench](http://wb.mysql.com/) — кроссплатформенная среда
проектирования баз данных для MySQL. С версии 5.2.34 (май 2011)
построена на Python 2.7, что, наверное, есть добрый знак для тех, кто
ей пользуется.

## Системы управления версиями (Source Control Management - SCM)

### Git

[Git](http://git-scm.com/)

### SVN

[Apache™ Subversion® (SVN)](http://subversion.apache.org/)

### Mercurial

[Mercurial](http://mercurial.selenic.com/)

### Bazaar

[Bazaar](http://bazaar.canonical.com/) — часть проекта GNU, спонсируется
Canonical.

## Управление проектами

### WEB-based

#### Trac

[Trac](http://trac.edgewall.org/)

#### Redmine

[Redmine](http://www.redmine.org/)

#### web2Project

[web2Project](http://web2project.net/)

#### OTRS

[OTRS](http://otrs.org/)

#### Dotproject

[Dotproject](http://dotproject.net/)

### Десктопные

#### Planner

[Planner](http://live.gnome.org/Planner)

#### OpenProj

[OpenProj](http://www.serena.com/products/openproj/)

#### KPlato

[KPlato в KOffice](http://www.koffice.org/kplato/)

#### TaskJuggler

[TaskJuggler](http://www.taskjuggler.org/)

#### GnoTime

[GnoTime](http://gttr.sourceforge.net/)

#### Rachota

[Rachota](http://rachota.sourceforge.net/)

## Компиляторы и системы управления сборкой

### GNU Autotools

[Autotools](http://www.lrde.epita.fr/~adl/autotools.html), или система
сборки GNU,— это набор программных средств, предназначенных для
упрощения переносимости исходного кода программ между
UNIX-подобными системами.

В средства сборки входят: Autoconf, Automake, и Libtool. Другие
средства, не входящие в состав набора, но обычно используемые
совместно с GNU Autotools: make, gettext, pkg-config, и, конечно, GNU
Compiler Collection, также называемый GCC.

### GCC

[GCC, the GNU Compiler Collection](http://gcc.gnu.org/) — что тут можно
ещё сказать... Это GCC.

### CMake

[CMake](http://www.cmake.org/)

### LLVM

[The LLVM Compiler Infrastructure](http://www.llvm.org/) — Clang,
dragonegg & llvm-gcc 4.2, LLDB, libc++ и libc++ ABI, compiler-rt, vmkit,
klee.

### G95

[G95](http://www.g95.org/) — компилятор Fortran 95.

### GNU Fortran (GFortran)

[GNU Fortran (GFortran)](http://gcc.gnu.org/fortran/) — входит в GCC (?)

### Intel® Compilers

[Intel®
Compilers](http://software.intel.com/en-us/articles/intel-compilers/)

### EKOPath 4

[EKOPath 4](http://www.pathscale.com/ekopath-compiler-suite) —
PathScale® EKOPath Compiler Suite недавно открытый компилятор.

## Интерпретаторы

## Отладчики

### GDB

[GDB: The GNU Project Debugger](http://www.gnu.org/software/gdb/)

#### Графические интерфейсы к GDB

##### Nemiver

[Nemiver](http://projects.gnome.org/nemiver/) — GUI к GDB, входящий в
состав GNOME. Подключаемые плагины, удаленная работа через gdbserver,
переключение в режим дизассемблера, подключение к запущенным процессам и
прочее.

##### KDbg

[KDbg](http://www.kdbg.org/) — интегрированный в KDE GUI к GDB.

##### DDD

[DDD](http://www.gnu.org/s/ddd/) — Display Data Debugger. Визуально
отображает ход исполнения другой программы в процессе ее работы;
или визуализирует состояние программы в момент краха.

### IDB

[Intel(R) Debugger for Linux\*
(IDB)](http://software.intel.com/en-us/articles/idb-linux/) — в составе
Intel® Composer XE for Linux and Intel® Parallel Studio XE (Напишите кто
знает, я ими не пользовался)
