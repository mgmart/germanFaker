package germanFaker

var germanMaleFirstNames = []string{"Alexander", "Benjamin", "Christian",
	"Daniel", "Erik", "Felix", "Günther", "Hans", "Ivan", "Jakob", "Karl",
	"Lukas", "Manuel", "Niklas", "Oliver", "Peter", "Quirin", "Robert",
	"Stefan", "Thomas", "Ulrich", "Viktor", "Werner", "Xaver", "Yannick",
	"Zacharias"}

var germanFemaleFirstNames = []string{"Anna", "Birgit", "Christina", "Daniela",
	"Eva", "Friederike", "Gisela", "Helga", "Ingrid", "Johanna", "Katja",
	"Laura", "Maria", "Nadine", "Olga", "Petra", "Quirina", "Renate",
	"Sabine", "Theresa", "Ulrike", "Victoria", "Waltraud", "Xenia",
	"Yvonne", "Zoe"}

var germanLastNames = []string{"Müller", "Schmidt", "Schneider", "Fischer",
	"Weber", "Meyer", "Wagner", "Becker", "Schulz", "Hoffmann", "Schäfer",
	"Koch", "Bauer", "Richter", "Klein", "Wolf", "Schröder", "Neumann",
	"Schwarz", "Zimmermann", "Braun", "Krüger", "Hofmann", "Hartmann", "Lange",
	"Schmitt", "Werner", "Schmitz", "Krause", "Meier", "Lehmann", "Schmid",
	"Schulze", "Maier", "Köhler", "Herrmann", "König", "Walter", "Mayer",
	"Huber", "Kaiser", "Fuchs", "Peters", "Lang", "Scholz", "Möller", "Weiß",
	"Jung", "Hahn", "Schubert"}

var cityInfos = []CityInfo{
	{Name: "Berlin", AreaCode: "030", PostCode: "10115"},
	{Name: "Hamburg", AreaCode: "040", PostCode: "20095"},
	{Name: "München", AreaCode: "089", PostCode: "80331"},
	{Name: "Löln", AreaCode: "0221", PostCode: "50667"},
	{Name: "Frankfurt", AreaCode: "069", PostCode: "60311"},
	{Name: "Stuttgart", AreaCode: "0711", PostCode: "70173"},
	{Name: "Düsseldorf", AreaCode: "0211", PostCode: "40213"},
	{Name: "Leipzig", AreaCode: "0341", PostCode: "04109"},
	{Name: "Dortmund", AreaCode: "0231", PostCode: "44135"},
	{Name: "Essen", AreaCode: "0201", PostCode: "45127"},
	{Name: "Bremen", AreaCode: "0421", PostCode: "28195"},
	{Name: "Dresden", AreaCode: "0351", PostCode: "01067"},
	{Name: "Hanover", AreaCode: "0511", PostCode: "30159"},
	{Name: "Nürnberg", AreaCode: "0911", PostCode: "90402"},
	{Name: "Duisburg", AreaCode: "0203", PostCode: "47051"},
	{Name: "Bochum", AreaCode: "0234", PostCode: "44787"},
	{Name: "Wuppertal", AreaCode: "0202", PostCode: "42103"},
	{Name: "Bielefeld", AreaCode: "0521", PostCode: "33602"},
	{Name: "Bonn", AreaCode: "0228", PostCode: "53111"},
	{Name: "Münster", AreaCode: "0251", PostCode: "48143"},
	{Name: "Karlsruhe", AreaCode: "0721", PostCode: "76133"},
	{Name: "Mannheim", AreaCode: "0621", PostCode: "68161"},
	{Name: "Augsburg", AreaCode: "0821", PostCode: "86150"},
	{Name: "Wiesbaden", AreaCode: "0611", PostCode: "65183"},
	{Name: "Gelsenkirchen", AreaCode: "0209", PostCode: "45879"},
	{Name: "Mönchengladbach", AreaCode: "02161", PostCode: "41061"},
	{Name: "Braunschweig", AreaCode: "0531", PostCode: "38100"},
	{Name: "Chemnitz", AreaCode: "0371", PostCode: "09111"},
	{Name: "Kiel", AreaCode: "0431", PostCode: "24103"},
	{Name: "Aachen", AreaCode: "0241", PostCode: "52062"},
	{Name: "Halle (Saale)", AreaCode: "0345", PostCode: "06108"},
	{Name: "Magdeburg", AreaCode: "0391", PostCode: "39104"},
	{Name: "Freiburg im Breisgau", AreaCode: "0761", PostCode: "79098"},
	{Name: "Krefeld", AreaCode: "02151", PostCode: "47798"},
	{Name: "Lübeck", AreaCode: "0451", PostCode: "23552"},
	{Name: "Oberhausen", AreaCode: "0208", PostCode: "46045"},
	{Name: "Erfurt", AreaCode: "0361", PostCode: "99084"},
	{Name: "Mainz", AreaCode: "06131", PostCode: "55116"},
	{Name: "Rostock", AreaCode: "0381", PostCode: "18055"},
	{Name: "Kassel", AreaCode: "0561", PostCode: "34117"},
	{Name: "Hagen", AreaCode: "02331", PostCode: "58095"},
	{Name: "Hamm", AreaCode: "02381", PostCode: "59065"},
	{Name: "Saarbrücken", AreaCode: "0681", PostCode: "66111"},
	{Name: "Mülheim an der Ruhr", AreaCode: "0208", PostCode: "45468"},
	{Name: "Potsdam", AreaCode: "0331", PostCode: "14467"},
	{Name: "Ludwigshafen am Rhein", AreaCode: "0621", PostCode: "67059"},
	{Name: "Oldenburg", AreaCode: "0441", PostCode: "26122"},
	{Name: "Leverkusen", AreaCode: "0214", PostCode: "51373"},
	{Name: "Osnabrück", AreaCode: "0541", PostCode: "49074"},
	{Name: "Solingen", AreaCode: "0212", PostCode: "42651"},
	{Name: "Heidelberg", AreaCode: "06221", PostCode: "69115"},
}

