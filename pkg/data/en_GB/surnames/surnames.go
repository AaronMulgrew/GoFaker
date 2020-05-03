package surnames

import (
	"math/rand"
)

// GenerateSurname simply generates first Names
// and picks the name from a large string array.
func GenerateSurname(seed int64) string {
	var allNames [500]string
	allNames[0] = "Savage"
	allNames[1] = "Winter"
	allNames[2] = "Metcalfe"
	allNames[3] = "Harper"
	allNames[4] = "Burgess"
	allNames[5] = "Bailey"
	allNames[6] = "Potts"
	allNames[7] = "Boyle"
	allNames[8] = "Brown"
	allNames[9] = "Jennings"
	allNames[10] = "Payne"
	allNames[11] = "Day"
	allNames[12] = "Holland"
	allNames[13] = "Higgins"
	allNames[14] = "Rhodes"
	allNames[15] = "Hancock"
	allNames[16] = "Howells"
	allNames[17] = "Fowler"
	allNames[18] = "Sims"
	allNames[19] = "Thomas"
	allNames[20] = "Parker"
	allNames[21] = "Bentley"
	allNames[22] = "Barnett"
	allNames[23] = "Manning"
	allNames[24] = "Collier"
	allNames[25] = "Holloway"
	allNames[26] = "Hartley"
	allNames[27] = "George"
	allNames[28] = "Tomlinson"
	allNames[29] = "Howard"
	allNames[30] = "Long"
	allNames[31] = "Farmer"
	allNames[32] = "Collins"
	allNames[33] = "Rice"
	allNames[34] = "Townsend"
	allNames[35] = "Rees"
	allNames[36] = "Bruce"
	allNames[37] = "Hammond"
	allNames[38] = "Ford"
	allNames[39] = "Tucker"
	allNames[40] = "Wallis"
	allNames[41] = "Hamilton"
	allNames[42] = "Ferguson"
	allNames[43] = "Hooper"
	allNames[44] = "Francis"
	allNames[45] = "Reeves"
	allNames[46] = "Barlow"
	allNames[47] = "Short"
	allNames[48] = "Cunningham"
	allNames[49] = "Hopkins"
	allNames[50] = "Nicholson"
	allNames[51] = "Archer"
	allNames[52] = "Green"
	allNames[53] = "Glover"
	allNames[54] = "Gibson"
	allNames[55] = "Spencer"
	allNames[56] = "Warner"
	allNames[57] = "Webb"
	allNames[58] = "Whitehouse"
	allNames[59] = "Dean"
	allNames[60] = "Griffiths"
	allNames[61] = "Clark"
	allNames[62] = "Hardy"
	allNames[63] = "Iqbal"
	allNames[64] = "Baldwin"
	allNames[66] = "Blake"
	allNames[67] = "Lees"
	allNames[68] = "Harvey"
	allNames[69] = "Clarke"
	allNames[70] = "Daniels"
	allNames[71] = "Browne"
	allNames[72] = "Macdonald"
	allNames[73] = "Kirk"
	allNames[74] = "Khan"
	allNames[75] = "Davidson"
	allNames[76] = "Dale"
	allNames[77] = "Sanders"
	allNames[78] = "Wilkins"
	allNames[79] = "Connor"
	allNames[80] = "Daly"
	allNames[81] = "Lane"
	allNames[82] = "Kennedy"
	allNames[83] = "Bray"
	allNames[84] = "Burrows"
	allNames[85] = "Hayes"
	allNames[86] = "Wyatt"
	allNames[87] = "Gould"
	allNames[88] = "Dyer"
	allNames[89] = "Nash"
	allNames[90] = "Bryan"
	allNames[91] = "Pope"
	allNames[92] = "Fraser"
	allNames[93] = "Steele"
	allNames[94] = "Walsh"
	allNames[95] = "Wade"
	allNames[96] = "Marsden"
	allNames[97] = "Humphries"
	allNames[99] = "Thompson"
	allNames[100] = "Lord"
	allNames[101] = "Coleman"
	allNames[102] = "Jarvis"
	allNames[103] = "Noble"
	allNames[104] = "Williamson"
	allNames[105] = "Carpenter"
	allNames[106] = "Gardner"
	allNames[107] = "Farrell"
	allNames[108] = "Clayton"
	allNames[109] = "Akhtar"
	allNames[110] = "Gallagher"
	allNames[111] = "Skinner"
	allNames[112] = "Birch"
	allNames[113] = "Kay"
	allNames[114] = "Barrett"
	allNames[115] = "Bates"
	allNames[116] = "Lucas"
	allNames[118] = "Chamberlain"
	allNames[119] = "Chapman"
	allNames[120] = "Ryan"
	allNames[121] = "Thorpe"
	allNames[122] = "Lawson"
	allNames[123] = "Howell"
	allNames[124] = "Martin"
	allNames[125] = "Kelly"
	allNames[126] = "Dobson"
	allNames[127] = "Stevens"
	allNames[128] = "Brennan"
	allNames[129] = "Lloyd"
	allNames[130] = "Quinn"
	allNames[131] = "Morton"
	allNames[132] = "Wilson"
	allNames[133] = "Barnes"
	allNames[134] = "Henry"
	allNames[135] = "Smith"
	allNames[136] = "Pritchard"
	allNames[137] = "Phillips"
	allNames[138] = "Dixon"
	allNames[139] = "Sharpe"
	allNames[140] = "Robertson"
	allNames[141] = "White"
	allNames[142] = "Bird"
	allNames[143] = "Abbott"
	allNames[144] = "Kirby"
	allNames[145] = "Hussain"
	allNames[146] = "Barber"
	allNames[147] = "Harris"
	allNames[148] = "Doyle"
	allNames[149] = "Jordan"
	allNames[150] = "Burns"
	allNames[151] = "Hodgson"
	allNames[152] = "Atkins"
	allNames[153] = "Stokes"
	allNames[154] = "Rogers"
	allNames[155] = "Parkes"
	allNames[156] = "Brookes"
	allNames[157] = "Herbert"
	allNames[158] = "Gordon"
	allNames[159] = "Kemp"
	allNames[160] = "Webster"
	allNames[161] = "Sinclair"
	allNames[162] = "Mclean"
	allNames[163] = "Saunders"
	allNames[164] = "Stephens"
	allNames[165] = "Newton"
	allNames[166] = "Potter"
	allNames[167] = "Storey"
	allNames[168] = "Stanley"
	allNames[169] = "Turnbull"
	allNames[170] = "Duncan"
	allNames[171] = "Rose"
	allNames[172] = "Mills"
	allNames[173] = "Sheppard"
	allNames[174] = "Butcher"
	allNames[175] = "Fry"
	allNames[176] = "Ross"
	allNames[177] = "Shepherd"
	allNames[178] = "Goodwin"
	allNames[179] = "Holt"
	allNames[180] = "Haynes"
	allNames[181] = "Cook"
	allNames[182] = "Ward"
	allNames[183] = "Godfrey"
	allNames[184] = "Stone"
	allNames[185] = "Dodd"
	allNames[186] = "Parsons"
	allNames[187] = "Ingram"
	allNames[188] = "Nixon"
	allNames[189] = "Evans"
	allNames[190] = "Hargreaves"
	allNames[191] = "Owen"
	allNames[192] = "Chan"
	allNames[193] = "Connolly"
	allNames[194] = "Charlton"
	allNames[195] = "Middleton"
	allNames[196] = "Hyde"
	allNames[197] = "Patel"
	allNames[198] = "Owens"
	allNames[199] = "Lamb"
	allNames[200] = "Palmer"
	allNames[201] = "Cooper"
	allNames[202] = "Mccarthy"
	allNames[203] = "Black"
	allNames[204] = "Dickinson"
	allNames[205] = "Gilbert"
	allNames[206] = "Leach"
	allNames[207] = "North"
	allNames[208] = "Byrne"
	allNames[209] = "Frost"
	allNames[210] = "Simmons"
	allNames[211] = "Matthews"
	allNames[212] = "Alexander"
	allNames[213] = "Ahmed"
	allNames[214] = "Gibbons"
	allNames[215] = "Stevenson"
	allNames[216] = "Rowley"
	allNames[217] = "Miles"
	allNames[218] = "Hanson"
	allNames[219] = "Bolton"
	allNames[220] = "Craig"
	allNames[221] = "Ali"
	allNames[222] = "Carroll"
	allNames[223] = "Allan"
	allNames[224] = "Sanderson"
	allNames[225] = "Fletcher"
	allNames[226] = "Burton"
	allNames[227] = "Oliver"
	allNames[228] = "Davison"
	allNames[229] = "Douglas"
	allNames[230] = "Field"
	allNames[231] = "Pickering"
	allNames[232] = "Pugh"
	allNames[233] = "Rowe"
	allNames[234] = "Mahmood"
	allNames[235] = "Sykes"
	allNames[236] = "Crawford"
	allNames[237] = "Williams"
	allNames[238] = "Parkin"
	allNames[239] = "Patterson"
	allNames[240] = "Power"
	allNames[241] = "Price"
	allNames[242] = "Murphy"
	allNames[243] = "Hale"
	allNames[244] = "Nicholls"
	allNames[245] = "Hall"
	allNames[246] = "Jones"
	allNames[247] = "Hughes"
	allNames[248] = "Stephenson"
	allNames[249] = "Morley"
	allNames[250] = "Knight"
	allNames[251] = "Kerr"
	allNames[252] = "Heath"
	allNames[253] = "Pollard"
	allNames[254] = "Lowe"
	allNames[256] = "Buckley"
	allNames[257] = "Bond"
	allNames[258] = "Dennis"
	allNames[259] = "Lewis"
	allNames[260] = "Weston"
	allNames[261] = "Joyce"
	allNames[262] = "Reynolds"
	allNames[263] = "Bishop"
	allNames[264] = "Norris"
	allNames[265] = "Barry"
	allNames[266] = "Whittaker"
	allNames[267] = "Carey"
	allNames[268] = "Hill"
	allNames[269] = "Kent"
	allNames[270] = "Ashton"
	allNames[271] = "Wilkinson"
	allNames[272] = "Powell"
	allNames[273] = "Henderson"
	allNames[274] = "Freeman"
	allNames[275] = "Dunn"
	allNames[276] = "Kaur"
	allNames[277] = "French"
	allNames[278] = "Parry"
	allNames[279] = "Walton"
	allNames[280] = "Fisher"
	allNames[281] = "Naylor"
	allNames[282] = "Duffy"
	allNames[283] = "Humphreys"
	allNames[284] = "Randall"
	allNames[285] = "Bevan"
	allNames[286] = "Doherty"
	allNames[287] = "Moore"
	allNames[288] = "Armstrong"
	allNames[289] = "Sullivan"
	allNames[290] = "Swift"
	allNames[291] = "Pearce"
	allNames[292] = "Tyler"
	allNames[293] = "Bradshaw"
	allNames[294] = "Allen"
	allNames[295] = "Mellor"
	allNames[296] = "Whitehead"
	allNames[297] = "Jackson"
	allNames[298] = "Grant"
	allNames[299] = "Fox"
	allNames[300] = "Wright"
	allNames[301] = "Anderson"
	allNames[302] = "Foster"
	allNames[303] = "Gibbs"
	allNames[304] = "Butler"
	allNames[305] = "Jenkins"
	allNames[306] = "John"
	allNames[307] = "Morrison"
	allNames[308] = "Talbot"
	allNames[309] = "Blackburn"
	allNames[310] = "Osborne"
	allNames[311] = "Flynn"
	allNames[312] = "Richards"
	allNames[313] = "Hurst"
	allNames[314] = "Bibi"
	allNames[315] = "Houghton"
	allNames[316] = "Johnson"
	allNames[317] = "Yates"
	allNames[318] = "Mistry"
	allNames[319] = "Donnelly"
	allNames[320] = "Parkinson"
	allNames[321] = "Thomson"
	allNames[322] = "Woods"
	allNames[323] = "Todd"
	allNames[324] = "Dawson"
	allNames[325] = "Hart"
	allNames[326] = "Graham"
	allNames[327] = "Berry"
	allNames[328] = "Willis"
	allNames[329] = "Miah"
	allNames[330] = "Brooks"
	allNames[331] = "Horton"
	allNames[332] = "Riley"
	allNames[333] = "Lambert"
	allNames[334] = "Waters"
	allNames[335] = "Lynch"
	allNames[336] = "Moss"
	allNames[337] = "Slater"
	allNames[338] = "Knowles"
	allNames[339] = "Benson"
	allNames[340] = "Adams"
	allNames[341] = "King"
	allNames[342] = "Davies"
	allNames[343] = "Richardson"
	allNames[344] = "Vincent"
	allNames[345] = "Holmes"
	allNames[346] = "Conway"
	allNames[347] = "Marshall"
	allNames[348] = "Faulkner"
	allNames[349] = "Garner"
	allNames[350] = "Booth"
	allNames[351] = "Harrison"
	allNames[352] = "Campbell"
	allNames[353] = "Cole"
	allNames[354] = "Goddard"
	allNames[355] = "Walters"
	allNames[356] = "Ellis"
	allNames[357] = "Edwards"
	allNames[358] = "Peters"
	allNames[359] = "Atkinson"
	allNames[360] = "Wood"
	allNames[361] = "Briggs"
	allNames[362] = "Elliott"
	allNames[363] = "Chandler"
	allNames[364] = "Hope"
	allNames[365] = "Hunter"
	allNames[366] = "Newman"
	allNames[367] = "Pratt"
	allNames[368] = "Rahman"
	allNames[369] = "Hicks"
	allNames[370] = "Cox"
	allNames[371] = "Reid"
	allNames[372] = "Morris"
	allNames[373] = "Banks"
	allNames[374] = "Myers"
	allNames[375] = "Mitchell"
	allNames[376] = "Davey"
	allNames[377] = "Peacock"
	allNames[378] = "Reed"
	allNames[379] = "Carter"
	allNames[380] = "Miller"
	allNames[381] = "Perkins"
	allNames[382] = "Read"
	allNames[383] = "Hilton"
	allNames[384] = "Moran"
	allNames[385] = "Welch"
	allNames[386] = "Vaughan"
	allNames[387] = "Clements"
	allNames[388] = "Griffin"
	allNames[389] = "Russell"
	allNames[391] = "Hobbs"
	allNames[392] = "Marsh"
	allNames[393] = "Porter"
	allNames[394] = "Gill"
	allNames[395] = "Leonard"
	allNames[396] = "Mckenzie"
	allNames[397] = "Thornton"
	allNames[398] = "Fitzgerald"
	allNames[399] = "Greenwood"
	allNames[400] = "Pearson"
	allNames[401] = "James"
	allNames[402] = "Coles"
	allNames[403] = "Roberts"
	allNames[404] = "Nelson"
	allNames[405] = "Forster"
	allNames[406] = "Gough"
	allNames[407] = "Mann"
	allNames[408] = "Law"
	allNames[409] = "Barker"
	allNames[410] = "Cartwright"
	allNames[411] = "Bradley"
	allNames[412] = "Sharp"
	allNames[413] = "Warren"
	allNames[414] = "Summers"
	allNames[415] = "Little"
	allNames[416] = "Perry"
	allNames[417] = "Fuller"
	allNames[418] = "West"
	allNames[419] = "Mason"
	allNames[420] = "Finch"
	allNames[421] = "Norton"
	allNames[422] = "Burke"
	allNames[423] = "Holden"
	allNames[424] = "Lee"
	allNames[425] = "Smart"
	allNames[426] = "Bull"
	allNames[427] = "Bryant"
	allNames[428] = "Gray"
	allNames[429] = "Watts"
	allNames[430] = "Brady"
	allNames[431] = "Baker"
	allNames[432] = "Barton"
	allNames[433] = "Davis"
	allNames[434] = "Baxter"
	allNames[435] = "Taylor"
	allNames[436] = "Carr"
	allNames[437] = "Wong"
	allNames[438] = "Cameron"
	allNames[439] = "Gardiner"
	allNames[440] = "Hawkins"
	allNames[441] = "Shaw"
	allNames[442] = "Wallace"
	allNames[443] = "Young"
	allNames[444] = "Shah"
	allNames[445] = "Gregory"
	allNames[446] = "Ball"
	allNames[447] = "Norman"
	allNames[448] = "Lawrence"
	allNames[449] = "Bowen"
	allNames[450] = "Wheeler"
	allNames[451] = "Bartlett"
	allNames[452] = "Sutton"
	allNames[453] = "Lyons"
	allNames[454] = "Hutchinson"
	allNames[455] = "Poole"
	allNames[456] = "Cooke"
	allNames[457] = "Franklin"
	allNames[458] = "Howe"
	allNames[459] = "Walker"
	allNames[460] = "Johnston"
	allNames[461] = "Austin"
	allNames[462] = "Chadwick"
	allNames[463] = "Bell"
	allNames[464] = "Wall"
	allNames[465] = "Woodward"
	allNames[466] = "Preston"
	allNames[467] = "Bennett"
	allNames[468] = "Murray"
	allNames[469] = "Begum"
	allNames[470] = "Mcdonald"
	allNames[471] = "Hudson"
	allNames[472] = "Cross"
	allNames[473] = "Singh"
	allNames[474] = "Howarth"
	allNames[475] = "Hewitt"
	allNames[476] = "Curtis"
	allNames[477] = "Harding"
	allNames[478] = "May"
	allNames[479] = "Wells"
	allNames[480] = "Giles"
	allNames[481] = "Watson"
	allNames[482] = "Nolan"
	allNames[483] = "Andrews"
	allNames[484] = "Hayward"
	allNames[485] = "Schofield"
	allNames[486] = "Hunt"
	allNames[487] = "Robson"
	allNames[488] = "Arnold"
	allNames[489] = "Morgan"
	allNames[490] = "Coates"
	allNames[491] = "Page"
	allNames[492] = "Simpson"
	allNames[493] = "Stewart"
	allNames[494] = "Robinson"
	allNames[495] = "Fleming"
	allNames[496] = "Scott"
	allNames[497] = "Chambers"
	allNames[498] = "Turner"
	allNames[499] = "Watkins"

	rand.Seed(seed)

	min := 0
	max := 500
	randIndex := rand.Intn(max-min) + min
	name := allNames[randIndex]
	return name
}
