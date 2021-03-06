package command

import (
	"flag"
)

type Cmd struct {
	ImgPath string
	Help bool
	Threshold int
	Filename string
	TmpImgName	string
	AscWidth int
	AscHeight int
	Ykr	bool	// true -> 使用AsciiYukkuri处理 ,false -> 使用Ascii
}


func ParseCmd() Cmd{
	cmd := Cmd{}
	flag.StringVar(&cmd.ImgPath ,"f" ,"" ,"specify an image file to transcode")
	flag.BoolVar(&cmd.Help ,"h" ,false ,"help")
	flag.IntVar(&cmd.Threshold ,"t" ,140 ,"set threshold for image grey processing")
	flag.StringVar(&cmd.Filename ,"n" ,"tmp.txt" ,"set output file name")
	flag.StringVar(&cmd.TmpImgName ,"m" ,"" ,"save temporary image")
	flag.IntVar(&cmd.AscHeight ,"H" ,-1 ,"set ascii max height")
	flag.IntVar(&cmd.AscWidth ,"W" ,-1 ,"set ascii max width")
	flag.BoolVar(&cmd.Ykr ,"y" ,false ,"make the output ascii more like yukkuri style,but more likely to be unrecognizable." )
	flag.Parse()
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	return cmd
}

