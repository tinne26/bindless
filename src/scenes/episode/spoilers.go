package episode

import "github.com/tinne26/bindless/src/lang"

// yo, don't spoil the game for yourself if you haven't played yet!

var episodeImgPaths = []string {
	"assets/scenes/city_mask.png",
	"assets/scenes/city_mask.png", // would have loved different images...
	"assets/scenes/city_mask.png", // ...but didn't have the time
	"assets/scenes/city_mask.png",
	"assets/scenes/city_mask.png",
	"assets/scenes/city_mask.png",
	"assets/scenes/city_mask.png",
	"assets/scenes/city_mask.png",
	"assets/scenes/city_mask.png",
}

var episodesRawText = []*lang.Text {
	lang.NewText(
		"\x08Ritapola outskirts, 29 June 3589\x07\x05\n" +
		"\x0BMirko ran right in front of the automaton.\x03\n" +
		"\x0B\x09S\x10C\x10A: \x01\x07Please remain clear of the walkway while I carry out my cleaning duties.\x03\x02\n" +
		"\x09Mirko: \x01\x07Wait a second! I've been looking for you all day!\x03\x02\n" +
		"\x09S\x10C\x10A: \x01\x07-Please remain clear of the walkway w\x0Dh\x0Ci\x0Dle I\x0C\x05\x04\x02\n" +
		"\x0BThe machine had suddenly frozen.\x03\x04\n" +
		"\x0B\x09S\x10C\x10A: \x01\x07...\x02\n" +
		"\x09S\x10C\x10A: \x01\x07Unable to identify individual.\x04\x02\n" +
		"\x09Mirko: \x01\x08Exactly what I was hoping for.\x04\x02\n" +
		"\x09S\x10C\x10A: \x01\x07Unable to identify individual, initiating emergency paralysis \x0Cs\x0De\x0Cq\x0D-#!\x0C\x04\x02\n\n" +
		"Mirko quickly grabbed his modulated MSP and slapped it onto the street cleaner automaton.\x05\n" +
		"\x0B\x09Mirko: \x01\x07Shut up already.\x02\n" +
		"\x0B\x08>>",
		"\x08Afueras de Ritapola, 29 de Junio de 3589\x07\x05\n" +
		"\x0BMirko saltó justo en frente del automata.\x03\n" +
		"\x0B\x09A\x10L\x10P: \x01\x07Por favor apártese de la vía mientras realizo mis tareas de limpieza.\x03\x02\n" +
		"\x09Mirko: \x01\x07Espera! Te he estado buscando durante todo el día!\x03\x02\n" +
		"\x09A\x10L\x10P: \x01\x07-Por favor apártese de la vía m\x0Di\x0Cen\x0Dtras\x0C\x05\x04\x02\n" +
		"\x0BLa máquina se congeló de repente.\x05\x04\n" +
		"\x0B\x09A\x10L\x10P: \x01\x07...\x02\n" +
		"\x09A\x10L\x10P: \x01\x07Incapaz de identificar al individuo.\x04\x02\n" +
		"\x09Mirko: \x01\x08Justo lo que esperaba.\x04\x02\n" +
		"\x09A\x10L\x10P: \x01\x07Incapaz de identificar al individuo, iniciando parálisis de emerge\x0C\x0Dn\x0Cc\x0D-#!\x0C\x05\x02\n\n" +
		"Mirko rápidamente cogió su unidad de MSP modulada y se la pegó al automata barrendero.\x05\n" +
		"\x0B\x09Mirko: \x01\x07Cállate ya.\x02\n" +
		"\x0B\x08>>",
		"\x08Afores de Ritapola, 29 de Juny de 3589\x07\x05\n" +
		"\x0BEn Mirko es va plantar just davant l'automata.\x03\n" +
		"\x0B\x09A\x10N\x10P: \x01\x07Si us plau aparti's de la via mentre realitzo les tasques de neteja.\x03\x02\n" +
		"\x09Mirko: \x01\x07Espera! T'he estat buscant durant tot el dia!\x03\x02\n" +
		"\x09A\x10N\x10P: \x01\x07-Si us plau aparti's de la via m\x0De\x0Cnt\x0Dre\x0C\x05\x04\x02\n" +
		"\x0BLa màquina es va congelar de cop.\x05\x04\n" +
		"\x0B\x09A\x10N\x10P: \x01\x07...\x02\n" +
		"\x09A\x10N\x10P: \x01\x07Incapaç d'identificar l'individu.\x04\x02\n" +
		"\x09Mirko: \x01\x08Precisament el que esperava.\x04\x02\n" +
		"\x09A\x10N\x10P: \x01\x07Incapaç d'identificar l'individu, iniciant paràlisis d'emergè\x0C\x0Dn\x0Cc\x0D-#!\x0C\x05\x02\n\n" +
		"En Mirko ràpidament va agafar la seva unitat MSP modulada i la va plantar sobre l'automata de neteja.\x05\n" +
		"\x0B\x09Mirko: \x01\x07Calla d'una vegada.\x02\n" +
		"\x0B\x08>>",
	),
	lang.NewText(
		"\x08Ritapola, MGNT Research Lab, 29 June 3589\x07\x05\n" +
		"\x0BA few hours after the S\x10C\x10A test, Mirko moved to the research facility that MGNT Industries had in the city.\n" +
		"\x0BThe magnetic unit of a cleaner automaton had almost no security, but here things could be different.\n" +
		"\x0BFor starters, he had to get past the door.\n" +
		"\x0B\x09Mirko: \x01\x08Will this even work..?\x02\n" +
		"\x0B\x08>>",
		"no traducido",
		"no traduït",
	),
	lang.NewText(
		"\x08Ritapola, inside MGNT Research Lab\x07\x05\n" +
		"\x0BMirko gained access to the facility and started searching the rooms, but a guard automaton detected him and quickly approached.\n" +
		"\x0B\x09Mirko: \x01\x08Aarrghh... wish I didn't have to meet you.\x07\x04\x02\n" +
		"\x0BThe automaton attempted to stop Mirko with magnetic commands, but like earlier, the attempts had no effect.\x04\n" +
		"\x0BMirko didn't have an MSP implanted in the spine anymore.\n" +
		"\x0B\x08>>",
		"no traducido",
		"no traduït",
	),
	lang.NewText(
		"\x08Ritapola, inside MGNT Research Lab\x07\x05\n" +
		"\x0BAfter disabling the automaton guard, Mirko continued the search.\n" +
		"\x0BWith only a few rooms left to examine, he finally found what he was looking for.\n" +
		"\x0B\x09Mirko: \x01\x08These are definitely not easy to come by. Let's hope they don't notice them missing too soon..\x07\x04\x02\n" +
		"\x0BMirko grabbed the two military-grade MSP units and promptly left the scene.\n" +
		"\x0B\x08>>",
		"no traducido",
		"no traduït",
	),
	lang.NewText(
		"\x08Ritapola, hidden drop-off location, 12 July 3589\x07\x05\n" +
		"\x0BLike every day, Mirko checked the hidden spot.\n" +
		"\x0B\x09Mirko: \x01\x08\x0D!\x0C\x02\x04\n" +
		"\x0B\x07He snatched the note and eagerly read it:\x04\n" +
		"\x0B\x09Jana: \x01\x08I can't guarantee it will be 100% safe, but it's the best I could do. I should be the one doing this, not you. I know Joseph's your brother, but it kills me having to stay here and wait. I'm even tempted to remove my MSP myself. I know, I know, don't worry, I won't.\x03\n" +
		"\x0BDon't push it too hard. And good luck out there.\x02\x05\n" +
		"\x0B\x07Jana had finally done it.\x04\n" +
		"\x0B\x08The ability \x09[Switch]\x08 has been unlocked.\n" +
		"\x0B\x08>>",
		"no traducido",
		"no traduït",
	),
	lang.NewText(
		"\x08Marunka Machart's cottage, 14 July 3589\x07\x05\n" +
		"\x0BMirko had been waiting in hidding for the last 6 hours. Marunka Machart, the Leadership member suspected of Joseph's disappearance, had left the place an hour ago.\x04\n" +
		"\x0BIt was time to find what was inside that place.\n" +
		"\x0BMirko approached silently and tried to sneak, but the attempt for stealth was futile. As soon as he jumped the fence, a guard automaton popped up from nowhere.\n" +
		"\x0BDisastrously for Mirko, the guard didn't go right away after him, but retreated to report the breach first.\n" +
		"\x0BAs soon as alarms went off, it turned again towards Mirko.\n" +
		"\x0B\x09Mirko: \x01\x08Tsk.\x02\x03\x07\n" +
		"\x0B\x08>>",
		"no traducido",
		"no traduït",
	),
	lang.NewText(
		"\x08Marunka Machart's cottage, 14 July 3589\x07\x05\n" +
		"\x0BMirko really struggled to disable the guard automaton, so there was no time to rest. Security would arrive soon.\x04\n" +
		"\x0BThe main floor looked clear, so Mirko went directly to the basement. There, a high-security door was preventing access.\n" +
		"\x0B\x09Mirko: \x01\x08Damn, this is \x0Ereally going to be\x0C a close call...\x02\x07\n" +
		"\x09Mirko: \x01\x08\"Don't push it too hard\"..? Sorry, Jana.\x02\x07\n" +
		"\x0B\x08>>",
		"no traducido",
		"no traduït",
	),
	lang.NewText(
		"\x08Marunka Machart's basement, 14 July 3589\x07\x05\n" +
		"\x0BThe interior was spacious. A man was sitting on a couch, reading a book. He immediately got up when he saw Mirko.\n" +
		"\x0B\x09Man: \x01\x07Who..? \x0EWho are you?\x0C\x02\x04\n" +
		"\x0BBefore Mirko could respond, a woman came running from another room.\n" +
		"\x0B\x09Woman: \x01\x07Did you come to rescue us?\x02\n" +
		"\x0BHer eyes told him everything. High-profile members of Leadership abusing MSP units to turn people into slaves.. was no longer just a rumor.\n" +
		"\x0B\x09Man: \x01\x07-Kayla, don't-\x02\x04\n" +
		"\x09Mirko: \x01\x07Something like that. But we don't have time. Is Joseph Tatar here?\x02\x04\n" +
		"\x09Kayla: \x01\x07Joseph was taken somewhere else two weeks ago.\x02\n" +
		"\x0BA third man came out of the hallway.\x05 As much as he wanted to find his brother, Mirko had to free them first.\n" +
		"\x0B\x08>>",
		"no traducido",
		"no traduït",
	),
}
