package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brian1917/vcodeapi"
)

//DECLARE VARIABLES
var credsFile, buildListFile, outputFileName string
var resultsFile *os.File
var sca vcodeapi.SoftwareCompositionAnalysis
var vulns []vcodeapi.Vulnerability
var errorCheck, err error

//GET USER INPUT
func init() {
	flag.StringVar(&credsFile, "credsFile", "", "Credentials file path")
	flag.StringVar(&buildListFile, "builds", "", "Location of the text file with builds on each line")
	flag.StringVar(&outputFileName, "outputFileName", "default", "Specific the name of the output file. Default is based on a timestamp. Providing this parameter will overwrite the file each run.")
}

func main() {

	start := time.Now()

	// PARSE FLAGS
	flag.Parse()

	// NAME THE OUTPUT FILE
	if outputFileName == "default" {
		outputFileName = "SCA_" + time.Now().Format("20060102_150405") + ".csv"
	}

	// CREATE A CSV FILE FOR RESULTS
	if resultsFile, err = os.Create(outputFileName); err != nil {
		log.Fatal(err)
	}
	defer resultsFile.Close()

	// CREATE THE WRITER
	writer := csv.NewWriter(resultsFile)
	defer writer.Flush()

	// WRITE CSV HEADERS
	headers := []string{"app_name", "build_id", "published_date", "component", "version", "cve_id", "severity", "cvss_score", "cwe_id", "affects_policy_compliance", "file_paths"}
	if err = writer.Write(headers); err != nil {
		log.Fatal(err)
	}

	// READ IN TEXT FILE OF BUILDS
	builds := getBuildList(buildListFile)

	// CYCLE THROUGH EACH BUILD AND GET THE SCA DATA
	buildCounter := 0

	for _, build := range builds {
		buildCounter++

		fmt.Printf("Processing Build ID %s (%v of %v)\n", build, buildCounter, len(builds))
		sca, err = vcodeapi.ParseSCAReport(credsFile, build)
		if err != nil {
			log.Fatal(err)
		}

		// CYCLE THROUGH COMPONENTS
		for _, vulnComponent := range sca.VulnerableComponents.Component {
			filePathCount := len(vulnComponent.FilePaths.FilePath)

			// CYCLE THROUGH VULNERABILITIES
			for _, vuln := range vulnComponent.Vulnerabilities.Vulnerability {

				// WRITE TO CSV
				entry := []string{sca.AppName, build, sca.PublishedDate, vulnComponent.FileName, vulnComponent.Version, vuln.CveID, vuln.Severity, vuln.CvssScore, vuln.CweID, vuln.VulnerabilityAffectsPolicyCompliance, strconv.Itoa(filePathCount)}
				if err = writer.Write(entry); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	fmt.Printf("Run time: %v \n", time.Since(start))
}
