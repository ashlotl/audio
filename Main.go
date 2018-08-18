package main
import "github.com/hegedustibor/htgo-tts"
import "dict"

func main() {
	dict.Initi("/home/wurst/go/src/dict/syllables")
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	set:=dict.SetOfKeys()
	for key:=0;key<len(set);key++{
		speech.Speak(set[key])
	}
}