/* Terminal colors (16 first used in escape sequence) */
static const char *colorname[] = {
	/* 8 normal colors */
	"#3D4654",
	"#FA6976",
	"#94BF55",
	"#FFC66D",
	"#54A2CC",
	"#C0ABFF",
	"#66CCAA",
	"#D6C9D5", /* default foreground and cursor colour */

	/* 8 bright colors */
	"#585A6E",
	"#FF937A",
	"#B3E667",
	"#FFDBA1",
	"#69CAFF",
	"#E2D9FF",
	"#7FFFD4",
	"#F0E2EF",

	[255] = 0,

	/* more colors can be added after 255 to use with DefaultXX */
	[256] = "#585A6E", /* default reverse cursor colour */
	[257] = "black", /* default background colour */
};


/*
* Default colors (colorname index)
* foreground, background, cursor, reverse cursor
*/
unsigned int defaultfg = 7;
unsigned int defaultbg = 257;
unsigned int defaultcs = 7;
static unsigned int defaultrcs = 256;
