package insight

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/stretchr/testify/assert"
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
func (m *MWS) Get() error {
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

	// * print each data value to console.
	// 		loop through lists
	// ! FIX LOOPING OF DATA MAP
	// ? json example: https://api.nasa.gov/insight_weather/?api_key=DEMO_KEY&feedtype=json&ver=1.0

	fields := reflect.TypeOf(contents)
	values := reflect.ValueOf(contents)
	for i := 0; i < len(contents.sols); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fmt.Print("Type:", field.Type, ",", field.Name, "=", value, "\n")

		var t assert.TestingT
		switch value.Kind() {
		case reflect.String:
			v := value.String()
			fmt.Print(v, "\n")
		case reflect.Int:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
		case reflect.Int32:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
		default:
			assert.Fail(t, "Not support type of struct")
		}
	}
	// what i'm used to doing
	// ! LF recursive func for maps in maps in maps etc..
	/* for i, v := range contents.sols {
		fmt.Printf("%b", v)
		if len(v[i]) {
			for x, y := range v[i] {
				fmt.Printf("%b", y)
				if y && y[x] {
					for b, c := range y[x] {
						fmt.Printf("%b", c)
					}
				}
			}
		}
	} */
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
