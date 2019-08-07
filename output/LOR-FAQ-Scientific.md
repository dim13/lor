## Чем построить график?

  - [gnuplot](http://www.gnuplot.info/), руководства, примеры, оболочки:
      - [gnuplot — not so Frequently Asked
        Questions](http://t16web.lanl.gov/Kawano/gnuplot/index-e.html),
        [зеркало](http://www.ualberta.ca/~xz10/gnuplot/index-e.html)
        если основной сайт не работает.
      - [руководство по рисованию более современных
        графиков](http://www.gnuplotting.org)
      - [ещё одно краткое
        руководство](http://en.wikibooks.org/wiki/Gnuplot)
      - [примеры необычных графиков и
        поверхностей](http://www.phyast.pitt.edu/~zov1/gnuplot/html/gallery.html)
      - [PlotDrop - простой GUI для визуализации серий данных из
        файлов](http://plotdrop.sourceforge.net)
  - [Veusz](http://home.gna.org/veusz/)
  - [GSEGrafix](http://www.gnu.org/software/gsegrafix/)
  - [LabPlot](http://labplot.sourceforge.net/)
  - [QtiPlot](http://soft.proindependent.com/qtiplot.html)
  - [SciDAVis](http://scidavis.sourceforge.net/) - форк QtiPlot
  - [Grace](http://plasma-gate.weizmann.ac.il/Grace/) - с использованием
    библиотеки [Open Motif](http://www.openmotif.org/)
  - [GraceGTK](http://gracegtk.sourceforge.net/) - форк Grace на GTK
  - [QtGrace](http://qtgrace.sourceforge.net/) - форк Grace на Qt
  - [matplotlib](http://matplotlib.sourceforge.net/) -
    Python-библиотека, имеющая сходный с Matlab интерфейс
  - [plplot](http://plplot.sourceforge.net/) - библиотека с биндингами к
    многим языкам.
  - [RLPlot](http://rlplot.sourceforge.net/)
  - [ngraph](http://sourceforge.net/projects/ngraph-gtk/) - построение
    2D графиков. Интерфейс английский, документация только на японском
  - [MathGL](http://mathgl.sourceforge.net/)
  - [SCaVis](http://jwork.org/scavis/)

для трёхмерных данных помимо вышеперечисленного есть:

  - [OpenDX](http://www.opendx.org/)
  - [MayaVi](http://code.enthought.com/projects/mayavi/)
  - [VisIt](https://wci.llnl.gov/codes/visit/)
  - [ParaView](http://www.paraview.org/)

## Как в gnuplot включить поддержку русского

выбрать кодировку символов можно с помощью команды:

    set encoding кодировка

Про русские шрифты написано здесь:
<http://mydebianblog.blogspot.com/2007/11/eps-gnuplot-latex.html>

Про utf8 тут: <http://statist.wald.intevation.org/utf8.html>

utf8 - дополнение

<http://surrender-zen-way.blogspot.com/2010/05/gnuplot-ubuntu-1004-utf8.html>
gnuplot в Ubuntu 10.04 теперь поддерживает UTF8\!

<http://www.gnuplot.info/faq/faq.html> Update: Version 4.4 contains
contains more complete support for UTF-8, including PostScript.

Не забываем про пакет psfrag в LaTeX'е, который позволяет полностью
подменять подписи.

## Как из графика в виде картинки получить табличку с данными?

Вручную:

  - [xyscan](http://star.physics.yale.edu/~ullrich/xyscanDistributionPage/)
  - [Engauge Digitizer](http://digitizer.sourceforge.net/)
  - [unplot.py](http://sovety.blogspot.com/2010/11/unplotpy-from-plots-to-tabular-data.html)
  - [SpectraScan](http://sourceforge.net/projects/spectrascan/)

## А чем можно создать пояснительные рисунки?

  - [CLUCalc](http://www.clucalc.info/) со скриптовым языком CLUScript
    (CLUViz, переработанный и улучшенный CLUCalc, не имеет версии под
    Linux)

## Чем просмотреть и обработать данные с атомно-силового или сканирующего туннельного микроскопа?

  - [Gwyddion](http://gwyddion.net/) поддерживает форматы NTMDT,
    AIST-NT, NanoScanTech, Femtoscan различных версий и многие другие
    (около 100 различных форматов). [Русская
    документация](http://gwyddion.net/documentation/user-guide-ru/)
    по работе с ним.

## Какие есть программы для работы с данными оптической микроскопии?

  - [BioImageXD](http://www.bioimagexd.net/)
  - [Micro-Manager](http://micro-manager.org/)
  - [JMicroVision](http://jmicrovision.com/) - проприетарная программа
    на Java для анализа изображений, заявлена поддержка работы с
    большими изображениями
  - [OME (Open Microscopy Environment)](http://www.openmicroscopy.org/)
    разрабатывает ряд открытых программных продуктов, в том числе
    [OMERO](http://www.openmicroscopy.org/site/products/omero),
    [VisBio](http://www.openmicroscopy.org/site/products/visbio) и
    другие
  - [BisQue](http://bioimage.ucsb.edu/bisque)

Для обработки изображений можно использовать
[ImageJ](http://rsbweb.nih.gov/ij/) написанное на Java.

Можно посмотреть дополнительные варианты [в этом
обзоре](http://www.leica-microsystems.com/science-lab/integrative-open-source-software-for-image-analysis-in-biology/).

## Чем быстро и удобно делать металлографию фотографий микроструктуры материалов с интерактивным исправлением измерений?

Металлография относится к оптической и электронной микроскопии, поэтому
в большинстве случаев будет достаточно программы
[ImageJ](http://rsbweb.nih.gov/ij/). Так, для определения объемных долей
фаз можно использовать плагин Grid или, в случае хорошей контрастности
изображения, использовать анализ частиц.

Однако, ImageJ на данный момент в поставке без плагинов не позволяет
делать неразрушающие измерения изображения с сохранением маркеров.
Эта возможность есть у таких программ как Gwyddion и проприетарная
ImagePro.

К счастью, есть расширение
[ObjectJ](http://simon.bio.uva.nl/objectj/index.html) для ImageJ. Данное
расширение позволяет сразу сохранять измерения отдельных размеров в
отдельных столбцах таблицы, что не позволяет делать ImagePro
(согласно отзывам пользователей этой программы).

Кроме этого поддержка неразрушающих маркеров есть в
[JMicroVision](http://jmicrovision.com), но разделения по типам
измерений, в отличие от ObjectJ, нет.

## Чем можно быстро открыть, исправить, обработать картинки гигантских разрешений?

  - [VIPS](http://www.vips.ecs.soton.ac.uk/index.php?title=VIPS)+[Nip2](http://www.vips.ecs.soton.ac.uk/index.php?title=Nip2)
    (первое - движок, второе - интерфейс)

## Чем можно проанализировать данные с масс-спектрометра/хроматографа?

MS:

  - [massXpert](http://massxpert.org/)
  - [mMass](http://www.mmass.org/)

MS+хроматография:

  - [OpenChrom](http://www.openchrom.net/)

## Какие есть аналоги Mathcad, Matlab, Maple, Mathematica?

Кроме Mathcad, остальные есть под Linux, но, естественно, стоят денег.

Matlab можно попробовать заменить на [Scilab](http://www.scilab.org/)
или [Octave](http://www.gnu.org/software/octave/) вместе с
[QtOctave](http://qtoctave.wordpress.com/) или [Octave
Workshop](http://www.ics.es.yamanashi.ac.jp/mirror/octave-workshop/).
Недостающий функционал в Octave добавляется обычно модулями с
[Octave-Forge](http://octave.sourceforge.net/).

Scilab более функционален, чем Octave, но синтаксис Scilab'a в большей
степени, чем синтаксис Octave, отличается от синтаксиса MatLab, однако
имеется конвертор M2SCI (Matlab2cilab) текстовых скриптов Matlab.

Возможности Scilab можно значительно расширить за счёт [внешних
модулей](http://atoms.scilab.org/), разрабатываемых сообществом и
энтузиастами. Среди них и нужно искать так нехватающие многим в основной
поставке тулбоксы для Wavelet анализа, обработки изображений и пр.

Об Octave можно подробнее почитать собственно в
[документации](http://www.gnu.org/software/octave/doc/interpreter/),
а также [на сайте проекта Xgu.ru](http://xgu.ru/wiki/GNU_Octave).

Еще есть [FreeMat](http://freemat.sourceforge.net/). Синтаксис такой же
как в Matlab, но реализован только базовый функционал, нет большого
числа дополнительных модулей, как в Octave.

Mathcad можно попробовать заменить на [SMath Studio](http://smath.info).

Для символьных вычислений есть [Maxima](http://maxima.sourceforge.net/)
и графические интерфейсы Xmaxima и
[wxMaxima](http://wxmaxima.sourceforge.net/). Примеры работы можно
посмотреть [здесь](http://www.csulb.edu/~woollett/). Также для
символьных вычислений есть [Axiom](http://axiom-developer.org/) (и
его форки [OpenAxiom](http://www.open-axiom.org/) и
[FriCAS](http://fricas.sourceforge.net/)),
[REDUCE](http://www.reduce-algebra.com/) и
[Mathomatic](http://www.mathomatic.org/) Активно развивается CAS,
написанная на Python:[Sympy](http://sympy.org/en/index.html). Ее
можно встраивать в свои приложения.

Для статистического анализа данных есть [R](http://www.r-project.org/) с
графическими интерфейсами [RKWard](http://rkward.sourceforge.net/) и [R
Commander](http://socserv.mcmaster.ca/jfox/Misc/Rcmdr/)

Кроме специализированных программ, достаточно популярны математические
библиотеки для обычных языков программирования, например,
[NumPy](http://numpy.scipy.org/) и [SciPy](http://scipy.org/) для
Python, [GSL](http://www.gnu.org/software/gsl/) для C.

Упомяну комбайн [Sage](http://www.sagemath.org/).

[1](http://en.wikipedia.org/wiki/Comparison_of_computer_algebra_systems)

## А чем можно сразу построить математическую модель задачи и решить получившиеся уравнения?

  - [ASCEND](http://ascend4.org/)

## Какой есть аналог GPSS?

О GPSS все давно забыли. Для моделирования систем массового обслуживания
есть, например, библиотека [SimPy](http://simpy.sourceforge.net/) для
Python.

## Чем можно нарисовать химическую формулу?

  - [BKChem](http://bkchem.zirael.org/)
  - [XDrawChem](http://xdrawchem.sourceforge.net/)
  - [ChemTool](http://ruby.chemie.uni-freiburg.de/~martin/chemtool/)
  - [JChemPaint](http://sourceforge.net/apps/mediawiki/cdk/index.php?title=JChemPaint)
  - [ChemDoodle](http://www.chemdoodle.com/) (несвободный)
  - [MarvinSketch](http://www.chemaxon.com/marvin/sketch/index.php)
    (несвободный)

## Чем нарисовать молекулу для её дальнейшего расчёта?

  - [Ghemical](http://www.uku.fi/~thassine/projects/ghemical/)
  - [Avogadro](http://avogadro.openmolecules.net/)
  - [PyMOL](http://www.pymol.org/)
  - [Gabedit](http://gabedit.sourceforge.net/)
  - [wxmacmolplt](http://www.scl.ameslab.gov/~brett/MacMolPlt/)
  - [CCP1GUI](http://www.cse.scitech.ac.uk/ccg/software/ccp1gui/)

## Чем считать молекулярную динамику?

Молекулярная механика:

  - [Gromacs](http://www.gromacs.org/)
  - [Amber](http://ambermd.org/)
  - [NAMD](http://www.ks.uiuc.edu/Research/namd/)
  - [LAMMPS](http://lammps.sandia.gov/)
  - [TINKER](http://dasher.wustl.edu/tinker/)
  - [HOOMD-blue](http://codeblue.umich.edu/hoomd-blue/)

[и многие
другие](http://en.wikipedia.org/wiki/List_of_software_for_molecular_mechanics_modeling)

Полуэмпирические методы:

  - [MOPAC (Molecular Orbital PACkage)](http://openmopac.net/)
  - [mopac7](http://www.uku.fi/~thassine/projects/download/current/)
    (свободная версия mopac)
  - [AMPAC](http://www.semichem.com/ampac/)

большая часть из нижеперечисленного тоже умеет полуэмпирические методы.

*Ab initio* методы (DFT, Хартри-Фок и т.д.):

  - [Firefly (бывший PC
    GAMESS)](http://classic.chem.msu.su/gran/gamess/index.html)
  - [GAMESS](http://www.msg.ameslab.gov/GAMESS/)
  - [ORCA](http://www.thch.uni-bonn.de/tc/orca/)
  - [Gaussian](http://www.gaussian.com/)
  - [NWChem](http://www.nwchem-sw.org/)
  - [MPQC (Massively Parallel Quantum Chemistry
    Program)](http://www.mpqc.org/)
  - [Molpro](http://www.molpro.net/)

[и многие
другие](http://en.wikipedia.org/wiki/Quantum_chemistry_computer_programs)

## А как её считать?

[Курс](http://kodomo.cmm.msu.ru/wiki/Main/Modelling) молекулярного
моделирования биополимеров Факультета Биоинженерии и
Биоинформатики МГУ.

Учебник по использованию PC GAMESS в связке с wxmacmolplt:
<http://classic.chem.msu.su/gran/gamess/marek/en/docs/PCG-Tutorial-Usage.pdf>

После переименования PC GAMESS в Firefly, на wxmacmolplt надо наложить
патчи для нормальной совместной работы. Взять их можно здесь:
<http://slackbuilds.org/repository/14.0/academic/wxmacmolplt/> Либо
можно использовать опцию -legacy в расчётах firefly, делающую
выходные файлы firefly совместимыми с gamess-us.

## Чем посмотреть результат расчётов молекулярной динамики?

Большая часть программ для рисования позволяет и посмотреть результат,
но есть и специализированные пакеты, предназначенные в основном только
для визуализации:

  - [PyMOL](http://www.pymol.org/)
  - [VMD](http://www.ks.uiuc.edu/Research/vmd/)
  - [Gabedit](http://gabedit.sourceforge.net/)
  - [Chimera](http://www.cgl.ucsf.edu/chimera/)
  - [Molekel](http://www.cscs.ch/molekel/)
  - [OVITO](http://ovito.org/) (поддерживает форматы файлов LAMMPS, XYZ
    и POSCAR)
  - [AtomEye](http://mt.seas.upenn.edu/Archive/Graphics/A/)
    (поддерживается LAMMPS (?))

и многие другие

Колебательные спектры красиво рисует
[GaussSum](http://gausssum.sourceforge.net/).

## Чем перевести из одного химического формата в другой?

  - [Open Babel](http://openbabel.org/)

## Какие есть программы для молекулярной биологии?

  - [UGENE](http://http://ugene.unipro.ru/) — набор инструментов
    молекулярного биолога.

## Какие программы-планетарии есть под linux?

  - [Stellarium](http://www.stellarium.org/)
  - [Celestia](http://www.shatters.net/celestia/)
  - [KStars](http://edu.kde.org/kstars/) из состава проекта KDE Edu
  - [Sky Chart/Cates du Ciel](http://www.ap-i.net/skychart/index.php)

## Какие интерактивные программы демонстрации физических явлений существуют?

  - [PhET](http://phet.colorado.edu/) — написано на Java, требуется
    перевод на русский
  - [Step](http://edu.kde.org/step/) из состава KDE Edu, в основном
    механика и термодинамика

## Какие есть свободные альтернативы пакетам ANSYS, COMSOL Multiphysics, CFD-ACE?

Все эти [CAE](http://ru.wikipedia.org/wiki/Computer-aided_engineering)
пакеты предназначены для серьезных инженерных расчетов и лицензии на
их использование стоят немалых денег. Но и для них есть свободные
альтернативы.

[Salome](http://www.salome-platform.org/home/presentation/overview/) и
набор решателей: OpenFOAM, Code-Aster, Code-Saturne.

[Elmer](http://www.csc.fi/english/pages/elmer), но в качестве генератора
сеток и постпроцессора удобнее использовать сторонние утилиты, например,
[Gmsh](http://geuz.org/gmsh/) и [ParaView](http://www.paraview.org/)
соответственно. Хотя в последнее время разработчики активно
допиливают графический интерфейс (ElmerGUI) и возможно в
ближайшем будущем сторонние утилиты использовать больше не будет
необходимости.

Так же стоит упомянуть специализированный дистрибутив
[CAELinux](http://caelinux.com/CMS/).

## Как рассчитать полупроводниковый лазер с вертикальным резонатором?

  - [CAMFR (CAvity Modelling FRamework)](http://camfr.sourceforge.net/)
    - быстрый, гибкий полностью векторный пакет для решения уравнений
    Масквелла (Maxwell) на C++ с биндингами на Python. Основное
    назначение - область нанофотоники.

<!-- end list -->

  - [Meep](http://ab-initio.mit.edu/wiki/index.php/Meep) Meep (or MEEP)
    is a free finite-difference time-domain (FDTD) simulation software
    package developed at MIT to model electromagnetic systems, along
    with our MPB eigenmode package.

<!-- end list -->

  - [openEMS](http://openems.de/start/index.php) -- пакет для решения
    уравнений Максвелла методом конечных разностей во временной
    области (FTDT). Поддерживает работу в прямоугольной и
    цилиндрической системе координат. Заявлена поддержка SSE и
    MPI (CUDA и OpenCL, к сожалению, не задействуются). Используется
    совместно с matlab или octave.

## Как аппроксимировать экспериментальные пики Гауссом, Лоренцом и др. кривыми?

  - [Fityk](http://www.unipress.waw.pl/fityk/) - приложение для
    нелинейной подгонки кривых методом наименьших квадратов.
  - Команда fit в gnuplot, только нужно сначала задать функцию которой
    аппроксимируется.

## С помощью чего можно организовать или упорядочить коллекцию статей? Какие есть аналоги papers?

  - [Mendeley](http://www.mendeley.com/) — freeware под все три основные
    платформы, распространяется без исходного кода. Данные для
    совместной работы (Shared Collection) хранятся в онлайне
    на серверах mendeley. Бесплатные учетные записи на сервисе
    mendeley предоставляют ограниченный функционал. В 2013 году Mendeley
    была куплена издательской компанией Elsevier.
  - [Zotero](http://www.zotero.org/) — инструмент для управления
    библиографическими данными. Существует в виде плагина для
    браузера, а также отдельного приложения. Интегрируется с
    LibreOffice, Apache OpenOffice и MS Office.
  - [Referencer](http://icculus.org/referencer/) — Удобный каталогизатор
    статей. Позволяет вводить и отслеживать библиографические данные,
    умеет импортировать/экспортировать подборки ссылок в BibTeX
    формат. "Вкусной" фишкой является сканирование текстов статей
    в pdf-формате и извлечение оттуда DOI, по которому онлайн можно
    определить библиографические данные (требуется бесплатная
    регистрация). Благодаря этому можно быстро каталогизировать
    большие объемы текстов статей, сохраненные локально.
  - [KBibTex](http://home.gna.org/kbibtex/) -- KDE-аналог referencer.
    Гораздо более настраиваемый и удобный. Работает напрямую с
    bib-файлом.
  - [Bibus](http://sourceforge.net/projects/bibus-biblio/) —
    библиографическая СУБД. Использует MySQL или SQLite в
    качестве хранилища, может вставлять ссылки в LibreOffice,
    OpenOffice.org и Microsoft Word, а также генерировать
    библиографический список.
  - [Zettelkasten](http://zettelkasten.danielluedecke.de/en/index.php) —
    (с нем. "карточный каталог") кроссплатформенное свободное
    приложение-каталогизатор на Java.
  - [JabRef](http://jabref.sourceforge.net/) — хранитель ссылок по типу
    EndNote на основе BibTeX (Java).
  - [Docear](http://www.docear.org/) — ("Dog-ear") пакет программ для
    работы с научной литературой (работает на большинстве платформ,
    где доступна Java 6). В нем интегрированы функции поиска,
    организации и создания научных текстов. А именно:
    электронная библиотека с поддержкой PDF, менеджер ссылок,
    инструмент для заметок, с mind maps в центральной роли. Более того,
    Docear работает с такими существующими инструментами, как Mendeley,
    Microsoft Word, Foxit Reader. Docear -- бесплатен, с открытым
    исходным кодом. Основан на Freeplane, а также JabRef и JPod.
    Финансируется Федеральным министерством технологии ФРГ.

## Где можно поискать свободно доступную современную литературу?

  - [intechopen](http://www.intechopen.com/books/) естественные и
    технические науки, впрочем, качество там довольно
    сомнительное.

<!-- end list -->

  - [Green Tea Press](http://www.greenteapress.com/) издательство
    объединяет различные книги в сериях Think X, и How to Think
    Like a Computer Scientist,Little Book of Semaphores. Книги
    начального уровня, но написаны специалистами в своей
    области. Доступны их исходные тексты в TeX.

<!-- end list -->

  - [МЦНМО](http://www.mccme.ru/free-books/)Часть книг в области
    математики распостраняется под свободными лицензиями

## Какой софт может предложить OpenSource психологу и нейрофизиологу?

  - [PsychoPy](http://www.psychopy.org/) - программа для проведения
    экспериментов в области восприятия и дальнейшей обработки
    результатов, написанная, как это ни странно, на Python.
  - [NeuroDebian](http://neuro.debian.net/) Проект развивающий
    репозиторий нейрофизиологического и психологического
    свободного софта. Так же они делают сборки Debian с
    предустановленным софтом
  - [OpenEEG](http://openeeg.sourceforge.net/doc/) Это проект по
    созданию дешевых приборов ЭЭГ диагностики и свободного софта
    для этой же цели
  - [EEGLAB](http://sccn.ucsd.edu/eeglab/index.html/) Инструмент для
    анализа ЭЭГ,являющийся расширением для Matlab
  - [ODIN](http://od1n.sourceforge.net/) объектно-ориентированная
    библиотека на С++ для анализа данных МРТ

## Обработка сигналов

  - [библиотека scipy](http://www.scipy.org/) -- данная библиотека на
    языке Python имеет несколько модулей, представляющих различные
    алгоритмы обработки сигналов.

Например scipy.signals реализует базовые операции над
сигналом(нормирование,поиск пиков и т.д.). Так же
данная библиотека поддерживает возможности по созданию цифровых
фильтров, расчету и моделированию аналоговых, расчету спектра (на
основе БПФ, метод Велча, модифицированный метод Велча,с применением
оконных функций и др.),расчету адаптивных фильтров. Многие функции
довольно близки по синтаксису к матлаб.

  - [GNU Radio](http://gnuradio.org/redmine/projects/gnuradio/wiki)
    Данный фреймворк предоставляет инструменты для анализа сигналов

## Что можно использовать для визуализации метаболических путей?

  - [Arcadia](http://arcadiapathways.sourceforge.net/) — просмотрщик,
    средство визуализации метаболических путей с поддержкой
    стандартов [SBML](http://www.sbml.org) и
    [SBGN](http://www.sbgn.org) с интерфейсом на Qt.

Прочее ПО для этой задачи можно найти на сайтах
[SBML](http://www.sbml.org) и [SBGN](http://www.sbgn.org).

[Category:Полезные советы](Category:Полезные_советы)