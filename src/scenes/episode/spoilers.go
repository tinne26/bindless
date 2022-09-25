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
		"\x08Mirko's home, Ritapola, 5 June 3589\x07\x05\n" +
		"\x0BMirko opened the door and invited her inside.\x03\n" +
		"\x0B\x09Jana: \x01\x07Good news, Mirko. The operating room will be free this thursday.\x02\n" +
		"\x09Mirko: \x01\x07So.. the time has finally come?\x04 That's good.\x02\x03\n" +
		"\x09Jana: \x01\x07Ye\x0Eah\x0C..\x02\n" +
		"\x0BHer voice trembled halfway.\x04 She looked away for an instant, but quickly resumed with her usual playful tone:\x03\n" +
		"\x0B\x09Jana: \x01\x07Oh, I've got your learning materials ready too. \x05I tried to make it accessible even for \x0Edummies like you\x0C, so make sure to engrave it in your brain!\x02\x04\n" +
		"\x0B\x08Knowing her, that \"dummy\" level probably escaped most average adults.\x03\n" +
		"\x0B\x09Mirko: \x01\x07Never thought the day where I'd get excited about homework would finally arrive!\x02\n" +
		"\x09Jana: \x01\x07\x0DHaha\x0C, definitely not the type!\x02\n" +
		"\x0B\x08>>",
		"\x08Casa de Mirko, Ritapola, 5 de junio de 3589\x07\x05\n" +
		"\x0BMirko abrió la puerta y la invitó a entrar.\x03\n" +
		"\x0B\x09Jana: \x01\x07Buenas noticias, Mirko. La sala de operaciones estará libre este jueves.\x02\n" +
		"\x09Mirko: \x01\x07Así que.. ha llegado la hora?\x04 Eso es bueno.\x02\x03\n" +
		"\x09Jana: \x01\x07Cl\x0Earo\x0C..\x02\n" +
		"\x0BSu voz tembló a la mitad.\x04 Jana apartó la mirada un instante, pero rápidamente recuperó su tono pillo habitual:\x03\n" +
		"\x0B\x09Jana: \x01\x07Oh, también tengo preparado tu material de estudio. \x05He intentado hacerlo accesible incluso para \x0Etortuguitos como tú\x0C, así que asegúrate de grabártelo bien en la cabeza!\x02\x04\n" +
		"\x0B\x08Conociéndola, lo más probable era que sus \"tortuguitos\" escaparan la mayoría de adultos promedio.\x03\n" +
		"\x0B\x09Mirko: \x01\x07Nunca pensé que llegaría el día en que me emocionaría por hacer los deberes!\x02\n" +
		"\x09Jana: \x01\x07\x0DJaja\x0C, definitivamente no das el pego!\x02\n" +
		"\x0B\x08>>",
		"\x08Casa d'en Mirko, Ritapola, 5 de juny de 3589\x07\x05\n" +
		"\x0BEn Mirko va obrir la porta i la va convidar a passar.\x03\n" +
		"\x0B\x09Jana: \x01\x07Bones notícies, Mirko. La sala d'operacions estarà lliure aquest dijous.\x02\n" +
		"\x09Mirko: \x01\x07Així que.. finalment ha arribat l'hora?\x04 Això és bo.\x02\x03\n" +
		"\x09Jana: \x01\x07\x0ESí\x0C..\x02\n" +
		"\x0BLa veu li va tremolar.\x04 La Jana va apartar la mirada un instant, però ràpidament va recuperar el seu to trapella habitual:\x03\n" +
		"\x0B\x09Jana: \x01\x07Ah, també tinc preparat el teu material d'estudi. \x05He intentat fer-lo accessible fins i tot per \x0Etortuguetes como tu\x0C, així que assegura't de gravar-te'l bé al cap!\x02\x04\n" +
		"\x0B\x08Coneixent-la, el més probable era que les seves \"tortuguetes\" escapessin la majoria d'adults.\x03\n" +
		"\x0B\x09Mirko: \x01\x07Mai m'hauria imaginat que arribaria el dia en que m'emocionaria per fer els deures!\x02\n" +
		"\x09Jana: \x01\x07\x0DHaha\x0C, definitivament no dones el perfil!\x02\n" +
		"\x0B\x08>>",
	),
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
		"\x09S\x10C\x10A: \x01\x07Unable to identify individual, initiating emergency paralysis \x0Cs\x0De\x0Cq\x0D-#!\x0C\x04\x02\n" +
		"\x0BMirko quickly grabbed his modulated MSP and slapped it onto the street cleaner automaton.\x05\n" +
		"\x0B\x09Mirko: \x01\x07Shut up already.\x02\n" +
		"\x0B\x08>>",
		"\x08Afueras de Ritapola, 29 de junio de 3589\x07\x05\n" +
		"\x0BMirko saltó justo en frente del autómata.\x03\n" +
		"\x0B\x09A\x10L\x10P: \x01\x07Por favor apártese de la vía mientras realizo mis tareas de limpieza.\x03\x02\n" +
		"\x09Mirko: \x01\x07Espera! Te he estado buscando durante todo el día!\x03\x02\n" +
		"\x09A\x10L\x10P: \x01\x07-Por favor apártese de la vía m\x0Di\x0Cen\x0Dtras\x0C\x05\x04\x02\n" +
		"\x0BLa máquina se congeló de repente.\x03\x04\n" +
		"\x0B\x09A\x10L\x10P: \x01\x07...\x02\n" +
		"\x09A\x10L\x10P: \x01\x07Incapaz de identificar individuo.\x04\x02\n" +
		"\x09Mirko: \x01\x08Justo lo que esperaba.\x04\x02\n" +
		"\x09A\x10L\x10P: \x01\x07Incapaz de identificar individuo, iniciando parálisis de emerge\x0C\x0Dn\x0Cc\x0D-#!\x0C\x04\x02\n" +
		"\x0BMirko rápidamente cogió su unidad de MSP modulada y se la pegó al autómata barrendero.\x05\n" +
		"\x0B\x09Mirko: \x01\x07Cállate ya.\x02\n" +
		"\x0B\x08>>",
		"\x08Afores de Ritapola, 29 de juny de 3589\x07\x05\n" +
		"\x0BEn Mirko es va moure just davant l'autòmata.\x03\n" +
		"\x0B\x09A\x10N\x10P: \x01\x07Si us plau aparti's de la via mentre realitzo les tasques de neteja.\x03\x02\n" +
		"\x09Mirko: \x01\x07Espera! T'he estat buscant durant tot el dia!\x03\x02\n" +
		"\x09A\x10N\x10P: \x01\x07-Si us plau aparti's de la via m\x0De\x0Cnt\x0Dre\x0C\x05\x04\x02\n" +
		"\x0BLa màquina es va congelar de cop.\x03\x04\n" +
		"\x0B\x09A\x10N\x10P: \x01\x07...\x02\n" +
		"\x09A\x10N\x10P: \x01\x07Incapaç d'identificar l'individu.\x04\x02\n" +
		"\x09Mirko: \x01\x08Precisament el que esperava.\x04\x02\n" +
		"\x09A\x10N\x10P: \x01\x07Incapaç d'identificar l'individu, iniciant paràlisis\nd'emergè\x0C\x0Dn\x0Cc\x0D-#!\x0C\x04\x02\n" +
		"\x0BEn Mirko ràpidament va agafar la seva unitat MSP modulada i la va plantar sobre l'autòmata de neteja.\x05\n" +
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
		"\x08Ritapola, Laboratorios MGNT, 29 de junio de 3589\x07\x05\n" +
		"\x0BPocas horas después de la prueba del A\x10L\x10P, Mirko se dirigió al complejo de laboratorios de investigación que Indústrias MGNT tenía en la ciudad.\n" +
		"\x0BLa unidad magnética de un autómata de limpieza tenía un sistema de seguridad muy básico, pero aquí las cosas podrían ser distintas.\n" +
		"\x0BPara empezar, Mirko necesitaba superar la puerta.\n" +
		"\x0B\x09Mirko: \x01\x08Funcionará esto..?\x02\n" +
		"\x0B\x08>>",
		"\x08Ritapola, Laboratoris MGNT, 29 de juny de 3589\x07\x05\n" +
		"\x0BPoques hores després de la prova del A\x10N\x10P, en Mirko es va dirigir al complex de laboratoris d'investigació que Indústries MGNT tenia a la ciutat.\n" +
		"\x0BLa unitat magnètica d'un autòmata de neteja tenia un sistema de seguretat molt bàsic, però aquí les coses podrien ser diferents.\n" +
		"\x0BPer començar, en Mirko necessitava superar la porta.\n" +
		"\x0B\x09Mirko: \x01\x08Funcionarà això..?\x02\n" +
		"\x0B\x08>>",
	),
	lang.NewText(
		"\x08Ritapola, inside MGNT Research Lab\x07\x05\n" +
		"\x0BMirko gained access to the facility and started searching the rooms, but a guard automaton detected him and quickly approached.\x03\n" +
		"\x0B\x09Mirko: \x01\x08Aarrghh... wish I didn't have to meet you.\x07\x04\x02\n" +
		"\x0BThe automaton attempted to stop Mirko with magnetic commands, but like earlier, the attempts had no effect.\x04\n" +
		"\x0BMirko didn't have an MSP implanted in the spine anymore.\n" +
		"\x0B\x08>>",
		"\x08Ritapola, dentro los Laboratorios MGNT\x07\x05\n" +
		"\x0BMirko consiguió acceder al complejo y empezó a buscar en cada habitación, pero un autómata de seguridad rápidamente lo detectó y se acercó a él.\x03\n" +
		"\x0B\x09Mirko: \x01\x08Aarrghh... hubiese preferido no tener que conocerte.\x07\x04\x02\n" +
		"\x0BEl autómata intentó detener a Mirko con comandos magnéticos, pero tal como antes, el intento fue en vano.\x04\n" +
		"\x0BMirko ya no tenía ninguna unidad MSP implantada en la columna.\n" +
		"\x0B\x08>>",
		"\x08Ritapola, dins els Laboratoris MGNT\x07\x05\n" +
		"\x0BEn Mirko va entrar al complex i va començar a buscar en cada habitació, però un autòmata de seguretat ràpidament el va detectar i es va apropar a ell.\x03\n" +
		"\x0B\x09Mirko: \x01\x08Aarrghh... hagués preferit no haver-te de conèixer.\x07\x04\x02\n" +
		"\x0BL'autòmata va tractar d'immobilitzar en Mirko amb directives magnètiques, però tal com abans, l'intent va ser inútil.\x04\n" +
		"\x0BEn Mirko ja no tenia cap unitat MSP implantada a la columna.\n" +
		"\x0B\x08>>",
	),
	lang.NewText(
		"\x08Ritapola, inside MGNT Research Lab\x07\x05\n" +
		"\x0BAfter disabling the automaton guard, Mirko continued the search.\n" +
		"\x0BWith only a few rooms left to examine, he finally found what he was looking for.\n" +
		"\x0B\x09Mirko: \x01\x08These are definitely not easy to come by. Let's hope they don't notice them missing too soon..\x07\x04\x02\n" +
		"\x0BMirko grabbed the two military-grade MSP units and promptly left the scene.\n" +
		"\x0B\x08>>",
		"\x08Ritapola, dentro los Laboratorios MGNT\x07\x05\n" +
		"\x0BTras desactivar el autómata de seguridad, Mirko retomó la búsqueda.\n" +
		"\x0BCon solo unas pocas habitaciones restantes, finalmente encontró lo que buscaba.\n" +
		"\x0B\x09Mirko: \x01\x08Estas no son nada fáciles de ver por ahí. Espero que no las echen de menos demasiado pronto..\x07\x04\x02\n" +
		"\x0BMirko cogió las dos unidades MSP militares e inmediatamente abandonó la escena.\n" +
		"\x0B\x08>>",
		"\x08Ritapola, dins els Laboratoris MGNT\x07\x05\n" +
		"\x0BDesprés de desactivar l'autòmata de seguretat, en Mirko va reiniciar la cerca.\n" +
		"\x0BQuan només quedaven quatre habitacions comptades, finalment va trobar el que buscava.\n" +
		"\x0B\x09Mirko: \x01\x08No s'en veuen sovint, d'aquestes. Espero que no les trobin a faltar massa aviat..\x07\x04\x02\n" +
		"\x0BEn Mirko va agafar les dues unitats MSP militars i immediatament va abandonar l'escena.\n" +
		"\x0B\x08>>",
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
		"\x08Ritapola, punto de entrega secreto, 12 de julio de 3589\x07\x05\n" +
		"\x0BComo cada día, Mirko revisó el escondrijo.\n" +
		"\x0B\x09Mirko: \x01\x08\x0D!\x0C\x02\x04\n" +
		"\x0B\x07Mirko cogió la nota y la leyó con afán:\x04\n" +
		"\x0B\x09Jana: \x01\x08No puedo garantizar que sea 100% seguro, pero es lo mejor que he podido hacer. Tendría que ser yo la que hiciera esto, no tú. Ya sé que Joseph es tu hermano, pero me mata tener que quedarme aquí esperando. A veces me quitaría mi MSP yo misma. Lo sé, lo sé, no te preocupes, me contendré.\x03\n" +
		"\x0BNo te excedas demasiado. Y buena suerte ahí fuera.\x02\x05\n" +
		"\x0B\x07Jana finalmente lo había conseguido.\x04\n" +
		"\x0B\x08La habilidad \x09[Switch]\x08 ha sido desbloqueada.\n" +
		"\x0B\x08>>",
		"\x08Ritapola, punt d'entrega secret, 12 de juliol de 3589\x07\x05\n" +
		"\x0BCom cada dia, en Mirko va revisar l'amagatall.\n" +
		"\x0B\x09Mirko: \x01\x08\x0D!\x0C\x02\x04\n" +
		"\x0B\x07En Mirko va treure la nota i va llegir ansiosament:\x04\n" +
		"\x0B\x09Jana: \x01\x08No puc garantitzar que sigui 100% segur, però és el millor que he pogut fer. Hauria de ser jo qui fes això, no tu. Ja sé que en Joseph és el teu germà, però em mata haver de quedar-me aquí esperant. A vegades em treuria la MSP jo mateixa. Ho sé, ho sé, no et preocupis, em contindré.\x03\n" +
		"\x0BNo t'excedeixis massa allà fora. Bona sort.\x02\x05\n" +
		"\x0B\x07La Jana finalment ho havia aconseguit.\x04\n" +
		"\x0B\x08L'habilitat \x09[Switch]\x08 ha estat desbloquejada.\n" +
		"\x0B\x08>>",
	),
	lang.NewText(
		"\x08Marunka Machart's cottage, 14 July 3589\x07\x05\n" +
		"\x0BMirko had been waiting in hidding for the last 6 hours. Marunka Machart, the Headsteer member suspected of Joseph's disappearance, had left the place an hour ago.\x04\n" +
		"\x0BIt was time to find what was inside that place.\x04\n" +
		"\x0BMirko approached silently and tried to sneak, but the attempt for stealth was futile. As soon as he jumped the fence, a guard automaton popped up from nowhere.\n" +
		"\x0BDisastrously for Mirko, the guard didn't go right away after him, but retreated to report the breach first.\n" +
		"\x0BAs soon as alarms went off, it turned again towards Mirko.\n" +
		"\x0B\x09Mirko: \x01\x08Tsk.\x02\x03\x07\n" +
		"\x0B\x08>>",
		"\x08Chalet rural de Marunka Machart, 14 de julio de 3589\x07\x05\n" +
		"\x0BMirko se había estado escondiendo durante las últimas 6 horas. Marunka Machart, la integrante de Cabecera sospechosa de la desaparición de Joseph, había salido una hora antes.\x04\n" +
		"\x0BEra el momento de ver qué pasaba ahí dentro.\x04\n" +
		"\x0BMirko se acercó sigilosamente, pero sus esfuerzos fueron en baladí. Tan pronto saltó la verja, un autómata apareció de la nada.\n" +
		"\x0BPor si faltaba más el autómata no fue a perseguirlo directamente, sinó que primero se retiró a comunicar la intrusión.\n" +
		"\x0BEn cuanto las alarmas sonaron, volvió a girarse hacia él.\n" +
		"\x0B\x09Mirko: \x01\x08Tsk.\x02\x03\x07\n" +
		"\x0B\x08>>",
		"\x08Xalet rural de Marunka Machart, 14 de juliol de 3589\x07\x05\n" +
		"\x0BEn Mirko s'havia estat amagant durant les últimes 6 hores. Marunka Machart, la integrant de Capçalera sospitosa de la desaparició d'en Joseph, havia sortit una hora abans.\x04\n" +
		"\x0BEra el moment de veure què passava allà dins.\x04\n" +
		"\x0BEn Mirko es va apropar sigil·losament, però la seva cautela va ser inútil. Tan aviat va saltar la tanca, un autòmata va aparèixer del no res.\n" +
		"\x0BPer a major desgràcia encara, l'autòmata no va començar a perseguir-lo directament, sinò que es va retirar a comunicar la intrusió primer.\n" +
		"\x0BQuan les alarmes estaven sonant, va tornar a girar-se cap a ell.\n" +
		"\x0B\x09Mirko: \x01\x08Tsk.\x02\x03\x07\n" +
		"\x0B\x08>>",
	),
	lang.NewText(
		"\x08Marunka Machart's cottage, 14 July 3589\x07\x05\n" +
		"\x0BMirko really struggled to disable the guard automaton, so there was no time to rest. Security would arrive soon.\x04\n" +
		"\x0BThe main floor looked clear, so Mirko went directly to the basement. There, a high-security door was preventing access.\n" +
		"\x0B\x09Mirko: \x01\x08Damn, this is \x0Ereally going to be\x0C a close call...\x02\x07\n" +
		"\x09Mirko: \x01\x08\"Don't push it too hard\"..? Sorry, Jana.\x02\x07\n" +
		"\x0B\x08>>",
		"\x08Chalet rural de Marunka Machart, 14 de julio de 3589\x07\x05\n" +
		"\x0BA Mirko le costó lo suyo desactivar el autómata y ya no quedaba tiempo que perder. Seguridad llegaría en cualquier momento.\x04\n" +
		"\x0BEl piso principal parecía vacío, así que Mirko bajó directamente al sótano. Allí, una puerta de alta seguridad impedía el acceso.\n" +
		"\x0B\x09Mirko: \x01\x08Maldita sea, \x0Erealmente esto\x0C va a ser muy justo...\x02\x07\n" +
		"\x09Mirko: \x01\x08\"No te excedas demasiado\"..? Lo siento, Jana.\x02\x07\n" +
		"\x0B\x08>>",
		"\x08Xalet rural de Marunka Machart, 14 de juliol de 3589\x07\x05\n" +
		"\x0BA en Mirko li va costar Déu i ajuda desactivar l'autòmata i ja no hi havia temps a pedre. Seguretat arribaria en qualsevol moment.\x04\n" +
		"\x0BEl pis principal semblava buit, així que en Mirko va baixar directament al soterrani. Allà, una porta d'alta seguretat impedia l'accés.\n" +
		"\x0B\x09Mirko: \x01\x08Merda, \x0Erealment això\x0C serà molt just...\x02\x07\n" +
		"\x09Mirko: \x01\x08\"No t'excedeixis massa\"..? Ho sento, Jana.\x02\x07\n" +
		"\x0B\x08>>",
	),
	lang.NewText(
		"\x08Marunka Machart's basement, 14 July 3589\x07\x05\n" +
		"\x0BThe interior was spacious. A man was sitting on a couch, reading a book. He immediately got up when he saw Mirko.\n" +
		"\x0B\x09Man: \x01\x07Who..? \x0EWho are you?\x0C\x02\x04\n" +
		"\x0BBefore Mirko could respond, a woman came running from another room.\n" +
		"\x0B\x09Woman: \x01\x07Did you come to rescue us?\x02\n" +
		"\x0BHer eyes told him everything. High-profile members of Headsteer abusing MSP units to turn people into slaves.. was no longer just a rumor.\n" +
		"\x0B\x09Man: \x01\x07-Kayla, don't-\x02\x04\n" +
		"\x09Mirko: \x01\x07Something like that. But we don't have time. Is Joseph Tatar here?\x02\x04\n" +
		"\x09Kayla: \x01\x07Joseph was taken somewhere else two weeks ago.\x02\n" +
		"\x0BA third man came out of the hallway.\x05 As much as he wanted to find his brother, Mirko had to free them first.\n" +
		"\x0B\x08>>",
		"\x08Sótano del chalet de Marunka Machart, 14 de julio de 3589\x07\x05\n" +
		"\x0BEl interior era espacioso. Un hombre se sentaba en un sofá leyendo un libro. Al ver a Mirko, se levantó de inmediato.\n" +
		"\x0B\x09Hombre: \x01\x07Quién..? \x0EQuién eres?\x0C\x02\x04\n" +
		"\x0BAntes de que Mirko pudiera responder, una mujer llegó corriendo desde otra habitación.\n" +
		"\x0B\x09Mujer: \x01\x07Has venido a rescatarnos?\x02\n" +
		"\x0BSus ojos lo decían todo. Lo de los altos cargos de Cabecera abusando las unidades MSP para esclavizar a personas.. ya no era solo un rumor.\n" +
		"\x0B\x09Hombre: \x01\x07-Kayla, no-\x02\x04\n" +
		"\x09Mirko: \x01\x07Algo así. Pero no hay tiempo. Está Joseph Tatar aquí?\x02\x04\n" +
		"\x09Kayla: \x01\x07Se lo llevaron a alguna otra parte hace dos semanas.\x02\n" +
		"\x0BUn tercer individuo salió del pasillo.\x05 Por más que quisiera encontrar a su hermano, tenía que ayudarlos a ellos primero.\n" +
		"\x0B\x08>>",
		"\x08Soterrani del xalet de Marunka Machart, 14 de juliol de 3589\x07\x05\n" +
		"\x0BL'interior era espaiós. Un home estava assegut al sofà llegint un llibre. Al veure en Mirko, l'home es va aixecar de cop.\n" +
		"\x0B\x09Home: \x01\x07Qui..? \x0EQui ets?\x0C\x02\x04\n" +
		"\x0BAbans que en Mirko pogués contestar, una dona va arribar corrent des d'una altra sala.\n" +
		"\x0B\x09Dona: \x01\x07Has vingut a rescatar-nos?\x02\n" +
		"\x0BEls seus ulls ho deien tot. Alts càrrecs de Capçalera abusant les unitats MSP para esclavitzar a persones.. ja no era només un rumor.\n" +
		"\x0B\x09Home: \x01\x07-Kayla, no-\x02\x04\n" +
		"\x09Mirko: \x01\x07Més o menys. Però no tenim temps. És aquí en Joseph Tatar?\x02\x04\n" +
		"\x09Kayla: \x01\x07Se'l van emportar a algun altre lloc fa dues setmanes.\x02\n" +
		"\x0BUn tercer individu va sortir del passadís.\x05 Per més que volgués trobar al seu germà, havia d'ajudar-los a ells primer.\n" +
		"\x0B\x08>>",
	),
}
