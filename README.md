# Storage Problem Study Case
ID finder tool
This code id developed in GO 
### Installation
```
https://github.com/GhadaM/StorageProblemStudyCase
```
### Usage
This tool have two endpoints:
```
http://localhost:8080/loadFile
```
This endpoint loads the initial *ids.csv*

```
http://localhost:8080/promotions/e0f30bca-8b8a-450f-928d-e407ac49e84b
```
This endpoint will find the id giving in the url 
### Logic behind
- The tool loads the *ids.csv. file 
- The tool loops through the lines and seprate them based on the month of the colomn *experation_date*
- The separation is done in Maps that are later written in corresponding csv
- When an id is in the Url , the tool reads through the csv files in the directories 
- The loops through the files at the same time 
- If the id is not found , a message is returned otherwise the record 
### Considerations 
- The separation into files based on the month helps the performance
- Using  Goroutines have huge performance benefits in peak periods
- Go's csv reader loads the file with each call to *loadFile*
- The following code(it is commented in the program) allows callt the function to load the csv file every 30 minutes
```
go func() {
		c := time.Tick(30 * time.Minute)
		for range c {
			GetInitialData()
		}
	}()
```