var adjectives = []string{"adäquat", "affektiert", "agil", "akribisch", "antagonistisch", "apathisch", "arriviert", "autokratisch", "banal", "brachial", "Contenance", "designiert", "desolat", "dediziert", "definitiv", "dezidiert", "diabolisch", "diametral", "differenziert", "diffizil", "diffus", "diskutabel", "distinguiert", "effektiv", "effizient", "elanvoll", "eloquent", "eminent", "essenziell", "evident", "exorbitant", "explizit", "expressiv", "fulminant", "generös", "gravierend", "heterogen", "homogen", "ikonisch", "illustrativ", "impraktikabel", "inadäquat", "inakzeptabel", "indiskutabel", "infernalisch", "informell", "initial", "irrelevant", "komplex", "kongenial", "konsistent", "konsterniert", "kontinuierlich", "konträr", "kurios", "lapidar", "legitim", "lethargisch", "loyal", "lukrativ", "maliziös", "maniriert", "marginal", "martialisch", "medioker", "melodramatisch", "morbid", "nebulös", "neuralgisch", "normativ", "obligatorisch", "obsolet", "omnipotent", "opportun", "opulent", "pekuniär", "penibel", "perfide", "pittoresk", "pointiert", "prädestiniert", "prägnant", "präsent", "prätentiös", "prekär", "prosaisch", "redundant", "relevant", "renitent", "renommiert", "respektabel", "restriktiv", "rudimentär", "sakrosankt", "satanisch", "saturiert", "servil", "skurril", "stringent", "subsidiär", "subtil", "substanziell", "superb", "theatralisch", "titanisch", "tolerabel", "tradiert", "trist", "trivial", "vakant", "vehement", "versiert"}

var substantives = []string{"Absenz", "Agilität", "Agonie", "Akribie", "Ambition", "Ambivalenz", "Analyse", "Antipathie", "Antithese", "Apathie", "Approximation", "Aspiration", "Assoziation", "Attitüde", "Attraktion", "Autarkie", "Authentizität", "Aversion", "Contenance", "Dedikation", "Dependance", "Desavouierung", "Destination", "Dezenz", "Dignität", "Diffamierung", "Differenz", "Direktive", "Diskrepanz", "Diskrimination", "Disproportionalität", "Distinktion", "Divergenz", "Doktrin", "Dualität", "Elementarität", "Eloquenz", "Euphorie", "Exkulpation", "Familiarität", "Fiktion", "Fluktuation", "Fortune", "Gravität", "Hypothese", "Imagination", "Implikation", "Imponderabilie", "Indifferenz", "Inkonsequenz", "Innovation", "Insistenz", "Integration", "Intervention", "Intimität", "Inspiration", "Invektive", "Präsenz", "Kohärenz", "Kohäsion", "Komposition", "Kontroverse", "Kontemplation", "Konvergenz", "Konversation", "Konzilianz", "Kreativität", "Krux", "Lakonie", "Lethargie", "Malaise", "Manie", "Maxime", "Mission", "Modifikation", "Modifizierung", "Obskurität", "Operation", "Perseveranz", "Phantasmagorie", "Phase", "Prämisse", "Präferenz", "Präpotenz", "Präsumtion", "Prävalenz", "Präzision", "Punktualität", "Ratio", "Reflexion", "Relation", "Relevanz", "Renitenz", "Reputation", "Retrospektive", "Rigorosität", "Schimäre", "Sentenz", "Servilität", "Signifikanz", "Suada", "Supposition", "Temperenz", "Tirade", "Transzendenz", "Usance", "Variante", "Virtualität", "Vita", "Zivilität"}

var companyForm = []string{
	//Personengesellschaften
	"GbR", "KG", "AG & Co. KG", "GmbH & Co. KG", "Limited & Co. KG",
	// "Stiftung & Co. KG", "Stiftung GmbH & Co. KG",
	"UG (haftungsbeschränkt) & Co. KG", "OHG", "GmbH & Co. OHG", "AG & Co. OHG",
	// "Partenreederei", "PartG", "PartG mbB", "Stille Gesellschaft",
	// Kapitalgesellschaften:
	"AG", "gAG", "GmbH",
	// "gGmbH", "InvAG",
	"KGaA",
	// "AG & Co. KGaA", "SE & Co. KGaA", "GmbH & Co. KGaA",
	// "Stiftung & Co. KGaA",
	// "REIT-AG",
	"UG (haftungsbeschränkt)",
	// Sonstige Rechtsformen:
	// "AöR", "eG", "Eigenbetrieb", "Einzelunternehmen", "e. V.", "KöR", "Regiebetrieb", "Stiftung", "VVaG"
}
