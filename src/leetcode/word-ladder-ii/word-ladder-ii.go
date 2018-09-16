package main

import "fmt"

// https://leetcode.com/problems/word-ladder-ii/description/

type MyPoint struct {
	word      string
	neighbors []*MyPoint
	level     int
	prevs     []*MyPoint
}

type MyPointGroup struct {
	pts       []*MyPoint
	startWord string
	endWord   string
	startPt   *MyPoint
	endPt     *MyPoint
	paths     [][]string
}

func (pg *MyPointGroup) add(word string) {
	pt := &MyPoint{word, make([]*MyPoint, 0), -1, make([]*MyPoint, 0)}

	if word == pg.startWord {
		pg.startPt = pt
		pt.level = 0
	}

	if word == pg.endWord {
		pg.endPt = pt
	}

	for _, otherPt := range pg.pts {
		pt.add2NeighborIf(otherPt)
	}

	pg.pts = append(pg.pts, pt)
}

func (pg *MyPointGroup) shortestTransform() {
	if pg.endPt == nil {
		return
	}

	ptQueue := make(chan *MyPoint, len(pg.pts))
	ptQueue <- pg.startPt

	for cont := true; cont; {
		select {
		case curPt := <-ptQueue:
			if pg.endPt != nil && (pg.endPt.level < 0 || pg.endPt.level > curPt.level) {
				for _, neighborPt := range curPt.neighbors {
					if neighborPt.level == -1 {
						neighborPt.level = curPt.level + 1
						neighborPt.prevs = append(neighborPt.prevs, curPt)
						ptQueue <- neighborPt
					} else if neighborPt.level == curPt.level+1 {
						neighborPt.prevs = append(neighborPt.prevs, curPt)
					}
				}
			}
		default:
			cont = false
		}
	}

	if pg.endPt != nil {
		path := make([]*MyPoint, pg.endPt.level+1)
		pg.yieldSeqs(pg.endPt, path)
	}
}

func (pg *MyPointGroup) yieldSeqs(curPt *MyPoint, path []*MyPoint) {
	if curPt.level >= 0 {
		path[curPt.level] = curPt
		if curPt == pg.startPt {
			wordPath := make([]string, len(path))
			for i := 0; i < len(path); i++ {
				wordPath[i] = path[i].word
			}
			pg.paths = append(pg.paths, wordPath)
		} else {
			for _, pt := range curPt.prevs {
				pg.yieldSeqs(pt, path)
			}
		}
	}
}

func dist(s1 string, s2 string) int {
	cnt := 0
	for idx, ch := range s1 {
		if string(ch) != string(s2[idx]) {
			cnt += 1
		}
	}
	return cnt
}

func (this *MyPoint) add2NeighborIf(that *MyPoint) {
	if dist(this.word, that.word) == 1 {
		this.neighbors = append(this.neighbors, that)
		that.neighbors = append(that.neighbors, this)
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	pg := &MyPointGroup{make([]*MyPoint, 0), beginWord, endWord, nil, nil, make([][]string, 0)}
	pg.add(beginWord)

	for _, word := range wordList {
		pg.add(word)
	}

	pg.shortestTransform()
	return pg.paths
}

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Println(findLadders("nape", "mild", []string{"dose", "ends", "dine", "jars", "prow", "soap", "guns", "hops", "cray", "hove", "ella", "hour", "lens", "jive", "wiry", "earl", "mara", "part", "flue", "putt", "rory", "bull", "york", "ruts", "lily", "vamp", "bask", "peer", "boat", "dens", "lyre", "jets", "wide", "rile", "boos", "down", "path", "onyx", "mows", "toke", "soto", "dork", "nape", "mans", "loin", "jots", "male", "sits", "minn", "sale", "pets", "hugo", "woke", "suds", "rugs", "vole", "warp", "mite", "pews", "lips", "pals", "nigh", "sulk", "vice", "clod", "iowa", "gibe", "shad", "carl", "huns", "coot", "sera", "mils", "rose", "orly", "ford", "void", "time", "eloy", "risk", "veep", "reps", "dolt", "hens", "tray", "melt", "rung", "rich", "saga", "lust", "yews", "rode", "many", "cods", "rape", "last", "tile", "nosy", "take", "nope", "toni", "bank", "jock", "jody", "diss", "nips", "bake", "lima", "wore", "kins", "cult", "hart", "wuss", "tale", "sing", "lake", "bogy", "wigs", "kari", "magi", "bass", "pent", "tost", "fops", "bags", "duns", "will", "tart", "drug", "gale", "mold", "disk", "spay", "hows", "naps", "puss", "gina", "kara", "zorn", "boll", "cams", "boas", "rave", "sets", "lego", "hays", "judy", "chap", "live", "bahs", "ohio", "nibs", "cuts", "pups", "data", "kate", "rump", "hews", "mary", "stow", "fang", "bolt", "rues", "mesh", "mice", "rise", "rant", "dune", "jell", "laws", "jove", "bode", "sung", "nils", "vila", "mode", "hued", "cell", "fies", "swat", "wags", "nate", "wist", "honk", "goth", "told", "oise", "wail", "tels", "sore", "hunk", "mate", "luke", "tore", "bond", "bast", "vows", "ripe", "fond", "benz", "firs", "zeds", "wary", "baas", "wins", "pair", "tags", "cost", "woes", "buns", "lend", "bops", "code", "eddy", "siva", "oops", "toed", "bale", "hutu", "jolt", "rife", "darn", "tape", "bold", "cope", "cake", "wisp", "vats", "wave", "hems", "bill", "cord", "pert", "type", "kroc", "ucla", "albs", "yoko", "silt", "pock", "drub", "puny", "fads", "mull", "pray", "mole", "talc", "east", "slay", "jamb", "mill", "dung", "jack", "lynx", "nome", "leos", "lade", "sana", "tike", "cali", "toge", "pled", "mile", "mass", "leon", "sloe", "lube", "kans", "cory", "burs", "race", "toss", "mild", "tops", "maze", "city", "sadr", "bays", "poet", "volt", "laze", "gold", "zuni", "shea", "gags", "fist", "ping", "pope", "cora", "yaks", "cosy", "foci", "plan", "colo", "hume", "yowl", "craw", "pied", "toga", "lobs", "love", "lode", "duds", "bled", "juts", "gabs", "fink", "rock", "pant", "wipe", "pele", "suez", "nina", "ring", "okra", "warm", "lyle", "gape", "bead", "lead", "jane", "oink", "ware", "zibo", "inns", "mope", "hang", "made", "fobs", "gamy", "fort", "peak", "gill", "dino", "dina", "tier"}))
}
