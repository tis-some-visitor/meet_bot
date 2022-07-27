package maps

var RussianTowns = map[string]int{"агидель": 0, "баймак": 1, "белебей": 2, "белорецк": 3, "бирск": 4, "благовещенск": 5, "давлеканово": 6, "дюртюли": 7, "ишимбай": 8, "кумертау": 9, "межгорье": 10, "мелеуз": 11, "нефтекамск": 12, "октябрьский": 13, "салават": 14, "сибай": 15, "стерлитамак": 16, "туймазы": 17, "уфа": 18, "учалы": 19, "янаул": 20, "бабушкин": 21, "гусиноозерск": 22, "закаменск": 23, "кяхта": 24, "северобайкальск": 25, "улан-удэ": 26, "горно-алтайск": 27, "буйнакск": 28, "дагестанские огни": 29, "дербент": 30, "избербаш": 31, "каспийск": 32, "кизилюрт": 33, "кизляр": 34, "махачкала": 35, "хасавюрт": 36, "южно-сухокумск": 37, "карабулак": 38, "магас": 39, "малгобек": 40, "назрань": 41, "баксан": 42, "майский": 43, "нальчик": 44, "нарткала": 45, "прохладный": 46, "терек": 47, "тырныауз": 48, "чегем": 49, "городовиковск": 50, "лагань": 51, "элиста": 52, "карачаевск": 53, "теберда": 54, "усть-джегута": 55, "черкесск": 56, "беломорск": 57, "кемь": 58, "кондопога": 59, "костомукша": 60, "лахденпохья": 61, "медвежьегорск": 62, "олонец": 63, "петрозаводск": 64, "питкяранта": 65, "пудож": 66, "сегежа": 67, "сортавала": 68, "суоярви": 69, "воркута": 70, "вуктыл": 71, "емва": 72, "инта": 73, "микунь": 74, "печора": 75, "сосногорск": 76, "сыктывкар": 77, "усинск": 78, "ухта": 79, "волжск": 80, "звенигово": 81, "йошкар-ола": 82, "козьмодемьянск": 83, "ардатов": 84, "инсар": 85, "ковылкино": 86, "краснослободск": 87, "рузаевка": 88, "саранск": 89, "темников": 90, "алдан": 91, "верхоянск": 92, "вилюйск": 93, "ленск": 94, "мирный": 95, "нерюнгри": 96, "нюрба": 97, "олекминск": 98, "покровск": 99, "среднеколымск": 100, "томмот": 101, "удачный": 102, "якутск": 103, "алагир": 104, "ардон": 105, "беслан": 106, "владикавказ": 107, "дигора": 108, "моздок": 109, "агрыз": 110, "азнакаево": 111, "альметьевск": 112, "арск": 113, "бавлы": 114, "болгар": 115, "бугульма": 116, "буинск": 117, "елабуга": 118, "заинск": 119, "зеленодольск": 120, "казань": 121, "лаишево": 122, "лениногорск": 123, "мамадыш": 124, "менделеевск": 125, "мензелинск": 126, "набережные челны": 127, "нижнекамск": 128, "нурлат": 129, "тетюши": 130, "чистополь": 131, "ак-довурак": 132, "кызыл": 133, "туран": 134, "чадан": 135, "шагонар": 136, "воткинск": 137, "глазов": 138, "ижевск": 139, "камбарка": 140, "можга": 141, "сарапул": 142, "абаза": 143, "абакан": 144, "саяногорск": 145, "сорск": 146, "черногорск": 147, "аргун": 148, "грозный": 149, "гудермес": 150, "урус-мартан": 151, "шали": 152, "алатырь": 153, "канаш": 154, "козловка": 155, "мариинский посад": 156, "новочебоксарск": 157, "цивильск": 158, "чебоксары": 159, "шумерля": 160, "ядрин": 161, "алейск": 162, "барнаул": 163, "белокуриха": 164, "бийск": 165, "горняк": 166, "заринск": 167, "змеиногорск": 168, "камень-на-оби": 169, "новоалтайск": 170, "рубцовск": 171, "славгород": 172, "яровое": 173, "абинск": 174, "анапа": 175, "апшеронск": 176, "армавир": 177, "белореченск": 178, "геленджик": 179, "горячий ключ": 180, "гулькевичи": 181, "ейск": 182, "кореновск": 183, "краснодар": 184, "кропоткин": 185, "крымск": 186, "курганинск": 187, "лабинск": 188, "новокубанск": 189, "новороссийск": 190, "приморско-ахтарск": 191, "славянск-на-кубани": 192, "сочи": 193, "темрюк": 194, "тимашевск": 195, "тихорецк": 196, "туапсе": 197, "усть-лабинск": 198, "хадыженск": 199, "артемовск": 200, "ачинск": 201, "боготол": 202, "бородино": 203, "дивногорск": 204, "дудинка": 205, "енисейск": 206, "железногорск": 207, "заозерный": 208, "зеленогорск": 209, "игарка": 210, "иланский": 211, "канск": 212, "кодинск": 213, "красноярск": 214, "лесосибирск": 215, "минусинск": 216, "назарово": 217, "норильск": 218, "сосновоборск": 219, "ужур": 220, "уяр": 221, "шарыпово": 222, "арсеньев": 223, "артем": 224, "большой камень": 225, "владивосток": 226, "дальнегорск": 227, "дальнереченск": 228, "лесозаводск": 229, "находка": 230, "партизанск": 231, "спасск-дальний": 232, "уссурийск": 233, "фокино": 234, "благодарный": 235, "буденновск": 236, "георгиевск": 237, "ессентуки": 238, "железноводск": 239, "зеленокумск": 240, "изобильный": 241, "ипатово": 242, "кисловодск": 243, "лермонтов": 244, "минеральные воды": 245, "михайловск": 246, "невинномысск": 247, "нефтекумск": 248, "новоалександровск": 249, "новопавловск": 250, "пятигорск": 251, "светлоград": 252, "ставрополь": 253, "амурск": 254, "бикин": 255, "вяземский": 256, "комсомольск-на-амуре": 257, "николаевск-на-амуре": 258, "советская гавань": 259, "хабаровск": 260, "белогорск": 261, "завитинск": 263, "зея": 264, "райчихинск": 265, "свободный": 266, "сковородино": 267, "тында": 268, "шимановск": 269, "архангельск": 270, "вельск": 271, "каргополь": 272, "коряжма": 273, "котлас": 274, "мезень": 275, "новодвинск": 277, "няндома": 278, "онега": 279, "северодвинск": 280, "сольвычегодск": 281, "шенкурск": 282, "астрахань": 283, "ахтубинск": 284, "ахтубинск-7": 285, "знаменск": 286, "камызяк": 287, "нариманов": 288, "харабали": 289, "алексеевка": 290, "белгород": 291, "бирюч": 292, "валуйки": 293, "грайворон": 294, "губкин": 295, "короча": 296, "новый оскол": 297, "старый оскол": 298, "строитель": 299, "шебекино": 300, "брянск": 301, "дятьково": 302, "жуковка": 303, "злынка": 304, "карачев": 305, "клинцы": 306, "мглин": 307, "новозыбков": 308, "почеп": 309, "севск": 310, "сельцо": 311, "стародуб": 312, "сураж": 313, "трубчевск": 314, "унеча": 315, "александров": 317, "владимир": 318, "вязники": 319, "гороховец": 320, "гусь-хрустальный": 321, "камешково": 322, "карабаново": 323, "киржач": 324, "ковров": 325, "кольчугино": 326, "костерево": 327, "курлово": 328, "лакинск": 329, "меленки": 330, "муром": 331, "петушки": 332, "покров": 333, "радужный": 334, "собинка": 335, "струнино": 336, "судогда": 337, "суздаль": 338, "юрьев-польский": 339, "волгоград": 340, "волжский": 341, "дубовка": 342, "жирновск": 343, "калач-на-дону": 344, "камышин": 345, "котельниково": 346, "котово": 347, "ленинск": 349, "михайловка": 350, "николаевск": 351, "новоаннинский": 352, "палласовка": 353, "петров вал": 354, "серафимович": 355, "суровикино": 356, "урюпинск": 357, "фролово": 358, "бабаево": 359, "белозерск": 360, "великий устюг": 361, "вологда": 362, "вытегра": 363, "грязовец": 364, "кадников": 365, "кириллов": 366, "красавино": 367, "никольск": 368, "сокол": 369, "тотьма": 370, "устюжна": 371, "харовск": 372, "череповец": 373, "бобров": 374, "богучар": 375, "борисоглебск": 376, "бутурлиновка": 377, "воронеж": 378, "воронеж-45": 379, "калач": 380, "лиски": 381, "нововоронеж": 382, "новохоперск": 383, "острогожск": 384, "павловск": 385, "поворино": 386, "россошь": 387, "семилуки": 388, "эртиль": 389, "вичуга": 390, "гаврилов-посад": 391, "заволжск": 392, "иваново": 393, "кинешма": 394, "комсомольск": 395, "кохма": 396, "наволоки": 397, "плес": 398, "приволжск": 399, "пучеж": 400, "родники": 401, "тейково": 402, "фурманов": 403, "шуя": 404, "южа": 405, "юрьевец": 406, "алзамай": 407, "ангарск": 408, "байкальск": 409, "бирюсинск": 410, "бодайбо": 411, "братск": 412, "вихоревка": 413, "железногорск-илимский": 414, "зима": 415, "иркутск": 416, "иркутск-45": 417, "киренск": 418, "нижнеудинск": 419, "саянск": 420, "свирск": 421, "слюдянка": 422, "тайшет": 423, "тулун": 424, "усолье-сибирское": 425, "усть-илимск": 426, "усть-кут": 427, "черемхово": 428, "шелехов": 429, "багратионовск": 430, "балтийск": 431, "гвардейск": 432, "гурьевск": 433, "гусев": 434, "зеленоградск": 435, "калининград": 436, "краснознаменск": 437, "ладушкин": 438, "мамоново": 439, "неман": 440, "нестеров": 441, "озерск": 442, "пионерский": 443, "полесск": 444, "правдинск": 445, "приморск": 446, "светлогорск": 447, "светлый": 448, "славск": 449, "советск": 450, "черняховск": 451, "балабаново": 452, "белоусово": 453, "боровск": 454, "боровск-1": 455, "ермолино": 456, "жиздра": 457, "жуков": 458, "калуга": 459, "киров": 460, "козельск": 461, "кондрово": 462, "кременки": 463, "людиново": 464, "малоярославец": 465, "медынь": 466, "мещовск": 467, "мосальск": 468, "обнинск": 469, "сосенский": 470, "спас-деменск": 471, "сухиничи": 472, "таруса": 473, "юхнов": 474, "юхнов-1": 475, "юхнов-2": 476, "вилючинск": 477, "елизово": 478, "петропавловск-камчатский": 479, "анжеро-судженск": 480, "белово": 481, "березовский": 482, "калтан": 484, "кемерово": 485, "киселевск": 486, "ленинск-кузнецкий": 487, "мариинск": 488, "междуреченск": 489, "мыски": 490, "новокузнецк": 491, "осинники": 492, "полысаево": 493, "прокопьевск": 494, "салаир": 495, "тайга": 496, "таштагол": 497, "топки": 498, "юрга": 499, "белая холуница": 500, "вятские поляны": 501, "зуевка": 502, "кирово-чепецк": 504, "кирс": 505, "котельнич": 506, "луза": 507, "малмыж": 508, "мураши": 509, "нолинск": 510, "омутнинск": 511, "орлов": 512, "слободской": 513, "сосновка": 515, "уржум": 516, "яранск": 517, "буй": 518, "волгореченск": 519, "галич": 520, "кологрив": 521, "кострома": 522, "макарьев": 523, "мантурово": 524, "нерехта": 525, "нея": 526, "солигалич": 527, "чухлома": 528, "шарья": 529, "далматово": 530, "катайск": 531, "курган": 532, "куртамыш": 533, "макушино": 534, "петухово": 535, "шадринск": 536, "шумиха": 537, "щучье": 538, "дмитриев": 539, "курск": 541, "курчатов": 542, "льгов": 543, "обоянь": 544, "рыльск": 545, "суджа": 546, "фатеж": 547, "щигры": 548, "бокситогорск": 549, "волосово": 550, "волхов": 551, "всеволожск": 552, "выборг": 553, "высоцк": 554, "гатчина": 555, "ивангород": 556, "каменногорск": 557, "кингисепп": 558, "кириши": 559, "кировск": 560, "коммунар": 561, "лодейное поле": 562, "луга": 563, "любань": 564, "никольское": 565, "новая ладога": 566, "отрадное": 567, "пикалево": 568, "подпорожье": 569, "приозерск": 571, "светогорск": 572, "сертолово": 573, "сланцы": 574, "сосновый бор": 575, "сясьстрой": 576, "тихвин": 577, "тосно": 578, "шлиссельбург": 579, "грязи": 580, "данков": 581, "елец": 582, "задонск": 583, "лебедянь": 584,
	"липецк": 585, "усмань": 586, "чаплыгин": 587, "магадан": 588, "сусуман": 589, "апрелевка": 590, "балашиха": 591,
	"бронницы": 592, "верея": 593, "видное": 594, "волоколамск": 595, "воскресенск": 596, "высоковск": 597, "голицыно": 598, "городской округ черноголовка": 599, "дедовск": 600, "дзержинский": 601, "дмитров": 602, "долгопрудный": 603, "домодедово": 604, "дрезна": 605, "дубна": 606, "егорьевск": 607, "железнодорожный": 608, "жуковский": 609, "зарайск": 610, "звенигород": 611, "ивантеевка": 612, "истра": 613, "истра-1": 614, "кашира": 615, "кашира-8": 616, "климовск": 617, "клин": 618, "коломна": 619, "королев": 620, "котельники": 621, "красноармейск": 622, "красногорск": 623, "краснозаводск": 624, "кубинка": 626, "куровское": 627, "ликино-дулево": 628, "лобня": 629, "лосино-петровский": 630, "луховицы": 631, "лыткарино": 632, "люберцы": 633, "можайск": 634, "мытищи": 635, "наро-фоминск": 636, "ногинск": 637, "одинцово": 638, "ожерелье": 639, "озеры": 640, "орехово-зуево": 641, "павловский посад": 642, "пересвет": 643, "подольск": 644, "протвино": 645, "пушкино": 646, "пущино": 647, "раменское": 648, "реутов": 649, "рошаль": 650, "руза": 651, "сергиев посад": 652, "сергиев посад-7": 653, "серпухов": 654, "снегири": 655, "солнечногорск": 656, "солнечногорск-2": 657, "солнечногорск-25": 658, "солнечногорск-30": 659, "солнечногорск-7": 660, "старая купавна": 661, "ступино": 662, "талдом": 663, "фрязино": 664, "химки": 665, "хотьково": 666, "черноголовка": 667, "чехов": 668, "чехов-2": 669, "чехов-3": 670, "чехов-8": 671, "шатура": 672, "щелково": 673, "электрогорск": 674, "электросталь": 675, "электроугли": 676, "юбилейный": 677, "яхрома": 678, "апатиты": 679, "гаджиево": 680, "заозерск": 681, "заполярный": 682, "кандалакша": 683, "ковдор": 685, "кола": 686, "мончегорск": 687, "мурманск": 688, "оленегорск": 689, "оленегорск-1": 690, "оленегорск-2": 691, "оленегорск-4": 692, "островной": 693, "полярные зори": 694, "полярный": 695, "североморск": 696, "снежногорск": 697, "арзамас": 698, "балахна": 699, "богородск": 700, "бор": 701, "ветлуга": 702, "володарск": 703, "ворсма": 704, "выкса": 705, "горбатов": 706, "городец": 707, "дзержинск": 708, "заволжье": 709, "княгинино": 710, "кстово": 711, "кулебаки": 712, "лукоянов": 713, "лысково": 714, "навашино": 715, "нижний новгород": 716, "павлово": 717, "первомайск": 718, "перевоз": 719, "саров": 720, "семенов": 721, "сергач": 722, "урень": 723, "чкаловск": 724, "шахунья": 725, "боровичи": 726, "валдай": 727, "великий новгород": 728, "малая вишера": 729, "окуловка": 730, "пестово": 731, "сольцы": 732, "сольцы 2": 733, "старая русса": 734, "холм": 735, "чудово": 736, "барабинск": 737, "бердск": 738, "болотное": 739, "искитим": 740, "карасук": 741, "каргат": 742, "куйбышев": 743, "купино": 744, "новосибирск": 745, "обь": 746, "татарск": 747, "тогучин": 748, "черепаново": 749, "чулым": 750, "чулым-3": 751, "исилькуль": 752, "калачинск": 753, "называевск": 754, "омск": 755, "тара": 756, "тюкалинск": 757, "абдулино": 758, "бугуруслан": 759, "бузулук": 760, "гай": 761, "кувандык": 762, "медногорск": 763, "новотроицк": 764, "оренбург": 765, "орск": 766, "соль-илецк": 767, "сорочинск": 768, "ясный": 769, "болхов": 770, "дмитровск": 771, "ливны": 772, "малоархангельск": 773, "мценск": 774, "новосиль": 775, "орел": 776, "белинский": 777, "городище": 778, "заречный": 779, "каменка": 780, "кузнецк": 781, "кузнецк-12": 782, "кузнецк-8": 783, "нижний ломов": 784, "пенза": 786, "сердобск": 787, "спасск": 788, "сурск": 789, "александровск": 790, "березники": 791, "верещагино": 792, "горнозаводск": 793, "гремячинск": 794, "губаха": 795, "добрянка": 796, "кизел": 797, "красновишерск": 798, "краснокамск": 799, "кудымкар": 800, "кунгур": 801, "лысьва": 802, "нытва": 803, "оса": 804, "оханск": 805, "очер": 806, "пермь": 807, "соликамск": 808, "усолье": 809, "чайковский": 810, "чердынь": 811, "чермоз": 812, "чернушка": 813, "чусовой": 814, "великие луки": 815, "великие луки-1": 816, "гдов": 817, "дно": 818, "невель": 819, "новоржев": 820, "новосокольники": 821, "опочка": 822, "остров": 823, "печоры": 824, "порхов": 825, "псков": 826, "пустошка": 827, "пыталово": 828, "себеж": 829, "азов": 830, "аксай": 831, "батайск": 832, "белая калитва": 833, "волгодонск": 834, "гуково": 835, "донецк": 836, "зверево": 837, "зерноград": 838, "каменск-шахтинский": 839, "константиновск": 840, "красный сулин": 841, "миллерово": 842, "морозовск": 843, "новочеркасск": 844, "новошахтинск": 845, "пролетарск": 846, "ростов-на-дону": 847, "сальск": 848, "семикаракорск": 849, "таганрог": 850, "цимлянск": 851, "шахты": 852, "касимов": 853, "кораблино": 854, "михайлов": 855, "новомичуринск": 856, "рыбное": 857, "ряжск": 858, "рязань": 859, "сасово": 860, "скопин": 861, "спас-клепики": 862, "спасск-рязанский": 863, "шацк": 864, "жигулевск": 865, "кинель": 866, "нефтегорск": 867, "новокуйбышевск": 868, "октябрьск": 869, "отрадный": 870, "похвистнево": 871, "самара": 872, "сызрань": 873, "тольятти": 874, "чапаевск": 875, "аркадак": 876, "аткарск": 877, "балаково": 878, "балашов": 879, "вольск": 880, "вольск-18": 881, "ершов": 882, "калининск": 883, "красный кут": 885, "маркс": 886, "новоузенск": 887, "петровск": 888, "пугачев": 889, "ртищево": 890, "саратов": 891, "хвалынск": 892, "шиханы": 893, "энгельс": 894, "энгельс-19": 895, "энгельс-2": 896, "александровск-сахалинский": 897, "анива": 898, "долинск": 899, "корсаков": 900, "курильск": 901, "макаров": 902, "невельск": 903, "оха": 904, "поронайск": 905, "северо-курильск": 906, "томари": 907, "углегорск": 908, "холмск": 909, "шахтерск": 910, "южно-сахалинск": 911, "алапаевск": 912, "арамиль": 913, "артемовский": 914, "асбест": 915, "богданович": 917, "верхний тагил": 918, "верхняя пышма": 919, "верхняя салда": 920, "верхняя тура": 921, "верхотурье": 922, "волчанск": 923, "дегтярск": 924, "екатеринбург": 925, "ивдель": 927, "ирбит": 928, "каменск-уральский": 929, "камышлов": 930, "карпинск": 931, "качканар": 932, "кировград": 933, "краснотурьинск": 934, "красноуральск": 935, "красноуфимск": 936, "кушва": 937, "лесной": 938, "невьянск": 940, "нижние серги": 941, "нижние серги-3": 942, "нижний тагил": 943, "нижняя салда": 944, "нижняя тура": 945, "новая ляля": 946, "новоуральск": 947, "первоуральск": 948, "полевской": 949, "ревда": 950, "реж": 951, "североуральск": 952, "серов": 953, "среднеуральск": 954, "сухой": 955, "сысерть": 956, "тавда": 957, "талица": 958, "туринск": 959, "велиж": 960, "вязьма": 961, "гагарин": 962, "демидов": 963, "десногорск": 964, "дорогобуж": 965, "духовщина": 966, "ельня": 967, "починок": 968, "рославль": 969, "рудня": 970, "сафоново": 971, "смоленск": 972, "сычевка": 973, "ярцево": 974, "жердевка": 975, "кирсанов": 976, "котовск": 977, "мичуринск": 978, "моршанск": 979, "рассказово": 980, "тамбов": 981, "уварово": 982, "андреаполь": 983, "бежецк": 984, "белый": 985, "бологое": 986, "весьегонск": 987, "вышний волочек": 988, "западная двина": 989, "зубцов": 990, "калязин": 991, "кашин": 992, "кимры": 993, "конаково": 994, "красный холм": 995, "кувшиново": 996, "лихославль": 997, "нелидово": 998, "осташков": 999, "ржев": 1000, "старица": 1001, "тверь": 1002, "торжок": 1003, "торопец": 1004, "удомля": 1005, "асино": 1006, "кедровый": 1007, "колпашево": 1008, "северск": 1009, "стрежевой": 1010, "томск": 1011, "алексин": 1012, "белев": 1013, "богородицк": 1014, "болохово": 1015, "венев": 1016, "донской": 1017, "ефремов": 1018, "кимовск": 1019, "киреевск": 1020, "липки": 1021, "новомосковск": 1022, "плавск": 1023, "суворов": 1025, "тула": 1026, "узловая": 1027, "чекалин": 1028, "щекино": 1029, "ясногорск": 1030, "заводоуковск": 1031, "ишим": 1032, "тобольск": 1033, "тюмень": 1034, "ялуторовск": 1035, "барыш": 1036, "димитровград": 1037, "инза": 1038, "новоульяновск": 1039, "сенгилей": 1040, "ульяновск": 1041, "аша": 1042, "бакал": 1043, "верхнеуральск": 1044, "верхний уфалей": 1045, "еманжелинск": 1046, "златоуст": 1047, "карабаш": 1048, "карталы": 1049, "касли": 1050, "катав-ивановск": 1051, "копейск": 1052, "коркино": 1053, "куса": 1054, "кыштым": 1055, "магнитогорск": 1056, "миасс": 1057, "миньяр": 1058, "нязепетровск": 1059, "пласт": 1061, "сатка": 1062, "сим": 1063, "снежинск": 1064, "трехгорный": 1065, "трехгорный-1": 1066, "троицк": 1067, "усть-катав": 1068, "чебаркуль": 1069, "челябинск": 1070, "южноуральск": 1071, "юрюзань": 1072, "балей": 1073, "борзя": 1074, "краснокаменск": 1075, "могоча": 1076, "нерчинск": 1077, "петровск-забайкальский": 1078, "сретенск": 1079, "хилок": 1080, "чита": 1081, "шилка": 1082, "гаврилов-ям": 1083, "данилов": 1084, "любим": 1085, "мышкин": 1086, "переславль-залесский": 1087, "пошехонье": 1088, "ростов": 1089, "рыбинск": 1090, "тутаев": 1091, "углич": 1092, "ярославль": 1093, "зеленоград": 1094, "москва": 1095, "московский": 1096, "щербинка": 1099, "колпино": 1101, "красное село": 1102, "кронштадт": 1103, "ломоносов": 1104, "петергоф": 1106, "пушкин": 1107, "санкт-петербург": 1108, "сестрорецк": 1109, "биробиджан": 1110, "облучье": 1111, "нарьян-мар": 1112, "белоярский": 1113, "когалым": 1114, "лангепас": 1115, "лянтор": 1116, "мегион": 1117, "нефтеюганск": 1118, "нижневартовск": 1119, "нягань": 1120, "покачи": 1121, "пыть-ях": 1122, "советский": 1124, "сургут": 1125, "урай": 1126, "ханты-мансийск": 1127, "югорск": 1128, "анадырь": 1129, "билибино": 1130, "певек": 1131, "губкинский": 1132, "лабытнанги": 1133, "муравленко": 1134, "надым": 1135, "новый уренгой": 1136, "ноябрьск": 1137, "салехард": 1138, "тарко-сале": 1139,
	"алупка": 1140, "алушта": 1141, "армянск": 1142, "армянськ": 1143, "бахчисарай": 1144, "джанкой": 1146, "евпатория": 1148, "керчь": 1149, "красноперекопск": 1150, "подгорное": 1152, "саки": 1153, "симферополь": 1155, "старый крым": 1156, "судак": 1157, "феодосия": 1158, "щелкино": 1159, "ялта": 1160, "инкерман": 1161, "севастополь": 1162,
	"республика адыгея": 0, "адыгея": 0, "республика алтай": 1, "алтай": 1, "республика башкортостан": 2, "башкортостан": 2,
	"республика бурятия": 3, "бурятия": 3, "республика дагестан": 4, "дагестан": 4, "республика ингушетия": 5, "ингушетия": 5,
	"кабардино-балкарская республика": 6, "республика калмыкия": 7, "карачаево-черкесская республика": 8,
	"кабардино-балкария": 6, "калмыкия": 7, "9. карачаево-черкессия": 8,
	"республика карелия": 9, "республика коми": 10, "республика крым": 11, "республика марий эл": 12,
	"карелия": 9, "коми": 10, "крым": 11, "марий эл": 12,
	"республика мордовия": 13, "республика саха": 14, "якутия": 14, "республика северная осетия": 15, "алания": 15,
	"мордовия": 13, "саха": 14, "северная осетия": 15,
	"республика татарстан": 16, "республика тыва": 17, "удмуртская республика": 18, "республика хакасия": 19,
	"татарстан": 16, "тыва": 17, "удмуртия": 18, "хакасия": 19,
	"чеченская республика": 20, "чувашская республика": 21, "алтайский край": 22, "забайкальский край": 23,
	"чечня": 20, "чувашия": 21, "забайкалье": 23,
	"камчатский край": 24, "краснодарский край": 25, "красноярский  край": 26, "пермский край": 27, "приморский край": 28,
	"камчатка": 24, "приморье": 28, "ставрополье": 29,
	"ставропольский край": 29, "хабаровский край": 30, "амурская область": 31, "архангельская область": 32,
	"астраханская область": 33, "белгородская область": 34, "брянская область": 35, "владимирская область": 36,
	"волгоградская область": 37, "вологодская область": 38, "воронежская область": 39, "ивановская область": 40,
	"иркутская область": 41, "калининградская область": 42, "калужская область": 43, "кемеровская область": 44,
	"кировская область": 45, "костромская область": 46, "курганская область": 47, "курская область": 48,
	"ленинградская область": 49, "ленобласть": 49, "липецкая область": 50, "магаданская область": 51,
	"московская область": 52, "мурманская область": 53, "нижегородская область": 54, "новгородская область": 55,
	"новосибирская область": 56, "омская область": 57, "оренбургская область": 58, "орловская область": 59,
	"пензенская область": 60, "псковская область": 61, "ростовская область": 62, "рязанская область": 63,
	"самарская область": 64, "саратовская область": 65, "сахалинская область": 66, "сахалин": 66,
	"свердловская область": 67, "смоленская область": 68, "тамбовская область": 69, "тверская область": 70,
	"томская область": 71, "тульская область": 72, "тюменская область": 73, "ульяновская область": 74,
	"челябинская область": 75, "ярославская область": 76, "еврейская автономная область": 80, "еврейская область": 80,
	"ненецкий автономный округ": 81, "ханты-мансийский автономный округ": 82, "чукотский автономный округ": 83,
	"ненецкий округ": 81, "ханты-мансийский округ": 82, "чукотка": 83, "ямало-ненецкий автономный округ": 84,
	"ямало-ненецкий округ": 84, "урал": 92, "кавказ": 3, "байкал": 80, "золотое кольцо": 3, "кижи": 80, "соловки": 80,
	"соловецкие острова": 80}

