package text

type pageKey int
const (
	Preamble      pageKey = 0
	ToBeContinued pageKey = 1
	Afterword     pageKey = 2
)

var gameTexts = []string {
	"\x08Preamble\x07\n" +
	"\x0BIn 3589, life was good for the average citizen: housing and food were guaranteed by the state, work was voluntary, streets were clean and most freedoms were respected.\n" +
	"\x0B\x09MSP Units\x07, implanted in the spine of every citizen, allowed \x09Leadership\x07 to maintain the social stability.\n" +
	"\x0BDeveloped back during the Primal Wars, \x09MSP Units\x07 were once the key to end the violent outbursts that plagued the country. Nowadays, these magnetic devices had evolved to allow even greater control over individuals, with certain monitoring functions mandated to remain always active.\n" +
	"\x0BMost people had little reason to complain, but for Mirko... things changed when his brother abandoned the love of his life and his family to go work for Marunka Machart, a member of Leadership.\x03\n" +
	"\x0BIt didn't make any sense. What if the rumors were true? He... had no choice.\n" +
	"\x0B\x08>>",
	"\x0BMirko managed to release two of the slaves, but got caught while trying to disable the MSP unit of the third one.\x04\n" +
	"\x0BSeeing the deployment of public safety units around the zone and the lack of contact from Mirko, Jana started fearing the worst.\x04\n" +
	"\x0BShe was a loose end, and they would surely come for her next. But... she hadn't given up just yet.\x04\n" +
	"\x0B\x08>> to be continued",
	"\x08Afterword\x07\n" +
	"\x0BThanks to Hajime Hoshi, the developer of Ebitengine, for being the true magnet for the Ebitengine community.\x04\n" +
	"\x0BThanks to all the people making cool games and libraries for Ebitengine.\x04 Or just hanging around!\x04\n" +
	"\x0BThanks to you for playing.\x04\n" +
	"\x0B\x08>>",
}
