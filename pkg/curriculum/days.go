package curriculum
var days = []Day{
	{
		Number: 1,
		Week:   "Week 1: Basic Movement",
		Title:  "hjkl — Move Without Arrow Keys",
		Summary: "Vim keeps your fingers on the home row: h moves left, l moves right, " +
			"j moves down (think: j looks like a hook dropping down), k moves up. " +
			"No mouse, no arrow keys — just h j k l.",
		Challenges: []Challenge{
			{
				Title: "First Steps",
				Instructions: "h moves the cursor one cell left, l moves it one cell right, " +
					"j moves it one cell down, k moves it one cell up. There's no diagonal " +
					"key — to move diagonally you just press two of them, like j then l. " +
					"Reach the G using only h j k l.",
				Tip: "Tip: h j k l sit right under your fingers on the home row (no need to move " +
					"your hand to an arrow-key cluster) — that's the whole reason Vim uses them.",
				NewKeys: []string{"h", "j", "k", "l"},
				Start: []string{
					".......",
					".......",
					".......",
					"....G..",
					".......",
					".......",
					".......",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{3, 4},
				Par:         7,
			},
			{
				Title: "Bigger Maze",
				Instructions: "Same four keys, longer trip: chain together presses of h j k l — " +
					"in any order — to cross the whole grid and land exactly on the G.",
				Tip: "Tip: you can prefix a motion with a number to repeat it, e.g. 6j moves down " +
					"6 cells in a single keystroke instead of pressing j six times — try it here.",
				NewKeys: []string{"h", "j", "k", "l"},
				Start: []string{
					"..........",
					"..........",
					"..........",
					"..........",
					"..........",
					"..........",
					".......G..",
					"..........",
					"..........",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{6, 7},
				Par:         13,
			},
			{
				Title: "Zigzag Descent",
				Instructions: "The grid is bigger again — this one takes more presses than either " +
					"of the first two combined. Cross it one cell at a time with h j k l to " +
					"land on the G.",
				Tip: "Tip: there's no wrong order to mix j and l here — go all the way down first, " +
					"then across, or interleave them, the destination is the same either way.",
				Start: []string{
					"............",
					"............",
					"............",
					"............",
					"............",
					"............",
					"............",
					"..........G.",
					"............",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{7, 10},
				Par:         17,
			},
			{
				Title: "Up And Over",
				Instructions: "So far you've only needed j and l. This time the G sits above and " +
					"to the left of your start — you'll need k and h instead, moving up and " +
					"left back toward the corner.",
				Tip: "Tip: h and l share the same row of the keyboard as j and k — once all " +
					"four feel natural, you never have to look down at your hands to move.",
				Start: []string{
					"........",
					".G......",
					"........",
					"........",
					"........",
					"........",
					"........",
					"........",
				},
				CursorStart: Pos{6, 6},
				Kind:        KindGoal,
				GoalPos:     Pos{1, 1},
				Par:         10,
			},
			{
				Title: "Full Traverse",
				Instructions: "The longest hjkl trip yet: start at the bottom-left corner and " +
					"climb all the way to the G in the top-right corner, crossing the full " +
					"height and width of the grid.",
				Tip: "Tip: this is exactly the kind of long, repetitive trip a count prefix " +
					"was built for — 9k then 13l gets you there in two keystrokes total instead " +
					"of twenty-two.",
				Start: []string{
					".............G",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
					"..............",
				},
				CursorStart: Pos{9, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 13},
				Par:         22,
			},
		},
	},
	{
		Number: 2,
		Week:   "Week 1: Basic Movement",
		Title:  "Word Motions — w, e, b",
		Summary: "Moving one letter at a time is slow. w jumps to the start of the next word, " +
			"e jumps to the end of the current/next word, b jumps back to the start of the " +
			"previous word. This is how fast Vim users cross a line in a few keystrokes.",
		Challenges: []Challenge{
			{
				Title: "Jump Forward",
				Instructions: "w moves the cursor to the very first character of the next word — " +
					"punctuation and symbols like $ count as words of their own. This line has " +
					"plenty of words to cross — press w repeatedly to land exactly on the $ at " +
					"the end.",
				Tip: "Tip: try it mentally on \"one two three\" — from \"one\", a single w lands " +
					"you on \"two\", another w lands you on \"three\".",
				NewKeys:     []string{"w"},
				Start:       []string{"jump over several small words to reach the $ sign"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 43},
				Par:         8,
			},
			{
				Title: "Jump Backward",
				Instructions: "b moves the cursor back to the start of a word. If the cursor " +
					"isn't already sitting on a word's first letter, the first b snaps back to " +
					"the start of the word you're currently inside — not the previous word. " +
					"Starting at the end of the line, use b to walk all the way back onto \"over\".",
				Tip: "Tip: this trips a lot of people up — from the middle of \"bank\", the first " +
					"b lands on the \"b\" of \"bank\" itself, and only the next b moves further back.",
				NewKeys:     []string{"b"},
				Start:       []string{"quick brown fox jumps over lazy dog near river bank"},
				CursorStart: Pos{0, 50},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 22},
				Par:         6,
			},
			{
				Title: "Land On the Last Letter",
				Instructions: "e moves the cursor to the END of the current or next word, instead " +
					"of the start like w does. Use e to hop word by word all the way to the last " +
					"letter of \"field\", the final word below.",
				Tip: "Tip: w vs e is the key distinction to remember — w lands on a word's " +
					"first letter, e lands on its last. Whichever is closer costs fewer presses.",
				NewKeys:     []string{"e"},
				Start:       []string{"reach the distant target across this wide open field"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 51},
				Par:         9,
			},
			{
				Title: "Word Hopscotch",
				Instructions: "w and e work exactly the same whether you start at the beginning " +
					"of a line or somewhere in the middle. Starting from \"four\" below, hop " +
					"forward word by word until you land on \"twelve\".",
				Tip: "Tip: word motions never care where on the line you started — only where " +
					"the word boundaries are from that point on.",
				Start:       []string{"one two three four five six seven eight nine ten eleven twelve"},
				CursorStart: Pos{0, 14},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 56},
				Par:         8,
			},
			{
				Title: "Backtrack Marathon",
				Instructions: "The longest word-motion trip yet. Starting at the very end of the " +
					"line (inside the last word), use b to walk all the way back to \"market\", " +
					"the very first word.",
				Tip: "Tip: remember the quirk from \"Jump Backward\" — your first b here snaps " +
					"to the start of the word you're already inside, not the one before it.",
				Start:       []string{"market square filled with fresh apples ripe pears sweet cherries golden plums"},
				CursorStart: Pos{0, 76},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 0},
				Par:         12,
			},
		},
	},
	{
		Number: 3,
		Week:   "Week 1: Basic Movement & Modes",
		Title:  "Insert Mode — i, A, o, O",
		Summary: "Normal mode is for moving and editing; Insert mode is for typing. " +
			"i enters Insert mode before the cursor, A jumps to end-of-line and enters " +
			"Insert mode, o opens a new line below and enters Insert mode, O opens one " +
			"above. Esc always returns you to Normal mode — you'll press it constantly.",
		Challenges: []Challenge{
			{
				Title: "Type Something",
				Instructions: "i switches from Normal mode into Insert mode, without moving the " +
					"cursor — anything you type from here on is inserted as regular text, just " +
					"like a normal text editor. Press i, type: welcome to vim insert mode — then " +
					"press Esc to drop back into Normal mode.",
				Tip: "Tip: Esc is one of the most-pressed keys in Vim — a lot of people remap " +
					"Caps Lock to Esc on their keyboard since it's easier to reach.",
				NewKeys:     []string{"i", "esc"},
				Start:       []string{""},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"welcome to vim insert mode"},
				Par:         28,
			},
			{
				Title: "Append At The End",
				Instructions: "A jumps straight to the end of the line and enters Insert mode there " +
					"— a shortcut for pressing $ and then a. No matter where your cursor " +
					"currently sits on the line, use A to append \" world of vim editing\" to " +
					"the end of it.",
				Tip: "Tip: lowercase a inserts right after the cursor's current spot; capital A " +
					"always jumps to end-of-line first, regardless of cursor position.",
				NewKeys:     []string{"A"},
				Start:       []string{"hello"},
				CursorStart: Pos{0, 2},
				Kind:        KindEdit,
				Target:      []string{"hello world of vim editing"},
				Par:         23,
			},
			{
				Title: "Open A New Line",
				Instructions: "o opens a brand-new, empty line directly below the cursor's current " +
					"line and drops you into Insert mode on it — handy for adding a line without " +
					"first navigating to end-of-line. Use it to insert a line reading: second " +
					"line goes right here — between the two lines below.",
				Tip: "Tip: capital O does the same thing but opens the new line ABOVE the " +
					"cursor instead of below — you'll use it in the very next challenge.",
				NewKeys:     []string{"o"},
				Start:       []string{"first line", "third line"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"first line", "second line goes right here", "third line"},
				Par:         29,
			},
			{
				Title: "Insert Above",
				Instructions: "O (capital) opens a new line ABOVE the cursor's current line, then " +
					"drops you into Insert mode on it — the mirror image of lowercase o. This " +
					"file is missing its opening line; use O to add \"first line\" above what's " +
					"already here.",
				Tip: "Tip: think of the a/A and i/I family — lowercase acts relative to the " +
					"cursor, capital jumps to an edge first. o/O follows the same before/after " +
					"pattern, just one line up or down instead of left or right.",
				NewKeys:     []string{"O"},
				Start:       []string{"second line", "third line"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"first line", "second line", "third line"},
				Par:         12,
			},
			{
				Title: "Sandwich The Line",
				Instructions: "Now combine all four: O to add a line above, A to append onto the " +
					"existing line, and o to add a line below — turning one lonely line into a " +
					"full three-course meal. Build: appetizer / main course dish / dessert.",
				Tip: "Tip: after O...Esc you land back on the line you just typed, not the " +
					"original one — you'll need to move down before appending to it.",
				Start:       []string{"main course"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"appetizer", "main course dish", "dessert"},
				Par:         28,
			},
		},
	},
	{
		Number: 4,
		Week:   "Week 1: Basic Movement",
		Title:  "gg, G, and x",
		Summary: "gg jumps to the very first line, G jumps to the very last line — instant " +
			"travel across a whole file, and a count prefix like 7gg jumps straight to line " +
			"7. x deletes the single character under the cursor, your first real edit command.",
		Challenges: []Challenge{
			{
				Title: "Jump To The Top",
				Instructions: "gg (tap g twice) jumps straight to line 1 of the file, no matter " +
					"how far down you are — instant travel instead of holding k. Use it to get " +
					"from the bottom of this list back to the top.",
				Tip:         "Tip: gg accepts a count too — 5gg jumps straight to line 5, exactly.",
				NewKeys:     []string{"gg"},
				Start:       []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
				CursorStart: Pos{9, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{0, 0},
				Par:         2,
			},
			{
				Title: "Jump To The Bottom",
				Instructions: "G jumps straight to the very last line of the file — the mirror " +
					"image of gg. Use it to go from the top all the way to the bottom in a " +
					"single keystroke.",
				Tip: "Tip: like gg, G takes a count — 3G jumps to exactly line 3, same as " +
					"typing :3 and pressing enter in command mode.",
				NewKeys:     []string{"G"},
				Start:       []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{9, 0},
				Par:         1,
			},
			{
				Title: "Delete The Typo",
				Instructions: "x deletes exactly one character — whichever one is currently under " +
					"the cursor — making it your fastest fix for a stray typo. All three lines " +
					"below start with a stray extra letter — delete each one so the lines read " +
					"hello / world / test.",
				Tip: "Tip: capital X deletes the character BEFORE the cursor instead of under " +
					"it — reach for X when you've typed one character too many.",
				NewKeys:     []string{"x"},
				Start:       []string{"xhello", "xworld", "xtest"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"hello", "world", "test"},
				Par:         5,
			},
			{
				Title: "Count Your Jumps",
				Instructions: "A number before gg or G jumps straight to that exact line — no " +
					"counting keystrokes with j needed. Use a single count-prefixed jump to land " +
					"on \"seven\", the 7th line of this list.",
				Tip: "Tip: this is the same trick as 5j from Day 1's tip, applied to gg/G — " +
					"any motion that repeats can usually take a count prefix.",
				Start:       []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
				CursorStart: Pos{0, 0},
				Kind:        KindGoal,
				GoalPos:     Pos{6, 0},
				Par:         3,
			},
			{
				Title: "Typo Hunt",
				Instructions: "Three lines in this longer file each have a stray letter tacked " +
					"onto the front. Use gg, G, j, and k to travel between them and x to delete " +
					"each one — you decide the fastest route.",
				Tip: "Tip: G to jump straight to the bottom typo, then work your way up with k, " +
					"is often faster than counting j's down from the top.",
				Start: []string{
					"first", "xsecond", "third", "fourth", "xfifth", "sixth", "seventh", "xeighth",
				},
				CursorStart: Pos{3, 0},
				Kind:        KindEdit,
				Target: []string{
					"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth",
				},
				Par: 9,
			},
		},
	},
	{
		Number: 5,
		Week:   "Week 1: Basic Movement & Modes",
		Title:  "dd, yy, p, P — Delete, Yank, Paste",
		Summary: "dd deletes (and stores) the whole current line. yy yanks (copies) it " +
			"without deleting. p pastes after the current line, P pastes before it. " +
			"Together these are how you move and duplicate lines without ever touching a mouse.",
		Challenges: []Challenge{
			{
				Title: "Delete The Line",
				Instructions: "dd deletes the entire current line in one shot — wherever the " +
					"cursor sits on that line, the whole line goes, and it's saved so you could " +
					"paste it elsewhere. Two lines below don't belong — navigate to each and " +
					"remove it.",
				Tip: "Tip: 3dd deletes 3 lines starting from the cursor — operators like d " +
					"combine with counts exactly like motions do.",
				NewKeys:     []string{"dd"},
				Start:       []string{"keep1", "BAD1", "keep2", "keep3", "BAD2", "keep4"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"keep1", "keep2", "keep3", "keep4"},
				Par:         7,
			},
			{
				Title: "Duplicate Below",
				Instructions: "yy yanks (copies) the current line without deleting it, and p pastes " +
					"whatever was last yanked or deleted directly below the cursor's line — " +
					"pressing p again pastes the SAME copy a second time. Yank the first line, " +
					"then paste it twice so it appears three times total.",
				Tip: "Tip: think of y as Vim's \"copy\" and d as its \"cut\" — both fill the same " +
					"register, and p is always the paste that follows either one, as many times " +
					"as you press it.",
				NewKeys:     []string{"yy", "p"},
				Start:       []string{"template", "other1", "other2"},
				CursorStart: Pos{1, 0},
				Kind:        KindEdit,
				Target:      []string{"template", "template", "template", "other1", "other2"},
				Par:         5,
			},
			{
				Title: "Paste Above",
				Instructions: "P (capital) pastes above the cursor's line instead of below it — " +
					"the only difference from lowercase p. Yank \"target\" below, then press P " +
					"twice to stack two extra copies above the original.",
				Tip: "Tip: lowercase p = paste after/below, capital P = paste before/above. " +
					"This same after-vs-before pattern (a/A, i/I, p/P) repeats all over Vim.",
				NewKeys:     []string{"P"},
				Start:       []string{"a", "b", "target", "c"},
				CursorStart: Pos{2, 0},
				Kind:        KindEdit,
				Target:      []string{"a", "b", "target", "target", "target", "c"},
				Par:         4,
			},
			{
				Title: "Relocate The Line",
				Instructions: "\"four\" is stuck in the wrong spot near the top — combine dd to " +
					"lift it out, G to jump to the very last line, and p to drop it back in at " +
					"the end where it belongs.",
				Tip: "Tip: dd doesn't just delete — it stores the line exactly like yy would, " +
					"ready for p or P whenever you need it.",
				Start:       []string{"one", "four", "two", "three"},
				CursorStart: Pos{1, 0},
				Kind:        KindEdit,
				Target:      []string{"one", "two", "three", "four"},
				Par:         4,
			},
			{
				Title: "Double Swap",
				Instructions: "This list has two separate pairs of lines swapped out of order. " +
					"Find both mistakes and fix each one using dd and p to lift a line out and " +
					"drop it back in the right place.",
				Tip: "Tip: fix them one at a time, top to bottom — a dd/p pair right after each " +
					"other is usually enough to swap two neighboring lines.",
				Start:       []string{"apple", "cherry", "banana", "date", "fig", "honey", "grape"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"apple", "banana", "cherry", "date", "fig", "grape", "honey"},
				Par:         9,
			},
		},
	},
	{
		Number: 6,
		Week:   "Week 2: More Operators",
		Title:  "dw, cw, ciw, diw — Operators Meet Word Motions",
		Summary: "Operators (d, c, y) and motions (w) combine freely — you're not limited " +
			"to the dd/yy pairs from Day 5. dw deletes a word forward. cw changes one, but " +
			"has a real Vim quirk: it stops at the end of the current word instead of also " +
			"eating the space after it. ciw and diw sidestep that entirely by targeting the " +
			"whole word around the cursor, wherever inside it you happen to be.",
		Challenges: []Challenge{
			{
				Title: "Delete The Extra Word",
				Instructions: "You already know d deletes and w moves by word — combine them and " +
					"dw deletes from the cursor to the start of the next word in a single motion. " +
					"Two stray words, \"very\" and \"extremely\", need to go from the line below.",
				Tip: "Tip: like dd, this combo takes a count too — d3w deletes three words " +
					"forward in one shot.",
				NewKeys:     []string{"dw"},
				Start:       []string{"quick very extremely fast car today"},
				CursorStart: Pos{0, 6},
				Kind:        KindEdit,
				Target:      []string{"quick fast car today"},
				Par:         4,
			},
			{
				Title: "The cw Quirk",
				Instructions: "c changes text just like d deletes it, but cw has a special " +
					"exception: it only replaces up to the end of the word under the cursor — it " +
					"leaves the trailing space alone, unlike dw. Fix both typos below, \"quikc\" " +
					"and \"borwn\", using cw.",
				Tip: "Tip: this exception exists so cw feels natural for fixing typos — without " +
					"it, retyping a word would also eat the space after it, forcing you to add it " +
					"back yourself.",
				NewKeys:     []string{"cw"},
				Start:       []string{"quikc borwn fox"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"quick brown fox"},
				Par:         17,
			},
			{
				Title: "Change Inner Word",
				Instructions: "ciw changes the whole word your cursor is inside, no matter which " +
					"letter you're on — unlike cw, you don't need to be sitting on the word's " +
					"first letter. Fix both \"woord\" and \"tiem\" below using ciw, starting from " +
					"the middle of each.",
				Tip: "Tip: iw also works with d and y (diw, yiw) — because it grabs just the " +
					"word itself, it's more predictable than w-based motions when the cursor " +
					"isn't sitting on a word's edge.",
				NewKeys:     []string{"ciw"},
				Start:       []string{"fix the woord and tiem now"},
				CursorStart: Pos{0, 10},
				Kind:        KindEdit,
				Target:      []string{"fix the word and time now"},
				Par:         18,
			},
			{
				Title: "The diw Gap",
				Instructions: "diw deletes just the inner word, same target as ciw but without " +
					"dropping into Insert mode afterward — and it leaves the surrounding spaces " +
					"exactly where they were, unlike daw which you'll meet tomorrow. Delete the " +
					"second \"there\" below with diw and watch the double space it leaves behind.",
				Tip: "Tip: that leftover double space is exactly the problem daw was built to " +
					"solve — keep this one in mind for Day 7.",
				NewKeys:     []string{"diw"},
				Start:       []string{"hello there there friend"},
				CursorStart: Pos{0, 14},
				Kind:        KindEdit,
				Target:      []string{"hello there  friend"},
				Par:         3,
			},
			{
				Title: "Word Surgery",
				Instructions: "Three lines, three different problems: an extra word to strip out, " +
					"a typo sitting at a word's first letter, and a typo buried in the middle of " +
					"a word. Pick the right tool for each — dw, cw, or ciw.",
				Tip: "Tip: if the cursor is already on the word's first letter, cw is enough — " +
					"only reach for ciw when you're starting from somewhere in the middle.",
				Start: []string{
					"remove very this word here",
					"please fix kwick now",
					"clean woord please",
				},
				CursorStart: Pos{0, 7},
				Kind:        KindEdit,
				Target: []string{
					"remove this word here",
					"please fix quick now",
					"clean word please",
				},
				Par: 22,
			},
		},
	},
	{
		Number: 7,
		Week:   "Week 2: More Operators",
		Title:  "daw, D, C — Around Word & End-of-Line Shortcuts",
		Summary: "iw grabs just the word itself; aw grabs the word plus the whitespace " +
			"around it, so daw cleans up after itself without leaving a double space behind " +
			"like diw did yesterday. D and C are shortcuts for d$ and c$ — delete or change " +
			"everything from the cursor to the end of the line in a single keystroke.",
		Challenges: []Challenge{
			{
				Title: "Delete Around Word",
				Instructions: "aw is like iw but also grabs the whitespace touching the word, so " +
					"daw removes a whole word cleanly without leaving a stray double space behind " +
					"like diw would. Use daw twice to remove both \"cat\" and \"junk\" from the " +
					"line below.",
				Tip: "Tip: yaw and caw follow the same pattern — y/c/d all pair with either the " +
					"i (inner) or a (around) text objects.",
				NewKeys:     []string{"daw"},
				Start:       []string{"please cat dog junk run fast"},
				CursorStart: Pos{0, 7},
				Kind:        KindEdit,
				Target:      []string{"please dog run fast"},
				Par:         7,
			},
			{
				Title: "Delete To End Of Line",
				Instructions: "D deletes from the cursor all the way to the end of the line — a " +
					"shortcut for d$. Both lines below have junk tacked onto the end starting at " +
					"the same column — use D on each to strip it off.",
				Tip: "Tip: D is one of a handful of capital-letter shortcuts for a lowercase " +
					"operator plus $ — C works the same way for c$, as you'll see next.",
				NewKeys:     []string{"D"},
				Start:       []string{"good stuff###delete me here", "nice today###forget this bit"},
				CursorStart: Pos{0, 10},
				Kind:        KindEdit,
				Target:      []string{"good stuff", "nice today"},
				Par:         3,
			},
			{
				Title: "Change To End Of Line",
				Instructions: "C deletes from the cursor to the end of the line and drops you into " +
					"Insert mode there — a shortcut for c$. Use C to fix the typo \"yuo\" on the " +
					"first line, then fix \"konw\" on the second by retyping its whole tail.",
				Tip: "Tip: same after-vs-before family as D/C — think of them as \"the rest of " +
					"the line, gone\" plus whether you land in Insert mode afterward.",
				NewKeys:     []string{"C"},
				Start:       []string{"nice to meet yuo", "great to konw you"},
				CursorStart: Pos{0, 13},
				Kind:        KindEdit,
				Target:      []string{"nice to meet you", "great to know you"},
				Par:         20,
			},
			{
				Title: "Combo Cleanup",
				Instructions: "One line, two fixes: a stray word to remove with daw, then junk " +
					"trailing after \"###\" to strip off with D. Use word motions to travel " +
					"between the two.",
				Tip: "Tip: \"###\" counts as its own word since it's a run of punctuation — w " +
					"lands you right at the start of it, exactly where D needs to strike.",
				Start:       []string{"please extra word here###junk tail data"},
				CursorStart: Pos{0, 7},
				Kind:        KindEdit,
				Target:      []string{"please word here"},
				Par:         6,
			},
			{
				Title: "Capstone",
				Instructions: "Two lines, both needing a different Day 7 tool: an extra word to " +
					"strip with daw on the first, then junk to trim with D on the second. Get " +
					"from one fix to the other using a word motion, not hjkl.",
				Tip: "Tip: b from inside a run of punctuation snaps to the start of that same " +
					"token — handy for lining the cursor up right before a \"###\" marker.",
				Start:       []string{"good words extra stuff here", "trim this###junk data now"},
				CursorStart: Pos{0, 11},
				Kind:        KindEdit,
				Target:      []string{"good words stuff here", "trim this"},
				Par:         6,
			},
		},
	},
	{
		Number: 8,
		Week:   "Week 2: Checkpoint",
		Title:  "Boss Challenge — Everything From Days 1-7",
		Summary: "No new keys today. Five tasks, each solvable only by combining motions, " +
			"operators, and Insert-mode commands from Days 1-7 — no hand-holding, no single " +
			"correct sequence spelled out for you.",
		Challenges: []Challenge{
			{
				Title: "Fix Two Typos",
				Instructions: "Two lines below have typos. Move between lines and fix each word " +
					"— figure out which operator and motion combo gets the job done fastest.",
				Tip: "Tip: you don't have to delete and retype an entire word to fix a typo " +
					"inside it — Day 6 covered exactly this.",
				Start: []string{
					"quikc brown fox",
					"jumps ovver the",
					"lazy dog now",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"quick brown fox",
					"jumps over the",
					"lazy dog now",
				},
				Par: 17,
			},
			{
				Title: "Sort Them Out",
				Instructions: "These three lines are out of alphabetical order. Move the odd one " +
					"out into its correct place — no retyping, just cut and paste.",
				Tip: "Tip: think back to Day 5's dd/p/P — deleting a line doesn't just remove " +
					"it, it stores it for pasting elsewhere.",
				Start:       []string{"banana", "apple", "cherry"},
				CursorStart: Pos{1, 0},
				Kind:        KindEdit,
				Target:      []string{"apple", "banana", "cherry"},
				Par:         5,
			},
			{
				Title: "Finish The Story",
				Instructions: "The first line is missing a word at the end, and the story is " +
					"missing its last line entirely. Fix both using Insert-mode commands from " +
					"Day 3.",
				Tip: "Tip: A jumps to end-of-line before inserting, o opens a fresh line below " +
					"— neither needs any cursor-walking first.",
				Start:       []string{"Hello", "Goodbye friend"},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target:      []string{"Hello world", "Goodbye friend", "See you soon"},
				Par:         23,
			},
			{
				Title: "One Line, Two Fixes",
				Instructions: "Both typos live on this single line. Fix the first one, then get " +
					"to the second using a word motion instead of hjkl — then fix that one too.",
				Tip: "Tip: w still works right where Normal mode left the cursor after Insert " +
					"mode ends — no need to reposition first.",
				Start:       []string{"The quikc brown fox jumps ovver lazy"},
				CursorStart: Pos{0, 4},
				Kind:        KindEdit,
				Target:      []string{"The quick brown fox jumps over lazy"},
				Par:         19,
			},
			{
				Title: "Final Boss",
				Instructions: "One line is in the wrong place, one word has a typo, one word " +
					"shouldn't be there, and one line has junk tacked onto the end. Fix all four " +
					"problems using whatever combination of Day 1-7 tools gets there fastest.",
				Tip: "Tip: reorder first with dd/p, then work top to bottom fixing what's left " +
					"— it's easier to fix typos in a list that's already in the right order.",
				Start: []string{
					"banana",
					"apple",
					"quikc brown fox",
					"jumps extra over",
					"the lazy dog###done",
				},
				CursorStart: Pos{0, 0},
				Kind:        KindEdit,
				Target: []string{
					"apple",
					"banana",
					"quick brown fox",
					"jumps over",
					"the lazy dog",
				},
				Par: 21,
			},
		},
	},
}
