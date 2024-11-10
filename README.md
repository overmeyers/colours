# colours

![example swatches in a terminal window](/swatches.png)

A dark colour pallette that attempts to keep the colour ratio to at least a WCAG 2.1[^1] AA rating of 4.5:1. For the standard colours, the ratio is calculated against the standard black as its base. For the bright variants, the ratio is calculated against the bright black as its base.

| swatch                                             | name             | rgb       | hsl                   | ratio to base |
|:--------------------------------------------------:|------------------|-----------|-----------------------|--------------:|
| ![#1a1e24](https://fakeimg.pl/40x40/1a1e24/1a1e24) | standard black   | `#1a1e24` | `hsl(216, 16%, 12%)`  | 1:1           |
| ![#e58792](https://fakeimg.pl/40x40/e58792/e58792) | standard red     | `#e58792` | `hsl(353, 64%, 71%)`  | 6.46:1        |
| ![#94bf55](https://fakeimg.pl/40x40/94bf55/94bf55) | standard green   | `#94bf55` |	`hsl(84, 45%, 54%)`   | 7.84:1        |
| ![#ffdf74](https://fakeimg.pl/40x40/ffdf74/ffdf74) | standard yellow  | `#ffdf74` | `hsl(46, 100%, 73%)`  | 12.80:1       |
| ![#51afff](https://fakeimg.pl/40x40/51afff/51afff) | standard blue    | `#51afff` | `hsl(208, 100%, 66%)` | 7.11:1        |
| ![#ae92ff](https://fakeimg.pl/40x40/ae92ff/ae92ff) | standard magenta | `#ae92ff` | `hsl(255, 100%, 79%)` | 6.66:1        |
| ![#9bdad3](https://fakeimg.pl/40x40/9bdad3/9bdad3) | standard cyan    | `#9bdad3` | `hsl(173, 46%, 73%)`  | 10.65:1       |
| ![#f1f1f1](https://fakeimg.pl/40x40/f1f1f1/f1f1f1) | standard white   | `#f1f1f1` | `hsl(0, 0%, 95%)`     | 14.81         |
| ![#3d4654](https://fakeimg.pl/40x40/3d4654/3d4654) | bright black     | `#3d4654` | `hsl(217, 16%, 28%)`  | 1:1           |
| ![#ff96a3](https://fakeimg.pl/40x40/ff96a3/ff96a3) | bright red       | `#ff96a3` | `hsl(353, 100%, 79%)` | 4.61:1        |
| ![#b9ed6b](https://fakeimg.pl/40x40/b9ed6b/b9ed6b) | bright green     | `#b9ed6b` |	`hsl(84, 78%, 67%)`   | 6.99:1        |
| ![#fff674](https://fakeimg.pl/40x40/fff674/fff674) | bright yellow    | `#fff674` | `hsl(56, 100%, 73%)`  | 8.48:1        |
| ![#7ad0ff](https://fakeimg.pl/40x40/7ad0ff/7ad0ff) | bright blue      | `#7ad0ff` | `hsl(201, 100%, 74%)` | 5.58:1        |
| ![#bcbaff](https://fakeimg.pl/40x40/bcbaff/bcbaff) | bright magenta   | `#bcbaff` | `hsl(242, 100%, 86%)` | 5.27:1        |
| ![#acf2ea](https://fakeimg.pl/40x40/acf2ea/acf2ea) | bright cyan      | `#acf2ea` | `hsl(173, 73%, 81%)`  | 7.56:1        |
| ![#fdfdfd](https://fakeimg.pl/40x40/fdfdfd/fdfdfd) | bright white     | `#fdfdfd` | `hsl(0, 0%, 99%)`     | 9.37          |


[^1]: 1.4.3 Contrast (Minimum), [“How to Meet WCAG (Quickref Reference)”](https://www.w3.org/WAI/WCAG22/quickref/?versions=2.1#contrast-minimum)
