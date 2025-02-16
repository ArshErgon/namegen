package namegen

var (
	maoriMaleFirstNames = []string{
	"Aata",
	"Ahika",
	"Aiali'i",
	"Akahata",
	"Amahau",
	"Anewa",
	"Aperaj",
	"Aperahama",
	"Arapeta",
	"Ari",
	"Atama",
	"Atawhai",
	"Eruera",
	"Etera",
	"Hahona",
	"Haimona",
	"Harata",
	"Hau",
	"Hehu",
	"Heketoro",
	"Henare",
	"Hohepa",
	"Huatare",
	"Hunapo",
	"Hunu",
	"Iarere",
	"Iehohapata",
	"Ieni",
	"Ietepere",
	"Ietoro",
	"Ihaia",
	"Ihaka",
	"Ihu",
	"Ikaroa",
	"Iorama",
	"Iorangi",
	"Iotama",
	"Iraia",
	"Irirangi",
	"Iwi",
	"Kaihautu",
	"Kamaka",
	"Kiritopa",
	"Koraka",
	"Maaka",
	"Mahaka",
	"Makareta",
	"Marcellin",
	"Marcellino",
	"Matiu",
	"Mere",
	"Mikaere",
	"Purta",
	"Puta",
	"Rangi",
	"Rapata",
	"Tangaroa",
	"Tawera",
	"Tayn",
	"Tipene",
	"Tui",
	"Wiremu",
	"Rawiri",
	"Rongo",
	"Ropata",
	"Ruru",
	"Tanemahuta",
	}
	maoriFemaleFirstNames = []string{
	"Ahere",
	"Ahorangi",
	"Ahurewa",
	"Aiali'i",
	"Aihe",
	"Aio",
	"Airini",
	"Aka",
	"Akenehi",
	"Akona",
	"Ametihita",
	"Amiria",
	"Anahira",
	"Aoatea",
	"Aperira",
	"Aputa",
	"Aramoana",
	"Aranga",
	"Arataki",
	"Arihi",
	"Arihia",
	"Arikiwi",
	"Arona",
	"Aronui",
	"Arorangi",
	"Ataahua",
	"Atarangi",
	"Atawhai",
	"Awhireinga",
	"Awi",
	"Emere",
	"Epa",
	"Erihapeti",
	"Haeata",
	"Haeatatanga",
	"Hahana",
	"Harata",
	"Hau",
	"Hauku",
	"Haumiatiketike",
	"Heeni",
	"Hine",
	"Hinewai",
	"Hoana",
	"Hokaka",
	"Horowai",
	"Huihana",
	"Hunapo",
	"Hunu",
	"Ihapera",
	"Ikaroa",
	"Iorangi",
	"Iraia",
	"Irirangi",
	"Iwa",
	"Iwi",
	"Kaheru",
	"Kahikatea",
	"Kahikatoa",
	"Kahu",
	"Kahuiwi",
	"Kahukura",
	"Kaihautu",
	"Kakarauri",
	"Kamaka",
	"Kaori",
	"Katarine",
	"Kiri",
	"Mahuika",
	"Makareta",
	"Marama",
	"Ngaio",
	"Ngaire",
	"Nyree",
	"Pania",
	"Pounamu",
	"Rongo",
	"Tino Aroha",
	"Tui",
	"Whina",
	"Wikitoria",
	}
	maoriLastNames = []string{
	"Abel",
	"Adsett",
	"Ahu",
	"Ahuahu", 
	"Ahuaiti",
	"Ahunuku",
	"Akena",
	"Akuhata", 
	"Akurangi",
	"Albert",
	"Alexander", 
	"Amaru",
	"Anaru",
	"Anderson",
	"Andrews",
	"Anihana",
	"Aperahama", 
	"Apiata",
	"Araiteuru", 
	"Arakopeka",
	"Arano",
	"Arapata",
	"Arapeta",
	"Aratiera",
	"Armstrong",
	"Ashby",
	"Aspinall",
	"Aupouri",
	"Averill",
	"Awa",
	"Awarau", 
	"Awaroa",
	"Babbington",
	"Baker",
	"Baldwin",
	"Balmforth",
	"Barber",
	"Barlow",
	"Berghan",
	"Birch",
	"Bishop",
	"Blackburn",
	"Bonsall",
	"Boxall",
	"Boyce",
	"Boyland",
	"Bristowe",
	"Broughton",
	"Brown",
	"Bruce",
	"Bryers",
	"Buckley",
	"Burch",
	"Burns",
	"Cafter",
	"Campbell",
	"Carter",
	"Cartman",
	"Cassidy",
	"Chapman",
	"Cherrington",
	"Christian",
	"Clark",
	"Clarke",
	"Clayton",
	"Clendon",
	"Clifford",
	"Cobbett",
	"Cochrane",
	"Coffe",
	"Coleman",
	"Colenso",
	"Collen",
	"Collier",
	"Cook",
	"Cooper",
	"Cope",
	"Couch",
	"Crawford",
	"Croft",
	"Crosbie",
	"Cross",
	"Cross",
	"Currie",
	"Davies",
	"Dawn",
	"De Har",
	"De Thierry",
	"De Tomatois",
	"Delamere",
	"Diamond",
	"Douglas",
	"Duncan",
	"Edmonds",
	"Eparaima",
	"Epiha",
	"Eruera",
	"Erueti",
	"Espy",
	"Etana",
	"Eteatara",
	"Fairlie",
	"Faithful",
	"Farley",
	"Faulkner",
	"Fergusson",
	"Ferris",
	"Flavell",
	"Foy",
	"Gardner",
	"George",
	"Gerard",
	"Gibbons",
	"Gilbert",
	"Gilshanon",
	"Going",
	"Goldsmith",
	"Gollan",
	"Grace",
	"Green",
	"Greening",
	"Grey",
	"Gudgeon",
	"Guilding",
	"Haara",
	"Haenga",
	"Haig",
	"Haimona",
	"Hakaraia",
	"Hale",
	"Hansey",
	"Hape",
	"Hapeta",
	"Haratua",
	"Harawira",
	"Hare-Harris",
	"Harepure",
	"Harri",
	"Harris",
	"Harrison",
	"Harvey",
	"Hau",
	"Haua",
	"Hauraki",
	"Haurangi",
	"Hawai",
	"Hawaikiirangi",
	"Hayes",
	"Heihei",
	"Heipounamu",
	"Heke",
	"Hemi",
	"Henare",
	"Henderson",
	"Hepehi",
	"Hepehi",
	"Hepehi",
	"Heperi",
	"Heperi",
	"Hepi",
	"Heremaia",
	"Heretini",
	"Herewaka",
	"Herewini",
	"Herivel",
	"Heta",
	"Hetaraka",
	"Heuheu",
	"Hika",
	"Hiku",
	"Hillman",
	"Hinekurangi",
	"Hinga",
	"Hipirini",
	"Hira",
	"Hirau",
	"Hodgkinson",
	"Hoekai",
	"Hoey",
	"Hohaia",
	"Hohepa",
	"Hokaanga",
	"Hokima",
	"Hona",
	"Hone",
	"Hongi",
	"Hooper",
	"Hopara",
	"Hopetaia",
	"Horahora",
	"Horo",
	"Horomona",
	"Hotere",
	"Hotini",
	"Houkamau",
	"Houston",
	"Howard",
	"Huahuake",
	"Hubbard",
	"Huihui",
	"Huinga",
	"Huirama",
	"Hulme",
	"Hurumanu",
	"Ihaka",
	"Iraia",
	"Irimana",
	"Jellick",
	"Johnston",
	"Johnstone",
	"Johnstone",
	"Jones",
	"Joseph",
	"Joyce",
	"Jury",
	"Kaahu",
	"Kahaki",
	"Kahotea",
	"Kahu",
	"Kahua",
	"Kahukura-Ao",
	"Kaihe",
	"Kaihe-Nori",
	"Kaipo",
	"Kaire",
	"Kaitoke",
	"Kakano",
	"Kake",
	"Kamau",
	"Kamizona",
	"Kanara",
	"Kanaraki",
	"Kanuta",
	"Kapa",
	"Karaha",
	"Karaitiana",
	"Karaka",
	"Karetu",
	"Karo",
	"Karora",
	"Karoro",
	"Katae",
	"Katene",
	"Kauere",
	"Kaui",
	"Kauwhata",
	"Kawana",
	"Kawhena",
	"Kawhia",
	"Kawiti",
	"Keepa",
	"Keiha",
	"Kekoki",
	"Kemp",
	"Kere",
	"Kerei",
	"Kihi",
	"Kima",
	"King",
	"Kingi",
	"Kino",
	"Kirikau",
	"Kirikiri",
	"Kiriwa",
	"Kitekaraka",
	"Kiwara",
	"Kiwikiwi",
	"Koara",
	"Kohai",
	"Kohere",
	"Kohuawaitai",
	"Koia",
	"Komene",
	"Koni",
	"Kopa",
	"Korimete",
	"Korohia",
	"Koroi",
	"Koroneho",
	"Koroua",
	"Koti",
	"Kotia",
	"Kowhai",
	"Kupe",
	"Kururangi",
	"Lawrence",
	"Leaf",
	"Leef",
	"Lemon",
	"Lockwood",
	"London",
	"Lovie",
	"Lundon",
	"Maaka",
	"MacDonald",
	"Maere",
	"Mahanga",
	"Mahia",
	"Mahinaroa",
	"Mahue",
	"Mahuika",
	"Mahuta",
	"Maihi",
	"Maika",
	"Maioha",
	"Makara",
	"Makehunia",
	"Makene",
	"Maki",
	"Makiri",
	"Mana",
	"Manawakume",
	"Mane",
	"Manihera",
	"Mano",
	"Manuel",
	"Manuera",
	"Mapi",
	"Mapuna",
	"Mara",
	"Maraki",
	"Mardon",
	"Marmon",
	"Marsden",
	"Marsh",
	"Marshall",
	"Martin",
	"Maru",
	"Marunui",
	"Matahiki",
	"Mataotao",
	"Matauru",
	"Matehe",
	"Matenga",
	"Materonea",
	"Matia",
	"Matiaha",
	"Mato",
	"Matthews",
	"Mauheni",
	"Maukau",
	"Maukino",
	"Maumau",
	"Maurirere",
	"Mauwhata",
	"Maxwell",
	"mcBride",
	"McCaskill",
	"McCrae",
	"McGhee",
	"McGuire",
	"McKenzie",
	"McLean",
	"McNab",
	"Meha",
	"Mete",
	"Mihaka",
	"Mika",
	"Mika",
	"Mikaera",
	"Mikaere",
	"Moana",
	"Moeke",
	"Moetara",
	"Mohi",
	"Mokena",
	"Moore",
	"Morete",
	"Morgan",
	"Morrell",
	"Morunga",
	"Motu",
	"Muhu",
	"Mulligan",
	"Munn",
	"Muramura",
	"Muriwai",
	"Murphy",
	"Murphy-Mapi",
	"Murray",
	"Naera",
	"Naere",
	"Napia",
	"Napier",
	"Napier-Nepia",
	"Nathan",
	"Nathan",
	"Neho",
	"Nehu",
	"Nehua",
	"Nehua-Bryers",
	"Nene",
	"Nepia",
	"Netana",
	"Ngakuru",
	"Ngakuru",
	"Ngakuru",
	"Ngapera",
	"Ngapiri",
	"Ngapua",
	"Ngarangi",
	"Ngare",
	"Ngarimu",
	"Ngarino",
	"Ngaro",
	"Ngaru",
	"Ngata",
	"Ngatai",
	"Ngatoko",
	"Ngatoro",
	"Ngatote",
	"Ngaturu",
	"Ngawai",
	"Ngawaka",
	"Ngawati",
	"Ngawharau",
	"Ngawhau",
	"Ngere",
	"Ngongohau",
	"Ngutuawa",
	"Nicholls",
	"Nihoniho",
	"Nikora",
	"Noanoa",
	"Nohoaaka",
	"Nohorua",
	"Nom",
	"Nopera",
	"North",
	"Northover",
	"O'Hara",
	"O'Neil",
	"Oates",
	"Ogle",
	"Olsen",
	"Onui",
	"Ormsby",
	"Otene",
	"Otini",
	"Owen",
	"Paati",
	"Paea",
	"Paemoana",
	"Paewai",
	"Pahau",
	"Pahewa",
	"Pahina",
	"Pahura",
	"Paikea",
	"Paikia",
	"Pain",
	"Paiwai",
	"Pakakohi",
	"Pakerau",
	"Paki",
	"Pakihi",
	"Paku",
	"Palmer",
	"Pangari",
	"Paora",
	"Papahi",
	"Papare",
	"Paparoa",
	"Paphia",
	"Paputene",
	"Para",
	"Paraha",
	"Paraki",
	"Parangi",
	"Paranihi",
	"Paraone",
	"Parapara",
	"Parata",
	"Paratene",
	"Pare",
	"Pareraukawa",
	"Parihi",
	"Patene",
	"Pateuru",
	"Patuone",
	"Pau",
	"Paura",
	"Pawhara",
	"Peeni",
	"Pehi",
	"Peihopa",
	"Peita",
	"Peka",
	"Pekama",
	"Pene",
	"Peneha",
	"Peneti",
	"Pepere",
	"Pera",
	"Pere",
	"Pereto",
	"Petiha",
	"Petiha Paku",
	"Pewhairangi",
	"Pickering",
	"Pihau",
	"Pihema",
	"Pihoihoi",
	"Piipi",
	"Pikimauri",
	"Piri",
	"Pirini",
	"Piripi",
	"Pita",
	"Pitman",
	"Poakatahi",
	"Poata",
	"Podlie",
	"Pohatu",
	"Pohe",
	"Pohipi",
	"Poito",
	"Pokai",
	"Pokiha",
	"Pomare",
	"Ponguru",
	"Porangi",
	"Porotiti",
	"Porter",
	"Potae",
	"Potaka",
	"Pou",
	"Pouaka",
	"Poutu",
	"Prihi",
	"Prime",
	"Puhi",
	"Puhipuhi",
	"Pukeroa",
	"Pumipi",
	"Pure",
	"Pureueke",
	"Puriri",
	"Puru",
	"Rackstraw",
	"Rahiri",
	"Rakena",
	"Rako",
	"Rameka",
	"Ranana",
	"Ranapia",
	"Rangarunga",
	"Rangi",
	"Rangikaputa",
	"Rangimahora",
	"Rangiuaia",
	"Rangiuia",
	"Rankin",
	"Rapaea",
	"Rapana",
	"Rapira",
	"Rare",
	"Raroa",
	"Rata",
	"Raumata",
	"Rawiri",
	"Raymond",
	"Reedy",
	"Rei",
	"Reihana",
	"Reka",
	"Reneti",
	"Repia",
	"Reti",
	"Retter",
	"Reupena",
	"Reweti",
	"Rewha",
	"Rewi",
	"Rewiri",
	"Ria",
	"Richards",
	"Riini",
	"Rika",
	"Rikihana",
	"Ripia",
	"Riwaiti",
	"Robinson",
	"Rodgers",
	"Rogan",
	"Rogers",
	"Rohe",
	"Rongo",
	"Ropata",
	"Ropiha",
	"Ross",
	"Round",
	"Ruarau",
	"Ruatapu",
	"Rudolph",
	"Ruka",
	"Rukopo",
	"Rukuata",
	"Rungarunga",
	"Rurawhe",
	"Ruru",
	"Rushbrook",
	"Rutene",
	"Ruwhiu",
	"Ryan",
	"Satchel",
	"Semmell",
	"Shedlock",
	"Shelford",
	"Sheppard",
	"Shortland",
	"Simon",
	"Skudder",
	"Slade",
	"Smith",
	"Smith",
	"Smythe",
	"Snowden",
	"Squire",
	"Stephens",
	"Stevens",
	"Stevenson",
	"Stewart",
	"Stickle",
	"Stirling",
	"Stowell",
	"Subritsky",
	"Sullivan",
	"Taaheke",
	"Taana",
	"Taare",
	"Taewa",
	"Tahana",
	"Tahau",
	"Tahena",
	"Tahere",
	"Tahi",
	"Tahiwi",
	"Tahua",
	"Tahuna",
	"Tai",
	"Taiapa",
	"Taiko",
	"Taimona",
	"Taingahue",
	"Tairua",
	"Taitoko",
	"Taiuru",
	"Taiwhanga",
	"Taka",
	"Tamaki",
	"Tamanuiotera",
	"Tamararo",
	"Tamati",
	"Tamawhaikai",
	"Tamepo",
	"Tamihana",
	"Tamitere",
	"Tana",
	"Tanapo",
	"Tane",
	"Taniere",
	"Taniora",
	"Taniwe",
	"Taniwha",
	"Taonui",
	"Tara",
	"Taranui",
	"Tarapata",
	"Tarapata",
	"Tarapata",
	"Tarapehu",
	"Tarawa",
	"Tarawau",
	"Tareha",
	"Tarewa",
	"Tari",
	"Taro",
	"Tatana",
	"Tataurangi",
	"Tau",
	"Tauehe",
	"Tauhinu",
	"Taukamo",
	"Taura",
	"Tauranga",
	"Taurewa",
	"Tauroa",
	"Tautahi",
	"Tautari",
	"Tautau",
	"Tautuhi",
	"Tawae",
	"Tawhai",
	"Tawhena",
	"Tawhera",
	"Tawhiti",
	"Taylor",
	"Te Apatahi",
	"Te Awarau",
	"Te Ehutu",
	"Te Ekewhare",
	"Te Haenga",
	"Te Hau",
	"Te Hei",
	"Te Hemoata",
	"Te Hiamo",
	"Te Hiko",
	"Te Hokio",
	"Te Ihi",
	"Te Irimana",
	"Te Kahaki",
	"Te Kahukino",
	"Te Kani",
	"Te Kiki",
	"Te Matauru",
	"Te Moana",
	"Te O-Paipa",
	"Te Ohaere",
	"Te Owai",
	"Te Pea",
	"Te Piri",
	"Te Rangi",
	"Te Rangitiwarawara",
	"Te Ripi",
	"Te Rure",
	"Te Ruru",
	"Te Taha",
	"Te Ururangi",
	"Te Waititi",
	"Te Whanau",
	"Te Whao",
	"Te Wheoro",
	"Tehei",
	"Teira",
	"Teremauri",
	"Tete",
	"Thomas",
	"Thompson",
	"Tiarite",
	"Tiatoa",
	"Tibble",
	"Tikaro",
	"Timu",
	"Tinirau",
	"Tipene",
	"Tirarau",
	"Titore",
	"To",
	"Toanui",
	"Toe",
	"Toenga",
	"Tohe",
	"Toheriri",
	"Tohia",
	"Tohu",
	"Toi",
	"Toia",
	"Toka",
	"Tokara",
	"Tokawahe",
	"Toki",
	"Tokomauri",
	"Toms",
	"Tongia",
	"Topi",
	"Topia",
	"Toroa",
	"Torona",
	"Toto",
	"Tripbook",
	"Tuari",
	"Tuau",
	"Tuauki",
	"Tuhaere",
	"Tuhaere-Paora",
	"Tuhaka",
	"Tuhiwai",
	"Tuhua",
	"Tui",
	"Tuoro",
	"Tupakihi",
	"Turaki",
	"Turatawhiti",
	"Turei",
	"Turoa",
	"Tuwhare",
	"Wa",
	"Waaka",
	"Waara",
	"Waata",
	"Waetford",
	"Wahawaha",
	"Waiakapakapa",
	"Waih",
	"Wainokenoke",
	"Waipapa",
	"Wairoa",
	"Waitai",
	"Waiti",
	"Waititi",
	"Waitoa",
	"Waiwhakaruru",
	"Wake",
	"Walker",
	"Wall",
	"Walters",
	"Wanoa",
	"Wara",
	"Ward",
	"Warena",
	"Waretini",
	"Watene",
	"Watikena",
	"Watters",
	"Wau",
	"Weaver",
	"Wehe",
	"Wehinuku",
	"Wehiwehi",
	"Weir",
	"Welsh",
	"Wepiha",
	"Werohia",
	"Whaanga",
	"Whai",
	"Whainoa",
	"Whaitiri",
	"Whakaawa",
	"Whakakana",
	"Whakataka",
	"Wharehinga",
	"Wharekawa",
	"Wharemate",
	"Wharenikau",
	"Wharerahi",
	"Wharerau",
	"Whareraupo",
	"Wharewaka",
	"Whata",
	"Whatarau",
	"Whatipu",
	"Whautere",
	"Wheau",
	"Wheoki",
	"White",
	"Whittaker",
	"Whiu",
	"Wiheta",
	"Wihongi",
	"Wikaira",
	"Wikeepa",
	"Wiki",
	"Williams",
	"Wilson",
	"Wimoka",
	"Winiana",
	"Winiata",
	"Winter",
	"Wiperi",
	"Wipou",
	"Wiremu",
	"Wirihana",
	"Witana",
	"Witehira",
	"Woods",
	"Wynyard",
	"Yates",
	"York",
	"Young",
	}
)