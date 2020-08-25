package insight

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// MWS ...
type MWS struct {
	Version  string
	Feedtype string
	Apikey   string
	HTTP     string
}

// New ...
func New(v string, ft string, ak string) *MWS {
	t := new(MWS)
	t.HTTP = "https://api.nasa.gov/insight_weather/"
	t.Version = v
	t.Feedtype = ft
	t.Apikey = ak
	return t
}

// POST ...
func (m *MWS) POST() string {
	if m.Version == "" || m.Feedtype == "" || m.Apikey == "" {
		fmt.Println("MWS.POST(): post variables not declared")
		return ""
	}
	r := m.HTTP + "?api_key=" + m.Apikey + "&feedtype=" + m.Feedtype + "&ver=" + m.Version
	return r
}

// Get ...
// get MWS data, pulled from the official InSight API
func (m *MWS) Get(savefetchedtofile bool, param ...string) error {
	// request json table from m.HTTP
	res, err := http.Get(m.POST())
	if err != nil {
		panic(err.Error())
	}

	// parse response to data
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var contents jsonContent
	json.Unmarshal([]byte(body), &contents)

	// @param download
	if savefetchedtofile == true {
		// create new file
		file, err := os.OpenFile(
			"test.txt",
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666,
		)
		if err != nil {
			panic(err.Error())
		}
		defer file.Close()

		// write json body to file
		byteSlice := body
		bytesWritten, err := file.Write(byteSlice)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Wrote %d bytes.\n", bytesWritten)
	}
	return nil
}

// jsonContent ...
// raw file, example: https://api.nasa.gov/insight_weather/?api_key=DEMO_KEY&feedtype=json&ver=1.0
type jsonContent struct {
	sols []SOL
}

// SOL ...
type SOL struct {
	index    int
	firstUTC string
	lastUTC  string
	season   string

	at         AT
	hws        HWS
	pre        PRE
	wd         []WD
	mostCommon WD
}

// AT ...
type AT struct {
	av float32
	ct int
	mn float32
	mx float32
}

// HWS ...
type HWS struct {
	av float32
	ct int
	mn float32
	mx float32
}

// PRE ...
type PRE struct {
	av float32
	ct int
	mn float32
	mx float32
}

// WD ...
type WD struct {
	index map[int]WDContent
}

// WDContent ...
type WDContent struct {
	CompassDegrees float32
	CompassPoint   string
	CompassRight   float32
	CompassUp      float32
	ct             int
}
