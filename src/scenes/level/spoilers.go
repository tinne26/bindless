package level

import "github.com/tinne26/bindless/src/lang"
import "github.com/tinne26/bindless/src/ui"

var levelTexts = map[int]*lang.Text {
	0: lang.NewText(
		"\x08MSP Hacking Tutorial ( 1 / 6 )\x07\n" +
		"\x0B\x09Attraction And Repulsion\x07\n" +
		"\x0BElectromagnets of different colors attract each other. Electromagnets of the same color repel each other.\n" +
		"\x0BAll motion in MSPs derives from these fundamental principles.",
		"\x08Tutorial de Jaqueo de MSPs ( 1 / 6 )\x07\n" +
		"\x0B\x09Atracción y Repulsión\x07\n" +
		"\x0BLos electroimanes de colores diferentes se atraen entre ellos. Los electroimanes del mismo color se repelen entre ellos.\n" +
		"\x0BTodo movimiento en los MSPs deriva de estos principios fundamentales.",
		"\x08Tutorial de Pirateig de MSPs ( 1 / 6 )\x07\n" +
		"\x0B\x09Atracció i Repulsió\x07\n" +
		"\x0BEls electroimants de colors diferents s'atrauen entre ells. Els electroimants del mateix color es repel·leixen entre ells.\n" +
		"\x0BTot moviment en els MSPs deriva d'aquests principis fonamentals.",
	),
	1: lang.NewText(
		"\x08MSP Hacking Tutorial ( 2 / 6 )\x07\n" +
		"\x0B\x09Electromagnet Types\x07\n" +
		"\x0BOnly small electromagnets that are floating can move; large electromagnets stay always anchored to a fixed location. \n" +
		"\x0BUnpowered electromagnets don't interact magnetically, but may still act as physical blockers.",
		"\x08MSP Tutorial de Jaqueo de MSPs ( 2 / 6 )\x07\n" +
		"\x0B\x09Tipos de Electroimanes\x07\n" +
		"\x0BSolo los electroimanes pequeños que están flotando pueden moverse; los electroimanes grandes están siempre fijos.\n" +
		"\x0BLos electroimanes inactivos no interaccionan magnéticamente, pero pueden actuar como obstáculos físicos igualmente.",
		"\x08MSP Tutorial de Pirateig de MSPs ( 2 / 6 )\x07\n" +
		"\x0B\x09Tipus d'electroimants\x07\n" +
		"\x0BNomés els electroimants menuts que estan flotant poden moure's; els electroimants grans són sempre fixos.\n" +
		"\x0BEls electroimants inactius no interaccionen magnèticament, però poden actuar com a obstacles físics igualment.",
	),
	2: lang.NewText(
		"\x08MSP Hacking Tutorial ( 3 / 6 )\x07\n" +
		"\x0B\x09The Grid\x07\n" +
		"\x0BElectromagnets don't interact with every other surrounding electromagnet, only those in the same row or column that are close enough.",
		"\x08MSP Tutorial de Jaqueo de MSPs ( 3 / 6 )\x07\n" +
		"\x0B\x09Las Casillas\x07\n" +
		"\x0BLos electroimanes no interactúan con todos los demás elementos a su alrededor, solo aquellos en la misma fila o columna que estén lo suficientemente cerca.",
		"\x08MSP Tutorial de Pirateig de MSPs ( 3 / 6 )\x07\n" +
		"\x0B\x09Les Caselles\x07\n" +
		"\x0BEls electroimants no interactuen amb tots els altres elements al seu voltant, només aquells de la mateixa fila o columna que es trobin prou a prop.",
	),
	3: lang.NewText(
		"\x08MSP Hacking Tutorial ( 4 / 6 )\x07\n" +
		"\x0B\x09Resolution\x07\n" +
		"\x0BTo successfully take control of an MSP you must manage to place an electromagnet on top of a circular field point of the same polarity.\n" + 
		"\x0BIn order to do this, you have to use the abilities on the bottom left part of the screen. You have a limited amount of charges per level.\n" +
		"\x0BSelect the abilities and find where to use them to solve this level.",
		"\x08MSP Tutorial de Jaqueo de MSPs ( 4 / 6 )\x07\n" +
		"\x0B\x09Resolución\x07\n" +
		"\x0BPara tomar el control de un MSP debes conseguir colocar un electroimán sobre cualquier objetivo circular de la misma polaridad.\n" +
		"\x0BPara hacer esto necesitas usar las habilidades de la parte inferior izquierda de la pantalla. El número de cargas por nivel es limitado.\n" +
		"\x0BSelecciona las habilidades y descubre dónde usarlas para resolver este nivel.",
		"\x08MSP Tutorial de Pirateig de MSPs ( 4 / 6 )\x07\n" +
		"\x0B\x09Resol·lució\x07\n" +
		"\x0BPer obtenir el control d'un MSP has d'aconseguir col·locar un electroimant sobre qualsevol objectiu circular de la mateixa polaritat.\n" +
		"\x0BPer fer-ho has d'utilitzar les habilitats de la part inferior esquerra de la pantalla. Tens un nombre limitat de càrregues per nivell.\n" +
		"\x0BSelecciona les habilitats i descobreix on fer-les servir per superar aquest nivell.",
	),
	4: lang.NewText(
		"\x08MSP Hacking Tutorial ( 5 / 6 )\x07\n" +
		"\x0B\x09Circuits And Power\x07\n" +
		"\x0BCircuits can be used to power electromagnets or transfer power between them.\n" + 
		"\x0BUse the abilities on the bottom left part of the screen to experiment with this. Take this opportunity to also improve your understanding on how \x09Dock\x07 and \x09Rewire\x07 operate.\n" +
		"\x0BAs an exercise, try to repeatedly get to the red mark without using the option to recharge abilities.",
		"\x08MSP Tutorial de Jaqueo de MSPs ( 5 / 6 )\x07\n" +
		"\x0B\x09Circuitos y Carga\x07\n" +
		"\x0BLos circuitos se pueden usar para activar los electroimanes o transferir carga entre ellos.\n" + 
		"\x0BUsa las habilidades de la parte inferior izquierda de la pantalla para experimentar con esto.\n" +
		"\x0BAprovecha también para entender mejor cómo funcionan \x09Dock\x07 (acoplar) y \x09Rewire\x07 (reconectar).\n" +
		"\x0BPara practicar, intenta alcanzar la marca roja repetidamente sin usar la opción de recarga de habilidades.",
		"\x08MSP Tutorial de Pirateig de MSPs ( 5 / 6 )\x07\n" +
		"\x0B\x09Circuits i Càrrega\x07\n" +
		"\x0BEls circuits es poden utilitzar per activar els electroimants o transferir càrrega entre ells.\n" + 
		"\x0BUtilitza les habilitats de la part inferior esquerra de la pantalla per experimentar amb això. Aprofita aquesta oportunitat per entendre millor com funcionen \x09Dock\x07 (acoblar) i \x09Rewire\x07 (reconnectar).\n" +
		"\x0BPer practicar, intenta arribar a la marca vermella repetidament sense fer servir l'opció de recàrrega d'habilitats.",
	),
	5: lang.NewText(
		"\x08MSP Hacking Tutorial ( 6 / 6 )\x07\n" +
		"\x0B\x09Synchrony\x07\n" +
		"\x0BElectromagnets always move in synchrony through the tiles. Correspondingly, used abilities aren't activated until the current movement cycle ends.\n" + 
		"\x0BSometimes finding a way forward requires more than the right sequence of actions.\n" +
		"\x0BSolve the level to complete the tutorial. If you fail, click on the bottom right menu icon to reset.",
		"\x08MSP Tutorial de Jaqueo de MSPs ( 6 / 6 )\x07\n" +
		"\x0B\x09Sincronía\x07\n" +
		"\x0BLos electroimanes se mueven en sincronía a través de las casillas. Correspondientemente, las habilidades usadas no se activan hasta que el ciclo de movimiento actual termina.\n" + 
		"\x0BEncontrar las soluciones puede requerir algo más que la secuencia de acciones correcta.\n" +
		"\x0BResuelve el nivel para completar el tutorial. Si fallas, pulsa el icono de la parte inferior derecha de la pantalla para reiniciar.",
		"\x08MSP Tutorial de Pirateig de MSPs ( 6 / 6 )\x07\n" +
		"\x0B\x09Sincronia\x07\n" +
		"\x0BEls electroimants es mouen en sincronia a través de les caselles. Corresponentment, les habilitats usades no s'activen fins que el cicle de moviment actual acaba.\n" + 
		"\x0BTrobar sol·lucions de vegades requereix alguna cosa més que la seqüència d'accions correcta.\n" +
		"\x0BResol el nivell per completar el tutorial. Si falles, clica la icona de la part inferior dreta de la pantalla per reiniciar.",
	),
	6: lang.NewText(
		"Street Cleaning Automaton\n\x08MSP V16.349.12 SHELL",
		"Autómata de Limpieza Pública\n\x08MSP V16.349.12 SHELL",
		"Autòmata de Neteja Pública\n\x08MSP V16.349.12 SHELL",
	),
	7: lang.NewText(
		"Street Cleaning Automaton\n\x08MSP V16.349.12 CORE",
		"Autómata de Limpieza Pública\n\x08MSP V16.349.12 CORE",
		"Autòmata de Neteja Pública\n\x08MSP V16.349.12 CORE",
	),
	8: lang.NewText(
		"BackSafe Door Model W\n\x08MSP V16.410.07 CORE",
		"BackSafe Puerta Modelo W\n\x08MSP V16.410.07 CORE",
		"BackSafe Porta Model W\n\x08MSP V16.410.07 CORE",
	),
	9: lang.NewText(
		"MGNT GentleGuard K2\n\x08MSP V16.388.65 SHELL",
		"MGNT GentleGuard K2\n\x08MSP V16.388.65 SHELL",
		"MGNT GentleGuard K2\n\x08MSP V16.388.65 SHELL",
	),
	10: lang.NewText(
		"MGNT GentleGuard K2\n\x08MSP V16.388.65 CORE",
		"MGNT GentleGuard K2\n\x08MSP V16.388.65 CORE",
		"MGNT GentleGuard K2\n\x08MSP V16.388.65 CORE",
	),
	11: lang.NewText(
		"\x08Modulated MSP Tutorial\x07\n" +
		"\x0B\x09Switch\x07\n" +
		"\x0BThe \x09Switch\x07 ability allows changing the polarity of a small electromagnet. It can be used on any powered small electromagnet, even if it's docked.",
		"\x08Tutorial MSP Modulada\x07\n" +
		"\x0B\x09Switch\x07\n" +
		"\x0BLa habilidad \x09Switch\x07 (conmutar) permite cambiar la polaridad de un electroimán pequeño. Se puede usar sobre cualquier electroimán pequeño activado, incluso si no está flotando.",
		"\x08Tutorial MSP Modulada\x07\n" +
		"\x0B\x09Switch\x07\n" +
		"\x0BL'habilitat \x09Switch\x07 (commutar) permet alternar la polaritat d'un electroimant menut. Es pot utilitzar sobre qualsevol electroimant, fins i tot si no està flotant.",
	),
}

