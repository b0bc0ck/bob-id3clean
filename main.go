package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/karrick/godirwalk"
	"gopkg.in/yaml.v2"
)

var P = flag.String("P", "/home/ftpd/glftpd/site/mp3/1969-01-01", "Full scan path")
var C = flag.Bool("C", false, "Cleanup (delete folders)")
var D = flag.Bool("D", false, "Debug mode")

type Config struct {
	Clean   []string
	Keepdir []string
}

func genre(file string) string {
	tag, _ := id3v2.Open(file, id3v2.Options{Parse: true})
	defer tag.Close()
	return tag.Genre()
}
func convgenre(num string) string {
	switch num {
	case "0":
		return "Blues"
	case "1":
		return "Classic Rock"
	case "2":
		return "Country"
	case "3":
		return "Dance"
	case "4":
		return "Disco"
	case "5":
		return "Funk"
	case "6":
		return "Grunge"
	case "7":
		return "Hip-Hop"
	case "8":
		return "Jazz"
	case "9":
		return "Metal"
	case "10":
		return "New Age"
	case "11":
		return "Olides"
	case "12":
		return "Other"
	case "13":
		return "Pop"
	case "14":
		return "R&B"
	case "15":
		return "Rap"
	case "16":
		return "Reggae"
	case "17":
		return "Rock"
	case "18":
		return "Techno"
	case "19":
		return "Industrial"
	case "20":
		return "Alternative"
	case "21":
		return "Ska"
	case "22":
		return "Death Metal"
	case "23":
		return "Pranks"
	case "24":
		return "Soundtrack"
	case "25":
		return "Euro-Techno"
	case "26":
		return "Ambient"
	case "27":
		return "Trip-Hop"
	case "28":
		return "Vocal"
	case "29":
		return "Jazz+Funk"
	case "30":
		return "Fusion"
	case "31":
		return "Trance"
	case "32":
		return "Classical"
	case "33":
		return "Instrumental"
	case "34":
		return "Acid"
	case "35":
		return "House"
	case "36":
		return "Game"
	case "37":
		return "Sound Clip"
	case "38":
		return "Gospel"
	case "39":
		return "Noise"
	case "40":
		return "Alt. Rock"
	case "41":
		return "Bass"
	case "42":
		return "Soul"
	case "43":
		return "Punk"
	case "44":
		return "Space"
	case "45":
		return "Meditative"
	case "46":
		return "Instrumental Pop"
	case "47":
		return "Instrumental Rock"
	case "48":
		return "Ethnic"
	case "49":
		return "Gothic"
	case "50":
		return "Darkwave"
	case "51":
		return "Techno-Industrial"
	case "52":
		return "Electronic"
	case "53":
		return "Pop-Folk"
	case "54":
		return "Eurodance"
	case "55":
		return "Dream"
	case "56":
		return "Southern Rock"
	case "57":
		return "Comedy"
	case "58":
		return "Cult"
	case "59":
		return "Gansta Rap"
	case "60":
		return "Top 40"
	case "61":
		return "Christian Rap"
	case "62":
		return "Pop/Funk"
	case "63":
		return "Jungle"
	case "64":
		return "Native American"
	case "65":
		return "Cabaret"
	case "66":
		return "New Wave"
	case "67":
		return "Psychedelic"
	case "68":
		return "Rave"
	case "69":
		return "Showtunes"
	case "70":
		return "Trailer"
	case "71":
		return "Lo-Fi"
	case "72":
		return "Tribal"
	case "73":
		return "Acid Punk"
	case "74":
		return "Acid Jazz"
	case "75":
		return "Pokla"
	case "76":
		return "Retro"
	case "77":
		return "Musical"
	case "78":
		return "Rock & Roll"
	case "79":
		return "Hard Rock"
	case "80":
		return "Folk"
	case "81":
		return "Folk-Rock"
	case "82":
		return "National Folk"
	case "83":
		return "Swing"
	case "84":
		return "Fast-Fusion"
	case "85":
		return "Bebop"
	case "86":
		return "Latin"
	case "87":
		return "Revival"
	case "88":
		return "Celtic"
	case "89":
		return "Bluegrass"
	case "90":
		return "Avantgarde"
	case "91":
		return "Gothic Rock"
	case "92":
		return "Progressive Rock"
	case "93":
		return "Psychedelic Rock"
	case "94":
		return "Symphonic Rock"
	case "95":
		return "Slow Rock"
	case "96":
		return "Big Band"
	case "97":
		return "Chorus"
	case "98":
		return "Easy Listening"
	case "99":
		return "Acoustic"
	case "100":
		return "Humor"
	case "101":
		return "Speech"
	case "102":
		return "Chanson"
	case "103":
		return "Opera"
	case "104":
		return "Chamber Music"
	case "105":
		return "Sonata"
	case "106":
		return "Symphony"
	case "107":
		return "Booty Bass"
	case "108":
		return "Primus"
	case "109":
		return "Porn Groove"
	case "110":
		return "Satire"
	case "111":
		return "Slow Jam"
	case "112":
		return "Club"
	case "113":
		return "Tango"
	case "114":
		return "Samba"
	case "115":
		return "Folklore"
	case "116":
		return "Ballad"
	case "117":
		return "Power Ballad"
	case "118":
		return "Rhythmic Soul"
	case "119":
		return "Freestyle"
	case "120":
		return "Duet"
	case "121":
		return "Punk Rock"
	case "122":
		return "Drum Solo"
	case "123":
		return "A Cappella"
	case "124":
		return "Euro-House"
	case "125":
		return "Dance Hall"
	case "126":
		return "Goa"
	case "127":
		return "Drum & Bass"
	case "128":
		return "Club-House"
	case "129":
		return "Hardcore"
	case "130":
		return "Terror"
	case "131":
		return "Indie"
	case "132":
		return "BritPop"
	case "133":
		return "Afro-Punk"
	case "134":
		return "Polsk Punk"
	case "135":
		return "Beat"
	case "136":
		return "Christian Gansta Rap"
	case "137":
		return "Heavy Metal"
	case "138":
		return "Black Metal"
	case "139":
		return "Crossover"
	case "140":
		return "Contemporary Christian"
	case "141":
		return "Christian Rock"
	case "142":
		return "Merengue"
	case "143":
		return "Salsa"
	case "144":
		return "Thrash Metal"
	case "145":
		return "Anime"
	case "146":
		return "JPop"
	case "147":
		return "Synthpop"
	case "148":
		return "Abstract"
	case "149":
		return "Art Rock"
	case "150":
		return "Baroque"
	case "151":
		return "Bhangra"
	case "152":
		return "Big Beat"
	case "153":
		return "Breakbeat"
	case "154":
		return "Chillout"
	case "155":
		return "Downtempo"
	case "156":
		return "Dub"
	case "157":
		return "EBM"
	case "158":
		return "Eclectic"
	case "159":
		return "Electro"
	case "160":
		return "Electroclash"
	case "161":
		return "Emo"
	case "162":
		return "Experimental"
	case "163":
		return "Garage"
	case "164":
		return "Global"
	case "165":
		return "IDM"
	case "166":
		return "Illbient"
	case "167":
		return "Industro-Goth"
	case "168":
		return "Jam Band"
	case "169":
		return "Krautrock"
	case "170":
		return "Leftfield"
	case "171":
		return "Lounge"
	case "172":
		return "Math Rock"
	case "173":
		return "New Romantic"
	case "174":
		return "Nu-Breakz"
	case "175":
		return "Post-Punk"
	case "176":
		return "Post-Rock"
	case "177":
		return "Psytrance"
	case "178":
		return "Shoegaze"
	case "179":
		return "Space Rock"
	case "180":
		return "Trop Rock"
	case "181":
		return "World Music"
	case "182":
		return "Neoclassical"
	case "183":
		return "Audiobook"
	case "184":
		return "Audio Theatre"
	case "185":
		return "Neue Deutsche Welle"
	case "186":
		return "Podcast"
	case "187":
		return "Indie Rock"
	case "188":
		return "G-Funk"
	case "189":
		return "Dubstep"
	case "190":
		return "Garage Rock"
	case "191":
		return "Psybient"
	default:
		return "Not Found"
	}
}

