# Veracode SCA Extracter

## Description
Creates a CSV file with SCA findings for all build IDs in the input file.

## Parameters
1.  **-credsFile**: REQUIRED. Credentials file with Veracode API ID/Key.
2. **-builds**: REQUIRED. Text file with list of build IDs on separate lines (recommended to grab from Applications Extract).
3. **-outputFileName**: OPTIONAL. Specify the name of the file for the CSV. The default is `SCA_20170101_150405.csv`, which includes a timestamp so a new file is created with each run. Specifying this parameter will overwrite the existing file with each run.

## Credentials File
Must be structured like the following:
```
[DEFAULT]
veracode_api_key_id = ID HERE
veracode_api_key_secret = SECRET HERE
```

## Executables
Executables for Windows, Mac, and Linux will be available in the releases section of the repository (https://github.com/brian1917/vcodeSCAExtractor/releases)
* For Windows, users download the EXE and from the command line run `vcodexsv.exe --help`.
* For Mac, download the executable, set it to be an executable: `chmod +x vcodecsv` and run `./vcodecsv --help`

## Third-party Packages
1. github.com/brian1917/vcodeapi
