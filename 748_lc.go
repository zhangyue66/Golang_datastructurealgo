import (
	//"fmt"
	"strings"
	"math"
    "unicode"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	dict := make(map[string]int)

	for i := 0; i < len(licensePlate); i++ {
        runeYZ := []rune(licensePlate)[i]
        if unicode.IsLetter(runeYZ) {
            att := strings.ToLower(string(runeYZ))
            dict[att] += 1
        }
	}

	res := []string{}
	
	for _,word := range words{
		match := 1
		for letter,cnt := range dict {
            if strings.Count(word,letter) < cnt {
                match = 0;
				break;
            }
				
		}
		if match == 0 {
			match =1;
			continue;
		} else {
			res = append(res,word)
		}
			
	}
	miw := ""
    //fmt.Println(res)
    //fmt.Println(dict)
	if len(res) == 1 {
        
		return res[0]
	} else {
		minLen := math.Inf(1)

        for _,ans := range res {
			minLen = math.Min(minLen,float64(len(ans)))
		}
		
        for _,ans := range res {
			if len(ans) == int(minLen) {
                
				miw = ans;
				break;
			}
		}
	}
	return miw
}