var Countries = map[string]int{"австрия": 2, "мадагаскар": 3, "мексика": 4, "суринам": 5, "лесото": 6, "мьянма": 7,
	"антарктика":         8,
	"виргинские острова": 9, "зимбабве": 10, "демократическая республика конго": 11, "молдавия": 12, "сейшелы": 13,
	"словакия": 14, "ямайка": 15, "туркмения": 16, "цар": 17, "ангола": 18, "белорусь": 19, "конго": 20,
	"народная республика конго": 20, "республика конго": 20, "малайзия": 21,
	"мали": 22, "намибия": 23, "швейцария": 24, "бангладеш": 25, "пакистан": 26, "саудовская аравия": 27,
	"коморские острова": 28, "коморы": 28, "афганистан": 29, "мальдивы": 30, "микронезия": 31, "австралия": 32, "андорра": 33,
	"боливия":  34,
	"иордания": 35, "колумбия": 36, "ангилья": 37, "китай": 38, "кнр": 38, "ирландия": 39, "северная ирландия": 39,
	"ливан": 40, "палау": 41,
	"юар": 42, "монако": 43, "сингапур": 44, "швеция": 45, "италия": 46, "судан": 47, "перу": 48, "бразилия": 49,
	"сальвадор": 50, "северная македония": 51, "сенегал": 52, "сирия": 53, "франция": 54, "вануату": 55, "монголия": 56,
	"сан-томе и принсипи": 57, "нигерия": 58, "узбекистан": 59, "болгария": 60, "сша": 221, "тувалу": 63,
	"туркменистан": 64, "эстония": 65, "гайана": 66, "канада": 67, "алжир": 68, "бахрейн": 69, "грузия": 70, "иран": 71,
	"корейская народно-демократическая республика": 164, "кндр": 164, "англия": 96, "исландия": 74, "науру": 75, "фиджи": 76,
	"экваториальная гвинея": 77, "ботсвана": 78, "венгрия": 79, "гренада": 80, "парагвай": 81, "пуэрто-рико": 82,
	"габон": 83, "канарские острова": 178, "канары": 178, "буркина-фасо": 85, "куба": 86, "оман": 87, "япония": 88,
	"кирибати":   89,
	"коста-рика": 91, "арктика": 92, "гваделупа": 93, "кот-д'ивуар": 94, "бруней": 95, "великобритания": 96,
	"восточный тимор": 97, "мальта": 98, "панама": 99, "словения": 100, "бутан": 101, "вьетнам": 102, "сербия": 103,
	"хорватия": 104, "венесуэла": 105, "румыния": 107, "эмираты": 152, "оаэ": 152, "гаити": 109, "гондурас": 110,
	"сент-винсент и гренадины": 111, "центральноафриканская республика": 17, "казахстан": 113, "гвинея": 114,
	"сан-марино": 115, "руанда": 116, "уругвай": 117, "йемен": 118, "тибет": 119, "теркс и кайкос": 120, "гана": 121,
	"литва": 122, "нидерланды": 123, "сейшельские острова": 13, "шри-ланка": 125, "британия": 96, "гватемала": 127,
	"индонезия": 128, "кипр": 129, "эритрея": 130, "израиль": 131, "маршалловы острова": 132, "тонга": 133,
	"черногория": 134, "польша": 135, "джибути": 136, "кабо-верде": 137, "гвинея-бисау": 138, "соломоновы острова": 139,
	"финляндия": 140, "дания": 141, "камерун": 142, "непал": 143, "того": 144, "украина": 145, "азербайджан": 146,
	"доминикана": 174, "чили": 148, "албания": 149, "египет": 150, "нигер": 151, "объединенные арабские эмираты": 152,
	"чехия": 153, "антигуа и барбуда": 154, "замбия": 155, "сомали": 156, "танзания": 157, "ливия": 158, "лихтенштейн": 159,
	"маврикий": 160, "мозамбик": 161, "аргентина": 162, "испания": 163, "марокко": 165, "сент-китс и невис": 166,
	"южный судан": 167, "латвия": 169, "молдова": 12, "новая зеландия": 171, "новая гвинея": 172,
	"доминиканская республика": 174, "гренландия": 175, "южная корея": 176, "тунис": 177,
	"бурунди": 179, "лаос": 180, "папуа - новая гвинея": 172, "папуа-новая гвинея": 172, "сент-люсия": 183, "мавритания": 184,
	"никарагуа": 185, "тринидад и тобаго": 186, "турция": 187, "филиппины": 188, "ирак": 189, "люксембург": 190,
	"бенин": 191, "македония": 51, "германия": 193, "норвегия": 194, "белоруссия": 19, "кувейт": 196, "сьерра-леоне": 197,
	"мартиника": 198, "греция": 199, "камбоджа": 200, "малави": 201, "эсватини": 202, "монсеррат": 203,
	"босния и герцеговина": 204, "гамбия": 205, "катар": 206, "таиланд": 207, "киргизия": 208, "чад": 209, "бельгия": 210,
	"кения": 211, "россия": 212, "шотландия": 213, "уэльс": 214, "таджикистан": 215, "уганда": 216, "белиз": 217,
	"северная корея": 164, "самоа": 219, "барбадос": 220, "соединенные штаты америки": 221, "эфиопия": 222,
	"южно-африканская республика": 42, "армения": 224, "багамские острова": 178, "багамы": 178, "доминика": 226,
	"либерия":    227,
	"антарктида": 228, "индия": 229, "португалия": 230, "эквадор": 231}