var prevText = lang.NewText("[ Previous ]", "[ Anterior ]", "[ Anterior ]")
var uiOptPrevDisabled = &ui.HChoice { Text: prevText }
var uiOptPrev = &ui.HChoice { Text: prevText }
var uiOptNext = &ui.HChoice {
	Text: lang.NewText("[ Next ]", "[ Siguiente ]", "[ Següent ]"),
}
var uiOptSolve = &ui.HChoice {
	Text: lang.NewText("[ Solve to continue ]", "[ Resuelve para seguir ]", "[ Resol per continuar ]"),
}
var uiOptSkip = &ui.HChoice {
	Text: lang.NewText("[ Skip tutorial ]", "[ Saltar el tutorial ]", "[ Saltar el tutorial ]"),
}
var uiRecharge = &ui.HChoice {
	Text: lang.NewText("[ Recharge abilities ]", "[ Recargar habilidades ]", "[ Recarregar habilitats ]"),
}

var LevelChoices = map[int][]*ui.HChoice {
	0: []*ui.HChoice{ uiOptPrev, uiOptNext },
	1: []*ui.HChoice{ uiOptPrev, uiOptNext },
	2: []*ui.HChoice{ uiOptPrev, uiOptNext },
	3: []*ui.HChoice{ uiOptPrev, uiOptSolve },
	4: []*ui.HChoice{ uiOptPrev, uiRecharge, uiOptNext },
	5: []*ui.HChoice{ uiOptPrev, uiOptSolve },
	11: []*ui.HChoice{ uiOptSkip, uiOptSolve },
}