func traverse(path string, cleangenres []string, keepdirs []string) {
	var rmdirs []string
	err := godirwalk.Walk(path, &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.IsDir() {
				// check if we match any of our keep directories
				for _, d := range keepdirs {
					match, _ := regexp.MatchString(strings.ToLower(d), strings.ToLower(string(de.Name())))
					if match == true {
						if *D == true {
							fmt.Printf("KEEPDIR : %s/%s\n", path, string(de.Name()))
						}
						return nil
					}
				}
				// traverse through the directory and get the genre
				root := path + "/" + string(de.Name())
				var files []string
				err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
					if filepath.Ext(path) == ".mp3" {
						files = append(files, path)
					}
					return nil
				})
				if err != nil {
					panic(err)
				}
				if len(files) > 0 {
					genre := genre(files[0])
					// check if we are gave a number for the genre, if so convert it to text
					if strings.Contains(genre, "(") {
						if strings.Contains(genre, ")") {
							re := regexp.MustCompile(`\((.*?)\)`)
							numgenre := re.FindString(genre)
							numgenre = strings.Trim(numgenre, "(")
							numgenre = strings.Trim(numgenre, ")")
							genre = convgenre(numgenre)
						}
					}
					for _, g := range cleangenres {
						if strings.ToLower(g) == strings.ToLower(genre) {
							if *D == true {
								fmt.Printf("DELETE  : %s/%s %s\n", path, string(de.Name()), genre)
							}
							rmdir := path + "/" + string(de.Name())
							rmdirs = append(rmdirs, rmdir)
							return nil
						}
					}
					if *D == true {
						fmt.Printf("KEEP    : %s/%s %s\n", path, string(de.Name()), genre)
					}
				}
			}
			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// do the actual deletion of the files
	if *C == true {
		for _, rmdir := range rmdirs {
			err = os.RemoveAll(rmdir)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func main() {
	flag.Parse()
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	filename, _ := filepath.Abs(exPath + "/bob-id3clean.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	traverse(*P, config.Clean, config.Keepdir)
}
