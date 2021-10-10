package main


import (
	"fmt"
	"io/ioutil"
    "log"
	"regexp"
	"os"
	"strings"
	"strconv"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/Open-Science-Global/poly/io/genbank"
	"github.com/Open-Science-Global/poly/finder"
	"github.com/Open-Science-Global/poly/transform"
	"github.com/Open-Science-Global/poly"
	"github.com/sethvargo/go-githubactions"
)

func main() {
  Execute()
}

var (
	input string
	output string
	pattern string
)


var rootCmd = &cobra.Command{
	Use:   "annotator",
	Short: "Annotator is github action to annotate problematic sequences from given Genbank file.",
	Long: `Annotator is github action to annotate problematic sequences from given Genbank file.`,
	Run: func(cmd *cobra.Command, args []string) {
		Script(input, output, pattern)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Directory where all the input genbank files will be read")
	rootCmd.PersistentFlags().StringVarP(&output, "ouput", "o", "", "Directory where all the output genbank files wil be written")
	rootCmd.PersistentFlags().StringVarP(&pattern, "pattern", "p", "", "Regex to selective filter specific files in the input folder")


	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("ouput")
	rootCmd.MarkFlagRequired("pattern")	
}

func Script(inputDir string, outputDir string, pattern string) {
	filesPath := getListFilesByPattern(inputDir, pattern)

	for _, filePath := range filesPath {
		sequence := genbank.Read(filePath)
		annotator(sequence, filePath, outputDir)
	} 
}

func annotator(sequence poly.Sequence, filePath string, outputDir string) {
	outputFile := filepath.Base(filePath)
	outputPath := outputDir + "/" + outputFile
	
	annotatedSequence := findProblematicSequences(sequence)
	str := fmt.Sprintf("%v", annotatedSequence)
	githubactions.Infof(str)
	genbank.Write(annotatedSequence, outputPath)
}

func getListFilesByPattern(inputDir string, pattern string) []string {
	files, err := ioutil.ReadDir(inputDir)
	
	if err != nil {
        log.Fatal(err)
    }
	var filesPath []string
    for _, f := range files {
		var validFile = regexp.MustCompile(pattern)
		if validFile.MatchString(f.Name()) {
			file := inputDir + "/" + f.Name()
			filesPath = append(filesPath, file)
		}
    }
	return filesPath
}

func findProblematicSequences(sequence poly.Sequence) poly.Sequence{
	var functions []func(string) []finder.Match

	functions = append(functions, finder.ForbiddenSequence(restrictionBindingSitesListFindProblems()))
	functions = append(functions, finder.ForbiddenSequence(homologySequencesFindProblems()))
	functions = append(functions, finder.RemoveRepeat(10))
	functions = append(functions, AvoidHairpin(20, 200))

	problems := finder.Find(strings.ToUpper(sequence.Sequence), functions)

	sequence = finder.AddMatchesToSequence(problems, sequence)
	return sequence
}

func AvoidHairpin(stemSize int, hairpinWindow int) func(string) []finder.Match {
	return func(sequence string) []finder.Match {
		var matches []finder.Match
		reverse := transform.ReverseComplement(sequence)
		for i := 0; i < len(sequence)-stemSize && len(sequence)-(i+hairpinWindow) >= 0; i++ {
			word := sequence[i : i+stemSize]
			rest := reverse[len(sequence)-(i+hairpinWindow) : len(sequence)-(i+stemSize)]
			if strings.Contains(rest, word) {
				location := strings.Index(rest, word)
				matches = append(matches, finder.Match{i, i + hairpinWindow - location - 1, "Harpin found in next " + strconv.Itoa(hairpinWindow) + "bp in reverse complementary sequence: " + word})
			}
		}
		return matches
	}
}

func homologySequencesFindProblems() []string {
	// I don't have to worry about TTTTTT and GGGGGG because I already try to find also by reverse complementary of each sequence in finder
	return []string{"AAAAAA", "CCCCCC"}
}

func restrictionBindingSitesListFindProblems() []string {
	BsaI_bind_5prime := "GGTCTC" //5" GGTCTC N|      3"
	//3" CCAGAG N NNNN| 5"

	BbsI_bind_5prime := "GAAGAC" //5" GAAGAC NN|
	//3" CTTCTG NN NNNN|

	BtgzI_bind_5prime := "GCGATG" //5" GCGATG NNN NNN NNN N|      3"
	//3" CGCTAC NNN NNN NNN N NNNN| 5"

	SapI_bind_5prime := "GCTCTTC" //5" GCTCTTC N|     3"
	//3" CGAGAAG N NNN| 5"

	BsmbI_bind_5prime := "CGTCTC" //5" CGTCTC N|      3"
	//3" GCAGAG N NNNN| 5"

	AarI_bind_5prime := "CACCTGC" //5" CACCTGC N NNN|       3"
	//3" GTGGACG N NNN NNNN|  5"

	PmeI_bind := "GTTTAAAC"

	HindIII := "AAGCTT"
	PstI := "CTGCAG"
	XbaI := "TCTAGA"
	BamHI := "GGATCC"
	SmaI := "CCCGGG"
	KpnI := "GGTACC"
	SacI := "GAGCTC"
	SalI := "GTCGAC"
	EcoRI := "GAATTC"
	SphI := "GCATGC"
	AvrII := "CCTAGG"
	SwaI := "ATTTAAAT"
	AscI := "GGCGCGCC"
	FseI := "GGCCGGCC"
	PacI := "TTAATTAA"
	SpeI := "ACTAGT"
	NotI := "GCGGCCGC"
	SanDI_A := "GGGACCC"
	SanDI_T := "GGGTCCC"
	BglII := "AGATCT"
	XhoI := "CTCGAG"
	ClaI := "ATCGAT"

	forbiddenSeqList := []string{
		BsaI_bind_5prime,
		BbsI_bind_5prime,
		SapI_bind_5prime,
		BsmbI_bind_5prime,
		BtgzI_bind_5prime,
		AarI_bind_5prime,
		PmeI_bind,
		HindIII,
		PstI,
		XbaI,
		BamHI,
		SmaI,
		KpnI,
		SacI,
		SalI,
		EcoRI,
		SphI,
		AvrII,
		SacI,
		SalI,
		SwaI,
		AscI,
		FseI,
		PacI,
		SpeI,
		NotI,
		SanDI_A,
		SanDI_T,
		BglII,
		XhoI,
		ClaI,
		"AAAAAA",
		"CCCCCC",
		"TTTTTT",
		"GGGGGG",
	}

	return forbiddenSeqList
}

